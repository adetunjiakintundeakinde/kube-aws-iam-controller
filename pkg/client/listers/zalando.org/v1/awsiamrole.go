/*
Copyright 2019 The Kubernetes Authors.

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
	v1 "github.com/mikkeloscar/kube-aws-iam-controller/pkg/apis/zalando.org/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AWSIAMRoleLister helps list AWSIAMRoles.
type AWSIAMRoleLister interface {
	// List lists all AWSIAMRoles in the indexer.
	List(selector labels.Selector) (ret []*v1.AWSIAMRole, err error)
	// AWSIAMRoles returns an object that can list and get AWSIAMRoles.
	AWSIAMRoles(namespace string) AWSIAMRoleNamespaceLister
	AWSIAMRoleListerExpansion
}

// aWSIAMRoleLister implements the AWSIAMRoleLister interface.
type aWSIAMRoleLister struct {
	indexer cache.Indexer
}

// NewAWSIAMRoleLister returns a new AWSIAMRoleLister.
func NewAWSIAMRoleLister(indexer cache.Indexer) AWSIAMRoleLister {
	return &aWSIAMRoleLister{indexer: indexer}
}

// List lists all AWSIAMRoles in the indexer.
func (s *aWSIAMRoleLister) List(selector labels.Selector) (ret []*v1.AWSIAMRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AWSIAMRole))
	})
	return ret, err
}

// AWSIAMRoles returns an object that can list and get AWSIAMRoles.
func (s *aWSIAMRoleLister) AWSIAMRoles(namespace string) AWSIAMRoleNamespaceLister {
	return aWSIAMRoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AWSIAMRoleNamespaceLister helps list and get AWSIAMRoles.
type AWSIAMRoleNamespaceLister interface {
	// List lists all AWSIAMRoles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AWSIAMRole, err error)
	// Get retrieves the AWSIAMRole from the indexer for a given namespace and name.
	Get(name string) (*v1.AWSIAMRole, error)
	AWSIAMRoleNamespaceListerExpansion
}

// aWSIAMRoleNamespaceLister implements the AWSIAMRoleNamespaceLister
// interface.
type aWSIAMRoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AWSIAMRoles in the indexer for a given namespace.
func (s aWSIAMRoleNamespaceLister) List(selector labels.Selector) (ret []*v1.AWSIAMRole, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AWSIAMRole))
	})
	return ret, err
}

// Get retrieves the AWSIAMRole from the indexer for a given namespace and name.
func (s aWSIAMRoleNamespaceLister) Get(name string) (*v1.AWSIAMRole, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsiamrole"), name)
	}
	return obj.(*v1.AWSIAMRole), nil
}