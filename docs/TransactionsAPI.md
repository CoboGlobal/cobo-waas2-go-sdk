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
[**ListFeeRates**](TransactionsAPI.md#ListFeeRates) | **Get** /transactions/fee_rates | Get fee rates
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
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The transaction ID.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
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

> CreateTransferTransaction201Response CreateContractCallTransaction(ctx).ContractCall(contractCall).Execute()

Call smart contract



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
	contractCall := *CoboWaas2.NewContractCall("f47ac10b-58cc-4372-a567-0e02b2c3d479", "ETH", CoboWaas2.ContractCallSource{MpcContractCallSource: CoboWaas2.NewMpcContractCallSource("Org-Controlled", "f47ac10b-58cc-4372-a567-0e02b2c3d479", "19AR6YWEGbSoY8UT9Ksy9WrmrZPD5sL4Ku", *CoboWaas2.NewMpcSigningGroup())}, *CoboWaas2.NewEstimateFeeContractCallDestination()) // ContractCall | The request body for making a contract call. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.CreateContractCallTransaction(ctx).ContractCall(contractCall).Execute()
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
 **contractCall** | [**ContractCall**](ContractCall.md) | The request body for making a contract call. | 

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

> CreateTransferTransaction201Response CreateMessageSignTransaction(ctx).SignMessage(signMessage).Execute()

Sign message



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
	signMessage := *CoboWaas2.NewSignMessage("f47ac10b-58cc-4372-a567-0e02b2c3d479", "ETH", *CoboWaas2.NewSignMessageSource(), *CoboWaas2.NewSignMessageDestination()) // SignMessage | The request body to create a message sign transaction (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.CreateMessageSignTransaction(ctx).SignMessage(signMessage).Execute()
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
 **signMessage** | [**SignMessage**](SignMessage.md) | The request body to create a message sign transaction | 

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

> CreateTransferTransaction201Response CreateTransferTransaction(ctx).Transfer(transfer).Execute()

Transfer token



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
	transfer := *CoboWaas2.NewTransfer("f47ac10b-58cc-4372-a567-0e02b2c3d479", CoboWaas2.TransferSource{BaseTransferSource: CoboWaas2.NewBaseTransferSource(CoboWaas2.WalletSubtype("Asset"), "f47ac10b-58cc-4372-a567-0e02b2c3d479")}, "ETH_USDT", CoboWaas2.TransferDestination{AddressTransferDestination: CoboWaas2.NewAddressTransferDestination(CoboWaas2.TransferDestinationType("Address"))}) // Transfer | The request body to create a transfer transaction (optional)

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
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The transaction ID.
	transactionRbf := *CoboWaas2.NewTransactionRbf("f47ac10b-58cc-4372-a567-0e02b2c3d479") // TransactionRbf | The request body to drop or to speed up transactions (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
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

> []EstimationFee EstimateFee(ctx).EstimateFee(estimateFee).Execute()

Estimate transaction fee



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
	estimateFee := CoboWaas2.EstimateFee{EstimateFeeContractCall: CoboWaas2.NewEstimateFeeContractCall("f47ac10b-58cc-4372-a567-0e02b2c3d479", "Transfer", "ETH", CoboWaas2.ContractCallSource{MpcContractCallSource: CoboWaas2.NewMpcContractCallSource("Org-Controlled", "f47ac10b-58cc-4372-a567-0e02b2c3d479", "19AR6YWEGbSoY8UT9Ksy9WrmrZPD5sL4Ku", *CoboWaas2.NewMpcSigningGroup())}, *CoboWaas2.NewEstimateFeeContractCallDestination())} // EstimateFee | The request body to estimate the transaction fee of a token transfer or a contract call. (optional)

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
	// response from `EstimateFee`: []EstimationFee
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.EstimateFee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEstimateFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **estimateFee** | [**EstimateFee**](EstimateFee.md) | The request body to estimate the transaction fee of a token transfer or a contract call. | 

### Return type

[**[]EstimationFee**](EstimationFee.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionById

> Transaction GetTransactionById(ctx, transactionId).Execute()

Get transaction information



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
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The transaction ID.

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
**transactionId** | **string** | The transaction ID. | 

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


## ListFeeRates

> []FeeRate ListFeeRates(ctx).ChainId(chainId).Execute()

Get fee rates



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
	chainId := "ETH" // string | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains).

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.ListFeeRates(ctx).ChainId(chainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListFeeRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFeeRates`: []FeeRate
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListFeeRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFeeRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). | 

### Return type

[**[]FeeRate**](FeeRate.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> ListTransactions200Response ListTransactions(ctx).RequestId(requestId).CoboId(coboId).TransactionId(transactionId).TransactionHash(transactionHash).Type_(type_).Status(status).SourceType(sourceType).SourceWalletId(sourceWalletId).SourceAddress(sourceAddress).DestinationType(destinationType).DestinationWalletId(destinationWalletId).DestinationAddress(destinationAddress).ChainIds(chainIds).TokenIds(tokenIds).AssetIds(assetIds).VaultId(vaultId).ProjectId(projectId).MinCreatedTimestamp(minCreatedTimestamp).MaxCreatedTimestamp(maxCreatedTimestamp).SortBy(sortBy).Direction(direction).Limit(limit).Before(before).After(after).Execute()

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
	requestId := "web_send_by_user_327_1610444045047" // string | The request ID that is used to track a withdrawal request. The request ID is provided by you and must be unique within your organization. (optional)
	coboId := "20231213122855000000000000000000" // string | The Cobo ID, which can be used to track a transaction. (optional)
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The transaction ID. (optional)
	transactionHash := "239861be9a4afe080c359b7fe4a1d035945ec46256b1a0f44d1267c71de8ec28" // string | The transaction hash. (optional)
	type_ := []CoboWaas2.TransactionType{CoboWaas2.TransactionType("Deposit")} // []TransactionType | The transaction type. Possible values include:    - `Deposit`: A deposit transaction.   - `Withdrawal`: A withdrawal transaction.   - `TokenApproval`: A transaction that grants permission to access your tokens.    - `ContractCall`: A transaction that interacts with a smart contract.   - `TransactionFeePayment`: A transaction that is initiated by Fee Station to pay your transaction fee.   - `RawMessage`: A transaction that signs a message.  (optional)
	status := []CoboWaas2.TransactionStatus{CoboWaas2.TransactionStatus("Submitted")} // []TransactionStatus | The transaction status. Possible values include:    - `Submitted`: The transaction is submitted.   - `PendingScreening`: The transaction is pending screening by Risk Control.    - `PendingAuthorization`: The transaction is pending approvals.   - `PendingSignature`: The transaction is pending signature.    - `Broadcasting`: The transaction is being broadcast.   - `Confirming`: The transaction is waiting for the required number of confirmations.   - `Completed`: The transaction is completed.   - `Failed`: The transaction failed.   - `Rejected`: The transaction is rejected.  (optional)
	sourceType := []CoboWaas2.TransactionSourceType{CoboWaas2.TransactionSourceType("Address")} // []TransactionSourceType | The type of transaction source. Possible values include:   - `Address`: An external address.   - `CustodialWallet`: A Custodial Wallet.   - `MPCWallet`: An MPC Wallet.   - `SafeWallet`: A Smart Contract Wallet (Safe{Wallet}).   - `ExchangeWallet`: An Exchange Wallet.   - `FeeStation`: A Fee Station.  (optional)
	sourceWalletId := []string{"Inner_example"} // []string | The wallet ID of the transaction source. (optional)
	sourceAddress := []string{"Inner_example"} // []string | The address of the transaction source. (optional)
	destinationType := []CoboWaas2.TransactionDestinationType{CoboWaas2.TransactionDestinationType("Address")} // []TransactionDestinationType | The transaction destination type. Possible values include:   - `Address`: An external address.    - `ContractCall`: A transaction that interacts with a smart contract.   - `MessageSign`: A transaction that signs a message.    - `CustodialWallet`: A Custodial Wallet.   - `MPCWallet`: An MPC Wallet.   - `SafeWallet`: A Smart Contract Wallets (Safe{Wallet}).   - `ExchangeWallet`: An Exchange Wallet.  (optional)
	destinationWalletId := []string{"Inner_example"} // []string | The wallet ID of the transaction destination. (optional)
	destinationAddress := []string{"Inner_example"} // []string | The address of the transaction destination. (optional)
	chainIds := "BTC,ETH" // string | A list of chain IDs, separated by comma. The chain ID is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). (optional)
	tokenIds := "ETH_USDT,ETH_USDC" // string | A list of token IDs, separated by comma. The token ID is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). (optional)
	assetIds := "USDT,USDC" // string | A list of asset IDs, separated by comma. (This concept applies to Exchange Wallets only) An asset is a digital representation of a valuable resource on a blockchain network. Exchange Wallets group your holdings by asset, even if the same asset exists on different blockchains. For example, if your Exchange Wallet has 1 USDT on Ethereum and 1 USDT on TRON, then your asset balance is 2 USDT. (optional)
	vaultId := []string{"Inner_example"} // []string | The MPC vault ID. (optional)
	projectId := []string{"Inner_example"} // []string | The MPC project ID. (optional)
	minCreatedTimestamp := int32(1635744000) // int32 | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. You can use this parameter to filter transactions created on or after the specified time. (optional)
	maxCreatedTimestamp := int32(1635744000) // int32 | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. You can use this parameter to filter transactions created on or before the specified time. (optional)
	sortBy := "timestamp" // string | The field used for sorting. (optional) (default to "")
	direction := "ASC" // string | The sort direction. Possible values include:   - `ASC`: Sort the results in ascending order.   - `DESC`: Sort the results in descending order.  (optional) (default to "")
	limit := int32(10) // int32 | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that end before the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` and `after` are both set to empty, the first slice is returned.  (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that start after the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` is set to empty and `after` is set to `last`, the last slice is returned.  (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.TransactionsAPI.ListTransactions(ctx).RequestId(requestId).CoboId(coboId).TransactionId(transactionId).TransactionHash(transactionHash).Type_(type_).Status(status).SourceType(sourceType).SourceWalletId(sourceWalletId).SourceAddress(sourceAddress).DestinationType(destinationType).DestinationWalletId(destinationWalletId).DestinationAddress(destinationAddress).ChainIds(chainIds).TokenIds(tokenIds).AssetIds(assetIds).VaultId(vaultId).ProjectId(projectId).MinCreatedTimestamp(minCreatedTimestamp).MaxCreatedTimestamp(maxCreatedTimestamp).SortBy(sortBy).Direction(direction).Limit(limit).Before(before).After(after).Execute()
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
 **requestId** | **string** | The request ID that is used to track a withdrawal request. The request ID is provided by you and must be unique within your organization. | 
 **coboId** | **string** | The Cobo ID, which can be used to track a transaction. | 
 **transactionId** | **string** | The transaction ID. | 
 **transactionHash** | **string** | The transaction hash. | 
 **type_** | [**[]TransactionType**](TransactionType.md) | The transaction type. Possible values include:    - &#x60;Deposit&#x60;: A deposit transaction.   - &#x60;Withdrawal&#x60;: A withdrawal transaction.   - &#x60;TokenApproval&#x60;: A transaction that grants permission to access your tokens.    - &#x60;ContractCall&#x60;: A transaction that interacts with a smart contract.   - &#x60;TransactionFeePayment&#x60;: A transaction that is initiated by Fee Station to pay your transaction fee.   - &#x60;RawMessage&#x60;: A transaction that signs a message.  | 
 **status** | [**[]TransactionStatus**](TransactionStatus.md) | The transaction status. Possible values include:    - &#x60;Submitted&#x60;: The transaction is submitted.   - &#x60;PendingScreening&#x60;: The transaction is pending screening by Risk Control.    - &#x60;PendingAuthorization&#x60;: The transaction is pending approvals.   - &#x60;PendingSignature&#x60;: The transaction is pending signature.    - &#x60;Broadcasting&#x60;: The transaction is being broadcast.   - &#x60;Confirming&#x60;: The transaction is waiting for the required number of confirmations.   - &#x60;Completed&#x60;: The transaction is completed.   - &#x60;Failed&#x60;: The transaction failed.   - &#x60;Rejected&#x60;: The transaction is rejected.  | 
 **sourceType** | [**[]TransactionSourceType**](TransactionSourceType.md) | The type of transaction source. Possible values include:   - &#x60;Address&#x60;: An external address.   - &#x60;CustodialWallet&#x60;: A Custodial Wallet.   - &#x60;MPCWallet&#x60;: An MPC Wallet.   - &#x60;SafeWallet&#x60;: A Smart Contract Wallet (Safe{Wallet}).   - &#x60;ExchangeWallet&#x60;: An Exchange Wallet.   - &#x60;FeeStation&#x60;: A Fee Station.  | 
 **sourceWalletId** | **[]string** | The wallet ID of the transaction source. | 
 **sourceAddress** | **[]string** | The address of the transaction source. | 
 **destinationType** | [**[]TransactionDestinationType**](TransactionDestinationType.md) | The transaction destination type. Possible values include:   - &#x60;Address&#x60;: An external address.    - &#x60;ContractCall&#x60;: A transaction that interacts with a smart contract.   - &#x60;MessageSign&#x60;: A transaction that signs a message.    - &#x60;CustodialWallet&#x60;: A Custodial Wallet.   - &#x60;MPCWallet&#x60;: An MPC Wallet.   - &#x60;SafeWallet&#x60;: A Smart Contract Wallets (Safe{Wallet}).   - &#x60;ExchangeWallet&#x60;: An Exchange Wallet.  | 
 **destinationWalletId** | **[]string** | The wallet ID of the transaction destination. | 
 **destinationAddress** | **[]string** | The address of the transaction destination. | 
 **chainIds** | **string** | A list of chain IDs, separated by comma. The chain ID is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). | 
 **tokenIds** | **string** | A list of token IDs, separated by comma. The token ID is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
 **assetIds** | **string** | A list of asset IDs, separated by comma. (This concept applies to Exchange Wallets only) An asset is a digital representation of a valuable resource on a blockchain network. Exchange Wallets group your holdings by asset, even if the same asset exists on different blockchains. For example, if your Exchange Wallet has 1 USDT on Ethereum and 1 USDT on TRON, then your asset balance is 2 USDT. | 
 **vaultId** | **[]string** | The MPC vault ID. | 
 **projectId** | **[]string** | The MPC project ID. | 
 **minCreatedTimestamp** | **int32** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. You can use this parameter to filter transactions created on or after the specified time. | 
 **maxCreatedTimestamp** | **int32** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. You can use this parameter to filter transactions created on or before the specified time. | 
 **sortBy** | **string** | The field used for sorting. | [default to &quot;&quot;]
 **direction** | **string** | The sort direction. Possible values include:   - &#x60;ASC&#x60;: Sort the results in ascending order.   - &#x60;DESC&#x60;: Sort the results in descending order.  | [default to &quot;&quot;]
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

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
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The transaction ID.
	transactionResend := *CoboWaas2.NewTransactionResend("f47ac10b-58cc-4372-a567-0e02b2c3d479") // TransactionResend | The request body to resend transactions (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
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
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	transactionId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The transaction ID.
	transactionRbf := *CoboWaas2.NewTransactionRbf("f47ac10b-58cc-4372-a567-0e02b2c3d479") // TransactionRbf | The request body to drop or to speed up transactions (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
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

