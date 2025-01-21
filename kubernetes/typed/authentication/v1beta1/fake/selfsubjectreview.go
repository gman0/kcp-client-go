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

	"github.com/kcp-dev/logicalcluster/v3"

	authenticationv1beta1 "k8s.io/api/authentication/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	authenticationv1beta1client "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var selfSubjectReviewsResource = schema.GroupVersionResource{Group: "authentication.k8s.io", Version: "v1beta1", Resource: "selfsubjectreviews"}
var selfSubjectReviewsKind = schema.GroupVersionKind{Group: "authentication.k8s.io", Version: "v1beta1", Kind: "SelfSubjectReview"}

type selfSubjectReviewsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *selfSubjectReviewsClusterClient) Cluster(clusterPath logicalcluster.Path) authenticationv1beta1client.SelfSubjectReviewInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &selfSubjectReviewsClient{Fake: c.Fake, ClusterPath: clusterPath}
}

type selfSubjectReviewsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *selfSubjectReviewsClient) Create(ctx context.Context, selfSubjectReview *authenticationv1beta1.SelfSubjectReview, opts metav1.CreateOptions) (*authenticationv1beta1.SelfSubjectReview, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(selfSubjectReviewsResource, c.ClusterPath, selfSubjectReview), &authenticationv1beta1.SelfSubjectReview{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authenticationv1beta1.SelfSubjectReview), err
}
