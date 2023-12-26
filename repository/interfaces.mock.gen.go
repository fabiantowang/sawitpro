// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockRepositoryInterface) AddUser(ctx context.Context, input AddUserInput) (AddUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", ctx, input)
	ret0, _ := ret[0].(AddUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockRepositoryInterfaceMockRecorder) AddUser(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockRepositoryInterface)(nil).AddUser), ctx, input)
}

// GetUserById mocks base method.
func (m *MockRepositoryInterface) GetUserById(ctx context.Context, id string) (GetUserByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id)
	ret0, _ := ret[0].(GetUserByIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserById), ctx, id)
}

// GetUserByPhone mocks base method.
func (m *MockRepositoryInterface) GetUserByPhone(ctx context.Context, phone string) (GetUserByPhoneOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPhone", ctx, phone)
	ret0, _ := ret[0].(GetUserByPhoneOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPhone indicates an expected call of GetUserByPhone.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserByPhone(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPhone", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserByPhone), ctx, phone)
}

// IncrementSuccessfulLogin mocks base method.
func (m *MockRepositoryInterface) IncrementSuccessfulLogin(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementSuccessfulLogin", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementSuccessfulLogin indicates an expected call of IncrementSuccessfulLogin.
func (mr *MockRepositoryInterfaceMockRecorder) IncrementSuccessfulLogin(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementSuccessfulLogin", reflect.TypeOf((*MockRepositoryInterface)(nil).IncrementSuccessfulLogin), ctx, id)
}

// UpdateUser mocks base method.
func (m *MockRepositoryInterface) UpdateUser(ctx context.Context, input UpdateUserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateUser(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateUser), ctx, input)
}