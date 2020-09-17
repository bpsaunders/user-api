// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bpsaunders/user-api/validators (interfaces: UserValidate)

// Package validators is a generated GoMock package.
package validators

import (
	models "github.com/bpsaunders/user-api/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserValidate is a mock of UserValidate interface
type MockUserValidate struct {
	ctrl     *gomock.Controller
	recorder *MockUserValidateMockRecorder
}

// MockUserValidateMockRecorder is the mock recorder for MockUserValidate
type MockUserValidateMockRecorder struct {
	mock *MockUserValidate
}

// NewMockUserValidate creates a new mock instance
func NewMockUserValidate(ctrl *gomock.Controller) *MockUserValidate {
	mock := &MockUserValidate{ctrl: ctrl}
	mock.recorder = &MockUserValidateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserValidate) EXPECT() *MockUserValidateMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockUserValidate) Validate(arg0 *models.User) []ValidationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].([]ValidationError)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockUserValidateMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockUserValidate)(nil).Validate), arg0)
}
