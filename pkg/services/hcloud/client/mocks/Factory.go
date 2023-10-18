// Code generated by mockery v2.35.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	hcloudclient "github.com/syself/cluster-api-provider-hetzner/pkg/services/hcloud/client"
)

// Factory is an autogenerated mock type for the Factory type
type Factory struct {
	mock.Mock
}

// NewClient provides a mock function with given fields: hcloudToken
func (_m *Factory) NewClient(hcloudToken string) hcloudclient.Client {
	ret := _m.Called(hcloudToken)

	var r0 hcloudclient.Client
	if rf, ok := ret.Get(0).(func(string) hcloudclient.Client); ok {
		r0 = rf(hcloudToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(hcloudclient.Client)
		}
	}

	return r0
}

// NewFactory creates a new instance of Factory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *Factory {
	mock := &Factory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
