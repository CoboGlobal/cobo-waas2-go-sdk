# \TransactionsAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTransactionById**](TransactionsAPI.md#CancelTransactionById) | **Post** /transactions/{transaction_id}/cancel | Cancel transaction
[**CreateContractCallTransaction**](TransactionsAPI.md#CreateContractCallTransaction) | **Post** /transactions/contract_call | Call smart contract
[**CreateMessageSignTransaction**](TransactionsAPI.md#CreateMessageSignTransaction) | **Post** /transactions/message_sign | Sign message
[**CreateTransferTransaction**](TransactionsAPI.md#CreateTransferTransaction) | **Post** /transactions/transfer | Transfer token
[**DropTransactionById**](TransactionsAPI.md#DropTransactionById) | **Post** /transactions/{transaction_id}/drop | Drop transaction
[**EstimateFee**](TransactionsAPI.md#EstimateFee) | **Post** /transactions/estimate_fee | Estimate transaction fee
[**GetTransactionById**](TransactionsAPI.md#GetTransactionById) | **Get** /transactions/{transaction_id} | Get transaction information
[**ListTransactions**](TransactionsAPI.md#ListTransactions) | **Get** /transactions | List all transactions
[**ResendTransactionById**](TransactionsAPI.md#ResendTransactionById) | **Post** /transactions/{transaction_id}/resend | Resend transaction
[**SpeedupTransactionById**](TransactionsAPI.md#SpeedupTransactionById) | **Post** /transactions/{transaction_id}/speedup | Speed up transaction



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
	waas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The transaction ID.

	configuration := waas2.NewConfiguration()
	apiClient := waas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
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

[CoboAuth](../README.md#CoboAuth)

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
	waas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	contractCallParams := *waas2.NewContractCallParams("f47ac10b-58cc-4372-a567-0e02b2c3d479", "ETH", waas2.ContractCallSource{MpcContractCallSource: waas2.NewMpcContractCallSource(waas2.ContractCallSourceType("Org-Controlled"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")}, waas2.ContractCallDestination{EvmContractCallDestination: waas2.NewEvmContractCallDestination(waas2.ContractCallDestinationType("EVM_Contract"), "0x0406db8351aa6839169bb363f63c2c808fee8f99", string([B@499c4d61))}) // ContractCallParams | The request body for making a contract call. (optional)

	configuration := waas2.NewConfiguration()
	apiClient := waas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
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
	waas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	messageSignParams := *waas2.NewMessageSignParams("f47ac10b-58cc-4372-a567-0e02b2c3d479", "ETH", waas2.MessageSignSource{MpcMessageSignSource: waas2.NewMpcMessageSignSource(waas2.MessageSignSourceType("Org-Controlled"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "19AR6YWEGbSoY8UT9Ksy9WrmrZPD5sL4Ku")}, waas2.MessageSignDestination{EvmEIP191MessageSignDestination: waas2.NewEvmEIP191MessageSignDestination(waas2.MessageSignDestinationType("EVM_EIP_191_Signature"), "YWFhYQ==")}) // MessageSignParams | The request body to create a message sign transaction (optional)

	configuration := waas2.NewConfiguration()
	apiClient := waas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
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
 **messageSignParams** | [**MessageSignParams**](MessageSignParams.md) | The request body to create a message sign transaction | 

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
	waas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transferParams := *waas2.NewTransferParams("f47ac10b-58cc-4372-a567-0e02b2c3d479", waas2.TransferSource{CustodialTransferSource: waas2.NewCustodialTransferSource(waas2.WalletSubtype("Asset"), "f47ac10b-58cc-4372-a567-0e02b2c3d479")}, "ETH_USDT", waas2.TransferDestination{AddressTransferDestination: waas2.NewAddressTransferDestination(waas2.TransferDestinationType("Address"))}) // TransferParams | The request body to create a transfer transaction (optional)

	configuration := waas2.NewConfiguration()
	apiClient := waas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
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
	waas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The transaction ID.
	transactionRbf := *waas2.NewTransactionRbf("f47ac10b-58cc-4372-a567-0e02b2c3d479", waas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: waas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", waas2.FeeType("Fixed"), "ETH")}) // TransactionRbf | The request body to drop or to speed up transactions (optional)

	configuration := waas2.NewConfiguration()
	apiClient := waas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
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

[CoboAuth](../README.md#CoboAuth)

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
	waas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	estimateFeeParams := waas2.EstimateFeeParams{EstimateContractCallFeeParams: waas2.NewEstimateContractCallFeeParams("f47ac10b-58cc-4372-a567-0e02b2c3d479", waas2.EstimateFeeRequestType("Transfer"), "ETH", waas2.ContractCallSource{MpcContractCallSource: waas2.NewMpcContractCallSource(waas2.ContractCallSourceType("Org-Controlled"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")}, waas2.ContractCallDestination{EvmContractCallDestination: waas2.NewEvmContractCallDestination(waas2.ContractCallDestinationType("EVM_Contract"), "0x0406db8351aa6839169bb363f63c2c808fee8f99", string([B@499c4d61))})} // EstimateFeeParams | The request body to estimate the transaction fee of a token transfer or a contract call. (optional)

	configuration := waas2.NewConfiguration()
	apiClient := waas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
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

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	waas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The transaction ID.

	configuration := waas2.NewConfiguration()
	apiClient := waas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
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

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> ListTransactions200Response ListTransactions(ctx).RequestId(requestId).CoboIds(coboIds).TransactionIds(transactionIds).TransactionHashes(transactionHashes).Types(types).Statuses(statuses).WalletIds(walletIds).ChainIds(chainIds).TokenIds(tokenIds).AssetIds(assetIds).VaultId(vaultId).ProjectId(projectId).MinCreatedTimestamp(minCreatedTimestamp).MaxCreatedTimestamp(maxCreatedTimestamp).Limit(limit).Before(before).After(after).Execute()

List all transactions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	waas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	requestId := "web_send_by_user_327_1610444045047" // string | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. (optional)
	coboIds := "20231213122855000000000000000000,20231213122955000000000000000000" // string | A list of Cobo IDs, separated by comma. A Cobo ID can be used to track a transaction. (optional)
	transactionIds := "f47ac10b-58cc-4372-a567-0e02b2c3d479,557918d2-632a-4fe1-932f-315711f05fe3" // string | A list of transaction IDs, separated by comma. (optional)
	transactionHashes := "239861be9a4afe080c359b7fe4a1d035945ec46256b1a0f44d1267c71de8ec28" // string | A list of transaction hashes, separated by comma. (optional)
	types := "Deposit,Withdrawal" // string | A list of transaction types, separated by comma. Possible values include:    - `Deposit`: A deposit transaction.   - `Withdrawal`: A withdrawal transaction.   - `ContractCall`: A transaction that interacts with a smart contract.   - `MessageSign`: A transaction that signs a message.    - `ExternalSafeTx`: A transaction to a Smart Contract Wallet (Safe{Wallet}) that requires one or multiple signatures to be executed.  (optional)
	statuses := "Completed,Failed" // string | A list of transaction statuses, separated by comma. Possible values include:    - `Submitted`: The transaction is submitted.   - `PendingScreening`: The transaction is pending screening by Risk Control.    - `PendingAuthorization`: The transaction is pending approvals.   - `PendingSignature`: The transaction is pending signature.    - `Broadcasting`: The transaction is being broadcast.   - `Confirming`: The transaction is waiting for the required number of confirmations.   - `Completed`: The transaction is completed.   - `Failed`: The transaction failed.   - `Rejected`: The transaction is rejected.   - `Pending`: The transaction is pending.  (optional)
	walletIds := "f47ac10b-58cc-4372-a567-0e02b2c3d479,1ddca562-8434-41c9-8809-d437bad9c868" // string | A list of wallet IDs, separated by comma. (optional)
	chainIds := "BTC,ETH" // string | A list of chain IDs, separated by comma. The chain ID is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](/v2/api-references/wallets/list-enabled-chains). (optional)
	tokenIds := "ETH_USDT,ETH_USDC" // string | A list of token IDs, separated by comma. The token ID is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](/v2/api-references/wallets/list-enabled-tokens). (optional)
	assetIds := "USDT,USDC" // string | (This concept applies to Exchange Wallets only) A list of asset IDs, separated by comma. An asset ID is the unique identifier of the asset held within your linked exchange account. (optional)
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | (This parameter is applicable to MPC Wallets only) The vault ID, which you can retrieve by calling [List all vaults](/v2/api-references/wallets--mpc-wallets/list-all-vaults). (optional)
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The project ID, which you can retrieve by calling [List all projects](/v2/api-references/wallets--mpc-wallets/list-all-projects).  (optional)
	minCreatedTimestamp := int64(1635744000000) // int64 | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. You can use this parameter to filter transactions created on or after the specified time. (optional)
	maxCreatedTimestamp := int64(1635744000000) // int64 | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. You can use this parameter to filter transactions created on or before the specified time. (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. For most operations, the value range is [1, 50]. (optional) (default to 10)
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1" // string | An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify `before` as `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1`, the request will retrieve a list of data objects that end before the object with the object ID `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  - If you set both `after` and `before`, an error will occur.  - If you leave both `before` and `after` empty, the first page of data is returned.  - If you set `before` to `infinity`, the last page of data is returned.  (optional)
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk" // string | An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify `after` as `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk`, the request will retrieve a list of data objects that start after the object with the object ID `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  - If you set both `after` and `before`, an error will occur.  - If you leave both `before` and `after` empty, the first page of data is returned.  (optional)

	configuration := waas2.NewConfiguration()
	apiClient := waas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.ListTransactions(ctx).RequestId(requestId).CoboIds(coboIds).TransactionIds(transactionIds).TransactionHashes(transactionHashes).Types(types).Statuses(statuses).WalletIds(walletIds).ChainIds(chainIds).TokenIds(tokenIds).AssetIds(assetIds).VaultId(vaultId).ProjectId(projectId).MinCreatedTimestamp(minCreatedTimestamp).MaxCreatedTimestamp(maxCreatedTimestamp).Limit(limit).Before(before).After(after).Execute()
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
 **types** | **string** | A list of transaction types, separated by comma. Possible values include:    - &#x60;Deposit&#x60;: A deposit transaction.   - &#x60;Withdrawal&#x60;: A withdrawal transaction.   - &#x60;ContractCall&#x60;: A transaction that interacts with a smart contract.   - &#x60;MessageSign&#x60;: A transaction that signs a message.    - &#x60;ExternalSafeTx&#x60;: A transaction to a Smart Contract Wallet (Safe{Wallet}) that requires one or multiple signatures to be executed.  | 
 **statuses** | **string** | A list of transaction statuses, separated by comma. Possible values include:    - &#x60;Submitted&#x60;: The transaction is submitted.   - &#x60;PendingScreening&#x60;: The transaction is pending screening by Risk Control.    - &#x60;PendingAuthorization&#x60;: The transaction is pending approvals.   - &#x60;PendingSignature&#x60;: The transaction is pending signature.    - &#x60;Broadcasting&#x60;: The transaction is being broadcast.   - &#x60;Confirming&#x60;: The transaction is waiting for the required number of confirmations.   - &#x60;Completed&#x60;: The transaction is completed.   - &#x60;Failed&#x60;: The transaction failed.   - &#x60;Rejected&#x60;: The transaction is rejected.   - &#x60;Pending&#x60;: The transaction is pending.  | 
 **walletIds** | **string** | A list of wallet IDs, separated by comma. | 
 **chainIds** | **string** | A list of chain IDs, separated by comma. The chain ID is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](/v2/api-references/wallets/list-enabled-chains). | 
 **tokenIds** | **string** | A list of token IDs, separated by comma. The token ID is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](/v2/api-references/wallets/list-enabled-tokens). | 
 **assetIds** | **string** | (This concept applies to Exchange Wallets only) A list of asset IDs, separated by comma. An asset ID is the unique identifier of the asset held within your linked exchange account. | 
 **vaultId** | **string** | (This parameter is applicable to MPC Wallets only) The vault ID, which you can retrieve by calling [List all vaults](/v2/api-references/wallets--mpc-wallets/list-all-vaults). | 
 **projectId** | **string** | The project ID, which you can retrieve by calling [List all projects](/v2/api-references/wallets--mpc-wallets/list-all-projects).  | 
 **minCreatedTimestamp** | **int64** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. You can use this parameter to filter transactions created on or after the specified time. | 
 **maxCreatedTimestamp** | **int64** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. You can use this parameter to filter transactions created on or before the specified time. | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify &#x60;before&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  - If you set &#x60;before&#x60; to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify &#x60;after&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 

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

> CreateTransferTransaction201Response ResendTransactionById(ctx, transactionId).TransactionResend(transactionResend).Execute()

Resend transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	waas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The transaction ID.
	transactionResend := *waas2.NewTransactionResend("f47ac10b-58cc-4372-a567-0e02b2c3d479") // TransactionResend | The request body to resend transactions (optional)

	configuration := waas2.NewConfiguration()
	apiClient := waas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
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

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	waas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The transaction ID.
	transactionRbf := *waas2.NewTransactionRbf("f47ac10b-58cc-4372-a567-0e02b2c3d479", waas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: waas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", waas2.FeeType("Fixed"), "ETH")}) // TransactionRbf | The request body to drop or to speed up transactions (optional)

	configuration := waas2.NewConfiguration()
	apiClient := waas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
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

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

