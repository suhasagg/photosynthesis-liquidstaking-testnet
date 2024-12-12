package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ModuleCdc defines the codec for the module
var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

// RegisterLegacyAminoCodec registers custom types with the provided LegacyAmino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgExecuteCronTasks{}, "photosynthesis/MsgExecuteCronTasks", nil)
	cdc.RegisterConcrete(&MsgResetTotalLiquidStake{}, "photosynthesis/MsgResetTotalLiquidStake", nil)
	cdc.RegisterConcrete(&MsgResetAllCompletedDepositRecords{}, "photosynthesis/MsgResetAllCompletedDepositRecords", nil)
	cdc.RegisterConcrete(&MsgSetRedeemTokenData{}, "photosynthesis/MsgSetRedeemTokenData", nil)
	cdc.RegisterConcrete(&MsgClearRedeemTokenData{}, "photosynthesis/MsgClearRedeemTokenData", nil)
}

// RegisterInterfaces registers the module's custom types with the interface registry.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgExecuteCronTasks{},
		&MsgResetTotalLiquidStake{},
		&MsgResetAllCompletedDepositRecords{},
		&MsgSetRedeemTokenData{},
		&MsgClearRedeemTokenData{},
	)
}

// Initialize the codec and register crypto and SDK types
func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)   // Register cryptographic types
	sdk.RegisterLegacyAminoCodec(amino) // Register Cosmos SDK types
	amino.Seal()                        // Seal the codec to prevent further modifications
}
