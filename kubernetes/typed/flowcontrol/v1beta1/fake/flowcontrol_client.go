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

	flowcontrolv1beta1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1"
	"k8s.io/client-go/rest"

	kcpflowcontrolv1beta1 "github.com/kcp-dev/client-go/kubernetes/typed/flowcontrol/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpflowcontrolv1beta1.FlowcontrolV1beta1ClusterInterface = (*FlowcontrolV1beta1ClusterClient)(nil)

type FlowcontrolV1beta1ClusterClient struct {
	*kcptesting.Fake
}

func (c *FlowcontrolV1beta1ClusterClient) Cluster(clusterPath logicalcluster.Path) flowcontrolv1beta1.FlowcontrolV1beta1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &FlowcontrolV1beta1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *FlowcontrolV1beta1ClusterClient) FlowSchemas() kcpflowcontrolv1beta1.FlowSchemaClusterInterface {
	return &flowSchemasClusterClient{Fake: c.Fake}
}

func (c *FlowcontrolV1beta1ClusterClient) PriorityLevelConfigurations() kcpflowcontrolv1beta1.PriorityLevelConfigurationClusterInterface {
	return &priorityLevelConfigurationsClusterClient{Fake: c.Fake}
}

var _ flowcontrolv1beta1.FlowcontrolV1beta1Interface = (*FlowcontrolV1beta1Client)(nil)

type FlowcontrolV1beta1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *FlowcontrolV1beta1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *FlowcontrolV1beta1Client) FlowSchemas() flowcontrolv1beta1.FlowSchemaInterface {
	return &flowSchemasClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *FlowcontrolV1beta1Client) PriorityLevelConfigurations() flowcontrolv1beta1.PriorityLevelConfigurationInterface {
	return &priorityLevelConfigurationsClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}
