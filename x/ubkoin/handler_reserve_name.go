package ubkoin

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regnull/ubkoin/x/ubkoin/keeper"
	"github.com/regnull/ubkoin/x/ubkoin/types"
)

func handleReserveName(ctx sdk.Context, k keeper.Keeper, msg *types.ReserveName) (*sdk.Result, error) {
	k.ReserveName(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
