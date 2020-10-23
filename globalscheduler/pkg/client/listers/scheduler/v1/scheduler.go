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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "k8s.io/kubernetes/globalscheduler/pkg/apis/scheduler/v1"
)

// SchedulerLister helps list Schedulers.
type SchedulerLister interface {
	// List lists all Schedulers in the indexer.
	List(selector labels.Selector) (ret []*v1.Scheduler, err error)
	// Schedulers returns an object that can list and get Schedulers.
	Schedulers(namespace string) SchedulerNamespaceLister
	SchedulersWithMultiTenancy(namespace string, tenant string) SchedulerNamespaceLister
	SchedulerListerExpansion
}

// schedulerLister implements the SchedulerLister interface.
type schedulerLister struct {
	indexer cache.Indexer
}

// NewSchedulerLister returns a new SchedulerLister.
func NewSchedulerLister(indexer cache.Indexer) SchedulerLister {
	return &schedulerLister{indexer: indexer}
}

// List lists all Schedulers in the indexer.
func (s *schedulerLister) List(selector labels.Selector) (ret []*v1.Scheduler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Scheduler))
	})
	return ret, err
}

// Schedulers returns an object that can list and get Schedulers.
func (s *schedulerLister) Schedulers(namespace string) SchedulerNamespaceLister {
	return schedulerNamespaceLister{indexer: s.indexer, namespace: namespace, tenant: "system"}
}

func (s *schedulerLister) SchedulersWithMultiTenancy(namespace string, tenant string) SchedulerNamespaceLister {
	return schedulerNamespaceLister{indexer: s.indexer, namespace: namespace, tenant: tenant}
}

// SchedulerNamespaceLister helps list and get Schedulers.
type SchedulerNamespaceLister interface {
	// List lists all Schedulers in the indexer for a given tenant/namespace.
	List(selector labels.Selector) (ret []*v1.Scheduler, err error)
	// Get retrieves the Scheduler from the indexer for a given tenant/namespace and name.
	Get(name string) (*v1.Scheduler, error)
	SchedulerNamespaceListerExpansion
}

// schedulerNamespaceLister implements the SchedulerNamespaceLister
// interface.
type schedulerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
	tenant    string
}

// List lists all Schedulers in the indexer for a given namespace.
func (s schedulerNamespaceLister) List(selector labels.Selector) (ret []*v1.Scheduler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.tenant, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Scheduler))
	})
	return ret, err
}

// Get retrieves the Scheduler from the indexer for a given namespace and name.
func (s schedulerNamespaceLister) Get(name string) (*v1.Scheduler, error) {
	key := s.tenant + "/" + s.namespace + "/" + name
	if s.tenant == "system" {
		key = s.namespace + "/" + name
	}
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("scheduler"), name)
	}
	return obj.(*v1.Scheduler), nil
}
