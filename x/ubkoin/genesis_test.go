package ubkoin_test

import (
	"testing"

	keepertest "github.com/regnull/ubkoin/testutil/keeper"
	"github.com/regnull/ubkoin/x/ubkoin"
	"github.com/regnull/ubkoin/x/ubkoin/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.UbkoinKeeper(t)
	ubkoin.InitGenesis(ctx, *k, genesisState)
	got := ubkoin.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
