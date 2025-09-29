# \OAuthAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExchangePermissionToken**](OAuthAPI.md#ExchangePermissionToken) | **Post** /oauth/permission_token/exchange | Exchange Permission Access Token by API Key
[**GetToken**](OAuthAPI.md#GetToken) | **Get** /oauth/token | Get Org Access Token
[**RefreshPermissionToken**](OAuthAPI.md#RefreshPermissionToken) | **Post** /oauth/permission_token/refresh | Refresh Permission Access Token by Permission Refresh Token
[**RefreshToken**](OAuthAPI.md#RefreshToken) | **Post** /oauth/token | Refresh Org Access Token



## ExchangePermissionToken

> ExchangePermissionToken201Response ExchangePermissionToken(ctx).ExchangePermissionTokenRequest(exchangePermissionTokenRequest).Execute()

Exchange Permission Access Token by API Key



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
	exchangePermissionTokenRequest := *coboWaas2.NewExchangePermissionTokenRequest("payment_orders_payin")

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
	resp, r, err := apiClient.OAuthAPI.ExchangePermissionToken(ctx).ExchangePermissionTokenRequest(exchangePermissionTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.ExchangePermissionToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExchangePermissionToken`: ExchangePermissionToken201Response
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.ExchangePermissionToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExchangePermissionTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchangePermissionTokenRequest** | [**ExchangePermissionTokenRequest**](ExchangePermissionTokenRequest.md) | The request body for exchanging an Permission Access Token. | 

### Return type

[**ExchangePermissionToken201Response**](ExchangePermissionToken201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToken

> GetToken2XXResponse GetToken(ctx).ClientId(clientId).OrgId(orgId).GrantType(grantType).Execute()

Get Org Access Token



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
	clientId := "pvSwS8iFrfK0oZrB0ugG54XPDOLEv0Ij"
	orgId := "e3986401-4aec-480a-973d-e775a4518413"
	grantType := "org_implicit"

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
	resp, r, err := apiClient.OAuthAPI.GetToken(ctx).ClientId(clientId).OrgId(orgId).GrantType(grantType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.GetToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToken`: GetToken2XXResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.GetToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | The client ID, a unique identifier to distinguish Cobo Portal Apps. You can get the client ID by retrieving the manifest file after publishing the app. | 
 **orgId** | **string** | Organization ID, a unique identifier to distinguish different organizations. You can get the organization ID from the callback message sent to the URL that was configured in the manifest file. | 
 **grantType** | **string** | The OAuth grant type. Set the value as &#x60;org_implicit&#x60;. | 

### Return type

[**GetToken2XXResponse**](GetToken2XXResponse.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshPermissionToken

> ExchangePermissionToken201Response RefreshPermissionToken(ctx).RefreshPermissionTokenRequest(refreshPermissionTokenRequest).Execute()

Refresh Permission Access Token by Permission Refresh Token



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
	refreshPermissionTokenRequest := *coboWaas2.NewRefreshPermissionTokenRequest("rK49jI0zt49gsttzscscik15Asmlpu1TdcxqguJJS8B9f6ilJEC0y3PbVqwsEAw5")

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
	resp, r, err := apiClient.OAuthAPI.RefreshPermissionToken(ctx).RefreshPermissionTokenRequest(refreshPermissionTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.RefreshPermissionToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshPermissionToken`: ExchangePermissionToken201Response
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.RefreshPermissionToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshPermissionTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshPermissionTokenRequest** | [**RefreshPermissionTokenRequest**](RefreshPermissionTokenRequest.md) | The request body for refreshing an Permission Access Token. | 

### Return type

[**ExchangePermissionToken201Response**](ExchangePermissionToken201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshToken

> RefreshToken2XXResponse RefreshToken(ctx).RefreshTokenRequest(refreshTokenRequest).Execute()

Refresh Org Access Token



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
	refreshTokenRequest := *coboWaas2.NewRefreshTokenRequest()

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
	resp, r, err := apiClient.OAuthAPI.RefreshToken(ctx).RefreshTokenRequest(refreshTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.RefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshToken`: RefreshToken2XXResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.RefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshTokenRequest** | [**RefreshTokenRequest**](RefreshTokenRequest.md) | The request body for refreshing an Org Access Token. | 

### Return type

[**RefreshToken2XXResponse**](RefreshToken2XXResponse.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

