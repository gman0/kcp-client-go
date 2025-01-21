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

	storagemigrationv1alpha1 "k8s.io/client-go/kubernetes/typed/storagemigration/v1alpha1"
	"k8s.io/client-go/rest"

	kcpstoragemigrationv1alpha1 "github.com/kcp-dev/client-go/kubernetes/typed/storagemigration/v1alpha1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpstoragemigrationv1alpha1.StoragemigrationV1alpha1ClusterInterface = (*StoragemigrationV1alpha1ClusterClient)(nil)

type StoragemigrationV1alpha1ClusterClient struct {
	*kcptesting.Fake
}

func (c *StoragemigrationV1alpha1ClusterClient) Cluster(clusterPath logicalcluster.Path) storagemigrationv1alpha1.StoragemigrationV1alpha1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &StoragemigrationV1alpha1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *StoragemigrationV1alpha1ClusterClient) StorageVersionMigrations() kcpstoragemigrationv1alpha1.StorageVersionMigrationClusterInterface {
	return &storageVersionMigrationsClusterClient{Fake: c.Fake}
}

var _ storagemigrationv1alpha1.StoragemigrationV1alpha1Interface = (*StoragemigrationV1alpha1Client)(nil)

type StoragemigrationV1alpha1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *StoragemigrationV1alpha1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *StoragemigrationV1alpha1Client) StorageVersionMigrations() storagemigrationv1alpha1.StorageVersionMigrationInterface {
	return &storageVersionMigrationsClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}
