# \SwapsAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSwapActivity**](SwapsAPI.md#CreateSwapActivity) | **Post** /swaps/swap | Create Swap Activity
[**CreateSwapQuote**](SwapsAPI.md#CreateSwapQuote) | **Post** /swaps/quote | Create Swap Quote
[**GetSwapActivity**](SwapsAPI.md#GetSwapActivity) | **Get** /swaps/activities/{activity_id} | Get Swap Activity Details
[**GetSwapQuote**](SwapsAPI.md#GetSwapQuote) | **Get** /swaps/quote | Get Current Swap Rate
[**ListEnableTokenPairs**](SwapsAPI.md#ListEnableTokenPairs) | **Get** /swaps/enabled_pairs | List Supported Token Pairs
[**ListSwapActivities**](SwapsAPI.md#ListSwapActivities) | **Get** /swaps/activities | List Swap Activities



## CreateSwapActivity

> SwapActivity CreateSwapActivity(ctx).CreateSwapActivityRequest(createSwapActivityRequest).Execute()

Create Swap Activity



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
	createSwapActivityRequest := *coboWaas2.NewCreateSwapActivityRequest("123e4567-e89b-12d3-a456-426614174000", "123e4567-e89b-12d3-a456-426614174001", "0.005")

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
	resp, r, err := apiClient.SwapsAPI.CreateSwapActivity(ctx).CreateSwapActivityRequest(createSwapActivityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapsAPI.CreateSwapActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSwapActivity`: SwapActivity
	fmt.Fprintf(os.Stdout, "Response from `SwapsAPI.CreateSwapActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSwapActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSwapActivityRequest** | [**CreateSwapActivityRequest**](CreateSwapActivityRequest.md) | The request body for creating a swap activity. | 

### Return type

[**SwapActivity**](SwapActivity.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSwapQuote

> CreateSwapQuote201Response CreateSwapQuote(ctx).CreateSwapQuoteRequest(createSwapQuoteRequest).Execute()

Create Swap Quote



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
	createSwapQuoteRequest := *coboWaas2.NewCreateSwapQuoteRequest("123e4567-e89b-12d3-a456-426614174000", "BTC", "ETH_WBTC")

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
	resp, r, err := apiClient.SwapsAPI.CreateSwapQuote(ctx).CreateSwapQuoteRequest(createSwapQuoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapsAPI.CreateSwapQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSwapQuote`: CreateSwapQuote201Response
	fmt.Fprintf(os.Stdout, "Response from `SwapsAPI.CreateSwapQuote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSwapQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSwapQuoteRequest** | [**CreateSwapQuoteRequest**](CreateSwapQuoteRequest.md) | The request body for creating a swap activity. | 

### Return type

[**CreateSwapQuote201Response**](CreateSwapQuote201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwapActivity

> SwapActivity GetSwapActivity(ctx, activityId).Execute()

Get Swap Activity Details



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
	activityId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"

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
	resp, r, err := apiClient.SwapsAPI.GetSwapActivity(ctx, activityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapsAPI.GetSwapActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwapActivity`: SwapActivity
	fmt.Fprintf(os.Stdout, "Response from `SwapsAPI.GetSwapActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**activityId** | **string** | The unique id of the activity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwapActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SwapActivity**](SwapActivity.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwapQuote

> SwapQuote GetSwapQuote(ctx).WalletId(walletId).PayTokenId(payTokenId).ReceiveTokenId(receiveTokenId).PayAmount(payAmount).ReceiveAmount(receiveAmount).Execute()

Get Current Swap Rate



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
	payTokenId := "ETH"
	receiveTokenId := "USDT"
	payAmount := "1.5"
	receiveAmount := "2000"

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
	resp, r, err := apiClient.SwapsAPI.GetSwapQuote(ctx).WalletId(walletId).PayTokenId(payTokenId).ReceiveTokenId(receiveTokenId).PayAmount(payAmount).ReceiveAmount(receiveAmount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapsAPI.GetSwapQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwapQuote`: SwapQuote
	fmt.Fprintf(os.Stdout, "Response from `SwapsAPI.GetSwapQuote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSwapQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletId** | **string** | The wallet ID. | 
 **payTokenId** | **string** | Unique id of the token to pay. | 
 **receiveTokenId** | **string** | Unique id of the token to receive. | 
 **payAmount** | **string** | The amount of pay token to swap. | 
 **receiveAmount** | **string** | The amount of token to receive. | 

### Return type

[**SwapQuote**](SwapQuote.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnableTokenPairs

> ListEnableTokenPairs200Response ListEnableTokenPairs(ctx).Limit(limit).Before(before).After(after).Execute()

List Supported Token Pairs



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
	resp, r, err := apiClient.SwapsAPI.ListEnableTokenPairs(ctx).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapsAPI.ListEnableTokenPairs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnableTokenPairs`: ListEnableTokenPairs200Response
	fmt.Fprintf(os.Stdout, "Response from `SwapsAPI.ListEnableTokenPairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEnableTokenPairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | This parameter specifies an object ID as a starting point for pagination, retrieving data before the specified object relative to the current dataset.    Suppose the current data is ordered as Object A, Object B, and Object C.  If you set &#x60;before&#x60; to the ID of Object C (&#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;), the response will include Object B and Object A.    **Notes**:   - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur. - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned. - If you set it to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | This parameter specifies an object ID as a starting point for pagination, retrieving data after the specified object relative to the current dataset.    Suppose the current data is ordered as Object A, Object B, and Object C. If you set &#x60;after&#x60; to the ID of Object A (&#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;), the response will include Object B and Object C.    **Notes**:   - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur. - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 

### Return type

[**ListEnableTokenPairs200Response**](ListEnableTokenPairs200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSwapActivities

> ListSwapActivities200Response ListSwapActivities(ctx).Status(status).MinUpdatedTimestamp(minUpdatedTimestamp).MaxUpdatedTimestamp(maxUpdatedTimestamp).Initiator(initiator).Limit(limit).Before(before).After(after).SortBy(sortBy).Direction(direction).Execute()

List Swap Activities



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
	status := "Success"
	minUpdatedTimestamp := int64(1635744000000)
	maxUpdatedTimestamp := int64(1635744000000)
	initiator := "steve@example.com"
	limit := int32(10)
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1"
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk"
	sortBy := "timestamp"
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
	resp, r, err := apiClient.SwapsAPI.ListSwapActivities(ctx).Status(status).MinUpdatedTimestamp(minUpdatedTimestamp).MaxUpdatedTimestamp(maxUpdatedTimestamp).Initiator(initiator).Limit(limit).Before(before).After(after).SortBy(sortBy).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapsAPI.ListSwapActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSwapActivities`: ListSwapActivities200Response
	fmt.Fprintf(os.Stdout, "Response from `SwapsAPI.ListSwapActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSwapActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **minUpdatedTimestamp** | **int64** | The start time of the query. All staking activities updated after the specified time will be retrieved. The time is in Unix timestamp format, measured in milliseconds. | 
 **maxUpdatedTimestamp** | **int64** | The end time of the query. All staking activities updated before the specified time will be retrieved. The time is in Unix timestamp format, measured in milliseconds. | 
 **initiator** | **string** | The activity initiator, which is your API key by default. You can also specify the initiator when creating the activity. | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | This parameter specifies an object ID as a starting point for pagination, retrieving data before the specified object relative to the current dataset.    Suppose the current data is ordered as Object A, Object B, and Object C.  If you set &#x60;before&#x60; to the ID of Object C (&#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;), the response will include Object B and Object A.    **Notes**:   - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur. - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned. - If you set it to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | This parameter specifies an object ID as a starting point for pagination, retrieving data after the specified object relative to the current dataset.    Suppose the current data is ordered as Object A, Object B, and Object C. If you set &#x60;after&#x60; to the ID of Object A (&#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;), the response will include Object B and Object C.    **Notes**:   - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur. - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 
 **sortBy** | **string** | The field used for sorting. | [default to &quot;&quot;]
 **direction** | **string** | The sort direction. Possible values include:   - &#x60;ASC&#x60;: Sort the results in ascending order.   - &#x60;DESC&#x60;: Sort the results in descending order.  | [default to &quot;&quot;]

### Return type

[**ListSwapActivities200Response**](ListSwapActivities200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

