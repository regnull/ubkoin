package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regnull/ubkoin/x/ubkoin/types"
)

func (k Keeper) GetName(ctx sdk.Context, name string) *types.NameRecord {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameRecordKey))
	nameRec := &types.NameRecord{}
	k.cdc.MustUnmarshal(store.Get(types.KeyPrefix(types.NameRecordKey+name)), nameRec)
	fmt.Printf("Got name: %+v\n", nameRec)
	return nameRec
}

func (k Keeper) GetAllName(ctx sdk.Context) (names []*types.NameRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameRecordKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.NameRecordKey))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		name := &types.NameRecord{}
		k.cdc.MustUnmarshal(iterator.Value(), name)
		names = append(names, name)
	}
	return
}
