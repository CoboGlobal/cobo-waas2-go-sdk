# \TransactionsAPI

All URIs are relative to *https://api.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSmartContractCallTransaction**](TransactionsAPI.md#CreateSmartContractCallTransaction) | **Post** /transactions/contract_call | Create a smart contract call transaction
[**CreateTransferTransaction**](TransactionsAPI.md#CreateTransferTransaction) | **Post** /transactions/transfer | Create a transfer transaction
[**DropTransactionById**](TransactionsAPI.md#DropTransactionById) | **Post** /transactions/{transaction_id}/drop | Drop a transaction by ID
[**EstimateFee**](TransactionsAPI.md#EstimateFee) | **Post** /transactions/estimate_fee | Estimate the fee for transaction
[**GetChainFeePrice**](TransactionsAPI.md#GetChainFeePrice) | **Get** /transactions/fee_price | Get the fee price data for chain and/or token(Hold, TBD after normalize fee settings)
[**GetTransactionById**](TransactionsAPI.md#GetTransactionById) | **Get** /transactions/{transaction_id} | Get transaction information by ID
[**ListTransactions**](TransactionsAPI.md#ListTransactions) | **Get** /transactions | List all transactions
[**ResendTransactionById**](TransactionsAPI.md#ResendTransactionById) | **Post** /transactions/{transaction_id}/resend | Resend a transaction by ID
[**RetryTransactionDoubleCheckById**](TransactionsAPI.md#RetryTransactionDoubleCheckById) | **Post** /transactions/{transaction_id}/callback_confirmation/retry | Retry up a transaction double-check by ID
[**SpeedupTransactionById**](TransactionsAPI.md#SpeedupTransactionById) | **Post** /transactions/{transaction_id}/speedup | Speed up a transaction by ID



## CreateSmartContractCallTransaction

> CreateTransferTransaction201Response CreateSmartContractCallTransaction(ctx).SmartContractCall(smartContractCall).Execute()

Create a smart contract call transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	smartContractCall := *CoboWaas2.NewSmartContractCall("f47ac10b-58cc-4372-a567-0e02b2c3d479", "Transfer", "f47ac10b-58cc-4372-a567-0e02b2c3d479", "19AR6YWEGbSoY8UT9Ksy9WrmrZPD5sL4Ku", "ETH", "bc1q0qfzuge7vr5s2xkczrjkccmxemlyyn8mhx298v", string([B@6aea99e7)) // SmartContractCall | The request body to create a smart contract transaction (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.CreateSmartContractCallTransaction(ctx).SmartContractCall(smartContractCall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.CreateSmartContractCallTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSmartContractCallTransaction`: CreateTransferTransaction201Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.CreateSmartContractCallTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSmartContractCallTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smartContractCall** | [**SmartContractCall**](SmartContractCall.md) | The request body to create a smart contract transaction | 

### Return type

[**CreateTransferTransaction201Response**](CreateTransferTransaction201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTransferTransaction

> CreateTransferTransaction201Response CreateTransferTransaction(ctx).Transfer(transfer).Execute()

Create a transfer transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transfer := *CoboWaas2.NewTransfer("f47ac10b-58cc-4372-a567-0e02b2c3d479", "Transfer", CoboWaas2.TransferSource{BaseTransferSource: CoboWaas2.NewBaseTransferSource(CoboWaas2.WalletSubtype("Asset"), "f47ac10b-58cc-4372-a567-0e02b2c3d479")}, "ETH_USDT", "1.5", CoboWaas2.TransferDestination{AddressTransferDestination: CoboWaas2.NewAddressTransferDestination(CoboWaas2.TransferDestinationType("Address"), "19AR6YWEGbSoY8UT9Ksy9WrmrZPD5sL4Ku")}) // Transfer | The request body to create a transfer transaction (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.CreateTransferTransaction(ctx).Transfer(transfer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.CreateTransferTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTransferTransaction`: CreateTransferTransaction201Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.CreateTransferTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransferTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transfer** | [**Transfer**](Transfer.md) | The request body to create a transfer transaction | 

### Return type

[**CreateTransferTransaction201Response**](CreateTransferTransaction201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DropTransactionById

> CreateTransferTransaction201Response DropTransactionById(ctx, transactionId).TransactionFee(transactionFee).Execute()

Drop a transaction by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the transaction
	transactionFee := CoboWaas2.TransactionFee{EvmEip1559TransactionFee: CoboWaas2.NewEvmEip1559TransactionFee("1", "0.1", "0.9", CoboWaas2.FeeType("Fixed"))} // TransactionFee | The request body of fee to initiate transaction (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.DropTransactionById(ctx, transactionId).TransactionFee(transactionFee).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.DropTransactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DropTransactionById`: CreateTransferTransaction201Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.DropTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**transactionId** | **string** | Unique id of the transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiDropTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionFee** | [**TransactionFee**](TransactionFee.md) | The request body of fee to initiate transaction | 

### Return type

[**CreateTransferTransaction201Response**](CreateTransferTransaction201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EstimateFee

> TransactionFee EstimateFee(ctx).EstimateFee(estimateFee).Execute()

Estimate the fee for transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	estimateFee := CoboWaas2.EstimateFee{SmartContractCall: CoboWaas2.NewSmartContractCall("f47ac10b-58cc-4372-a567-0e02b2c3d479", "Transfer", "f47ac10b-58cc-4372-a567-0e02b2c3d479", "19AR6YWEGbSoY8UT9Ksy9WrmrZPD5sL4Ku", "ETH", "bc1q0qfzuge7vr5s2xkczrjkccmxemlyyn8mhx298v", string([B@6aea99e7))} // EstimateFee | The request body to estimate fee of transfer or call transaction (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.EstimateFee(ctx).EstimateFee(estimateFee).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.EstimateFee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EstimateFee`: TransactionFee
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.EstimateFee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEstimateFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **estimateFee** | [**EstimateFee**](EstimateFee.md) | The request body to estimate fee of transfer or call transaction | 

### Return type

[**TransactionFee**](TransactionFee.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChainFeePrice

> ChainFeePrice GetChainFeePrice(ctx).ChainId(chainId).TokenId(tokenId).Execute()

Get the fee price data for chain and/or token(Hold, TBD after normalize fee settings)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	chainId := "ETH" // string | Unique id of the chain (optional)
	tokenId := "ETH_USDT" // string | Unique id of the token (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.GetChainFeePrice(ctx).ChainId(chainId).TokenId(tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetChainFeePrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChainFeePrice`: ChainFeePrice
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetChainFeePrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChainFeePriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chainId** | **string** | Unique id of the chain | 
 **tokenId** | **string** | Unique id of the token | 

### Return type

[**ChainFeePrice**](ChainFeePrice.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionById

> Transaction GetTransactionById(ctx, transactionId).Execute()

Get transaction information by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the transaction

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.GetTransactionById(ctx, transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionById`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**transactionId** | **string** | Unique id of the transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> ListTransactions200Response ListTransactions(ctx).RequestId(requestId).CoboId(coboId).TransactionId(transactionId).TransactionHash(transactionHash).Type_(type_).Status(status).WalletId(walletId).ChainId(chainId).TokenId(tokenId).AssetId(assetId).MinCreatedTimestamp(minCreatedTimestamp).MaxCreatedTimestamp(maxCreatedTimestamp).SortBy(sortBy).Direction(direction).Limit(limit).Before(before).After(after).Execute()

List all transactions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	requestId := "web_send_by_user_327_1610444045047" // string | Request ID (optional)
	coboId := "20231213122855000000000000000000" // string | Cobo ID (optional)
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the transaction (optional)
	transactionHash := "239861be9a4afe080c359b7fe4a1d035945ec46256b1a0f44d1267c71de8ec28" // string | Transaction hash (optional)
	type_ := CoboWaas2.TransactionType("Deposit") // TransactionType | The type of a transaction (optional)
	status := CoboWaas2.TransactionStatus("Submitted") // TransactionStatus | The status of a transaction (optional)
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the wallet (optional)
	chainId := "ETH" // string | Unique id of the chain (optional)
	tokenId := "ETH_USDT" // string | Unique id of the token (optional)
	assetId := "USDT" // string | Unique id of the asset (optional)
	minCreatedTimestamp := int32(1635744000) // int32 | The minimum transaction creation timestamp in Unix epoch seconds (optional)
	maxCreatedTimestamp := int32(1635744000) // int32 | The maximum transaction creation timestamp in Unix epoch seconds (optional)
	sortBy := "timestamp" // string | Field of sort by (optional) (default to "")
	direction := "ASC" // string | Direction to sort by (optional) (default to "")
	limit := int32(10) // int32 | The maximum number of objects to return. The value range is [1, 50]. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `foo`, the request will retrieve a list of data objects that end before the object with the object ID `foo`. You can set this parameter to the value of `pagination.after` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `bar`, the request will retrieve a list of data objects that start after the object with the object ID `bar`. You can set this parameter to the value of `pagination.before` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.ListTransactions(ctx).RequestId(requestId).CoboId(coboId).TransactionId(transactionId).TransactionHash(transactionHash).Type_(type_).Status(status).WalletId(walletId).ChainId(chainId).TokenId(tokenId).AssetId(assetId).MinCreatedTimestamp(minCreatedTimestamp).MaxCreatedTimestamp(maxCreatedTimestamp).SortBy(sortBy).Direction(direction).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransactions`: ListTransactions200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Request ID | 
 **coboId** | **string** | Cobo ID | 
 **transactionId** | **string** | Unique id of the transaction | 
 **transactionHash** | **string** | Transaction hash | 
 **type_** | [**TransactionType**](TransactionType.md) | The type of a transaction | 
 **status** | [**TransactionStatus**](TransactionStatus.md) | The status of a transaction | 
 **walletId** | **string** | Unique id of the wallet | 
 **chainId** | **string** | Unique id of the chain | 
 **tokenId** | **string** | Unique id of the token | 
 **assetId** | **string** | Unique id of the asset | 
 **minCreatedTimestamp** | **int32** | The minimum transaction creation timestamp in Unix epoch seconds | 
 **maxCreatedTimestamp** | **int32** | The maximum transaction creation timestamp in Unix epoch seconds | 
 **sortBy** | **string** | Field of sort by | [default to &quot;&quot;]
 **direction** | **string** | Direction to sort by | [default to &quot;&quot;]
 **limit** | **int32** | The maximum number of objects to return. The value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;foo&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;foo&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;bar&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;bar&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 

### Return type

[**ListTransactions200Response**](ListTransactions200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendTransactionById

> CreateTransferTransaction201Response ResendTransactionById(ctx, transactionId).Execute()

Resend a transaction by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the transaction

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.ResendTransactionById(ctx, transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ResendTransactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendTransactionById`: CreateTransferTransaction201Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ResendTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**transactionId** | **string** | Unique id of the transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateTransferTransaction201Response**](CreateTransferTransaction201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryTransactionDoubleCheckById

> CreateTransferTransaction201Response RetryTransactionDoubleCheckById(ctx, transactionId).Execute()

Retry up a transaction double-check by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the transaction

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.RetryTransactionDoubleCheckById(ctx, transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.RetryTransactionDoubleCheckById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryTransactionDoubleCheckById`: CreateTransferTransaction201Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.RetryTransactionDoubleCheckById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**transactionId** | **string** | Unique id of the transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryTransactionDoubleCheckByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateTransferTransaction201Response**](CreateTransferTransaction201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpeedupTransactionById

> CreateTransferTransaction201Response SpeedupTransactionById(ctx, transactionId).TransactionFee(transactionFee).Execute()

Speed up a transaction by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the transaction
	transactionFee := CoboWaas2.TransactionFee{EvmEip1559TransactionFee: CoboWaas2.NewEvmEip1559TransactionFee("1", "0.1", "0.9", CoboWaas2.FeeType("Fixed"))} // TransactionFee | The request body of fee to initiate transaction (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.SpeedupTransactionById(ctx, transactionId).TransactionFee(transactionFee).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.SpeedupTransactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpeedupTransactionById`: CreateTransferTransaction201Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.SpeedupTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**transactionId** | **string** | Unique id of the transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpeedupTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionFee** | [**TransactionFee**](TransactionFee.md) | The request body of fee to initiate transaction | 

### Return type

[**CreateTransferTransaction201Response**](CreateTransferTransaction201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

