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

package v1alpha1

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v2"

	flowcontrolv1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamflowcontrolv1alpha1informers "k8s.io/client-go/informers/flowcontrol/v1alpha1"
	upstreamflowcontrolv1alpha1listers "k8s.io/client-go/listers/flowcontrol/v1alpha1"
	"k8s.io/client-go/tools/cache"

	clientset "github.com/kcp-dev/client-go/clients/clientset/versioned"
	"github.com/kcp-dev/client-go/clients/informers/internalinterfaces"
	flowcontrolv1alpha1listers "github.com/kcp-dev/client-go/clients/listers/flowcontrol/v1alpha1"
)

// PriorityLevelConfigurationClusterInformer provides access to a shared informer and lister for
// PriorityLevelConfigurations.
type PriorityLevelConfigurationClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamflowcontrolv1alpha1informers.PriorityLevelConfigurationInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() flowcontrolv1alpha1listers.PriorityLevelConfigurationClusterLister
}

type priorityLevelConfigurationClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPriorityLevelConfigurationClusterInformer constructs a new informer for PriorityLevelConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPriorityLevelConfigurationClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredPriorityLevelConfigurationClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPriorityLevelConfigurationClusterInformer constructs a new informer for PriorityLevelConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPriorityLevelConfigurationClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowcontrolV1alpha1().PriorityLevelConfigurations().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowcontrolV1alpha1().PriorityLevelConfigurations().Watch(context.TODO(), options)
			},
		},
		&flowcontrolv1alpha1.PriorityLevelConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *priorityLevelConfigurationClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredPriorityLevelConfigurationClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *priorityLevelConfigurationClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&flowcontrolv1alpha1.PriorityLevelConfiguration{}, f.defaultInformer)
}

func (f *priorityLevelConfigurationClusterInformer) Lister() flowcontrolv1alpha1listers.PriorityLevelConfigurationClusterLister {
	return flowcontrolv1alpha1listers.NewPriorityLevelConfigurationClusterLister(f.Informer().GetIndexer())
}

func (f *priorityLevelConfigurationClusterInformer) Cluster(cluster logicalcluster.Name) upstreamflowcontrolv1alpha1informers.PriorityLevelConfigurationInformer {
	return &priorityLevelConfigurationInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type priorityLevelConfigurationInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamflowcontrolv1alpha1listers.PriorityLevelConfigurationLister
}

func (f *priorityLevelConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *priorityLevelConfigurationInformer) Lister() upstreamflowcontrolv1alpha1listers.PriorityLevelConfigurationLister {
	return f.lister
}
