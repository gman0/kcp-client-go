
Replace removed `coordination/v1alpha1` packages with `v1alpha2`:
```sh
git grep -wl '"k8s.io/api/coordination/v1alpha1"' -- '*.go' | xargs sd -F '"k8s.io/api/coordination/v1alpha1"' '"k8s.io/api/coordination/v1alpha2"'
git grep -wl '"k8s.io/client-go/informers/coordination/v1alpha1"' -- '*.go' | xargs sd -F '"k8s.io/client-go/informers/coordination/v1alpha1"' '"k8s.io/client-go/informers/coordination/v1alpha2"'
git grep -wl '"k8s.io/client-go/listers/coordination/v1alpha1"' -- '*.go' | xargs sd -F '"k8s.io/client-go/listers/coordination/v1alpha1"' '"k8s.io/client-go/listers/coordination/v1alpha2"'
git grep -wl '"k8s.io/client-go/kubernetes/typed/coordination/v1alpha1"' -- '*.go' | xargs sd -F '"k8s.io/client-go/kubernetes/typed/coordination/v1alpha1"' '"k8s.io/client-go/kubernetes/typed/coordination/v1alpha2"'
git grep -wl '"k8s.io/client-go/applyconfigurations/coordination/v1alpha1"' -- '*.go' | xargs sd -F '"k8s.io/client-go/applyconfigurations/coordination/v1alpha1"' '"k8s.io/client-go/applyconfigurations/coordination/v1alpha2"'
```

- Coordination.v1alpha1 API is dropped and replaced with coordination.v1alpha2. Old coordination.v1alpha1 types must be deleted before upgrade ([#127857](https://github.com/kubernetes/kubernetes/pull/127857), [@Jefftree](https://github.com/Jefftree)) [SIG API Machinery, Etcd, Scheduling and Testing]


# Missing commits

Author: Ben Luddy <bluddy@redhat.com>
Date:   Thu Nov 7 00:05:03 2024 -0500

    Fix content type fallback when a client defaults to CBOR.

    With the ClientsAllowCBOR client-go feature gate enabled, a 415 response to a CBOR-encoded REST
    causes all subsequent requests from the client to fall back to a JSON request encoding. This
    mechanism had only worked as intended when CBOR was explicitly configured in the
    ClientContentConfig. When both ClientsAllowCBOR and ClientsPreferCBOR are enabled, an
    unconfigured (empty) content type defaults to CBOR instead of JSON. Both ways of configuring a
    client to use the CBOR request encoding are now subject to the same fallback mechanism.

    Kubernetes-commit: a77f4c7ba2e761461daaf115a38903fc91916dd6

commit c57dbd8decb083cef12d85b224b8db2000d12d32
Author: Kubernetes Publisher <k8s-publishing-bot@users.noreply.github.com>
Date:   Wed Nov 6 23:17:35 2024 +0000

    Merge pull request #128503 from benluddy/cbor-codecs-featuregate

    KEP-4222: Wire serving codecs to CBOR feature gate.

    Kubernetes-commit: 6399c32669c62cfbf7c33b14b77d6781ce1cce27


commit 334e30739dff78f897fa59b5da1149dcb241c33d
Author: Ben Luddy <bluddy@redhat.com>
Date:   Fri Nov 1 16:05:32 2024 -0400

    Wire serving codecs to CBOR feature gate.

    Integration testing has to this point relied on patching serving codecs for built-in APIs. The
    test-only patching is removed and replaced by feature gated checks at runtime.

    Kubernetes-commit: 439d2f7b4028638b3d8d9261bb046c3ba8d9bfcb

commit 10c2fdb7f4cf6f43d04242ca27f0669beadcf43d
Author: Ben Luddy <bluddy@redhat.com>
Date:   Fri Nov 1 13:14:06 2024 -0400

    Use application/cbor-seq media type in streaming CBOR responses.

    The media type application/cbor describes exactly one encoded item. As a new (to Kubernetes) format
    with no existing clients, streaming/watch responses will use the application/cbor-seq media
    type. CBOR watch responses conform to the specification of CBOR Sequences and are encoded as the
    concatenation of zero or more items with no additional framing.

    Kubernetes-commit: 504f14998e920ca8837b3310094b3da11c62a070
