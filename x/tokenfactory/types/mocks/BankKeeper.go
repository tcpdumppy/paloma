// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper is an autogenerated mock type for the BankKeeper type
type BankKeeper struct {
	mock.Mock
}

// BurnCoins provides a mock function with given fields: ctx, moduleName, amt
func (_m *BankKeeper) BurnCoins(ctx context.Context, moduleName string, amt types.Coins) error {
	ret := _m.Called(ctx, moduleName, amt)

	if len(ret) == 0 {
		panic("no return value specified for BurnCoins")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.Coins) error); ok {
		r0 = rf(ctx, moduleName, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllBalances provides a mock function with given fields: ctx, addr
func (_m *BankKeeper) GetAllBalances(ctx context.Context, addr types.AccAddress) types.Coins {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for GetAllBalances")
	}

	var r0 types.Coins
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress) types.Coins); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Coins)
		}
	}

	return r0
}

// GetBalance provides a mock function with given fields: ctx, addr, denom
func (_m *BankKeeper) GetBalance(ctx context.Context, addr types.AccAddress, denom string) types.Coin {
	ret := _m.Called(ctx, addr, denom)

	if len(ret) == 0 {
		panic("no return value specified for GetBalance")
	}

	var r0 types.Coin
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, string) types.Coin); ok {
		r0 = rf(ctx, addr, denom)
	} else {
		r0 = ret.Get(0).(types.Coin)
	}

	return r0
}

// GetDenomMetaData provides a mock function with given fields: ctx, denom
func (_m *BankKeeper) GetDenomMetaData(ctx context.Context, denom string) (banktypes.Metadata, bool) {
	ret := _m.Called(ctx, denom)

	if len(ret) == 0 {
		panic("no return value specified for GetDenomMetaData")
	}

	var r0 banktypes.Metadata
	var r1 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) (banktypes.Metadata, bool)); ok {
		return rf(ctx, denom)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) banktypes.Metadata); ok {
		r0 = rf(ctx, denom)
	} else {
		r0 = ret.Get(0).(banktypes.Metadata)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) bool); ok {
		r1 = rf(ctx, denom)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// HasBalance provides a mock function with given fields: ctx, addr, amt
func (_m *BankKeeper) HasBalance(ctx context.Context, addr types.AccAddress, amt types.Coin) bool {
	ret := _m.Called(ctx, addr, amt)

	if len(ret) == 0 {
		panic("no return value specified for HasBalance")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, types.Coin) bool); ok {
		r0 = rf(ctx, addr, amt)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HasSupply provides a mock function with given fields: ctx, denom
func (_m *BankKeeper) HasSupply(ctx context.Context, denom string) bool {
	ret := _m.Called(ctx, denom)

	if len(ret) == 0 {
		panic("no return value specified for HasSupply")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, denom)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MintCoins provides a mock function with given fields: ctx, moduleName, amt
func (_m *BankKeeper) MintCoins(ctx context.Context, moduleName string, amt types.Coins) error {
	ret := _m.Called(ctx, moduleName, amt)

	if len(ret) == 0 {
		panic("no return value specified for MintCoins")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.Coins) error); ok {
		r0 = rf(ctx, moduleName, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendCoins provides a mock function with given fields: ctx, fromAddr, toAddr, amt
func (_m *BankKeeper) SendCoins(ctx context.Context, fromAddr types.AccAddress, toAddr types.AccAddress, amt types.Coins) error {
	ret := _m.Called(ctx, fromAddr, toAddr, amt)

	if len(ret) == 0 {
		panic("no return value specified for SendCoins")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, types.AccAddress, types.Coins) error); ok {
		r0 = rf(ctx, fromAddr, toAddr, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendCoinsFromAccountToModule provides a mock function with given fields: ctx, senderAddr, recipientModule, amt
func (_m *BankKeeper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	ret := _m.Called(ctx, senderAddr, recipientModule, amt)

	if len(ret) == 0 {
		panic("no return value specified for SendCoinsFromAccountToModule")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, string, types.Coins) error); ok {
		r0 = rf(ctx, senderAddr, recipientModule, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendCoinsFromModuleToAccount provides a mock function with given fields: ctx, senderModule, recipientAddr, amt
func (_m *BankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	ret := _m.Called(ctx, senderModule, recipientAddr, amt)

	if len(ret) == 0 {
		panic("no return value specified for SendCoinsFromModuleToAccount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.AccAddress, types.Coins) error); ok {
		r0 = rf(ctx, senderModule, recipientAddr, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDenomMetaData provides a mock function with given fields: ctx, denomMetaData
func (_m *BankKeeper) SetDenomMetaData(ctx context.Context, denomMetaData banktypes.Metadata) {
	_m.Called(ctx, denomMetaData)
}

// SpendableCoins provides a mock function with given fields: ctx, addr
func (_m *BankKeeper) SpendableCoins(ctx context.Context, addr types.AccAddress) types.Coins {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for SpendableCoins")
	}

	var r0 types.Coins
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress) types.Coins); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Coins)
		}
	}

	return r0
}

// NewBankKeeper creates a new instance of BankKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBankKeeper(t interface {
	mock.TestingT
	Cleanup(func())
},
) *BankKeeper {
	mock := &BankKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
