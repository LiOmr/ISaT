// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2024_2_SaraFun/internal/pkg/personalities/usecase/user (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2024_2_SaraFun/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockRepository) AddUser(arg0 context.Context, arg1 models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockRepositoryMockRecorder) AddUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockRepository)(nil).AddUser), arg0, arg1)
}

// ChangePassword mocks base method.
func (m *MockRepository) ChangePassword(arg0 context.Context, arg1 int, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockRepositoryMockRecorder) ChangePassword(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockRepository)(nil).ChangePassword), arg0, arg1, arg2)
}

// CheckUsernameExists mocks base method.
func (m *MockRepository) CheckUsernameExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUsernameExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUsernameExists indicates an expected call of CheckUsernameExists.
func (mr *MockRepositoryMockRecorder) CheckUsernameExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUsernameExists", reflect.TypeOf((*MockRepository)(nil).CheckUsernameExists), arg0, arg1)
}

// GetFeedList mocks base method.
func (m *MockRepository) GetFeedList(arg0 context.Context, arg1 int, arg2 []int) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeedList", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeedList indicates an expected call of GetFeedList.
func (mr *MockRepositoryMockRecorder) GetFeedList(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeedList", reflect.TypeOf((*MockRepository)(nil).GetFeedList), arg0, arg1, arg2)
}

// GetProfileIdByUserId mocks base method.
func (m *MockRepository) GetProfileIdByUserId(arg0 context.Context, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileIdByUserId", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileIdByUserId indicates an expected call of GetProfileIdByUserId.
func (mr *MockRepositoryMockRecorder) GetProfileIdByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileIdByUserId", reflect.TypeOf((*MockRepository)(nil).GetProfileIdByUserId), arg0, arg1)
}

// GetUserByUsername mocks base method.
func (m *MockRepository) GetUserByUsername(arg0 context.Context, arg1 string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockRepositoryMockRecorder) GetUserByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockRepository)(nil).GetUserByUsername), arg0, arg1)
}

// GetUserIdByUsername mocks base method.
func (m *MockRepository) GetUserIdByUsername(arg0 context.Context, arg1 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIdByUsername", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIdByUsername indicates an expected call of GetUserIdByUsername.
func (mr *MockRepositoryMockRecorder) GetUserIdByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIdByUsername", reflect.TypeOf((*MockRepository)(nil).GetUserIdByUsername), arg0, arg1)
}

// GetUserList mocks base method.
func (m *MockRepository) GetUserList(arg0 context.Context, arg1 int) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserList", arg0, arg1)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserList indicates an expected call of GetUserList.
func (mr *MockRepositoryMockRecorder) GetUserList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserList", reflect.TypeOf((*MockRepository)(nil).GetUserList), arg0, arg1)
}

// GetUsernameByUserId mocks base method.
func (m *MockRepository) GetUsernameByUserId(arg0 context.Context, arg1 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsernameByUserId", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsernameByUserId indicates an expected call of GetUsernameByUserId.
func (mr *MockRepositoryMockRecorder) GetUsernameByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsernameByUserId", reflect.TypeOf((*MockRepository)(nil).GetUsernameByUserId), arg0, arg1)
}
