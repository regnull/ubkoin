package keeper

import (
	"github.com/regnull/ubkoin/x/ubkoin/types"
)

var _ types.QueryServer = Keeper{}
