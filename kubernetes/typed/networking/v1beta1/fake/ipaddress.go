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

package fake

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsnetworkingv1beta1 "k8s.io/client-go/applyconfigurations/networking/v1beta1"
	networkingv1beta1client "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var iPAddressesResource = schema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1beta1", Resource: "ipaddresses"}
var iPAddressesKind = schema.GroupVersionKind{Group: "networking.k8s.io", Version: "v1beta1", Kind: "IPAddress"}

type iPAddressesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *iPAddressesClusterClient) Cluster(clusterPath logicalcluster.Path) networkingv1beta1client.IPAddressInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &iPAddressesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of IPAddresses that match those selectors across all clusters.
func (c *iPAddressesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*networkingv1beta1.IPAddressList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(iPAddressesResource, iPAddressesKind, logicalcluster.Wildcard, opts), &networkingv1beta1.IPAddressList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1beta1.IPAddressList{ListMeta: obj.(*networkingv1beta1.IPAddressList).ListMeta}
	for _, item := range obj.(*networkingv1beta1.IPAddressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested IPAddresses across all clusters.
func (c *iPAddressesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(iPAddressesResource, logicalcluster.Wildcard, opts))
}

type iPAddressesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *iPAddressesClient) Create(ctx context.Context, iPAddress *networkingv1beta1.IPAddress, opts metav1.CreateOptions) (*networkingv1beta1.IPAddress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(iPAddressesResource, c.ClusterPath, iPAddress), &networkingv1beta1.IPAddress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1beta1.IPAddress), err
}

func (c *iPAddressesClient) Update(ctx context.Context, iPAddress *networkingv1beta1.IPAddress, opts metav1.UpdateOptions) (*networkingv1beta1.IPAddress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(iPAddressesResource, c.ClusterPath, iPAddress), &networkingv1beta1.IPAddress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1beta1.IPAddress), err
}

func (c *iPAddressesClient) UpdateStatus(ctx context.Context, iPAddress *networkingv1beta1.IPAddress, opts metav1.UpdateOptions) (*networkingv1beta1.IPAddress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(iPAddressesResource, c.ClusterPath, "status", iPAddress), &networkingv1beta1.IPAddress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1beta1.IPAddress), err
}

func (c *iPAddressesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(iPAddressesResource, c.ClusterPath, name, opts), &networkingv1beta1.IPAddress{})
	return err
}

func (c *iPAddressesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(iPAddressesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &networkingv1beta1.IPAddressList{})
	return err
}

func (c *iPAddressesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*networkingv1beta1.IPAddress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(iPAddressesResource, c.ClusterPath, name), &networkingv1beta1.IPAddress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1beta1.IPAddress), err
}

// List takes label and field selectors, and returns the list of IPAddresses that match those selectors.
func (c *iPAddressesClient) List(ctx context.Context, opts metav1.ListOptions) (*networkingv1beta1.IPAddressList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(iPAddressesResource, iPAddressesKind, c.ClusterPath, opts), &networkingv1beta1.IPAddressList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1beta1.IPAddressList{ListMeta: obj.(*networkingv1beta1.IPAddressList).ListMeta}
	for _, item := range obj.(*networkingv1beta1.IPAddressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *iPAddressesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(iPAddressesResource, c.ClusterPath, opts))
}

func (c *iPAddressesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*networkingv1beta1.IPAddress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(iPAddressesResource, c.ClusterPath, name, pt, data, subresources...), &networkingv1beta1.IPAddress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1beta1.IPAddress), err
}

func (c *iPAddressesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsnetworkingv1beta1.IPAddressApplyConfiguration, opts metav1.ApplyOptions) (*networkingv1beta1.IPAddress, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(iPAddressesResource, c.ClusterPath, *name, types.ApplyPatchType, data), &networkingv1beta1.IPAddress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1beta1.IPAddress), err
}

func (c *iPAddressesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsnetworkingv1beta1.IPAddressApplyConfiguration, opts metav1.ApplyOptions) (*networkingv1beta1.IPAddress, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(iPAddressesResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &networkingv1beta1.IPAddress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1beta1.IPAddress), err
}
