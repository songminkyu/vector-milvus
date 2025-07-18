// Code generated by mockery v2.53.3. DO NOT EDIT.

package mock_streaming

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/milvus-io/milvus/pkg/v2/streaming/util/types"
)

// MockLocal is an autogenerated mock type for the Local type
type MockLocal struct {
	mock.Mock
}

type MockLocal_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLocal) EXPECT() *MockLocal_Expecter {
	return &MockLocal_Expecter{mock: &_m.Mock}
}

// GetLatestMVCCTimestampIfLocal provides a mock function with given fields: ctx, vchannel
func (_m *MockLocal) GetLatestMVCCTimestampIfLocal(ctx context.Context, vchannel string) (uint64, error) {
	ret := _m.Called(ctx, vchannel)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestMVCCTimestampIfLocal")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (uint64, error)); ok {
		return rf(ctx, vchannel)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) uint64); ok {
		r0 = rf(ctx, vchannel)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, vchannel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLocal_GetLatestMVCCTimestampIfLocal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestMVCCTimestampIfLocal'
type MockLocal_GetLatestMVCCTimestampIfLocal_Call struct {
	*mock.Call
}

// GetLatestMVCCTimestampIfLocal is a helper method to define mock.On call
//   - ctx context.Context
//   - vchannel string
func (_e *MockLocal_Expecter) GetLatestMVCCTimestampIfLocal(ctx interface{}, vchannel interface{}) *MockLocal_GetLatestMVCCTimestampIfLocal_Call {
	return &MockLocal_GetLatestMVCCTimestampIfLocal_Call{Call: _e.mock.On("GetLatestMVCCTimestampIfLocal", ctx, vchannel)}
}

func (_c *MockLocal_GetLatestMVCCTimestampIfLocal_Call) Run(run func(ctx context.Context, vchannel string)) *MockLocal_GetLatestMVCCTimestampIfLocal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockLocal_GetLatestMVCCTimestampIfLocal_Call) Return(_a0 uint64, _a1 error) *MockLocal_GetLatestMVCCTimestampIfLocal_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLocal_GetLatestMVCCTimestampIfLocal_Call) RunAndReturn(run func(context.Context, string) (uint64, error)) *MockLocal_GetLatestMVCCTimestampIfLocal_Call {
	_c.Call.Return(run)
	return _c
}

// GetMetricsIfLocal provides a mock function with given fields: ctx
func (_m *MockLocal) GetMetricsIfLocal(ctx context.Context) (*types.StreamingNodeMetrics, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetMetricsIfLocal")
	}

	var r0 *types.StreamingNodeMetrics
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*types.StreamingNodeMetrics, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *types.StreamingNodeMetrics); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.StreamingNodeMetrics)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLocal_GetMetricsIfLocal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMetricsIfLocal'
type MockLocal_GetMetricsIfLocal_Call struct {
	*mock.Call
}

// GetMetricsIfLocal is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockLocal_Expecter) GetMetricsIfLocal(ctx interface{}) *MockLocal_GetMetricsIfLocal_Call {
	return &MockLocal_GetMetricsIfLocal_Call{Call: _e.mock.On("GetMetricsIfLocal", ctx)}
}

func (_c *MockLocal_GetMetricsIfLocal_Call) Run(run func(ctx context.Context)) *MockLocal_GetMetricsIfLocal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockLocal_GetMetricsIfLocal_Call) Return(_a0 *types.StreamingNodeMetrics, _a1 error) *MockLocal_GetMetricsIfLocal_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLocal_GetMetricsIfLocal_Call) RunAndReturn(run func(context.Context) (*types.StreamingNodeMetrics, error)) *MockLocal_GetMetricsIfLocal_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLocal creates a new instance of MockLocal. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLocal(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLocal {
	mock := &MockLocal{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
