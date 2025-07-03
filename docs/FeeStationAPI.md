# \FeeStationAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EstimateFeeStationFee**](FeeStationAPI.md#EstimateFeeStationFee) | **Post** /fee_station/transactions/estimate_fee | Estimate fee for Fee Station transaction
[**GetFeeStationTransactionById**](FeeStationAPI.md#GetFeeStationTransactionById) | **Get** /fee_station/transactions/{transaction_id} | Get Fee Station transaction information
[**ListFeeStationAddresses**](FeeStationAPI.md#ListFeeStationAddresses) | **Get** /fee_station/addresses | List Fee Station addresses
[**ListFeeStationTransactions**](FeeStationAPI.md#ListFeeStationTransactions) | **Get** /fee_station/transactions | List all Fee Station transactions
[**ListTokenBalancesForFeeStation**](FeeStationAPI.md#ListTokenBalancesForFeeStation) | **Get** /fee_station/tokens | List Fee Station token balances



## EstimateFeeStationFee

> EstimatedFixedFee EstimateFeeStationFee(ctx).FeeStationTransfer(feeStationTransfer).Execute()

Estimate fee for Fee Station transaction



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
	feeStationTransfer := *coboWaas2.NewFeeStationTransfer("ETH_USDT")

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
	resp, r, err := apiClient.FeeStationAPI.EstimateFeeStationFee(ctx).FeeStationTransfer(feeStationTransfer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeeStationAPI.EstimateFeeStationFee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EstimateFeeStationFee`: EstimatedFixedFee
	fmt.Fprintf(os.Stdout, "Response from `FeeStationAPI.EstimateFeeStationFee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEstimateFeeStationFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feeStationTransfer** | [**FeeStationTransfer**](FeeStationTransfer.md) | The information about a Fee Station top-up transaction. | 

### Return type

[**EstimatedFixedFee**](EstimatedFixedFee.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeeStationTransactionById

> TransactionDetail GetFeeStationTransactionById(ctx, transactionId).Execute()

Get Fee Station transaction information



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
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"

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
	resp, r, err := apiClient.FeeStationAPI.GetFeeStationTransactionById(ctx, transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeeStationAPI.GetFeeStationTransactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeeStationTransactionById`: TransactionDetail
	fmt.Fprintf(os.Stdout, "Response from `FeeStationAPI.GetFeeStationTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**transactionId** | **string** | The transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeeStationTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransactionDetail**](TransactionDetail.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeeStationAddresses

> ListAddresses200Response ListFeeStationAddresses(ctx).ChainIds(chainIds).Addresses(addresses).Limit(limit).Before(before).After(after).Execute()

List Fee Station addresses



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
	chainIds := "BTC,ETH"
	addresses := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045,0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"
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
	resp, r, err := apiClient.FeeStationAPI.ListFeeStationAddresses(ctx).ChainIds(chainIds).Addresses(addresses).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeeStationAPI.ListFeeStationAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFeeStationAddresses`: ListAddresses200Response
	fmt.Fprintf(os.Stdout, "Response from `FeeStationAPI.ListFeeStationAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFeeStationAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chainIds** | **string** | A list of chain IDs, separated by comma. The chain ID is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | 
 **addresses** | **string** | A list of wallet addresses, separated by comma. For addresses requiring a memo, append the memo after the address using the &#39;|&#39; separator (e.g., \&quot;address|memo\&quot;). | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 

### Return type

[**ListAddresses200Response**](ListAddresses200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeeStationTransactions

> ListTransactions200Response ListFeeStationTransactions(ctx).RequestId(requestId).CoboIds(coboIds).TransactionIds(transactionIds).TransactionHashes(transactionHashes).Types(types).Statuses(statuses).ChainIds(chainIds).TokenIds(tokenIds).AssetIds(assetIds).MinCreatedTimestamp(minCreatedTimestamp).MaxCreatedTimestamp(maxCreatedTimestamp).Limit(limit).Before(before).After(after).Direction(direction).Execute()

List all Fee Station transactions



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
	requestId := "web_send_by_user_327_1610444045047"
	coboIds := "20231213122855000000000000000000,20231213122955000000000000000000"
	transactionIds := "f47ac10b-58cc-4372-a567-0e02b2c3d479,557918d2-632a-4fe1-932f-315711f05fe3"
	transactionHashes := "239861be9a4afe080c359b7fe4a1d035945ec46256b1a0f44d1267c71de8ec28"
	types := "Deposit,Withdrawal"
	statuses := "Completed,Failed"
	chainIds := "BTC,ETH"
	tokenIds := "ETH_USDT,ETH_USDC"
	assetIds := "USDT,USDC"
	minCreatedTimestamp := int64(1635744000000)
	maxCreatedTimestamp := int64(1635744000000)
	limit := int32(10)
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1"
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk"
	direction := "ASC"

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
	resp, r, err := apiClient.FeeStationAPI.ListFeeStationTransactions(ctx).RequestId(requestId).CoboIds(coboIds).TransactionIds(transactionIds).TransactionHashes(transactionHashes).Types(types).Statuses(statuses).ChainIds(chainIds).TokenIds(tokenIds).AssetIds(assetIds).MinCreatedTimestamp(minCreatedTimestamp).MaxCreatedTimestamp(maxCreatedTimestamp).Limit(limit).Before(before).After(after).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeeStationAPI.ListFeeStationTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFeeStationTransactions`: ListTransactions200Response
	fmt.Fprintf(os.Stdout, "Response from `FeeStationAPI.ListFeeStationTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFeeStationTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | 
 **coboIds** | **string** | A list of Cobo IDs, separated by comma. A Cobo ID can be used to track a transaction. | 
 **transactionIds** | **string** | A list of transaction IDs, separated by comma. | 
 **transactionHashes** | **string** | A list of transaction hashes, separated by comma. | 
 **types** | **string** | A list of transaction types for Fee Station, separated by comma. Possible values include:    - &#x60;Deposit&#x60;: A deposit transaction.   - &#x60;Withdrawal&#x60;: A withdrawal transaction.  | 
 **statuses** | **string** | A list of transaction statuses, separated by comma. Possible values include:    - &#x60;Submitted&#x60;: The transaction is submitted.   - &#x60;PendingScreening&#x60;: The transaction is pending screening by Risk Control.    - &#x60;PendingAuthorization&#x60;: The transaction is pending approvals.   - &#x60;PendingSignature&#x60;: The transaction is pending signature.    - &#x60;Broadcasting&#x60;: The transaction is being broadcast.   - &#x60;Confirming&#x60;: The transaction is waiting for the required number of confirmations.   - &#x60;Completed&#x60;: The transaction is completed.   - &#x60;Failed&#x60;: The transaction failed.   - &#x60;Rejected&#x60;: The transaction is rejected.   - &#x60;Pending&#x60;: The transaction is waiting to be included in the next block of the blockchain.  | 
 **chainIds** | **string** | A list of chain IDs, separated by comma. The chain ID is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | 
 **tokenIds** | **string** | A list of token IDs, separated by comma. The token ID is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
 **assetIds** | **string** | (This concept applies to Exchange Wallets only) A list of asset IDs, separated by comma. An asset ID is the unique identifier of the asset held within your linked exchange account. | 
 **minCreatedTimestamp** | **int64** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. You can use this parameter to filter transactions created on or after the specified time.  If not provided, the default value is 90 days before the current time. This default value is subject to change.  | 
 **maxCreatedTimestamp** | **int64** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. You can use this parameter to filter transactions created on or before the specified time.  If not provided, the default value is the current time. This default value is subject to change.  | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 
 **direction** | **string** | The sort direction. Possible values include:   - &#x60;ASC&#x60;: Sort the results in ascending order.   - &#x60;DESC&#x60;: Sort the results in descending order.  | [default to &quot;ASC&quot;]

### Return type

[**ListTransactions200Response**](ListTransactions200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokenBalancesForFeeStation

> ListTokenBalancesForFeeStation200Response ListTokenBalancesForFeeStation(ctx).TokenIds(tokenIds).Limit(limit).Before(before).After(after).Execute()

List Fee Station token balances



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
	resp, r, err := apiClient.FeeStationAPI.ListTokenBalancesForFeeStation(ctx).TokenIds(tokenIds).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeeStationAPI.ListTokenBalancesForFeeStation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTokenBalancesForFeeStation`: ListTokenBalancesForFeeStation200Response
	fmt.Fprintf(os.Stdout, "Response from `FeeStationAPI.ListTokenBalancesForFeeStation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTokenBalancesForFeeStationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenIds** | **string** | A list of token IDs, separated by comma. The token ID is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 

### Return type

[**ListTokenBalancesForFeeStation200Response**](ListTokenBalancesForFeeStation200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

