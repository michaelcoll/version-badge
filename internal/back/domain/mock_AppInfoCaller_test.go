// Code generated by mockery v2.20.0. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockAppInfoCaller is an autogenerated mock type for the AppInfoCaller type
type MockAppInfoCaller struct {
	mock.Mock
}

type MockAppInfoCaller_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAppInfoCaller) EXPECT() *MockAppInfoCaller_Expecter {
	return &MockAppInfoCaller_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, env, appName
func (_m *MockAppInfoCaller) Get(ctx context.Context, env string, appName string) (*AppInfo, error) {
	ret := _m.Called(ctx, env, appName)

	var r0 *AppInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*AppInfo, error)); ok {
		return rf(ctx, env, appName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *AppInfo); ok {
		r0 = rf(ctx, env, appName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AppInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, env, appName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAppInfoCaller_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAppInfoCaller_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - env string
//   - appName string
func (_e *MockAppInfoCaller_Expecter) Get(ctx interface{}, env interface{}, appName interface{}) *MockAppInfoCaller_Get_Call {
	return &MockAppInfoCaller_Get_Call{Call: _e.mock.On("Get", ctx, env, appName)}
}

func (_c *MockAppInfoCaller_Get_Call) Run(run func(ctx context.Context, env string, appName string)) *MockAppInfoCaller_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockAppInfoCaller_Get_Call) Return(_a0 *AppInfo, _a1 error) *MockAppInfoCaller_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAppInfoCaller_Get_Call) RunAndReturn(run func(context.Context, string, string) (*AppInfo, error)) *MockAppInfoCaller_Get_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockAppInfoCaller interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAppInfoCaller creates a new instance of MockAppInfoCaller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAppInfoCaller(t mockConstructorTestingTNewMockAppInfoCaller) *MockAppInfoCaller {
	mock := &MockAppInfoCaller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
