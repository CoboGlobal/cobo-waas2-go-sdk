# Go API client for CoboWaas2

Cobo WaaS 2.0 enables you to programmatically access Cobo's full suite of
crypto wallet technologies with powerful and flexible access controls.

# Wallet technologies
- Custodial Wallet
- MPC Wallet
- Smart Contract Wallet (Based on Safe{Wallet})
- Exchange Wallet

# Risk Control technologies
- Workflow
- Access Control List (ACL)

# Risk Control targets
- Wallet Management
  - User/team and their permission management
  - Risk control configurations, e.g. whitelist, blacklist, rate-limiting etc.
- Blockchain Interaction
  - Crypto transfer
  - Smart Contract Invocation

# Important
HTTPS only. RESTful, resource oriented

# Get Started
Set up your APIs or get authorization

# Authentication and Authorization
CoboAuth

# Request and Response
application/json

# Error Handling

### Common error codes
| Error Code | Description |
| -- | -- |

### API-specific error codes
For error codes that are dedicated to a specific API, see the Error codes section in each API specification, for example, /v3/wallets.

# Rate and Usage Limiting

# Idempotent Request

# Pagination
# Support
[Developer Hub](https://cobo.com/developers)


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v2
- Package version: 0.1.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.cobo.com/waas](https://www.cobo.com/waas)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using Dev environment set context value `CoboWaas2.ContextEnv` of type `int`.

```go
ctx := context.WithValue(context.Background(), CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
```

For using Dev environment set context value `CoboWaas2.ContextEnv` of type `int`.

```go
ctx := context.WithValue(context.Background(), CoboWaas2.ContextEnv, CoboWaas2.ProdEnv)
```

you can also set server URL using context values `CoboWaas2.ContextServerHost` of type `string`

```go
ctx := context.WithValue(context.Background(), CoboWaas2.ContextServerHost, "https://api[.xxxx].cobo.com/v2")
```

## Set Api Signer
```go
ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
    Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.cobo.com/v3*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DevelopersWebhooksAPI* | [**GetWebhookEvent**](docs/DevelopersWebhooksAPI.md#getwebhookevent) | **Get** /webhooks/events/{event_id} | Retrieve webhook event information by event ID.
*DevelopersWebhooksAPI* | [**GetWebhookEventLogs**](docs/DevelopersWebhooksAPI.md#getwebhookeventlogs) | **Get** /webhooks/events/{event_id}/logs | List webhook event logs by event ID.
*DevelopersWebhooksAPI* | [**ListEvents**](docs/DevelopersWebhooksAPI.md#listevents) | **Get** /webhooks/events | List triggered events.
*DevelopersWebhooksAPI* | [**ListWebhookEventDefinitions**](docs/DevelopersWebhooksAPI.md#listwebhookeventdefinitions) | **Get** /webhooks/events/definitions | List all supported event definitions.
*DevelopersWebhooksAPI* | [**RetryWebhookEvent**](docs/DevelopersWebhooksAPI.md#retrywebhookevent) | **Post** /webhooks/events/{event_id}/retry | Retry webhook event by event ID.
*TransactionsAPI* | [**CreateSignMessageTransaction**](docs/TransactionsAPI.md#createsignmessagetransaction) | **Post** /transactions/sign | Create a sign message transaction
*TransactionsAPI* | [**CreateSmartContractCallTransaction**](docs/TransactionsAPI.md#createsmartcontractcalltransaction) | **Post** /transactions/call | Create a smart contract call transaction
*TransactionsAPI* | [**CreateTransferTransaction**](docs/TransactionsAPI.md#createtransfertransaction) | **Post** /transactions/transfer | Create a transfer transaction
*TransactionsAPI* | [**DropTransactionById**](docs/TransactionsAPI.md#droptransactionbyid) | **Post** /transactions/{transaction_id}/drop | Drop a transaction by ID
*TransactionsAPI* | [**EstimateFee**](docs/TransactionsAPI.md#estimatefee) | **Post** /transactions/estimate_fee | Estimate the fee for transaction
*TransactionsAPI* | [**GetChainFeePrice**](docs/TransactionsAPI.md#getchainfeeprice) | **Get** /transactions/fee_price | Get the fee price data for chain and/or token(Hold, TBD after normalize fee settings)
*TransactionsAPI* | [**GetTransactionById**](docs/TransactionsAPI.md#gettransactionbyid) | **Get** /transactions/{transaction_id} | Get transaction information by ID
*TransactionsAPI* | [**ListTransactions**](docs/TransactionsAPI.md#listtransactions) | **Get** /transactions | List all transactions
*TransactionsAPI* | [**ResendTransactionById**](docs/TransactionsAPI.md#resendtransactionbyid) | **Post** /transactions/{transaction_id}/resend | Resend a transaction by ID
*TransactionsAPI* | [**RetryTransactionDoubleCheckById**](docs/TransactionsAPI.md#retrytransactiondoublecheckbyid) | **Post** /transactions/{transaction_id}/double_check/retry | Retry up a transaction double-check by ID
*TransactionsAPI* | [**SpeedupTransactionById**](docs/TransactionsAPI.md#speeduptransactionbyid) | **Post** /transactions/{transaction_id}/speedup | Speed up a transaction by ID
*TransactionsAPI* | [**UpdateTransacitonById**](docs/TransactionsAPI.md#updatetransacitonbyid) | **Put** /transactions/{transaction_id} | Update transaction by ID
*WalletsAPI* | [**AddWalletAddress**](docs/WalletsAPI.md#addwalletaddress) | **Post** /wallets/{wallet_id}/addresses | Add address to a wallet
*WalletsAPI* | [**CreateWallet**](docs/WalletsAPI.md#createwallet) | **Post** /wallets | Create new wallet
*WalletsAPI* | [**DeleteWalletById**](docs/WalletsAPI.md#deletewalletbyid) | **Delete** /wallets/{wallet_id} | Delete a wallet by ID
*WalletsAPI* | [**GetAddressValidity**](docs/WalletsAPI.md#getaddressvalidity) | **Get** /wallets/address/validity | Get the given address validity for token
*WalletsAPI* | [**GetAssets**](docs/WalletsAPI.md#getassets) | **Get** /wallets/assets | List the metadata of assets
*WalletsAPI* | [**GetChains**](docs/WalletsAPI.md#getchains) | **Get** /wallets/chains | List the metadata of chain
*WalletsAPI* | [**GetMaxSendValue**](docs/WalletsAPI.md#getmaxsendvalue) | **Get** /wallets/{wallet_id}/max_sendable_value | Get max sendable Vaule
*WalletsAPI* | [**GetSpendableList**](docs/WalletsAPI.md#getspendablelist) | **Get** /wallets/{wallet_id}/spendables | List the spendable utxo
*WalletsAPI* | [**GetSupportedChains**](docs/WalletsAPI.md#getsupportedchains) | **Get** /wallets/supported_chains | List the supported chains by wallet subtype
*WalletsAPI* | [**GetSupportedTokens**](docs/WalletsAPI.md#getsupportedtokens) | **Get** /wallets/supported_tokens | List the supported tokens by wallet subtype and chain id if specified
*WalletsAPI* | [**GetTokens**](docs/WalletsAPI.md#gettokens) | **Get** /wallets/tokens | List the metadata of tokens
*WalletsAPI* | [**GetWalletAddressById**](docs/WalletsAPI.md#getwalletaddressbyid) | **Get** /wallets/{wallet_id}/addresses/{address_id} | Get address information by ID
*WalletsAPI* | [**GetWalletAddressTokenBalances**](docs/WalletsAPI.md#getwalletaddresstokenbalances) | **Get** /wallets/{wallet_id}/addresses/{address_id}/tokens | List the token balance by address in the wallets(to be specific)
*WalletsAPI* | [**GetWalletById**](docs/WalletsAPI.md#getwalletbyid) | **Get** /wallets/{wallet_id} | Get wallet information by ID
*WalletsAPI* | [**GetWalletTokenBalances**](docs/WalletsAPI.md#getwallettokenbalances) | **Get** /wallets/{wallet_id}/tokens | List the token balance in the wallets(to be specific)
*WalletsAPI* | [**ListAddresses**](docs/WalletsAPI.md#listaddresses) | **Get** /wallets/{wallet_id}/addresses | List wallet addresses by wallet ID
*WalletsAPI* | [**ListWallets**](docs/WalletsAPI.md#listwallets) | **Get** /wallets | List all wallets
*WalletsAPI* | [**UpdateWalletById**](docs/WalletsAPI.md#updatewalletbyid) | **Put** /wallets/{wallet_id} | Update wallet by ID
*WalletsExchangeWalletAPI* | [**GetExchangeSupportedAssets**](docs/WalletsExchangeWalletAPI.md#getexchangesupportedassets) | **Get** /wallets/exchanges/{exchange_id}/supported_assets | List the supported assets by exchange id
*WalletsExchangeWalletAPI* | [**GetExchangeSupportedChains**](docs/WalletsExchangeWalletAPI.md#getexchangesupportedchains) | **Get** /wallets/exchanges/{exchange_id}/assets/supported_chains | List the supported chains by exchange id and asset id
*WalletsExchangeWalletAPI* | [**GetExchangeWalletAssetBalances**](docs/WalletsExchangeWalletAPI.md#getexchangewalletassetbalances) | **Get** /wallets/exchanges/{wallet_id}/assets | List the asset balance in exchange wallet
*WalletsExchangeWalletAPI* | [**LinkSubAccountsByWalletId**](docs/WalletsExchangeWalletAPI.md#linksubaccountsbywalletid) | **Post** /wallets/{wallet_id}/exchanges/subaccounts | Link exchange sub accounts by wallet id
*WalletsExchangeWalletAPI* | [**ListExchanges**](docs/WalletsExchangeWalletAPI.md#listexchanges) | **Get** /wallets/exchanges/settings | List of exchanges
*WalletsExchangeWalletAPI* | [**ListSubAccountsByApikey**](docs/WalletsExchangeWalletAPI.md#listsubaccountsbyapikey) | **Get** /wallets/exchanges/{exchange_id}/subaccounts | List exchange sub accounts by apikey
*WalletsExchangeWalletAPI* | [**ListSubAccountsByWalletId**](docs/WalletsExchangeWalletAPI.md#listsubaccountsbywalletid) | **Get** /wallets/{wallet_id}/exchanges/subaccounts | List exchange sub accounts by wallet id
*WalletsMPCWalletAPI* | [**CancelTssRequest**](docs/WalletsMPCWalletAPI.md#canceltssrequest) | **Put** /wallets/mpc/vaults/{vault_id}/tss_requests/{tss_request_id} | cancel tss request
*WalletsMPCWalletAPI* | [**CreateKeyGroup**](docs/WalletsMPCWalletAPI.md#createkeygroup) | **Post** /wallets/mpc/vaults/{vault_id}/key_groups | create a mpc key group
*WalletsMPCWalletAPI* | [**CreateMpcProject**](docs/WalletsMPCWalletAPI.md#creatempcproject) | **Post** /wallets/mpc/projects | Create a mpc project
*WalletsMPCWalletAPI* | [**CreateMpcVault**](docs/WalletsMPCWalletAPI.md#creatempcvault) | **Post** /wallets/mpc/vaults | Create a mpc vault
*WalletsMPCWalletAPI* | [**CreateTssRequest**](docs/WalletsMPCWalletAPI.md#createtssrequest) | **Post** /wallets/mpc/vaults/{vault_id}/tss_requests | Create a tss request to generate key secrets for a tss group
*WalletsMPCWalletAPI* | [**DeleteKeyGroup**](docs/WalletsMPCWalletAPI.md#deletekeygroup) | **Delete** /wallets/mpc/vaults/{vault_id}/key_groups/{key_group_id} | delete a mpc key group
*WalletsMPCWalletAPI* | [**GetKeyGroup**](docs/WalletsMPCWalletAPI.md#getkeygroup) | **Get** /wallets/mpc/vaults/{vault_id}/key_groups/{key_group_id} | get a mpc key group
*WalletsMPCWalletAPI* | [**GetMpcProject**](docs/WalletsMPCWalletAPI.md#getmpcproject) | **Get** /wallets/mpc/projects/{project_id} | get a mpc project
*WalletsMPCWalletAPI* | [**GetMpcVault**](docs/WalletsMPCWalletAPI.md#getmpcvault) | **Get** /wallets/mpc/vaults/{vault_id} | get a mpc vault
*WalletsMPCWalletAPI* | [**GetTssRequest**](docs/WalletsMPCWalletAPI.md#gettssrequest) | **Get** /wallets/mpc/vaults/{vault_id}/tss_requests/{tss_request_id} | get a tss request
*WalletsMPCWalletAPI* | [**ListCoboKeyHolder**](docs/WalletsMPCWalletAPI.md#listcobokeyholder) | **Get** /wallets/mpc/cobo_key_holders | List all cobo key holders
*WalletsMPCWalletAPI* | [**ListKeyGroup**](docs/WalletsMPCWalletAPI.md#listkeygroup) | **Get** /wallets/mpc/vaults/{vault_id}/key_groups | List all mpc key groups
*WalletsMPCWalletAPI* | [**ListMpcProject**](docs/WalletsMPCWalletAPI.md#listmpcproject) | **Get** /wallets/mpc/projects | List all mpc projects
*WalletsMPCWalletAPI* | [**ListMpcVault**](docs/WalletsMPCWalletAPI.md#listmpcvault) | **Get** /wallets/mpc/vaults | List all mpc vaults
*WalletsMPCWalletAPI* | [**ListTssRequest**](docs/WalletsMPCWalletAPI.md#listtssrequest) | **Get** /wallets/mpc/vaults/{vault_id}/tss_requests | List tss request information of a vault
*WalletsMPCWalletAPI* | [**ModifyMpcVault**](docs/WalletsMPCWalletAPI.md#modifympcvault) | **Put** /wallets/mpc/vaults/{vault_id} | Modify a mpc vault
*WalletsMPCWalletAPI* | [**UpdateKeyGroup**](docs/WalletsMPCWalletAPI.md#updatekeygroup) | **Put** /wallets/mpc/vaults/{vault_id}/key_groups/{key_group_id} | update a mpc key group
*WalletsMPCWalletAPI* | [**UpdateMpcProject**](docs/WalletsMPCWalletAPI.md#updatempcproject) | **Put** /wallets/mpc/projects/{project_id} | update a mpc project


## Documentation For Models

 - [AddWalletAddressRequest](docs/AddWalletAddressRequest.md)
 - [AddressEncoding](docs/AddressEncoding.md)
 - [AddressInfo](docs/AddressInfo.md)
 - [AddressTransferDestination](docs/AddressTransferDestination.md)
 - [AddressTransferSource](docs/AddressTransferSource.md)
 - [AssetBalance](docs/AssetBalance.md)
 - [AssetInfo](docs/AssetInfo.md)
 - [BaseCreateWallet](docs/BaseCreateWallet.md)
 - [BaseTransferSource](docs/BaseTransferSource.md)
 - [ChainFeePrice](docs/ChainFeePrice.md)
 - [ChainInfo](docs/ChainInfo.md)
 - [CreateCustodialWallet](docs/CreateCustodialWallet.md)
 - [CreateExchangeWallet](docs/CreateExchangeWallet.md)
 - [CreateKeyGroupRequest](docs/CreateKeyGroupRequest.md)
 - [CreateKeyGroupRequestKeyHoldersInner](docs/CreateKeyGroupRequestKeyHoldersInner.md)
 - [CreateMpcProjectRequest](docs/CreateMpcProjectRequest.md)
 - [CreateMpcVaultRequest](docs/CreateMpcVaultRequest.md)
 - [CreateMpcWallet](docs/CreateMpcWallet.md)
 - [CreateSafeWallet](docs/CreateSafeWallet.md)
 - [CreateSmartContractWallet](docs/CreateSmartContractWallet.md)
 - [CreateTransferTransaction201Response](docs/CreateTransferTransaction201Response.md)
 - [CreateTssRequestRequest](docs/CreateTssRequestRequest.md)
 - [CreateTssRequestRequestDetailParams](docs/CreateTssRequestRequestDetailParams.md)
 - [CreatedWallet](docs/CreatedWallet.md)
 - [CurveType](docs/CurveType.md)
 - [CustodialWalletInfo](docs/CustodialWalletInfo.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [EstimateFee](docs/EstimateFee.md)
 - [EstimationFee](docs/EstimationFee.md)
 - [EvmEip1559Fee](docs/EvmEip1559Fee.md)
 - [EvmEip1559FeeBasePrice](docs/EvmEip1559FeeBasePrice.md)
 - [EvmEip1559FeePrice](docs/EvmEip1559FeePrice.md)
 - [EvmEip1559FeeSlow](docs/EvmEip1559FeeSlow.md)
 - [EvmEip1559TransactionFee](docs/EvmEip1559TransactionFee.md)
 - [EvmLegacyFee](docs/EvmLegacyFee.md)
 - [EvmLegacyFeeBasePrice](docs/EvmLegacyFeeBasePrice.md)
 - [EvmLegacyFeePrice](docs/EvmLegacyFeePrice.md)
 - [EvmLegacyFeeSlow](docs/EvmLegacyFeeSlow.md)
 - [EvmLegacyTransactionFee](docs/EvmLegacyTransactionFee.md)
 - [ExchangeId](docs/ExchangeId.md)
 - [ExchangeTransferDestination](docs/ExchangeTransferDestination.md)
 - [ExchangeTransferSource](docs/ExchangeTransferSource.md)
 - [ExchangeWalletInfo](docs/ExchangeWalletInfo.md)
 - [ExchangeWalletInfoAllOfSubAccounts](docs/ExchangeWalletInfoAllOfSubAccounts.md)
 - [FeeAmount](docs/FeeAmount.md)
 - [FeeData](docs/FeeData.md)
 - [FeeType](docs/FeeType.md)
 - [FixedFee](docs/FixedFee.md)
 - [GetAddressValidity200Response](docs/GetAddressValidity200Response.md)
 - [GetAssets200Response](docs/GetAssets200Response.md)
 - [GetChains200Response](docs/GetChains200Response.md)
 - [GetExchangeWalletAssetBalances200Response](docs/GetExchangeWalletAssetBalances200Response.md)
 - [GetTokens200Response](docs/GetTokens200Response.md)
 - [GetWalletTokenBalances200Response](docs/GetWalletTokenBalances200Response.md)
 - [KeyGroup](docs/KeyGroup.md)
 - [KeyGroupStatus](docs/KeyGroupStatus.md)
 - [KeyGroupType](docs/KeyGroupType.md)
 - [KeyHolder](docs/KeyHolder.md)
 - [KeyHolderStatus](docs/KeyHolderStatus.md)
 - [KeyHolderType](docs/KeyHolderType.md)
 - [LinkSubAccountsByWalletIdRequest](docs/LinkSubAccountsByWalletIdRequest.md)
 - [ListAddresses200Response](docs/ListAddresses200Response.md)
 - [ListEvents200Response](docs/ListEvents200Response.md)
 - [ListExchanges200ResponseInner](docs/ListExchanges200ResponseInner.md)
 - [ListTransactions200Response](docs/ListTransactions200Response.md)
 - [ListWallets200Response](docs/ListWallets200Response.md)
 - [ListWebhookEventDefinitions200ResponseInner](docs/ListWebhookEventDefinitions200ResponseInner.md)
 - [MPCProject](docs/MPCProject.md)
 - [MPCVault](docs/MPCVault.md)
 - [MPCVaultType](docs/MPCVaultType.md)
 - [MPCWalletInfo](docs/MPCWalletInfo.md)
 - [MaxSendValue](docs/MaxSendValue.md)
 - [ModifyMpcVaultRequest](docs/ModifyMpcVaultRequest.md)
 - [MpcSigningGroup](docs/MpcSigningGroup.md)
 - [MpcTransferSource](docs/MpcTransferSource.md)
 - [Network](docs/Network.md)
 - [Pagination](docs/Pagination.md)
 - [RetryWebhookEvent201Response](docs/RetryWebhookEvent201Response.md)
 - [RootPubkey](docs/RootPubkey.md)
 - [SafeTransferSource](docs/SafeTransferSource.md)
 - [SafeTransferSourceAllOfDelegate](docs/SafeTransferSourceAllOfDelegate.md)
 - [SafeWallet](docs/SafeWallet.md)
 - [SafeWalletAllOfInitiator](docs/SafeWalletAllOfInitiator.md)
 - [SignMessage](docs/SignMessage.md)
 - [SmartContractCall](docs/SmartContractCall.md)
 - [SmartContractWalletInfo](docs/SmartContractWalletInfo.md)
 - [SmartContractWalletOperationType](docs/SmartContractWalletOperationType.md)
 - [SmartContractWalletType](docs/SmartContractWalletType.md)
 - [TSSGroupId](docs/TSSGroupId.md)
 - [TSSRequest](docs/TSSRequest.md)
 - [TSSRequestStatus](docs/TSSRequestStatus.md)
 - [TSSRequestType](docs/TSSRequestType.md)
 - [TokenBalance](docs/TokenBalance.md)
 - [TokenBalanceBalance](docs/TokenBalanceBalance.md)
 - [TokenInfo](docs/TokenInfo.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionAddress](docs/TransactionAddress.md)
 - [TransactionAddressType](docs/TransactionAddressType.md)
 - [TransactionApprover](docs/TransactionApprover.md)
 - [TransactionDetails](docs/TransactionDetails.md)
 - [TransactionFee](docs/TransactionFee.md)
 - [TransactionInitiatorType](docs/TransactionInitiatorType.md)
 - [TransactionSigner](docs/TransactionSigner.md)
 - [TransactionStatus](docs/TransactionStatus.md)
 - [TransactionSubStatus](docs/TransactionSubStatus.md)
 - [TransactionTimeline](docs/TransactionTimeline.md)
 - [TransactionTokeApproval](docs/TransactionTokeApproval.md)
 - [TransactionToken](docs/TransactionToken.md)
 - [TransactionType](docs/TransactionType.md)
 - [Transfer](docs/Transfer.md)
 - [TransferDestination](docs/TransferDestination.md)
 - [TransferDestinationType](docs/TransferDestinationType.md)
 - [TransferSource](docs/TransferSource.md)
 - [UTXO](docs/UTXO.md)
 - [UpdateMpcProjectRequest](docs/UpdateMpcProjectRequest.md)
 - [UpdateWalletByIdRequest](docs/UpdateWalletByIdRequest.md)
 - [UtxoFee](docs/UtxoFee.md)
 - [UtxoFeeBasePrice](docs/UtxoFeeBasePrice.md)
 - [UtxoFeePrice](docs/UtxoFeePrice.md)
 - [UtxoFeeSlow](docs/UtxoFeeSlow.md)
 - [UtxoTransactionFee](docs/UtxoTransactionFee.md)
 - [WalletInfo](docs/WalletInfo.md)
 - [WalletSubtype](docs/WalletSubtype.md)
 - [WalletType](docs/WalletType.md)
 - [WebhookEvent](docs/WebhookEvent.md)
 - [WebhookEventLog](docs/WebhookEventLog.md)
 - [WebhookEventStatus](docs/WebhookEventStatus.md)
 - [WebhookEventType](docs/WebhookEventType.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### CoboAuth

- **Type**: API key
- **API key parameter name**: BIZ-API-KEY
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: BIZ-API-KEY and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		CoboWaas2.ContextAPIKeys,
		map[string]CoboWaas2.APIKey{
			"BIZ-API-KEY": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@cobo.com

