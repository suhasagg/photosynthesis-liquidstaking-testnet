package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgExecuteCronTasks{}

// NewMsgExecuteCronTasks creates a new MsgExecuteCronTasks instance.
func NewMsgExecuteCronTasks(creator sdk.AccAddress) *MsgExecuteCronTasks {
	return &MsgExecuteCronTasks{
		Creator: creator.String(),
	}
}

// Route returns the name of the module
func (msg *MsgExecuteCronTasks) Route() string { return RouterKey }

// Type returns the action
func (msg *MsgExecuteCronTasks) Type() string { return "ExecuteCronTasks" }

// GetSigners defines whose signature is required
func (msg *MsgExecuteCronTasks) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg *MsgExecuteCronTasks) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic performs stateless validation checks
func (msg *MsgExecuteCronTasks) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator cannot be empty")
	}
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgResetTotalLiquidStake{}

// NewMsgResetTotalLiquidStake creates a new MsgResetTotalLiquidStake instance.
func NewMsgResetTotalLiquidStake(creator sdk.AccAddress) *MsgResetTotalLiquidStake {
	return &MsgResetTotalLiquidStake{
		Creator: creator.String(),
	}
}

// Route returns the name of the module
func (msg *MsgResetTotalLiquidStake) Route() string { return RouterKey }

// Type returns the action
func (msg *MsgResetTotalLiquidStake) Type() string { return "ResetTotalLiquidStake" }

// GetSigners defines whose signature is required
func (msg *MsgResetTotalLiquidStake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg *MsgResetTotalLiquidStake) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic performs stateless validation checks
func (msg *MsgResetTotalLiquidStake) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator cannot be empty")
	}
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// NewMsgResetAllCompletedDepositRecords creates a new MsgResetAllCompletedDepositRecords instance.
func NewMsgResetAllCompletedDepositRecords(creator sdk.AccAddress) *MsgResetAllCompletedDepositRecords {
	return &MsgResetAllCompletedDepositRecords{
		Creator: creator.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg *MsgResetAllCompletedDepositRecords) Route() string {
	return RouterKey
}

// Type implements the sdk.Msg interface.
func (msg *MsgResetAllCompletedDepositRecords) Type() string {
	return "ResetAllCompletedDepositRecords"
}

// GetSigners implements the sdk.Msg interface.
func (msg *MsgResetAllCompletedDepositRecords) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg *MsgResetAllCompletedDepositRecords) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg *MsgResetAllCompletedDepositRecords) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator cannot be empty")
	}
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgSetRedeemTokenData(creator, contractAddr, amount string) *MsgSetRedeemTokenData {
	return &MsgSetRedeemTokenData{
		Creator:         creator,
		ContractAddress: contractAddr,
		Amount:          amount,
	}
}

func (msg *MsgSetRedeemTokenData) Route() string { return RouterKey }
func (msg *MsgSetRedeemTokenData) Type() string  { return "SetRedeemTokenData" }

func (msg *MsgSetRedeemTokenData) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}
	if _, err := sdk.AccAddressFromBech32(msg.ContractAddress); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid contract address")
	}
	if _, ok := sdk.NewIntFromString(msg.Amount); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}
	return nil
}

// GetSigners defines whose signature is required
func (msg *MsgSetRedeemTokenData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func NewMsgClearRedeemTokenData(creator string) *MsgClearRedeemTokenData {
	return &MsgClearRedeemTokenData{Creator: creator}
}

func (msg *MsgClearRedeemTokenData) Route() string { return RouterKey }
func (msg *MsgClearRedeemTokenData) Type() string  { return "ClearRedeemTokenData" }

func (msg *MsgClearRedeemTokenData) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}
	return nil
}

// GetSigners defines whose signature is required
func (msg *MsgClearRedeemTokenData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}
