package v4_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	v4 "github.com/irisnet/irismod/modules/coinswap/migrations/v4"
	coinswaptypes "github.com/irisnet/irismod/modules/coinswap/types"
	"github.com/irisnet/irismod/simapp"
)

func TestMigrate(t *testing.T) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	legacySubspace := app.GetSubspace(coinswaptypes.ModuleName)

	params := coinswaptypes.DefaultParams()
	legacySubspace.SetParamSet(ctx, &params)

	err := v4.Migrate(
		ctx,
		app.CoinswapKeeper,
		legacySubspace,
	)
	require.NoError(t, err)

	expParams := app.CoinswapKeeper.GetParams(ctx)
	require.Equal(t, expParams, params, "v4.Migrate failed")

}
