/*
Copyright 2022 The KCP Authors.

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

package metadatainformer

import (
	"context"
	"sync"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	thirdpartyinformers "github.com/kcp-dev/apimachinery/third_party/informers"
	kcpinformers "github.com/kcp-dev/client-go/clients/informers"
	kcpmetadata "github.com/kcp-dev/client-go/clients/metadata"
	kcpmetadatalisters "github.com/kcp-dev/client-go/clients/metadata/metadatalister"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/metadata/metadatainformer"
	"k8s.io/client-go/tools/cache"
)

// NewSharedInformerFactory constructs a new instance of sharedInformerFactory for all namespaces.
func NewSharedInformerFactory(client kcpmetadata.ClusterInterface, defaultResync time.Duration) SharedInformerFactory {
	return NewFilteredSharedInformerFactory(client, defaultResync, nil)
}

// NewFilteredSharedInformerFactory constructs a new instance of sharedInformerFactory.
// Listers obtained via this factory will be subject to the same filters as specified here.
func NewFilteredSharedInformerFactory(client kcpmetadata.ClusterInterface, defaultResync time.Duration, tweakListOptions metadatainformer.TweakListOptionsFunc) SharedInformerFactory {
	return &sharedInformerFactory{
		client:           client,
		defaultResync:    defaultResync,
		informers:        map[schema.GroupVersionResource]kcpinformers.GenericClusterInformer{},
		startedInformers: make(map[schema.GroupVersionResource]bool),
		tweakListOptions: tweakListOptions,
	}
}

type sharedInformerFactory struct {
	client        kcpmetadata.ClusterInterface
	defaultResync time.Duration

	lock      sync.Mutex
	informers map[schema.GroupVersionResource]kcpinformers.GenericClusterInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[schema.GroupVersionResource]bool
	tweakListOptions metadatainformer.TweakListOptionsFunc
}

var _ SharedInformerFactory = &sharedInformerFactory{}

func (f *sharedInformerFactory) ForResource(gvr schema.GroupVersionResource) kcpinformers.GenericClusterInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	key := gvr
	informer, exists := f.informers[key]
	if exists {
		return informer
	}

	informer = NewFilteredDynamicInformer(f.client, gvr, f.defaultResync, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc}, f.tweakListOptions)
	f.informers[key] = informer

	return informer
}

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			go informer.Informer().Run(stopCh)
			f.startedInformers[informerType] = true
		}
	}
}

// WaitForCacheSync waits for all started informers' cache were synced.
func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[schema.GroupVersionResource]bool {
	informers := func() map[schema.GroupVersionResource]cache.SharedIndexInformer {
		f.lock.Lock()
		defer f.lock.Unlock()

		informers := map[schema.GroupVersionResource]cache.SharedIndexInformer{}
		for informerType, informer := range f.informers {
			if f.startedInformers[informerType] {
				informers[informerType] = informer.Informer()
			}
		}
		return informers
	}()

	res := map[schema.GroupVersionResource]bool{}
	for informType, informer := range informers {
		res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
	}
	return res
}

// NewFilteredDynamicInformer constructs a new informer for a dynamic type.
func NewFilteredDynamicInformer(client kcpmetadata.ClusterInterface, gvr schema.GroupVersionResource, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions metadatainformer.TweakListOptionsFunc) kcpinformers.GenericClusterInformer {
	return &dynamicInformer{
		gvr: gvr,
		informer: thirdpartyinformers.NewSharedIndexInformer(
			&cache.ListWatch{
				ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
					if tweakListOptions != nil {
						tweakListOptions(&options)
					}
					return client.Resource(gvr).List(context.TODO(), options)
				},
				WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
					if tweakListOptions != nil {
						tweakListOptions(&options)
					}
					return client.Resource(gvr).Watch(context.TODO(), options)
				},
			},
			&unstructured.Unstructured{},
			resyncPeriod,
			indexers,
		),
	}
}

type dynamicInformer struct {
	informer cache.SharedIndexInformer
	gvr      schema.GroupVersionResource
}

var _ kcpinformers.GenericClusterInformer = &dynamicInformer{}

func (d *dynamicInformer) Informer() cache.SharedIndexInformer {
	return d.informer
}

func (d *dynamicInformer) Lister() kcpcache.GenericClusterLister {
	return kcpmetadatalisters.NewRuntimeObjectShim(kcpmetadatalisters.New(d.informer.GetIndexer(), d.gvr))
}
