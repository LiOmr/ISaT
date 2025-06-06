// Code generated by MockGen. DO NOT EDIT.
// Source: personalities_grpc.pb.go

// Package mock_gen is a generated GoMock package.
package mock_gen

import (
	context "context"
	reflect "reflect"

	gen "github.com/go-park-mail-ru/2024_2_SaraFun/internal/pkg/personalities/delivery/grpc/gen"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockPersonalitiesClient is a mock of PersonalitiesClient interface.
type MockPersonalitiesClient struct {
	ctrl     *gomock.Controller
	recorder *MockPersonalitiesClientMockRecorder
}

// MockPersonalitiesClientMockRecorder is the mock recorder for MockPersonalitiesClient.
type MockPersonalitiesClientMockRecorder struct {
	mock *MockPersonalitiesClient
}

// NewMockPersonalitiesClient creates a new mock instance.
func NewMockPersonalitiesClient(ctrl *gomock.Controller) *MockPersonalitiesClient {
	mock := &MockPersonalitiesClient{ctrl: ctrl}
	mock.recorder = &MockPersonalitiesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonalitiesClient) EXPECT() *MockPersonalitiesClientMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method.
func (m *MockPersonalitiesClient) ChangePassword(ctx context.Context, in *gen.ChangePasswordRequest, opts ...grpc.CallOption) (*gen.ChangePasswordResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangePassword", varargs...)
	ret0, _ := ret[0].(*gen.ChangePasswordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockPersonalitiesClientMockRecorder) ChangePassword(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockPersonalitiesClient)(nil).ChangePassword), varargs...)
}

// CheckPassword mocks base method.
func (m *MockPersonalitiesClient) CheckPassword(ctx context.Context, in *gen.CheckPasswordRequest, opts ...grpc.CallOption) (*gen.CheckPasswordResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckPassword", varargs...)
	ret0, _ := ret[0].(*gen.CheckPasswordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPassword indicates an expected call of CheckPassword.
func (mr *MockPersonalitiesClientMockRecorder) CheckPassword(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPassword", reflect.TypeOf((*MockPersonalitiesClient)(nil).CheckPassword), varargs...)
}

// CheckUsernameExists mocks base method.
func (m *MockPersonalitiesClient) CheckUsernameExists(ctx context.Context, in *gen.CheckUsernameExistsRequest, opts ...grpc.CallOption) (*gen.CheckUsernameExistsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckUsernameExists", varargs...)
	ret0, _ := ret[0].(*gen.CheckUsernameExistsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUsernameExists indicates an expected call of CheckUsernameExists.
func (mr *MockPersonalitiesClientMockRecorder) CheckUsernameExists(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUsernameExists", reflect.TypeOf((*MockPersonalitiesClient)(nil).CheckUsernameExists), varargs...)
}

// CreateProfile mocks base method.
func (m *MockPersonalitiesClient) CreateProfile(ctx context.Context, in *gen.CreateProfileRequest, opts ...grpc.CallOption) (*gen.CreateProfileResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateProfile", varargs...)
	ret0, _ := ret[0].(*gen.CreateProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProfile indicates an expected call of CreateProfile.
func (mr *MockPersonalitiesClientMockRecorder) CreateProfile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfile", reflect.TypeOf((*MockPersonalitiesClient)(nil).CreateProfile), varargs...)
}

// DeleteProfile mocks base method.
func (m *MockPersonalitiesClient) DeleteProfile(ctx context.Context, in *gen.DeleteProfileRequest, opts ...grpc.CallOption) (*gen.DeleteProfileResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteProfile", varargs...)
	ret0, _ := ret[0].(*gen.DeleteProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProfile indicates an expected call of DeleteProfile.
func (mr *MockPersonalitiesClientMockRecorder) DeleteProfile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProfile", reflect.TypeOf((*MockPersonalitiesClient)(nil).DeleteProfile), varargs...)
}

// GetFeedList mocks base method.
func (m *MockPersonalitiesClient) GetFeedList(ctx context.Context, in *gen.GetFeedListRequest, opts ...grpc.CallOption) (*gen.GetFeedListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFeedList", varargs...)
	ret0, _ := ret[0].(*gen.GetFeedListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeedList indicates an expected call of GetFeedList.
func (mr *MockPersonalitiesClientMockRecorder) GetFeedList(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeedList", reflect.TypeOf((*MockPersonalitiesClient)(nil).GetFeedList), varargs...)
}

// GetProfile mocks base method.
func (m *MockPersonalitiesClient) GetProfile(ctx context.Context, in *gen.GetProfileRequest, opts ...grpc.CallOption) (*gen.GetProfileResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProfile", varargs...)
	ret0, _ := ret[0].(*gen.GetProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockPersonalitiesClientMockRecorder) GetProfile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockPersonalitiesClient)(nil).GetProfile), varargs...)
}

// GetProfileIDByUserID mocks base method.
func (m *MockPersonalitiesClient) GetProfileIDByUserID(ctx context.Context, in *gen.GetProfileIDByUserIDRequest, opts ...grpc.CallOption) (*gen.GetProfileIDByUserIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProfileIDByUserID", varargs...)
	ret0, _ := ret[0].(*gen.GetProfileIDByUserIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileIDByUserID indicates an expected call of GetProfileIDByUserID.
func (mr *MockPersonalitiesClientMockRecorder) GetProfileIDByUserID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileIDByUserID", reflect.TypeOf((*MockPersonalitiesClient)(nil).GetProfileIDByUserID), varargs...)
}

// GetUserIDByUsername mocks base method.
func (m *MockPersonalitiesClient) GetUserIDByUsername(ctx context.Context, in *gen.GetUserIDByUsernameRequest, opts ...grpc.CallOption) (*gen.GetUserIDByUsernameResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserIDByUsername", varargs...)
	ret0, _ := ret[0].(*gen.GetUserIDByUsernameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIDByUsername indicates an expected call of GetUserIDByUsername.
func (mr *MockPersonalitiesClientMockRecorder) GetUserIDByUsername(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIDByUsername", reflect.TypeOf((*MockPersonalitiesClient)(nil).GetUserIDByUsername), varargs...)
}

// GetUsernameByUserID mocks base method.
func (m *MockPersonalitiesClient) GetUsernameByUserID(ctx context.Context, in *gen.GetUsernameByUserIDRequest, opts ...grpc.CallOption) (*gen.GetUsernameByUserIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsernameByUserID", varargs...)
	ret0, _ := ret[0].(*gen.GetUsernameByUserIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsernameByUserID indicates an expected call of GetUsernameByUserID.
func (mr *MockPersonalitiesClientMockRecorder) GetUsernameByUserID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsernameByUserID", reflect.TypeOf((*MockPersonalitiesClient)(nil).GetUsernameByUserID), varargs...)
}

// RegisterUser mocks base method.
func (m *MockPersonalitiesClient) RegisterUser(ctx context.Context, in *gen.RegisterUserRequest, opts ...grpc.CallOption) (*gen.RegisterUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterUser", varargs...)
	ret0, _ := ret[0].(*gen.RegisterUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockPersonalitiesClientMockRecorder) RegisterUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockPersonalitiesClient)(nil).RegisterUser), varargs...)
}

// UpdateProfile mocks base method.
func (m *MockPersonalitiesClient) UpdateProfile(ctx context.Context, in *gen.UpdateProfileRequest, opts ...grpc.CallOption) (*gen.UpdateProfileResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateProfile", varargs...)
	ret0, _ := ret[0].(*gen.UpdateProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockPersonalitiesClientMockRecorder) UpdateProfile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockPersonalitiesClient)(nil).UpdateProfile), varargs...)
}

// MockPersonalitiesServer is a mock of PersonalitiesServer interface.
type MockPersonalitiesServer struct {
	ctrl     *gomock.Controller
	recorder *MockPersonalitiesServerMockRecorder
}

// MockPersonalitiesServerMockRecorder is the mock recorder for MockPersonalitiesServer.
type MockPersonalitiesServerMockRecorder struct {
	mock *MockPersonalitiesServer
}

// NewMockPersonalitiesServer creates a new mock instance.
func NewMockPersonalitiesServer(ctrl *gomock.Controller) *MockPersonalitiesServer {
	mock := &MockPersonalitiesServer{ctrl: ctrl}
	mock.recorder = &MockPersonalitiesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonalitiesServer) EXPECT() *MockPersonalitiesServerMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method.
func (m *MockPersonalitiesServer) ChangePassword(arg0 context.Context, arg1 *gen.ChangePasswordRequest) (*gen.ChangePasswordResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", arg0, arg1)
	ret0, _ := ret[0].(*gen.ChangePasswordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockPersonalitiesServerMockRecorder) ChangePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockPersonalitiesServer)(nil).ChangePassword), arg0, arg1)
}

// CheckPassword mocks base method.
func (m *MockPersonalitiesServer) CheckPassword(arg0 context.Context, arg1 *gen.CheckPasswordRequest) (*gen.CheckPasswordResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPassword", arg0, arg1)
	ret0, _ := ret[0].(*gen.CheckPasswordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPassword indicates an expected call of CheckPassword.
func (mr *MockPersonalitiesServerMockRecorder) CheckPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPassword", reflect.TypeOf((*MockPersonalitiesServer)(nil).CheckPassword), arg0, arg1)
}

// CheckUsernameExists mocks base method.
func (m *MockPersonalitiesServer) CheckUsernameExists(arg0 context.Context, arg1 *gen.CheckUsernameExistsRequest) (*gen.CheckUsernameExistsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUsernameExists", arg0, arg1)
	ret0, _ := ret[0].(*gen.CheckUsernameExistsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUsernameExists indicates an expected call of CheckUsernameExists.
func (mr *MockPersonalitiesServerMockRecorder) CheckUsernameExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUsernameExists", reflect.TypeOf((*MockPersonalitiesServer)(nil).CheckUsernameExists), arg0, arg1)
}

// CreateProfile mocks base method.
func (m *MockPersonalitiesServer) CreateProfile(arg0 context.Context, arg1 *gen.CreateProfileRequest) (*gen.CreateProfileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfile", arg0, arg1)
	ret0, _ := ret[0].(*gen.CreateProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProfile indicates an expected call of CreateProfile.
func (mr *MockPersonalitiesServerMockRecorder) CreateProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfile", reflect.TypeOf((*MockPersonalitiesServer)(nil).CreateProfile), arg0, arg1)
}

// DeleteProfile mocks base method.
func (m *MockPersonalitiesServer) DeleteProfile(arg0 context.Context, arg1 *gen.DeleteProfileRequest) (*gen.DeleteProfileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProfile", arg0, arg1)
	ret0, _ := ret[0].(*gen.DeleteProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProfile indicates an expected call of DeleteProfile.
func (mr *MockPersonalitiesServerMockRecorder) DeleteProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProfile", reflect.TypeOf((*MockPersonalitiesServer)(nil).DeleteProfile), arg0, arg1)
}

// GetFeedList mocks base method.
func (m *MockPersonalitiesServer) GetFeedList(arg0 context.Context, arg1 *gen.GetFeedListRequest) (*gen.GetFeedListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeedList", arg0, arg1)
	ret0, _ := ret[0].(*gen.GetFeedListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeedList indicates an expected call of GetFeedList.
func (mr *MockPersonalitiesServerMockRecorder) GetFeedList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeedList", reflect.TypeOf((*MockPersonalitiesServer)(nil).GetFeedList), arg0, arg1)
}

// GetProfile mocks base method.
func (m *MockPersonalitiesServer) GetProfile(arg0 context.Context, arg1 *gen.GetProfileRequest) (*gen.GetProfileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", arg0, arg1)
	ret0, _ := ret[0].(*gen.GetProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockPersonalitiesServerMockRecorder) GetProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockPersonalitiesServer)(nil).GetProfile), arg0, arg1)
}

// GetProfileIDByUserID mocks base method.
func (m *MockPersonalitiesServer) GetProfileIDByUserID(arg0 context.Context, arg1 *gen.GetProfileIDByUserIDRequest) (*gen.GetProfileIDByUserIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileIDByUserID", arg0, arg1)
	ret0, _ := ret[0].(*gen.GetProfileIDByUserIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileIDByUserID indicates an expected call of GetProfileIDByUserID.
func (mr *MockPersonalitiesServerMockRecorder) GetProfileIDByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileIDByUserID", reflect.TypeOf((*MockPersonalitiesServer)(nil).GetProfileIDByUserID), arg0, arg1)
}

// GetUserIDByUsername mocks base method.
func (m *MockPersonalitiesServer) GetUserIDByUsername(arg0 context.Context, arg1 *gen.GetUserIDByUsernameRequest) (*gen.GetUserIDByUsernameResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIDByUsername", arg0, arg1)
	ret0, _ := ret[0].(*gen.GetUserIDByUsernameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIDByUsername indicates an expected call of GetUserIDByUsername.
func (mr *MockPersonalitiesServerMockRecorder) GetUserIDByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIDByUsername", reflect.TypeOf((*MockPersonalitiesServer)(nil).GetUserIDByUsername), arg0, arg1)
}

// GetUsernameByUserID mocks base method.
func (m *MockPersonalitiesServer) GetUsernameByUserID(arg0 context.Context, arg1 *gen.GetUsernameByUserIDRequest) (*gen.GetUsernameByUserIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsernameByUserID", arg0, arg1)
	ret0, _ := ret[0].(*gen.GetUsernameByUserIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsernameByUserID indicates an expected call of GetUsernameByUserID.
func (mr *MockPersonalitiesServerMockRecorder) GetUsernameByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsernameByUserID", reflect.TypeOf((*MockPersonalitiesServer)(nil).GetUsernameByUserID), arg0, arg1)
}

// RegisterUser mocks base method.
func (m *MockPersonalitiesServer) RegisterUser(arg0 context.Context, arg1 *gen.RegisterUserRequest) (*gen.RegisterUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", arg0, arg1)
	ret0, _ := ret[0].(*gen.RegisterUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockPersonalitiesServerMockRecorder) RegisterUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockPersonalitiesServer)(nil).RegisterUser), arg0, arg1)
}

// UpdateProfile mocks base method.
func (m *MockPersonalitiesServer) UpdateProfile(arg0 context.Context, arg1 *gen.UpdateProfileRequest) (*gen.UpdateProfileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", arg0, arg1)
	ret0, _ := ret[0].(*gen.UpdateProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockPersonalitiesServerMockRecorder) UpdateProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockPersonalitiesServer)(nil).UpdateProfile), arg0, arg1)
}

// mustEmbedUnimplementedPersonalitiesServer mocks base method.
func (m *MockPersonalitiesServer) mustEmbedUnimplementedPersonalitiesServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPersonalitiesServer")
}

// mustEmbedUnimplementedPersonalitiesServer indicates an expected call of mustEmbedUnimplementedPersonalitiesServer.
func (mr *MockPersonalitiesServerMockRecorder) mustEmbedUnimplementedPersonalitiesServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedPersonalitiesServer", reflect.TypeOf((*MockPersonalitiesServer)(nil).mustEmbedUnimplementedPersonalitiesServer))
}

// MockUnsafePersonalitiesServer is a mock of UnsafePersonalitiesServer interface.
type MockUnsafePersonalitiesServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafePersonalitiesServerMockRecorder
}

// MockUnsafePersonalitiesServerMockRecorder is the mock recorder for MockUnsafePersonalitiesServer.
type MockUnsafePersonalitiesServerMockRecorder struct {
	mock *MockUnsafePersonalitiesServer
}

// NewMockUnsafePersonalitiesServer creates a new mock instance.
func NewMockUnsafePersonalitiesServer(ctrl *gomock.Controller) *MockUnsafePersonalitiesServer {
	mock := &MockUnsafePersonalitiesServer{ctrl: ctrl}
	mock.recorder = &MockUnsafePersonalitiesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafePersonalitiesServer) EXPECT() *MockUnsafePersonalitiesServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedPersonalitiesServer mocks base method.
func (m *MockUnsafePersonalitiesServer) mustEmbedUnimplementedPersonalitiesServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPersonalitiesServer")
}

// mustEmbedUnimplementedPersonalitiesServer indicates an expected call of mustEmbedUnimplementedPersonalitiesServer.
func (mr *MockUnsafePersonalitiesServerMockRecorder) mustEmbedUnimplementedPersonalitiesServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedPersonalitiesServer", reflect.TypeOf((*MockUnsafePersonalitiesServer)(nil).mustEmbedUnimplementedPersonalitiesServer))
}
