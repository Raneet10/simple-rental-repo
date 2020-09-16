package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var MinPlacePrice = sdk.Coins{sdk.NewInt64Coin("rentaltoken", 1)}

type Rental struct {
	Price   sdk.Coins      `json:"price"`
	Owner   sdk.AccAddress `json:"owner"`
	Current sdk.AccAddress `json:"current"`
	Booked  bool           `json: "booked"`
}

func NewRental() Rental {
	return Rental{
		Price:  MinPlacePrice,
		Booked: false,
	}
}

func (r Rental) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
    Price: %s
    Booked: %s
    Current: %s`, r.Owner, r.Price, r.Booked, r.Current))
}
