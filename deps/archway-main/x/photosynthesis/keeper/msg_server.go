package keeper

import (
	"context"
	"fmt"

	"github.com/archway-network/archway/x/photosynthesis/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgServer implements the module gRPC messaging service.
type MsgServer struct {
	keeper PhotosynthesisKeeper
}

// NewMsgServer creates a new gRPC messaging server.
func NewMsgServer(keeper PhotosynthesisKeeper) *MsgServer {
	return &MsgServer{
		keeper: keeper,
	}
}

// handles MsgExecuteCronTasks messages.
func (k MsgServer) ExecuteCronTasks(goCtx context.Context, msg *types.MsgExecuteCronTasks) (*types.MsgExecuteCronTasksResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.keeper.CronJobTriggerTasks(ctx)
	if err != nil {
		return nil, err
	}

	return &types.MsgExecuteCronTasksResponse{}, nil
}

// ResetTotalLiquidStake handles MsgResetTotalLiquidStake messages.
func (k MsgServer) ResetTotalLiquidStake(goCtx context.Context, msg *types.MsgResetTotalLiquidStake) (*types.MsgResetTotalLiquidStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Call the keeper to reset the total liquid stake
	k.keeper.ResetTotalLiquidStake(ctx)

	return &types.MsgResetTotalLiquidStakeResponse{}, nil
}

// ResetAllCompletedDepositRecords handles MsgResetTotalLiquidStake messages.
func (k MsgServer) ResetAllCompletedDepositRecords(goCtx context.Context, msg *types.MsgResetAllCompletedDepositRecords) (*types.MsgResetAllCompletedDepositRecordsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Call the keeper to reset the total liquid stake
	k.keeper.ResetAllCompletedDepositRecords(ctx)

	return &types.MsgResetAllCompletedDepositRecordsResponse{}, nil
}

// SetRedeemTokenData handles MsgSetRedeemTokenData messages.
func (m MsgServer) SetRedeemTokenData(goCtx context.Context, msg *types.MsgSetRedeemTokenData) (*types.MsgSetRedeemTokenDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Parse contract address
	contractAddr, err := sdk.AccAddressFromBech32(msg.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid contract address %s: %w", msg.ContractAddress, err)
	}

	// Parse amount
	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return nil, fmt.Errorf("invalid amount %s:", msg.Amount)
	}

	// Validate amount
	if !amount.IsPositive() {
		return nil, fmt.Errorf("amount must be positive")
	}

	// Set redeem token data via keeper
	if err := m.keeper.SetRedeemTokenData(ctx, contractAddr, amount); err != nil {
		return nil, fmt.Errorf("failed to set redeem token data: %w", err)
	}

	return &types.MsgSetRedeemTokenDataResponse{}, nil
}

// ClearRedeemTokenData handles MsgClearRedeemTokenData messages.
func (m MsgServer) ClearRedeemTokenData(goCtx context.Context, msg *types.MsgClearRedeemTokenData) (*types.MsgClearRedeemTokenDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Clear redeem token data via keeper
	m.keeper.ClearRedeemTokenData(ctx)

	return &types.MsgClearRedeemTokenDataResponse{}, nil
}
