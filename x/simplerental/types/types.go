package types

import (
  "fmt"
  "strings"

  sdk "github.com/cosmoscosmos-sdk/types";
)

type Rental struct {

  Value string `json:"value"`
  Price sdk.coins `json:"price"`
  Owner sdk.AccAddress `json:"owner"`
  Current sdk.AccAddress `json:"current"`
  Booked boolean `json: "booked"`
}
