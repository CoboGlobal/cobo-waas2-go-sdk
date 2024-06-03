# \WalletsMPCWalletAPI

All URIs are relative to *https://api.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTssRequest**](WalletsMPCWalletAPI.md#CancelTssRequest) | **Put** /wallets/mpc/vaults/{vault_id}/tss_requests/{tss_request_id} | Cancel a tss request
[**CreateKeyGroup**](WalletsMPCWalletAPI.md#CreateKeyGroup) | **Post** /wallets/mpc/vaults/{vault_id}/key_groups | Create a mpc key group
[**CreateMpcProject**](WalletsMPCWalletAPI.md#CreateMpcProject) | **Post** /wallets/mpc/projects | Create a mpc project
[**CreateMpcVault**](WalletsMPCWalletAPI.md#CreateMpcVault) | **Post** /wallets/mpc/vaults | Create a mpc vault
[**CreateTssRequest**](WalletsMPCWalletAPI.md#CreateTssRequest) | **Post** /wallets/mpc/vaults/{vault_id}/tss_requests | Create a tss request to generate key secrets for a tss group
[**DeleteKeyGroup**](WalletsMPCWalletAPI.md#DeleteKeyGroup) | **Delete** /wallets/mpc/vaults/{vault_id}/key_groups/{key_group_id} | Delete a mpc key group
[**GetKeyGroup**](WalletsMPCWalletAPI.md#GetKeyGroup) | **Get** /wallets/mpc/vaults/{vault_id}/key_groups/{key_group_id} | Get a mpc key group information by group id
[**GetMpcProject**](WalletsMPCWalletAPI.md#GetMpcProject) | **Get** /wallets/mpc/projects/{project_id} | Get a mpc project information
[**GetMpcVault**](WalletsMPCWalletAPI.md#GetMpcVault) | **Get** /wallets/mpc/vaults/{vault_id} | Get a mpc vault information
[**GetTssRequest**](WalletsMPCWalletAPI.md#GetTssRequest) | **Get** /wallets/mpc/vaults/{vault_id}/tss_requests/{tss_request_id} | Get a tss request information
[**ListCoboKeyHolder**](WalletsMPCWalletAPI.md#ListCoboKeyHolder) | **Get** /wallets/mpc/cobo_key_holders | List all cobo key holders
[**ListKeyGroup**](WalletsMPCWalletAPI.md#ListKeyGroup) | **Get** /wallets/mpc/vaults/{vault_id}/key_groups | List all mpc key groups
[**ListMpcProject**](WalletsMPCWalletAPI.md#ListMpcProject) | **Get** /wallets/mpc/projects | List all mpc projects
[**ListMpcVault**](WalletsMPCWalletAPI.md#ListMpcVault) | **Get** /wallets/mpc/vaults | List all mpc vaults
[**ListTssRequest**](WalletsMPCWalletAPI.md#ListTssRequest) | **Get** /wallets/mpc/vaults/{vault_id}/tss_requests | List tss request information by vault ID
[**ModifyMpcVault**](WalletsMPCWalletAPI.md#ModifyMpcVault) | **Put** /wallets/mpc/vaults/{vault_id} | Update a mpc vault information
[**UpdateKeyGroup**](WalletsMPCWalletAPI.md#UpdateKeyGroup) | **Put** /wallets/mpc/vaults/{vault_id}/key_groups/{key_group_id} | Update a mpc key group information
[**UpdateMpcProject**](WalletsMPCWalletAPI.md#UpdateMpcProject) | **Put** /wallets/mpc/projects/{project_id} | Update a mpc project



## CancelTssRequest

> TSSRequest CancelTssRequest(ctx, vaultId, tssRequestId).TssRequestAction(tssRequestAction).Execute()

Cancel a tss request



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	tssRequestId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the tss request
	tssRequestAction := "cancel" // string | The action of tss request.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.CancelTssRequest(ctx, vaultId, tssRequestId).TssRequestAction(tssRequestAction).Execute()
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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
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

Create a mpc key group



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	createKeyGroupRequest := *CoboWaas2.NewCreateKeyGroupRequest(CoboWaas2.KeyGroupType("MainKeyGroup"), int32(123), int32(123), []CoboWaas2.CreateKeyGroupRequestKeyHoldersInner{*CoboWaas2.NewCreateKeyGroupRequestKeyHoldersInner()}) // CreateKeyGroupRequest | The request body to create a mpc key group (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.CreateKeyGroup(ctx, vaultId).CreateKeyGroupRequest(createKeyGroupRequest).Execute()
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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
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
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	createMpcProjectRequest := *CoboWaas2.NewCreateMpcProjectRequest(int32(123), int32(123)) // CreateMpcProjectRequest | The request body to create a mpc project (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.CreateMpcProject(ctx).CreateMpcProjectRequest(createMpcProjectRequest).Execute()
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
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	createMpcVaultRequest := *CoboWaas2.NewCreateMpcVaultRequest("My mpc vault", CoboWaas2.MPCVaultType("OrgControlled")) // CreateMpcVaultRequest | The request body to create a mpc vault (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.CreateMpcVault(ctx).CreateMpcVaultRequest(createMpcVaultRequest).Execute()
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
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	createTssRequestRequest := *CoboWaas2.NewCreateTssRequestRequest(CoboWaas2.TSSRequestType("KeyGen"), "TargetKeyGroupId_example") // CreateTssRequestRequest | The request body to create a tss request (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.CreateTssRequest(ctx, vaultId).CreateTssRequestRequest(createTssRequestRequest).Execute()
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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
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

Delete a mpc key group



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	keyGroupId := "880311524363903326" // string | Unique id of the tss group

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.DeleteKeyGroup(ctx, vaultId, keyGroupId).Execute()
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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
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

Get a mpc key group information by group id



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	keyGroupId := "880311524363903326" // string | Unique id of the tss group

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.GetKeyGroup(ctx, vaultId, keyGroupId).Execute()
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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
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

Get a mpc project information



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
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc project

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.GetMpcProject(ctx, projectId).Execute()
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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
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

Get a mpc vault information



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.GetMpcVault(ctx, vaultId).Execute()
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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
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

Get a tss request information



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	tssRequestId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the tss request

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.GetTssRequest(ctx, vaultId, tssRequestId).Execute()
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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
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
	resp, r, err := apiClient.WalletsMPCWalletAPI.ListCoboKeyHolder(ctx).Execute()
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
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	keyGroupType := CoboWaas2.KeyGroupType("MainKeyGroup") // KeyGroupType | The type of key group. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.ListKeyGroup(ctx, vaultId).KeyGroupType(keyGroupType).Execute()
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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
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
	resp, r, err := apiClient.WalletsMPCWalletAPI.ListMpcProject(ctx).Execute()
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
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc project (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.ListMpcVault(ctx).ProjectId(projectId).Execute()
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

List tss request information by vault ID



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	targetKeyGroupId := "880311524363903326" // string | The target key group id of tss request. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.ListTssRequest(ctx, vaultId).TargetKeyGroupId(targetKeyGroupId).Execute()
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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
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

Update a mpc vault information



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	modifyMpcVaultRequest := *CoboWaas2.NewModifyMpcVaultRequest("my mpc vault name") // ModifyMpcVaultRequest | The request body to update a mpc vault (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.ModifyMpcVault(ctx, vaultId).ModifyMpcVaultRequest(modifyMpcVaultRequest).Execute()
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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
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

Update a mpc key group information



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc vault
	keyGroupId := "880311524363903326" // string | Unique id of the tss group
	updateKeyGroupAction := "UpgradeToMainKeyGroup" // string | The action of update key group

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.UpdateKeyGroup(ctx, vaultId, keyGroupId).UpdateKeyGroupAction(updateKeyGroupAction).Execute()
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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
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

Update a mpc project



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
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | Unique id of the mpc project
	updateMpcProjectRequest := *CoboWaas2.NewUpdateMpcProjectRequest("My mpc new project name") // UpdateMpcProjectRequest | The request body to update a mpc project (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletAPI.UpdateMpcProject(ctx, projectId).UpdateMpcProjectRequest(updateMpcProjectRequest).Execute()
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
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
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

