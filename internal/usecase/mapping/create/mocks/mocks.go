// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package test

import (
	"context"

	"github.com/enesanbar/url-shortener/internal/domain"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/create"
	mock "github.com/stretchr/testify/mock"
)

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// FindByCode provides a mock function for the type MockRepository
func (_mock *MockRepository) FindByCode(ctx context.Context, code string) (*domain.Mapping, error) {
	ret := _mock.Called(ctx, code)

	if len(ret) == 0 {
		panic("no return value specified for FindByCode")
	}

	var r0 *domain.Mapping
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) (*domain.Mapping, error)); ok {
		return returnFunc(ctx, code)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) *domain.Mapping); ok {
		r0 = returnFunc(ctx, code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Mapping)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = returnFunc(ctx, code)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockRepository_FindByCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByCode'
type MockRepository_FindByCode_Call struct {
	*mock.Call
}

// FindByCode is a helper method to define mock.On call
//   - ctx
//   - code
func (_e *MockRepository_Expecter) FindByCode(ctx interface{}, code interface{}) *MockRepository_FindByCode_Call {
	return &MockRepository_FindByCode_Call{Call: _e.mock.On("FindByCode", ctx, code)}
}

func (_c *MockRepository_FindByCode_Call) Run(run func(ctx context.Context, code string)) *MockRepository_FindByCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_FindByCode_Call) Return(mapping *domain.Mapping, err error) *MockRepository_FindByCode_Call {
	_c.Call.Return(mapping, err)
	return _c
}

func (_c *MockRepository_FindByCode_Call) RunAndReturn(run func(ctx context.Context, code string) (*domain.Mapping, error)) *MockRepository_FindByCode_Call {
	_c.Call.Return(run)
	return _c
}

// Store provides a mock function for the type MockRepository
func (_mock *MockRepository) Store(ctx context.Context, mapping *domain.Mapping) (*domain.Mapping, error) {
	ret := _mock.Called(ctx, mapping)

	if len(ret) == 0 {
		panic("no return value specified for Store")
	}

	var r0 *domain.Mapping
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *domain.Mapping) (*domain.Mapping, error)); ok {
		return returnFunc(ctx, mapping)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *domain.Mapping) *domain.Mapping); ok {
		r0 = returnFunc(ctx, mapping)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Mapping)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *domain.Mapping) error); ok {
		r1 = returnFunc(ctx, mapping)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockRepository_Store_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Store'
type MockRepository_Store_Call struct {
	*mock.Call
}

// Store is a helper method to define mock.On call
//   - ctx
//   - mapping
func (_e *MockRepository_Expecter) Store(ctx interface{}, mapping interface{}) *MockRepository_Store_Call {
	return &MockRepository_Store_Call{Call: _e.mock.On("Store", ctx, mapping)}
}

func (_c *MockRepository_Store_Call) Run(run func(ctx context.Context, mapping *domain.Mapping)) *MockRepository_Store_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Mapping))
	})
	return _c
}

func (_c *MockRepository_Store_Call) Return(mapping1 *domain.Mapping, err error) *MockRepository_Store_Call {
	_c.Call.Return(mapping1, err)
	return _c
}

func (_c *MockRepository_Store_Call) RunAndReturn(run func(ctx context.Context, mapping *domain.Mapping) (*domain.Mapping, error)) *MockRepository_Store_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

type MockService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockService) EXPECT() *MockService_Expecter {
	return &MockService_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function for the type MockService
func (_mock *MockService) Execute(ctx context.Context, input *create.Request) (*domain.Mapping, error) {
	ret := _mock.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *domain.Mapping
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *create.Request) (*domain.Mapping, error)); ok {
		return returnFunc(ctx, input)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *create.Request) *domain.Mapping); ok {
		r0 = returnFunc(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Mapping)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *create.Request) error); ok {
		r1 = returnFunc(ctx, input)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockService_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockService_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx
//   - input
func (_e *MockService_Expecter) Execute(ctx interface{}, input interface{}) *MockService_Execute_Call {
	return &MockService_Execute_Call{Call: _e.mock.On("Execute", ctx, input)}
}

func (_c *MockService_Execute_Call) Run(run func(ctx context.Context, input *create.Request)) *MockService_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*create.Request))
	})
	return _c
}

func (_c *MockService_Execute_Call) Return(mapping *domain.Mapping, err error) *MockService_Execute_Call {
	_c.Call.Return(mapping, err)
	return _c
}

func (_c *MockService_Execute_Call) RunAndReturn(run func(ctx context.Context, input *create.Request) (*domain.Mapping, error)) *MockService_Execute_Call {
	_c.Call.Return(run)
	return _c
}
