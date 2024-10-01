package metrix

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/palomachain/paloma/v2/x/metrix/keeper"
	"github.com/palomachain/paloma/v2/x/metrix/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	return genesis
}
