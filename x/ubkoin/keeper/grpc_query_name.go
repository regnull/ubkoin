// x/block/keeper/grpc_query_post.go
package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/regnull/ubkoin/x/ubkoin/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NameAll(c context.Context, req *types.QueryAllNameRequest) (*types.QueryAllNameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var names []*types.NameRecord
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, types.KeyPrefix(types.NameRecordKey))

	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var name types.NameRecord
		if err := k.cdc.Unmarshal(value, &name); err != nil {
			return err
		}

		names = append(names, &name)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNameResponse{Record: names, Pagination: pageRes}, nil
}

func (k Keeper) Name(c context.Context, req *types.QueryGetNameRequest) (*types.QueryGetNameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var name types.NameRecord
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameRecordKey))
	k.cdc.MustUnmarshal(store.Get(types.KeyPrefix(types.NameRecordKey+req.Name)), &name)
	fmt.Printf("Got name: %+v\n", name)

	return &types.QueryGetNameResponse{Record: &name}, nil
}
