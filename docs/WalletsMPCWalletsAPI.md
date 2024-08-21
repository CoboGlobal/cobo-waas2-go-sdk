# \WalletsMPCWalletsAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTssRequestById**](WalletsMPCWalletsAPI.md#CancelTssRequestById) | **Post** /wallets/mpc/vaults/{vault_id}/tss_requests/{tss_request_id}/cancel | Cancel TSS request
[**CreateKeyShareHolderGroup**](WalletsMPCWalletsAPI.md#CreateKeyShareHolderGroup) | **Post** /wallets/mpc/vaults/{vault_id}/key_share_holder_groups | Create key share holder group
[**CreateMpcProject**](WalletsMPCWalletsAPI.md#CreateMpcProject) | **Post** /wallets/mpc/projects | Create project
[**CreateMpcVault**](WalletsMPCWalletsAPI.md#CreateMpcVault) | **Post** /wallets/mpc/vaults | Create vault
[**CreateTssRequest**](WalletsMPCWalletsAPI.md#CreateTssRequest) | **Post** /wallets/mpc/vaults/{vault_id}/tss_requests | Create TSS request
[**DeleteKeyShareHolderGroupById**](WalletsMPCWalletsAPI.md#DeleteKeyShareHolderGroupById) | **Post** /wallets/mpc/vaults/{vault_id}/key_share_holder_groups/{key_share_holder_group_id}/delete | Delete key share holder group
[**GetKeyShareHolderGroupById**](WalletsMPCWalletsAPI.md#GetKeyShareHolderGroupById) | **Get** /wallets/mpc/vaults/{vault_id}/key_share_holder_groups/{key_share_holder_group_id} | Get key share holder group information
[**GetMpcProjectById**](WalletsMPCWalletsAPI.md#GetMpcProjectById) | **Get** /wallets/mpc/projects/{project_id} | Get project information
[**GetMpcVaultById**](WalletsMPCWalletsAPI.md#GetMpcVaultById) | **Get** /wallets/mpc/vaults/{vault_id} | Get vault information
[**GetTssRequestById**](WalletsMPCWalletsAPI.md#GetTssRequestById) | **Get** /wallets/mpc/vaults/{vault_id}/tss_requests/{tss_request_id} | Get TSS request
[**ListCoboKeyHolders**](WalletsMPCWalletsAPI.md#ListCoboKeyHolders) | **Get** /wallets/mpc/cobo_key_share_holders | List all Cobo key share holders
[**ListKeyShareHolderGroups**](WalletsMPCWalletsAPI.md#ListKeyShareHolderGroups) | **Get** /wallets/mpc/vaults/{vault_id}/key_share_holder_groups | List all key share holder groups
[**ListMpcProjects**](WalletsMPCWalletsAPI.md#ListMpcProjects) | **Get** /wallets/mpc/projects | List all projects
[**ListMpcVaults**](WalletsMPCWalletsAPI.md#ListMpcVaults) | **Get** /wallets/mpc/vaults | List all vaults
[**ListTssRequests**](WalletsMPCWalletsAPI.md#ListTssRequests) | **Get** /wallets/mpc/vaults/{vault_id}/tss_requests | List TSS requests
[**UpdateKeyShareHolderGroupById**](WalletsMPCWalletsAPI.md#UpdateKeyShareHolderGroupById) | **Put** /wallets/mpc/vaults/{vault_id}/key_share_holder_groups/{key_share_holder_group_id} | Update key share holder group
[**UpdateMpcProjectById**](WalletsMPCWalletsAPI.md#UpdateMpcProjectById) | **Put** /wallets/mpc/projects/{project_id} | Update project name
[**UpdateMpcVaultById**](WalletsMPCWalletsAPI.md#UpdateMpcVaultById) | **Put** /wallets/mpc/vaults/{vault_id} | Update vault name



## CancelTssRequestById

> TSSRequest CancelTssRequestById(ctx, vaultId, tssRequestId).Execute()

Cancel TSS request



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	tssRequestId := "20240711114129000132315000003970"

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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.CancelTssRequestById(ctx, vaultId, tssRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.CancelTssRequestById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelTssRequestById`: TSSRequest
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.CancelTssRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID, which you can retrieve by calling [List all vaults](/v2/api-references/wallets--mpc-wallet/list-all-mpc-vaults). | 
**tssRequestId** | **string** | The TSS request ID, which you can retrieve by calling [List TSS requests](/v2/api-references/wallets--mpc-wallets/list-tss-requests). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelTssRequestByIdRequest struct via the builder pattern


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


## CreateKeyShareHolderGroup

> KeyShareHolderGroup CreateKeyShareHolderGroup(ctx, vaultId).CreateKeyShareHolderGroupRequest(createKeyShareHolderGroupRequest).Execute()

Create key share holder group



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	createKeyShareHolderGroupRequest := *coboWaas2.NewCreateKeyShareHolderGroupRequest(coboWaas2.KeyShareHolderGroupType("MainGroup"), int32(3), int32(2), []coboWaas2.CreateKeyShareHolder{*coboWaas2.NewCreateKeyShareHolder()})

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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.CreateKeyShareHolderGroup(ctx, vaultId).CreateKeyShareHolderGroupRequest(createKeyShareHolderGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.CreateKeyShareHolderGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKeyShareHolderGroup`: KeyShareHolderGroup
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.CreateKeyShareHolderGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID, which you can retrieve by calling [List all vaults](/v2/api-references/wallets--mpc-wallet/list-all-mpc-vaults). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyShareHolderGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createKeyShareHolderGroupRequest** | [**CreateKeyShareHolderGroupRequest**](CreateKeyShareHolderGroupRequest.md) | The request body to create a key share holder group. | 

### Return type

[**KeyShareHolderGroup**](KeyShareHolderGroup.md)

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
    coboWaas2 "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
    "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2/crypto"
)

func main() {
	createMpcProjectRequest := *coboWaas2.NewCreateMpcProjectRequest("Project name", int32(3), int32(2))

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
    coboWaas2 "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
    "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2/crypto"
)

func main() {
	createMpcVaultRequest := *coboWaas2.NewCreateMpcVaultRequest("My vault", coboWaas2.MPCVaultType("Org-Controlled"))

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
    coboWaas2 "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
    "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2/crypto"
)

func main() {
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	createTssRequestRequest := *coboWaas2.NewCreateTssRequestRequest(coboWaas2.TSSRequestType("KeyGen"), "a1bf161f-8b60-4f61-9c35-6434b8654437")

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
**vaultId** | **string** | The vault ID, which you can retrieve by calling [List all vaults](/v2/api-references/wallets--mpc-wallet/list-all-mpc-vaults). | 

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


## DeleteKeyShareHolderGroupById

> DeleteKeyShareHolderGroupById201Response DeleteKeyShareHolderGroupById(ctx, vaultId, keyShareHolderGroupId).Execute()

Delete key share holder group



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	keyShareHolderGroupId := "e8257ac8-76b8-4d1e-a1f9-eec4cb931dce"

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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.DeleteKeyShareHolderGroupById(ctx, vaultId, keyShareHolderGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.DeleteKeyShareHolderGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKeyShareHolderGroupById`: DeleteKeyShareHolderGroupById201Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.DeleteKeyShareHolderGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID, which you can retrieve by calling [List all vaults](/v2/api-references/wallets--mpc-wallet/list-all-mpc-vaults). | 
**keyShareHolderGroupId** | **string** | The key share holder group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyShareHolderGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteKeyShareHolderGroupById201Response**](DeleteKeyShareHolderGroupById201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyShareHolderGroupById

> KeyShareHolderGroup GetKeyShareHolderGroupById(ctx, vaultId, keyShareHolderGroupId).Execute()

Get key share holder group information



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	keyShareHolderGroupId := "e8257ac8-76b8-4d1e-a1f9-eec4cb931dce"

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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.GetKeyShareHolderGroupById(ctx, vaultId, keyShareHolderGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.GetKeyShareHolderGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeyShareHolderGroupById`: KeyShareHolderGroup
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.GetKeyShareHolderGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID, which you can retrieve by calling [List all vaults](/v2/api-references/wallets--mpc-wallet/list-all-mpc-vaults). | 
**keyShareHolderGroupId** | **string** | The key share holder group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyShareHolderGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KeyShareHolderGroup**](KeyShareHolderGroup.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMpcProjectById

> MPCProject GetMpcProjectById(ctx, projectId).Execute()

Get project information



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
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"

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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.GetMpcProjectById(ctx, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.GetMpcProjectById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMpcProjectById`: MPCProject
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.GetMpcProjectById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**projectId** | **string** | The project ID, which you can retrieve by calling [List all projects](/v2/api-references/wallets--mpc-wallets/list-all-projects). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMpcProjectByIdRequest struct via the builder pattern


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


## GetMpcVaultById

> MPCVault GetMpcVaultById(ctx, vaultId).Execute()

Get vault information



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"

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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.GetMpcVaultById(ctx, vaultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.GetMpcVaultById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMpcVaultById`: MPCVault
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.GetMpcVaultById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID, which you can retrieve by calling [List all vaults](/v2/api-references/wallets--mpc-wallet/list-all-mpc-vaults). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMpcVaultByIdRequest struct via the builder pattern


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


## GetTssRequestById

> TSSRequest GetTssRequestById(ctx, vaultId, tssRequestId).Execute()

Get TSS request



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	tssRequestId := "20240711114129000132315000003970"

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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.GetTssRequestById(ctx, vaultId, tssRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.GetTssRequestById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTssRequestById`: TSSRequest
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.GetTssRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID, which you can retrieve by calling [List all vaults](/v2/api-references/wallets--mpc-wallet/list-all-mpc-vaults). | 
**tssRequestId** | **string** | The TSS request ID, which you can retrieve by calling [List TSS requests](/v2/api-references/wallets--mpc-wallets/list-tss-requests). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTssRequestByIdRequest struct via the builder pattern


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


## ListCoboKeyHolders

> []KeyShareHolder ListCoboKeyHolders(ctx).Execute()

List all Cobo key share holders



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
	// Initialize the API client
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.Background()

    // Select the development environment. To use the production environment, replace coboWaas2.DevEnv with coboWaas2.ProdEnv
	ctx = context.WithValue(ctx, coboWaas2.ContextEnv, coboWaas2.DevEnv)
    // Replace `<YOUR_PRIVATE_KEY>` with your private key
	ctx = context.WithValue(ctx, coboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_PRIVATE_KEY>",
	})
	resp, r, err := apiClient.WalletsMPCWalletsAPI.ListCoboKeyHolders(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.ListCoboKeyHolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCoboKeyHolders`: []KeyShareHolder
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.ListCoboKeyHolders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCoboKeyHoldersRequest struct via the builder pattern


### Return type

[**[]KeyShareHolder**](KeyShareHolder.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeyShareHolderGroups

> ListKeyShareHolderGroups200Response ListKeyShareHolderGroups(ctx, vaultId).KeyShareHolderGroupType(keyShareHolderGroupType).Limit(limit).Before(before).After(after).Execute()

List all key share holder groups



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	keyShareHolderGroupType := coboWaas2.KeyShareHolderGroupType("MainGroup")
	limit := int32(10)
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1"
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk"

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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.ListKeyShareHolderGroups(ctx, vaultId).KeyShareHolderGroupType(keyShareHolderGroupType).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.ListKeyShareHolderGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKeyShareHolderGroups`: ListKeyShareHolderGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.ListKeyShareHolderGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID, which you can retrieve by calling [List all vaults](/v2/api-references/wallets--mpc-wallet/list-all-mpc-vaults). | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKeyShareHolderGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyShareHolderGroupType** | [**KeyShareHolderGroupType**](KeyShareHolderGroupType.md) | The key share holder group type. Possible values include: - &#x60;MainGroup&#x60;: The [Main Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups#main-group).  - &#x60;SigningGroup&#x60;: The [Signing Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups#signing-group).  - &#x60;RecoveryGroup&#x60;: The [Recovery Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups#recovery-group).  **Note**: If this parameter is left empty, all key share holder group types will be retrieved.  | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify &#x60;before&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  - If you set &#x60;before&#x60; to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify &#x60;after&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 

### Return type

[**ListKeyShareHolderGroups200Response**](ListKeyShareHolderGroups200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMpcProjects

> ListMpcProjects200Response ListMpcProjects(ctx).Limit(limit).Before(before).After(after).Execute()

List all projects



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
	limit := int32(10)
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1"
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk"

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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.ListMpcProjects(ctx).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.ListMpcProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMpcProjects`: ListMpcProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.ListMpcProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMpcProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify &#x60;before&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  - If you set &#x60;before&#x60; to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify &#x60;after&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 

### Return type

[**ListMpcProjects200Response**](ListMpcProjects200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMpcVaults

> ListMpcVaults200Response ListMpcVaults(ctx).VaultType(vaultType).ProjectId(projectId).Limit(limit).Before(before).After(after).Execute()

List all vaults



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
	vaultType := coboWaas2.MPCVaultType("Org-Controlled")
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	limit := int32(10)
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1"
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk"

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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.ListMpcVaults(ctx).VaultType(vaultType).ProjectId(projectId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.ListMpcVaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMpcVaults`: ListMpcVaults200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.ListMpcVaults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMpcVaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vaultType** | [**MPCVaultType**](MPCVaultType.md) | The vault type. Possible values include: - &#x60;Org-Controlled&#x60;: This vault is a collection of [Organization-Controlled Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#organization-controlled-wallets).  - &#x60;User-Controlled&#x60;: This vault is a collection of [User-Controlled Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#user-controlled-wallets).  | 
 **projectId** | **string** | The project ID, which you can retrieve by calling [List all projects](/v2/api-references/wallets--mpc-wallets/list-all-projects).  | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify &#x60;before&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  - If you set &#x60;before&#x60; to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify &#x60;after&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 

### Return type

[**ListMpcVaults200Response**](ListMpcVaults200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTssRequests

> ListTssRequests200Response ListTssRequests(ctx, vaultId).KeyShareHolderGroupId(keyShareHolderGroupId).Limit(limit).Before(before).After(after).Execute()

List TSS requests



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	keyShareHolderGroupId := "a3a45e99-5a12-444f-867a-ffe0ebb1bb30"
	limit := int32(10)
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1"
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk"

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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.ListTssRequests(ctx, vaultId).KeyShareHolderGroupId(keyShareHolderGroupId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.ListTssRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTssRequests`: ListTssRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.ListTssRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID, which you can retrieve by calling [List all vaults](/v2/api-references/wallets--mpc-wallet/list-all-mpc-vaults). | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTssRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyShareHolderGroupId** | **string** | The key share holder group ID of the TSS request, which you can retrieve by calling [List all key share holder groups](/v2/api-references/wallets--mpc-wallets/list-all-key-share-holder-groups). | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify &#x60;before&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  - If you set &#x60;before&#x60; to &#x60;infinity&#x60;, the last page of data is returned.  | 
 **after** | **string** | An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify &#x60;after&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  | 

### Return type

[**ListTssRequests200Response**](ListTssRequests200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKeyShareHolderGroupById

> KeyShareHolderGroup UpdateKeyShareHolderGroupById(ctx, vaultId, keyShareHolderGroupId).UpdateKeyShareHolderGroupByIdRequest(updateKeyShareHolderGroupByIdRequest).Execute()

Update key share holder group



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	keyShareHolderGroupId := "e8257ac8-76b8-4d1e-a1f9-eec4cb931dce"
	updateKeyShareHolderGroupByIdRequest := *coboWaas2.NewUpdateKeyShareHolderGroupByIdRequest(coboWaas2.UpdateGroupAction("UpgradeToMainGroup"))

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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.UpdateKeyShareHolderGroupById(ctx, vaultId, keyShareHolderGroupId).UpdateKeyShareHolderGroupByIdRequest(updateKeyShareHolderGroupByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.UpdateKeyShareHolderGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKeyShareHolderGroupById`: KeyShareHolderGroup
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.UpdateKeyShareHolderGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID, which you can retrieve by calling [List all vaults](/v2/api-references/wallets--mpc-wallet/list-all-mpc-vaults). | 
**keyShareHolderGroupId** | **string** | The key share holder group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyShareHolderGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateKeyShareHolderGroupByIdRequest** | [**UpdateKeyShareHolderGroupByIdRequest**](UpdateKeyShareHolderGroupByIdRequest.md) |  | 

### Return type

[**KeyShareHolderGroup**](KeyShareHolderGroup.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMpcProjectById

> MPCProject UpdateMpcProjectById(ctx, projectId).UpdateMpcProjectByIdRequest(updateMpcProjectByIdRequest).Execute()

Update project name



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
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	updateMpcProjectByIdRequest := *coboWaas2.NewUpdateMpcProjectByIdRequest("New project name")

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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.UpdateMpcProjectById(ctx, projectId).UpdateMpcProjectByIdRequest(updateMpcProjectByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.UpdateMpcProjectById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMpcProjectById`: MPCProject
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.UpdateMpcProjectById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**projectId** | **string** | The project ID, which you can retrieve by calling [List all projects](/v2/api-references/wallets--mpc-wallets/list-all-projects). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMpcProjectByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMpcProjectByIdRequest** | [**UpdateMpcProjectByIdRequest**](UpdateMpcProjectByIdRequest.md) | The request body to update a project&#39;s name. | 

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


## UpdateMpcVaultById

> MPCVault UpdateMpcVaultById(ctx, vaultId).UpdateMpcVaultByIdRequest(updateMpcVaultByIdRequest).Execute()

Update vault name



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
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	updateMpcVaultByIdRequest := *coboWaas2.NewUpdateMpcVaultByIdRequest("The new name of the vault")

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
	resp, r, err := apiClient.WalletsMPCWalletsAPI.UpdateMpcVaultById(ctx, vaultId).UpdateMpcVaultByIdRequest(updateMpcVaultByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsMPCWalletsAPI.UpdateMpcVaultById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMpcVaultById`: MPCVault
	fmt.Fprintf(os.Stdout, "Response from `WalletsMPCWalletsAPI.UpdateMpcVaultById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**vaultId** | **string** | The vault ID, which you can retrieve by calling [List all vaults](/v2/api-references/wallets--mpc-wallet/list-all-mpc-vaults). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMpcVaultByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMpcVaultByIdRequest** | [**UpdateMpcVaultByIdRequest**](UpdateMpcVaultByIdRequest.md) | The request body to update a vault&#39;s name. | 

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

