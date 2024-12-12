package keeper

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Define constants for prefixes and keys used in the keeper
const (
	LiquidStakeQueuePrefix           = "liquidStakeQueue"
	LiquidStakeRecordsPrefix         = "liquidStakeRecords"
	RewardsDistributionToDappsKey    = "rewardsDistributionToDapps"
	RewardsWithdrawalToDappsKey      = "rewardsWithdrawalToDapps"
	CumulativeLiquidityAmountKey     = "cumulativeLiquidity"
	LiquidityDistributionLogPrefix   = "liquidityDistributionLog"
	RedeemedTokensDistributionPrefix = "redeemedTokensDistribution"
	LatestRedemptionTimeKey          = "latestRedemptionTime"
	LatestRedemptionTimeStoreKey     = "latestRedemptionTimeKey"
)

// Default parameter values
const (
	DefaultLiquidStakingInterval       int64 = 2 // 30 minutes
	DefaultArchLiquidStakeInterval     int64 = 3 // 120 minutes
	DefaultRedemptionRateQueryInterval int64 = 1 // 15 minutes
	DefaultRedemptionRateThreshold     int64 = 105
	DefaultRedemptionIntervalThreshold int64 = 30 // 60 minutes
	DefaultRewardsWithdrawalInterval   int64 = 1  // 15 minutes
)

// Parameter keys
var (
	KeyLiquidStakingInterval       = []byte("LiquidStakingInterval")
	KeyArchLiquidStakeInterval     = []byte("ArchLiquidStakeInterval")
	KeyRedemptionRateQueryInterval = []byte("RedemptionRateQueryInterval")
	KeyRedemptionRateThreshold     = []byte("RedemptionRateThreshold")
	KeyRedemptionIntervalThreshold = []byte("RedemptionIntervalThreshold")
	KeyRewardsWithdrawalInterval   = []byte("RewardsWithdrawalInterval")
)

// Params defines the parameters for the photosynthesis module.
type Params struct {
	LiquidStakingInterval       int64 `json:"liquid_staking_interval" yaml:"liquid_staking_interval"`
	ArchLiquidStakeInterval     int64 `json:"arch_liquid_stake_interval" yaml:"arch_liquid_stake_interval"`
	RedemptionRateQueryInterval int64 `json:"redemption_rate_query_interval" yaml:"redemption_rate_query_interval"`
	RedemptionRateThreshold     int64 `json:"redemption_rate_threshold" yaml:"redemption_rate_threshold"`
	RedemptionIntervalThreshold int64 `json:"redemption_interval_threshold" yaml:"redemption_interval_threshold"`
	RewardsWithdrawalInterval   int64 `json:"rewards_withdrawal_interval" yaml:"rewards_withdrawal_interval"`
}

// ParamKeyTable returns the KeyTable for module parameters.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params object.
func NewParams(
	liquidStakingInterval int64,
	archLiquidStakeInterval int64,
	redemptionRateQueryInterval int64,
	redemptionRateThreshold int64,
	redemptionIntervalThreshold int64,
	rewardsWithdrawalInterval int64,
) Params {
	return Params{
		LiquidStakingInterval:       liquidStakingInterval,
		ArchLiquidStakeInterval:     archLiquidStakeInterval,
		RedemptionRateQueryInterval: redemptionRateQueryInterval,
		RedemptionRateThreshold:     redemptionRateThreshold,
		RedemptionIntervalThreshold: redemptionIntervalThreshold,
		RewardsWithdrawalInterval:   rewardsWithdrawalInterval,
	}
}

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return NewParams(
		DefaultLiquidStakingInterval,
		DefaultArchLiquidStakeInterval,
		DefaultRedemptionRateQueryInterval,
		DefaultRedemptionRateThreshold,
		DefaultRedemptionIntervalThreshold,
		DefaultRewardsWithdrawalInterval,
	)
}

// Validate validates the module parameters.
func (p Params) Validate() error {
	if err := validatePositiveInt64(p.LiquidStakingInterval); err != nil {
		return fmt.Errorf("invalid LiquidStakingInterval: %w", err)
	}
	if err := validatePositiveInt64(p.ArchLiquidStakeInterval); err != nil {
		return fmt.Errorf("invalid ArchLiquidStakeInterval: %w", err)
	}
	if err := validatePositiveInt64(p.RedemptionRateQueryInterval); err != nil {
		return fmt.Errorf("invalid RedemptionRateQueryInterval: %w", err)
	}
	if err := validatePositiveInt64(p.RedemptionRateThreshold); err != nil {
		return fmt.Errorf("invalid RedemptionRateThreshold: %w", err)
	}
	if err := validatePositiveInt64(p.RedemptionIntervalThreshold); err != nil {
		return fmt.Errorf("invalid RedemptionIntervalThreshold: %w", err)
	}
	if err := validatePositiveInt64(p.RewardsWithdrawalInterval); err != nil {
		return fmt.Errorf("invalid RewardsWithdrawalInterval: %w", err)
	}
	return nil
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs of module parameters.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyLiquidStakingInterval, &p.LiquidStakingInterval, validatePositiveInt64),
		paramtypes.NewParamSetPair(KeyArchLiquidStakeInterval, &p.ArchLiquidStakeInterval, validatePositiveInt64),
		paramtypes.NewParamSetPair(KeyRedemptionRateQueryInterval, &p.RedemptionRateQueryInterval, validatePositiveInt64),
		paramtypes.NewParamSetPair(KeyRedemptionRateThreshold, &p.RedemptionRateThreshold, validatePositiveInt64),
		paramtypes.NewParamSetPair(KeyRedemptionIntervalThreshold, &p.RedemptionIntervalThreshold, validatePositiveInt64),
		paramtypes.NewParamSetPair(KeyRewardsWithdrawalInterval, &p.RewardsWithdrawalInterval, validatePositiveInt64),
	}
}

// Validation functions
func validatePositiveInt64(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return fmt.Errorf("parameter must be positive: %d", v)
	}
	return nil
}
