package rest

import (
	"cangbinggu.io/buliangren/pkg/apis/tiankuixing"
	tiankuixingv1 "cangbinggu.io/buliangren/pkg/apis/tiankuixing/v1"
	tiankuixinginternalclient "cangbinggu.io/buliangren/pkg/generated/clientset/internalversion/typed/tiankuixing/internalversion"
	updateconfigstorage "cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/registry/updateconfig/storage"
	"cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/storage"
	"fmt"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	restclient "k8s.io/client-go/rest"
)

type StorageProvider struct {
	LoopbackClientConfig *restclient.Config
}

var _ storage.RESTStorageProvider = &StorageProvider{}

func (s *StorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool) {
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(tiankuixing.Group, tiankuixing.Scheme, tiankuixing.ParameterCodec, tiankuixing.Codecs)

	if storageMap, err := s.v1Storage(apiResourceConfigSource, restOptionsGetter, s.LoopbackClientConfig); err != nil {
		return genericapiserver.APIGroupInfo{}, false
	} else if len(storageMap) > 0 {
		fmt.Println("888888888888888888888888")
		apiGroupInfo.VersionedResourcesStorageMap[tiankuixingv1.SchemeGroupVersion.Version] = storageMap
		fmt.Printf("888888888888888888888888:%s", storageMap)
	}

	return apiGroupInfo, true
}

func (s *StorageProvider) GroupName() string {
	return tiankuixing.Group
}

func (s *StorageProvider) v1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter, loopbackClientConfig *restclient.Config) (map[string]rest.Storage, error) {
	client := tiankuixinginternalclient.NewForConfigOrDie(loopbackClientConfig)
	strorageMap := make(map[string]rest.Storage)

	tiankuixingREST, err := updateconfigstorage.NewStorage(restOptionsGetter, client)
	if err != nil {
		return strorageMap, err
	}
	strorageMap["updateconfigs"] = tiankuixingREST.UpdateConfig

	return strorageMap, nil
}
