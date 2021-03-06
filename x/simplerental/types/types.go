package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var MinPlacePrice = sdk.Coins{sdk.NewInt64Coin("pricetoken", 5)}
var MinPlaceRent = sdk.Coins{sdk.NewInt64Coin("rentaltoken", 1)}

type Rental struct {
	Price      sdk.Coins      `json:"price"`
	RentPerDay sdk.Coins      `json:"rentperday"`
	Advance    sdk.Coins      `json:"advance"`
	Owner      sdk.AccAddress `json:"owner"`
	Current    sdk.AccAddress `json:"current"`
	Booked     bool           `json: "booked"`
}

func NewRental() Rental {
	return Rental{
		Price:      MinPlacePrice,
		RentPerDay: MinPlaceRent,
		Booked:     false,
	}
}

func (r Rental) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
    Price: %s
    Booked: %t
    Current: %s
		RentPerDay: %s
		Advance: %s`, r.Owner, r.Price, r.Booked, r.Current, r.RentPerDay, r.Advance))
}
