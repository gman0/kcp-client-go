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

package v1beta1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	networkingv1beta1listers "k8s.io/client-go/listers/networking/v1beta1"
	"k8s.io/client-go/tools/cache"
)

// IngressClusterLister can list Ingresses across all workspaces, or scope down to a IngressLister for one workspace.
// All objects returned here must be treated as read-only.
type IngressClusterLister interface {
	// List lists all Ingresses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networkingv1beta1.Ingress, err error)
	// Cluster returns a lister that can list and get Ingresses in one workspace.
	Cluster(clusterName logicalcluster.Name) networkingv1beta1listers.IngressLister
	IngressClusterListerExpansion
}

type ingressClusterLister struct {
	indexer cache.Indexer
}

// NewIngressClusterLister returns a new IngressClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewIngressClusterLister(indexer cache.Indexer) *ingressClusterLister {
	return &ingressClusterLister{indexer: indexer}
}

// List lists all Ingresses in the indexer across all workspaces.
func (s *ingressClusterLister) List(selector labels.Selector) (ret []*networkingv1beta1.Ingress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*networkingv1beta1.Ingress))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get Ingresses.
func (s *ingressClusterLister) Cluster(clusterName logicalcluster.Name) networkingv1beta1listers.IngressLister {
	return &ingressLister{indexer: s.indexer, clusterName: clusterName}
}

// ingressLister implements the networkingv1beta1listers.IngressLister interface.
type ingressLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all Ingresses in the indexer for a workspace.
func (s *ingressLister) List(selector labels.Selector) (ret []*networkingv1beta1.Ingress, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*networkingv1beta1.Ingress))
	})
	return ret, err
}

// Ingresses returns an object that can list and get Ingresses in one namespace.
func (s *ingressLister) Ingresses(namespace string) networkingv1beta1listers.IngressNamespaceLister {
	return &ingressNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// ingressNamespaceLister implements the networkingv1beta1listers.IngressNamespaceLister interface.
type ingressNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// List lists all Ingresses in the indexer for a given workspace and namespace.
func (s *ingressNamespaceLister) List(selector labels.Selector) (ret []*networkingv1beta1.Ingress, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*networkingv1beta1.Ingress))
	})
	return ret, err
}

// Get retrieves the Ingress from the indexer for a given workspace, namespace and name.
func (s *ingressNamespaceLister) Get(name string) (*networkingv1beta1.Ingress, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(networkingv1beta1.Resource("ingresses"), name)
	}
	return obj.(*networkingv1beta1.Ingress), nil
}
