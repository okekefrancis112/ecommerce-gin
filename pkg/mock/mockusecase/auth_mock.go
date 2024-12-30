// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interfaces/auth.go

// Package mockusecase is a generated GoMock package.
package mockusecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	request "github.com/okekefrancis112/ecommerce-gin/pkg/api/handler/request"
	domain "github.com/okekefrancis112/ecommerce-gin/pkg/domain"
	token "github.com/okekefrancis112/ecommerce-gin/pkg/service/token"
	interfaces "github.com/okekefrancis112/ecommerce-gin/pkg/usecase/interfaces"
)

// MockAuthUseCase is a mock of AuthUseCase interface.
type MockAuthUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockAuthUseCaseMockRecorder
}

// MockAuthUseCaseMockRecorder is the mock recorder for MockAuthUseCase.
type MockAuthUseCaseMockRecorder struct {
	mock *MockAuthUseCase
}

// NewMockAuthUseCase creates a new mock instance.
func NewMockAuthUseCase(ctrl *gomock.Controller) *MockAuthUseCase {
	mock := &MockAuthUseCase{ctrl: ctrl}
	mock.recorder = &MockAuthUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthUseCase) EXPECT() *MockAuthUseCaseMockRecorder {
	return m.recorder
}

// AdminLogin mocks base method.
func (m *MockAuthUseCase) AdminLogin(ctx context.Context, loginDetails request.Login) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdminLogin", ctx, loginDetails)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AdminLogin indicates an expected call of AdminLogin.
func (mr *MockAuthUseCaseMockRecorder) AdminLogin(ctx, loginDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminLogin", reflect.TypeOf((*MockAuthUseCase)(nil).AdminLogin), ctx, loginDetails)
}

// GenerateAccessToken mocks base method.
func (m *MockAuthUseCase) GenerateAccessToken(ctx context.Context, tokenParams interfaces.GenerateTokenParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", ctx, tokenParams)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockAuthUseCaseMockRecorder) GenerateAccessToken(ctx, tokenParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockAuthUseCase)(nil).GenerateAccessToken), ctx, tokenParams)
}

// GenerateRefreshToken mocks base method.
func (m *MockAuthUseCase) GenerateRefreshToken(ctx context.Context, tokenParams interfaces.GenerateTokenParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", ctx, tokenParams)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockAuthUseCaseMockRecorder) GenerateRefreshToken(ctx, tokenParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockAuthUseCase)(nil).GenerateRefreshToken), ctx, tokenParams)
}

// GoogleLogin mocks base method.
func (m *MockAuthUseCase) GoogleLogin(ctx context.Context, user domain.User) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GoogleLogin", ctx, user)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GoogleLogin indicates an expected call of GoogleLogin.
func (mr *MockAuthUseCaseMockRecorder) GoogleLogin(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GoogleLogin", reflect.TypeOf((*MockAuthUseCase)(nil).GoogleLogin), ctx, user)
}

// LoginOtpVerify mocks base method.
func (m *MockAuthUseCase) LoginOtpVerify(ctx context.Context, otpVerifyDetails request.OTPVerify) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginOtpVerify", ctx, otpVerifyDetails)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginOtpVerify indicates an expected call of LoginOtpVerify.
func (mr *MockAuthUseCaseMockRecorder) LoginOtpVerify(ctx, otpVerifyDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginOtpVerify", reflect.TypeOf((*MockAuthUseCase)(nil).LoginOtpVerify), ctx, otpVerifyDetails)
}

// SingUpOtpVerify mocks base method.
func (m *MockAuthUseCase) SingUpOtpVerify(ctx context.Context, otpVerifyDetails request.OTPVerify) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SingUpOtpVerify", ctx, otpVerifyDetails)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SingUpOtpVerify indicates an expected call of SingUpOtpVerify.
func (mr *MockAuthUseCaseMockRecorder) SingUpOtpVerify(ctx, otpVerifyDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SingUpOtpVerify", reflect.TypeOf((*MockAuthUseCase)(nil).SingUpOtpVerify), ctx, otpVerifyDetails)
}

// UserLogin mocks base method.
func (m *MockAuthUseCase) UserLogin(ctx context.Context, loginDetails request.Login) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserLogin", ctx, loginDetails)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLogin indicates an expected call of UserLogin.
func (mr *MockAuthUseCaseMockRecorder) UserLogin(ctx, loginDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLogin", reflect.TypeOf((*MockAuthUseCase)(nil).UserLogin), ctx, loginDetails)
}

// UserLoginOtpSend mocks base method.
func (m *MockAuthUseCase) UserLoginOtpSend(ctx context.Context, loginDetails request.OTPLogin) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserLoginOtpSend", ctx, loginDetails)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLoginOtpSend indicates an expected call of UserLoginOtpSend.
func (mr *MockAuthUseCaseMockRecorder) UserLoginOtpSend(ctx, loginDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLoginOtpSend", reflect.TypeOf((*MockAuthUseCase)(nil).UserLoginOtpSend), ctx, loginDetails)
}

// UserSignUp mocks base method.
func (m *MockAuthUseCase) UserSignUp(ctx context.Context, signUpDetails domain.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", ctx, signUpDetails)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockAuthUseCaseMockRecorder) UserSignUp(ctx, signUpDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockAuthUseCase)(nil).UserSignUp), ctx, signUpDetails)
}

// VerifyAndGetRefreshTokenSession mocks base method.
func (m *MockAuthUseCase) VerifyAndGetRefreshTokenSession(ctx context.Context, refreshToken string, usedFor token.UserType) (domain.RefreshSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyAndGetRefreshTokenSession", ctx, refreshToken, usedFor)
	ret0, _ := ret[0].(domain.RefreshSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyAndGetRefreshTokenSession indicates an expected call of VerifyAndGetRefreshTokenSession.
func (mr *MockAuthUseCaseMockRecorder) VerifyAndGetRefreshTokenSession(ctx, refreshToken, usedFor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAndGetRefreshTokenSession", reflect.TypeOf((*MockAuthUseCase)(nil).VerifyAndGetRefreshTokenSession), ctx, refreshToken, usedFor)
}
