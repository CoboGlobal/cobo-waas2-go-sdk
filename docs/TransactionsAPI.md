# \TransactionsAPI

All URIs are relative to *https://api.cobo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSignMessageTransaction**](TransactionsAPI.md#CreateSignMessageTransaction) | **Post** /transactions/sign | Create a sign message transaction
[**CreateSmartContractCallTransaction**](TransactionsAPI.md#CreateSmartContractCallTransaction) | **Post** /transactions/call | Create a smart contract call transaction
[**CreateTransferTransaction**](TransactionsAPI.md#CreateTransferTransaction) | **Post** /transactions/transfer | Create a transfer transaction
[**DropTransactionById**](TransactionsAPI.md#DropTransactionById) | **Post** /transactions/{transaction_id}/drop | Drop a transaction by ID
[**EstimateFee**](TransactionsAPI.md#EstimateFee) | **Post** /transactions/estimate_fee | Estimate the fee for transaction
[**GetChainFeePrice**](TransactionsAPI.md#GetChainFeePrice) | **Get** /transactions/fee_price | Get the fee price data for chain and/or token(Hold, TBD after normalize fee settings)
[**GetTransactionById**](TransactionsAPI.md#GetTransactionById) | **Get** /transactions/{transaction_id} | Get transaction information by ID
[**ListTransactions**](TransactionsAPI.md#ListTransactions) | **Get** /transactions | List all transactions
[**ResendTransactionById**](TransactionsAPI.md#ResendTransactionById) | **Post** /transactions/{transaction_id}/resend | Resend a transaction by ID
[**RetryTransactionDoubleCheckById**](TransactionsAPI.md#RetryTransactionDoubleCheckById) | **Post** /transactions/{transaction_id}/double_check/retry | Retry up a transaction double-check by ID
[**SpeedupTransactionById**](TransactionsAPI.md#SpeedupTransactionById) | **Post** /transactions/{transaction_id}/speedup | Speed up a transaction by ID
[**UpdateTransacitonById**](TransactionsAPI.md#UpdateTransacitonById) | **Put** /transactions/{transaction_id} | Update transaction by ID



## CreateSignMessageTransaction

> CreateTransferTransaction201Response CreateSignMessageTransaction(ctx).SignMessage(signMessage).Execute()

Create a sign message transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	signMessage := *openapiclient.NewSignMessage("f47ac10b-58cc-4372-a567-0e02b2c3d479", "Transfer", "ETH") // SignMessage | The request body to create a message sign transaction (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.CreateSignMessageTransaction(context.Background()).SignMessage(signMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.CreateSignMessageTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSignMessageTransaction`: CreateTransferTransaction201Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.CreateSignMessageTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSignMessageTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signMessage** | [**SignMessage**](SignMessage.md) | The request body to create a message sign transaction | 

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
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	smartContractCall := *openapiclient.NewSmartContractCall("f47ac10b-58cc-4372-a567-0e02b2c3d479", "Transfer", "f47ac10b-58cc-4372-a567-0e02b2c3d479", "19AR6YWEGbSoY8UT9Ksy9WrmrZPD5sL4Ku", "ETH", "bc1q0qfzuge7vr5s2xkczrjkccmxemlyyn8mhx298v", string([B@47058864)) // SmartContractCall | The request body to create a smart contract transaction (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.CreateSmartContractCallTransaction(context.Background()).SmartContractCall(smartContractCall).Execute()
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
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	transfer := *openapiclient.NewTransfer("f47ac10b-58cc-4372-a567-0e02b2c3d479", "Transfer", openapiclient.TransferSource{BaseTransferSource: openapiclient.NewBaseTransferSource(openapiclient.WalletSubtype("Asset"), "f47ac10b-58cc-4372-a567-0e02b2c3d479")}, "ETH_USDT", "1.5", openapiclient.TransferDestination{AddressTransferDestination: openapiclient.NewAddressTransferDestination(openapiclient.TransferDestinationType("Address"), "19AR6YWEGbSoY8UT9Ksy9WrmrZPD5sL4Ku")}) // Transfer | The request body to create a transfer transaction (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.CreateTransferTransaction(context.Background()).Transfer(transfer).Execute()
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
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the transaction
	transactionFee := openapiclient.TransactionFee{EvmEip1559TransactionFee: openapiclient.NewEvmEip1559TransactionFee("1", int32(123), int32(123), openapiclient.FeeType("Fixed"))} // TransactionFee | The request body of fee to initiate transaction (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.DropTransactionById(context.Background(), transactionId).TransactionFee(transactionFee).Execute()
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	estimateFee := openapiclient.EstimateFee{SmartContractCall: openapiclient.NewSmartContractCall("f47ac10b-58cc-4372-a567-0e02b2c3d479", "Transfer", "f47ac10b-58cc-4372-a567-0e02b2c3d479", "19AR6YWEGbSoY8UT9Ksy9WrmrZPD5sL4Ku", "ETH", "bc1q0qfzuge7vr5s2xkczrjkccmxemlyyn8mhx298v", string([B@47058864))} // EstimateFee | The request body to estimate fee of transfer or call transaction (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.EstimateFee(context.Background()).EstimateFee(estimateFee).Execute()
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
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	chainId := "ETH" // string | Unique id of the chain (optional)
	tokenId := "ETH_USDT" // string | Unique id of the token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetChainFeePrice(context.Background()).ChainId(chainId).TokenId(tokenId).Execute()
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

> TransactionDetails GetTransactionById(ctx, transactionId).Execute()

Get transaction information by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionById(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionById`: TransactionDetails
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Unique id of the transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransactionDetails**](TransactionDetails.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> ListTransactions200Response ListTransactions(ctx).RequestId(requestId).SortBy(sortBy).Direction(direction).Limit(limit).Before(before).After(after).Execute()

List all transactions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	requestId := "web_send_by_user_327_1610444045047" // string | Request ID (optional)
	sortBy := "timestamp" // string | Field of sort by (optional) (default to "")
	direction := "ASC" // string | Direction to sort by (optional) (default to "")
	limit := int32(10) // int32 | size of page to return (pagination) (optional) (default to 10)
	before := "123" // string | Cursor string received from previous request (optional) (default to "")
	after := "123" // string | Cursor string received from previous request (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.ListTransactions(context.Background()).RequestId(requestId).SortBy(sortBy).Direction(direction).Limit(limit).Before(before).After(after).Execute()
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
 **sortBy** | **string** | Field of sort by | [default to &quot;&quot;]
 **direction** | **string** | Direction to sort by | [default to &quot;&quot;]
 **limit** | **int32** | size of page to return (pagination) | [default to 10]
 **before** | **string** | Cursor string received from previous request | [default to &quot;&quot;]
 **after** | **string** | Cursor string received from previous request | [default to &quot;&quot;]

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
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.ResendTransactionById(context.Background(), transactionId).Execute()
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.RetryTransactionDoubleCheckById(context.Background(), transactionId).Execute()
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the transaction
	transactionFee := openapiclient.TransactionFee{EvmEip1559TransactionFee: openapiclient.NewEvmEip1559TransactionFee("1", int32(123), int32(123), openapiclient.FeeType("Fixed"))} // TransactionFee | The request body of fee to initiate transaction (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.SpeedupTransactionById(context.Background(), transactionId).TransactionFee(transactionFee).Execute()
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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


## UpdateTransacitonById

> TransactionDetails UpdateTransacitonById(ctx, transactionId).TransactionDetails(transactionDetails).Execute()

Update transaction by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the transaction
	transactionDetails := *openapiclient.NewTransactionDetails() // TransactionDetails | The request body to update a address (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.UpdateTransacitonById(context.Background(), transactionId).TransactionDetails(transactionDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.UpdateTransacitonById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTransacitonById`: TransactionDetails
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.UpdateTransacitonById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Unique id of the transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransacitonByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionDetails** | [**TransactionDetails**](TransactionDetails.md) | The request body to update a address | 

### Return type

[**TransactionDetails**](TransactionDetails.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

