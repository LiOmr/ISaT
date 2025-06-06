// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2024_2_SaraFun/internal/pkg/image/usecase (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	multipart "mime/multipart"
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

// DeleteImage mocks base method.
func (m *MockRepository) DeleteImage(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteImage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteImage indicates an expected call of DeleteImage.
func (mr *MockRepositoryMockRecorder) DeleteImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImage", reflect.TypeOf((*MockRepository)(nil).DeleteImage), arg0, arg1)
}

// GetFirstImage mocks base method.
func (m *MockRepository) GetFirstImage(arg0 context.Context, arg1 int) (models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirstImage", arg0, arg1)
	ret0, _ := ret[0].(models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirstImage indicates an expected call of GetFirstImage.
func (mr *MockRepositoryMockRecorder) GetFirstImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirstImage", reflect.TypeOf((*MockRepository)(nil).GetFirstImage), arg0, arg1)
}

// GetImageLinksByUserId mocks base method.
func (m *MockRepository) GetImageLinksByUserId(arg0 context.Context, arg1 int) ([]models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageLinksByUserId", arg0, arg1)
	ret0, _ := ret[0].([]models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageLinksByUserId indicates an expected call of GetImageLinksByUserId.
func (mr *MockRepositoryMockRecorder) GetImageLinksByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageLinksByUserId", reflect.TypeOf((*MockRepository)(nil).GetImageLinksByUserId), arg0, arg1)
}

// SaveImage mocks base method.
func (m *MockRepository) SaveImage(arg0 context.Context, arg1 multipart.File, arg2 string, arg3, arg4 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveImage", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveImage indicates an expected call of SaveImage.
func (mr *MockRepositoryMockRecorder) SaveImage(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveImage", reflect.TypeOf((*MockRepository)(nil).SaveImage), arg0, arg1, arg2, arg3, arg4)
}

// UpdateOrdNumbers mocks base method.
func (m *MockRepository) UpdateOrdNumbers(arg0 context.Context, arg1 []models.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrdNumbers", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrdNumbers indicates an expected call of UpdateOrdNumbers.
func (mr *MockRepositoryMockRecorder) UpdateOrdNumbers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrdNumbers", reflect.TypeOf((*MockRepository)(nil).UpdateOrdNumbers), arg0, arg1)
}
