# \ReconciliationAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReconciliationLedger**](ReconciliationAPI.md#GetReconciliationLedger) | **Get** /recon/ledger | Get reconciliation ledger
[**ListReconciliationStatements**](ReconciliationAPI.md#ListReconciliationStatements) | **Get** /recon/statements | List reconciliation daily statements



## GetReconciliationLedger

> GetReconciliationLedger200Response GetReconciliationLedger(ctx).TransactionIds(transactionIds).Execute()

Get reconciliation ledger



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
	transactionIds := "f47ac10b-58cc-4372-a567-0e02b2c3d479,557918d2-632a-4fe1-932f-315711f05fe3"

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
	resp, r, err := apiClient.ReconciliationAPI.GetReconciliationLedger(ctx).TransactionIds(transactionIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReconciliationAPI.GetReconciliationLedger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReconciliationLedger`: GetReconciliationLedger200Response
	fmt.Fprintf(os.Stdout, "Response from `ReconciliationAPI.GetReconciliationLedger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReconciliationLedgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionIds** | **string** | A list of transaction IDs, separated by comma. You can specify 1 to 100 transaction IDs. | 

### Return type

[**GetReconciliationLedger200Response**](GetReconciliationLedger200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReconciliationStatements

> ListReconciliationStatements200Response ListReconciliationStatements(ctx).StartDate(startDate).EndDate(endDate).WalletIds(walletIds).TokenIds(tokenIds).Limit(limit).Before(before).After(after).Execute()

List reconciliation daily statements



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    coboWaas2 "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
    "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2/crypto"
)

func main() {
	startDate := time.Now()
	endDate := time.Now()
	walletIds := "f47ac10b-58cc-4372-a567-0e02b2c3d479,1ddca562-8434-41c9-8809-d437bad9c868"
	tokenIds := "ETH_USDT,ETH_USDC"
	limit := int32(10)
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1"
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk"

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
	resp, r, err := apiClient.ReconciliationAPI.ListReconciliationStatements(ctx).StartDate(startDate).EndDate(endDate).WalletIds(walletIds).TokenIds(tokenIds).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReconciliationAPI.ListReconciliationStatements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReconciliationStatements`: ListReconciliationStatements200Response
	fmt.Fprintf(os.Stdout, "Response from `ReconciliationAPI.ListReconciliationStatements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReconciliationStatementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | The start date of the reconciliation period (inclusive), in YYYY-MM-DD format (UTC). The range between &#x60;start_date&#x60; and &#x60;end_date&#x60; must not exceed 366 days. | 
 **endDate** | **string** | The end date of the reconciliation period (inclusive), in YYYY-MM-DD format (UTC). The range between &#x60;start_date&#x60; and &#x60;end_date&#x60; must not exceed 366 days. | 
 **walletIds** | **string** | A list of wallet IDs, separated by comma. | 
 **tokenIds** | **string** | A list of token IDs, separated by comma. The token ID is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 

### Return type

[**ListReconciliationStatements200Response**](ListReconciliationStatements200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

