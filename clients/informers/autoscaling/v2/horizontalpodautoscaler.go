//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v2

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v2"

	autoscalingv2 "k8s.io/api/autoscaling/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamautoscalingv2informers "k8s.io/client-go/informers/autoscaling/v2"
	upstreamautoscalingv2listers "k8s.io/client-go/listers/autoscaling/v2"
	"k8s.io/client-go/tools/cache"

	clientset "github.com/kcp-dev/client-go/clients/clientset/versioned"
	"github.com/kcp-dev/client-go/clients/informers/internalinterfaces"
	autoscalingv2listers "github.com/kcp-dev/client-go/clients/listers/autoscaling/v2"
)

// HorizontalPodAutoscalerClusterInformer provides access to a shared informer and lister for
// HorizontalPodAutoscalers.
type HorizontalPodAutoscalerClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamautoscalingv2informers.HorizontalPodAutoscalerInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() autoscalingv2listers.HorizontalPodAutoscalerClusterLister
}

type horizontalPodAutoscalerClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewHorizontalPodAutoscalerClusterInformer constructs a new informer for HorizontalPodAutoscaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHorizontalPodAutoscalerClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredHorizontalPodAutoscalerClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredHorizontalPodAutoscalerClusterInformer constructs a new informer for HorizontalPodAutoscaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHorizontalPodAutoscalerClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV2().HorizontalPodAutoscalers().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV2().HorizontalPodAutoscalers().Watch(context.TODO(), options)
			},
		},
		&autoscalingv2.HorizontalPodAutoscaler{},
		resyncPeriod,
		indexers,
	)
}

func (f *horizontalPodAutoscalerClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredHorizontalPodAutoscalerClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc},
		f.tweakListOptions,
	)
}

func (f *horizontalPodAutoscalerClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&autoscalingv2.HorizontalPodAutoscaler{}, f.defaultInformer)
}

func (f *horizontalPodAutoscalerClusterInformer) Lister() autoscalingv2listers.HorizontalPodAutoscalerClusterLister {
	return autoscalingv2listers.NewHorizontalPodAutoscalerClusterLister(f.Informer().GetIndexer())
}

func (f *horizontalPodAutoscalerClusterInformer) Cluster(cluster logicalcluster.Name) upstreamautoscalingv2informers.HorizontalPodAutoscalerInformer {
	return &horizontalPodAutoscalerInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type horizontalPodAutoscalerInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamautoscalingv2listers.HorizontalPodAutoscalerLister
}

func (f *horizontalPodAutoscalerInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *horizontalPodAutoscalerInformer) Lister() upstreamautoscalingv2listers.HorizontalPodAutoscalerLister {
	return f.lister
}
