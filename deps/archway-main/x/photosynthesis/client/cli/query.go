package cli

import (
	"context"

	"github.com/archway-network/archway/x/photosynthesis/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

// GetQueryCmd builds the query command group for the Photosynthesis module.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the Photosynthesis module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// Add individual query commands
	cmd.AddCommand(
		CmdQueryParams(),
		CmdQueryTotalLiquidStake(),
		CmdQueryStakeRatio(),
		CmdQueryAllStakeRatios(),
		CmdQueryAllRedeemTokenRatios(),
		CmdQueryTotalRedemptionAmount(),
	)

	// Add query flags to the command group
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryParams implements the command to query module parameters
func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query the current photosynthesis parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryTotalLiquidStake queries the total liquid stake amount
func CmdQueryTotalLiquidStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-liquid-stake",
		Short: "Query the total liquid stake amount",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.TotalLiquidStake(context.Background(), &types.QueryTotalLiquidStakeRequest{})
			if err != nil {
				return err
			}

			// Print the result in JSON format
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryStakeRatio implements the command to query the stake ratio of a contract.
func CmdQueryStakeRatio() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake-ratio [contract-address]",
		Short: "Query the stake ratio of a contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			contractAddress := args[0]

			res, err := queryClient.StakeRatio(context.Background(), &types.QueryStakeRatioRequest{
				ContractAddress: contractAddress,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryAllStakeRatios implements the command to query stake ratios for all contracts.
func CmdQueryAllStakeRatios() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-stake-ratios",
		Short: "Query the stake ratios of all contracts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AllStakeRatios(context.Background(), &types.QueryAllStakeRatiosRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryAllRedeemTokenRatios queries redeem token ratios for all contracts.
func CmdQueryAllRedeemTokenRatios() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-redeem-token-ratios",
		Short: "Query redeem token ratios for all contracts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.RedeemTokenRatios(context.Background(), &types.QueryRedeemTokenRatiosRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryTotalRedemptionAmount queries the total redemption amount.
func CmdQueryTotalRedemptionAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-redemption-amount",
		Short: "Query the total redemption amount across all contracts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.TotalRedemptionAmount(context.Background(), &types.QueryTotalRedemptionAmountRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
