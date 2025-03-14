package bindings

import (
	sdkerrors "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/v2/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/palomachain/paloma/v2/util/libwasm"
	bindingstypes "github.com/palomachain/paloma/v2/x/tokenfactory/bindings/types"
	tokenfactorykeeper "github.com/palomachain/paloma/v2/x/tokenfactory/keeper"
	tokenfactorytypes "github.com/palomachain/paloma/v2/x/tokenfactory/types"
)

func NewMessenger(bank *bankkeeper.BaseKeeper, tokenFactory *tokenfactorykeeper.Keeper) libwasm.Messenger[bindingstypes.Message] {
	return &customMessenger{
		bank:         bank,
		tokenFactory: tokenFactory,
	}
}

type customMessenger struct {
	bank         *bankkeeper.BaseKeeper
	tokenFactory *tokenfactorykeeper.Keeper
}

var _ libwasm.Messenger[bindingstypes.Message] = (*customMessenger)(nil)

func (m *customMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, contractMsg bindingstypes.Message) ([]sdk.Event, [][]byte, [][]*codectypes.Any, error) {
	switch {
	case contractMsg.CreateDenom != nil:
		return m.createDenom(ctx, contractAddr, contractMsg.CreateDenom)
	case contractMsg.MintTokens != nil:
		return m.mintTokens(ctx, contractAddr, contractMsg.MintTokens)
	case contractMsg.ChangeAdmin != nil:
		return m.changeAdmin(ctx, contractAddr, contractMsg.ChangeAdmin)
	case contractMsg.BurnTokens != nil:
		return m.burnTokens(ctx, contractAddr, contractMsg.BurnTokens)
	case contractMsg.SetMetadata != nil:
		return m.setMetadata(ctx, contractAddr, contractMsg.SetMetadata)
	}

	return nil, nil, nil, libwasm.ErrUnrecognizedMessage
}

func (m *customMessenger) createDenom(ctx sdk.Context, contractAddr sdk.AccAddress, createDenom *bindingstypes.CreateDenom) ([]sdk.Event, [][]byte, [][]*codectypes.Any, error) {
	bz, err := PerformCreateDenom(m.tokenFactory, m.bank, ctx, contractAddr, createDenom)
	if err != nil {
		return nil, nil, nil, sdkerrors.Wrap(err, "perform create denom")
	}
	return nil, [][]byte{bz}, nil, nil
}

func PerformCreateDenom(f *tokenfactorykeeper.Keeper, b *bankkeeper.BaseKeeper, ctx sdk.Context, contractAddr sdk.AccAddress, createDenom *bindingstypes.CreateDenom) ([]byte, error) {
	if createDenom == nil {
		return nil, wasmvmtypes.InvalidRequest{Err: "create denom null create denom"}
	}

	msgServer := tokenfactorykeeper.NewMsgServerImpl(*f)
	msgCreateDenom := tokenfactorytypes.NewMsgCreateDenom(contractAddr.String(), createDenom.Subdenom)
	if err := msgCreateDenom.ValidateBasic(); err != nil {
		return nil, sdkerrors.Wrap(err, "failed validating MsgCreateDenom")
	}

	resp, err := msgServer.CreateDenom(ctx, msgCreateDenom)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "creating denom")
	}

	if createDenom.Metadata != nil {
		newDenom := resp.NewTokenDenom
		err := PerformSetMetadata(f, b, ctx, contractAddr, newDenom, *createDenom.Metadata)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "setting metadata")
		}
	}

	return resp.Marshal()
}

func (m *customMessenger) mintTokens(ctx sdk.Context, contractAddr sdk.AccAddress, mint *bindingstypes.MintTokens) ([]sdk.Event, [][]byte, [][]*codectypes.Any, error) {
	err := PerformMint(m.tokenFactory, m.bank, ctx, contractAddr, mint)
	if err != nil {
		return nil, nil, nil, sdkerrors.Wrap(err, "perform mint")
	}
	return nil, nil, nil, nil
}

func PerformMint(f *tokenfactorykeeper.Keeper, b *bankkeeper.BaseKeeper, ctx sdk.Context, contractAddr sdk.AccAddress, mint *bindingstypes.MintTokens) error {
	if mint == nil {
		return wasmvmtypes.InvalidRequest{Err: "mint token null mint"}
	}
	rcpt, err := parseAddress(mint.MintToAddress)
	if err != nil {
		return err
	}

	coin := sdk.Coin{Denom: mint.Denom, Amount: mint.Amount}
	sdkMsg := tokenfactorytypes.NewMsgMint(contractAddr.String(), coin)
	if err = sdkMsg.ValidateBasic(); err != nil {
		return err
	}

	msgServer := tokenfactorykeeper.NewMsgServerImpl(*f)
	_, err = msgServer.Mint(ctx, sdkMsg)
	if err != nil {
		return sdkerrors.Wrap(err, "minting coins from message")
	}
	err = b.SendCoins(ctx, contractAddr, rcpt, sdk.NewCoins(coin))
	if err != nil {
		return sdkerrors.Wrap(err, "sending newly minted coins from message")
	}
	return nil
}

func (m *customMessenger) changeAdmin(ctx sdk.Context, contractAddr sdk.AccAddress, changeAdmin *bindingstypes.ChangeAdmin) ([]sdk.Event, [][]byte, [][]*codectypes.Any, error) {
	err := ChangeAdmin(m.tokenFactory, ctx, contractAddr, changeAdmin)
	if err != nil {
		return nil, nil, nil, sdkerrors.Wrap(err, "failed to change admin")
	}
	return nil, nil, nil, nil
}

func ChangeAdmin(f *tokenfactorykeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, changeAdmin *bindingstypes.ChangeAdmin) error {
	if changeAdmin == nil {
		return wasmvmtypes.InvalidRequest{Err: "changeAdmin is nil"}
	}
	newAdminAddr, err := parseAddress(changeAdmin.NewAdminAddress)
	if err != nil {
		return err
	}

	changeAdminMsg := tokenfactorytypes.NewMsgChangeAdmin(contractAddr.String(), changeAdmin.Denom, newAdminAddr.String())
	if err := changeAdminMsg.ValidateBasic(); err != nil {
		return err
	}

	msgServer := tokenfactorykeeper.NewMsgServerImpl(*f)
	_, err = msgServer.ChangeAdmin(ctx, changeAdminMsg)
	if err != nil {
		return sdkerrors.Wrap(err, "failed changing admin from message")
	}
	return nil
}

func (m *customMessenger) burnTokens(ctx sdk.Context, contractAddr sdk.AccAddress, burn *bindingstypes.BurnTokens) ([]sdk.Event, [][]byte, [][]*codectypes.Any, error) {
	err := PerformBurn(m.tokenFactory, ctx, contractAddr, burn)
	if err != nil {
		return nil, nil, nil, sdkerrors.Wrap(err, "perform burn")
	}
	return nil, nil, nil, nil
}

func PerformBurn(f *tokenfactorykeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, burn *bindingstypes.BurnTokens) error {
	if burn == nil {
		return wasmvmtypes.InvalidRequest{Err: "burn token null mint"}
	}
	if burn.BurnFromAddress != "" && burn.BurnFromAddress != contractAddr.String() {
		return wasmvmtypes.InvalidRequest{Err: "BurnFromAddress must be \"\""}
	}

	coin := sdk.Coin{Denom: burn.Denom, Amount: burn.Amount}
	sdkMsg := tokenfactorytypes.NewMsgBurn(contractAddr.String(), coin)
	if err := sdkMsg.ValidateBasic(); err != nil {
		return err
	}

	msgServer := tokenfactorykeeper.NewMsgServerImpl(*f)
	_, err := msgServer.Burn(ctx, sdkMsg)
	if err != nil {
		return sdkerrors.Wrap(err, "burning coins from message")
	}
	return nil
}

func (m *customMessenger) setMetadata(ctx sdk.Context, contractAddr sdk.AccAddress, setMetadata *bindingstypes.SetMetadata) ([]sdk.Event, [][]byte, [][]*codectypes.Any, error) {
	err := PerformSetMetadata(m.tokenFactory, m.bank, ctx, contractAddr, setMetadata.Denom, setMetadata.Metadata)
	if err != nil {
		return nil, nil, nil, sdkerrors.Wrap(err, "perform create denom")
	}
	return nil, nil, nil, nil
}

func PerformSetMetadata(f *tokenfactorykeeper.Keeper, b *bankkeeper.BaseKeeper, ctx sdk.Context, contractAddr sdk.AccAddress, denom string, metadata bindingstypes.Metadata) error {
	auth, err := f.GetAuthorityMetadata(ctx, denom)
	if err != nil {
		return err
	}
	if auth.Admin != contractAddr.String() {
		return wasmvmtypes.InvalidRequest{Err: "only admin can set metadata"}
	}

	// ensure we are setting proper denom metadata (bank uses Base field, fill it if missing)
	if metadata.Base == "" {
		metadata.Base = denom
	} else if metadata.Base != denom {
		// this is the key that we set
		return wasmvmtypes.InvalidRequest{Err: "Base must be the same as denom"}
	}

	bankMetadata := wasmMetadataToSdk(metadata)
	if err := bankMetadata.Validate(); err != nil {
		return err
	}

	b.SetDenomMetaData(ctx, bankMetadata)
	return nil
}

func getFullDenom(contract string, subDenom string) (string, error) {
	if _, err := parseAddress(contract); err != nil {
		return "", err
	}
	fullDenom, err := tokenfactorytypes.GetTokenDenom(contract, subDenom)
	if err != nil {
		return "", sdkerrors.Wrap(err, "validate sub-denom")
	}

	return fullDenom, nil
}

func parseAddress(addr string) (sdk.AccAddress, error) {
	parsed, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "address from bech32")
	}
	err = sdk.VerifyAddressFormat(parsed)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "verify address format")
	}
	return parsed, nil
}

func wasmMetadataToSdk(metadata bindingstypes.Metadata) banktypes.Metadata {
	denoms := []*banktypes.DenomUnit{}
	for _, unit := range metadata.DenomUnits {
		denoms = append(denoms, &banktypes.DenomUnit{
			Denom:    unit.Denom,
			Exponent: unit.Exponent,
			Aliases:  unit.Aliases,
		})
	}
	return banktypes.Metadata{
		Description: metadata.Description,
		Display:     metadata.Display,
		Base:        metadata.Base,
		Name:        metadata.Name,
		Symbol:      metadata.Symbol,
		DenomUnits:  denoms,
	}
}

func sdkMetadataToWasm(metadata banktypes.Metadata) *bindingstypes.Metadata {
	denoms := []bindingstypes.DenomUnit{}
	for _, unit := range metadata.DenomUnits {
		denoms = append(denoms, bindingstypes.DenomUnit{
			Denom:    unit.Denom,
			Exponent: unit.Exponent,
			Aliases:  unit.Aliases,
		})
	}
	return &bindingstypes.Metadata{
		Description: metadata.Description,
		Display:     metadata.Display,
		Base:        metadata.Base,
		Name:        metadata.Name,
		Symbol:      metadata.Symbol,
		DenomUnits:  denoms,
	}
}
