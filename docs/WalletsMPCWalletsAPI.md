# \WalletsMPCWalletsAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTssRequest**](WalletsMPCWalletsAPI.md#CancelTssRequest) | **Post** /wallets/mpc/vaults/{vault_id}/tss_requests/{tss_request_id}/cancel | Cancel TSS request
[**CreateKeyGroup**](WalletsMPCWalletsAPI.md#CreateKeyGroup) | **Post** /wallets/mpc/vaults/{vault_id}/key_groups | Create key share group
[**CreateMpcProject**](WalletsMPCWalletsAPI.md#CreateMpcProject) | **Post** /wallets/mpc/projects | Create project
[**CreateMpcVault**](WalletsMPCWalletsAPI.md#CreateMpcVault) | **Post** /wallets/mpc/vaults | Create vault
[**CreateTssRequest**](WalletsMPCWalletsAPI.md#CreateTssRequest) | **Post** /wallets/mpc/vaults/{vault_id}/tss_requests | Create TSS request
[**DeleteKeyGroup**](WalletsMPCWalletsAPI.md#DeleteKeyGroup) | **Delete** /wallets/mpc/vaults/{vault_id}/key_groups/{key_share_group_id} | Delete key share group
[**GetKeyGroup**](WalletsMPCWalletsAPI.md#GetKeyGroup) | **Get** /wallets/mpc/vaults/{vault_id}/key_groups/{key_share_group_id} | Get key share group information
[**GetMpcProject**](WalletsMPCWalletsAPI.md#GetMpcProject) | **Get** /wallets/mpc/projects/{project_id} | Get project information
[**GetMpcVault**](WalletsMPCWalletsAPI.md#GetMpcVault) | **Get** /wallets/mpc/vaults/{vault_id} | Get vault information
[**GetTssRequest**](WalletsMPCWalletsAPI.md#GetTssRequest) | **Get** /wallets/mpc/vaults/{vault_id}/tss_requests/{tss_request_id} | Get TSS request
[**ListCoboKeyHolder**](WalletsMPCWalletsAPI.md#ListCoboKeyHolder) | **Get** /wallets/mpc/cobo_key_holders | List all Cobo key share holders
[**ListKeyGroup**](WalletsMPCWalletsAPI.md#ListKeyGroup) | **Get** /wallets/mpc/vaults/{vault_id}/key_groups | List all key share groups
[**ListMpcProject**](WalletsMPCWalletsAPI.md#ListMpcProject) | **Get** /wallets/mpc/projects | List all projects
[**ListMpcVault**](WalletsMPCWalletsAPI.md#ListMpcVault) | **Get** /wallets/mpc/vaults | List all vaults
[**ListTssRequest**](WalletsMPCWalletsAPI.md#ListTssRequest) | **Get** /wallets/mpc/vaults/{vault_id}/tss_requests | List TSS requests
[**ModifyMpcVault**](WalletsMPCWalletsAPI.md#ModifyMpcVault) | **Put** /wallets/mpc/vaults/{vault_id} | Update vault name
[**UpdateKeyGroup**](WalletsMPCWalletsAPI.md#UpdateKeyGroup) | **Put** /wallets/mpc/vaults/{vault_id}/key_groups/{key_share_group_id} | Update key share group
[**UpdateMpcProject**](WalletsMPCWalletsAPI.md#UpdateMpcProject) | **Put** /wallets/mpc/projects/{project_id} | Update project name



## CancelTssRequest

> TSSRequest CancelTssRequest(ctx, vaultId, tssRequestId).Execute()

Cancel TSS request



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The vault ID.
	tssRequestId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The TSS request ID.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.CancelTssRequest(ctx, vaultId, tssRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.CancelTssRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelTssRequest`: TSSRequest
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.CancelTssRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID. | 
**tssRequestId** | **string** | The TSS request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelTssRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TSSRequest**](TSSRequest.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKeyGroup

> KeyGroup CreateKeyGroup(ctx, vaultId).CreateKeyGroupRequest(createKeyGroupRequest).Execute()

Create key share group



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The vault ID.
	createKeyGroupRequest := *CoboWaas2.NewCreateKeyGroupRequest(CoboWaas2.KeyGroupType("MainGroup"), int32(3), int32(2), []CoboWaas2.CreateKeyGroupRequestKeyHoldersInner{*CoboWaas2.NewCreateKeyGroupRequestKeyHoldersInner()}) // CreateKeyGroupRequest | The request body to create a key share group. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.CreateKeyGroup(ctx, vaultId).CreateKeyGroupRequest(createKeyGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.CreateKeyGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKeyGroup`: KeyGroup
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.CreateKeyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createKeyGroupRequest** | [**CreateKeyGroupRequest**](CreateKeyGroupRequest.md) | The request body to create a key share group. | 

### Return type

[**KeyGroup**](KeyGroup.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMpcProject

> MPCProject CreateMpcProject(ctx).CreateMpcProjectRequest(createMpcProjectRequest).Execute()

Create project



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
	createMpcProjectRequest := *CoboWaas2.NewCreateMpcProjectRequest("Project name", int32(3), int32(2)) // CreateMpcProjectRequest | The request body to create a project. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.CreateMpcProject(ctx).CreateMpcProjectRequest(createMpcProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.CreateMpcProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMpcProject`: MPCProject
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.CreateMpcProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMpcProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMpcProjectRequest** | [**CreateMpcProjectRequest**](CreateMpcProjectRequest.md) | The request body to create a project. | 

### Return type

[**MPCProject**](MPCProject.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMpcVault

> MPCVault CreateMpcVault(ctx).CreateMpcVaultRequest(createMpcVaultRequest).Execute()

Create vault



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
	createMpcVaultRequest := *CoboWaas2.NewCreateMpcVaultRequest("Name of new vault", CoboWaas2.MPCVaultType("OrgControlled")) // CreateMpcVaultRequest | The request body to create a vault. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.CreateMpcVault(ctx).CreateMpcVaultRequest(createMpcVaultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.CreateMpcVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMpcVault`: MPCVault
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.CreateMpcVault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMpcVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMpcVaultRequest** | [**CreateMpcVaultRequest**](CreateMpcVaultRequest.md) | The request body to create a vault. | 

### Return type

[**MPCVault**](MPCVault.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTssRequest

> TSSRequest CreateTssRequest(ctx, vaultId).CreateTssRequestRequest(createTssRequestRequest).Execute()

Create TSS request



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The vault ID.
	createTssRequestRequest := *CoboWaas2.NewCreateTssRequestRequest(CoboWaas2.TSSRequestType("KeyGen"), "123456789012345678") // CreateTssRequestRequest | The request body to create a TSS request. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.CreateTssRequest(ctx, vaultId).CreateTssRequestRequest(createTssRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.CreateTssRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTssRequest`: TSSRequest
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.CreateTssRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTssRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTssRequestRequest** | [**CreateTssRequestRequest**](CreateTssRequestRequest.md) | The request body to create a TSS request. | 

### Return type

[**TSSRequest**](TSSRequest.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKeyGroup

> DeleteKeyGroup(ctx, vaultId, keyShareGroupId).Execute()

Delete key share group



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The vault ID.
	keyShareGroupId := "880311524363903326" // string | The key share group ID.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	r, err := apiClient.WalletsMPCWalletsAPI.DeleteKeyGroup(ctx, vaultId, keyShareGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.DeleteKeyGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID. | 
**keyShareGroupId** | **string** | The key share group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyGroup

> KeyGroup GetKeyGroup(ctx, vaultId, keyShareGroupId).Execute()

Get key share group information



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The vault ID.
	keyShareGroupId := "880311524363903326" // string | The key share group ID.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.GetKeyGroup(ctx, vaultId, keyShareGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.GetKeyGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeyGroup`: KeyGroup
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.GetKeyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID. | 
**keyShareGroupId** | **string** | The key share group ID. | 

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

Get project information



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
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The project ID.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.GetMpcProject(ctx, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.GetMpcProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMpcProject`: MPCProject
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.GetMpcProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**projectId** | **string** | The project ID. | 

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

Get vault information



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The vault ID.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.GetMpcVault(ctx, vaultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.GetMpcVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMpcVault`: MPCVault
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.GetMpcVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID. | 

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

Get TSS request



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The vault ID.
	tssRequestId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The TSS request ID.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.GetTssRequest(ctx, vaultId, tssRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.GetTssRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTssRequest`: TSSRequest
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.GetTssRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID. | 
**tssRequestId** | **string** | The TSS request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTssRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TSSRequest**](TSSRequest.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCoboKeyHolder

> []KeyHolder ListCoboKeyHolder(ctx).Execute()

List all Cobo key share holders



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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.ListCoboKeyHolder(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.ListCoboKeyHolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCoboKeyHolder`: []KeyHolder
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.ListCoboKeyHolder`: %v\n", resp)
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

List all key share groups



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The vault ID.
	keyGroupType := CoboWaas2.KeyGroupType("MainGroup") // KeyGroupType | The selected key share group type to retrieve. Possible values include: - `MainKeyGroup`: Only [Main Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups) will be retrieved.  - `SigningKeyGroup`: Only [Signing Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups) will be retrieved.  - `RecoveryKeyGroup`: Only [Recovery Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups) will be retrieved.  **Note:** If left empty, all key share group types will be retrieved.  (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.ListKeyGroup(ctx, vaultId).KeyGroupType(keyGroupType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.ListKeyGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKeyGroup`: []KeyGroup
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.ListKeyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKeyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyGroupType** | [**KeyGroupType**](KeyGroupType.md) | The selected key share group type to retrieve. Possible values include: - &#x60;MainKeyGroup&#x60;: Only [Main Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups) will be retrieved.  - &#x60;SigningKeyGroup&#x60;: Only [Signing Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups) will be retrieved.  - &#x60;RecoveryKeyGroup&#x60;: Only [Recovery Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups) will be retrieved.  **Note:** If left empty, all key share group types will be retrieved.  | 

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

List all projects



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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.ListMpcProject(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.ListMpcProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMpcProject`: []MPCProject
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.ListMpcProject`: %v\n", resp)
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

List all vaults



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
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The project ID. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.ListMpcVault(ctx).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.ListMpcVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMpcVault`: []MPCVault
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.ListMpcVault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMpcVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | The project ID. | 

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

> []TSSRequest ListTssRequest(ctx, vaultId).TargetKeyGroupId(targetKeyGroupId).Execute()

List TSS requests



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The vault ID.
	targetKeyGroupId := "880311524363903326" // string | The target key share group ID of the TSS request.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.ListTssRequest(ctx, vaultId).TargetKeyGroupId(targetKeyGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.ListTssRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTssRequest`: []TSSRequest
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.ListTssRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTssRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetKeyGroupId** | **string** | The target key share group ID of the TSS request. | 

### Return type

[**[]TSSRequest**](TSSRequest.md)

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

Update vault name



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The vault ID.
	modifyMpcVaultRequest := *CoboWaas2.NewModifyMpcVaultRequest("The new name of the vault") // ModifyMpcVaultRequest | The request body to update a vault's name. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.ModifyMpcVault(ctx, vaultId).ModifyMpcVaultRequest(modifyMpcVaultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.ModifyMpcVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMpcVault`: MPCVault
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.ModifyMpcVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMpcVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyMpcVaultRequest** | [**ModifyMpcVaultRequest**](ModifyMpcVaultRequest.md) | The request body to update a vault&#39;s name. | 

### Return type

[**MPCVault**](MPCVault.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKeyGroup

> KeyGroup UpdateKeyGroup(ctx, vaultId, keyShareGroupId).UpdateKeyGroupRequest(updateKeyGroupRequest).Execute()

Update key share group



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The vault ID.
	keyShareGroupId := "880311524363903326" // string | The key share group ID.
	updateKeyGroupRequest := *CoboWaas2.NewUpdateKeyGroupRequest("UpgradeToMainGroup") // UpdateKeyGroupRequest | The available actions of key share group update. Possible values include: - `UpgradeToMainGroup`: This changes the specified key share group's type to [Main Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups).  (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.UpdateKeyGroup(ctx, vaultId, keyShareGroupId).UpdateKeyGroupRequest(updateKeyGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.UpdateKeyGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKeyGroup`: KeyGroup
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.UpdateKeyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID. | 
**keyShareGroupId** | **string** | The key share group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateKeyGroupRequest** | [**UpdateKeyGroupRequest**](UpdateKeyGroupRequest.md) | The available actions of key share group update. Possible values include: - &#x60;UpgradeToMainGroup&#x60;: This changes the specified key share group&#39;s type to [Main Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups).  | 

### Return type

[**KeyGroup**](KeyGroup.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMpcProject

> MPCProject UpdateMpcProject(ctx, projectId).UpdateMpcProjectRequest(updateMpcProjectRequest).Execute()

Update project name



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
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The project ID.
	updateMpcProjectRequest := *CoboWaas2.NewUpdateMpcProjectRequest("New project name") // UpdateMpcProjectRequest | The request body to update a project's name. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.UpdateMpcProject(ctx, projectId).UpdateMpcProjectRequest(updateMpcProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.UpdateMpcProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMpcProject`: MPCProject
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.UpdateMpcProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**projectId** | **string** | The project ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMpcProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMpcProjectRequest** | [**UpdateMpcProjectRequest**](UpdateMpcProjectRequest.md) | The request body to update a project&#39;s name. | 

### Return type

[**MPCProject**](MPCProject.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

