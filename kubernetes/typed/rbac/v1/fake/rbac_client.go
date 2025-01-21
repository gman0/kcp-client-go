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

	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	"k8s.io/client-go/rest"

	kcprbacv1 "github.com/kcp-dev/client-go/kubernetes/typed/rbac/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcprbacv1.RbacV1ClusterInterface = (*RbacV1ClusterClient)(nil)

type RbacV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *RbacV1ClusterClient) Cluster(clusterPath logicalcluster.Path) rbacv1.RbacV1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &RbacV1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *RbacV1ClusterClient) Roles() kcprbacv1.RoleClusterInterface {
	return &rolesClusterClient{Fake: c.Fake}
}

func (c *RbacV1ClusterClient) RoleBindings() kcprbacv1.RoleBindingClusterInterface {
	return &roleBindingsClusterClient{Fake: c.Fake}
}

func (c *RbacV1ClusterClient) ClusterRoles() kcprbacv1.ClusterRoleClusterInterface {
	return &clusterRolesClusterClient{Fake: c.Fake}
}

func (c *RbacV1ClusterClient) ClusterRoleBindings() kcprbacv1.ClusterRoleBindingClusterInterface {
	return &clusterRoleBindingsClusterClient{Fake: c.Fake}
}

var _ rbacv1.RbacV1Interface = (*RbacV1Client)(nil)

type RbacV1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *RbacV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *RbacV1Client) Roles(namespace string) rbacv1.RoleInterface {
	return &rolesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *RbacV1Client) RoleBindings(namespace string) rbacv1.RoleBindingInterface {
	return &roleBindingsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *RbacV1Client) ClusterRoles() rbacv1.ClusterRoleInterface {
	return &clusterRolesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *RbacV1Client) ClusterRoleBindings() rbacv1.ClusterRoleBindingInterface {
	return &clusterRoleBindingsClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}
