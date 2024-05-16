# \DevelopersWebhooksAPI

All URIs are relative to *https://api.cobo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebhookEvent**](DevelopersWebhooksAPI.md#GetWebhookEvent) | **Get** /webhooks/events/{event_id} | Retrieve webhook event information by event ID.
[**GetWebhookEventLogs**](DevelopersWebhooksAPI.md#GetWebhookEventLogs) | **Get** /webhooks/events/{event_id}/logs | List webhook event logs by event ID.
[**ListEvents**](DevelopersWebhooksAPI.md#ListEvents) | **Get** /webhooks/events | List triggered events.
[**ListWebhookEventDefinitions**](DevelopersWebhooksAPI.md#ListWebhookEventDefinitions) | **Get** /webhooks/events/definitions | List all supported event definitions.
[**RetryWebhookEvent**](DevelopersWebhooksAPI.md#RetryWebhookEvent) | **Post** /webhooks/events/{event_id}/retry | Retry webhook event by event ID.



## GetWebhookEvent

> WebhookEvent GetWebhookEvent(ctx, eventId).Execute()

Retrieve webhook event information by event ID.



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
	eventId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the webhook event, get event IDs by calling `List triggered events`.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.GetWebhookEvent(ctx, eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.GetWebhookEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookEvent`: WebhookEvent
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.GetWebhookEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | Unique id of the webhook event, get event IDs by calling &#x60;List triggered events&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookEvent**](WebhookEvent.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookEventLogs

> []WebhookEventLog GetWebhookEventLogs(ctx, eventId).Execute()

List webhook event logs by event ID.



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
	eventId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the webhook event, get event IDs by calling `List triggered events`.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.GetWebhookEventLogs(ctx, eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.GetWebhookEventLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookEventLogs`: []WebhookEventLog
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.GetWebhookEventLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | Unique id of the webhook event, get event IDs by calling &#x60;List triggered events&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookEventLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]WebhookEventLog**](WebhookEventLog.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> ListEvents200Response ListEvents(ctx).Status(status).Type_(type_).Limit(limit).Before(before).After(after).Execute()

List triggered events.



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
	status := CoboWaas2.WebhookEventStatus("Success") // WebhookEventStatus | The status of event. (optional)
	type_ := CoboWaas2.WebhookEventType("asset_wallet.outbound.created") // WebhookEventType | The type of event. Get event types by calling `List all supported event definitions`.  (optional)
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
	resp, r, err := apiClient.DevelopersWebhooksAPI.ListEvents(ctx).Status(status).Type_(type_).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.ListEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEvents`: ListEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.ListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**WebhookEventStatus**](WebhookEventStatus.md) | The status of event. | 
 **type_** | [**WebhookEventType**](WebhookEventType.md) | The type of event. Get event types by calling &#x60;List all supported event definitions&#x60;.  | 
 **limit** | **int32** | size of page to return (pagination) | [default to 10]
 **before** | **string** | Cursor string received from previous request | [default to &quot;&quot;]
 **after** | **string** | Cursor string received from previous request | [default to &quot;&quot;]

### Return type

[**ListEvents200Response**](ListEvents200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookEventDefinitions

> []ListWebhookEventDefinitions200ResponseInner ListWebhookEventDefinitions(ctx).Execute()

List all supported event definitions.



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
	resp, r, err := apiClient.DevelopersWebhooksAPI.ListWebhookEventDefinitions(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.ListWebhookEventDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhookEventDefinitions`: []ListWebhookEventDefinitions200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.ListWebhookEventDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookEventDefinitionsRequest struct via the builder pattern


### Return type

[**[]ListWebhookEventDefinitions200ResponseInner**](ListWebhookEventDefinitions200ResponseInner.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryWebhookEvent

> RetryWebhookEvent201Response RetryWebhookEvent(ctx, eventId).Execute()

Retry webhook event by event ID.



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
	eventId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the webhook event, get event IDs by calling `List triggered events`.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.RetryWebhookEvent(ctx, eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.RetryWebhookEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryWebhookEvent`: RetryWebhookEvent201Response
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.RetryWebhookEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | Unique id of the webhook event, get event IDs by calling &#x60;List triggered events&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryWebhookEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetryWebhookEvent201Response**](RetryWebhookEvent201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

