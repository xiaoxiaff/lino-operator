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
	v1 "github.com/lino-network/lino-operator/pkg/apis/samplecrd/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FullnodeLister helps list Fullnodes.
type FullnodeLister interface {
	// List lists all Fullnodes in the indexer.
	List(selector labels.Selector) (ret []*v1.Fullnode, err error)
	// Fullnodes returns an object that can list and get Fullnodes.
	Fullnodes(namespace string) FullnodeNamespaceLister
	FullnodeListerExpansion
}

// fullnodeLister implements the FullnodeLister interface.
type fullnodeLister struct {
	indexer cache.Indexer
}

// NewFullnodeLister returns a new FullnodeLister.
func NewFullnodeLister(indexer cache.Indexer) FullnodeLister {
	return &fullnodeLister{indexer: indexer}
}

// List lists all Fullnodes in the indexer.
func (s *fullnodeLister) List(selector labels.Selector) (ret []*v1.Fullnode, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Fullnode))
	})
	return ret, err
}

// Fullnodes returns an object that can list and get Fullnodes.
func (s *fullnodeLister) Fullnodes(namespace string) FullnodeNamespaceLister {
	return fullnodeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FullnodeNamespaceLister helps list and get Fullnodes.
type FullnodeNamespaceLister interface {
	// List lists all Fullnodes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Fullnode, err error)
	// Get retrieves the Fullnode from the indexer for a given namespace and name.
	Get(name string) (*v1.Fullnode, error)
	FullnodeNamespaceListerExpansion
}

// fullnodeNamespaceLister implements the FullnodeNamespaceLister
// interface.
type fullnodeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Fullnodes in the indexer for a given namespace.
func (s fullnodeNamespaceLister) List(selector labels.Selector) (ret []*v1.Fullnode, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Fullnode))
	})
	return ret, err
}

// Get retrieves the Fullnode from the indexer for a given namespace and name.
func (s fullnodeNamespaceLister) Get(name string) (*v1.Fullnode, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("fullnode"), name)
	}
	return obj.(*v1.Fullnode), nil
}
