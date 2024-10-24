# \WalletsExchangeWalletAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAssetBalancesForExchangeWallet**](WalletsExchangeWalletAPI.md#ListAssetBalancesForExchangeWallet) | **Get** /wallets/{wallet_id}/exchanges/assets | List asset balances
[**ListExchanges**](WalletsExchangeWalletAPI.md#ListExchanges) | **Get** /wallets/exchanges | List supported exchanges
[**ListSupportedAssetsForExchange**](WalletsExchangeWalletAPI.md#ListSupportedAssetsForExchange) | **Get** /wallets/exchanges/{exchange_id}/assets | List supported assets
[**ListSupportedChainsForExchange**](WalletsExchangeWalletAPI.md#ListSupportedChainsForExchange) | **Get** /wallets/exchanges/{exchange_id}/assets/{asset_id}/chains | List supported chains



## ListAssetBalancesForExchangeWallet

> ListAssetBalancesForExchangeWallet200Response ListAssetBalancesForExchangeWallet(ctx, walletId).TradingAccountTypes(tradingAccountTypes).AssetIds(assetIds).Limit(limit).Before(before).After(after).Execute()

List asset balances



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
	tradingAccountTypes := "Trading,Funding"
	assetIds := "USDT,USDC"
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
	resp, r, err := apiClient.WalletsExchangeWalletAPI.ListAssetBalancesForExchangeWallet(ctx, walletId).TradingAccountTypes(tradingAccountTypes).AssetIds(assetIds).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsExchangeWalletAPI.ListAssetBalancesForExchangeWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssetBalancesForExchangeWallet`: ListAssetBalancesForExchangeWallet200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsExchangeWalletAPI.ListAssetBalancesForExchangeWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAssetBalancesForExchangeWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tradingAccountTypes** | **string** | A list of trading account types, separated by comma. You can get the the supported trading account types of an exchange by calling [List supported exchanges](/v2/api-references/wallets--exchange-wallet/list-supported-exchanges). | 
 **assetIds** | **string** | (This concept applies to Exchange Wallets only) A list of asset IDs, separated by comma. An asset ID is the unique identifier of the asset held within your linked exchange account. | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify &#x60;before&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  - If you set &#x60;before&#x60; to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify &#x60;after&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 

### Return type

[**ListAssetBalancesForExchangeWallet200Response**](ListAssetBalancesForExchangeWallet200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExchanges

> []ListExchanges200ResponseInner ListExchanges(ctx).Execute()

List supported exchanges



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
	resp, r, err := apiClient.WalletsExchangeWalletAPI.ListExchanges(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsExchangeWalletAPI.ListExchanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExchanges`: []ListExchanges200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `WalletsExchangeWalletAPI.ListExchanges`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListExchangesRequest struct via the builder pattern


### Return type

[**[]ListExchanges200ResponseInner**](ListExchanges200ResponseInner.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportedAssetsForExchange

> ListSupportedAssetsForExchange200Response ListSupportedAssetsForExchange(ctx, exchangeId).Limit(limit).Before(before).After(after).Execute()

List supported assets



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
	exchangeId := coboWaas2.ExchangeId("binance")
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
	resp, r, err := apiClient.WalletsExchangeWalletAPI.ListSupportedAssetsForExchange(ctx, exchangeId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsExchangeWalletAPI.ListSupportedAssetsForExchange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSupportedAssetsForExchange`: ListSupportedAssetsForExchange200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsExchangeWalletAPI.ListSupportedAssetsForExchange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**exchangeId** | [**ExchangeId**](.md) | The ID of the exchange. Possible values include:   - &#x60;binance&#x60;: Binance.   - &#x60;okx&#x60;: OKX.   - &#x60;deribit&#x60;: Deribit.   - &#x60;bybit&#x60;: Bybit.   - &#x60;gate&#x60;: Gate.io   - &#x60;bitget&#x60;: Bitget   - &#x60;bitmart&#x60;: BitMart   - &#x60;bitfinex&#x60;: Bitfinex  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportedAssetsForExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify &#x60;before&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  - If you set &#x60;before&#x60; to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify &#x60;after&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 

### Return type

[**ListSupportedAssetsForExchange200Response**](ListSupportedAssetsForExchange200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportedChainsForExchange

> []ChainInfo ListSupportedChainsForExchange(ctx, exchangeId, assetId).Execute()

List supported chains



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
	exchangeId := coboWaas2.ExchangeId("binance")
	assetId := "USDT"

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
	resp, r, err := apiClient.WalletsExchangeWalletAPI.ListSupportedChainsForExchange(ctx, exchangeId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsExchangeWalletAPI.ListSupportedChainsForExchange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSupportedChainsForExchange`: []ChainInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsExchangeWalletAPI.ListSupportedChainsForExchange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**exchangeId** | [**ExchangeId**](.md) | The ID of the exchange. Possible values include:   - &#x60;binance&#x60;: Binance.   - &#x60;okx&#x60;: OKX.   - &#x60;deribit&#x60;: Deribit.   - &#x60;bybit&#x60;: Bybit.   - &#x60;gate&#x60;: Gate.io   - &#x60;bitget&#x60;: Bitget   - &#x60;bitmart&#x60;: BitMart   - &#x60;bitfinex&#x60;: Bitfinex  | 
**assetId** | **string** | (This concept applies to Exchange Wallets only) The asset ID. An asset ID is the unique identifier of the asset held within your linked exchange account. You can get the ID of the assets supported by an exchanges by calling [List supported assets](/v2/api-references/wallets--exchange-wallet/list-supported-assets). | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportedChainsForExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ChainInfo**](ChainInfo.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

