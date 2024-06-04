# \DevelopersWebhooksAPI

All URIs are relative to *https://api.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebhookEvent**](DevelopersWebhooksAPI.md#GetWebhookEvent) | **Get** /webhooks/events/{event_id} | Retrieve event by ID
[**GetWebhookEventLogs**](DevelopersWebhooksAPI.md#GetWebhookEventLogs) | **Get** /webhooks/events/{event_id}/logs | List event logs by ID
[**ListEvents**](DevelopersWebhooksAPI.md#ListEvents) | **Get** /webhooks/events | List all events
[**RetryWebhookEvent**](DevelopersWebhooksAPI.md#RetryWebhookEvent) | **Post** /webhooks/events/{event_id}/retry | Retry event by ID



## GetWebhookEvent

> WebhookEvent GetWebhookEvent(ctx, eventId).Execute()

Retrieve event by ID



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
	eventId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The event ID. You can obtain a list of event IDs by calling [List all events](/api-references/v2/developers--webhooks/list-all-events).

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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**eventId** | **string** | The event ID. You can obtain a list of event IDs by calling [List all events](/api-references/v2/developers--webhooks/list-all-events). | 

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

List event logs by ID



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
	eventId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The event ID. You can obtain a list of event IDs by calling [List all events](/api-references/v2/developers--webhooks/list-all-events).

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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**eventId** | **string** | The event ID. You can obtain a list of event IDs by calling [List all events](/api-references/v2/developers--webhooks/list-all-events). | 

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

List all events



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
	status := CoboWaas2.WebhookEventStatus("Success") // WebhookEventStatus | The event status. Possible values include: - `Success`: The event has been delivered, and the webhook endpoint has responded to the event. - `Retrying`: The event has been delivered, but the webhook endpoint has not responded. In this case, the WaaS service will retry delivering the event. - `Failed`: The event cannot be delivered and the WaaS service will stop retrying. This may occur if the number of retries reaches 10, or if the event has been delivered but the webhook endpoint responded with an error.  (optional)
	type_ := CoboWaas2.WebhookEventType("asset_wallet.outbound.created") // WebhookEventType | The event type.  (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The value range is [1, 50]. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `foo`, the request will retrieve a list of data objects that end before the object with the object ID `foo`. You can set this parameter to the value of `pagination.after` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `bar`, the request will retrieve a list of data objects that start after the object with the object ID `bar`. You can set this parameter to the value of `pagination.before` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)

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
 **status** | [**WebhookEventStatus**](WebhookEventStatus.md) | The event status. Possible values include: - &#x60;Success&#x60;: The event has been delivered, and the webhook endpoint has responded to the event. - &#x60;Retrying&#x60;: The event has been delivered, but the webhook endpoint has not responded. In this case, the WaaS service will retry delivering the event. - &#x60;Failed&#x60;: The event cannot be delivered and the WaaS service will stop retrying. This may occur if the number of retries reaches 10, or if the event has been delivered but the webhook endpoint responded with an error.  | 
 **type_** | [**WebhookEventType**](WebhookEventType.md) | The event type.  | 
 **limit** | **int32** | The maximum number of objects to return. The value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;foo&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;foo&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;bar&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;bar&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 

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


## RetryWebhookEvent

> RetryWebhookEvent201Response RetryWebhookEvent(ctx, eventId).Execute()

Retry event by ID



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
	eventId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The event ID. You can obtain a list of event IDs by calling [List all events](/api-references/v2/developers--webhooks/list-all-events).

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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**eventId** | **string** | The event ID. You can obtain a list of event IDs by calling [List all events](/api-references/v2/developers--webhooks/list-all-events). | 

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

