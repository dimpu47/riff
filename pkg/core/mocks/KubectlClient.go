// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import core "github.com/projectriff/riff/pkg/core"
import mock "github.com/stretchr/testify/mock"

// KubectlClient is an autogenerated mock type for the KubectlClient type
type KubectlClient struct {
	mock.Mock
}

// NamespaceInit provides a mock function with given fields: manifests, options
func (_m *KubectlClient) NamespaceInit(manifests map[string]*core.Manifest, options core.NamespaceInitOptions) error {
	ret := _m.Called(manifests, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]*core.Manifest, core.NamespaceInitOptions) error); ok {
		r0 = rf(manifests, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SystemInstall provides a mock function with given fields: manifests, options
func (_m *KubectlClient) SystemInstall(manifests map[string]*core.Manifest, options core.SystemInstallOptions) (bool, error) {
	ret := _m.Called(manifests, options)

	var r0 bool
	if rf, ok := ret.Get(0).(func(map[string]*core.Manifest, core.SystemInstallOptions) bool); ok {
		r0 = rf(manifests, options)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]*core.Manifest, core.SystemInstallOptions) error); ok {
		r1 = rf(manifests, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemUninstall provides a mock function with given fields: options
func (_m *KubectlClient) SystemUninstall(options core.SystemUninstallOptions) (bool, error) {
	ret := _m.Called(options)

	var r0 bool
	if rf, ok := ret.Get(0).(func(core.SystemUninstallOptions) bool); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.SystemUninstallOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}