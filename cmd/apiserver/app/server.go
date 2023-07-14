package app

import (
	"cangbinggu.io/buliangren/cmd/apiserver/app/config"
	"cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/apiserver"
	updateconfigprovider "cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/provider/updateconfig"
	"fmt"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

func CreateServerChain(cfg *config.Config) (*genericapiserver.GenericAPIServer, error) {
	apiServerConfig := createAPIServerConfig(cfg)
	apiServer, err := CreateAPIServer(apiServerConfig, genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}

	if err := registerHandler(apiServer); err != nil {
		return nil, err
	}

	apiServer.GenericAPIServer.AddPostStartHookOrDie("strat tiankuixing informers", func(context genericapiserver.PostStartHookContext) error {
		cfg.VersionedSharedInformerFactory.Start(context.StopCh)
		return nil
	})

	return apiServer.GenericAPIServer, nil
}

func CreateAPIServer(apiServerConfig *apiserver.Config, delegateAPIServer genericapiserver.DelegationTarget) (*apiserver.APIServer, error) {
	return apiServerConfig.Complete().New(delegateAPIServer)
}

func createAPIServerConfig(cfg *config.Config) *apiserver.Config {
	return &apiserver.Config{
		GenericConfig: &genericapiserver.RecommendedConfig{
			Config: *cfg.GenericAPIServerConfig,
		},
		ExtraConfig: apiserver.ExtraConfig{
			ServerName:              cfg.ServerName,
			APIResourceConfigSource: cfg.StorageFactory.APIResourceConfigSource,
			StorageFactory:          cfg.StorageFactory,
			VersionedInformers:      cfg.VersionedSharedInformerFactory,
		},
	}
}

func createFilterChain(apiServer *genericapiserver.GenericAPIServer) {

}

func registerHandler(apiServer *apiserver.APIServer) error {
	createFilterChain(apiServer.GenericAPIServer)

	updateconfigprovider.RegisterHandler(apiServer.GenericAPIServer.Handler.NonGoRestfulMux)

	fmt.Printf("All of http handlers registerd,paths: %v", apiServer.GenericAPIServer.Handler.ListedPaths())
	return nil
}
