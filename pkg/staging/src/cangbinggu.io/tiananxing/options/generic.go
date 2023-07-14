package options

import (
	"fmt"
	netutil "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/sets"
	genericapiserveroptions "k8s.io/apiserver/pkg/server/options"
	"net"
	"os"
)

type GenericOptions struct {
	*genericapiserveroptions.ServerRunOptions
	ExternalPort   int
	ExternalScheme string
	ExternalCAFile string
}

func NewGenericOptions() *GenericOptions {
	return &GenericOptions{
		ServerRunOptions: genericapiserveroptions.NewServerRunOptions(),
		ExternalScheme:   "https",
	}
}

func (o *GenericOptions) DefaultAdvertiseAddress(secure *SecureServingOptions) error {
	if o.AdvertiseAddress == nil || o.AdvertiseAddress.IsUnspecified() {
		hostIP, err := netutil.ResolveBindAddress(secure.BindAddress)
		if err != nil {
			return fmt.Errorf("unable to find suitable network address.error='%v'. Try to set the AdvertiseAddress directly or provide a valid BindAddress to fix this", err)
		}
		o.AdvertiseAddress = hostIP
	}
	return nil
}

func CompleteGenericAndSecureOptions(genericOpts *GenericOptions, secureServingOpts *SecureServingOptions) error {
	// set defaults
	if err := genericOpts.DefaultAdvertiseAddress(secureServingOpts); err != nil {
		return err
	}
	if err := secureServingOpts.MaybeDefaultWithSelfSignedCerts(genericOpts.AdvertiseAddress.String(), []string{"localhost", "localhost.localdomain"}, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	if len(genericOpts.ExternalHost) == 0 {
		if len(genericOpts.AdvertiseAddress) > 0 {
			genericOpts.ExternalHost = genericOpts.AdvertiseAddress.String()
		} else {
			if hostname, err := os.Hostname(); err == nil {
				genericOpts.ExternalHost = hostname
			} else {
				return fmt.Errorf("error finding host name: %v", err)
			}
		}
		fmt.Printf("External host was not specified, using %v", genericOpts.ExternalHost)
	}

	if genericOpts.ExternalPort == 0 {
		genericOpts.ExternalPort = secureServingOpts.BindPort
		fmt.Printf("External port was not specified, using binding port %d", secureServingOpts.BindPort)
	}

	if genericOpts.ExternalScheme == "" {
		genericOpts.ExternalScheme = "https"
		fmt.Printf("External scheme was not specified, using default scheme `HTTPS`")
	} else {
		schemes := sets.NewString("http", "https")
		if !schemes.Has(genericOpts.ExternalScheme) {
			return fmt.Errorf("error matching external scheme: %s, must be http or https", genericOpts.ExternalScheme)
		}
	}

	if genericOpts.ExternalScheme == "http" {
		if genericOpts.ExternalCAFile != "" {
			return fmt.Errorf("cannot set CA file when external exposure is HTTP")
		}
	} else {
		if genericOpts.ExternalCAFile == "" {
			fmt.Printf("External CA file was not specified, using server certificate file: %s", secureServingOpts.ServerCert.CertKey.CertFile)
			genericOpts.ExternalCAFile = secureServingOpts.ServerCert.CertKey.CertFile
		}
	}

	return nil
}
