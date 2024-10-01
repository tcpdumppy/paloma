// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/palomachain/paloma/v2/x/evm/types"
	mock "github.com/stretchr/testify/mock"
)

// SkywayKeeper is an autogenerated mock type for the SkywayKeeper type
type SkywayKeeper struct {
	mock.Mock
}

// CastAllERC20ToDenoms provides a mock function with given fields: ctx
func (_m *SkywayKeeper) CastAllERC20ToDenoms(ctx context.Context) ([]types.ERC20Record, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CastAllERC20ToDenoms")
	}

	var r0 []types.ERC20Record
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]types.ERC20Record, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []types.ERC20Record); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.ERC20Record)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CastChainERC20ToDenoms provides a mock function with given fields: ctx, chainReferenceID
func (_m *SkywayKeeper) CastChainERC20ToDenoms(ctx context.Context, chainReferenceID string) ([]types.ERC20Record, error) {
	ret := _m.Called(ctx, chainReferenceID)

	if len(ret) == 0 {
		panic("no return value specified for CastChainERC20ToDenoms")
	}

	var r0 []types.ERC20Record
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]types.ERC20Record, error)); ok {
		return rf(ctx, chainReferenceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []types.ERC20Record); ok {
		r0 = rf(ctx, chainReferenceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.ERC20Record)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, chainReferenceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastObservedSkywayNonce provides a mock function with given fields: ctx, chainReferenceID
func (_m *SkywayKeeper) GetLastObservedSkywayNonce(ctx context.Context, chainReferenceID string) (uint64, error) {
	ret := _m.Called(ctx, chainReferenceID)

	if len(ret) == 0 {
		panic("no return value specified for GetLastObservedSkywayNonce")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (uint64, error)); ok {
		return rf(ctx, chainReferenceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) uint64); ok {
		r0 = rf(ctx, chainReferenceID)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, chainReferenceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSkywayKeeper creates a new instance of SkywayKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSkywayKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *SkywayKeeper {
	mock := &SkywayKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
