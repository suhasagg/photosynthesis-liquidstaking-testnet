package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TaskTime represents the last processed time of a task.
type TaskTime struct {
	Key  string    `json:"key" yaml:"key"`
	Time time.Time `json:"time" yaml:"time"`
}

// NewGenesisState creates a new GenesisState object.
func NewGenesisState(
	params Params,
	depositRecords []DepositRecord, // Updated to value type
	contracts []Contract, // Updated to value type
	cumulativeLiquidityAmount sdk.Coin,
) *GenesisState {
	return &GenesisState{
		Params:                    params,
		DepositRecords:            depositRecords, // No longer need dereferencing
		Contracts:                 contracts,      // No longer need dereferencing
		CumulativeLiquidityAmount: cumulativeLiquidityAmount,
	}
}

// DefaultGenesisState returns the default genesis state.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:                    DefaultParams(),
		DepositRecords:            []DepositRecord{}, // Updated to value type
		Contracts:                 []Contract{},      // Updated to value type
		CumulativeLiquidityAmount: sdk.NewCoin("uarch", sdk.NewInt(0)),
	}
}

// Validate performs basic genesis state validation.
func (s GenesisState) Validate() error {
	// Validate Params
	if err := s.Params.Validate(); err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}

	// Validate deposit records
	for _, record := range s.DepositRecords {
		if err := record.Validate(); err != nil {
			return fmt.Errorf("invalid deposit record: %w", err)
		}
	}

	// Validate contracts
	for _, contract := range s.Contracts {
		if err := contract.Validate(); err != nil {
			return fmt.Errorf("invalid contract: %w", err)
		}
	}

	return nil
}

// Validate function for DepositRecord
func (dr DepositRecord) Validate() error { // Updated to value type
	if _, err := sdk.AccAddressFromBech32(dr.ContractAddress); err != nil {
		return fmt.Errorf("invalid contract address: %w", err)
	}
	if dr.Amount <= 0 {
		return fmt.Errorf("invalid amount: %d", dr.Amount)
	}
	if dr.Status == "" {
		return fmt.Errorf("status cannot be empty")
	}
	return nil
}

// Validate function for Contract
func (c Contract) Validate() error { // Updated to value type
	if _, err := sdk.AccAddressFromBech32(c.Address); err != nil {
		return fmt.Errorf("invalid contract address: %w", err)
	}
	if _, err := sdk.AccAddressFromBech32(c.LiquidityProviderAddress); err != nil {
		return fmt.Errorf("invalid liquidity provider address: %w", err)
	}
	if _, err := sdk.AccAddressFromBech32(c.RewardsAddress); err != nil {
		return fmt.Errorf("invalid rewards address: %w", err)
	}
	if c.MinimumRewardAmount == 0 {
		return fmt.Errorf("minimum reward amount cannot be zero")
	}
	return nil
}
