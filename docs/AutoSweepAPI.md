# \AutoSweepAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutoSweepTask**](AutoSweepAPI.md#CreateAutoSweepTask) | **Post** /auto_sweep/tasks | create auto sweep task
[**CreateWalletSweepToAddresses**](AutoSweepAPI.md#CreateWalletSweepToAddresses) | **Post** /auto_sweep/sweep_to_addresses | create wallet sweep to addresses
[**GetAutoSweepTaskById**](AutoSweepAPI.md#GetAutoSweepTaskById) | **Get** /auto_sweep/tasks/{task_id} | Get auto sweep task information
[**ListAutoSweepTask**](AutoSweepAPI.md#ListAutoSweepTask) | **Get** /auto_sweep/tasks | List wallet auto sweep task
[**ListWalletSweepToAddresses**](AutoSweepAPI.md#ListWalletSweepToAddresses) | **Get** /auto_sweep/sweep_to_addresses | List wallet sweep to addresses



## CreateAutoSweepTask

> AutoSweepTask CreateAutoSweepTask(ctx).CreateAutoSweepTask(createAutoSweepTask).Execute()

create auto sweep task



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
	createAutoSweepTask := *coboWaas2.NewCreateAutoSweepTask("f47ac10b-58cc-4372-a567-0e02b2c3d479", "ETH_USDT")

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
	resp, r, err := apiClient.AutoSweepAPI.CreateAutoSweepTask(ctx).CreateAutoSweepTask(createAutoSweepTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoSweepAPI.CreateAutoSweepTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutoSweepTask`: AutoSweepTask
	fmt.Fprintf(os.Stdout, "Response from `AutoSweepAPI.CreateAutoSweepTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoSweepTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAutoSweepTask** | [**CreateAutoSweepTask**](CreateAutoSweepTask.md) | The request body to generates a new sweep to addresses within a specified wallet. | 

### Return type

[**AutoSweepTask**](AutoSweepTask.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWalletSweepToAddresses

> SweepToAddress CreateWalletSweepToAddresses(ctx).CreateSweepToAddress(createSweepToAddress).Execute()

create wallet sweep to addresses



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
	createSweepToAddress := *coboWaas2.NewCreateSweepToAddress("f47ac10b-58cc-4372-a567-0e02b2c3d479", "ETH")

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
	resp, r, err := apiClient.AutoSweepAPI.CreateWalletSweepToAddresses(ctx).CreateSweepToAddress(createSweepToAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoSweepAPI.CreateWalletSweepToAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWalletSweepToAddresses`: SweepToAddress
	fmt.Fprintf(os.Stdout, "Response from `AutoSweepAPI.CreateWalletSweepToAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWalletSweepToAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSweepToAddress** | [**CreateSweepToAddress**](CreateSweepToAddress.md) | The request body to generates a new sweep to addresses within a specified wallet. | 

### Return type

[**SweepToAddress**](SweepToAddress.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoSweepTaskById

> AutoSweepTask GetAutoSweepTaskById(ctx, taskId).Execute()

Get auto sweep task information



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
	taskId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"

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
	resp, r, err := apiClient.AutoSweepAPI.GetAutoSweepTaskById(ctx, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoSweepAPI.GetAutoSweepTaskById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoSweepTaskById`: AutoSweepTask
	fmt.Fprintf(os.Stdout, "Response from `AutoSweepAPI.GetAutoSweepTaskById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**taskId** | **string** | The auto sweep task ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoSweepTaskByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoSweepTask**](AutoSweepTask.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAutoSweepTask

> ListAutoSweepTask200Response ListAutoSweepTask(ctx).WalletId(walletId).TokenId(tokenId).TaskIds(taskIds).MinCreatedTimestamp(minCreatedTimestamp).MaxCreatedTimestamp(maxCreatedTimestamp).Limit(limit).Before(before).After(after).Direction(direction).Execute()

List wallet auto sweep task



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	tokenId := "ETH_USDT"
	taskIds := "f47ac10b-58cc-4372-a567-0e02b2c3d479,1ddca562-8434-41c9-8809-d437bad9c868"
	minCreatedTimestamp := int64(1635744000000)
	maxCreatedTimestamp := int64(1635744000000)
	limit := int32(10)
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1"
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk"
	direction := "ASC"

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
	resp, r, err := apiClient.AutoSweepAPI.ListAutoSweepTask(ctx).WalletId(walletId).TokenId(tokenId).TaskIds(taskIds).MinCreatedTimestamp(minCreatedTimestamp).MaxCreatedTimestamp(maxCreatedTimestamp).Limit(limit).Before(before).After(after).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoSweepAPI.ListAutoSweepTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAutoSweepTask`: ListAutoSweepTask200Response
	fmt.Fprintf(os.Stdout, "Response from `AutoSweepAPI.ListAutoSweepTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAutoSweepTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletId** | **string** | The wallet ID. | 
 **tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
 **taskIds** | **string** | A list of auto sweep task IDs, separated by comma. | 
 **minCreatedTimestamp** | **int64** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. You can use this parameter to filter transactions created on or after the specified time. | 
 **maxCreatedTimestamp** | **int64** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. You can use this parameter to filter transactions created on or before the specified time. | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | This parameter specifies an object ID as a starting point for pagination, retrieving data before the specified object relative to the current dataset.    Suppose the current data is ordered as Object A, Object B, and Object C.  If you set &#x60;before&#x60; to the ID of Object C (&#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;), the response will include Object B and Object A.    **Notes**:   - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur. - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned. - If you set it to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | This parameter specifies an object ID as a starting point for pagination, retrieving data after the specified object relative to the current dataset.    Suppose the current data is ordered as Object A, Object B, and Object C. If you set &#x60;after&#x60; to the ID of Object A (&#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;), the response will include Object B and Object C.    **Notes**:   - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur. - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 
 **direction** | **string** | The sort direction. Possible values include:   - &#x60;ASC&#x60;: Sort the results in ascending order.   - &#x60;DESC&#x60;: Sort the results in descending order.  | [default to &quot;ASC&quot;]

### Return type

[**ListAutoSweepTask200Response**](ListAutoSweepTask200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWalletSweepToAddresses

> ListWalletSweepToAddresses200Response ListWalletSweepToAddresses(ctx).WalletId(walletId).Execute()

List wallet sweep to addresses



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"

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
	resp, r, err := apiClient.AutoSweepAPI.ListWalletSweepToAddresses(ctx).WalletId(walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoSweepAPI.ListWalletSweepToAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWalletSweepToAddresses`: ListWalletSweepToAddresses200Response
	fmt.Fprintf(os.Stdout, "Response from `AutoSweepAPI.ListWalletSweepToAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWalletSweepToAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletId** | **string** | The wallet ID. | 

### Return type

[**ListWalletSweepToAddresses200Response**](ListWalletSweepToAddresses200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

