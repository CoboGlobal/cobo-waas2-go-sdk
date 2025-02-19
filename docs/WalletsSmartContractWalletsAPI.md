# \WalletsSmartContractWalletsAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSafeWalletDelegates**](WalletsSmartContractWalletsAPI.md#ListSafeWalletDelegates) | **Post** /wallets/{wallet_id}/smart_contracts/delegates | List Delegates



## ListSafeWalletDelegates

> []CoboSafeDelegate ListSafeWalletDelegates(ctx, walletId).SafeWalletDelegates(safeWalletDelegates).Execute()

List Delegates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    coboWaas2 "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
    "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2/crypto"
)

func main() {
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	safeWalletDelegates := coboWaas2.SafeWalletDelegates{SafeWalletDelegatesContractCall: coboWaas2.NewSafeWalletDelegatesContractCall(coboWaas2.EstimateFeeRequestType("Transfer"))}

	configuration := coboWaas2.NewConfiguration()
	// Initialize the API client
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.Background()

    // Select the development environment. To use the production environment, replace coboWaas2.DevEnv with coboWaas2.ProdEnv
	ctx = context.WithValue(ctx, coboWaas2.ContextEnv, coboWaas2.DevEnv)
    // Replace `<YOUR_PRIVATE_KEY>` with your private key
	ctx = context.WithValue(ctx, coboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_PRIVATE_KEY>",
	})
	resp, r, err := apiClient.WalletsSmartContractWalletsAPI.ListSafeWalletDelegates(ctx, walletId).SafeWalletDelegates(safeWalletDelegates).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsSmartContractWalletsAPI.ListSafeWalletDelegates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSafeWalletDelegates`: []CoboSafeDelegate
	fmt.Fprintf(os.Stdout, "Response from `WalletsSmartContractWalletsAPI.ListSafeWalletDelegates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSafeWalletDelegatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **safeWalletDelegates** | [**SafeWalletDelegates**](SafeWalletDelegates.md) | The request body to query the Delegates of a Safe{Wallet}. | 

### Return type

[**[]CoboSafeDelegate**](CoboSafeDelegate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

