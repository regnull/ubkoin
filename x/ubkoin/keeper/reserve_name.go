// x/blog/keeper/post.go
package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regnull/ubkoin/x/ubkoin/types"
)

func (k Keeper) ReserveName(ctx sdk.Context, msg types.ReserveName) {
	rec := &types.NameRecord{
		Name:     msg.GetName(),
		OwnerKey: []byte(msg.GetOwnerAddress()),
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameRecordKey))
	key := types.KeyPrefix(types.NameRecordKey + msg.GetName())
	value := k.cdc.MustMarshal(rec)
	store.Set(key, value)
}
