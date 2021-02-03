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

package v1

import (
	fmt "fmt"
	strings "strings"
	sync "sync"
	"time"

	errors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	diff "k8s.io/apimachinery/pkg/util/diff"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "k8s.io/code-generator/_examples/crd/apis/example/v1"
	scheme "k8s.io/code-generator/_examples/crd/clientset/versioned/scheme"
	klog "k8s.io/klog"
	autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
)

// ClusterTestTypesGetter has a method to return a ClusterTestTypeInterface.
// A group's client should implement this interface.
type ClusterTestTypesGetter interface {
	ClusterTestTypes() ClusterTestTypeInterface
}

// ClusterTestTypeInterface has methods to work with ClusterTestType resources.
type ClusterTestTypeInterface interface {
	Create(*v1.ClusterTestType) (*v1.ClusterTestType, error)
	Update(*v1.ClusterTestType) (*v1.ClusterTestType, error)
	UpdateStatus(*v1.ClusterTestType) (*v1.ClusterTestType, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ClusterTestType, error)
	List(opts metav1.ListOptions) (*v1.ClusterTestTypeList, error)
	Watch(opts metav1.ListOptions) watch.AggregatedWatchInterface
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterTestType, err error)
	GetScale(clusterTestTypeName string, options metav1.GetOptions) (*autoscaling.Scale, error)
	UpdateScale(clusterTestTypeName string, scale *autoscaling.Scale) (*autoscaling.Scale, error)

	ClusterTestTypeExpansion
}

// clusterTestTypes implements ClusterTestTypeInterface
type clusterTestTypes struct {
	client  rest.Interface
	clients []rest.Interface
}

// newClusterTestTypes returns a ClusterTestTypes
func newClusterTestTypes(c *ExampleV1Client) *clusterTestTypes {
	return &clusterTestTypes{
		client:  c.RESTClient(),
		clients: c.RESTClients(),
	}
}

// Get takes name of the clusterTestType, and returns the corresponding clusterTestType object, and an error if there is any.
func (c *clusterTestTypes) Get(name string, options metav1.GetOptions) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}
	err = c.client.Get().
		Resource("clustertesttypes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)

	return
}

// List takes label and field selectors, and returns the list of ClusterTestTypes that match those selectors.
func (c *clusterTestTypes) List(opts metav1.ListOptions) (result *v1.ClusterTestTypeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterTestTypeList{}

	wgLen := 1
	// When resource version is not empty, it reads from api server local cache
	// Need to check all api server partitions
	if opts.ResourceVersion != "" && len(c.clients) > 1 {
		wgLen = len(c.clients)
	}

	if wgLen > 1 {
		var listLock sync.Mutex

		var wg sync.WaitGroup
		wg.Add(wgLen)
		results := make(map[int]*v1.ClusterTestTypeList)
		errs := make(map[int]error)
		for i, client := range c.clients {
			go func(c *clusterTestTypes, ci rest.Interface, opts metav1.ListOptions, lock *sync.Mutex, pos int, resultMap map[int]*v1.ClusterTestTypeList, errMap map[int]error) {
				r := &v1.ClusterTestTypeList{}
				err := ci.Get().
					Resource("clustertesttypes").
					VersionedParams(&opts, scheme.ParameterCodec).
					Timeout(timeout).
					Do().
					Into(r)

				lock.Lock()
				resultMap[pos] = r
				errMap[pos] = err
				lock.Unlock()
				wg.Done()
			}(c, client, opts, &listLock, i, results, errs)
		}
		wg.Wait()

		// consolidate list result
		itemsMap := make(map[string]*v1.ClusterTestType)
		for j := 0; j < wgLen; j++ {
			currentErr, isOK := errs[j]
			if isOK && currentErr != nil {
				if !(errors.IsForbidden(currentErr) && strings.Contains(currentErr.Error(), "no relationship found between node")) {
					err = currentErr
					return
				} else {
					continue
				}
			}

			currentResult, _ := results[j]
			if result.ResourceVersion == "" {
				result.TypeMeta = currentResult.TypeMeta
				result.ListMeta = currentResult.ListMeta
			} else {
				isNewer, errCompare := diff.RevisionStrIsNewer(currentResult.ResourceVersion, result.ResourceVersion)
				if errCompare != nil {
					err = errors.NewInternalError(fmt.Errorf("Invalid resource version [%v]", errCompare))
					return
				} else if isNewer {
					// Since the lists are from different api servers with different partition. When used in list and watch,
					// we cannot watch from the biggest resource version. Leave it to watch for adjustment.
					result.ResourceVersion = currentResult.ResourceVersion
				}
			}
			for _, item := range currentResult.Items {
				if _, exist := itemsMap[item.ResourceVersion]; !exist {
					itemsMap[item.ResourceVersion] = &item
				}
			}
		}

		for _, item := range itemsMap {
			result.Items = append(result.Items, *item)
		}
		return
	}

	// The following is used for single api server partition and/or resourceVersion is empty
	// When resourceVersion is empty, objects are read from ETCD directly and will get full
	// list of data if no permission issue. The list needs to done sequential to avoid increasing
	// system load.
	err = c.client.Get().
		Resource("clustertesttypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	if err == nil {
		return
	}

	if !(errors.IsForbidden(err) && strings.Contains(err.Error(), "no relationship found between node")) {
		return
	}

	// Found api server that works with this list, keep the client
	for _, client := range c.clients {
		if client == c.client {
			continue
		}

		err = client.Get().
			Resource("clustertesttypes").
			VersionedParams(&opts, scheme.ParameterCodec).
			Timeout(timeout).
			Do().
			Into(result)

		if err == nil {
			c.client = client
			return
		}

		if err != nil && errors.IsForbidden(err) &&
			strings.Contains(err.Error(), "no relationship found between node") {
			klog.V(6).Infof("Skip error %v in list", err)
			continue
		}
	}

	return
}

// Watch returns a watch.Interface that watches the requested clusterTestTypes.
func (c *clusterTestTypes) Watch(opts metav1.ListOptions) watch.AggregatedWatchInterface {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	aggWatch := watch.NewAggregatedWatcher()
	for _, client := range c.clients {
		watcher, err := client.Get().
			Resource("clustertesttypes").
			VersionedParams(&opts, scheme.ParameterCodec).
			Timeout(timeout).
			Watch()
		if err != nil && opts.AllowPartialWatch && errors.IsForbidden(err) {
			// watch error was not returned properly in error message. Skip when partial watch is allowed
			klog.V(6).Infof("Watch error for partial watch %v. options [%+v]", err, opts)
			continue
		}
		aggWatch.AddWatchInterface(watcher, err)
	}
	return aggWatch
}

// Create takes the representation of a clusterTestType and creates it.  Returns the server's representation of the clusterTestType, and an error, if there is any.
func (c *clusterTestTypes) Create(clusterTestType *v1.ClusterTestType) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}

	err = c.client.Post().
		Resource("clustertesttypes").
		Body(clusterTestType).
		Do().
		Into(result)

	return
}

// Update takes the representation of a clusterTestType and updates it. Returns the server's representation of the clusterTestType, and an error, if there is any.
func (c *clusterTestTypes) Update(clusterTestType *v1.ClusterTestType) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}

	err = c.client.Put().
		Resource("clustertesttypes").
		Name(clusterTestType.Name).
		Body(clusterTestType).
		Do().
		Into(result)

	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterTestTypes) UpdateStatus(clusterTestType *v1.ClusterTestType) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}

	err = c.client.Put().
		Resource("clustertesttypes").
		Name(clusterTestType.Name).
		SubResource("status").
		Body(clusterTestType).
		Do().
		Into(result)

	return
}

// Delete takes name of the clusterTestType and deletes it. Returns an error if one occurs.
func (c *clusterTestTypes) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clustertesttypes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterTestTypes) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clustertesttypes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterTestType.
func (c *clusterTestTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}
	err = c.client.Patch(pt).
		Resource("clustertesttypes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)

	return
}

// GetScale takes name of the clusterTestType, and returns the corresponding autoscaling.Scale object, and an error if there is any.
func (c *clusterTestTypes) GetScale(clusterTestTypeName string, options metav1.GetOptions) (result *autoscaling.Scale, err error) {
	result = &autoscaling.Scale{}
	err = c.client.Get().
		Resource("clustertesttypes").
		Name(clusterTestTypeName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)

	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *clusterTestTypes) UpdateScale(clusterTestTypeName string, scale *autoscaling.Scale) (result *autoscaling.Scale, err error) {
	result = &autoscaling.Scale{}

	err = c.client.Put().
		Resource("clustertesttypes").
		Name(clusterTestTypeName).
		SubResource("scale").
		Body(scale).
		Do().
		Into(result)

	return
}
