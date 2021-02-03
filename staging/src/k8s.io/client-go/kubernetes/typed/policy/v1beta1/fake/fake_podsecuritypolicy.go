/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "k8s.io/api/policy/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodSecurityPolicies implements PodSecurityPolicyInterface
type FakePodSecurityPolicies struct {
	Fake *FakePolicyV1beta1
	te   string
}

var podsecuritypoliciesResource = schema.GroupVersionResource{Group: "policy", Version: "v1beta1", Resource: "podsecuritypolicies"}

var podsecuritypoliciesKind = schema.GroupVersionKind{Group: "policy", Version: "v1beta1", Kind: "PodSecurityPolicy"}

// Get takes name of the podSecurityPolicy, and returns the corresponding podSecurityPolicy object, and an error if there is any.
func (c *FakePodSecurityPolicies) Get(name string, options v1.GetOptions) (result *v1beta1.PodSecurityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewTenantGetAction(podsecuritypoliciesResource, name, c.te), &v1beta1.PodSecurityPolicy{})

	if obj == nil {
		return nil, err
	}

	return obj.(*v1beta1.PodSecurityPolicy), err
}

// List takes label and field selectors, and returns the list of PodSecurityPolicies that match those selectors.
func (c *FakePodSecurityPolicies) List(opts v1.ListOptions) (result *v1beta1.PodSecurityPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewTenantListAction(podsecuritypoliciesResource, podsecuritypoliciesKind, opts, c.te), &v1beta1.PodSecurityPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PodSecurityPolicyList{ListMeta: obj.(*v1beta1.PodSecurityPolicyList).ListMeta}
	for _, item := range obj.(*v1beta1.PodSecurityPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.AggregatedWatchInterface that watches the requested podSecurityPolicies.
func (c *FakePodSecurityPolicies) Watch(opts v1.ListOptions) watch.AggregatedWatchInterface {
	aggWatch := watch.NewAggregatedWatcher()
	watcher, err := c.Fake.
		InvokesWatch(testing.NewTenantWatchAction(podsecuritypoliciesResource, opts, c.te))

	aggWatch.AddWatchInterface(watcher, err)
	return aggWatch
}

// Create takes the representation of a podSecurityPolicy and creates it.  Returns the server's representation of the podSecurityPolicy, and an error, if there is any.
func (c *FakePodSecurityPolicies) Create(podSecurityPolicy *v1beta1.PodSecurityPolicy) (result *v1beta1.PodSecurityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewTenantCreateAction(podsecuritypoliciesResource, podSecurityPolicy, c.te), &v1beta1.PodSecurityPolicy{})

	if obj == nil {
		return nil, err
	}

	return obj.(*v1beta1.PodSecurityPolicy), err
}

// Update takes the representation of a podSecurityPolicy and updates it. Returns the server's representation of the podSecurityPolicy, and an error, if there is any.
func (c *FakePodSecurityPolicies) Update(podSecurityPolicy *v1beta1.PodSecurityPolicy) (result *v1beta1.PodSecurityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewTenantUpdateAction(podsecuritypoliciesResource, podSecurityPolicy, c.te), &v1beta1.PodSecurityPolicy{})

	if obj == nil {
		return nil, err
	}

	return obj.(*v1beta1.PodSecurityPolicy), err
}

// Delete takes name of the podSecurityPolicy and deletes it. Returns an error if one occurs.
func (c *FakePodSecurityPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewTenantDeleteAction(podsecuritypoliciesResource, name, c.te), &v1beta1.PodSecurityPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodSecurityPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {

	action := testing.NewTenantDeleteCollectionAction(podsecuritypoliciesResource, listOptions, c.te)

	_, err := c.Fake.Invokes(action, &v1beta1.PodSecurityPolicyList{})
	return err
}

// Patch applies the patch and returns the patched podSecurityPolicy.
func (c *FakePodSecurityPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.PodSecurityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewTenantPatchSubresourceAction(podsecuritypoliciesResource, c.te, name, pt, data, subresources...), &v1beta1.PodSecurityPolicy{})

	if obj == nil {
		return nil, err
	}

	return obj.(*v1beta1.PodSecurityPolicy), err
}
