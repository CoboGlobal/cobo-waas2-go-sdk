/*
Cobo Wallet as a Service 2.0

Testing PrimeBrokerAPIService

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

func Test_cobo_waas2_PrimeBrokerAPIService(t *testing.T) {

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), coboWaas2.ContextServerHost, "https://api[.xxxx].cobo.com/v2")
	ctx = context.WithValue(ctx, coboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})

	t.Run("Test PrimeBrokerAPIService ChangeGuardPubkey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		resp, httpRes, err := apiClient.PrimeBrokerAPI.ChangeGuardPubkey(ctx, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrimeBrokerAPIService CreateGuardPubkey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		resp, httpRes, err := apiClient.PrimeBrokerAPI.CreateGuardPubkey(ctx, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrimeBrokerAPIService CreatePrimeBrokerAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		resp, httpRes, err := apiClient.PrimeBrokerAPI.CreatePrimeBrokerAddress(ctx, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrimeBrokerAPIService DeleteGuardPubkey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		resp, httpRes, err := apiClient.PrimeBrokerAPI.DeleteGuardPubkey(ctx, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrimeBrokerAPIService QueryApprovalStatement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var statementId string

		resp, httpRes, err := apiClient.PrimeBrokerAPI.QueryApprovalStatement(ctx, statementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PrimeBrokerAPIService QueryGuardPubkey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		resp, httpRes, err := apiClient.PrimeBrokerAPI.QueryGuardPubkey(ctx, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
