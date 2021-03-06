// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/api/core/v1"
)

// limitRangeLister is an autogenerated mock type for the limitRangeLister type
type limitRangeLister struct {
	mock.Mock
}

// Create provides a mock function with given fields: namespace, name, limitRangeInput
func (_m *limitRangeLister) Create(namespace string, name string, limitRangeInput gqlschema.LimitRangeInput) (*v1.LimitRange, error) {
	ret := _m.Called(namespace, name, limitRangeInput)

	var r0 *v1.LimitRange
	if rf, ok := ret.Get(0).(func(string, string, gqlschema.LimitRangeInput) *v1.LimitRange); ok {
		r0 = rf(namespace, name, limitRangeInput)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.LimitRange)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, gqlschema.LimitRangeInput) error); ok {
		r1 = rf(namespace, name, limitRangeInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ns
func (_m *limitRangeLister) List(ns string) ([]*v1.LimitRange, error) {
	ret := _m.Called(ns)

	var r0 []*v1.LimitRange
	if rf, ok := ret.Get(0).(func(string) []*v1.LimitRange); ok {
		r0 = rf(ns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.LimitRange)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
