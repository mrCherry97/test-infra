// Code generated by mockery v2.38.0. DO NOT EDIT.

package oidcmocks

import (
	jwt "github.com/go-jose/go-jose/v4/jwt"
	mock "github.com/stretchr/testify/mock"
)

// MockClaimsInterface is an autogenerated mock type for the ClaimsInterface type
type MockClaimsInterface struct {
	mock.Mock
}

type MockClaimsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClaimsInterface) EXPECT() *MockClaimsInterface_Expecter {
	return &MockClaimsInterface_Expecter{mock: &_m.Mock}
}

// Validate provides a mock function with given fields: e
func (_m *MockClaimsInterface) Validate(e jwt.Expected) error {
	ret := _m.Called(e)

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(jwt.Expected) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClaimsInterface_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type MockClaimsInterface_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - e jwt.Expected
func (_e *MockClaimsInterface_Expecter) Validate(e interface{}) *MockClaimsInterface_Validate_Call {
	return &MockClaimsInterface_Validate_Call{Call: _e.mock.On("Validate", e)}
}

func (_c *MockClaimsInterface_Validate_Call) Run(run func(e jwt.Expected)) *MockClaimsInterface_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(jwt.Expected))
	})
	return _c
}

func (_c *MockClaimsInterface_Validate_Call) Return(_a0 error) *MockClaimsInterface_Validate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClaimsInterface_Validate_Call) RunAndReturn(run func(jwt.Expected) error) *MockClaimsInterface_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClaimsInterface creates a new instance of MockClaimsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClaimsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClaimsInterface {
	mock := &MockClaimsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}