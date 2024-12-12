package keeper

import (
	"github.com/archway-network/archway/x/photosynthesis/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the Photosynthesis module's state at genesis.
func (k PhotosynthesisKeeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) {
	k.SetParams(ctx, data.Params)

	// Initialize deposit records
	for _, record := range data.DepositRecords {
		contractAddr, err := sdk.AccAddressFromBech32(record.ContractAddress)
		if err != nil {
			panic(err) // Handle errors more gracefully in a real scenario
		}
		k.SetContractLiquidStakeDeposits(ctx, contractAddr.String(), []*types.DepositRecord{&record}) // Storing as pointers
	}

}

// ExportGenesis exports the Photosynthesis module's state for genesis.
func (k PhotosynthesisKeeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)

	// Collect deposit records (convert pointers to values)
	var depositRecords []types.DepositRecord
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.LiquidStakeRecordsPrefix))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var depositRecord types.DepositRecord
		k.cdc.MustUnmarshal(iterator.Value(), &depositRecord)
		depositRecords = append(depositRecords, depositRecord) // Collect as values
	}

	// Collect contracts
	var contracts []types.Contract
	contractIterator := sdk.KVStorePrefixIterator(store, []byte(types.ContractPrefix))
	defer contractIterator.Close()

	for ; contractIterator.Valid(); contractIterator.Next() {
		var contract types.Contract
		k.cdc.MustUnmarshal(contractIterator.Value(), &contract)
		contracts = append(contracts, contract)
	}

	return &types.GenesisState{
		Params:         params,
		DepositRecords: depositRecords, // Use slice of values, not pointers
		Contracts:      contracts,
	}
}
