package keeper

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/archway-network/archway/x/photosynthesis/types"
	rewardkeeper "github.com/archway-network/archway/x/rewards/keeper"
	rewardstypes "github.com/archway-network/archway/x/rewards/types"

	"github.com/tendermint/tendermint/libs/log"
)

// Task keys for tracking last processed times
const (
	LastLiquidStakingDAppRewardsTimeKey = "last_liquid_staking_dapp_rewards_time"
	LastArchLiquidStakeIntervalTimeKey  = "last_arch_liquid_stake_interval_time"
	LastRedemptionRateQueryTimeKey      = "last_redemption_rate_query_time"
	LastRewardsWithdrawalTimeKey        = "last_rewards_withdrawal_time"
	RedeemTokenPrefix                   = "redeem_token"
	RedeemTokenRatioPrefix              = "redeem_token_ratio"
)

// PhotosynthesisKeeper manages the state and provides methods for the Photosynthesis module.
type PhotosynthesisKeeper struct {
	storeKey     storetypes.StoreKey
	cdc          codec.BinaryCodec
	paramStore   paramtypes.Subspace
	bankKeeper   bankkeeper.Keeper
	rewardKeeper rewardkeeper.Keeper
	logger       log.Logger
}

// NewPhotosynthesisKeeper creates a new PhotosynthesisKeeper.
func NewPhotosynthesisKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	paramStore paramtypes.Subspace,
	bankKeeper bankkeeper.Keeper,
	rewardKeeper rewardkeeper.Keeper,
	logger log.Logger,
) PhotosynthesisKeeper {
	return PhotosynthesisKeeper{
		storeKey:     storeKey,
		cdc:          cdc,
		paramStore:   paramStore.WithKeyTable(types.ParamKeyTable()),
		bankKeeper:   bankKeeper,
		rewardKeeper: rewardKeeper,
		logger:       logger.With("module", "x/photosynthesis"),
	}
}

// ---- Set/Get Last Processed Time in KVStore ----

// SetLastProcessingTime sets the last processed time for a task.
func (k PhotosynthesisKeeper) SetLastProcessingTime(ctx sdk.Context, key string, t time.Time) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(t.Unix()))
	store.Set([]byte(key), bz)
}

// GetLastProcessingTime retrieves the last processed time for a task.
func (k PhotosynthesisKeeper) GetLastProcessingTime(ctx sdk.Context, key string) (time.Time, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(key))
	if bz == nil {
		return time.Time{}, fmt.Errorf("last processing time not found for key: %s", key)
	}
	lastTimeUnix := int64(binary.BigEndian.Uint64(bz))
	return time.Unix(lastTimeUnix, 0), nil
}

// ShouldProcessTask checks if enough time has passed since the last execution of a specific task.
func (k PhotosynthesisKeeper) ShouldProcessTask(ctx sdk.Context, key string, intervalInMinutes int64) bool {
	lastProcessedTime, err := k.GetLastProcessingTime(ctx, key)
	if err != nil {
		// If no last processed time is found, process the task.
		return true
	}
	currentTime := ctx.BlockTime()
	timeSinceLastProcess := currentTime.Sub(lastProcessedTime).Minutes()

	// If more than intervalInMinutes have passed, it's time to process the task.
	return timeSinceLastProcess >= float64(intervalInMinutes)
}

// ---- Params for Intervals ----

// GetParams retrieves the module parameters.
func (k PhotosynthesisKeeper) GetParams(ctx sdk.Context) types.Params {
	var params types.Params
	k.paramStore.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the module parameters.
func (k PhotosynthesisKeeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramStore.SetParamSet(ctx, &params)
}

// ---- Logger ----

// Logger returns a module-specific logger.
func (k PhotosynthesisKeeper) Logger(ctx sdk.Context) log.Logger {
	return k.logger
}

// ---- Task Handlers ----

func (k PhotosynthesisKeeper) HandleLiquidStakingDAppRewardsTask(ctx sdk.Context) error {
	params := k.GetParams(ctx)

	if !k.ShouldProcessTask(ctx, LastLiquidStakingDAppRewardsTimeKey, params.LiquidStakingInterval) {
		k.Logger(ctx).Info("LiquidStakingDAppRewardsTask skipped as interval has not passed.")
		return nil
	}

	k.Logger(ctx).Info("LiquidStakingDAppRewardsTask processing started.")
	// Call the handler function for Liquid Staking DApp Rewards
	k.handleLiquidStakingDAppRewards(ctx)

	// Set the last processed time after completion
	k.SetLastProcessingTime(ctx, LastLiquidStakingDAppRewardsTimeKey, ctx.BlockTime())
	k.Logger(ctx).Info("LiquidStakingDAppRewardsTask processing completed.")
	return nil
}

func (k PhotosynthesisKeeper) HandleArchLiquidStakeIntervalTask(ctx sdk.Context) error {
	params := k.GetParams(ctx)

	if !k.ShouldProcessTask(ctx, LastArchLiquidStakeIntervalTimeKey, params.ArchLiquidStakeInterval) {
		k.Logger(ctx).Info("ArchLiquidStakeIntervalTask skipped as interval has not passed.")
		return nil
	}

	k.Logger(ctx).Info("ArchLiquidStakeIntervalTask processing started.")
	// Call the handler function for Arch Liquid Stake Interval
	k.handleArchLiquidStakeInterval(ctx)

	// Set the last processed time after completion
	k.SetLastProcessingTime(ctx, LastArchLiquidStakeIntervalTimeKey, ctx.BlockTime())
	k.Logger(ctx).Info("ArchLiquidStakeIntervalTask processing completed.")
	return nil
}

func (k PhotosynthesisKeeper) HandleRedemptionRateQueryTask(ctx sdk.Context) error {
	params := k.GetParams(ctx)

	if !k.ShouldProcessTask(ctx, LastRedemptionRateQueryTimeKey, params.RedemptionRateQueryInterval) {
		k.Logger(ctx).Info("RedemptionRateQueryTask skipped as interval has not passed.")
		return nil
	}

	k.Logger(ctx).Info("RedemptionRateQueryTask processing started.")
	// Call the handler function for Redemption Rate Query
	k.handleRedemptionRateQuery(ctx)

	// Set the last processed time after completion
	k.SetLastProcessingTime(ctx, LastRedemptionRateQueryTimeKey, ctx.BlockTime())
	k.Logger(ctx).Info("RedemptionRateQueryTask processing completed.")
	return nil
}

func (k PhotosynthesisKeeper) HandleRewardsWithdrawalTask(ctx sdk.Context) error {
	params := k.GetParams(ctx)

	if !k.ShouldProcessTask(ctx, LastRewardsWithdrawalTimeKey, params.RewardsWithdrawalInterval) {
		k.Logger(ctx).Info("RewardsWithdrawalTask skipped as interval has not passed.")
		return nil
	}

	k.Logger(ctx).Info("RewardsWithdrawalTask processing started.")
	// Call the handler function for Rewards Withdrawal
	k.handleRewardsWithdrawal(ctx)

	// Set the last processed time after completion
	k.SetLastProcessingTime(ctx, LastRewardsWithdrawalTimeKey, ctx.BlockTime())
	k.Logger(ctx).Info("RewardsWithdrawalTask processing completed.")
	return nil
}

// ---- Cron Job Trigger Function ----

// CronJobTriggerTasks is the external-facing method to be triggered by a cron job.
func (k PhotosynthesisKeeper) CronJobTriggerTasks(ctx sdk.Context) error {
	k.Logger(ctx).Info("Cron job triggered for task processing.")

	// Process each task based on the intervals
	if err := k.HandleLiquidStakingDAppRewardsTask(ctx); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error in LiquidStakingDAppRewardsTask: %v", err))
	}

	if err := k.HandleArchLiquidStakeIntervalTask(ctx); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error in ArchLiquidStakeIntervalTask: %v", err))
	}

	if err := k.HandleRedemptionRateQueryTask(ctx); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error in RedemptionRateQueryTask: %v", err))
	}

	if err := k.HandleRewardsWithdrawalTask(ctx); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error in RewardsWithdrawalTask: %v", err))
	}

	k.Logger(ctx).Info("Cron job completed task processing successfully.")
	return nil
}

// ---- Placeholder Task Implementations ----

// handleLiquidStakingDAppRewards contains the logic for Liquid Staking DApp Rewards task.
func (k PhotosynthesisKeeper) handleLiquidStakingDAppRewards(ctx sdk.Context) {
	k.Logger(ctx).Info("Executing handleLiquidStakingDAppRewards")

	// Retrieve cumulative rewards
	rewardMap := k.GetCumulativeRewardAmount(ctx)

	// Iterate over contracts and process rewards
	state := k.rewardKeeper.GetState()
	contractmeta := state.ContractMetadataState(ctx)

	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		if meta.MinimumRewardAmount > 0 && meta.RewardsAddress != "" {
			amount, exists := rewardMap[meta.RewardsAddress]
			if !exists {
				k.Logger(ctx).Info(fmt.Sprintf("RewardsAddress %s not found in reward map", meta.RewardsAddress))
				return false
			}

			if uint64(amount) >= meta.MinimumRewardAmount {
				record := k.CreateContractLiquidStakeDepositRecord(ctx, amount, sdk.MustAccAddressFromBech32(meta.RewardsAddress))
				k.Logger(ctx).Info(fmt.Sprintf("Created LiquidStakeDepositRecord: %v", record))

				// Wrap the record in a slice
				records := []*types.DepositRecord{record}
				err := k.SetContractLiquidStakeDeposits(ctx, meta.RewardsAddress, records)
				if err != nil {
					k.Logger(ctx).Error(fmt.Sprintf("Error enqueuing LiquidStakeRecord: %v", err))
					return false
				}

				types.EmitLiquidStakeDepositRecordCreatedEvent(ctx, record.String(), record.Amount)
				k.Logger(ctx).Info("Emitted LiquidStakeDepositRecordCreatedEvent")
			}
		}
		return false
	})
}

// handleArchLiquidStakeInterval contains the logic for Arch Liquid Stake Interval task.
func (k PhotosynthesisKeeper) handleArchLiquidStakeInterval(ctx sdk.Context) {
	k.Logger(ctx).Info("Executing handleArchLiquidStakeInterval")

	// Retrieve total liquid stake amount
	totalLiquidStake, err := k.GetTotalLiquidStake(ctx)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error getting total liquid stake: %v", err))
		return
	}

	liquidityAmount := totalLiquidStake.Int64()
	k.Logger(ctx).Info(fmt.Sprintf("Total Liquid Stake: %d", liquidityAmount))

	// Distribute liquidity tokens
	k.DistributeLiquidity(ctx, liquidityAmount)
	k.Logger(ctx).Info(fmt.Sprintf("Distributed Liquidity Amount: %d", liquidityAmount))
}

// handleRedemptionRateQuery contains the logic for Redemption Rate Query task.
func (k PhotosynthesisKeeper) handleRedemptionRateQuery(ctx sdk.Context) {
	k.Logger(ctx).Info("Executing handleRedemptionRateQuery")

	// Query the redemption rate
	redemptionRate, err := k.QueryRedemptionRate(ctx)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error querying redemption rate: %v", err))
		return
	}

	params := k.GetParams(ctx)

	if redemptionRate >= 14 {
		timeSinceRedemption := k.GetTimeSinceLatestRedemption(ctx)
		if timeSinceRedemption >= params.RedemptionIntervalThreshold {
			cumLiquidityAmount, err := k.GetCumulativeLiquidityAmount(ctx)
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("Error getting cumulative liquidity amount: %v", err))
				return
			}

			_, err = k.RedeemLiquidTokens(ctx, &types.Coin{Amount: cumLiquidityAmount.Int64()})
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("Error redeeming liquid tokens: %v", err))
				return
			}

			err = k.DistributeRedeemedTokens(ctx)
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("Error distributing redeemed tokens: %v", err))
				return
			}

			err = k.DeleteRedemptionRecord(ctx)
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("Error deleting redemption record: %v", err))
				return
			}
		}
	}
}

// handleRewardsWithdrawal contains the logic for Rewards Withdrawal task.
func (k PhotosynthesisKeeper) handleRewardsWithdrawal(ctx sdk.Context) {
	k.Logger(ctx).Info("Executing handleRewardsWithdrawal")

	state := k.rewardKeeper.GetState()
	_, records := state.RewardsRecord(ctx).Export()
	k.Logger(ctx).Info(fmt.Sprintf("Processing reward records for contract %v", records))
	totalRewards := sdk.NewCoins()
	var latestTimestamp time.Time
	var blockHeight int64

	rewardMap, rewardsRecords := k.GetTotalRewardRecords(ctx)

	contractmeta := state.ContractMetadataState(ctx)

	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		if meta.RewardsWithdrawalInterval > 0 {
			k.Logger(ctx).Info(fmt.Sprintf("Processing rewards withdrawal for contract %s", meta.RewardsAddress))

			recordCounter := 0
			for _, record := range records {
				recordCounter++
				totalRewards = totalRewards.Add(record.Rewards...)
				if recordCounter == 1 {
					latestTimestamp = record.CalculatedTime
				}
				latestTimestamp = record.CalculatedTime
				blockHeight = record.CalculatedHeight
			}

			if !totalRewards.IsZero() {
				data := fmt.Sprintf("%s,%d,%d,%s,%d",
					meta.RewardsAddress,
					totalRewards.AmountOf("uarch").Int64()-int64(rewardMap[meta.RewardsAddress]),
					int64(recordCounter)-int64(rewardsRecords[meta.RewardsAddress]),
					latestTimestamp.Format(time.RFC3339),
					blockHeight,
				)
				k.Logger(ctx).Info(fmt.Sprintf("Rewards data: %v", data))

				// Store data in the KVStore
				k.AddRewardsDistributionToDapps(ctx, data)
				k.AddRewardsWithdrawalToDapps(ctx, data)
			}
		}
		return false
	})
}

// ---- Additional Functions ----

// CreateContractLiquidStakeDepositRecord creates a new DepositRecord.
func (k PhotosynthesisKeeper) CreateContractLiquidStakeDepositRecord(ctx sdk.Context, amount int64, rewardAddress sdk.AccAddress) *types.DepositRecord {
	k.Logger(ctx).Info(fmt.Sprintf("Creating DepositRecord: RewardAddress %s, Amount %d", rewardAddress, amount))

	depositRecord := &types.DepositRecord{
		ContractAddress: rewardAddress.String(),
		Amount:          amount,
		Status:          "pending",
	}

	return depositRecord
}

// EnqueueLiquidStakeRecord enqueues a DepositRecord for processing.
func (k PhotosynthesisKeeper) EnqueueLiquidStakeRecord(ctx sdk.Context, record *types.DepositRecord) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.LiquidStakeQueuePrefix))
	key := []byte(record.ContractAddress)
	bz := k.cdc.MustMarshal(record)
	store.Set(key, bz)
	k.Logger(ctx).Info(fmt.Sprintf("Enqueued LiquidStakeRecord: %s", record.String()))
	return nil
}

// GetTotalLiquidStake retrieves the total liquid stake amount.
func (k PhotosynthesisKeeper) GetTotalLiquidStake(ctx sdk.Context) (sdk.Int, error) {
	k.Logger(ctx).Info("---------------Start Get Total Liquid Stake----------------")
	totalLiquidStake := sdk.ZeroInt()

	// Iterate through all contracts
	contractmeta := k.rewardKeeper.GetState().ContractMetadataState(ctx)
	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		// Retrieve deposit records for the contract
		depositRecords, err := k.GetContractLiquidStakeDeposits(ctx, meta.RewardsAddress)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error retrieving deposit records for contract %s: %v", meta.RewardsAddress, err))
			return false
		}

		// Sum up the liquid stake for the contract
		contractLiquidStake := sdk.ZeroInt()
		var updatedRecords []*types.DepositRecord

		for _, record := range depositRecords {
			if record.Status == "pending" {
				contractLiquidStake = contractLiquidStake.Add(sdk.NewInt(record.Amount))
				record.Status = "completed"
				updatedRecords = append(updatedRecords, record)
			}
			k.Logger(ctx).Info(fmt.Sprintf("Contract Address:%v, Liquid stake deposit record amount:%v, record status:%v", meta.ContractAddress, record.Amount, record.Status))
		}

		// Update the deposit records in the store if any records were updated
		if len(updatedRecords) > 0 {
			err := k.SetContractLiquidStakeDeposits(ctx, meta.RewardsAddress, updatedRecords)
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("Error updating deposit records for contract %s: %v", meta.RewardsAddress, err))
				return false
			}
		}

		// Add the contract's liquid stake to the total liquid stake
		totalLiquidStake = totalLiquidStake.Add(contractLiquidStake)
		k.Logger(ctx).Info(fmt.Sprintf("Total liquid stake after processing contract %s: %v", meta.RewardsAddress, totalLiquidStake))
		return false
	})
	// Store the updated total liquid stake in the keeper's store
	k.SetTotalLiquidStakeData(ctx, totalLiquidStake)
	k.Logger(ctx).Info("---------------Finish Get Total Liquid Stake----------------")
	return totalLiquidStake, nil
}

// SetTotalLiquidStakeData stores the total liquid stake amount in the keeper's store.
func (k PhotosynthesisKeeper) SetTotalLiquidStakeData(ctx sdk.Context, totalLiquidStake sdk.Int) {
	store := ctx.KVStore(k.storeKey)

	// Log before storing the value
	k.Logger(ctx).Info("Setting total liquid stake in store", "totalLiquidStake", totalLiquidStake.String())

	bz := totalLiquidStake.BigInt().Bytes()
	store.Set([]byte("totalLiquidStake"), bz)

	// Log after storing the value
	k.Logger(ctx).Info("Total liquid stake successfully stored")
}

// GetTotalLiquidStakeData retrieves the total liquid stake amount from the keeper's store.
func (k PhotosynthesisKeeper) GetTotalLiquidStakeData(ctx sdk.Context) sdk.Int {
	store := ctx.KVStore(k.storeKey)

	// Log before retrieving the value
	k.Logger(ctx).Info("Getting total liquid stake from store")

	bz := store.Get([]byte("totalLiquidStake"))
	if bz == nil {
		k.Logger(ctx).Info("Total liquid stake not found in store; returning zero")
		return sdk.ZeroInt()
	}

	totalLiquidStake := sdk.NewIntFromBigInt(new(big.Int).SetBytes(bz))

	// Log after retrieving the value
	k.Logger(ctx).Info("Total liquid stake successfully retrieved", "totalLiquidStake", totalLiquidStake.String())

	return totalLiquidStake
}

// ResetTotalLiquidStake resets the total liquid stake to zero.
func (k PhotosynthesisKeeper) ResetTotalLiquidStake(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)

	// Log before resetting the value
	k.Logger(ctx).Info("Resetting total liquid stake to zero")

	// Set the total liquid stake to zero
	zeroStake := sdk.NewInt(0)
	bz := zeroStake.BigInt().Bytes()
	store.Set([]byte("totalLiquidStake"), bz)

	// Log after resetting the value
	k.Logger(ctx).Info("Total liquid stake successfully reset to zero")
}

// SetContractLiquidStakeDeposits appends new deposit records to the existing records for a contract.
func (k PhotosynthesisKeeper) SetContractLiquidStakeDeposits(ctx sdk.Context, contractAddr string, newRecords []*types.DepositRecord) error {
	k.Logger(ctx).Info(fmt.Sprintf("Setting Contract Liquid Stake Deposits for contract: %s", contractAddr))

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.LiquidStakeRecordsPrefix))
	key := []byte(contractAddr)

	// Fetch existing records
	existingRecords, err := k.GetContractLiquidStakeDeposits(ctx, contractAddr)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error fetching existing records for contract %s: %v", contractAddr, err))
		return err
	}
	k.Logger(ctx).Info(fmt.Sprintf("Existing records for contract %s: %v records", contractAddr, existingRecords))

	// Append new records to existing ones
	allRecords := append(existingRecords, newRecords...)
	k.Logger(ctx).Info(fmt.Sprintf("Total records after appending new records for contract %s: %v records", contractAddr, allRecords))

	// Create the deposit records object
	var depositRecords types.DepositRecords
	depositRecords.Records = allRecords

	// Marshal and set in the store
	bz, err := k.cdc.Marshal(&depositRecords)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error marshaling deposit records for contract %s: %v", contractAddr, err))
		return err
	}
	k.Logger(ctx).Info(fmt.Sprintf("Marshaled deposit records for contract %s", contractAddr))

	store.Set(key, bz)
	k.Logger(ctx).Info(fmt.Sprintf("Successfully set deposit records for contract %s", contractAddr))
	return nil
}

// GetContractLiquidStakeDeposits retrieves the deposit records for a contract.
func (k PhotosynthesisKeeper) GetContractLiquidStakeDeposits(ctx sdk.Context, contractAddr string) ([]*types.DepositRecord, error) {
	k.Logger(ctx).Info(fmt.Sprintf("Getting Contract Liquid Stake Deposits for contract: %s", contractAddr))

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.LiquidStakeRecordsPrefix))
	key := []byte(contractAddr)
	bz := store.Get(key)
	if bz == nil {
		k.Logger(ctx).Info(fmt.Sprintf("No deposit records found for contract %s", contractAddr))
		return nil, nil
	}

	var depositRecords types.DepositRecords
	err := k.cdc.Unmarshal(bz, &depositRecords)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error unmarshaling deposit records for contract %s: %v", contractAddr, err))
		return nil, err
	}
	k.Logger(ctx).Info(fmt.Sprintf("Retrieved %v deposit records for contract %s", depositRecords.Records, contractAddr))

	return depositRecords.Records, nil
}

// GetCumulativeRewardAmount retrieves the cumulative reward amounts from the store.
func (k PhotosynthesisKeeper) GetCumulativeRewardAmount(ctx sdk.Context) map[string]int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RewardsDistributionToDappsKey))
	iter := store.Iterator(nil, nil)
	defer iter.Close()

	rewardMap := make(map[string]int64)
	for ; iter.Valid(); iter.Next() {
		data := string(iter.Value())
		fields := strings.Split(data, ",")
		if len(fields) < 2 {
			continue
		}
		address := fields[0]
		amount, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error parsing amount for address %s: %v", address, err))
			continue
		}
		rewardMap[address] += amount
	}

	// Clear the store after reading
	//store.DeletePrefix(nil)
	k.Logger(ctx).Info(fmt.Sprintf("Cumulative Reward Map: %v", rewardMap))
	return rewardMap
}

/* DistributeLiquidity distributes liquidity tokens to DApps based on their stake.
func (k PhotosynthesisKeeper) DistributeLiquidity(ctx sdk.Context, liquidityAmount int64) {
	k.Logger(ctx).Info(fmt.Sprintf("Starting Liquidity Distribution: Amount %d", liquidityAmount))

	contractmeta := k.rewardKeeper.GetState().ContractMetadataState(ctx)
	cumulativeStakes := make(map[string]sdk.Int)
	totalStake := sdk.NewInt(0)

	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		contractStake, err := k.GetContractStake(ctx, sdk.MustAccAddressFromBech32(meta.RewardsAddress))
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error getting stake for contract %s: %v", meta.RewardsAddress, err))
			return false
		}
		cumulativeStakes[meta.LiquidityProviderAddress] = contractStake
		totalStake = totalStake.Add(contractStake)
		return false
	})

	// Proceed only if there is a total stake to distribute
	if totalStake.IsZero() {
		k.Logger(ctx).Info("Total stake is zero. Skipping distribution.")
		return
	}

	// Open the distribution log store
	distributionStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.LiquidityDistributionLogPrefix))

	// Distribute liquidity tokens proportionally
	for contractAddr, contractStake := range cumulativeStakes {
		stakeProportion := sdk.NewDecFromInt(contractStake).Quo(sdk.NewDecFromInt(totalStake))
		liquidityTokensAmount := stakeProportion.MulInt64(liquidityAmount).TruncateInt()

		// Update the distribution log
		distributionData := fmt.Sprintf("%s,%d", contractAddr, liquidityTokensAmount.Int64())
		distributionStore.Set([]byte(contractAddr), []byte(distributionData))

		k.Logger(ctx).Info(fmt.Sprintf("Distributed %d liquidity tokens to %s", liquidityTokensAmount.Int64(), contractAddr))
	}

	k.Logger(ctx).Info("Finished Liquidity Distribution")
}
*/
// QueryRedemptionRate queries the current redemption rate.
func (k PhotosynthesisKeeper) QueryRedemptionRate(ctx sdk.Context) (int64, error) {
	// Implement logic to query the redemption rate
	// This is a placeholder implementation
	redemptionRate := int64(105) // Replace with actual logic
	return redemptionRate, nil
}

// GetTimeSinceLatestRedemption calculates the time since the latest redemption.
func (k PhotosynthesisKeeper) GetTimeSinceLatestRedemption(ctx sdk.Context) int64 {
	latestRedemptionTime := k.GetLatestRedemptionTime(ctx)
	currentTime := ctx.BlockTime()
	timeDifference := currentTime.Sub(latestRedemptionTime).Seconds()
	return int64(timeDifference)
}

// GetLatestRedemptionTime retrieves the latest redemption time from the store.
func (k PhotosynthesisKeeper) GetLatestRedemptionTime(ctx sdk.Context) time.Time {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.LatestRedemptionTimeKey))
	bz := store.Get([]byte(types.LatestRedemptionTimeStoreKey))
	if bz == nil {
		return time.Time{}
	}
	latestRedemptionTime := int64(binary.BigEndian.Uint64(bz))
	return time.Unix(latestRedemptionTime, 0)
}

// GetCumulativeLiquidityAmount retrieves the cumulative liquidity amount.
func (k PhotosynthesisKeeper) GetCumulativeLiquidityAmount(ctx sdk.Context) (sdk.Int, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CumulativeLiquidityAmountKey))
	bz := store.Get([]byte("cumulative_liquidity"))
	if bz == nil {
		return sdk.ZeroInt(), errors.New("cumulative liquidity amount not set")
	}

	var coins types.Coin
	err := k.cdc.Unmarshal(bz, &coins)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error unmarshaling cumulative liquidity amount: %v", err))
		return sdk.ZeroInt(), err
	}

	k.Logger(ctx).Info(fmt.Sprintf("Cumulative Liquidity Amount: %d", coins.Amount))
	return sdk.NewInt(coins.Amount), nil
}

// RedeemLiquidTokens redeems liquid tokens.
func (k PhotosynthesisKeeper) RedeemLiquidTokens(ctx sdk.Context, amount *types.Coin) (sdk.Int, error) {
	// Implement logic to redeem liquid tokens
	// This is a placeholder implementation
	redeemedAmount := sdk.NewInt(amount.Amount) // Replace with actual logic
	return redeemedAmount, nil
}

// DistributeRedeemedTokens distributes redeemed tokens to all contracts based on their stake.
func (k PhotosynthesisKeeper) DistributeRedeemedTokens(ctx sdk.Context) error {
	k.Logger(ctx).Info("Starting Distribution of Redeemed Tokens")

	contractmeta := k.rewardKeeper.GetState().ContractMetadataState(ctx)
	cumulativeStakes := make(map[string]sdk.Int)
	totalStake := sdk.NewInt(0)

	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		contractStake, err := k.GetContractStake(ctx, sdk.MustAccAddressFromBech32(meta.RewardsAddress))
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error getting stake for contract %s: %v", meta.RewardsAddress, err))
			return false
		}
		cumulativeStakes[meta.LiquidityProviderAddress] = contractStake
		totalStake = totalStake.Add(contractStake)
		return false
	})

	if totalStake.IsZero() {
		k.Logger(ctx).Info("Total stake is zero. Skipping distribution of redeemed tokens.")
		return nil
	}

	distributionStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RedeemedTokensDistributionPrefix))

	for contractAddr, stake := range cumulativeStakes {
		stakeProportion := sdk.NewDecFromInt(stake).Quo(sdk.NewDecFromInt(totalStake))
		redeemedTokensAmount := stakeProportion.MulInt64(k.getRedeemedTokensAmount(ctx)).TruncateInt()

		distributionData := fmt.Sprintf("%s,%d", contractAddr, redeemedTokensAmount.Int64())
		distributionStore.Set([]byte(contractAddr), []byte(distributionData))

		k.Logger(ctx).Info(fmt.Sprintf("Distributed %d redeemed tokens to %s", redeemedTokensAmount.Int64(), contractAddr))
	}

	k.Logger(ctx).Info("Finished Distribution of Redeemed Tokens")
	return nil
}

// getRedeemedTokensAmount retrieves the total redeemed tokens amount.
func (k PhotosynthesisKeeper) getRedeemedTokensAmount(ctx sdk.Context) int64 {
	// Implement logic to get total redeemed tokens amount
	// This is a placeholder implementation
	return 500000 // Replace with actual logic
}

// DeleteRedemptionRecord deletes the latest redemption record from the store.
func (k PhotosynthesisKeeper) DeleteRedemptionRecord(ctx sdk.Context) error {
	//store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RedemptionRecordPrefix))
	//store.DeletePrefix(nil)
	k.Logger(ctx).Info("Deleted RedemptionRecord")
	return nil
}

// GetTotalRewardRecords retrieves total reward records and counts.
func (k PhotosynthesisKeeper) GetTotalRewardRecords(ctx sdk.Context) (map[string]int, map[string]int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RewardsDistributionToDappsKey))
	iter := store.Iterator(nil, nil)
	defer iter.Close()

	rewardMap := make(map[string]int)
	rewardRecords := make(map[string]int)

	for ; iter.Valid(); iter.Next() {
		data := string(iter.Value())
		fields := strings.Split(data, ",")
		if len(fields) < 2 {
			continue
		}
		address := fields[0]
		amount, err := strconv.Atoi(fields[1])
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error parsing amount for address %s: %v", address, err))
			continue
		}
		rewardMap[address] += amount
		// Assuming the third field is the number of records
		if len(fields) >= 3 {
			records, err := strconv.Atoi(fields[2])
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("Error parsing records for address %s: %v", address, err))
				continue
			}
			rewardRecords[address] += records
		}
	}

	// Clear the store after reading
	//store.DeletePrefix(nil)
	k.Logger(ctx).Info(fmt.Sprintf("Total Reward Map: %v", rewardMap))
	k.Logger(ctx).Info(fmt.Sprintf("Total Reward Records: %v", rewardRecords))
	return rewardMap, rewardRecords
}

// AddRewardsDistributionToDapps adds rewards distribution data to the store.
func (k PhotosynthesisKeeper) AddRewardsDistributionToDapps(ctx sdk.Context, data string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RewardsDistributionToDappsKey))
	key := []byte(strconv.FormatInt(ctx.BlockHeight(), 10))
	store.Set(key, []byte(data))
	k.Logger(ctx).Info("Added Rewards Distribution to Dapps to the store")
}

// AddRewardsWithdrawalToDapps adds rewards withdrawal data to the store.
func (k PhotosynthesisKeeper) AddRewardsWithdrawalToDapps(ctx sdk.Context, data string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RewardsWithdrawalToDappsKey))
	key := []byte(strconv.FormatInt(ctx.BlockHeight(), 10))
	store.Set(key, []byte(data))
	k.Logger(ctx).Info("Added Rewards Withdrawal to Dapps to the store")
}

// ResetCompletedDepositRecords resets or removes completed deposit records for a contract.
func (k PhotosynthesisKeeper) ResetCompletedDepositRecords(ctx sdk.Context, contractAddr string) error {
	k.Logger(ctx).Info(fmt.Sprintf("Resetting completed deposit records for contract: %s", contractAddr))

	// Retrieve existing deposit records for the contract
	depositRecords, err := k.GetContractLiquidStakeDeposits(ctx, contractAddr)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error retrieving deposit records for contract %s: %v", contractAddr, err))
		return err
	}

	// Filter out the completed records
	var pendingRecords []*types.DepositRecord
	for _, record := range depositRecords {
		if record.Status != "completed" {
			pendingRecords = append(pendingRecords, record)
		} else {
			k.Logger(ctx).Info(fmt.Sprintf("Removing completed deposit record: %v", record))
		}
	}

	// Update the deposit records in the store with only pending records
	err = k.SetContractLiquidStakeDeposits(ctx, contractAddr, pendingRecords)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error updating deposit records for contract %s: %v", contractAddr, err))
		return err
	}

	k.Logger(ctx).Info(fmt.Sprintf("Successfully reset completed deposit records for contract: %s", contractAddr))
	return nil
}

// ResetAllCompletedDepositRecords resets completed deposit records for all contracts.
func (k PhotosynthesisKeeper) ResetAllCompletedDepositRecords(ctx sdk.Context) error {
	k.Logger(ctx).Info("Resetting completed deposit records for all contracts")

	// Iterate over all contracts
	contractmeta := k.rewardKeeper.GetState().ContractMetadataState(ctx)
	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		contractAddr := meta.RewardsAddress
		// Reset completed deposit records for each contract
		err := k.ResetCompletedDepositRecords(ctx, contractAddr)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error resetting deposit records for contract %s: %v", contractAddr, err))
		}
		return false
	})

	k.Logger(ctx).Info("Successfully reset completed deposit records for all contracts")
	return nil
}

// BeginBlocker is now simplified and does not handle task processing directly.
func (k PhotosynthesisKeeper) BeginBlocker(ctx sdk.Context) abci.ResponseBeginBlock {
	// Retain lightweight tasks here if needed.
	k.Logger(ctx).Info("BeginBlocker invoked - No task processing handled.")
	return abci.ResponseBeginBlock{}
}

// Define a new prefix for storing contract stakes
const (
	ContractStakePrefix = "contract_stake"
)

// AddContractStake adds an amount to the contract's stake during the cycle.
func (k PhotosynthesisKeeper) AddContractStake(ctx sdk.Context, contractAddr string, amount sdk.Int) {
	// Retrieve current stake
	currentStake := k.GetContractStakeData(ctx, contractAddr)
	// Add the amount
	newStake := currentStake.Add(amount)
	// Store the new stake
	k.SetContractStakeData(ctx, contractAddr, newStake)
}

// GetContractStakeData retrieves the stake of a contract from the store.
func (k PhotosynthesisKeeper) GetContractStakeData(ctx sdk.Context, contractAddr string) sdk.Int {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(ContractStakePrefix))
	bz := store.Get([]byte(contractAddr))
	if bz == nil {
		return sdk.ZeroInt()
	}
	return sdk.NewIntFromBigInt(new(big.Int).SetBytes(bz))
}

// SetContractStakeData stores the stake of a contract in the store.
func (k PhotosynthesisKeeper) SetContractStakeData(ctx sdk.Context, contractAddr string, stake sdk.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(ContractStakePrefix))
	store.Set([]byte(contractAddr), stake.BigInt().Bytes())
}

// ResetContractStakes resets the stakes for all contracts.
func (k PhotosynthesisKeeper) ResetContractStakes(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(ContractStakePrefix))
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		store.Delete(iterator.Key())
	}
	k.Logger(ctx).Info("Reset all contract stakes to zero")
}

// Modify GetContractStake to use the stored stake data
func (k PhotosynthesisKeeper) GetContractStake(ctx sdk.Context, address sdk.AccAddress) (sdk.Int, error) {
	contractStake := k.GetContractStakeData(ctx, address.String())
	return contractStake, nil
}

// Add a new prefix for storing stake ratios
const (
	StakeRatioPrefix = "stake_ratio"
)

// Modify DistributeLiquidity to store stake ratios
func (k PhotosynthesisKeeper) DistributeLiquidity(ctx sdk.Context, liquidityAmount int64) {
	k.Logger(ctx).Info(fmt.Sprintf("Starting Liquidity Distribution: Amount %d", liquidityAmount))

	// Retrieve total stakes per contract
	contractmeta := k.rewardKeeper.GetState().ContractMetadataState(ctx)
	cumulativeStakes := make(map[string]sdk.Int)
	totalStake := sdk.NewInt(0)

	contractmeta.IterateContractMetadata(func(meta *rewardstypes.ContractMetadata) (stop bool) {
		contractStake, err := k.GetContractStake(ctx, sdk.MustAccAddressFromBech32(meta.RewardsAddress))
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error getting stake for contract %s: %v", meta.RewardsAddress, err))
			return false
		}
		if contractStake.IsPositive() {
			cumulativeStakes[meta.RewardsAddress] = contractStake
			totalStake = totalStake.Add(contractStake)
		}
		return false
	})

	// Proceed only if there is a total stake to distribute
	if totalStake.IsZero() {
		k.Logger(ctx).Info("Total stake is zero. Skipping distribution.")
		return
	}

	// Open the distribution log store and stake ratio store
	distributionStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.LiquidityDistributionLogPrefix))
	stakeRatioStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(StakeRatioPrefix))

	// Distribute liquidity tokens proportionally
	for contractAddr, contractStake := range cumulativeStakes {
		stakeProportion := sdk.NewDecFromInt(contractStake).Quo(sdk.NewDecFromInt(totalStake))
		liquidityTokensAmount := stakeProportion.MulInt64(liquidityAmount).TruncateInt()

		// Update the distribution log
		distributionData := fmt.Sprintf("%s,%d", contractAddr, liquidityTokensAmount.Int64())
		distributionStore.Set([]byte(contractAddr), []byte(distributionData))

		// Store the stake ratio
		stakeRatioStore.Set([]byte(contractAddr), []byte(stakeProportion.String()))

		k.Logger(ctx).Info(fmt.Sprintf("Distributed %d liquidity tokens to %s with stake ratio %s", liquidityTokensAmount.Int64(), contractAddr, stakeProportion.String()))
	}

	k.Logger(ctx).Info("Finished Liquidity Distribution")
}

// SetStakeRatio sets the stake ratio for a contract in the store.
func (k PhotosynthesisKeeper) SetStakeRatio(ctx sdk.Context, contractAddr sdk.AccAddress, stakeRatio sdk.Dec) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(StakeRatioPrefix))
	key := contractAddr.Bytes()

	// Serialize stake ratio
	bz, err := stakeRatio.MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal stake ratio for contract %s: %w", contractAddr.String(), err)
	}

	store.Set(key, bz)
	return nil
}

// GetStakeRatio retrieves the stake ratio for a contract from the store.
func (k PhotosynthesisKeeper) GetStakeRatio(ctx sdk.Context, contractAddr sdk.AccAddress) (sdk.Dec, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(StakeRatioPrefix))
	key := contractAddr.Bytes()
	bz := store.Get(key)

	if bz == nil {
		return sdk.Dec{}, fmt.Errorf("stake ratio not found for contract: %s", contractAddr.String())
	}

	var stakeRatio sdk.Dec
	if err := stakeRatio.UnmarshalJSON(bz); err != nil {
		return sdk.Dec{}, fmt.Errorf("failed to unmarshal stake ratio for contract %s: %w", contractAddr.String(), err)
	}

	return stakeRatio, nil
}

// GetAllStakeRatios retrieves the stake ratios for all staking contracts.
func (k PhotosynthesisKeeper) GetAllStakeRatios(ctx sdk.Context) ([]*types.StakeRatio, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(StakeRatioPrefix))
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var stakeRatios []*types.StakeRatio
	for ; iterator.Valid(); iterator.Next() {
		contractAddrBytes := iterator.Key()
		bz := iterator.Value()

		var stakeRatio sdk.Dec
		if err := stakeRatio.UnmarshalJSON(bz); err != nil {
			return nil, fmt.Errorf("failed to unmarshal stake ratio: %w", err)
		}

		contractAddr := sdk.AccAddress(contractAddrBytes)

		stakeRatios = append(stakeRatios, &types.StakeRatio{
			ContractAddress: contractAddr.String(),
			StakeRatio:      stakeRatio.String(),
		})
	}

	return stakeRatios, nil
}

// SetRedeemTokenData stores the redeem token amount for a contract in the store.
func (k PhotosynthesisKeeper) SetRedeemTokenData(ctx sdk.Context, contractAddr sdk.AccAddress, amount sdk.Int) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(RedeemTokenPrefix))
	key := contractAddr.Bytes()

	bz := amount.BigInt().Bytes()
	store.Set(key, bz)

	k.Logger(ctx).Info(fmt.Sprintf("Set redeem token amount for contract: %s, amount: %s", contractAddr.String(), amount.String()))
	return nil
}

// GetRedeemTokenData retrieves the redeem token amount for a contract from the store.
func (k PhotosynthesisKeeper) GetRedeemTokenData(ctx sdk.Context, contractAddr sdk.AccAddress) (sdk.Int, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(RedeemTokenPrefix))
	key := contractAddr.Bytes()
	bz := store.Get(key)

	if bz == nil {
		return sdk.ZeroInt(), fmt.Errorf("redeem token amount not found for contract: %s", contractAddr.String())
	}

	amount := sdk.NewIntFromBigInt(new(big.Int).SetBytes(bz))
	return amount, nil
}

// GetTotalRedemptionAmount calculates the total redemption amount from all contracts.
func (k PhotosynthesisKeeper) GetTotalRedemptionAmount(ctx sdk.Context) (sdk.Int, error) {
	k.Logger(ctx).Info("Calculating total redemption amount")

	totalRedeemAmount := sdk.ZeroInt()

	// Iterate over all contracts with redeem token data
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(RedeemTokenPrefix))
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		amount := sdk.NewIntFromBigInt(new(big.Int).SetBytes(bz))
		totalRedeemAmount = totalRedeemAmount.Add(amount)
	}

	k.Logger(ctx).Info(fmt.Sprintf("Total redemption amount: %s", totalRedeemAmount.String()))
	return totalRedeemAmount, nil
}

// ComputeRedeemTokenRatios computes and stores the redeem ratios for all contracts.
func (k PhotosynthesisKeeper) ComputeRedeemTokenRatios(ctx sdk.Context) error {
	k.Logger(ctx).Info("Starting computation of redeem token ratios")

	// Initialize variables
	totalRedeemAmount := sdk.ZeroInt()
	redeemAmounts := make(map[string]sdk.Int)

	// Iterate over all contracts with redeem token data
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(RedeemTokenPrefix))
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		contractAddrBytes := iterator.Key()
		bz := iterator.Value()

		amount := sdk.NewIntFromBigInt(new(big.Int).SetBytes(bz))
		contractAddr := sdk.AccAddress(contractAddrBytes)
		redeemAmounts[contractAddr.String()] = amount
		totalRedeemAmount = totalRedeemAmount.Add(amount)
	}

	if totalRedeemAmount.IsZero() {
		k.Logger(ctx).Info("Total redeem amount is zero. Skipping computing redeem ratios.")
		return nil
	}

	// Compute and store ratios
	ratioStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(RedeemTokenRatioPrefix))

	for contractAddrStr, amount := range redeemAmounts {
		redeemRatio := sdk.NewDecFromInt(amount).Quo(sdk.NewDecFromInt(totalRedeemAmount))
		contractAddr, err := sdk.AccAddressFromBech32(contractAddrStr)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Invalid contract address: %s", contractAddrStr))
			continue
		}

		// Store the redeem ratio
		key := contractAddr.Bytes()
		bz, err := redeemRatio.Marshal()
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error marshaling redeem ratio for contract %s: %v", contractAddrStr, err))
			continue
		}
		ratioStore.Set(key, bz)
		k.Logger(ctx).Info(fmt.Sprintf("Stored redeem ratio for contract %s: %s", contractAddrStr, redeemRatio.String()))
	}

	k.Logger(ctx).Info("Computed and stored redeem token ratios.")
	return nil
}

// SetRedeemTokenRatio sets the redeem ratio for a contract in the store.
func (k PhotosynthesisKeeper) SetRedeemTokenRatio(ctx sdk.Context, contractAddr sdk.AccAddress, redeemRatio sdk.Dec) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(RedeemTokenRatioPrefix))
	key := contractAddr.Bytes()

	// Serialize redeem ratio
	bz, err := redeemRatio.Marshal()
	if err != nil {
		return fmt.Errorf("failed to marshal redeem ratio for contract %s: %w", contractAddr.String(), err)
	}

	store.Set(key, bz)
	k.Logger(ctx).Info(fmt.Sprintf("Set redeem ratio for contract %s: %s", contractAddr.String(), redeemRatio.String()))
	return nil
}

// GetRedeemTokenRatio retrieves the redeem ratio for a contract from the store.
func (k PhotosynthesisKeeper) GetRedeemTokenRatio(ctx sdk.Context, contractAddr sdk.AccAddress) (sdk.Dec, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(RedeemTokenRatioPrefix))
	key := contractAddr.Bytes()
	bz := store.Get(key)

	if bz == nil {
		return sdk.Dec{}, fmt.Errorf("redeem ratio not found for contract: %s", contractAddr.String())
	}

	var redeemRatio sdk.Dec
	if err := redeemRatio.Unmarshal(bz); err != nil {
		return sdk.Dec{}, fmt.Errorf("failed to unmarshal redeem ratio for contract %s: %w", contractAddr.String(), err)
	}

	return redeemRatio, nil
}

// GetAllRedeemTokenRatios retrieves the redeem ratios for all contracts.
func (k PhotosynthesisKeeper) GetAllRedeemTokenRatios(ctx sdk.Context) ([]*types.RedeemTokenRatio, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(RedeemTokenRatioPrefix))
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var redeemRatios []*types.RedeemTokenRatio
	for ; iterator.Valid(); iterator.Next() {
		contractAddrBytes := iterator.Key()
		bz := iterator.Value()

		var redeemRatio sdk.Dec
		if err := redeemRatio.Unmarshal(bz); err != nil {
			return nil, fmt.Errorf("failed to unmarshal redeem ratio: %w", err)
		}

		contractAddr := sdk.AccAddress(contractAddrBytes)

		redeemRatios = append(redeemRatios, &types.RedeemTokenRatio{
			ContractAddress: contractAddr.String(),
			RedeemRatio:     redeemRatio.String(),
		})
	}

	return redeemRatios, nil
}

// ClearRedeemTokenData clears the redeem token data and ratios from the store.
func (k PhotosynthesisKeeper) ClearRedeemTokenData(ctx sdk.Context) {
	// Clear redeem token amounts
	amountStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(RedeemTokenPrefix))
	amountIterator := amountStore.Iterator(nil, nil)
	defer amountIterator.Close()

	for ; amountIterator.Valid(); amountIterator.Next() {
		amountStore.Delete(amountIterator.Key())
	}

	// Clear redeem token ratios
	ratioStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(RedeemTokenRatioPrefix))
	ratioIterator := ratioStore.Iterator(nil, nil)
	defer ratioIterator.Close()

	for ; ratioIterator.Valid(); ratioIterator.Next() {
		ratioStore.Delete(ratioIterator.Key())
	}

	k.Logger(ctx).Info("Cleared redeem token data and ratios from the store")
}
