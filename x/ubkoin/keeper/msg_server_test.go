package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/regnull/ubkoin/testutil/keeper"
	"github.com/regnull/ubkoin/x/ubkoin/keeper"
	"github.com/regnull/ubkoin/x/ubkoin/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.UbkoinKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
