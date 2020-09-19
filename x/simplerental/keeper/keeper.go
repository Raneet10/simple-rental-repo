package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/raneet10/simple-rental-repo/x/simplerental/types"
)

// Keeper of the simplerental store
type Keeper struct {
	CoinKeeper types.BankKeeper
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

func (k Keeper) SetRentalPlaceDetails(ctx sdk.Context, name string, rental types.Rental) {

	if rental.Owner.Empty() {
		return
	}

	store := ctx.KVStrore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(rental))
}

func (k Keeper) GetRentalPlaceDetails(ctx sdk.Context, name string) types.Rental {

	store := ctx.KVStore(k.storeKey)
	if !k.IsPlacePresent(ctx, name) {
		return types.NewRental()
	}
	bz := store.get([]byte(name))
	var rental types.Rental
	k.cdc.MustUnmarshalBinaryBare(bz, &rental)
	return rental
}

func (k Keeper) IsPlacePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}
