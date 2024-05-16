# \WalletsAPI

All URIs are relative to *https://api.cobo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWalletAddress**](WalletsAPI.md#AddWalletAddress) | **Post** /wallets/{wallet_id}/addresses | Add address to a wallet
[**CreateWallet**](WalletsAPI.md#CreateWallet) | **Post** /wallets | Create new wallet
[**DeleteWalletById**](WalletsAPI.md#DeleteWalletById) | **Delete** /wallets/{wallet_id} | Delete a wallet by ID
[**GetAddressValidity**](WalletsAPI.md#GetAddressValidity) | **Get** /wallets/address/validity | Get the given address validity for token
[**GetAssets**](WalletsAPI.md#GetAssets) | **Get** /wallets/assets | List the metadata of assets
[**GetChains**](WalletsAPI.md#GetChains) | **Get** /wallets/chains | List the metadata of chain
[**GetMaxSendValue**](WalletsAPI.md#GetMaxSendValue) | **Get** /wallets/{wallet_id}/max_sendable_value | Get max sendable Vaule
[**GetSpendableList**](WalletsAPI.md#GetSpendableList) | **Get** /wallets/{wallet_id}/spendables | List the spendable utxo
[**GetSupportedChains**](WalletsAPI.md#GetSupportedChains) | **Get** /wallets/supported_chains | List the supported chains by wallet subtype
[**GetSupportedTokens**](WalletsAPI.md#GetSupportedTokens) | **Get** /wallets/supported_tokens | List the supported tokens by wallet subtype and chain id if specified
[**GetTokens**](WalletsAPI.md#GetTokens) | **Get** /wallets/tokens | List the metadata of tokens
[**GetWalletAddressById**](WalletsAPI.md#GetWalletAddressById) | **Get** /wallets/{wallet_id}/addresses/{address_id} | Get address information by ID
[**GetWalletAddressTokenBalances**](WalletsAPI.md#GetWalletAddressTokenBalances) | **Get** /wallets/{wallet_id}/addresses/{address_id}/tokens | List the token balance by address in the wallets(to be specific)
[**GetWalletById**](WalletsAPI.md#GetWalletById) | **Get** /wallets/{wallet_id} | Get wallet information by ID
[**GetWalletTokenBalances**](WalletsAPI.md#GetWalletTokenBalances) | **Get** /wallets/{wallet_id}/tokens | List the token balance in the wallets(to be specific)
[**ListAddresses**](WalletsAPI.md#ListAddresses) | **Get** /wallets/{wallet_id}/addresses | List wallet addresses by wallet ID
[**ListWallets**](WalletsAPI.md#ListWallets) | **Get** /wallets | List all wallets
[**UpdateWalletById**](WalletsAPI.md#UpdateWalletById) | **Put** /wallets/{wallet_id} | Update wallet by ID



## AddWalletAddress

> []AddressInfo AddWalletAddress(ctx, walletId).AddWalletAddressRequest(addWalletAddressRequest).Execute()

Add address to a wallet



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the wallet
	addWalletAddressRequest := *CoboWaas2.NewAddWalletAddressRequest("ETH_USDT", int32(1)) // AddWalletAddressRequest | The request body to add address for a wallet (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.AddWalletAddress(ctx, walletId).AddWalletAddressRequest(addWalletAddressRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.AddWalletAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddWalletAddress`: []AddressInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.AddWalletAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Unique id of the wallet | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWalletAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addWalletAddressRequest** | [**AddWalletAddressRequest**](AddWalletAddressRequest.md) | The request body to add address for a wallet | 

### Return type

[**[]AddressInfo**](AddressInfo.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWallet

> WalletInfo CreateWallet(ctx).CreatedWallet(createdWallet).Execute()

Create new wallet



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
	createdWallet := CoboWaas2.CreatedWallet{CreateCustodialWallet: CoboWaas2.NewCreateCustodialWallet("My WaaS 2.0 Wallet", "WalletType_example", "WalletSubtype_example")} // CreatedWallet | The request body to create a wallet (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.CreateWallet(ctx).CreatedWallet(createdWallet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CreateWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWallet`: WalletInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CreateWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdWallet** | [**CreatedWallet**](CreatedWallet.md) | The request body to create a wallet | 

### Return type

[**WalletInfo**](WalletInfo.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWalletById

> DeleteWalletById(ctx, walletId).Execute()

Delete a wallet by ID



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the wallet

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	r, err := apiClient.WalletsAPI.DeleteWalletById(ctx, walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.DeleteWalletById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Unique id of the wallet | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWalletByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressValidity

> GetAddressValidity200Response GetAddressValidity(ctx).TokenId(tokenId).AddressStr(addressStr).Execute()

Get the given address validity for token



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
	tokenId := "ETH_USDT" // string | Unique id of the token
	addressStr := "0x0000000000000000000000000000000000000000" // string | The address string

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetAddressValidity(ctx).TokenId(tokenId).AddressStr(addressStr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetAddressValidity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressValidity`: GetAddressValidity200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetAddressValidity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressValidityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenId** | **string** | Unique id of the token | 
 **addressStr** | **string** | The address string | 

### Return type

[**GetAddressValidity200Response**](GetAddressValidity200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssets

> GetAssets200Response GetAssets(ctx).AssetId(assetId).Limit(limit).Before(before).After(after).Execute()

List the metadata of assets



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
	assetId := "USDT" // string | Unique id of the asset (optional)
	limit := int32(10) // int32 | size of page to return (pagination) (optional) (default to 10)
	before := "123" // string | Cursor string received from previous request (optional) (default to "")
	after := "123" // string | Cursor string received from previous request (optional) (default to "")

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetAssets(ctx).AssetId(assetId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssets`: GetAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetId** | **string** | Unique id of the asset | 
 **limit** | **int32** | size of page to return (pagination) | [default to 10]
 **before** | **string** | Cursor string received from previous request | [default to &quot;&quot;]
 **after** | **string** | Cursor string received from previous request | [default to &quot;&quot;]

### Return type

[**GetAssets200Response**](GetAssets200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChains

> GetChains200Response GetChains(ctx).ChainId(chainId).Limit(limit).Before(before).After(after).Execute()

List the metadata of chain



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
	limit := int32(10) // int32 | size of page to return (pagination) (optional) (default to 10)
	before := "123" // string | Cursor string received from previous request (optional) (default to "")
	after := "123" // string | Cursor string received from previous request (optional) (default to "")

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetChains(ctx).ChainId(chainId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetChains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChains`: GetChains200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetChains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chainId** | **string** | Unique id of the chain | 
 **limit** | **int32** | size of page to return (pagination) | [default to 10]
 **before** | **string** | Cursor string received from previous request | [default to &quot;&quot;]
 **after** | **string** | Cursor string received from previous request | [default to &quot;&quot;]

### Return type

[**GetChains200Response**](GetChains200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxSendValue

> MaxSendValue GetMaxSendValue(ctx, walletId).ToAddress(toAddress).FromAddress(fromAddress).Execute()

Get max sendable Vaule



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the wallet
	toAddress := "2N2xFZtbCFB6Nb3Pj9Sxsx5mX2fxX3yEgkE" // string | address
	fromAddress := "2N2xFZtbCFB6Nb3Pj9Sxsx5mX2fxX3yEgkE" // string | address (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetMaxSendValue(ctx, walletId).ToAddress(toAddress).FromAddress(fromAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetMaxSendValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaxSendValue`: MaxSendValue
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetMaxSendValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Unique id of the wallet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxSendValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **toAddress** | **string** | address | 
 **fromAddress** | **string** | address | 

### Return type

[**MaxSendValue**](MaxSendValue.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpendableList

> []UTXO GetSpendableList(ctx, walletId).TokenId(tokenId).AddressStr(addressStr).Execute()

List the spendable utxo



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the wallet
	tokenId := "ETH_USDT" // string | Unique id of the token
	addressStr := "2N2xFZtbCFB6Nb3Pj9Sxsx5mX2fxX3yEgkE" // string | address (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetSpendableList(ctx, walletId).TokenId(tokenId).AddressStr(addressStr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetSpendableList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpendableList`: []UTXO
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetSpendableList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Unique id of the wallet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpendableListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenId** | **string** | Unique id of the token | 
 **addressStr** | **string** | address | 

### Return type

[**[]UTXO**](UTXO.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedChains

> GetChains200Response GetSupportedChains(ctx).WalletSubtype(walletSubtype).Limit(limit).Before(before).After(after).Execute()

List the supported chains by wallet subtype



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
	walletSubtype := CoboWaas2.WalletSubtype("Asset") // WalletSubtype | Wallet subtype to query
	limit := int32(10) // int32 | size of page to return (pagination) (optional) (default to 10)
	before := "123" // string | Cursor string received from previous request (optional) (default to "")
	after := "123" // string | Cursor string received from previous request (optional) (default to "")

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetSupportedChains(ctx).WalletSubtype(walletSubtype).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetSupportedChains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportedChains`: GetChains200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetSupportedChains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedChainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletSubtype** | [**WalletSubtype**](WalletSubtype.md) | Wallet subtype to query | 
 **limit** | **int32** | size of page to return (pagination) | [default to 10]
 **before** | **string** | Cursor string received from previous request | [default to &quot;&quot;]
 **after** | **string** | Cursor string received from previous request | [default to &quot;&quot;]

### Return type

[**GetChains200Response**](GetChains200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedTokens

> GetTokens200Response GetSupportedTokens(ctx).WalletSubtype(walletSubtype).ChainId(chainId).Limit(limit).Before(before).After(after).Execute()

List the supported tokens by wallet subtype and chain id if specified



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
	walletSubtype := CoboWaas2.WalletSubtype("Asset") // WalletSubtype | Wallet subtype to query
	chainId := "ETH" // string | Unique id of the chain (optional)
	limit := int32(10) // int32 | size of page to return (pagination) (optional) (default to 10)
	before := "123" // string | Cursor string received from previous request (optional) (default to "")
	after := "123" // string | Cursor string received from previous request (optional) (default to "")

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetSupportedTokens(ctx).WalletSubtype(walletSubtype).ChainId(chainId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetSupportedTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportedTokens`: GetTokens200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetSupportedTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletSubtype** | [**WalletSubtype**](WalletSubtype.md) | Wallet subtype to query | 
 **chainId** | **string** | Unique id of the chain | 
 **limit** | **int32** | size of page to return (pagination) | [default to 10]
 **before** | **string** | Cursor string received from previous request | [default to &quot;&quot;]
 **after** | **string** | Cursor string received from previous request | [default to &quot;&quot;]

### Return type

[**GetTokens200Response**](GetTokens200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokens

> GetTokens200Response GetTokens(ctx).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()

List the metadata of tokens



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
	tokenId := "ETH_USDT" // string | Unique id of the token (optional)
	limit := int32(10) // int32 | size of page to return (pagination) (optional) (default to 10)
	before := "123" // string | Cursor string received from previous request (optional) (default to "")
	after := "123" // string | Cursor string received from previous request (optional) (default to "")

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetTokens(ctx).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokens`: GetTokens200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenId** | **string** | Unique id of the token | 
 **limit** | **int32** | size of page to return (pagination) | [default to 10]
 **before** | **string** | Cursor string received from previous request | [default to &quot;&quot;]
 **after** | **string** | Cursor string received from previous request | [default to &quot;&quot;]

### Return type

[**GetTokens200Response**](GetTokens200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletAddressById

> AddressInfo GetWalletAddressById(ctx, walletId, addressId).Execute()

Get address information by ID



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the wallet
	addressId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the address

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetWalletAddressById(ctx, walletId, addressId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletAddressById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletAddressById`: AddressInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletAddressById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Unique id of the wallet | 
**addressId** | **string** | Unique id of the address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletAddressByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AddressInfo**](AddressInfo.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletAddressTokenBalances

> GetWalletTokenBalances200Response GetWalletAddressTokenBalances(ctx, walletId, addressId).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()

List the token balance by address in the wallets(to be specific)



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the wallet
	addressId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the address
	tokenId := "ETH_USDT" // string | Unique id of the token (optional)
	limit := int32(10) // int32 | size of page to return (pagination) (optional) (default to 10)
	before := "123" // string | Cursor string received from previous request (optional) (default to "")
	after := "123" // string | Cursor string received from previous request (optional) (default to "")

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetWalletAddressTokenBalances(ctx, walletId, addressId).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletAddressTokenBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletAddressTokenBalances`: GetWalletTokenBalances200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletAddressTokenBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Unique id of the wallet | 
**addressId** | **string** | Unique id of the address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletAddressTokenBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenId** | **string** | Unique id of the token | 
 **limit** | **int32** | size of page to return (pagination) | [default to 10]
 **before** | **string** | Cursor string received from previous request | [default to &quot;&quot;]
 **after** | **string** | Cursor string received from previous request | [default to &quot;&quot;]

### Return type

[**GetWalletTokenBalances200Response**](GetWalletTokenBalances200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletById

> WalletInfo GetWalletById(ctx, walletId).Execute()

Get wallet information by ID



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the wallet

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetWalletById(ctx, walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletById`: WalletInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Unique id of the wallet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WalletInfo**](WalletInfo.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletTokenBalances

> GetWalletTokenBalances200Response GetWalletTokenBalances(ctx, walletId).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()

List the token balance in the wallets(to be specific)



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the wallet
	tokenId := "ETH_USDT" // string | Unique id of the token (optional)
	limit := int32(10) // int32 | size of page to return (pagination) (optional) (default to 10)
	before := "123" // string | Cursor string received from previous request (optional) (default to "")
	after := "123" // string | Cursor string received from previous request (optional) (default to "")

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetWalletTokenBalances(ctx, walletId).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletTokenBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletTokenBalances`: GetWalletTokenBalances200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletTokenBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Unique id of the wallet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletTokenBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenId** | **string** | Unique id of the token | 
 **limit** | **int32** | size of page to return (pagination) | [default to 10]
 **before** | **string** | Cursor string received from previous request | [default to &quot;&quot;]
 **after** | **string** | Cursor string received from previous request | [default to &quot;&quot;]

### Return type

[**GetWalletTokenBalances200Response**](GetWalletTokenBalances200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAddresses

> ListAddresses200Response ListAddresses(ctx, walletId).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()

List wallet addresses by wallet ID



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the wallet
	tokenId := "ETH_USDT" // string | Unique id of the token (optional)
	limit := int32(10) // int32 | size of page to return (pagination) (optional) (default to 10)
	before := "123" // string | Cursor string received from previous request (optional) (default to "")
	after := "123" // string | Cursor string received from previous request (optional) (default to "")

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.ListAddresses(ctx, walletId).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.ListAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAddresses`: ListAddresses200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.ListAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Unique id of the wallet | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenId** | **string** | Unique id of the token | 
 **limit** | **int32** | size of page to return (pagination) | [default to 10]
 **before** | **string** | Cursor string received from previous request | [default to &quot;&quot;]
 **after** | **string** | Cursor string received from previous request | [default to &quot;&quot;]

### Return type

[**ListAddresses200Response**](ListAddresses200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWallets

> ListWallets200Response ListWallets(ctx).WalletType(walletType).WalletSubtype(walletSubtype).VaultId(vaultId).Limit(limit).Before(before).After(after).Execute()

List all wallets



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
	walletType := CoboWaas2.WalletType("Custodial") // WalletType | Wallet type to query (optional)
	walletSubtype := CoboWaas2.WalletSubtype("Asset") // WalletSubtype | Wallet subtype to query (optional)
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault (optional)
	limit := int32(10) // int32 | size of page to return (pagination) (optional) (default to 10)
	before := "123" // string | Cursor string received from previous request (optional) (default to "")
	after := "123" // string | Cursor string received from previous request (optional) (default to "")

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.ListWallets(ctx).WalletType(walletType).WalletSubtype(walletSubtype).VaultId(vaultId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.ListWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWallets`: ListWallets200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.ListWallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletType** | [**WalletType**](WalletType.md) | Wallet type to query | 
 **walletSubtype** | [**WalletSubtype**](WalletSubtype.md) | Wallet subtype to query | 
 **vaultId** | **string** | Unique id of the mpc vault | 
 **limit** | **int32** | size of page to return (pagination) | [default to 10]
 **before** | **string** | Cursor string received from previous request | [default to &quot;&quot;]
 **after** | **string** | Cursor string received from previous request | [default to &quot;&quot;]

### Return type

[**ListWallets200Response**](ListWallets200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWalletById

> WalletInfo UpdateWalletById(ctx, walletId).UpdateWalletByIdRequest(updateWalletByIdRequest).Execute()

Update wallet by ID



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the wallet
	updateWalletByIdRequest := *CoboWaas2.NewUpdateWalletByIdRequest() // UpdateWalletByIdRequest | The request body to update a wallet (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.UpdateWalletById(ctx, walletId).UpdateWalletByIdRequest(updateWalletByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.UpdateWalletById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWalletById`: WalletInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.UpdateWalletById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Unique id of the wallet | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWalletByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWalletByIdRequest** | [**UpdateWalletByIdRequest**](UpdateWalletByIdRequest.md) | The request body to update a wallet | 

### Return type

[**WalletInfo**](WalletInfo.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

