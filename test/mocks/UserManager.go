// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	models "github.com/Conty111/CarsCatalog/internal/models"
	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// UserManager is an autogenerated mock type for the UserManager type
type UserManager struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *UserManager) CreateUser(user *models.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByID provides a mock function with given fields: id
func (_m *UserManager) DeleteByID(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByFullName provides a mock function with given fields: name, surname, patronymic
func (_m *UserManager) GetByFullName(name string, surname string, patronymic string) (*models.User, error) {
	ret := _m.Called(name, surname, patronymic)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string, string, string) *models.User); ok {
		r0 = rf(name, surname, patronymic)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(name, surname, patronymic)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *UserManager) GetByID(id uuid.UUID) (*models.User, error) {
	ret := _m.Called(id)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: id, updates
func (_m *UserManager) UpdateByID(id uuid.UUID, updates interface{}) error {
	ret := _m.Called(id, updates)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, interface{}) error); ok {
		r0 = rf(id, updates)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserManager creates a new instance of UserManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserManager(t mockConstructorTestingTNewUserManager) *UserManager {
	mock := &UserManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}