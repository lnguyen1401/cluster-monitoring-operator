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

package v1alpha1

import (
	v1alpha1 "k8s.io/api/discovery/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EndpointSliceLister helps list EndpointSlices.
// All objects returned here must be treated as read-only.
type EndpointSliceLister interface {
	// List lists all EndpointSlices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EndpointSlice, err error)
	// EndpointSlices returns an object that can list and get EndpointSlices.
	EndpointSlices(namespace string) EndpointSliceNamespaceLister
	EndpointSliceListerExpansion
}

// endpointSliceLister implements the EndpointSliceLister interface.
type endpointSliceLister struct {
	indexer cache.Indexer
}

// NewEndpointSliceLister returns a new EndpointSliceLister.
func NewEndpointSliceLister(indexer cache.Indexer) EndpointSliceLister {
	return &endpointSliceLister{indexer: indexer}
}

// List lists all EndpointSlices in the indexer.
func (s *endpointSliceLister) List(selector labels.Selector) (ret []*v1alpha1.EndpointSlice, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EndpointSlice))
	})
	return ret, err
}

// EndpointSlices returns an object that can list and get EndpointSlices.
func (s *endpointSliceLister) EndpointSlices(namespace string) EndpointSliceNamespaceLister {
	return endpointSliceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EndpointSliceNamespaceLister helps list and get EndpointSlices.
// All objects returned here must be treated as read-only.
type EndpointSliceNamespaceLister interface {
	// List lists all EndpointSlices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EndpointSlice, err error)
	// Get retrieves the EndpointSlice from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EndpointSlice, error)
	EndpointSliceNamespaceListerExpansion
}

// endpointSliceNamespaceLister implements the EndpointSliceNamespaceLister
// interface.
type endpointSliceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EndpointSlices in the indexer for a given namespace.
func (s endpointSliceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EndpointSlice, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EndpointSlice))
	})
	return ret, err
}

// Get retrieves the EndpointSlice from the indexer for a given namespace and name.
func (s endpointSliceNamespaceLister) Get(name string) (*v1alpha1.EndpointSlice, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("endpointslice"), name)
	}
	return obj.(*v1alpha1.EndpointSlice), nil
}