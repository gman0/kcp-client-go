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
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	networkingv1alpha1 "k8s.io/api/networking/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	networkingv1alpha1listers "k8s.io/client-go/listers/networking/v1alpha1"
	"k8s.io/client-go/tools/cache"
)

// IPAddressClusterLister can list IPAddresses across all workspaces, or scope down to a IPAddressLister for one workspace.
// All objects returned here must be treated as read-only.
type IPAddressClusterLister interface {
	// List lists all IPAddresses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networkingv1alpha1.IPAddress, err error)
	// Cluster returns a lister that can list and get IPAddresses in one workspace.
	Cluster(clusterName logicalcluster.Name) networkingv1alpha1listers.IPAddressLister
	IPAddressClusterListerExpansion
}

type iPAddressClusterLister struct {
	indexer cache.Indexer
}

// NewIPAddressClusterLister returns a new IPAddressClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewIPAddressClusterLister(indexer cache.Indexer) *iPAddressClusterLister {
	return &iPAddressClusterLister{indexer: indexer}
}

// List lists all IPAddresses in the indexer across all workspaces.
func (s *iPAddressClusterLister) List(selector labels.Selector) (ret []*networkingv1alpha1.IPAddress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*networkingv1alpha1.IPAddress))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get IPAddresses.
func (s *iPAddressClusterLister) Cluster(clusterName logicalcluster.Name) networkingv1alpha1listers.IPAddressLister {
	return &iPAddressLister{indexer: s.indexer, clusterName: clusterName}
}

// iPAddressLister implements the networkingv1alpha1listers.IPAddressLister interface.
type iPAddressLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all IPAddresses in the indexer for a workspace.
func (s *iPAddressLister) List(selector labels.Selector) (ret []*networkingv1alpha1.IPAddress, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*networkingv1alpha1.IPAddress))
	})
	return ret, err
}

// Get retrieves the IPAddress from the indexer for a given workspace and name.
func (s *iPAddressLister) Get(name string) (*networkingv1alpha1.IPAddress, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(networkingv1alpha1.Resource("ipaddresses"), name)
	}
	return obj.(*networkingv1alpha1.IPAddress), nil
}
