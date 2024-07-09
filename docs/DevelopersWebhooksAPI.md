# \DevelopersWebhooksAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookEndpoint**](DevelopersWebhooksAPI.md#CreateWebhookEndpoint) | **Post** /webhooks/endpoints | Register webhook endpoint
[**GetWebhookEndpointById**](DevelopersWebhooksAPI.md#GetWebhookEndpointById) | **Get** /webhooks/endpoints/{endpoint_id} | Get webhook endpoint information
[**GetWebhookEvent**](DevelopersWebhooksAPI.md#GetWebhookEvent) | **Get** /webhooks/endpoints/{endpoint_id}/events/{event_id} | Retrieve event information
[**GetWebhookEventLogs**](DevelopersWebhooksAPI.md#GetWebhookEventLogs) | **Get** /webhooks/endpoints/{endpoint_id}/events/{event_id}/logs | List event logs
[**ListEvents**](DevelopersWebhooksAPI.md#ListEvents) | **Get** /webhooks/endpoints/{endpoint_id}/events | List all events
[**ListWebhookEndpoints**](DevelopersWebhooksAPI.md#ListWebhookEndpoints) | **Get** /webhooks/endpoints | List webhook endpoints
[**ListWebhookEventDefinitions**](DevelopersWebhooksAPI.md#ListWebhookEventDefinitions) | **Get** /webhooks/events/definitions | Get webhook event types
[**RetryWebhookEvent**](DevelopersWebhooksAPI.md#RetryWebhookEvent) | **Post** /webhooks/endpoints/{endpoint_id}/events/{event_id}/retry | Retry event
[**UpdateWebhookEndpoint**](DevelopersWebhooksAPI.md#UpdateWebhookEndpoint) | **Put** /webhooks/endpoints/{endpoint_id} | Update webhook endpoint



## CreateWebhookEndpoint

> Endpoint CreateWebhookEndpoint(ctx).CreateWebhookEndpointRequest(createWebhookEndpointRequest).Execute()

Register webhook endpoint



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
	createWebhookEndpointRequest := *CoboWaas2.NewCreateWebhookEndpointRequest("https://example.com/webhook", []CoboWaas2.WebhookEventType{CoboWaas2.WebhookEventType("wallets.transaction.created")}) // CreateWebhookEndpointRequest | The request body to register a webhook endpoint. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.CreateWebhookEndpoint(ctx).CreateWebhookEndpointRequest(createWebhookEndpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.CreateWebhookEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhookEndpoint`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.CreateWebhookEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebhookEndpointRequest** | [**CreateWebhookEndpointRequest**](CreateWebhookEndpointRequest.md) | The request body to register a webhook endpoint. | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookEndpointById

> Endpoint GetWebhookEndpointById(ctx, endpointId).Execute()

Get webhook endpoint information



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
	endpointId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints).

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.GetWebhookEndpointById(ctx, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.GetWebhookEndpointById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookEndpointById`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.GetWebhookEndpointById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**endpointId** | **string** | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookEndpointByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookEvent

> WebhookEvent GetWebhookEvent(ctx, eventId, endpointId).Execute()

Retrieve event information



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
	eventId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The event ID. You can obtain a list of event IDs by calling [List all events](/v2/api-references/developers--webhooks/list-all-events).
	endpointId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints).

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.GetWebhookEvent(ctx, eventId, endpointId).Execute()
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
**eventId** | **string** | The event ID. You can obtain a list of event IDs by calling [List all events](/v2/api-references/developers--webhooks/list-all-events). | 
**endpointId** | **string** | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints). | 

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

> []WebhookEventLog GetWebhookEventLogs(ctx, eventId, endpointId).Execute()

List event logs



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
	eventId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The event ID. You can obtain a list of event IDs by calling [List all events](/v2/api-references/developers--webhooks/list-all-events).
	endpointId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints).

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.GetWebhookEventLogs(ctx, eventId, endpointId).Execute()
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
**eventId** | **string** | The event ID. You can obtain a list of event IDs by calling [List all events](/v2/api-references/developers--webhooks/list-all-events). | 
**endpointId** | **string** | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints). | 

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

> ListEvents200Response ListEvents(ctx, endpointId).Status(status).Type_(type_).Limit(limit).Before(before).After(after).Execute()

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
	endpointId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints).
	status := CoboWaas2.WebhookEventStatus("Success") // WebhookEventStatus |  (optional)
	type_ := CoboWaas2.WebhookEventType("wallets.transaction.created") // WebhookEventType |  (optional)
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
	resp, r, err := apiClient.DevelopersWebhooksAPI.ListEvents(ctx, endpointId).Status(status).Type_(type_).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.ListEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEvents`: ListEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.ListEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**endpointId** | **string** | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints). | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**WebhookEventStatus**](WebhookEventStatus.md) |  | 
 **type_** | [**WebhookEventType**](WebhookEventType.md) |  | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

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


## ListWebhookEndpoints

> ListWebhookEndpoints200Response ListWebhookEndpoints(ctx).Status(status).EventType(eventType).Limit(limit).Before(before).After(after).Execute()

List webhook endpoints



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
	status := CoboWaas2.EndpointStatus("STATUS_ACTIVE") // EndpointStatus |  (optional)
	eventType := CoboWaas2.WebhookEventType("wallets.transaction.created") // WebhookEventType |  (optional)
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
	resp, r, err := apiClient.DevelopersWebhooksAPI.ListWebhookEndpoints(ctx).Status(status).EventType(eventType).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.ListWebhookEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhookEndpoints`: ListWebhookEndpoints200Response
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.ListWebhookEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**EndpointStatus**](EndpointStatus.md) |  | 
 **eventType** | [**WebhookEventType**](WebhookEventType.md) |  | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**ListWebhookEndpoints200Response**](ListWebhookEndpoints200Response.md)

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

Get webhook event types



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

> RetryWebhookEvent201Response RetryWebhookEvent(ctx, eventId, endpointId).Execute()

Retry event



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
	eventId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The event ID. You can obtain a list of event IDs by calling [List all events](/v2/api-references/developers--webhooks/list-all-events).
	endpointId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints).

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.RetryWebhookEvent(ctx, eventId, endpointId).Execute()
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
**eventId** | **string** | The event ID. You can obtain a list of event IDs by calling [List all events](/v2/api-references/developers--webhooks/list-all-events). | 
**endpointId** | **string** | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints). | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryWebhookEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RetryWebhookEvent201Response**](RetryWebhookEvent201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookEndpoint

> Endpoint UpdateWebhookEndpoint(ctx, endpointId).UpdateWebhookEndpointRequest(updateWebhookEndpointRequest).Execute()

Update webhook endpoint



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
	endpointId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints).
	updateWebhookEndpointRequest := *CoboWaas2.NewUpdateWebhookEndpointRequest() // UpdateWebhookEndpointRequest | The request body to update a webhook endpoint. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.UpdateWebhookEndpoint(ctx, endpointId).UpdateWebhookEndpointRequest(updateWebhookEndpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.UpdateWebhookEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookEndpoint`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.UpdateWebhookEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**endpointId** | **string** | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWebhookEndpointRequest** | [**UpdateWebhookEndpointRequest**](UpdateWebhookEndpointRequest.md) | The request body to update a webhook endpoint. | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

