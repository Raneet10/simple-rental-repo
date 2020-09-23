package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// verify interface at compile time
var _ sdk.Msg = &MsgSetPlaceRent{}

//const RouterKey = ModuleName

// Msg<Action> - struct for unjailing jailed validator
type MsgSetPlaceRent struct {
	Name       string         `json:"name"`
	RentPerDay string         `json:"rent"`
	Owner      sdk.AccAddress `json:"owner"`
}

// NewMsg<Action> creates a new Msg<Action> instance
func NewMsgSetPlaceRent(Name string, RentPerDay string, Owner sdk.AccAddress) MsgSetPlaceRent {
	return MsgSetPlaceRent{
		Name:       Name,
		RentPerDay: RentPerDay,
		Owner:      Owner,
	}
}

// nolint
func (msg MsgSetPlaceRent) Route() string { return RouterKey }
func (msg MsgSetPlaceRent) Type() string  { return "set_place_rent" }

func (msg MsgSetPlaceRent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgSetPlaceRent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgSetPlaceRent) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}

	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Place Name can't be empty ")
	}

	rent, _ := strconv.Atoi(msg.RentPerDay)
	if len(msg.RentPerDay) == 0 || rent <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Rent can't be empty, zero and/or negative ")

	}

	//IF THE ABOVE DOESN'T WORK, USE THIS AND CHANGE RentPerDay TO sdk.Coins
	/*if !msg.RentPerDay.IsAllPositive() || msg.RentPerDay.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Rent can't be empty and/or zero")
	}*/

	return nil
}
