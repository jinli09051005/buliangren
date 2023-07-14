package rest

import (
	"cangbinggu.io/buliangren/pkg/apis/tiankuixing"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	restclient "k8s.io/client-go/rest"
)

type LegacyRESTStorageProvider struct {
}

func (c LegacyRESTStorageProvider) NewLegacyRESTStorage(restOptionsGetter generic.RESTOptionsGetter, loopbackClientConfig *restclient.Config) (*genericapiserver.APIGroupInfo, error) {
	apiGroupInfo := &genericapiserver.APIGroupInfo{
		PrioritizedVersions:          tiankuixing.Scheme.PrioritizedVersionsForGroup(""),
		VersionedResourcesStorageMap: map[string]map[string]rest.Storage{},
		Scheme:                       tiankuixing.Scheme,
		ParameterCodec:               tiankuixing.ParameterCodec,
		NegotiatedSerializer:         tiankuixing.Codecs,
	}

	//tiananxingCLient := tiankuixinginternalclient.NewForConfigOrDie(loopbackClientConfig)

	restStorageMap := map[string]rest.Storage{}

	apiGroupInfo.VersionedResourcesStorageMap["v1"] = restStorageMap
	return apiGroupInfo, nil
}
