package simplerental

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the simplerental type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {

		case MsgSetPlaceRent:
			return handleMsgSetPlaceRent(ctx, keeper, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handle<Action> does x
func handleMsgSetPlaceRent(ctx sdk.Context, k Keeper, msg Msg) (*sdk.Result, error) {

	if !msg.Owner.Equals(k.GetOwner(ctx, msg.Name)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	// msg events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventSetRent,
			sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	)

	rent, _ := strconv.Atoi(msg.RentPerDay)
	k.SetRentPerDay(ctx, msg.Name, sdk.Coins{sdk.NewInt64Coin("rentaltoken", rent)})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
