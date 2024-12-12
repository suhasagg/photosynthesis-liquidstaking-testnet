package cli

import (
	"fmt"

	"github.com/archway-network/archway/x/photosynthesis/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

// GetTxCmd builds the transaction command group for the Photosynthesis module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Transaction commands for the Photosynthesis module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// Add subcommands
	cmd.AddCommand(CmdExecuteCronTasks())
	cmd.AddCommand(CmdResetTotalLiquidStake())
	cmd.AddCommand(CmdResetAllCompletedDepositRecords())
	cmd.AddCommand(CmdSetRedeemTokenData())   // Add SetRedeemTokenData command
	cmd.AddCommand(CmdClearRedeemTokenData()) // Add ClearRedeemTokenData command

	cmd.SetHelpTemplate(fmt.Sprintf(`%s
Transaction commands for the photosynthesis module.

Usage:
  %s [command]

Available Commands:
  execute-cron-tasks                   Manually trigger the execution of cron tasks
  reset-total-liquid-stake             Reset the total liquid stake to zero
  set-redeem-token-data                Set redeem token data for a contract
  clear-redeem-token-data              Clear all redeem token data

Use "%s [command] --help" for more information about a command.
`, cmd.Short, cmd.Use, cmd.Use))

	return cmd
}

// CmdExecuteCronTasks defines the command to trigger cron tasks manually
func CmdExecuteCronTasks() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute-cron-tasks",
		Short: "Manually trigger the execution of cron tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get the client context
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Convert the from address to string
			creator := clientCtx.GetFromAddress()

			// Construct the message with the creator's address
			msg := types.NewMsgExecuteCronTasks(creator)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			// Generate or broadcast the transaction
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	// Add transaction flags
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdResetTotalLiquidStake defines the command to reset the total liquid stake
func CmdResetTotalLiquidStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset-total-liquid-stake",
		Short: "Reset the total liquid stake to zero",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get the client context
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Convert the from address to string
			creator := clientCtx.GetFromAddress()

			// Construct the message with the creator's address
			msg := types.NewMsgResetTotalLiquidStake(creator)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			// Generate or broadcast the transaction
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	// Add transaction flags
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdResetAllCompletedDepositRecords defines the command to reset all completed deposit records
func CmdResetAllCompletedDepositRecords() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset-all-completed-deposit-records",
		Short: "Reset all completed deposit records",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get the client context
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Convert the from address to string
			creator := clientCtx.GetFromAddress()

			// Construct the message with the creator's address
			msg := types.NewMsgResetAllCompletedDepositRecords(creator)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			// Generate or broadcast the transaction
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	// Add transaction flags
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdSetRedeemTokenData defines the command to set redeem token data.
func CmdSetRedeemTokenData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-redeem-token-data [contract-address] [amount]",
		Short: "Set the redeem token data for a specific contract",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			contractAddress := args[0]
			amount := args[1]

			msg := types.NewMsgSetRedeemTokenData(creator, contractAddress, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdClearRedeemTokenData defines the command to clear redeem token data.
func CmdClearRedeemTokenData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear-redeem-token-data",
		Short: "Clear all redeem token data from the store",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			msg := types.NewMsgClearRedeemTokenData(creator)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
