package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stretchr/testify/suite"

	e2eTesting "github.com/archway-network/archway/e2e/testing"
	"github.com/archway-network/archway/x/photosynthesis/types"
)

type KeeperTestSuite struct {
	suite.Suite

	chain *e2eTesting.TestChain
}

func (s *KeeperTestSuite) SetupTest() {
	s.chain = e2eTesting.NewTestChain(s.T(), 1)
}

/*
	func (s *KeeperTestSuite) TestGRPC_ListContracts() {
		ctx, k := s.chain.GetContext(), s.chain.GetApp().PhotosynthesisKeeper
		querySrvr := NewQueryServer(k)

		s.Run("ok: empty list", func() {
			req := &types.QueryListContractsRequest{}
			res, err := querySrvr.ListContracts(sdk.WrapSDKContext(ctx), req)
			s.Require().NoError(err)
			s.Require().Equal(0, len(res.ContractAddresses))
		})
	}
*/
func (s *KeeperTestSuite) TestGRPC_ArchLiquidStakeInterval() {
	ctx, k := s.chain.GetContext(), s.chain.GetApp().PhotosynthesisKeeper
	querySrvr := NewQueryServer(k)

	s.Run("ok: gets interval", func() {
		req := &types.QueryArchLiquidStakeIntervalRequest{}
		res, err := querySrvr.QueryArchLiquidStakeInterval(ctx, req)
		s.Require().NoError(err)
		s.Require().Equal(uint64(0), res.ArchLiquidStakeInterval)
	})
}

func (s *KeeperTestSuite) TestGRPC_CumulativeLiquidityAmount() {
	ctx, k := s.chain.GetContext(), s.chain.GetApp().PhotosynthesisKeeper
	querySrvr := NewQueryServer(k)

	s.Run("ok: gets amount", func() {
		//req := &types.QueryCumulativeLiquidityAmountRequest{}
		addr := sdk.AccAddress("archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m")
		epoch := int64(0)
		res, err := querySrvr.QueryCumulativeLiquidityAmount(ctx, epoch, addr)
		s.Require().NoError(err)
		s.Require().Equal(uint64(0), res.CumulativeLiquidityAmount)
	})
}

func (s *KeeperTestSuite) TestGRPC_LatestRedemptionRecord() {
	ctx, k := s.chain.GetContext(), s.chain.GetApp().PhotosynthesisKeeper
	querySrvr := NewQueryServer(k)

	s.Run("err: no redemption record found", func() {
		req := &types.QueryLatestRedemptionRecordRequest{}
		_, err := querySrvr.QueryLatestRedemptionRecord(ctx, req)
		s.Require().Error(err)
		s.Require().Equal(status.Error(codes.NotFound, "no redemption record found"), err)
	})
}

func (s *KeeperTestSuite) TestGRPC_RedemptionInterval() {
	ctx, k := s.chain.GetContext(), s.chain.GetApp().PhotosynthesisKeeper
	querySrvr := NewQueryServer(k)

	s.Run("ok: gets interval", func() {
		req := &types.QueryRedemptionIntervalRequest{}
		res, err := querySrvr.QueryRedemptionInterval(ctx, req)
		s.Require().NoError(err)
		s.Require().Equal(uint64(0), res.RedemptionInterval)
	})
}
