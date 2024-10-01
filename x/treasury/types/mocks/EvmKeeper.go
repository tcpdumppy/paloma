// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/palomachain/paloma/v2/x/evm/types"
)

// EvmKeeper is an autogenerated mock type for the EvmKeeper type
type EvmKeeper struct {
	mock.Mock
}

// GetChainInfo provides a mock function with given fields: ctx, targetChainReferenceID
func (_m *EvmKeeper) GetChainInfo(ctx context.Context, targetChainReferenceID string) (*types.ChainInfo, error) {
	ret := _m.Called(ctx, targetChainReferenceID)

	if len(ret) == 0 {
		panic("no return value specified for GetChainInfo")
	}

	var r0 *types.ChainInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*types.ChainInfo, error)); ok {
		return rf(ctx, targetChainReferenceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.ChainInfo); ok {
		r0 = rf(ctx, targetChainReferenceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ChainInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, targetChainReferenceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEvmKeeper creates a new instance of EvmKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEvmKeeper(t interface {
	mock.TestingT
	Cleanup(func())
},
) *EvmKeeper {
	mock := &EvmKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
