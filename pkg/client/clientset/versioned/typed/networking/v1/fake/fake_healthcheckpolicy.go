/*
* Copyright 2024 Google LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     https://www.apache.org/licenses/LICENSE-2.0
*
*     Unless required by applicable law or agreed to in writing, software
*     distributed under the License is distributed on an "AS IS" BASIS,
*     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*     See the License for the specific language governing permissions and
*     limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/GoogleCloudPlatform/gke-gateway-api/apis/networking/v1"
	networkingv1 "github.com/GoogleCloudPlatform/gke-gateway-api/pkg/client/clientset/versioned/typed/networking/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeHealthCheckPolicies implements HealthCheckPolicyInterface
type fakeHealthCheckPolicies struct {
	*gentype.FakeClientWithList[*v1.HealthCheckPolicy, *v1.HealthCheckPolicyList]
	Fake *FakeNetworkingV1
}

func newFakeHealthCheckPolicies(fake *FakeNetworkingV1, namespace string) networkingv1.HealthCheckPolicyInterface {
	return &fakeHealthCheckPolicies{
		gentype.NewFakeClientWithList[*v1.HealthCheckPolicy, *v1.HealthCheckPolicyList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("healthcheckpolicies"),
			v1.SchemeGroupVersion.WithKind("HealthCheckPolicy"),
			func() *v1.HealthCheckPolicy { return &v1.HealthCheckPolicy{} },
			func() *v1.HealthCheckPolicyList { return &v1.HealthCheckPolicyList{} },
			func(dst, src *v1.HealthCheckPolicyList) { dst.ListMeta = src.ListMeta },
			func(list *v1.HealthCheckPolicyList) []*v1.HealthCheckPolicy {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1.HealthCheckPolicyList, items []*v1.HealthCheckPolicy) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
