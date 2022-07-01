// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	consensustypes "github.com/palomachain/paloma/x/consensus/types"

	mock "github.com/stretchr/testify/mock"

	testing "testing"

	types "github.com/cosmos/cosmos-sdk/types"
)

// ConsensusKeeper is an autogenerated mock type for the ConsensusKeeper type
type ConsensusKeeper struct {
	mock.Mock
}

// PutMessageForSigning provides a mock function with given fields: ctx, queueTypeName, msg
func (_m *ConsensusKeeper) PutMessageForSigning(ctx types.Context, queueTypeName string, msg consensustypes.ConsensusMsg) error {
	ret := _m.Called(ctx, queueTypeName, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, string, consensustypes.ConsensusMsg) error); ok {
		r0 = rf(ctx, queueTypeName, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveConsensusQueue provides a mock function with given fields: ctx, queueTypeName
func (_m *ConsensusKeeper) RemoveConsensusQueue(ctx types.Context, queueTypeName string) error {
	ret := _m.Called(ctx, queueTypeName)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, string) error); ok {
		r0 = rf(ctx, queueTypeName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewConsensusKeeper creates a new instance of ConsensusKeeper. It also registers a cleanup function to assert the mocks expectations.
func NewConsensusKeeper(t testing.TB) *ConsensusKeeper {
	mock := &ConsensusKeeper{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
