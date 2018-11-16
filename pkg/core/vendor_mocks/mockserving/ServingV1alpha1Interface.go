// Code generated by mockery v1.0.0. DO NOT EDIT.

package mockserving

import mock "github.com/stretchr/testify/mock"
import rest "k8s.io/client-go/rest"
import v1alpha1 "github.com/knative/serving/pkg/client/clientset/versioned/typed/serving/v1alpha1"

// ServingV1alpha1Interface is an autogenerated mock type for the ServingV1alpha1Interface type
type ServingV1alpha1Interface struct {
	mock.Mock
}

// Configurations provides a mock function with given fields: namespace
func (_m *ServingV1alpha1Interface) Configurations(namespace string) v1alpha1.ConfigurationInterface {
	ret := _m.Called(namespace)

	var r0 v1alpha1.ConfigurationInterface
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ConfigurationInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ConfigurationInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *ServingV1alpha1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}

// Revisions provides a mock function with given fields: namespace
func (_m *ServingV1alpha1Interface) Revisions(namespace string) v1alpha1.RevisionInterface {
	ret := _m.Called(namespace)

	var r0 v1alpha1.RevisionInterface
	if rf, ok := ret.Get(0).(func(string) v1alpha1.RevisionInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.RevisionInterface)
		}
	}

	return r0
}

// Routes provides a mock function with given fields: namespace
func (_m *ServingV1alpha1Interface) Routes(namespace string) v1alpha1.RouteInterface {
	ret := _m.Called(namespace)

	var r0 v1alpha1.RouteInterface
	if rf, ok := ret.Get(0).(func(string) v1alpha1.RouteInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.RouteInterface)
		}
	}

	return r0
}

// Services provides a mock function with given fields: namespace
func (_m *ServingV1alpha1Interface) Services(namespace string) v1alpha1.ServiceInterface {
	ret := _m.Called(namespace)

	var r0 v1alpha1.ServiceInterface
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ServiceInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ServiceInterface)
		}
	}

	return r0
}
