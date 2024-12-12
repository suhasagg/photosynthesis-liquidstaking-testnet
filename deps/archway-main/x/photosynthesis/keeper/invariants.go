package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/crisis/types"
)

// RegisterInvariants registers the Photosynthesis module invariants.
func RegisterInvariants(ir sdk.InvariantRegistry, k PhotosynthesisKeeper) {
	ir.RegisterRoute(types.ModuleName, "liquid-stake", LiquidStakeInvariant(k))
	ir.RegisterRoute(types.ModuleName, "reward-records", RewardRecordsInvariant(k))
}

// LiquidStakeInvariant checks that the liquid stake values are consistent.
func LiquidStakeInvariant(k PhotosynthesisKeeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		totalLiquidStake, err := k.GetTotalLiquidStake(ctx)
		if err != nil {
			return fmt.Sprintf("error retrieving total liquid stake: %v", err), true
		}

		if totalLiquidStake.IsNegative() {
			return fmt.Sprintf("total liquid stake is negative: %v", totalLiquidStake), true
		}

		return "All liquid stake amounts are valid", false
	}
}

// RewardRecordsInvariant checks that the reward records are consistent.
func RewardRecordsInvariant(k PhotosynthesisKeeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		rewardMap, _ := k.GetTotalRewardRecords(ctx)

		for address, amount := range rewardMap {
			if amount < 0 {
				return fmt.Sprintf("Negative reward amount for address %s: %d", address, amount), true
			}
		}

		// Further checks can be added here (e.g., ensuring rewards and records are balanced).

		return "All reward records are valid", false
	}
}
