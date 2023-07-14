package apiserver

import (
	tiankuixingv1 "cangbinggu.io/buliangren/pkg/apis/tiankuixing/v1"
	versionedinformers "cangbinggu.io/buliangren/pkg/generated/informers/externalversions"
	corerest "cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/registry/core/rest"
	tiankuixingrest "cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/registry/rest"
	"cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/storage"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apiserver/pkg/registry/generic"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
)

type ExtraConfig struct {
	ServerName              string
	APIResourceConfigSource serverstorage.APIResourceConfigSource
	StorageFactory          serverstorage.StorageFactory
	VersionedInformers      versionedinformers.SharedInformerFactory
}

type Config struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}

type CompletedConfig struct {
	*completedConfig
}

type APIServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

func (cfg *Config) Complete() CompletedConfig {
	c := completedConfig{
		cfg.GenericConfig.Complete(),
		&cfg.ExtraConfig,
	}

	return CompletedConfig{&c}
}

func (c completedConfig) New(delegationTarget genericapiserver.DelegationTarget) (*APIServer, error) {
	//初始化，创建go-restful的Container，初始化apiServerHandler，API Server预先注册了一些默认path
	s, err := c.GenericConfig.New(c.ExtraConfig.ServerName, delegationTarget)
	if err != nil {
		return nil, err
	}

	m := &APIServer{
		GenericAPIServer: s,
	}

	if c.ExtraConfig.APIResourceConfigSource.VersionEnabled(corev1.SchemeGroupVersion) {
		legacyRESTStorageProvider := corerest.LegacyRESTStorageProvider{}
		m.InstallLegacyAPI(&c, c.GenericConfig.RESTOptionsGetter, legacyRESTStorageProvider)
	}

	// /apis开头版本的api注册到Container中
	restStorageProviders := []storage.RESTStorageProvider{
		&tiankuixingrest.StorageProvider{
			LoopbackClientConfig: c.GenericConfig.LoopbackClientConfig,
		},
	}

	m.InstallAPIs(c.ExtraConfig.APIResourceConfigSource, c.GenericConfig.RESTOptionsGetter, restStorageProviders...)

	return m, err
}

func (m *APIServer) InstallLegacyAPI(c *completedConfig, restOptionsGetter generic.RESTOptionsGetter, legacyRESTStorageProvider corerest.LegacyRESTStorageProvider) {
	apiGroupInfo, err := legacyRESTStorageProvider.NewLegacyRESTStorage(restOptionsGetter, c.GenericConfig.LoopbackClientConfig)
	if err != nil {
		fmt.Printf("Error building core storage: %v", err)
	}

	if err := m.GenericAPIServer.InstallLegacyAPIGroup(genericapiserver.DefaultLegacyAPIPrefix, apiGroupInfo); err != nil {
		fmt.Printf("Error in registering group versions: %v", err)
	}
}

func (m *APIServer) InstallAPIs(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter, restStorageProviders ...storage.RESTStorageProvider) {
	var apiGroupsInfo []genericapiserver.APIGroupInfo

	for _, restStorageBuilder := range restStorageProviders {
		groupName := restStorageBuilder.GroupName()
		if !apiResourceConfigSource.AnyVersionForGroupEnabled(groupName) {
			fmt.Printf("Skipping disabled API group %q", groupName)
			continue
		}
		fmt.Println("jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")
		apiGroupInfo, enabled := restStorageBuilder.NewRESTStorage(apiResourceConfigSource, restOptionsGetter)
		fmt.Printf("apiGroupInfo=%v,enabled=%v", apiGroupsInfo, enabled)
		fmt.Println("kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk")
		if !enabled {
			fmt.Printf("Problem initializing API group %q,skipping.", groupName)
			continue
		}
		fmt.Printf("Enabled API group %q.", groupName)
		if postHookProvider, ok := restStorageBuilder.(genericapiserver.PostStartHookProvider); ok {
			name, hook, err := postHookProvider.PostStartHook()
			if err != nil {
				fmt.Printf("Error building PostStartHook: %v", err)
			}
			m.GenericAPIServer.AddPostStartHookOrDie(name, hook)
		}
		apiGroupsInfo = append(apiGroupsInfo, apiGroupInfo)
	}

	for i := range apiGroupsInfo {
		fmt.Printf("22222222222222222222222=============================")
		fmt.Printf("apiGroupsInfo[i] ============================= %v", &apiGroupsInfo[i])
		fmt.Printf("1111111111111111111111111111=============================")
		if err := m.GenericAPIServer.InstallAPIGroup(&apiGroupsInfo[i]); err != nil {
			fmt.Printf("3333333333333333333333333333=============================")
			fmt.Printf("Error in registering group versions: %v", err)
		}
	}
}

func DefaultAPIResourceConfigSource() *serverstorage.ResourceConfig {
	ret := serverstorage.NewResourceConfig()
	ret.EnableVersions(
		tiankuixingv1.SchemeGroupVersion,
	)
	return ret
}
