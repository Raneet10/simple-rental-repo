package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

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

func NewKeeper(CoinKeeper types.BankKeeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {

	return Keeper{
		CoinKeeper: CoinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("Module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetRentalPlaceDetails(ctx sdk.Context, name string, rental types.Rental) {

	if rental.Owner.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(rental))
}

func (k Keeper) GetRentalPlaceDetails(ctx sdk.Context, name string) types.Rental {

	store := ctx.KVStore(k.storeKey)
	if !k.IsPlacePresent(ctx, name) {
		return types.NewRental()
	}
	bz := store.Get([]byte(name))
	var rental types.Rental
	k.cdc.MustUnmarshalBinaryBare(bz, &rental)
	return rental
}

func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {

	return k.GetRentalPlaceDetails(ctx, name).Price
}

func (k Keeper) GetRentPerDay(ctx sdk.Context, name string) sdk.Coins {

	return k.GetRentalPlaceDetails(ctx, name).RentPerDay
}

func (k Keeper) GetAdvancePrice(ctx sdk.Context, name string) sdk.Coins {

	return k.GetRentalPlaceDetails(ctx, name).Advance
}

func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {

	return k.GetRentalPlaceDetails(ctx, name).Owner
}

func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {

	return !k.GetRentalPlaceDetails(ctx, name).Owner.Empty()
}

func (k Keeper) GetCurrent(ctx sdk.Context, name string) sdk.AccAddress {

	return k.GetRentalPlaceDetails(ctx, name).Current
}

func (k Keeper) GetIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

func (k Keeper) IsPlaceBooked(ctx sdk.Context, name string) bool {

	return k.GetRentalPlaceDetails(ctx, name).Booked
}

func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {

	rental := k.GetRentalPlaceDetails(ctx, name)
	rental.Price = price
	k.SetRentalPlaceDetails(ctx, name, rental)
}

func (k Keeper) SetRentPerDay(ctx sdk.Context, name string, rent sdk.Coins) {

	rental := k.GetRentalPlaceDetails(ctx, name)
	rental.RentPerDay = rent
	k.SetRentalPlaceDetails(ctx, name, rental)
}

func (k Keeper) SetAdvance(ctx sdk.Context, name string, advance sdk.Coins) {

	rental := k.GetRentalPlaceDetails(ctx, name)
	rental.Advance = advance
	k.SetRentalPlaceDetails(ctx, name, rental)
}

func (k Keeper) SetOwner(ctx sdk.Context, name string, newOwner sdk.AccAddress) {

	rental := k.GetRentalPlaceDetails(ctx, name)
	rental.Owner = newOwner
	k.SetRentalPlaceDetails(ctx, name, rental)
}

func (k Keeper) SetCurrent(ctx sdk.Context, name string, current sdk.AccAddress) {

	rental := k.GetRentalPlaceDetails(ctx, name)
	rental.Current = current
	k.SetRentalPlaceDetails(ctx, name, rental)
}

func (k Keeper) SetBookedStatus(ctx sdk.Context, name string, booked bool) {

	rental := k.GetRentalPlaceDetails(ctx, name)
	rental.Booked = booked
	k.SetRentalPlaceDetails(ctx, name, rental)
}

func (k Keeper) DeletePlace(ctx sdk.Context, name string) {

	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(name))
}

func (k Keeper) IsPlacePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}
