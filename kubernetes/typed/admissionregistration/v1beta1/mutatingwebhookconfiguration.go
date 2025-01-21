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
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	admissionregistrationv1beta1client "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
)

// MutatingWebhookConfigurationsClusterGetter has a method to return a MutatingWebhookConfigurationClusterInterface.
// A group's cluster client should implement this interface.
type MutatingWebhookConfigurationsClusterGetter interface {
	MutatingWebhookConfigurations() MutatingWebhookConfigurationClusterInterface
}

// MutatingWebhookConfigurationClusterInterface can operate on MutatingWebhookConfigurations across all clusters,
// or scope down to one cluster and return a admissionregistrationv1beta1client.MutatingWebhookConfigurationInterface.
type MutatingWebhookConfigurationClusterInterface interface {
	Cluster(logicalcluster.Path) admissionregistrationv1beta1client.MutatingWebhookConfigurationInterface
	List(ctx context.Context, opts metav1.ListOptions) (*admissionregistrationv1beta1.MutatingWebhookConfigurationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type mutatingWebhookConfigurationsClusterInterface struct {
	clientCache kcpclient.Cache[*admissionregistrationv1beta1client.AdmissionregistrationV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *mutatingWebhookConfigurationsClusterInterface) Cluster(clusterPath logicalcluster.Path) admissionregistrationv1beta1client.MutatingWebhookConfigurationInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).MutatingWebhookConfigurations()
}

// List returns the entire collection of all MutatingWebhookConfigurations across all clusters.
func (c *mutatingWebhookConfigurationsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*admissionregistrationv1beta1.MutatingWebhookConfigurationList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).MutatingWebhookConfigurations().List(ctx, opts)
}

// Watch begins to watch all MutatingWebhookConfigurations across all clusters.
func (c *mutatingWebhookConfigurationsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).MutatingWebhookConfigurations().Watch(ctx, opts)
}
