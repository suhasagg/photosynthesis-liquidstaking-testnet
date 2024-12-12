package keeper

import (
	"context"
	"fmt"

	"github.com/archway-network/archway/x/photosynthesis/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper implements the QueryServer interface for the Photosynthesis module.
type Keeper struct {
	PhotosynthesisKeeper
}

// NewQueryServerImpl returns an implementation of the QueryServer interface for the provided Keeper.
type QueryServer struct {
	keeper PhotosynthesisKeeper
}

// NewQueryServer returns an implementation of the QueryServer interface.
func NewQueryServer(keeper PhotosynthesisKeeper) *QueryServer {
	return &QueryServer{keeper: keeper}
}

// Params returns the module parameters.
func (s *QueryServer) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := s.keeper.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}

// TotalLiquidStake returns the total liquid stake amount.
func (s *QueryServer) TotalLiquidStake(c context.Context, req *types.QueryTotalLiquidStakeRequest) (*types.QueryTotalLiquidStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	totalStake := s.keeper.GetTotalLiquidStakeData(ctx)

	// Convert sdk.Int to string to avoid issues with protobuf encoding of big integers
	totalStakeStr := totalStake.String()

	return &types.QueryTotalLiquidStakeResponse{TotalLiquidStake: totalStakeStr}, nil
}

// DepositRecords returns deposit records for a contract.
func (s *QueryServer) DepositRecords(c context.Context, req *types.QueryDepositRecordsRequest) (*types.QueryDepositRecordsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	contractAddr, err := sdk.AccAddressFromBech32(req.ContractAddress)
	if err != nil {
		return nil, err
	}

	deposits, err := s.keeper.GetContractLiquidStakeDeposits(ctx, contractAddr.String())
	if err != nil {
		return nil, err
	}

	// Include pagination handling if required (not shown here)
	return &types.QueryDepositRecordsResponse{DepositRecords: deposits}, nil
}

// RedemptionRateThreshold returns the redemption rate threshold.
func (s *QueryServer) RedemptionRateThreshold(c context.Context, req *types.QueryRedemptionRateThresholdRequest) (*types.QueryRedemptionRateThresholdResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	params := s.keeper.GetParams(ctx)
	return &types.QueryRedemptionRateThresholdResponse{RedemptionRateThreshold: params.RedemptionRateThreshold}, nil
}

// RedemptionIntervalThreshold returns the redemption interval threshold.
func (s *QueryServer) RedemptionIntervalThreshold(c context.Context, req *types.QueryRedemptionIntervalThresholdRequest) (*types.QueryRedemptionIntervalThresholdResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	params := s.keeper.GetParams(ctx)
	return &types.QueryRedemptionIntervalThresholdResponse{RedemptionIntervalThreshold: params.RedemptionIntervalThreshold}, nil
}

// AllStakeRatios handles the AllStakeRatios gRPC query.
func (q *QueryServer) AllStakeRatios(ctx context.Context, req *types.QueryAllStakeRatiosRequest) (*types.QueryAllStakeRatiosResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	stakeRatios, err := q.keeper.GetAllStakeRatios(sdkCtx)
	if err != nil {
		return nil, fmt.Errorf("failed to get stake ratios: %w", err)
	}

	return &types.QueryAllStakeRatiosResponse{
		StakeRatios: stakeRatios,
	}, nil
}

// StakeRatio handles the StakeRatio gRPC query for a specific contract address.
func (q *QueryServer) StakeRatio(ctx context.Context, req *types.QueryStakeRatioRequest) (*types.QueryStakeRatioResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check if the contract address is provided in the request
	if req.ContractAddress == "" {
		return nil, fmt.Errorf("contract address is required")
	}

	// Parse the contract address from the request
	contractAddr, err := sdk.AccAddressFromBech32(req.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid contract address: %w", err)
	}

	// Get the stake ratio for the specific contract address
	stakeRatio, err := q.keeper.GetStakeRatio(sdkCtx, contractAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get stake ratio for contract %s: %w", req.ContractAddress, err)
	}

	return &types.QueryStakeRatioResponse{
		ContractAddress: req.ContractAddress,
		StakeRatio:      stakeRatio.String(),
	}, nil
}

func (q *QueryServer) RedeemTokenRatios(c context.Context, req *types.QueryRedeemTokenRatiosRequest) (*types.QueryRedeemTokenRatiosResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	redeemRatios, err := q.keeper.GetAllRedeemTokenRatios(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryRedeemTokenRatiosResponse{RedeemRatios: redeemRatios}, nil
}

// TotalRedemptionAmount queries the total redemption amount across all contracts.
func (q *QueryServer) TotalRedemptionAmount(ctx context.Context, req *types.QueryTotalRedemptionAmountRequest) (*types.QueryTotalRedemptionAmountResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	totalAmount, err := q.keeper.GetTotalRedemptionAmount(sdkCtx)
	if err != nil {
		return nil, fmt.Errorf("failed to get total redemption amount: %w", err)
	}

	return &types.QueryTotalRedemptionAmountResponse{
		TotalAmount: totalAmount.String(),
	}, nil
}
