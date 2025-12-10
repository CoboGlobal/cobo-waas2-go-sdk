# \BatchPayoutsAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelBatchPayout**](BatchPayoutsAPI.md#CancelBatchPayout) | **Post** /batch_payouts/payouts/{payout_id}/cancel | Cancel batch payout
[**CreateBatchPayout**](BatchPayoutsAPI.md#CreateBatchPayout) | **Post** /batch_payouts/payouts | Create batch payout
[**DropBatchPayout**](BatchPayoutsAPI.md#DropBatchPayout) | **Post** /batch_payouts/payouts/{payout_id}/drop | Drop batch payout
[**EstimateBatchPayoutFee**](BatchPayoutsAPI.md#EstimateBatchPayoutFee) | **Post** /batch_payouts/estimate_fee | Estimate batch payout fee
[**GetBatchPayout**](BatchPayoutsAPI.md#GetBatchPayout) | **Get** /batch_payouts/payouts/{payout_id} | Get batch payout
[**ListBatchPayouts**](BatchPayoutsAPI.md#ListBatchPayouts) | **Get** /batch_payouts/payouts | List batch payouts
[**RevokeBatchPayout**](BatchPayoutsAPI.md#RevokeBatchPayout) | **Post** /batch_payouts/payouts/{payout_id}/revoke | Revoke batch payout
[**SpeedUpBatchPayout**](BatchPayoutsAPI.md#SpeedUpBatchPayout) | **Post** /batch_payouts/payouts/{payout_id}/speedup | Speed up batch payout



## CancelBatchPayout

> PayoutDetail CancelBatchPayout(ctx, payoutId).CancelPayoutBody(cancelPayoutBody).Execute()

Cancel batch payout



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
	payoutId := "e3986401-4aec-480a-973d-e775a4518413"
	cancelPayoutBody := *coboWaas2.NewCancelPayoutBody()

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
	resp, r, err := apiClient.BatchPayoutsAPI.CancelBatchPayout(ctx, payoutId).CancelPayoutBody(cancelPayoutBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPayoutsAPI.CancelBatchPayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelBatchPayout`: PayoutDetail
	fmt.Fprintf(os.Stdout, "Response from `BatchPayoutsAPI.CancelBatchPayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**payoutId** | **string** | The payout ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelBatchPayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelPayoutBody** | [**CancelPayoutBody**](CancelPayoutBody.md) | The request body to cancel a batch payout. | 

### Return type

[**PayoutDetail**](PayoutDetail.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBatchPayout

> CreateBatchPayout201Response CreateBatchPayout(ctx).CreateBatchPayoutRequest(createBatchPayoutRequest).Execute()

Create batch payout



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
	createBatchPayoutRequest := *coboWaas2.NewCreateBatchPayoutRequest("Monthly vendor payments - January 2024", "ETH_USDT", coboWaas2.PayoutMode("Normal"), *coboWaas2.NewPayoutSource("123e4567-e89b-12d3-a456-426614174003"), []coboWaas2.PayoutDestination{*coboWaas2.NewPayoutDestination("0xabc123def456ghi789jkl012mno345pqr678stu901", "1.500000")})

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
	resp, r, err := apiClient.BatchPayoutsAPI.CreateBatchPayout(ctx).CreateBatchPayoutRequest(createBatchPayoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPayoutsAPI.CreateBatchPayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBatchPayout`: CreateBatchPayout201Response
	fmt.Fprintf(os.Stdout, "Response from `BatchPayoutsAPI.CreateBatchPayout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchPayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBatchPayoutRequest** | [**CreateBatchPayoutRequest**](CreateBatchPayoutRequest.md) | The request body to create a batch payout. | 

### Return type

[**CreateBatchPayout201Response**](CreateBatchPayout201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DropBatchPayout

> PayoutDetail DropBatchPayout(ctx, payoutId).PayoutRbfBody(payoutRbfBody).Execute()

Drop batch payout



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
	payoutId := "e3986401-4aec-480a-973d-e775a4518413"
	payoutRbfBody := *coboWaas2.NewPayoutRbfBody()

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
	resp, r, err := apiClient.BatchPayoutsAPI.DropBatchPayout(ctx, payoutId).PayoutRbfBody(payoutRbfBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPayoutsAPI.DropBatchPayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DropBatchPayout`: PayoutDetail
	fmt.Fprintf(os.Stdout, "Response from `BatchPayoutsAPI.DropBatchPayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**payoutId** | **string** | The payout ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDropBatchPayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payoutRbfBody** | [**PayoutRbfBody**](PayoutRbfBody.md) | The request body to speed up or drop a batch payout. | 

### Return type

[**PayoutDetail**](PayoutDetail.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EstimateBatchPayoutFee

> []PayoutEstimatedFee EstimateBatchPayoutFee(ctx).EstimateBatchPayoutFeeRequest(estimateBatchPayoutFeeRequest).Execute()

Estimate batch payout fee



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
	estimateBatchPayoutFeeRequest := *coboWaas2.NewEstimateBatchPayoutFeeRequest("ETH_USDT", coboWaas2.PayoutMode("Normal"), *coboWaas2.NewPayoutSource("123e4567-e89b-12d3-a456-426614174003"), []coboWaas2.PayoutDestination{*coboWaas2.NewPayoutDestination("0xabc123def456ghi789jkl012mno345pqr678stu901", "1.500000")})

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
	resp, r, err := apiClient.BatchPayoutsAPI.EstimateBatchPayoutFee(ctx).EstimateBatchPayoutFeeRequest(estimateBatchPayoutFeeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPayoutsAPI.EstimateBatchPayoutFee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EstimateBatchPayoutFee`: []PayoutEstimatedFee
	fmt.Fprintf(os.Stdout, "Response from `BatchPayoutsAPI.EstimateBatchPayoutFee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEstimateBatchPayoutFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **estimateBatchPayoutFeeRequest** | [**EstimateBatchPayoutFeeRequest**](EstimateBatchPayoutFeeRequest.md) | The request body to estimate the fee of a batch payout. | 

### Return type

[**[]PayoutEstimatedFee**](PayoutEstimatedFee.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchPayout

> PayoutDetail GetBatchPayout(ctx, payoutId).Execute()

Get batch payout



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
	payoutId := "e3986401-4aec-480a-973d-e775a4518413"

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
	resp, r, err := apiClient.BatchPayoutsAPI.GetBatchPayout(ctx, payoutId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPayoutsAPI.GetBatchPayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchPayout`: PayoutDetail
	fmt.Fprintf(os.Stdout, "Response from `BatchPayoutsAPI.GetBatchPayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**payoutId** | **string** | The payout ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchPayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PayoutDetail**](PayoutDetail.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBatchPayouts

> ListBatchPayouts200Response ListBatchPayouts(ctx).Limit(limit).Before(before).After(after).Execute()

List batch payouts



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
	resp, r, err := apiClient.BatchPayoutsAPI.ListBatchPayouts(ctx).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPayoutsAPI.ListBatchPayouts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBatchPayouts`: ListBatchPayouts200Response
	fmt.Fprintf(os.Stdout, "Response from `BatchPayoutsAPI.ListBatchPayouts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBatchPayoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | This parameter specifies an object ID as a starting point for pagination, retrieving data before the specified object relative to the current dataset.    Suppose the current data is ordered as Object A, Object B, and Object C.  If you set &#x60;before&#x60; to the ID of Object C (&#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;), the response will include Object B and Object A.    **Notes**:   - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur. - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned. - If you set it to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | This parameter specifies an object ID as a starting point for pagination, retrieving data after the specified object relative to the current dataset.    Suppose the current data is ordered as Object A, Object B, and Object C. If you set &#x60;after&#x60; to the ID of Object A (&#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;), the response will include Object B and Object C.    **Notes**:   - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur. - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 

### Return type

[**ListBatchPayouts200Response**](ListBatchPayouts200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeBatchPayout

> PayoutDetail RevokeBatchPayout(ctx, payoutId).Execute()

Revoke batch payout



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
	payoutId := "e3986401-4aec-480a-973d-e775a4518413"

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
	resp, r, err := apiClient.BatchPayoutsAPI.RevokeBatchPayout(ctx, payoutId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPayoutsAPI.RevokeBatchPayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeBatchPayout`: PayoutDetail
	fmt.Fprintf(os.Stdout, "Response from `BatchPayoutsAPI.RevokeBatchPayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**payoutId** | **string** | The payout ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeBatchPayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PayoutDetail**](PayoutDetail.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpeedUpBatchPayout

> PayoutDetail SpeedUpBatchPayout(ctx, payoutId).PayoutRbfBody(payoutRbfBody).Execute()

Speed up batch payout



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
	payoutId := "e3986401-4aec-480a-973d-e775a4518413"
	payoutRbfBody := *coboWaas2.NewPayoutRbfBody()

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
	resp, r, err := apiClient.BatchPayoutsAPI.SpeedUpBatchPayout(ctx, payoutId).PayoutRbfBody(payoutRbfBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchPayoutsAPI.SpeedUpBatchPayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpeedUpBatchPayout`: PayoutDetail
	fmt.Fprintf(os.Stdout, "Response from `BatchPayoutsAPI.SpeedUpBatchPayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**payoutId** | **string** | The payout ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpeedUpBatchPayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payoutRbfBody** | [**PayoutRbfBody**](PayoutRbfBody.md) | The request body to speed up or drop a batch payout. | 

### Return type

[**PayoutDetail**](PayoutDetail.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

