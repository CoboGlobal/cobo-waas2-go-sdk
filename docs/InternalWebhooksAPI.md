# \InternalWebhooksAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookEvent**](InternalWebhooksAPI.md#CreateWebhookEvent) | **Post** /internal/webhook_event | Create webhook event



## CreateWebhookEvent

> CreateWebhookEventInfo CreateWebhookEvent(ctx).CreateWebhookEventParams(createWebhookEventParams).Execute()

Create webhook event



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
	createWebhookEventParams := *coboWaas2.NewCreateWebhookEventParams("ChannelId_example", coboWaas2.WebhookEventType("wallets.transaction.created"), map[string]interface{}{"key": interface{}(123)}, "Uuid_example")

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
	resp, r, err := apiClient.InternalWebhooksAPI.CreateWebhookEvent(ctx).CreateWebhookEventParams(createWebhookEventParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalWebhooksAPI.CreateWebhookEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhookEvent`: CreateWebhookEventInfo
	fmt.Fprintf(os.Stdout, "Response from `InternalWebhooksAPI.CreateWebhookEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebhookEventParams** | [**CreateWebhookEventParams**](CreateWebhookEventParams.md) | The request body to create a webhook event | 

### Return type

[**CreateWebhookEventInfo**](CreateWebhookEventInfo.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

