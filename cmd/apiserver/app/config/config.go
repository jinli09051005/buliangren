package config

import (
	"cangbinggu.io/buliangren/cmd/apiserver/app/options"
	"cangbinggu.io/buliangren/pkg/apis/tiankuixing"
	versionedclientset "cangbinggu.io/buliangren/pkg/generated/clientset/versioned"
	versionedinformers "cangbinggu.io/buliangren/pkg/generated/informers/externalversions"
	generatedopenapi "cangbinggu.io/buliangren/pkg/generated/openapi"
	"cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/apiserver"
	"cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/openapi"
	"cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/storage"
	"fmt"
	"k8s.io/apimachinery/pkg/util/sets"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/filters"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"time"
)

const (
	title   = "Cangbinggu Buliangren Engine Tiankuixing API"
	license = "Apache 2.0"
)

type Config struct {
	ServerName                     string
	GenericAPIServerConfig         *genericapiserver.Config
	VersionedSharedInformerFactory versionedinformers.SharedInformerFactory
	StorageFactory                 *serverstorage.DefaultStorageFactory
}

func CreateConfigFromOptions(serverName string, opts *options.Options) (*Config, error) {
	genericAPIServerConfig := genericapiserver.NewConfig(tiankuixing.Codecs)
	genericAPIServerConfig.LongRunningFunc = filters.BasicLongRunningRequestCheck(
		sets.NewString("watch", "proxy"),
		sets.NewString("attach", "exec", "proxy", "log", "portforward"),
	)

	genericAPIServerConfig.MergedResourceConfig = apiserver.DefaultAPIResourceConfigSource()
	//fmt.Printf("gggggggggggggggggggggg======================GroupVersionConfigs:%v", genericAPIServerConfig.MergedResourceConfig.GroupVersionConfigs)
	//fmt.Printf("gggggggggggggggggggggg======================ResourceConfigs:%v", genericAPIServerConfig.MergedResourceConfig.ResourceConfigs)
	genericAPIServerConfig.EnableIndex = false
	genericAPIServerConfig.EnableProfiling = false

	if err := opts.SecureServing.ApplyTo(&genericAPIServerConfig.SecureServing, &genericAPIServerConfig.LoopbackClientConfig); err != nil {
		return nil, err
	}
	// openAPI
	openapi.SetupOpenAPI(genericAPIServerConfig, generatedopenapi.GetOpenAPIDefinitions, title, license)
	// storageFactory
	storageFactoryConfig := storage.NewFactoryConfig(tiankuixing.Codecs, tiankuixing.Scheme)
	storageFactoryConfig.APIResourceConfig = genericAPIServerConfig.MergedResourceConfig
	opts.ETCD.ServerList = []string{"http://etcd-svc:2379"}
	//opts.SecureServing.ServerCert.CertDirectory = "/apiserver.local.config/certificates"
	opts.SecureServing.ServerCert.CertKey.CertFile = "/apiserver.local.config/certificates/tls.crt"
	opts.SecureServing.ServerCert.CertKey.KeyFile = "/apiserver.local.config/certificates/tls.key"
	opts.Generic.ExternalPort = 443
	completedStorageFactoryConfig, err := storageFactoryConfig.Complete(opts.ETCD)
	if err != nil {
		return nil, err
	}
	storageFactory, err := completedStorageFactoryConfig.New()
	if err != nil {
		return nil, err
	}
	if err := opts.ETCD.ApplyWithStorageFactoryTo(storageFactory, genericAPIServerConfig); err != nil {
		return nil, err
	}
	genericAPIServerConfig.LoopbackClientConfig.ContentConfig.ContentType = "application/vnd.kubernetes.protobuf"
	kubeClientConfig := genericAPIServerConfig.LoopbackClientConfig
	clientGoExternalClient, err := versionedclientset.NewForConfig(kubeClientConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create real external clientset: %v", err)
	}

	versionedInformers := versionedinformers.NewSharedInformerFactory(clientGoExternalClient, 10*time.Minute)

	return &Config{
		ServerName:                     serverName,
		GenericAPIServerConfig:         genericAPIServerConfig,
		VersionedSharedInformerFactory: versionedInformers,
		StorageFactory:                 storageFactory,
	}, nil
}
