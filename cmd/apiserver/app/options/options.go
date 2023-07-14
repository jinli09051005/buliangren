package options

import (
	cangbingguoptions "cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/options"
	storageoptions "cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/storage/options"
	"github.com/spf13/pflag"
	genericapiserveroptions "k8s.io/apiserver/pkg/server/options"
)

// AddFlags adds flags for a specific server to the specified FlagSet object.
func (o *Options) AddFlags(fs *pflag.FlagSet) {
	o.SecureServing.AddFlags(fs)
	o.ETCD.AddFlags(fs)
}

// ApplyFlags parsing parameters from the command line or configuration file
// to the options instance.
func (o *Options) ApplyFlags() []error {
	var errs []error
	return errs
}

type Options struct {
	SecureServing *cangbingguoptions.SecureServingOptions
	Audit         *genericapiserveroptions.AuditOptions
	ETCD          *storageoptions.ETCDStorageOptions
	Generic       *cangbingguoptions.GenericOptions
}

func NewOptions(serverName string) *Options {
	return &Options{
		SecureServing: cangbingguoptions.NewSecureServingOptions(serverName, 443),
		Audit:         genericapiserveroptions.NewAuditOptions(),
		ETCD:          storageoptions.NewETCDStorageOptions("tiankuixing/updateconfig"),
		Generic:       cangbingguoptions.NewGenericOptions(),
	}
}

func (o *Options) Complete() error {
	if err := cangbingguoptions.CompleteGenericAndSecureOptions(o.Generic, o.SecureServing); err != nil {
		return err
	}
	return nil
}
