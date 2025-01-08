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
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	admissionregistrationv1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	admissionregistrationv1alpha1listers "k8s.io/client-go/listers/admissionregistration/v1alpha1"
	"k8s.io/client-go/tools/cache"
)

// MutatingAdmissionPolicyBindingClusterLister can list MutatingAdmissionPolicyBindings across all workspaces, or scope down to a MutatingAdmissionPolicyBindingLister for one workspace.
// All objects returned here must be treated as read-only.
type MutatingAdmissionPolicyBindingClusterLister interface {
	// List lists all MutatingAdmissionPolicyBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*admissionregistrationv1alpha1.MutatingAdmissionPolicyBinding, err error)
	// Cluster returns a lister that can list and get MutatingAdmissionPolicyBindings in one workspace.
	Cluster(clusterName logicalcluster.Name) admissionregistrationv1alpha1listers.MutatingAdmissionPolicyBindingLister
	MutatingAdmissionPolicyBindingClusterListerExpansion
}

type mutatingAdmissionPolicyBindingClusterLister struct {
	indexer cache.Indexer
}

// NewMutatingAdmissionPolicyBindingClusterLister returns a new MutatingAdmissionPolicyBindingClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewMutatingAdmissionPolicyBindingClusterLister(indexer cache.Indexer) *mutatingAdmissionPolicyBindingClusterLister {
	return &mutatingAdmissionPolicyBindingClusterLister{indexer: indexer}
}

// List lists all MutatingAdmissionPolicyBindings in the indexer across all workspaces.
func (s *mutatingAdmissionPolicyBindingClusterLister) List(selector labels.Selector) (ret []*admissionregistrationv1alpha1.MutatingAdmissionPolicyBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*admissionregistrationv1alpha1.MutatingAdmissionPolicyBinding))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get MutatingAdmissionPolicyBindings.
func (s *mutatingAdmissionPolicyBindingClusterLister) Cluster(clusterName logicalcluster.Name) admissionregistrationv1alpha1listers.MutatingAdmissionPolicyBindingLister {
	return &mutatingAdmissionPolicyBindingLister{indexer: s.indexer, clusterName: clusterName}
}

// mutatingAdmissionPolicyBindingLister implements the admissionregistrationv1alpha1listers.MutatingAdmissionPolicyBindingLister interface.
type mutatingAdmissionPolicyBindingLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all MutatingAdmissionPolicyBindings in the indexer for a workspace.
func (s *mutatingAdmissionPolicyBindingLister) List(selector labels.Selector) (ret []*admissionregistrationv1alpha1.MutatingAdmissionPolicyBinding, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*admissionregistrationv1alpha1.MutatingAdmissionPolicyBinding))
	})
	return ret, err
}

// Get retrieves the MutatingAdmissionPolicyBinding from the indexer for a given workspace and name.
func (s *mutatingAdmissionPolicyBindingLister) Get(name string) (*admissionregistrationv1alpha1.MutatingAdmissionPolicyBinding, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(admissionregistrationv1alpha1.Resource("mutatingadmissionpolicybindings"), name)
	}
	return obj.(*admissionregistrationv1alpha1.MutatingAdmissionPolicyBinding), nil
}
