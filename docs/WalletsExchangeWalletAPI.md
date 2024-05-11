# \WalletsExchangeWalletAPI

All URIs are relative to *https://api.cobo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExchangeSupportedAssets**](WalletsExchangeWalletAPI.md#GetExchangeSupportedAssets) | **Get** /wallets/exchanges/{exchange_id}/supported_assets | List the supported assets by exchange id
[**GetExchangeSupportedChains**](WalletsExchangeWalletAPI.md#GetExchangeSupportedChains) | **Get** /wallets/exchanges/{exchange_id}/assets/supported_chains | List the supported chains by exchange id and asset id
[**GetExchangeWalletAssetBalances**](WalletsExchangeWalletAPI.md#GetExchangeWalletAssetBalances) | **Get** /wallets/exchanges/{wallet_id}/assets | List the asset balance in exchange wallet
[**LinkSubAccountsByWalletId**](WalletsExchangeWalletAPI.md#LinkSubAccountsByWalletId) | **Post** /wallets/{wallet_id}/exchanges/subaccounts | Link exchange sub accounts by wallet id
[**ListExchanges**](WalletsExchangeWalletAPI.md#ListExchanges) | **Get** /wallets/exchanges/settings | List of exchanges
[**ListSubAccountsByApikey**](WalletsExchangeWalletAPI.md#ListSubAccountsByApikey) | **Get** /wallets/exchanges/{exchange_id}/subaccounts | List exchange sub accounts by apikey
[**ListSubAccountsByWalletId**](WalletsExchangeWalletAPI.md#ListSubAccountsByWalletId) | **Get** /wallets/{wallet_id}/exchanges/subaccounts | List exchange sub accounts by wallet id



## GetExchangeSupportedAssets

> GetAssets200Response GetExchangeSupportedAssets(ctx, exchangeId).Limit(limit).Before(before).After(after).Execute()

List the supported assets by exchange id



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
	exchangeId := openapiclient.ExchangeId("binance") // ExchangeId | Exchange ID to query
	limit := int32(10) // int32 | size of page to return (pagination) (optional) (default to 10)
	before := "123" // string | Cursor string received from previous request (optional) (default to "")
	after := "123" // string | Cursor string received from previous request (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsExchangeWalletAPI.GetExchangeSupportedAssets(context.Background(), exchangeId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsExchangeWalletAPI.GetExchangeSupportedAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExchangeSupportedAssets`: GetAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsExchangeWalletAPI.GetExchangeSupportedAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchangeId** | [**ExchangeId**](.md) | Exchange ID to query | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeSupportedAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	exchangeId := openapiclient.ExchangeId("binance") // ExchangeId | Exchange ID to query
	assetId := "USDT" // string | Unique id of the asset
	limit := int32(10) // int32 | size of page to return (pagination) (optional) (default to 10)
	before := "123" // string | Cursor string received from previous request (optional) (default to "")
	after := "123" // string | Cursor string received from previous request (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsExchangeWalletAPI.GetExchangeSupportedChains(context.Background(), exchangeId).AssetId(assetId).Limit(limit).Before(before).After(after).Execute()
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchangeId** | [**ExchangeId**](.md) | Exchange ID to query | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeSupportedChainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetId** | **string** | Unique id of the asset | 
 **limit** | **int32** | size of page to return (pagination) | [default to 10]
 **before** | **string** | Cursor string received from previous request | [default to &quot;&quot;]
 **after** | **string** | Cursor string received from previous request | [default to &quot;&quot;]

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
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the wallet
	subWalletId := "SPOT" // string | Unique id of the wallet
	assetId := "USDT" // string | Unique id of the asset (optional)
	limit := int32(10) // int32 | size of page to return (pagination) (optional) (default to 10)
	before := "123" // string | Cursor string received from previous request (optional) (default to "")
	after := "123" // string | Cursor string received from previous request (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsExchangeWalletAPI.GetExchangeWalletAssetBalances(context.Background(), walletId).SubWalletId(subWalletId).AssetId(assetId).Limit(limit).Before(before).After(after).Execute()
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Unique id of the wallet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeWalletAssetBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subWalletId** | **string** | Unique id of the wallet | 
 **assetId** | **string** | Unique id of the asset | 
 **limit** | **int32** | size of page to return (pagination) | [default to 10]
 **before** | **string** | Cursor string received from previous request | [default to &quot;&quot;]
 **after** | **string** | Cursor string received from previous request | [default to &quot;&quot;]

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
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the wallet
	linkSubAccountsByWalletIdRequest := *openapiclient.NewLinkSubAccountsByWalletIdRequest() // LinkSubAccountsByWalletIdRequest | Request body for linking subaccounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsExchangeWalletAPI.LinkSubAccountsByWalletId(context.Background(), walletId).LinkSubAccountsByWalletIdRequest(linkSubAccountsByWalletIdRequest).Execute()
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Unique id of the wallet | 

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
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsExchangeWalletAPI.ListExchanges(context.Background()).Execute()
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
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	exchangeId := openapiclient.ExchangeId("binance") // ExchangeId | Exchange ID to query
	apikey := "d8f062da-39f4-4a11-8b9d-12595854237f" // string | The API Key for the exchange (optional)
	secret := "75B4F636193162488A3728B4A5797708" // string | The API Secret for the exchange. (optional)
	passphrase := "A3DBHJV" // string | The API passphrase for the exchange wallet. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsExchangeWalletAPI.ListSubAccountsByApikey(context.Background(), exchangeId).Apikey(apikey).Secret(secret).Passphrase(passphrase).Execute()
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the wallet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsExchangeWalletAPI.ListSubAccountsByWalletId(context.Background(), walletId).Execute()
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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Unique id of the wallet | 

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

