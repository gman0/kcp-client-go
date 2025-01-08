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

package fake

import (
	"github.com/kcp-dev/logicalcluster/v3"

	resourcev1alpha3 "k8s.io/client-go/kubernetes/typed/resource/v1alpha3"
	"k8s.io/client-go/rest"

	kcpresourcev1alpha3 "github.com/kcp-dev/client-go/kubernetes/typed/resource/v1alpha3"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpresourcev1alpha3.ResourceV1alpha3ClusterInterface = (*ResourceV1alpha3ClusterClient)(nil)

type ResourceV1alpha3ClusterClient struct {
	*kcptesting.Fake
}

func (c *ResourceV1alpha3ClusterClient) Cluster(clusterPath logicalcluster.Path) resourcev1alpha3.ResourceV1alpha3Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &ResourceV1alpha3Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *ResourceV1alpha3ClusterClient) ResourceSlices() kcpresourcev1alpha3.ResourceSliceClusterInterface {
	return &resourceSlicesClusterClient{Fake: c.Fake}
}

func (c *ResourceV1alpha3ClusterClient) ResourceClaims() kcpresourcev1alpha3.ResourceClaimClusterInterface {
	return &resourceClaimsClusterClient{Fake: c.Fake}
}

func (c *ResourceV1alpha3ClusterClient) DeviceClasses() kcpresourcev1alpha3.DeviceClassClusterInterface {
	return &deviceClassesClusterClient{Fake: c.Fake}
}

func (c *ResourceV1alpha3ClusterClient) ResourceClaimTemplates() kcpresourcev1alpha3.ResourceClaimTemplateClusterInterface {
	return &resourceClaimTemplatesClusterClient{Fake: c.Fake}
}

var _ resourcev1alpha3.ResourceV1alpha3Interface = (*ResourceV1alpha3Client)(nil)

type ResourceV1alpha3Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *ResourceV1alpha3Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *ResourceV1alpha3Client) ResourceSlices() resourcev1alpha3.ResourceSliceInterface {
	return &resourceSlicesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *ResourceV1alpha3Client) ResourceClaims(namespace string) resourcev1alpha3.ResourceClaimInterface {
	return &resourceClaimsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *ResourceV1alpha3Client) DeviceClasses() resourcev1alpha3.DeviceClassInterface {
	return &deviceClassesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *ResourceV1alpha3Client) ResourceClaimTemplates(namespace string) resourcev1alpha3.ResourceClaimTemplateInterface {
	return &resourceClaimTemplatesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
