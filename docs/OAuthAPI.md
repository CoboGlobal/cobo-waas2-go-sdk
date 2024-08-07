# \OAuthAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetToken**](OAuthAPI.md#GetToken) | **Get** /oauth/token | Get Access Token
[**RefreshToken**](OAuthAPI.md#RefreshToken) | **Post** /oauth/token | Refresh Access Token



## GetToken

> GetToken200Response GetToken(ctx).ClientId(clientId).OrgId(orgId).GrantType(grantType).Execute()

Get Access Token



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
	clientId := "pvSwS8iFrfK0oZrB0ugG54XPDOLEv0Ij" // string | The App ID, a unique identifier to distinguish Cobo Portal Apps. You can get the App ID by retrieving the Manifest file after receiving the notification of app launch approval.
	orgId := "e3986401-4aec-480a-973d-e775a4518413" // string | Org ID, a unique identifier to distinguish different organizations. You can get the Org ID by retrieving the Manifest file after receiving the notification of app launch approval.
	grantType := "org_implicit" // string | The type of the permission granting. To get an access token, you need to set the value as `org_implicit`.

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, coboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, coboWaas2.ContextEnv, coboWaas2.DevEnv)
	ctx = context.WithValue(ctx, coboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.OAuthAPI.GetToken(ctx).ClientId(clientId).OrgId(orgId).GrantType(grantType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.GetToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToken`: GetToken200Response
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.GetToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | The App ID, a unique identifier to distinguish Cobo Portal Apps. You can get the App ID by retrieving the Manifest file after receiving the notification of app launch approval. | 
 **orgId** | **string** | Org ID, a unique identifier to distinguish different organizations. You can get the Org ID by retrieving the Manifest file after receiving the notification of app launch approval. | 
 **grantType** | **string** | The type of the permission granting. To get an access token, you need to set the value as &#x60;org_implicit&#x60;. | 

### Return type

[**GetToken200Response**](GetToken200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshToken

> GetToken200Response RefreshToken(ctx).RefreshTokenRequest(refreshTokenRequest).Execute()

Refresh Access Token



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
	refreshTokenRequest := *coboWaas2.NewRefreshTokenRequest() // RefreshTokenRequest | The request body for refreshing an access token.

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, coboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, coboWaas2.ContextEnv, coboWaas2.DevEnv)
	ctx = context.WithValue(ctx, coboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.OAuthAPI.RefreshToken(ctx).RefreshTokenRequest(refreshTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.RefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshToken`: GetToken200Response
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.RefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshTokenRequest** | [**RefreshTokenRequest**](RefreshTokenRequest.md) | The request body for refreshing an access token. | 

### Return type

[**GetToken200Response**](GetToken200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

