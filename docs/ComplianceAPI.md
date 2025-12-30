# \ComplianceAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKyaScreenings**](ComplianceAPI.md#CreateKyaScreenings) | **Post** /compliance/kya/screenings | Create KYA address screening requests
[**GetDispositionStatus**](ComplianceAPI.md#GetDispositionStatus) | **Get** /compliance/funds/disposition | Get disposition status
[**GetKyaScreening**](ComplianceAPI.md#GetKyaScreening) | **Get** /compliance/kya/screenings/{screening_id} | Get KYA address screening result
[**GetKytScreeningStatus**](ComplianceAPI.md#GetKytScreeningStatus) | **Get** /compliance/kyt/screenings/status | Get KYT screening status
[**IsolateFunds**](ComplianceAPI.md#IsolateFunds) | **Post** /compliance/funds/disposition/isolate | Isolate funds
[**ListKyaScreenings**](ComplianceAPI.md#ListKyaScreenings) | **Get** /compliance/kya/screenings | List KYA address screening results
[**RefundFunds**](ComplianceAPI.md#RefundFunds) | **Post** /compliance/funds/disposition/refund | Refund funds
[**SubmitKytManualReview**](ComplianceAPI.md#SubmitKytManualReview) | **Post** /compliance/kyt/screenings/manual_review | Submit KYT manual review result
[**SubmitKytScreeningDecisions**](ComplianceAPI.md#SubmitKytScreeningDecisions) | **Post** /compliance/kyt/screenings/decisions | Submit KYT screening decision
[**UnfreezeFunds**](ComplianceAPI.md#UnfreezeFunds) | **Post** /compliance/funds/disposition/unfreeze | Unfreeze frozen funds



## CreateKyaScreenings

> []KyaScreeningResult CreateKyaScreenings(ctx).CreateKyaScreeningsBody(createKyaScreeningsBody).Execute()

Create KYA address screening requests



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
	createKyaScreeningsBody := *coboWaas2.NewCreateKyaScreeningsBody([]coboWaas2.KyaScreeningRequest{*coboWaas2.NewKyaScreeningRequest("scan_eth_addr1_20240721", "0x1234567890abcdef1234567890abcdef12345678", "ETH")})

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
	resp, r, err := apiClient.ComplianceAPI.CreateKyaScreenings(ctx).CreateKyaScreeningsBody(createKyaScreeningsBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CreateKyaScreenings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKyaScreenings`: []KyaScreeningResult
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CreateKyaScreenings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKyaScreeningsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKyaScreeningsBody** | [**CreateKyaScreeningsBody**](CreateKyaScreeningsBody.md) | The request body to create KYA address screening requests. | 

### Return type

[**[]KyaScreeningResult**](KyaScreeningResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDispositionStatus

> DispositionQueryResponse GetDispositionStatus(ctx).TransactionId(transactionId).Execute()

Get disposition status



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
	resp, r, err := apiClient.ComplianceAPI.GetDispositionStatus(ctx).TransactionId(transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetDispositionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDispositionStatus`: DispositionQueryResponse
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetDispositionStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDispositionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **string** | The unique identifier (UUID) of the transaction to retrieve KYT screening status information for. | 

### Return type

[**DispositionQueryResponse**](DispositionQueryResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKyaScreening

> KyaScreeningResult GetKyaScreening(ctx, screeningId).Execute()

Get KYA address screening result



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
	screeningId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"

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
	resp, r, err := apiClient.ComplianceAPI.GetKyaScreening(ctx, screeningId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetKyaScreening``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKyaScreening`: KyaScreeningResult
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetKyaScreening`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**screeningId** | **string** | The unique identifier (UUID) of the address screening request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKyaScreeningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KyaScreeningResult**](KyaScreeningResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKytScreeningStatus

> KytScreeningsTransaction GetKytScreeningStatus(ctx).TransactionId(transactionId).Execute()

Get KYT screening status



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
	resp, r, err := apiClient.ComplianceAPI.GetKytScreeningStatus(ctx).TransactionId(transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetKytScreeningStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKytScreeningStatus`: KytScreeningsTransaction
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetKytScreeningStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKytScreeningStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **string** | The unique identifier (UUID) of the transaction to retrieve KYT screening status information for. | 

### Return type

[**KytScreeningsTransaction**](KytScreeningsTransaction.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsolateFunds

> DispositionResponse IsolateFunds(ctx).IsolateDisposition(isolateDisposition).Execute()

Isolate funds



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
	isolateDisposition := *coboWaas2.NewIsolateDisposition("f47ac10b-58cc-4372-a567-0e02b2c3d479", "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7", "1.5")

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
	resp, r, err := apiClient.ComplianceAPI.IsolateFunds(ctx).IsolateDisposition(isolateDisposition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.IsolateFunds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsolateFunds`: DispositionResponse
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.IsolateFunds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIsolateFundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isolateDisposition** | [**IsolateDisposition**](IsolateDisposition.md) | The request body to isolate funds. | 

### Return type

[**DispositionResponse**](DispositionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKyaScreenings

> ListKyaScreenings200Response ListKyaScreenings(ctx).ScreeningIds(screeningIds).Limit(limit).Before(before).After(after).Execute()

List KYA address screening results



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
	screeningIds := "f47ac10b-58cc-4372-a567-0e02b2c3d479,a1b2c3d4-e5f6-4321-8765-fedcba987654"
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
	resp, r, err := apiClient.ComplianceAPI.ListKyaScreenings(ctx).ScreeningIds(screeningIds).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.ListKyaScreenings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKyaScreenings`: ListKyaScreenings200Response
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.ListKyaScreenings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKyaScreeningsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **screeningIds** | **string** | A comma-separated list of screening request IDs to filter specific screening results. Maximum 50 IDs allowed to ensure URL length stays within standard web server limits (typically 8KB).  Each ID must be in standard UUID format (36 characters fixed length).  Example: &#x60;f47ac10b-58cc-4372-a567-0e02b2c3d479,a1b2c3d4-e5f6-4321-8765-fedcba987654&#x60;  | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 

### Return type

[**ListKyaScreenings200Response**](ListKyaScreenings200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundFunds

> DispositionResponse RefundFunds(ctx).RefundDisposition(refundDisposition).Execute()

Refund funds



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
	refundDisposition := *coboWaas2.NewRefundDisposition("f47ac10b-58cc-4372-a567-0e02b2c3d479", "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7", "1.5")

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
	resp, r, err := apiClient.ComplianceAPI.RefundFunds(ctx).RefundDisposition(refundDisposition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.RefundFunds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundFunds`: DispositionResponse
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.RefundFunds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundFundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundDisposition** | [**RefundDisposition**](RefundDisposition.md) | The request body to refund funds. | 

### Return type

[**DispositionResponse**](DispositionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitKytManualReview

> SubmitKytResponse SubmitKytManualReview(ctx).SubmitKytScreeningsReviewBody(submitKytScreeningsReviewBody).Execute()

Submit KYT manual review result



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
	submitKytScreeningsReviewBody := *coboWaas2.NewSubmitKytScreeningsReviewBody("f47ac10b-58cc-4372-a567-0e02b2c3d479", coboWaas2.KytScreeningsReviewType("Approval"))

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
	resp, r, err := apiClient.ComplianceAPI.SubmitKytManualReview(ctx).SubmitKytScreeningsReviewBody(submitKytScreeningsReviewBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.SubmitKytManualReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitKytManualReview`: SubmitKytResponse
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.SubmitKytManualReview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitKytManualReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitKytScreeningsReviewBody** | [**SubmitKytScreeningsReviewBody**](SubmitKytScreeningsReviewBody.md) | The request body to submit a manual review result for a KYT screening case that requires human analysis. | 

### Return type

[**SubmitKytResponse**](SubmitKytResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitKytScreeningDecisions

> SubmitKytResponse SubmitKytScreeningDecisions(ctx).SubmitKytScreeningsDecisionsBody(submitKytScreeningsDecisionsBody).Execute()

Submit KYT screening decision



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
	submitKytScreeningsDecisionsBody := *coboWaas2.NewSubmitKytScreeningsDecisionsBody("f47ac10b-58cc-4372-a567-0e02b2c3d479", coboWaas2.KytScreeningsDecisionsType("Approval"))

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
	resp, r, err := apiClient.ComplianceAPI.SubmitKytScreeningDecisions(ctx).SubmitKytScreeningsDecisionsBody(submitKytScreeningsDecisionsBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.SubmitKytScreeningDecisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitKytScreeningDecisions`: SubmitKytResponse
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.SubmitKytScreeningDecisions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitKytScreeningDecisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitKytScreeningsDecisionsBody** | [**SubmitKytScreeningsDecisionsBody**](SubmitKytScreeningsDecisionsBody.md) | The request body to submit a KYT screening decision for a specific transaction based on external compliance review. | 

### Return type

[**SubmitKytResponse**](SubmitKytResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnfreezeFunds

> DispositionResponse UnfreezeFunds(ctx).UnfreezeDisposition(unfreezeDisposition).Execute()

Unfreeze frozen funds



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
	unfreezeDisposition := *coboWaas2.NewUnfreezeDisposition("f47ac10b-58cc-4372-a567-0e02b2c3d479")

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
	resp, r, err := apiClient.ComplianceAPI.UnfreezeFunds(ctx).UnfreezeDisposition(unfreezeDisposition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.UnfreezeFunds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnfreezeFunds`: DispositionResponse
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.UnfreezeFunds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnfreezeFundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unfreezeDisposition** | [**UnfreezeDisposition**](UnfreezeDisposition.md) | The request body to unfreeze funds. | 

### Return type

[**DispositionResponse**](DispositionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

