# \TransactionsAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BroadcastSignedTransactions**](TransactionsAPI.md#BroadcastSignedTransactions) | **Post** /transactions/broadcast | Broadcast signed transactions
[**CancelTransactionById**](TransactionsAPI.md#CancelTransactionById) | **Post** /transactions/{transaction_id}/cancel | Cancel transaction
[**CheckLoopTransfers**](TransactionsAPI.md#CheckLoopTransfers) | **Get** /transactions/check_loop_transfers | Check Cobo Loop transfers
[**CreateContractCallTransaction**](TransactionsAPI.md#CreateContractCallTransaction) | **Post** /transactions/contract_call | Call smart contract
[**CreateMessageSignTransaction**](TransactionsAPI.md#CreateMessageSignTransaction) | **Post** /transactions/message_sign | Sign message
[**CreateTransferTransaction**](TransactionsAPI.md#CreateTransferTransaction) | **Post** /transactions/transfer | Transfer token
[**DropTransactionById**](TransactionsAPI.md#DropTransactionById) | **Post** /transactions/{transaction_id}/drop | Drop transaction
[**EstimateFee**](TransactionsAPI.md#EstimateFee) | **Post** /transactions/estimate_fee | Estimate transaction fee
[**GetTransactionApprovalDetail**](TransactionsAPI.md#GetTransactionApprovalDetail) | **Get** /transactions/{transaction_id}/approval_detail | Get transaction approval details
[**GetTransactionById**](TransactionsAPI.md#GetTransactionById) | **Get** /transactions/{transaction_id} | Get transaction information
[**ListTransactionApprovalDetails**](TransactionsAPI.md#ListTransactionApprovalDetails) | **Get** /transactions/approval_details | List transaction approval details
[**ListTransactions**](TransactionsAPI.md#ListTransactions) | **Get** /transactions | List all transactions
[**ResendTransactionById**](TransactionsAPI.md#ResendTransactionById) | **Post** /transactions/{transaction_id}/resend | Resend transaction
[**SignAndBroadcastTransactionById**](TransactionsAPI.md#SignAndBroadcastTransactionById) | **Post** /transactions/{transaction_id}/sign_and_broadcast | Sign and broadcast transaction
[**SpeedupTransactionById**](TransactionsAPI.md#SpeedupTransactionById) | **Post** /transactions/{transaction_id}/speedup | Speed up transaction



## BroadcastSignedTransactions

> []BroadcastSignedTransactions201ResponseInner BroadcastSignedTransactions(ctx).BroadcastSignedTransactionsRequest(broadcastSignedTransactionsRequest).Execute()

Broadcast signed transactions



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
	broadcastSignedTransactionsRequest := *coboWaas2.NewBroadcastSignedTransactionsRequest()

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
	resp, r, err := apiClient.TransactionsAPI.BroadcastSignedTransactions(ctx).BroadcastSignedTransactionsRequest(broadcastSignedTransactionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.BroadcastSignedTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastSignedTransactions`: []BroadcastSignedTransactions201ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.BroadcastSignedTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastSignedTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **broadcastSignedTransactionsRequest** | [**BroadcastSignedTransactionsRequest**](BroadcastSignedTransactionsRequest.md) | The request body to broadcast a list of signed transactions. | 

### Return type

[**[]BroadcastSignedTransactions201ResponseInner**](BroadcastSignedTransactions201ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelTransactionById

> CreateTransferTransaction201Response CancelTransactionById(ctx, transactionId).Execute()

Cancel transaction



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
	resp, r, err := apiClient.TransactionsAPI.CancelTransactionById(ctx, transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.CancelTransactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelTransactionById`: CreateTransferTransaction201Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.CancelTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**transactionId** | **string** | The transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateTransferTransaction201Response**](CreateTransferTransaction201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckLoopTransfers

> []CheckLoopTransfers200ResponseInner CheckLoopTransfers(ctx).TokenId(tokenId).SourceWalletId(sourceWalletId).DestinationAddresses(destinationAddresses).Execute()

Check Cobo Loop transfers



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
	tokenId := "ETH_USDT"
	sourceWalletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	destinationAddresses := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045,0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"

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
	resp, r, err := apiClient.TransactionsAPI.CheckLoopTransfers(ctx).TokenId(tokenId).SourceWalletId(sourceWalletId).DestinationAddresses(destinationAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.CheckLoopTransfers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckLoopTransfers`: []CheckLoopTransfers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.CheckLoopTransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckLoopTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
 **sourceWalletId** | **string** | The wallet ID of the transaction source. | 
 **destinationAddresses** | **string** | A list of destination addresses, separated by comma. | 

### Return type

[**[]CheckLoopTransfers200ResponseInner**](CheckLoopTransfers200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContractCallTransaction

> CreateTransferTransaction201Response CreateContractCallTransaction(ctx).ContractCallParams(contractCallParams).Execute()

Call smart contract



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
	contractCallParams := *coboWaas2.NewContractCallParams("f47ac10b-58cc-4372-a567-0e02b2c3d479", "ETH", coboWaas2.ContractCallSource{CustodialWeb3ContractCallSource: coboWaas2.NewCustodialWeb3ContractCallSource(coboWaas2.ContractCallSourceType("Web3"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")}, coboWaas2.ContractCallDestination{EvmContractCallDestination: coboWaas2.NewEvmContractCallDestination(coboWaas2.ContractCallDestinationType("EVM_Contract"), "0x0406db8351aa6839169bb363f63c2c808fee8f99", "0xa22cb4650000000000000000000000001e0049783f008a0085193e00003d00cd54003c71000000000000000000000000000000000000000000000000000000000000DEMO")})

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
	resp, r, err := apiClient.TransactionsAPI.CreateContractCallTransaction(ctx).ContractCallParams(contractCallParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.CreateContractCallTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContractCallTransaction`: CreateTransferTransaction201Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.CreateContractCallTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContractCallTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractCallParams** | [**ContractCallParams**](ContractCallParams.md) | The request body for making a contract call. | 

### Return type

[**CreateTransferTransaction201Response**](CreateTransferTransaction201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessageSignTransaction

> CreateTransferTransaction201Response CreateMessageSignTransaction(ctx).MessageSignParams(messageSignParams).Execute()

Sign message



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
	messageSignParams := *coboWaas2.NewMessageSignParams("f47ac10b-58cc-4372-a567-0e02b2c3d479", "ETH", coboWaas2.MessageSignSource{CustodialWeb3MessageSignSource: coboWaas2.NewCustodialWeb3MessageSignSource(coboWaas2.MessageSignSourceType("Web3"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "19AR6YWEGbSoY8UT9Ksy9WrmrZPD5sL4Ku")}, coboWaas2.MessageSignDestination{BTCEIP191MessageSignDestination: coboWaas2.NewBTCEIP191MessageSignDestination(coboWaas2.MessageSignDestinationType("EVM_EIP_191_Signature"), "YWFhYQ==")})

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
	resp, r, err := apiClient.TransactionsAPI.CreateMessageSignTransaction(ctx).MessageSignParams(messageSignParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.CreateMessageSignTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMessageSignTransaction`: CreateTransferTransaction201Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.CreateMessageSignTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageSignTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageSignParams** | [**MessageSignParams**](MessageSignParams.md) | The request body to create a message signing transaction | 

### Return type

[**CreateTransferTransaction201Response**](CreateTransferTransaction201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTransferTransaction

> CreateTransferTransaction201Response CreateTransferTransaction(ctx).TransferParams(transferParams).Execute()

Transfer token



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
	transferParams := *coboWaas2.NewTransferParams("f47ac10b-58cc-4372-a567-0e02b2c3d479", coboWaas2.TransferSource{CustodialTransferSource: coboWaas2.NewCustodialTransferSource(coboWaas2.WalletSubtype("Asset"), "f47ac10b-58cc-4372-a567-0e02b2c3d479")}, "ETH_USDT", coboWaas2.TransferDestination{AddressTransferDestination: coboWaas2.NewAddressTransferDestination(coboWaas2.TransferDestinationType("Address"))})

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
	resp, r, err := apiClient.TransactionsAPI.CreateTransferTransaction(ctx).TransferParams(transferParams).Execute()
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
 **transferParams** | [**TransferParams**](TransferParams.md) | The request body to create a transfer transaction | 

### Return type

[**CreateTransferTransaction201Response**](CreateTransferTransaction201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DropTransactionById

> CreateTransferTransaction201Response DropTransactionById(ctx, transactionId).TransactionRbf(transactionRbf).Execute()

Drop transaction



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
	transactionRbf := *coboWaas2.NewTransactionRbf("f47ac10b-58cc-4372-a567-0e02b2c3d479", coboWaas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: coboWaas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", coboWaas2.FeeType("Fixed"), "ETH")})

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
	resp, r, err := apiClient.TransactionsAPI.DropTransactionById(ctx, transactionId).TransactionRbf(transactionRbf).Execute()
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
**transactionId** | **string** | The transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDropTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionRbf** | [**TransactionRbf**](TransactionRbf.md) | The request body to drop or to speed up transactions | 

### Return type

[**CreateTransferTransaction201Response**](CreateTransferTransaction201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EstimateFee

> EstimatedFee EstimateFee(ctx).EstimateFeeParams(estimateFeeParams).Execute()

Estimate transaction fee



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
	estimateFeeParams := coboWaas2.EstimateFeeParams{EstimateContractCallFeeParams: coboWaas2.NewEstimateContractCallFeeParams(coboWaas2.EstimateFeeRequestType("Transfer"), "ETH", coboWaas2.ContractCallSource{CustodialWeb3ContractCallSource: coboWaas2.NewCustodialWeb3ContractCallSource(coboWaas2.ContractCallSourceType("Web3"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")}, coboWaas2.ContractCallDestination{EvmContractCallDestination: coboWaas2.NewEvmContractCallDestination(coboWaas2.ContractCallDestinationType("EVM_Contract"), "0x0406db8351aa6839169bb363f63c2c808fee8f99", "0xa22cb4650000000000000000000000001e0049783f008a0085193e00003d00cd54003c71000000000000000000000000000000000000000000000000000000000000DEMO")})}

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
	resp, r, err := apiClient.TransactionsAPI.EstimateFee(ctx).EstimateFeeParams(estimateFeeParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.EstimateFee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EstimateFee`: EstimatedFee
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.EstimateFee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEstimateFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **estimateFeeParams** | [**EstimateFeeParams**](EstimateFeeParams.md) | The request body to estimate the transaction fee of a token transfer or a contract call. | 

### Return type

[**EstimatedFee**](EstimatedFee.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionApprovalDetail

> TransactionApprovalDetail GetTransactionApprovalDetail(ctx, transactionId).Execute()

Get transaction approval details



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
	resp, r, err := apiClient.TransactionsAPI.GetTransactionApprovalDetail(ctx, transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionApprovalDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionApprovalDetail`: TransactionApprovalDetail
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionApprovalDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**transactionId** | **string** | The transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionApprovalDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransactionApprovalDetail**](TransactionApprovalDetail.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionById

> TransactionDetail GetTransactionById(ctx, transactionId).Execute()

Get transaction information



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
	resp, r, err := apiClient.TransactionsAPI.GetTransactionById(ctx, transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionById`: TransactionDetail
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**transactionId** | **string** | The transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionByIdRequest struct via the builder pattern


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


## ListTransactionApprovalDetails

> ListTransactionApprovalDetails200Response ListTransactionApprovalDetails(ctx).TransactionIds(transactionIds).CoboIds(coboIds).Execute()

List transaction approval details



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
	coboIds := "20231213122855000000000000000000,20231213122955000000000000000000"

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
	resp, r, err := apiClient.TransactionsAPI.ListTransactionApprovalDetails(ctx).TransactionIds(transactionIds).CoboIds(coboIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListTransactionApprovalDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransactionApprovalDetails`: ListTransactionApprovalDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListTransactionApprovalDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionApprovalDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionIds** | **string** | A list of transaction IDs, separated by comma. | 
 **coboIds** | **string** | A list of Cobo IDs, separated by comma. A Cobo ID can be used to track a transaction. | 

### Return type

[**ListTransactionApprovalDetails200Response**](ListTransactionApprovalDetails200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> ListTransactions200Response ListTransactions(ctx).RequestId(requestId).CoboIds(coboIds).TransactionIds(transactionIds).TransactionHashes(transactionHashes).Types(types).Statuses(statuses).WalletIds(walletIds).ChainIds(chainIds).TokenIds(tokenIds).AssetIds(assetIds).VaultId(vaultId).ProjectId(projectId).MinCreatedTimestamp(minCreatedTimestamp).MaxCreatedTimestamp(maxCreatedTimestamp).Limit(limit).Before(before).After(after).Direction(direction).Execute()

List all transactions



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
	walletIds := "f47ac10b-58cc-4372-a567-0e02b2c3d479,1ddca562-8434-41c9-8809-d437bad9c868"
	chainIds := "BTC,ETH"
	tokenIds := "ETH_USDT,ETH_USDC"
	assetIds := "USDT,USDC"
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
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
	resp, r, err := apiClient.TransactionsAPI.ListTransactions(ctx).RequestId(requestId).CoboIds(coboIds).TransactionIds(transactionIds).TransactionHashes(transactionHashes).Types(types).Statuses(statuses).WalletIds(walletIds).ChainIds(chainIds).TokenIds(tokenIds).AssetIds(assetIds).VaultId(vaultId).ProjectId(projectId).MinCreatedTimestamp(minCreatedTimestamp).MaxCreatedTimestamp(maxCreatedTimestamp).Limit(limit).Before(before).After(after).Direction(direction).Execute()
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
 **requestId** | **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | 
 **coboIds** | **string** | A list of Cobo IDs, separated by comma. A Cobo ID can be used to track a transaction. | 
 **transactionIds** | **string** | A list of transaction IDs, separated by comma. | 
 **transactionHashes** | **string** | A list of transaction hashes, separated by comma. | 
 **types** | **string** | A list of transaction types, separated by comma. Possible values include:    - &#x60;Deposit&#x60;: A deposit transaction.   - &#x60;Withdrawal&#x60;: A withdrawal transaction.   - &#x60;ContractCall&#x60;: A transaction that interacts with a smart contract.   - &#x60;MessageSign&#x60;: A transaction that signs a message.    - &#x60;ExternalSafeTx&#x60;: A transaction to a Smart Contract Wallet (Safe{Wallet}) that requires one or multiple signatures to be executed.   - &#x60;Stake&#x60;: A transaction that creates a staking request.   - &#x60;Unstake&#x60;: A transaction that creates a unstaking request.  | 
 **statuses** | **string** | A list of transaction statuses, separated by comma. Possible values include:    - &#x60;Submitted&#x60;: The transaction is submitted.   - &#x60;PendingScreening&#x60;: The transaction is pending screening by Risk Control.    - &#x60;PendingAuthorization&#x60;: The transaction is pending approvals.   - &#x60;PendingSignature&#x60;: The transaction is pending signature.    - &#x60;Broadcasting&#x60;: The transaction is being broadcast.   - &#x60;Confirming&#x60;: The transaction is waiting for the required number of confirmations.   - &#x60;Completed&#x60;: The transaction is completed.   - &#x60;Failed&#x60;: The transaction failed.   - &#x60;Rejected&#x60;: The transaction is rejected.   - &#x60;Pending&#x60;: The transaction is waiting to be included in the next block of the blockchain.  | 
 **walletIds** | **string** | A list of wallet IDs, separated by comma. | 
 **chainIds** | **string** | A list of chain IDs, separated by comma. The chain ID is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | 
 **tokenIds** | **string** | A list of token IDs, separated by comma. The token ID is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
 **assetIds** | **string** | (This concept applies to Exchange Wallets only) A list of asset IDs, separated by comma. An asset ID is the unique identifier of the asset held within your linked exchange account. | 
 **vaultId** | **string** | The vault ID, which you can retrieve by calling [List all vaults](https://www.cobo.com/developers/v2/api-references/wallets--mpc-wallets/list-all-vaults). | 
 **projectId** | **string** | The project ID, which you can retrieve by calling [List all projects](https://www.cobo.com/developers/v2/api-references/wallets--mpc-wallets/list-all-projects).  | 
 **minCreatedTimestamp** | **int64** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. You can use this parameter to filter transactions created on or after the specified time. | 
 **maxCreatedTimestamp** | **int64** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. You can use this parameter to filter transactions created on or before the specified time. | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | This parameter specifies an object ID as a starting point for pagination, retrieving data before the specified object relative to the current dataset.    Suppose the current data is ordered as Object A, Object B, and Object C.  If you set &#x60;before&#x60; to the ID of Object C (&#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;), the response will include Object B and Object A.    **Notes**:   - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur. - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned. - If you set it to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | This parameter specifies an object ID as a starting point for pagination, retrieving data after the specified object relative to the current dataset.    Suppose the current data is ordered as Object A, Object B, and Object C. If you set &#x60;after&#x60; to the ID of Object A (&#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;), the response will include Object B and Object C.    **Notes**:   - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur. - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 
 **direction** | **string** | The sort direction. Possible values include:   - &#x60;ASC&#x60;: Sort the results in ascending order.   - &#x60;DESC&#x60;: Sort the results in descending order.  | [default to &quot;&quot;]

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


## ResendTransactionById

> CreateTransferTransaction201Response ResendTransactionById(ctx, transactionId).TransactionResend(transactionResend).Execute()

Resend transaction



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
	transactionResend := *coboWaas2.NewTransactionResend("f47ac10b-58cc-4372-a567-0e02b2c3d479")

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
	resp, r, err := apiClient.TransactionsAPI.ResendTransactionById(ctx, transactionId).TransactionResend(transactionResend).Execute()
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
**transactionId** | **string** | The transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionResend** | [**TransactionResend**](TransactionResend.md) | The request body to resend transactions | 

### Return type

[**CreateTransferTransaction201Response**](CreateTransferTransaction201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignAndBroadcastTransactionById

> CreateTransferTransaction201Response SignAndBroadcastTransactionById(ctx, transactionId).Execute()

Sign and broadcast transaction



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
	resp, r, err := apiClient.TransactionsAPI.SignAndBroadcastTransactionById(ctx, transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.SignAndBroadcastTransactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignAndBroadcastTransactionById`: CreateTransferTransaction201Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.SignAndBroadcastTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**transactionId** | **string** | The transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignAndBroadcastTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateTransferTransaction201Response**](CreateTransferTransaction201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpeedupTransactionById

> CreateTransferTransaction201Response SpeedupTransactionById(ctx, transactionId).TransactionRbf(transactionRbf).Execute()

Speed up transaction



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
	transactionRbf := *coboWaas2.NewTransactionRbf("f47ac10b-58cc-4372-a567-0e02b2c3d479", coboWaas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: coboWaas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", coboWaas2.FeeType("Fixed"), "ETH")})

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
	resp, r, err := apiClient.TransactionsAPI.SpeedupTransactionById(ctx, transactionId).TransactionRbf(transactionRbf).Execute()
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
**transactionId** | **string** | The transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpeedupTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionRbf** | [**TransactionRbf**](TransactionRbf.md) | The request body to drop or to speed up transactions | 

### Return type

[**CreateTransferTransaction201Response**](CreateTransferTransaction201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

