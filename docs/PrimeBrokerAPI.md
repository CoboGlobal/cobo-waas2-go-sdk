# \PrimeBrokerAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeGuardPubkey**](PrimeBrokerAPI.md#ChangeGuardPubkey) | **Put** /prime_broker/user/{user_id}/guard_pubkey | Change Guard pubkey binding
[**CreateGuardPubkey**](PrimeBrokerAPI.md#CreateGuardPubkey) | **Post** /prime_broker/user/{user_id}/guard_pubkey | Create Guard pubkey binding
[**CreatePrimeBrokerAddress**](PrimeBrokerAPI.md#CreatePrimeBrokerAddress) | **Post** /prime_broker/user/{user_id}/addresses | Bind addresses to a broker user
[**DeleteGuardPubkey**](PrimeBrokerAPI.md#DeleteGuardPubkey) | **Post** /prime_broker/user/{user_id}/guard_pubkey/delete | Delete Guard pubkey binding
[**QueryApprovalStatement**](PrimeBrokerAPI.md#QueryApprovalStatement) | **Get** /prime_broker/approval_statement/{statement_id} | Query approval statement
[**QueryGuardPubkey**](PrimeBrokerAPI.md#QueryGuardPubkey) | **Get** /prime_broker/user/{user_id}/guard_pubkey | Query a Guard pubkey



## ChangeGuardPubkey

> ChangeGuardPubkey200Response ChangeGuardPubkey(ctx, userId).Execute()

Change Guard pubkey binding



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
	userId := "168108513539918"

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
	resp, r, err := apiClient.PrimeBrokerAPI.ChangeGuardPubkey(ctx, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrimeBrokerAPI.ChangeGuardPubkey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeGuardPubkey`: ChangeGuardPubkey200Response
	fmt.Fprintf(os.Stdout, "Response from `PrimeBrokerAPI.ChangeGuardPubkey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**userId** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeGuardPubkeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChangeGuardPubkey200Response**](ChangeGuardPubkey200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGuardPubkey

> ChangeGuardPubkey200Response CreateGuardPubkey(ctx, userId).Execute()

Create Guard pubkey binding



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
	userId := "168108513539918"

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
	resp, r, err := apiClient.PrimeBrokerAPI.CreateGuardPubkey(ctx, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrimeBrokerAPI.CreateGuardPubkey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGuardPubkey`: ChangeGuardPubkey200Response
	fmt.Fprintf(os.Stdout, "Response from `PrimeBrokerAPI.CreateGuardPubkey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**userId** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGuardPubkeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChangeGuardPubkey200Response**](ChangeGuardPubkey200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrimeBrokerAddress

> CreatePrimeBrokerAddress201Response CreatePrimeBrokerAddress(ctx, userId).CreatePrimeBrokerAddressRequest(createPrimeBrokerAddressRequest).Execute()

Bind addresses to a broker user



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
	userId := "168108513539918"
	createPrimeBrokerAddressRequest := *coboWaas2.NewCreatePrimeBrokerAddressRequest()

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
	resp, r, err := apiClient.PrimeBrokerAPI.CreatePrimeBrokerAddress(ctx, userId).CreatePrimeBrokerAddressRequest(createPrimeBrokerAddressRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrimeBrokerAPI.CreatePrimeBrokerAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrimeBrokerAddress`: CreatePrimeBrokerAddress201Response
	fmt.Fprintf(os.Stdout, "Response from `PrimeBrokerAPI.CreatePrimeBrokerAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**userId** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrimeBrokerAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPrimeBrokerAddressRequest** | [**CreatePrimeBrokerAddressRequest**](CreatePrimeBrokerAddressRequest.md) | The request body to bind addresses to a broker user. | 

### Return type

[**CreatePrimeBrokerAddress201Response**](CreatePrimeBrokerAddress201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGuardPubkey

> DeleteGuardPubkey201Response DeleteGuardPubkey(ctx, userId).Execute()

Delete Guard pubkey binding



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
	userId := "168108513539918"

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
	resp, r, err := apiClient.PrimeBrokerAPI.DeleteGuardPubkey(ctx, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrimeBrokerAPI.DeleteGuardPubkey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGuardPubkey`: DeleteGuardPubkey201Response
	fmt.Fprintf(os.Stdout, "Response from `PrimeBrokerAPI.DeleteGuardPubkey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**userId** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuardPubkeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteGuardPubkey201Response**](DeleteGuardPubkey201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryApprovalStatement

> QueryApprovalStatement200Response QueryApprovalStatement(ctx, statementId).Execute()

Query approval statement



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
	statementId := "168108513539918"

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
	resp, r, err := apiClient.PrimeBrokerAPI.QueryApprovalStatement(ctx, statementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrimeBrokerAPI.QueryApprovalStatement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryApprovalStatement`: QueryApprovalStatement200Response
	fmt.Fprintf(os.Stdout, "Response from `PrimeBrokerAPI.QueryApprovalStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**statementId** | **string** | The approval statement ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryApprovalStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryApprovalStatement200Response**](QueryApprovalStatement200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryGuardPubkey

> QueryGuardPubkey200Response QueryGuardPubkey(ctx, userId).Execute()

Query a Guard pubkey



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
	userId := "168108513539918"

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
	resp, r, err := apiClient.PrimeBrokerAPI.QueryGuardPubkey(ctx, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrimeBrokerAPI.QueryGuardPubkey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryGuardPubkey`: QueryGuardPubkey200Response
	fmt.Fprintf(os.Stdout, "Response from `PrimeBrokerAPI.QueryGuardPubkey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**userId** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryGuardPubkeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryGuardPubkey200Response**](QueryGuardPubkey200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

