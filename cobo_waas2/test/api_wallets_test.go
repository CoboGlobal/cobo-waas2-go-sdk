/*
Cobo Wallet as a Service 2.0

Testing WalletsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package cobo_waas2

import (
	"context"
	"testing"

	coboWaas2 "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
	"github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2/crypto"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func Test_cobo_waas2_WalletsAPIService(t *testing.T) {

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), coboWaas2.ContextServerHost, "https://api[.xxxx].cobo.com/v2")
	ctx = context.WithValue(ctx, coboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})

	t.Run("Test WalletsAPIService CheckAddressChainsValidity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsAPI.CheckAddressChainsValidity(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService CheckAddressValidity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsAPI.CheckAddressValidity(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService CheckAddressesValidity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsAPI.CheckAddressesValidity(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService CreateAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.CreateAddress(ctx, walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService CreateWallet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsAPI.CreateWallet(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService DeleteWalletById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.DeleteWalletById(ctx, walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetChainById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var chainId string

		resp, httpRes, err := apiClient.WalletsAPI.GetChainById(ctx, chainId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetMaxTransferableValue", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetMaxTransferableValue(ctx, walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetTokenById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tokenId string

		resp, httpRes, err := apiClient.WalletsAPI.GetTokenById(ctx, tokenId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetWalletById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetWalletById(ctx, walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService ListAddresses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.ListAddresses(ctx, walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService ListEnabledChains", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsAPI.ListEnabledChains(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService ListEnabledTokens", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsAPI.ListEnabledTokens(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService ListSupportedChains", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsAPI.ListSupportedChains(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService ListSupportedTokens", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsAPI.ListSupportedTokens(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService ListTokenBalancesForAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string
		var address string

		resp, httpRes, err := apiClient.WalletsAPI.ListTokenBalancesForAddress(ctx, walletId, address).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService ListTokenBalancesForWallet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.ListTokenBalancesForWallet(ctx, walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService ListUtxos", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.ListUtxos(ctx, walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService ListWallets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsAPI.ListWallets(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService LockUtxos", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.LockUtxos(ctx, walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService UnlockUtxos", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.UnlockUtxos(ctx, walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService UpdateWalletById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.UpdateWalletById(ctx, walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
