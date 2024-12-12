package ante

import (
	"fmt"

	"github.com/archway-network/archway/x/photosynthesis/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewPhotosynthesisAnteHandler returns an AnteHandler for the photosynthesis module.
func NewPhotosynthesisAnteHandler() sdk.AnteHandler {
	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (newCtx sdk.Context, err error) {
		// Process each message in the transaction
		for _, msg := range tx.GetMsgs() {
			switch msg := msg.(type) {
			case *types.MsgExecuteCronTasks:
				// Validate the MsgExecuteCronTasks
				if err := validateMsgExecuteCronTasks(ctx, msg); err != nil {
					return ctx, err
				}
			default:
				// If the message is not of the expected type, return an error
				return ctx, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("unrecognized message type: %T", msg))
			}
		}

		// Proceed with next AnteHandler in the chain (if any)
		return next(ctx, tx, simulate)
	}
}

// validateMsgExecuteCronTasks performs basic validation checks for MsgExecuteCronTasks.
func validateMsgExecuteCronTasks(ctx sdk.Context, msg *types.MsgExecuteCronTasks) error {
	// Ensure that the sender (creator) is a valid Bech32 address
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	// Additional checks (e.g., permissions, balances, etc.) can be added here.
	// Example: You could validate the sender has specific permissions or tokens.

	return nil
}

// next is a placeholder for the next AnteHandler in the chain (if any). This can be replaced with your actual next handler.
func next(ctx sdk.Context, tx sdk.Tx, simulate bool) (newCtx sdk.Context, err error) {
	// This is a placeholder to proceed with the next AnteHandler in the chain if needed.
	return ctx, nil
}
