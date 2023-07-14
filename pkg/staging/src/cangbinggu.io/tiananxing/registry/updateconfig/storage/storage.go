package storage

import (
	"cangbinggu.io/buliangren/pkg/apis/tiankuixing"
	tiankuixinginternalclient "cangbinggu.io/buliangren/pkg/generated/clientset/internalversion/typed/tiankuixing/internalversion"
	updateconfigstrategy "cangbinggu.io/buliangren/pkg/staging/src/cangbinggu.io/tiananxing/registry/updateconfig"
	"context"
	"go.uber.org/zap"
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	genericregistry "k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"log"
)

type TiankuixingStorage struct {
	UpdateConfig *REST
}

func NewStorage(optsGetter genericregistry.RESTOptionsGetter, tiankuixingclient tiankuixinginternalclient.TiankuixingInterface) (*TiankuixingStorage, error) {
	strategy := updateconfigstrategy.NewStrategy(tiankuixingclient)
	store := &registry.Store{
		NewFunc: func() runtime.Object {
			return &tiankuixing.UpdateConfig{}
		},
		NewListFunc: func() runtime.Object {
			return &tiankuixing.UpdateConfigList{}
		},
		DefaultQualifiedResource: tiankuixing.Resource("updateconfigs"),
		ReturnDeletedObject:      true,

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,
	}
	store.TableConvertor = rest.NewDefaultTableConvertor(store.DefaultQualifiedResource)
	options := &genericregistry.StoreOptions{
		RESTOptions: optsGetter,
		AttrFunc:    updateconfigstrategy.GetAttrs,
	}

	if err := store.CompleteWithOptions(options); err != nil {
		log.Panic("Failed to create updateconfig etcd rest storage", zap.Error(err))
		return nil, err
	}

	return &TiankuixingStorage{
		UpdateConfig: &REST{
			store,
		},
	}, nil
}

type REST struct {
	*registry.Store
}

var _ rest.ShortNamesProvider = &REST{}

func (r *REST) ShortNames() []string {
	return []string{"upcs"}
}

func (r *REST) List(ctx context.Context, options *metainternal.ListOptions) (runtime.Object, error) {
	return r.Store.List(ctx, options)
}

func (r *REST) DeleteCollection(ctx context.Context, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions, listOptions *metainternal.ListOptions) (runtime.Object, error) {
	return r.Store.DeleteCollection(ctx, deleteValidation, options, listOptions)
}

func (r *REST) Get(ctx context.Context, updateconfigName string, options *metav1.GetOptions) (runtime.Object, error) {
	obj, err := r.Store.Get(ctx, updateconfigName, options)
	if err != nil {
		return nil, err
	}
	updateconfig := obj.(*tiankuixing.UpdateConfig)
	return updateconfig, err
}
