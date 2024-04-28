// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	external_api "github.com/Conty111/CarsCatalog/internal/external_api"
	mock "github.com/stretchr/testify/mock"
)

// ExternalAPIClient is an autogenerated mock type for the ExternalAPIClient type
type ExternalAPIClient struct {
	mock.Mock
}

// GetCarInfo provides a mock function with given fields: regNum
func (_m *ExternalAPIClient) GetCarInfo(regNum string) (*external_api.CarData, error) {
	ret := _m.Called(regNum)

	var r0 *external_api.CarData
	if rf, ok := ret.Get(0).(func(string) *external_api.CarData); ok {
		r0 = rf(regNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*external_api.CarData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(regNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewExternalAPIClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewExternalAPIClient creates a new instance of ExternalAPIClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExternalAPIClient(t mockConstructorTestingTNewExternalAPIClient) *ExternalAPIClient {
	mock := &ExternalAPIClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}