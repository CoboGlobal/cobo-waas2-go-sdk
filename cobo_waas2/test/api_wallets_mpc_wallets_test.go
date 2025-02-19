/*
Cobo Wallet as a Service 2.0

Testing WalletsMPCWalletsAPIService

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

func Test_cobo_waas2_WalletsMPCWalletsAPIService(t *testing.T) {

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), coboWaas2.ContextServerHost, "https://api[.xxxx].cobo.com/v2")
	ctx = context.WithValue(ctx, coboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})

	t.Run("Test WalletsMPCWalletsAPIService CancelTssRequestById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vaultId string
		var tssRequestId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.CancelTssRequestById(ctx, vaultId, tssRequestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService CreateKeyShareHolderGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vaultId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.CreateKeyShareHolderGroup(ctx, vaultId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService CreateMpcProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.CreateMpcProject(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService CreateMpcVault", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.CreateMpcVault(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService CreateTssRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vaultId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.CreateTssRequest(ctx, vaultId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService DeleteKeyShareHolderGroupById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vaultId string
		var keyShareHolderGroupId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.DeleteKeyShareHolderGroupById(ctx, vaultId, keyShareHolderGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService GetKeyShareHolderByTssNodeId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vaultId string
		var tssNodeId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.GetKeyShareHolderByTssNodeId(ctx, vaultId, tssNodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService GetKeyShareHolderGroupById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vaultId string
		var keyShareHolderGroupId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.GetKeyShareHolderGroupById(ctx, vaultId, keyShareHolderGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService GetMpcProjectById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.GetMpcProjectById(ctx, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService GetMpcVaultById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vaultId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.GetMpcVaultById(ctx, vaultId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService GetTssRequestById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vaultId string
		var tssRequestId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.GetTssRequestById(ctx, vaultId, tssRequestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService ListCoboKeyHolders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.ListCoboKeyHolders(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService ListKeyShareHolderGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vaultId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.ListKeyShareHolderGroups(ctx, vaultId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService ListKeyShareHolders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vaultId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.ListKeyShareHolders(ctx, vaultId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService ListMpcProjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.ListMpcProjects(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService ListMpcVaults", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.ListMpcVaults(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService ListTssRequests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vaultId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.ListTssRequests(ctx, vaultId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService UpdateKeyShareHolderGroupById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vaultId string
		var keyShareHolderGroupId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.UpdateKeyShareHolderGroupById(ctx, vaultId, keyShareHolderGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService UpdateMpcProjectById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.UpdateMpcProjectById(ctx, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsMPCWalletsAPIService UpdateMpcVaultById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vaultId string

		resp, httpRes, err := apiClient.WalletsMPCWalletsAPI.UpdateMpcVaultById(ctx, vaultId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
