package storage

import (
	storageoptions "cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/storage/options"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/server/options/encryptionconfig"
	"k8s.io/apiserver/pkg/server/resourceconfig"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	"strings"
)

var SpecialDefaultResourcePrefixes = map[schema.GroupResource]string{}

func NewFactoryConfig(codecs runtime.StorageSerializer, scheme *runtime.Scheme) *FactoryConfig {
	var resources []schema.GroupVersionResource
	return &FactoryConfig{
		Serializer:                codecs,
		DefaultResourceEncoding:   serverstorage.NewDefaultResourceEncodingConfig(scheme),
		ResourceEncodingOverrides: resources,
	}
}

type FactoryConfig struct {
	StorageConfig                    storagebackend.Config
	APIResourceConfig                *serverstorage.ResourceConfig
	DefaultResourceEncoding          *serverstorage.DefaultResourceEncodingConfig
	DefaultStorageMediaType          string
	Serializer                       runtime.StorageSerializer
	ResourceEncodingOverrides        []schema.GroupVersionResource
	ETCDServersOverrides             []string
	EncryptionProviderConfigFilePath string
}

func (c *FactoryConfig) Complete(etcdOptions *storageoptions.ETCDStorageOptions) (*CompletedFactoryConfig, error) {
	c.StorageConfig = storagebackend.Config{
		Type:   storagebackend.StorageTypeETCD3,
		Prefix: etcdOptions.Prefix,
		Transport: storagebackend.TransportConfig{
			ServerList:    etcdOptions.ServerList,
			KeyFile:       etcdOptions.KeyFile,
			CertFile:      etcdOptions.CertFile,
			TrustedCAFile: etcdOptions.CAFile,
		},
		Paging:                etcdOptions.Paging,
		Codec:                 etcdOptions.Codec,
		EncodeVersioner:       etcdOptions.EncodeVersioner,
		Transformer:           etcdOptions.Transformer,
		CompactionInterval:    etcdOptions.CompactionInterval,
		CountMetricPollPeriod: etcdOptions.CountMetricPollPeriod,
	}
	c.DefaultStorageMediaType = etcdOptions.DefaultStorageMediaType
	c.ETCDServersOverrides = etcdOptions.ETCDServersOverrides
	c.EncryptionProviderConfigFilePath = etcdOptions.EncryptionProviderConfigFilePath
	return &CompletedFactoryConfig{c}, nil
}

type CompletedFactoryConfig struct {
	*FactoryConfig
}

func (c *CompletedFactoryConfig) New() (*serverstorage.DefaultStorageFactory, error) {
	resourceEncodingConfig := resourceconfig.MergeResourceEncodingConfigs(c.DefaultResourceEncoding, c.ResourceEncodingOverrides)
	storageFactory := serverstorage.NewDefaultStorageFactory(
		c.StorageConfig,
		c.DefaultStorageMediaType,
		c.Serializer,
		resourceEncodingConfig,
		c.APIResourceConfig,
		SpecialDefaultResourcePrefixes)

	for _, override := range c.ETCDServersOverrides {
		tokens := strings.Split(override, "#")
		apiresource := strings.Split(tokens[0], "/")

		group := apiresource[0]
		resource := apiresource[1]
		groupResource := schema.GroupResource{Group: group, Resource: resource}

		servers := strings.Split(tokens[1], ";")
		storageFactory.SetEtcdLocation(groupResource, servers)
	}
	if len(c.EncryptionProviderConfigFilePath) != 0 {
		transformerOverrides, err := encryptionconfig.GetTransformerOverrides(c.EncryptionProviderConfigFilePath)
		if err != nil {
			return nil, err
		}
		for groupResource, transformer := range transformerOverrides {
			storageFactory.SetTransformer(groupResource, transformer)
		}
	}
	return storageFactory, nil
}
