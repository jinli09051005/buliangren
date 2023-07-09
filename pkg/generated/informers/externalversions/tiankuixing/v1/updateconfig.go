/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	tiankuixingv1 "cangbinggu.io/buliangren/pkg/apis/tiankuixing/v1"
	versioned "cangbinggu.io/buliangren/pkg/generated/clientset/versioned"
	internalinterfaces "cangbinggu.io/buliangren/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "cangbinggu.io/buliangren/pkg/generated/listers/tiankuixing/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// UpdateConfigInformer provides access to a shared informer and lister for
// UpdateConfigs.
type UpdateConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.UpdateConfigLister
}

type updateConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewUpdateConfigInformer constructs a new informer for UpdateConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewUpdateConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredUpdateConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredUpdateConfigInformer constructs a new informer for UpdateConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredUpdateConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TiankuixingV1().UpdateConfigs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TiankuixingV1().UpdateConfigs(namespace).Watch(context.TODO(), options)
			},
		},
		&tiankuixingv1.UpdateConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *updateConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredUpdateConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *updateConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tiankuixingv1.UpdateConfig{}, f.defaultInformer)
}

func (f *updateConfigInformer) Lister() v1.UpdateConfigLister {
	return v1.NewUpdateConfigLister(f.Informer().GetIndexer())
}
