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

	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/rest"

	kcpappsv1 "github.com/kcp-dev/client-go/kubernetes/typed/apps/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpappsv1.AppsV1ClusterInterface = (*AppsV1ClusterClient)(nil)

type AppsV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *AppsV1ClusterClient) Cluster(clusterPath logicalcluster.Path) appsv1.AppsV1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &AppsV1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *AppsV1ClusterClient) StatefulSets() kcpappsv1.StatefulSetClusterInterface {
	return &statefulSetsClusterClient{Fake: c.Fake}
}

func (c *AppsV1ClusterClient) Deployments() kcpappsv1.DeploymentClusterInterface {
	return &deploymentsClusterClient{Fake: c.Fake}
}

func (c *AppsV1ClusterClient) DaemonSets() kcpappsv1.DaemonSetClusterInterface {
	return &daemonSetsClusterClient{Fake: c.Fake}
}

func (c *AppsV1ClusterClient) ReplicaSets() kcpappsv1.ReplicaSetClusterInterface {
	return &replicaSetsClusterClient{Fake: c.Fake}
}

func (c *AppsV1ClusterClient) ControllerRevisions() kcpappsv1.ControllerRevisionClusterInterface {
	return &controllerRevisionsClusterClient{Fake: c.Fake}
}

var _ appsv1.AppsV1Interface = (*AppsV1Client)(nil)

type AppsV1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *AppsV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *AppsV1Client) StatefulSets(namespace string) appsv1.StatefulSetInterface {
	return &statefulSetsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *AppsV1Client) Deployments(namespace string) appsv1.DeploymentInterface {
	return &deploymentsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *AppsV1Client) DaemonSets(namespace string) appsv1.DaemonSetInterface {
	return &daemonSetsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *AppsV1Client) ReplicaSets(namespace string) appsv1.ReplicaSetInterface {
	return &replicaSetsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *AppsV1Client) ControllerRevisions(namespace string) appsv1.ControllerRevisionInterface {
	return &controllerRevisionsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
