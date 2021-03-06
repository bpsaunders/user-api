// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bpsaunders/user-api/transformers (interfaces: UserTransform)

// Package transformers is a generated GoMock package.
package transformers

import (
	models "github.com/bpsaunders/user-api/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserTransform is a mock of UserTransform interface
type MockUserTransform struct {
	ctrl     *gomock.Controller
	recorder *MockUserTransformMockRecorder
}

// MockUserTransformMockRecorder is the mock recorder for MockUserTransform
type MockUserTransformMockRecorder struct {
	mock *MockUserTransform
}

// NewMockUserTransform creates a new mock instance
func NewMockUserTransform(ctrl *gomock.Controller) *MockUserTransform {
	mock := &MockUserTransform{ctrl: ctrl}
	mock.recorder = &MockUserTransformMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserTransform) EXPECT() *MockUserTransformMockRecorder {
	return m.recorder
}

// ToEntity mocks base method
func (m *MockUserTransform) ToEntity(arg0 *models.User) *models.UserDao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToEntity", arg0)
	ret0, _ := ret[0].(*models.UserDao)
	return ret0
}

// ToEntity indicates an expected call of ToEntity
func (mr *MockUserTransformMockRecorder) ToEntity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToEntity", reflect.TypeOf((*MockUserTransform)(nil).ToEntity), arg0)
}

// ToRest mocks base method
func (m *MockUserTransform) ToRest(arg0 *models.UserDao) *models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToRest", arg0)
	ret0, _ := ret[0].(*models.User)
	return ret0
}

// ToRest indicates an expected call of ToRest
func (mr *MockUserTransformMockRecorder) ToRest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToRest", reflect.TypeOf((*MockUserTransform)(nil).ToRest), arg0)
}

// ToRestArray mocks base method
func (m *MockUserTransform) ToRestArray(arg0 *[]*models.UserDao) *[]*models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToRestArray", arg0)
	ret0, _ := ret[0].(*[]*models.User)
	return ret0
}

// ToRestArray indicates an expected call of ToRestArray
func (mr *MockUserTransformMockRecorder) ToRestArray(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToRestArray", reflect.TypeOf((*MockUserTransform)(nil).ToRestArray), arg0)
}
