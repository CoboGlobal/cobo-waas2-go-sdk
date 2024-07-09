# \WalletsExchangeWalletAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExchangeSupportedAssets**](WalletsExchangeWalletAPI.md#GetExchangeSupportedAssets) | **Get** /wallets/exchanges/{exchange_id}/supported_assets | List the supported assets by exchange id
[**GetExchangeSupportedChains**](WalletsExchangeWalletAPI.md#GetExchangeSupportedChains) | **Get** /wallets/exchanges/{exchange_id}/assets/supported_chains | List the supported chains by exchange id and asset id
[**GetExchangeWalletAssetBalances**](WalletsExchangeWalletAPI.md#GetExchangeWalletAssetBalances) | **Get** /wallets/{wallet_id}/exchanges/assets | List the asset balance in exchange wallet
[**LinkSubAccountsByWalletId**](WalletsExchangeWalletAPI.md#LinkSubAccountsByWalletId) | **Post** /wallets/{wallet_id}/exchanges/subaccounts | Link exchange sub accounts by wallet id
[**ListExchanges**](WalletsExchangeWalletAPI.md#ListExchanges) | **Get** /wallets/exchanges/settings | List of exchanges
[**ListSubAccountsByApikey**](WalletsExchangeWalletAPI.md#ListSubAccountsByApikey) | **Get** /wallets/exchanges/{exchange_id}/subaccounts | List exchange sub accounts by apikey
[**ListSubAccountsByWalletId**](WalletsExchangeWalletAPI.md#ListSubAccountsByWalletId) | **Get** /wallets/{wallet_id}/exchanges/subaccounts | List exchange sub accounts by wallet id



## GetExchangeSupportedAssets

> GetExchangeSupportedAssets200Response GetExchangeSupportedAssets(ctx, exchangeId).Limit(limit).Before(before).After(after).Execute()

List the supported assets by exchange id



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
	exchangeId := CoboWaas2.ExchangeId("binance") // ExchangeId | Exchange ID to query
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
	resp, r, err := apiClient.WalletsExchangeWalletAPI.GetExchangeSupportedAssets(ctx, exchangeId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsExchangeWalletAPI.GetExchangeSupportedAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExchangeSupportedAssets`: GetExchangeSupportedAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsExchangeWalletAPI.GetExchangeSupportedAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**exchangeId** | [**ExchangeId**](.md) | Exchange ID to query | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeSupportedAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**GetExchangeSupportedAssets200Response**](GetExchangeSupportedAssets200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeSupportedChains

> []ChainInfo GetExchangeSupportedChains(ctx, exchangeId).AssetId(assetId).Limit(limit).Before(before).After(after).Execute()

List the supported chains by exchange id and asset id



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
	exchangeId := CoboWaas2.ExchangeId("binance") // ExchangeId | Exchange ID to query
	assetId := "USDT" // string | (This concept applies to Exchange Wallets only) The asset ID. An asset is a digital representation of a valuable resource on a blockchain network. Exchange Wallets group your holdings by asset, even if the same asset exists on different blockchains. For example, if your Exchange Wallet has 1 USDT on Ethereum and 1 USDT on TRON, then your asset balance is 2 USDT.
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
	resp, r, err := apiClient.WalletsExchangeWalletAPI.GetExchangeSupportedChains(ctx, exchangeId).AssetId(assetId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsExchangeWalletAPI.GetExchangeSupportedChains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExchangeSupportedChains`: []ChainInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsExchangeWalletAPI.GetExchangeSupportedChains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**exchangeId** | [**ExchangeId**](.md) | Exchange ID to query | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeSupportedChainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetId** | **string** | (This concept applies to Exchange Wallets only) The asset ID. An asset is a digital representation of a valuable resource on a blockchain network. Exchange Wallets group your holdings by asset, even if the same asset exists on different blockchains. For example, if your Exchange Wallet has 1 USDT on Ethereum and 1 USDT on TRON, then your asset balance is 2 USDT. | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

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


## GetExchangeWalletAssetBalances

> GetExchangeWalletAssetBalances200Response GetExchangeWalletAssetBalances(ctx, walletId).SubWalletId(subWalletId).AssetId(assetId).Limit(limit).Before(before).After(after).Execute()

List the asset balance in exchange wallet



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.
	subWalletId := "SPOT" // string | Unique id of the wallet
	assetId := "USDT" // string | (This concept applies to Exchange Wallets only) The asset ID. An asset is a digital representation of a valuable resource on a blockchain network. Exchange Wallets group your holdings by asset, even if the same asset exists on different blockchains. For example, if your Exchange Wallet has 1 USDT on Ethereum and 1 USDT on TRON, then your asset balance is 2 USDT. (optional)
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
	resp, r, err := apiClient.WalletsExchangeWalletAPI.GetExchangeWalletAssetBalances(ctx, walletId).SubWalletId(subWalletId).AssetId(assetId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsExchangeWalletAPI.GetExchangeWalletAssetBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExchangeWalletAssetBalances`: GetExchangeWalletAssetBalances200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsExchangeWalletAPI.GetExchangeWalletAssetBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeWalletAssetBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subWalletId** | **string** | Unique id of the wallet | 
 **assetId** | **string** | (This concept applies to Exchange Wallets only) The asset ID. An asset is a digital representation of a valuable resource on a blockchain network. Exchange Wallets group your holdings by asset, even if the same asset exists on different blockchains. For example, if your Exchange Wallet has 1 USDT on Ethereum and 1 USDT on TRON, then your asset balance is 2 USDT. | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**GetExchangeWalletAssetBalances200Response**](GetExchangeWalletAssetBalances200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkSubAccountsByWalletId

> []string LinkSubAccountsByWalletId(ctx, walletId).LinkSubAccountsByWalletIdRequest(linkSubAccountsByWalletIdRequest).Execute()

Link exchange sub accounts by wallet id



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.
	linkSubAccountsByWalletIdRequest := *CoboWaas2.NewLinkSubAccountsByWalletIdRequest() // LinkSubAccountsByWalletIdRequest | Request body for linking subaccounts (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsExchangeWalletAPI.LinkSubAccountsByWalletId(ctx, walletId).LinkSubAccountsByWalletIdRequest(linkSubAccountsByWalletIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsExchangeWalletAPI.LinkSubAccountsByWalletId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkSubAccountsByWalletId`: []string
	fmt.Fprintf(os.Stdout, "Response from `WalletsExchangeWalletAPI.LinkSubAccountsByWalletId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkSubAccountsByWalletIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **linkSubAccountsByWalletIdRequest** | [**LinkSubAccountsByWalletIdRequest**](LinkSubAccountsByWalletIdRequest.md) | Request body for linking subaccounts | 

### Return type

**[]string**

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExchanges

> []ListExchanges200ResponseInner ListExchanges(ctx).Execute()

List of exchanges



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

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
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


## ListSubAccountsByApikey

> []string ListSubAccountsByApikey(ctx, exchangeId).Apikey(apikey).Secret(secret).Passphrase(passphrase).Execute()

List exchange sub accounts by apikey



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
	exchangeId := CoboWaas2.ExchangeId("binance") // ExchangeId | Exchange ID to query
	apikey := "d8f062da-39f4-4a11-8b9d-12595854237f" // string | The API Key for the exchange (optional)
	secret := "75B4F636193162488A3728B4A5797708" // string | The API Secret for the exchange. (optional)
	passphrase := "A3DBHJV" // string | The API passphrase for the exchange wallet. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsExchangeWalletAPI.ListSubAccountsByApikey(ctx, exchangeId).Apikey(apikey).Secret(secret).Passphrase(passphrase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsExchangeWalletAPI.ListSubAccountsByApikey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubAccountsByApikey`: []string
	fmt.Fprintf(os.Stdout, "Response from `WalletsExchangeWalletAPI.ListSubAccountsByApikey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**exchangeId** | [**ExchangeId**](.md) | Exchange ID to query | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubAccountsByApikeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apikey** | **string** | The API Key for the exchange | 
 **secret** | **string** | The API Secret for the exchange. | 
 **passphrase** | **string** | The API passphrase for the exchange wallet. | 

### Return type

**[]string**

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubAccountsByWalletId

> []string ListSubAccountsByWalletId(ctx, walletId).Execute()

List exchange sub accounts by wallet id



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsExchangeWalletAPI.ListSubAccountsByWalletId(ctx, walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsExchangeWalletAPI.ListSubAccountsByWalletId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubAccountsByWalletId`: []string
	fmt.Fprintf(os.Stdout, "Response from `WalletsExchangeWalletAPI.ListSubAccountsByWalletId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubAccountsByWalletIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

