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

package v1beta3

import (
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	flowcontrolv1beta3 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta3"
	"k8s.io/client-go/rest"
)

type FlowcontrolV1beta3ClusterInterface interface {
	FlowcontrolV1beta3ClusterScoper
	FlowSchemasClusterGetter
	PriorityLevelConfigurationsClusterGetter
}

type FlowcontrolV1beta3ClusterScoper interface {
	Cluster(logicalcluster.Path) flowcontrolv1beta3.FlowcontrolV1beta3Interface
}

type FlowcontrolV1beta3ClusterClient struct {
	clientCache kcpclient.Cache[*flowcontrolv1beta3.FlowcontrolV1beta3Client]
}

func (c *FlowcontrolV1beta3ClusterClient) Cluster(clusterPath logicalcluster.Path) flowcontrolv1beta3.FlowcontrolV1beta3Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(clusterPath)
}

func (c *FlowcontrolV1beta3ClusterClient) FlowSchemas() FlowSchemaClusterInterface {
	return &flowSchemasClusterInterface{clientCache: c.clientCache}
}

func (c *FlowcontrolV1beta3ClusterClient) PriorityLevelConfigurations() PriorityLevelConfigurationClusterInterface {
	return &priorityLevelConfigurationsClusterInterface{clientCache: c.clientCache}
}

// NewForConfig creates a new FlowcontrolV1beta3ClusterClient for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*FlowcontrolV1beta3ClusterClient, error) {
	client, err := rest.HTTPClientFor(c)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(c, client)
}

// NewForConfigAndClient creates a new FlowcontrolV1beta3ClusterClient for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*FlowcontrolV1beta3ClusterClient, error) {
	cache := kcpclient.NewCache(c, h, &kcpclient.Constructor[*flowcontrolv1beta3.FlowcontrolV1beta3Client]{
		NewForConfigAndClient: flowcontrolv1beta3.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.Name("root").Path()); err != nil {
		return nil, err
	}
	return &FlowcontrolV1beta3ClusterClient{clientCache: cache}, nil
}

// NewForConfigOrDie creates a new FlowcontrolV1beta3ClusterClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *FlowcontrolV1beta3ClusterClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
