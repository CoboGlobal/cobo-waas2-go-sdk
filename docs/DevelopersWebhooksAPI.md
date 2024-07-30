# \DevelopersWebhooksAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookEndpoint**](DevelopersWebhooksAPI.md#CreateWebhookEndpoint) | **Post** /webhooks/endpoints | Register webhook endpoint
[**GetWebhookEndpointById**](DevelopersWebhooksAPI.md#GetWebhookEndpointById) | **Get** /webhooks/endpoints/{endpoint_id} | Get webhook endpoint information
[**GetWebhookEventById**](DevelopersWebhooksAPI.md#GetWebhookEventById) | **Get** /webhooks/endpoints/{endpoint_id}/events/{event_id} | Retrieve event information
[**ListWebhookEndpoints**](DevelopersWebhooksAPI.md#ListWebhookEndpoints) | **Get** /webhooks/endpoints | List webhook endpoints
[**ListWebhookEventDefinitions**](DevelopersWebhooksAPI.md#ListWebhookEventDefinitions) | **Get** /webhooks/events/definitions | Get webhook event types
[**ListWebhookEventLogs**](DevelopersWebhooksAPI.md#ListWebhookEventLogs) | **Get** /webhooks/endpoints/{endpoint_id}/events/{event_id}/logs | List webhook event logs
[**ListWebhookEvents**](DevelopersWebhooksAPI.md#ListWebhookEvents) | **Get** /webhooks/endpoints/{endpoint_id}/events | List all webhook events
[**RetryWebhookEventById**](DevelopersWebhooksAPI.md#RetryWebhookEventById) | **Post** /webhooks/endpoints/{endpoint_id}/events/{event_id}/retry | Retry event
[**UpdateWebhookEndpointById**](DevelopersWebhooksAPI.md#UpdateWebhookEndpointById) | **Put** /webhooks/endpoints/{endpoint_id} | Update webhook endpoint



## CreateWebhookEndpoint

> WebhookEndpoint CreateWebhookEndpoint(ctx).CreateWebhookEndpointRequest(createWebhookEndpointRequest).Execute()

Register webhook endpoint



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
	createWebhookEndpointRequest := *coboWaas2.NewCreateWebhookEndpointRequest("https://example.com/webhook", []coboWaas2.WebhookEventType{coboWaas2.WebhookEventType("wallets.transaction.created")}) // CreateWebhookEndpointRequest | The request body to register a webhook endpoint. (optional)

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.CreateWebhookEndpoint(ctx).CreateWebhookEndpointRequest(createWebhookEndpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.CreateWebhookEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhookEndpoint`: WebhookEndpoint
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

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookEndpointById

> WebhookEndpoint GetWebhookEndpointById(ctx, endpointId).Execute()

Get webhook endpoint information



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
	endpointId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints).

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.GetWebhookEndpointById(ctx, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.GetWebhookEndpointById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookEndpointById`: WebhookEndpoint
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

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookEventById

> WebhookEvent GetWebhookEventById(ctx, eventId, endpointId).Execute()

Retrieve event information



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
	eventId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The event ID. You can obtain a list of event IDs by calling [List all events](/v2/api-references/developers--webhooks/list-all-events).
	endpointId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints).

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.GetWebhookEventById(ctx, eventId, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.GetWebhookEventById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookEventById`: WebhookEvent
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.GetWebhookEventById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**eventId** | **string** | The event ID. You can obtain a list of event IDs by calling [List all events](/v2/api-references/developers--webhooks/list-all-events). | 
**endpointId** | **string** | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookEventByIdRequest struct via the builder pattern


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
	coboWaas2 "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
    "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2/crypto"
)

func main() {
	status := coboWaas2.WebhookEndpointStatus("STATUS_ACTIVE") // WebhookEndpointStatus |  (optional)
	eventType := coboWaas2.WebhookEventType("wallets.transaction.created") // WebhookEventType |  (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. For most operations, the value range is [1, 50]. (optional) (default to 10)
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1" // string | An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify `before` as `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1`, the request will retrieve a list of data objects that end before the object with the object ID `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  - If you set both `after` and `before`, an error will occur.  - If you leave both `before` and `after` empty, the first page of data is returned.  - If you set `before` to `infinity`, the last page of data is returned.  (optional)
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk" // string | An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify `after` as `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk`, the request will retrieve a list of data objects that start after the object with the object ID `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  - If you set both `after` and `before`, an error will occur.  - If you leave both `before` and `after` empty, the first page of data is returned.  (optional)

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
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
 **status** | [**WebhookEndpointStatus**](WebhookEndpointStatus.md) |  | 
 **eventType** | [**WebhookEventType**](WebhookEventType.md) |  | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify &#x60;before&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  - If you set &#x60;before&#x60; to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify &#x60;after&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 

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
	coboWaas2 "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
    "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2/crypto"
)

func main() {

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
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


## ListWebhookEventLogs

> ListWebhookEventLogs200Response ListWebhookEventLogs(ctx, eventId, endpointId).Limit(limit).Before(before).After(after).Execute()

List webhook event logs



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
	eventId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The event ID. You can obtain a list of event IDs by calling [List all events](/v2/api-references/developers--webhooks/list-all-events).
	endpointId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints).
	limit := int32(10) // int32 | The maximum number of objects to return. For most operations, the value range is [1, 50]. (optional) (default to 10)
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1" // string | An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify `before` as `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1`, the request will retrieve a list of data objects that end before the object with the object ID `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  - If you set both `after` and `before`, an error will occur.  - If you leave both `before` and `after` empty, the first page of data is returned.  - If you set `before` to `infinity`, the last page of data is returned.  (optional)
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk" // string | An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify `after` as `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk`, the request will retrieve a list of data objects that start after the object with the object ID `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  - If you set both `after` and `before`, an error will occur.  - If you leave both `before` and `after` empty, the first page of data is returned.  (optional)

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.ListWebhookEventLogs(ctx, eventId, endpointId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.ListWebhookEventLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhookEventLogs`: ListWebhookEventLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.ListWebhookEventLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**eventId** | **string** | The event ID. You can obtain a list of event IDs by calling [List all events](/v2/api-references/developers--webhooks/list-all-events). | 
**endpointId** | **string** | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints). | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookEventLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify &#x60;before&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  - If you set &#x60;before&#x60; to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify &#x60;after&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 

### Return type

[**ListWebhookEventLogs200Response**](ListWebhookEventLogs200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookEvents

> ListWebhookEvents200Response ListWebhookEvents(ctx, endpointId).Status(status).Type_(type_).Limit(limit).Before(before).After(after).Execute()

List all webhook events



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
	endpointId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints).
	status := coboWaas2.WebhookEventStatus("Success") // WebhookEventStatus |  (optional)
	type_ := coboWaas2.WebhookEventType("wallets.transaction.created") // WebhookEventType |  (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. For most operations, the value range is [1, 50]. (optional) (default to 10)
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1" // string | An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify `before` as `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1`, the request will retrieve a list of data objects that end before the object with the object ID `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  - If you set both `after` and `before`, an error will occur.  - If you leave both `before` and `after` empty, the first page of data is returned.  - If you set `before` to `infinity`, the last page of data is returned.  (optional)
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk" // string | An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify `after` as `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk`, the request will retrieve a list of data objects that start after the object with the object ID `RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  - If you set both `after` and `before`, an error will occur.  - If you leave both `before` and `after` empty, the first page of data is returned.  (optional)

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.ListWebhookEvents(ctx, endpointId).Status(status).Type_(type_).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.ListWebhookEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhookEvents`: ListWebhookEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.ListWebhookEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**endpointId** | **string** | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints). | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**WebhookEventStatus**](WebhookEventStatus.md) |  | 
 **type_** | [**WebhookEventType**](WebhookEventType.md) |  | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify &#x60;before&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  - If you set &#x60;before&#x60; to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify &#x60;after&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 

### Return type

[**ListWebhookEvents200Response**](ListWebhookEvents200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryWebhookEventById

> RetryWebhookEventById201Response RetryWebhookEventById(ctx, eventId, endpointId).Execute()

Retry event



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
	eventId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The event ID. You can obtain a list of event IDs by calling [List all events](/v2/api-references/developers--webhooks/list-all-events).
	endpointId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints).

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.RetryWebhookEventById(ctx, eventId, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.RetryWebhookEventById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryWebhookEventById`: RetryWebhookEventById201Response
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.RetryWebhookEventById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**eventId** | **string** | The event ID. You can obtain a list of event IDs by calling [List all events](/v2/api-references/developers--webhooks/list-all-events). | 
**endpointId** | **string** | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints). | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryWebhookEventByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RetryWebhookEventById201Response**](RetryWebhookEventById201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookEndpointById

> WebhookEndpoint UpdateWebhookEndpointById(ctx, endpointId).UpdateWebhookEndpointByIdRequest(updateWebhookEndpointByIdRequest).Execute()

Update webhook endpoint



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
	endpointId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints).
	updateWebhookEndpointByIdRequest := *coboWaas2.NewUpdateWebhookEndpointByIdRequest() // UpdateWebhookEndpointByIdRequest | The request body to update a webhook endpoint. (optional)

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, waas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, waas2.ContextEnv, waas2.DevEnv)
	ctx = context.WithValue(ctx, waas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.DevelopersWebhooksAPI.UpdateWebhookEndpointById(ctx, endpointId).UpdateWebhookEndpointByIdRequest(updateWebhookEndpointByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopersWebhooksAPI.UpdateWebhookEndpointById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookEndpointById`: WebhookEndpoint
	fmt.Fprintf(os.Stdout, "Response from `DevelopersWebhooksAPI.UpdateWebhookEndpointById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**endpointId** | **string** | The webhook endpoint ID. You can retrieve a list of webhook endpoint IDs by calling [List webhook endpoints](/v2/api-references/developers--webhooks/list-webhook-endpoints). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookEndpointByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWebhookEndpointByIdRequest** | [**UpdateWebhookEndpointByIdRequest**](UpdateWebhookEndpointByIdRequest.md) | The request body to update a webhook endpoint. | 

### Return type

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

