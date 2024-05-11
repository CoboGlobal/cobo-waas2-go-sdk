# \WalletsMPCWalletAPI

All URIs are relative to *https://api.cobo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTssRequest**](WalletsMPCWalletAPI.md#CancelTssRequest) | **Put** /wallets/mpc/vaults/{vault_id}/tss_requests/{tss_request_id} | cancel tss request
[**CreateKeyGroup**](WalletsMPCWalletAPI.md#CreateKeyGroup) | **Post** /wallets/mpc/vaults/{vault_id}/key_groups | create a mpc key group
[**CreateMpcProject**](WalletsMPCWalletAPI.md#CreateMpcProject) | **Post** /wallets/mpc/projects | Create a mpc project
[**CreateMpcVault**](WalletsMPCWalletAPI.md#CreateMpcVault) | **Post** /wallets/mpc/vaults | Create a mpc vault
[**CreateTssRequest**](WalletsMPCWalletAPI.md#CreateTssRequest) | **Post** /wallets/mpc/vaults/{vault_id}/tss_requests | Create a tss request to generate key secrets for a tss group
[**DeleteKeyGroup**](WalletsMPCWalletAPI.md#DeleteKeyGroup) | **Delete** /wallets/mpc/vaults/{vault_id}/key_groups/{key_group_id} | delete a mpc key group
[**GetKeyGroup**](WalletsMPCWalletAPI.md#GetKeyGroup) | **Get** /wallets/mpc/vaults/{vault_id}/key_groups/{key_group_id} | get a mpc key group
[**GetMpcProject**](WalletsMPCWalletAPI.md#GetMpcProject) | **Get** /wallets/mpc/projects/{project_id} | get a mpc project
[**GetMpcVault**](WalletsMPCWalletAPI.md#GetMpcVault) | **Get** /wallets/mpc/vaults/{vault_id} | get a mpc vault
[**GetTssRequest**](WalletsMPCWalletAPI.md#GetTssRequest) | **Get** /wallets/mpc/vaults/{vault_id}/tss_requests/{tss_request_id} | get a tss request
[**ListCoboKeyHolder**](WalletsMPCWalletAPI.md#ListCoboKeyHolder) | **Get** /wallets/mpc/cobo_key_holders | List all cobo key holders
[**ListKeyGroup**](WalletsMPCWalletAPI.md#ListKeyGroup) | **Get** /wallets/mpc/vaults/{vault_id}/key_groups | List all mpc key groups
[**ListMpcProject**](WalletsMPCWalletAPI.md#ListMpcProject) | **Get** /wallets/mpc/projects | List all mpc projects
[**ListMpcVault**](WalletsMPCWalletAPI.md#ListMpcVault) | **Get** /wallets/mpc/vaults | List all mpc vaults
[**ListTssRequest**](WalletsMPCWalletAPI.md#ListTssRequest) | **Get** /wallets/mpc/vaults/{vault_id}/tss_requests | List tss request information of a vault
[**ModifyMpcVault**](WalletsMPCWalletAPI.md#ModifyMpcVault) | **Put** /wallets/mpc/vaults/{vault_id} | Modify a mpc vault
[**UpdateKeyGroup**](WalletsMPCWalletAPI.md#UpdateKeyGroup) | **Put** /wallets/mpc/vaults/{vault_id}/key_groups/{key_group_id} | update a mpc key group
[**UpdateMpcProject**](WalletsMPCWalletAPI.md#UpdateMpcProject) | **Put** /wallets/mpc/projects/{project_id} | update a mpc project



## CancelTssRequest

> TSSRequest CancelTssRequest(ctx, vaultId, tssRequestId).TssRequestAction(tssRequestAction).Execute()

cancel tss request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	tssRequestId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the tss request
	tssRequestAction := "cancel" // string | The action of tss request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.CancelTssRequest(context.Background(), vaultId, tssRequestId).TssRequestAction(tssRequestAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.CancelTssRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelTssRequest`: TSSRequest
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.CancelTssRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** | Unique id of the mpc vault | 
**tssRequestId** | **string** | Unique id of the tss request | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelTssRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tssRequestAction** | **string** | The action of tss request. | 

### Return type

[**TSSRequest**](TSSRequest.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKeyGroup

> KeyGroup CreateKeyGroup(ctx, vaultId).CreateKeyGroupRequest(createKeyGroupRequest).Execute()

create a mpc key group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	createKeyGroupRequest := *openapiclient.NewCreateKeyGroupRequest(openapiclient.KeyGroupType("MainKeyGroup"), int32(123), int32(123), []openapiclient.CreateKeyGroupRequestKeyHoldersInner{*openapiclient.NewCreateKeyGroupRequestKeyHoldersInner()}) // CreateKeyGroupRequest | The request body to create a mpc key group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.CreateKeyGroup(context.Background(), vaultId).CreateKeyGroupRequest(createKeyGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.CreateKeyGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKeyGroup`: KeyGroup
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.CreateKeyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** | Unique id of the mpc vault | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createKeyGroupRequest** | [**CreateKeyGroupRequest**](CreateKeyGroupRequest.md) | The request body to create a mpc key group | 

### Return type

[**KeyGroup**](KeyGroup.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMpcProject

> MPCProject CreateMpcProject(ctx).CreateMpcProjectRequest(createMpcProjectRequest).Execute()

Create a mpc project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	createMpcProjectRequest := *openapiclient.NewCreateMpcProjectRequest(int32(123), int32(123)) // CreateMpcProjectRequest | The request body to create a mpc project (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.CreateMpcProject(context.Background()).CreateMpcProjectRequest(createMpcProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.CreateMpcProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMpcProject`: MPCProject
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.CreateMpcProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMpcProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMpcProjectRequest** | [**CreateMpcProjectRequest**](CreateMpcProjectRequest.md) | The request body to create a mpc project | 

### Return type

[**MPCProject**](MPCProject.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMpcVault

> MPCVault CreateMpcVault(ctx).CreateMpcVaultRequest(createMpcVaultRequest).Execute()

Create a mpc vault



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	createMpcVaultRequest := *openapiclient.NewCreateMpcVaultRequest("My mpc vault", openapiclient.MPCVaultType("OrgControlled")) // CreateMpcVaultRequest | The request body to create a mpc vault (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.CreateMpcVault(context.Background()).CreateMpcVaultRequest(createMpcVaultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.CreateMpcVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMpcVault`: MPCVault
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.CreateMpcVault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMpcVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMpcVaultRequest** | [**CreateMpcVaultRequest**](CreateMpcVaultRequest.md) | The request body to create a mpc vault | 

### Return type

[**MPCVault**](MPCVault.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTssRequest

> TSSRequest CreateTssRequest(ctx, vaultId).CreateTssRequestRequest(createTssRequestRequest).Execute()

Create a tss request to generate key secrets for a tss group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	createTssRequestRequest := *openapiclient.NewCreateTssRequestRequest(openapiclient.TSSRequestType("KeyGen"), "TargetKeyGroupId_example") // CreateTssRequestRequest | The request body to create a tss request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.CreateTssRequest(context.Background(), vaultId).CreateTssRequestRequest(createTssRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.CreateTssRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTssRequest`: TSSRequest
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.CreateTssRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** | Unique id of the mpc vault | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTssRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTssRequestRequest** | [**CreateTssRequestRequest**](CreateTssRequestRequest.md) | The request body to create a tss request | 

### Return type

[**TSSRequest**](TSSRequest.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKeyGroup

> KeyGroup DeleteKeyGroup(ctx, vaultId, keyGroupId).Execute()

delete a mpc key group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	keyGroupId := "880311524363903326" // string | Unique id of the tss group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.DeleteKeyGroup(context.Background(), vaultId, keyGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.DeleteKeyGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKeyGroup`: KeyGroup
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.DeleteKeyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** | Unique id of the mpc vault | 
**keyGroupId** | **string** | Unique id of the tss group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KeyGroup**](KeyGroup.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyGroup

> KeyGroup GetKeyGroup(ctx, vaultId, keyGroupId).Execute()

get a mpc key group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	keyGroupId := "880311524363903326" // string | Unique id of the tss group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.GetKeyGroup(context.Background(), vaultId, keyGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.GetKeyGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeyGroup`: KeyGroup
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.GetKeyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** | Unique id of the mpc vault | 
**keyGroupId** | **string** | Unique id of the tss group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KeyGroup**](KeyGroup.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMpcProject

> MPCProject GetMpcProject(ctx, projectId).Execute()

get a mpc project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.GetMpcProject(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.GetMpcProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMpcProject`: MPCProject
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.GetMpcProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique id of the mpc project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMpcProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MPCProject**](MPCProject.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMpcVault

> MPCVault GetMpcVault(ctx, vaultId).Execute()

get a mpc vault



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.GetMpcVault(context.Background(), vaultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.GetMpcVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMpcVault`: MPCVault
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.GetMpcVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** | Unique id of the mpc vault | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMpcVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MPCVault**](MPCVault.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTssRequest

> TSSRequest GetTssRequest(ctx, vaultId, tssRequestId).Execute()

get a tss request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	tssRequestId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the tss request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.GetTssRequest(context.Background(), vaultId, tssRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.GetTssRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTssRequest`: TSSRequest
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.GetTssRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** | Unique id of the mpc vault | 
**tssRequestId** | **string** | Unique id of the tss request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTssRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TSSRequest**](TSSRequest.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCoboKeyHolder

> []KeyHolder ListCoboKeyHolder(ctx).Execute()

List all cobo key holders



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.ListCoboKeyHolder(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.ListCoboKeyHolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCoboKeyHolder`: []KeyHolder
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.ListCoboKeyHolder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCoboKeyHolderRequest struct via the builder pattern


### Return type

[**[]KeyHolder**](KeyHolder.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeyGroup

> []KeyGroup ListKeyGroup(ctx, vaultId).KeyGroupType(keyGroupType).Execute()

List all mpc key groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	keyGroupType := openapiclient.KeyGroupType("MainKeyGroup") // KeyGroupType | The type of key group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.ListKeyGroup(context.Background(), vaultId).KeyGroupType(keyGroupType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.ListKeyGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKeyGroup`: []KeyGroup
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.ListKeyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** | Unique id of the mpc vault | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKeyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyGroupType** | [**KeyGroupType**](KeyGroupType.md) | The type of key group. | 

### Return type

[**[]KeyGroup**](KeyGroup.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMpcProject

> []MPCProject ListMpcProject(ctx).Execute()

List all mpc projects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.ListMpcProject(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.ListMpcProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMpcProject`: []MPCProject
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.ListMpcProject`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMpcProjectRequest struct via the builder pattern


### Return type

[**[]MPCProject**](MPCProject.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMpcVault

> []MPCVault ListMpcVault(ctx).ProjectId(projectId).Execute()

List all mpc vaults



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc project (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.ListMpcVault(context.Background()).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.ListMpcVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMpcVault`: []MPCVault
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.ListMpcVault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMpcVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | Unique id of the mpc project | 

### Return type

[**[]MPCVault**](MPCVault.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTssRequest

> TSSRequest ListTssRequest(ctx, vaultId).TargetKeyGroupId(targetKeyGroupId).Execute()

List tss request information of a vault



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	targetKeyGroupId := "880311524363903326" // string | The target key group id of tss request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.ListTssRequest(context.Background(), vaultId).TargetKeyGroupId(targetKeyGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.ListTssRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTssRequest`: TSSRequest
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.ListTssRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** | Unique id of the mpc vault | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTssRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetKeyGroupId** | **string** | The target key group id of tss request. | 

### Return type

[**TSSRequest**](TSSRequest.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMpcVault

> MPCVault ModifyMpcVault(ctx, vaultId).ModifyMpcVaultRequest(modifyMpcVaultRequest).Execute()

Modify a mpc vault



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	modifyMpcVaultRequest := *openapiclient.NewModifyMpcVaultRequest("my mpc vault name") // ModifyMpcVaultRequest | The request body to update a mpc vault (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.ModifyMpcVault(context.Background(), vaultId).ModifyMpcVaultRequest(modifyMpcVaultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.ModifyMpcVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMpcVault`: MPCVault
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.ModifyMpcVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** | Unique id of the mpc vault | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMpcVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyMpcVaultRequest** | [**ModifyMpcVaultRequest**](ModifyMpcVaultRequest.md) | The request body to update a mpc vault | 

### Return type

[**MPCVault**](MPCVault.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKeyGroup

> KeyGroup UpdateKeyGroup(ctx, vaultId, keyGroupId).UpdateKeyGroupAction(updateKeyGroupAction).Execute()

update a mpc key group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	keyGroupId := "880311524363903326" // string | Unique id of the tss group
	updateKeyGroupAction := "UpgradeToMainKeyGroup" // string | The action of update key group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.UpdateKeyGroup(context.Background(), vaultId, keyGroupId).UpdateKeyGroupAction(updateKeyGroupAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.UpdateKeyGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKeyGroup`: KeyGroup
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.UpdateKeyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** | Unique id of the mpc vault | 
**keyGroupId** | **string** | Unique id of the tss group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateKeyGroupAction** | **string** | The action of update key group | 

### Return type

[**KeyGroup**](KeyGroup.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMpcProject

> MPCProject UpdateMpcProject(ctx, projectId).UpdateMpcProjectRequest(updateMpcProjectRequest).Execute()

update a mpc project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CoboGlobal/cobo-waas2-go-api"
)

func main() {
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc project
	updateMpcProjectRequest := *openapiclient.NewUpdateMpcProjectRequest("My mpc new project name") // UpdateMpcProjectRequest | The request body to update a mpc project (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsMPCWalletAPI.UpdateMpcProject(context.Background(), projectId).UpdateMpcProjectRequest(updateMpcProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletAPI.UpdateMpcProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMpcProject`: MPCProject
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletAPI.UpdateMpcProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique id of the mpc project | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMpcProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMpcProjectRequest** | [**UpdateMpcProjectRequest**](UpdateMpcProjectRequest.md) | The request body to update a mpc project | 

### Return type

[**MPCProject**](MPCProject.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

