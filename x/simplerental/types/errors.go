package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrRentalPlaceDoesNotExist   = sdkerrors.Register(ModuleName, 1, "The Rental Place does not exist")
	ErrRentalPlaceCannotBeBooked = sdkerrors.Register(ModuleName, 2, "The Rental Place is already booked")
	ErrSufficientAdvanceNotPayed = sdkerrors.Register(ModuleName, 3, "Sufficient Advance is not payed")
)
