package photosynthesis

import (
	"errors"

	"github.com/archway-network/archway/x/photosynthesis/keeper"
	"github.com/archway-network/archway/x/photosynthesis/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler creates a new handler for the photosynthesis module's messages.
func NewHandler(k keeper.PhotosynthesisKeeper) sdk.Handler {
	msgServer := keeper.NewMsgServer(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// Handle MsgExecuteCronTasks.
		case *types.MsgExecuteCronTasks:
			res, err := msgServer.ExecuteCronTasks(sdk.WrapSDKContext(ctx), msg)
			if err != nil {
				return nil, err
			}
			return sdk.WrapServiceResult(ctx, res, err)
			// Handle MsgResetTotalLiquidStake (new message type).
		case *types.MsgResetTotalLiquidStake:
			res, err := msgServer.ResetTotalLiquidStake(sdk.WrapSDKContext(ctx), msg)
			if err != nil {
				return nil, err
			}
			return sdk.WrapServiceResult(ctx, res, err)
			// Handle MsgResetAllCompletedDepositRecords.
		case *types.MsgResetAllCompletedDepositRecords:
			res, err := msgServer.ResetAllCompletedDepositRecords(sdk.WrapSDKContext(ctx), msg)
			if err != nil {
				return nil, err
			}
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, errors.New("invalid message type")
		}
	}
}

// RegisterLegacyAminoCodec registers the module's types on the given LegacyAmino codec.
//func (a AppModuleBasic) RegisterCodec(amino *codec.LegacyAmino) {
//
//}

/*
// DefaultGenesis returns default genesis state as raw bytes for the module.
func (a AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the module.
func (a AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var state types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &state); err != nil {
		return fmt.Errorf("failed to unmarshal x/%s genesis state: %w", types.ModuleName, err)
	}

	return state.Validate()
}
*/

/*
// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, serveMux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), serveMux, types.NewQueryClient(clientCtx)); err != nil {
		panic(fmt.Errorf("registering query handler for x/%s: %w", types.ModuleName, err))
	}
}
*/

/*
// RegisterInvariants registers	cdc        codec.Codec
 the module invariants.
func (a AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	keeper.RegisterInvariants(ir, a.keeper)
}
*/

// RegisterServices registers the module services.
/*
func (a AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterQueryServer(cfg.QueryServer(), keeper.NewMsgServer(a.keeper))
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServer(a.keeper))
}
*/
/*
// InitGenesis performs genesis initialization for the module. It returns no validator updates.
func (a AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, bz json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(bz, &genesisState)

	//a.keeper.InitGenesis(ctx, &genesisState)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the module.
func (a AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	//	state := a.keeper.ExportGenesis(ctx)
	//return cdc.MustMarshalJSON(state)
	return nil
}
*/
