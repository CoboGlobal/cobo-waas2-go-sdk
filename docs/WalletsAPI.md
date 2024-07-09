# \WalletsAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAddressValidity**](WalletsAPI.md#CheckAddressValidity) | **Get** /wallets/address/check_validity | Check address validity
[**CreateWallet**](WalletsAPI.md#CreateWallet) | **Post** /wallets | Create wallet
[**DeleteWalletById**](WalletsAPI.md#DeleteWalletById) | **Post** /wallets/{wallet_id}/delete | Delete wallet
[**GenerateWalletAddress**](WalletsAPI.md#GenerateWalletAddress) | **Post** /wallets/{wallet_id}/addresses | Generate new addresses in wallet
[**GetAddressById**](WalletsAPI.md#GetAddressById) | **Get** /wallets/{wallet_id}/addresses/{address} | Get address information
[**GetChainById**](WalletsAPI.md#GetChainById) | **Get** /wallets/chains/{chain_id} | Get chain information
[**GetChains**](WalletsAPI.md#GetChains) | **Get** /wallets/chains | List all supported chains
[**GetEnabledChains**](WalletsAPI.md#GetEnabledChains) | **Get** /wallets/enabled_chains | List organization enabled chains
[**GetEnabledTokens**](WalletsAPI.md#GetEnabledTokens) | **Get** /wallets/enabled_tokens | List organization enabled tokens
[**GetMaxTransferableValue**](WalletsAPI.md#GetMaxTransferableValue) | **Get** /wallets/{wallet_id}/max_transferable_value | Get maximum transferable value
[**GetSpendableList**](WalletsAPI.md#GetSpendableList) | **Get** /wallets/{wallet_id}/spendables | List spendable UTXOs
[**GetSupportedChains**](WalletsAPI.md#GetSupportedChains) | **Get** /wallets/supported_chains | List wallet supported chains
[**GetSupportedTokens**](WalletsAPI.md#GetSupportedTokens) | **Get** /wallets/supported_tokens | List wallet supported tokens
[**GetTokenById**](WalletsAPI.md#GetTokenById) | **Get** /wallets/tokens/{token_id} | Get token information
[**GetTokens**](WalletsAPI.md#GetTokens) | **Get** /wallets/tokens | List all supported tokens
[**GetWalletAddressTokenBalances**](WalletsAPI.md#GetWalletAddressTokenBalances) | **Get** /wallets/{wallet_id}/addresses/{address}/tokens | List token balances by address
[**GetWalletById**](WalletsAPI.md#GetWalletById) | **Get** /wallets/{wallet_id} | Get wallet information
[**GetWalletTokenBalances**](WalletsAPI.md#GetWalletTokenBalances) | **Get** /wallets/{wallet_id}/tokens | List token balances by wallet
[**ListAddresses**](WalletsAPI.md#ListAddresses) | **Get** /wallets/{wallet_id}/addresses | List wallet addresses
[**ListWallets**](WalletsAPI.md#ListWallets) | **Get** /wallets | List all wallets
[**LockSpendableList**](WalletsAPI.md#LockSpendableList) | **Post** /wallets/{wallet_id}/spendables/lock | Lock UTXOs
[**UnlockSpendableList**](WalletsAPI.md#UnlockSpendableList) | **Post** /wallets/{wallet_id}/spendables/unlock | Unlock UTXOs
[**UpdateWalletById**](WalletsAPI.md#UpdateWalletById) | **Put** /wallets/{wallet_id} | Update wallet



## CheckAddressValidity

> CheckAddressValidity200Response CheckAddressValidity(ctx).TokenId(tokenId).Address(address).Execute()

Check address validity



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
	tokenId := "ETH_USDT" // string | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens).
	address := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045" // string | The wallet address.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.CheckAddressValidity(ctx).TokenId(tokenId).Address(address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CheckAddressValidity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckAddressValidity`: CheckAddressValidity200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CheckAddressValidity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAddressValidityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
 **address** | **string** | The wallet address. | 

### Return type

[**CheckAddressValidity200Response**](CheckAddressValidity200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWallet

> WalletInfo CreateWallet(ctx).CreatedWallet(createdWallet).Execute()

Create wallet



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
	createdWallet := CoboWaas2.CreatedWallet{CreateCustodialWallet: CoboWaas2.NewCreateCustodialWallet("My WaaS 2.0 Wallet", CoboWaas2.WalletType("Custodial"), CoboWaas2.WalletSubtype("Asset"))} // CreatedWallet | The request body to create a wallet (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.CreateWallet(ctx).CreatedWallet(createdWallet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CreateWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWallet`: WalletInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CreateWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdWallet** | [**CreatedWallet**](CreatedWallet.md) | The request body to create a wallet | 

### Return type

[**WalletInfo**](WalletInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWalletById

> DeleteWalletById200Response DeleteWalletById(ctx, walletId).Execute()

Delete wallet



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.DeleteWalletById(ctx, walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.DeleteWalletById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWalletById`: DeleteWalletById200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.DeleteWalletById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWalletByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteWalletById200Response**](DeleteWalletById200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateWalletAddress

> []AddressInfo GenerateWalletAddress(ctx, walletId).GenerateWalletAddressRequest(generateWalletAddressRequest).Execute()

Generate new addresses in wallet



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.
	generateWalletAddressRequest := *CoboWaas2.NewGenerateWalletAddressRequest("ETH_USDT", int32(1)) // GenerateWalletAddressRequest | The request body to generates addresses within a specified wallet. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GenerateWalletAddress(ctx, walletId).GenerateWalletAddressRequest(generateWalletAddressRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GenerateWalletAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateWalletAddress`: []AddressInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GenerateWalletAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateWalletAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateWalletAddressRequest** | [**GenerateWalletAddressRequest**](GenerateWalletAddressRequest.md) | The request body to generates addresses within a specified wallet. | 

### Return type

[**[]AddressInfo**](AddressInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressById

> ListAddresses200Response GetAddressById(ctx, walletId, address).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()

Get address information



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.
	address := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045" // string | The wallet address.
	tokenId := "ETH_USDT" // string | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that end before the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` and `after` are both set to empty, the first slice is returned.  (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that start after the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` is set to empty and `after` is set to `last`, the last slice is returned.  (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetAddressById(ctx, walletId, address).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetAddressById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressById`: ListAddresses200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetAddressById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 
**address** | **string** | The wallet address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**ListAddresses200Response**](ListAddresses200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChainById

> ChainInfo GetChainById(ctx, chainId).Execute()

Get chain information



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
	chainId := "ETH" // string | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains).

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetChainById(ctx, chainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetChainById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChainById`: ChainInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetChainById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**chainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChainByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChainInfo**](ChainInfo.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChains

> GetChains200Response GetChains(ctx).ChainIds(chainIds).Limit(limit).Before(before).After(after).Execute()

List all supported chains



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
	chainIds := "BTC,ETH" // string | A list of chain IDs, separated by comma. The chain ID is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that end before the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` and `after` are both set to empty, the first slice is returned.  (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that start after the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` is set to empty and `after` is set to `last`, the last slice is returned.  (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetChains(ctx).ChainIds(chainIds).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetChains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChains`: GetChains200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetChains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chainIds** | **string** | A list of chain IDs, separated by comma. The chain ID is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**GetChains200Response**](GetChains200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnabledChains

> GetChains200Response GetEnabledChains(ctx).WalletType(walletType).WalletSubtype(walletSubtype).Limit(limit).Before(before).After(after).Execute()

List organization enabled chains



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
	walletType := CoboWaas2.WalletType("Custodial") // WalletType | The wallet type.  - `Custodial`: [Custodial Wallets](https://manuals.cobo.com/en/portal/custodial-wallets/introduction)  - `MPC`: [MPC Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction)  - `SmartContract`: [Smart Contract Wallets](https://manuals.cobo.com/en/portal/smart-contract-wallets/introduction)  - `Exchange`: [Exchange Wallets](https://manuals.cobo.com/en/portal/exchange-wallets/introduction)  (optional)
	walletSubtype := CoboWaas2.WalletSubtype("Asset") // WalletSubtype | The wallet subtype.  - `Asset`: Custodial Wallets (Asset Wallets)  - `Web3`: Custodial Wallets (Web3 Wallets)  - `Main`: Exchange Wallets (Main Account)  - `Sub`: Exchange Wallets (Sub Account)  - `Org-Controlled`: MPC Wallets (Organization-Controlled Wallets)  - `User-Controlled`: MPC Wallets (User-Controlled Wallets)  - `Safe{Wallet}`: Smart Contract Wallets (Safe{Wallet})  (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that end before the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` and `after` are both set to empty, the first slice is returned.  (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that start after the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` is set to empty and `after` is set to `last`, the last slice is returned.  (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetEnabledChains(ctx).WalletType(walletType).WalletSubtype(walletSubtype).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetEnabledChains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnabledChains`: GetChains200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetEnabledChains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnabledChainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletType** | [**WalletType**](WalletType.md) | The wallet type.  - &#x60;Custodial&#x60;: [Custodial Wallets](https://manuals.cobo.com/en/portal/custodial-wallets/introduction)  - &#x60;MPC&#x60;: [MPC Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction)  - &#x60;SmartContract&#x60;: [Smart Contract Wallets](https://manuals.cobo.com/en/portal/smart-contract-wallets/introduction)  - &#x60;Exchange&#x60;: [Exchange Wallets](https://manuals.cobo.com/en/portal/exchange-wallets/introduction)  | 
 **walletSubtype** | [**WalletSubtype**](WalletSubtype.md) | The wallet subtype.  - &#x60;Asset&#x60;: Custodial Wallets (Asset Wallets)  - &#x60;Web3&#x60;: Custodial Wallets (Web3 Wallets)  - &#x60;Main&#x60;: Exchange Wallets (Main Account)  - &#x60;Sub&#x60;: Exchange Wallets (Sub Account)  - &#x60;Org-Controlled&#x60;: MPC Wallets (Organization-Controlled Wallets)  - &#x60;User-Controlled&#x60;: MPC Wallets (User-Controlled Wallets)  - &#x60;Safe{Wallet}&#x60;: Smart Contract Wallets (Safe{Wallet})  | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**GetChains200Response**](GetChains200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnabledTokens

> GetTokens200Response GetEnabledTokens(ctx).WalletType(walletType).WalletSubtype(walletSubtype).ChainId(chainId).Limit(limit).Before(before).After(after).Execute()

List organization enabled tokens



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
	walletType := CoboWaas2.WalletType("Custodial") // WalletType | The wallet type.  - `Custodial`: [Custodial Wallets](https://manuals.cobo.com/en/portal/custodial-wallets/introduction)  - `MPC`: [MPC Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction)  - `SmartContract`: [Smart Contract Wallets](https://manuals.cobo.com/en/portal/smart-contract-wallets/introduction)  - `Exchange`: [Exchange Wallets](https://manuals.cobo.com/en/portal/exchange-wallets/introduction)  (optional)
	walletSubtype := CoboWaas2.WalletSubtype("Asset") // WalletSubtype | The wallet subtype.  - `Asset`: Custodial Wallets (Asset Wallets)  - `Web3`: Custodial Wallets (Web3 Wallets)  - `Main`: Exchange Wallets (Main Account)  - `Sub`: Exchange Wallets (Sub Account)  - `Org-Controlled`: MPC Wallets (Organization-Controlled Wallets)  - `User-Controlled`: MPC Wallets (User-Controlled Wallets)  - `Safe{Wallet}`: Smart Contract Wallets (Safe{Wallet})  (optional)
	chainId := "ETH" // string | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that end before the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` and `after` are both set to empty, the first slice is returned.  (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that start after the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` is set to empty and `after` is set to `last`, the last slice is returned.  (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetEnabledTokens(ctx).WalletType(walletType).WalletSubtype(walletSubtype).ChainId(chainId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetEnabledTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnabledTokens`: GetTokens200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetEnabledTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnabledTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletType** | [**WalletType**](WalletType.md) | The wallet type.  - &#x60;Custodial&#x60;: [Custodial Wallets](https://manuals.cobo.com/en/portal/custodial-wallets/introduction)  - &#x60;MPC&#x60;: [MPC Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction)  - &#x60;SmartContract&#x60;: [Smart Contract Wallets](https://manuals.cobo.com/en/portal/smart-contract-wallets/introduction)  - &#x60;Exchange&#x60;: [Exchange Wallets](https://manuals.cobo.com/en/portal/exchange-wallets/introduction)  | 
 **walletSubtype** | [**WalletSubtype**](WalletSubtype.md) | The wallet subtype.  - &#x60;Asset&#x60;: Custodial Wallets (Asset Wallets)  - &#x60;Web3&#x60;: Custodial Wallets (Web3 Wallets)  - &#x60;Main&#x60;: Exchange Wallets (Main Account)  - &#x60;Sub&#x60;: Exchange Wallets (Sub Account)  - &#x60;Org-Controlled&#x60;: MPC Wallets (Organization-Controlled Wallets)  - &#x60;User-Controlled&#x60;: MPC Wallets (User-Controlled Wallets)  - &#x60;Safe{Wallet}&#x60;: Smart Contract Wallets (Safe{Wallet})  | 
 **chainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**GetTokens200Response**](GetTokens200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxTransferableValue

> MaxTransferableValue GetMaxTransferableValue(ctx, walletId).TokenId(tokenId).FeeRate(feeRate).ToAddress(toAddress).FromAddress(fromAddress).Execute()

Get maximum transferable value



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.
	tokenId := "ETH_USDT" // string | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens).
	feeRate := "0.001" // string | The fee rate in sats/vByte or fee_price in gwei.
	toAddress := "2N2xFZtbCFB6Nb3Pj9Sxsx5mX2fxX3yEgkE" // string | The recipient's address.
	fromAddress := "2N2xFZtbCFB6Nb3Pj9Sxsx5mX2fxX3yEgkE" // string | The sender's address. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetMaxTransferableValue(ctx, walletId).TokenId(tokenId).FeeRate(feeRate).ToAddress(toAddress).FromAddress(fromAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetMaxTransferableValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaxTransferableValue`: MaxTransferableValue
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetMaxTransferableValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxTransferableValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
 **feeRate** | **string** | The fee rate in sats/vByte or fee_price in gwei. | 
 **toAddress** | **string** | The recipient&#39;s address. | 
 **fromAddress** | **string** | The sender&#39;s address. | 

### Return type

[**MaxTransferableValue**](MaxTransferableValue.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpendableList

> GetSpendableList200Response GetSpendableList(ctx, walletId).TokenId(tokenId).Address(address).Execute()

List spendable UTXOs



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.
	tokenId := "ETH_USDT" // string | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens).
	address := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045" // string | The wallet address. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetSpendableList(ctx, walletId).TokenId(tokenId).Address(address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetSpendableList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpendableList`: GetSpendableList200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetSpendableList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpendableListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
 **address** | **string** | The wallet address. | 

### Return type

[**GetSpendableList200Response**](GetSpendableList200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedChains

> GetChains200Response GetSupportedChains(ctx).WalletType(walletType).WalletSubtype(walletSubtype).Limit(limit).Before(before).After(after).Execute()

List wallet supported chains



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
	walletType := CoboWaas2.WalletType("Custodial") // WalletType | The wallet type.  - `Custodial`: [Custodial Wallets](https://manuals.cobo.com/en/portal/custodial-wallets/introduction)  - `MPC`: [MPC Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction)  - `SmartContract`: [Smart Contract Wallets](https://manuals.cobo.com/en/portal/smart-contract-wallets/introduction)  - `Exchange`: [Exchange Wallets](https://manuals.cobo.com/en/portal/exchange-wallets/introduction)  (optional)
	walletSubtype := CoboWaas2.WalletSubtype("Asset") // WalletSubtype | The wallet subtype.  - `Asset`: Custodial Wallets (Asset Wallets)  - `Web3`: Custodial Wallets (Web3 Wallets)  - `Main`: Exchange Wallets (Main Account)  - `Sub`: Exchange Wallets (Sub Account)  - `Org-Controlled`: MPC Wallets (Organization-Controlled Wallets)  - `User-Controlled`: MPC Wallets (User-Controlled Wallets)  - `Safe{Wallet}`: Smart Contract Wallets (Safe{Wallet})  (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that end before the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` and `after` are both set to empty, the first slice is returned.  (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that start after the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` is set to empty and `after` is set to `last`, the last slice is returned.  (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetSupportedChains(ctx).WalletType(walletType).WalletSubtype(walletSubtype).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetSupportedChains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportedChains`: GetChains200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetSupportedChains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedChainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletType** | [**WalletType**](WalletType.md) | The wallet type.  - &#x60;Custodial&#x60;: [Custodial Wallets](https://manuals.cobo.com/en/portal/custodial-wallets/introduction)  - &#x60;MPC&#x60;: [MPC Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction)  - &#x60;SmartContract&#x60;: [Smart Contract Wallets](https://manuals.cobo.com/en/portal/smart-contract-wallets/introduction)  - &#x60;Exchange&#x60;: [Exchange Wallets](https://manuals.cobo.com/en/portal/exchange-wallets/introduction)  | 
 **walletSubtype** | [**WalletSubtype**](WalletSubtype.md) | The wallet subtype.  - &#x60;Asset&#x60;: Custodial Wallets (Asset Wallets)  - &#x60;Web3&#x60;: Custodial Wallets (Web3 Wallets)  - &#x60;Main&#x60;: Exchange Wallets (Main Account)  - &#x60;Sub&#x60;: Exchange Wallets (Sub Account)  - &#x60;Org-Controlled&#x60;: MPC Wallets (Organization-Controlled Wallets)  - &#x60;User-Controlled&#x60;: MPC Wallets (User-Controlled Wallets)  - &#x60;Safe{Wallet}&#x60;: Smart Contract Wallets (Safe{Wallet})  | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**GetChains200Response**](GetChains200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedTokens

> GetTokens200Response GetSupportedTokens(ctx).WalletType(walletType).WalletSubtype(walletSubtype).ChainId(chainId).Limit(limit).Before(before).After(after).Execute()

List wallet supported tokens



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
	walletType := CoboWaas2.WalletType("Custodial") // WalletType | The wallet type.  - `Custodial`: [Custodial Wallets](https://manuals.cobo.com/en/portal/custodial-wallets/introduction)  - `MPC`: [MPC Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction)  - `SmartContract`: [Smart Contract Wallets](https://manuals.cobo.com/en/portal/smart-contract-wallets/introduction)  - `Exchange`: [Exchange Wallets](https://manuals.cobo.com/en/portal/exchange-wallets/introduction)  (optional)
	walletSubtype := CoboWaas2.WalletSubtype("Asset") // WalletSubtype | The wallet subtype.  - `Asset`: Custodial Wallets (Asset Wallets)  - `Web3`: Custodial Wallets (Web3 Wallets)  - `Main`: Exchange Wallets (Main Account)  - `Sub`: Exchange Wallets (Sub Account)  - `Org-Controlled`: MPC Wallets (Organization-Controlled Wallets)  - `User-Controlled`: MPC Wallets (User-Controlled Wallets)  - `Safe{Wallet}`: Smart Contract Wallets (Safe{Wallet})  (optional)
	chainId := "ETH" // string | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that end before the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` and `after` are both set to empty, the first slice is returned.  (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that start after the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` is set to empty and `after` is set to `last`, the last slice is returned.  (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetSupportedTokens(ctx).WalletType(walletType).WalletSubtype(walletSubtype).ChainId(chainId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetSupportedTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportedTokens`: GetTokens200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetSupportedTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletType** | [**WalletType**](WalletType.md) | The wallet type.  - &#x60;Custodial&#x60;: [Custodial Wallets](https://manuals.cobo.com/en/portal/custodial-wallets/introduction)  - &#x60;MPC&#x60;: [MPC Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction)  - &#x60;SmartContract&#x60;: [Smart Contract Wallets](https://manuals.cobo.com/en/portal/smart-contract-wallets/introduction)  - &#x60;Exchange&#x60;: [Exchange Wallets](https://manuals.cobo.com/en/portal/exchange-wallets/introduction)  | 
 **walletSubtype** | [**WalletSubtype**](WalletSubtype.md) | The wallet subtype.  - &#x60;Asset&#x60;: Custodial Wallets (Asset Wallets)  - &#x60;Web3&#x60;: Custodial Wallets (Web3 Wallets)  - &#x60;Main&#x60;: Exchange Wallets (Main Account)  - &#x60;Sub&#x60;: Exchange Wallets (Sub Account)  - &#x60;Org-Controlled&#x60;: MPC Wallets (Organization-Controlled Wallets)  - &#x60;User-Controlled&#x60;: MPC Wallets (User-Controlled Wallets)  - &#x60;Safe{Wallet}&#x60;: Smart Contract Wallets (Safe{Wallet})  | 
 **chainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**GetTokens200Response**](GetTokens200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenById

> GetTokens200Response GetTokenById(ctx, tokenId).Execute()

Get token information



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
	tokenId := "ETH_USDT" // string | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens).

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetTokenById(ctx, tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetTokenById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenById`: GetTokens200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetTokenById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTokens200Response**](GetTokens200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokens

> GetTokens200Response GetTokens(ctx).TokenIds(tokenIds).Limit(limit).Before(before).After(after).Execute()

List all supported tokens



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
	tokenIds := "ETH_USDT,ETH_USDC" // string | A list of token IDs, separated by comma. The token ID is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that end before the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` and `after` are both set to empty, the first slice is returned.  (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that start after the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` is set to empty and `after` is set to `last`, the last slice is returned.  (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetTokens(ctx).TokenIds(tokenIds).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokens`: GetTokens200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenIds** | **string** | A list of token IDs, separated by comma. The token ID is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**GetTokens200Response**](GetTokens200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletAddressTokenBalances

> GetWalletTokenBalances200Response GetWalletAddressTokenBalances(ctx, walletId, address).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()

List token balances by address



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.
	address := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045" // string | The wallet address.
	tokenId := "ETH_USDT" // string | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that end before the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` and `after` are both set to empty, the first slice is returned.  (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that start after the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` is set to empty and `after` is set to `last`, the last slice is returned.  (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetWalletAddressTokenBalances(ctx, walletId, address).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletAddressTokenBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletAddressTokenBalances`: GetWalletTokenBalances200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletAddressTokenBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 
**address** | **string** | The wallet address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletAddressTokenBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**GetWalletTokenBalances200Response**](GetWalletTokenBalances200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletById

> WalletInfo GetWalletById(ctx, walletId).Execute()

Get wallet information



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetWalletById(ctx, walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletById`: WalletInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WalletInfo**](WalletInfo.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletTokenBalances

> GetWalletTokenBalances200Response GetWalletTokenBalances(ctx, walletId).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()

List token balances by wallet



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.
	tokenId := "ETH_USDT" // string | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that end before the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` and `after` are both set to empty, the first slice is returned.  (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that start after the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` is set to empty and `after` is set to `last`, the last slice is returned.  (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetWalletTokenBalances(ctx, walletId).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletTokenBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletTokenBalances`: GetWalletTokenBalances200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletTokenBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletTokenBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**GetWalletTokenBalances200Response**](GetWalletTokenBalances200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAddresses

> ListAddresses200Response ListAddresses(ctx, walletId).TokenId(tokenId).Addresses(addresses).Limit(limit).Before(before).After(after).Execute()

List wallet addresses



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.
	tokenId := "ETH_USDT" // string | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). (optional)
	addresses := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045,0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97" // string | A list of wallet addresses, separated by comma. (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that end before the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` and `after` are both set to empty, the first slice is returned.  (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that start after the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` is set to empty and `after` is set to `last`, the last slice is returned.  (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.ListAddresses(ctx, walletId).TokenId(tokenId).Addresses(addresses).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.ListAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAddresses`: ListAddresses200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.ListAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
 **addresses** | **string** | A list of wallet addresses, separated by comma. | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**ListAddresses200Response**](ListAddresses200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWallets

> ListWallets200Response ListWallets(ctx).WalletType(walletType).WalletSubtype(walletSubtype).ProjectId(projectId).VaultId(vaultId).Limit(limit).Before(before).After(after).Execute()

List all wallets



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
	walletType := CoboWaas2.WalletType("Custodial") // WalletType | The wallet type.  - `Custodial`: [Custodial Wallets](https://manuals.cobo.com/en/portal/custodial-wallets/introduction)  - `MPC`: [MPC Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction)  - `SmartContract`: [Smart Contract Wallets](https://manuals.cobo.com/en/portal/smart-contract-wallets/introduction)  - `Exchange`: [Exchange Wallets](https://manuals.cobo.com/en/portal/exchange-wallets/introduction)  (optional)
	walletSubtype := CoboWaas2.WalletSubtype("Asset") // WalletSubtype | The wallet subtype.  - `Asset`: Custodial Wallets (Asset Wallets)  - `Web3`: Custodial Wallets (Web3 Wallets)  - `Main`: Exchange Wallets (Main Account)  - `Sub`: Exchange Wallets (Sub Account)  - `Org-Controlled`: MPC Wallets (Organization-Controlled Wallets)  - `User-Controlled`: MPC Wallets (User-Controlled Wallets)  - `Safe{Wallet}`: Smart Contract Wallets (Safe{Wallet})  (optional)
	projectId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The project ID. (optional)
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The MPC vault ID. This parameter is applicable to MPC Wallets only. (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that end before the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.after` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` and `after` are both set to empty, the first slice is returned.  (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`, the request will retrieve a list of data objects that start after the object with the object ID `8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f`. You can set this parameter to the value of `pagination.before` in the response of the previous request.  If you set both `after` or `before`, only the setting of `before` will take effect.  If the `before` is set to empty and `after` is set to `last`, the last slice is returned.  (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.ListWallets(ctx).WalletType(walletType).WalletSubtype(walletSubtype).ProjectId(projectId).VaultId(vaultId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.ListWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWallets`: ListWallets200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.ListWallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletType** | [**WalletType**](WalletType.md) | The wallet type.  - &#x60;Custodial&#x60;: [Custodial Wallets](https://manuals.cobo.com/en/portal/custodial-wallets/introduction)  - &#x60;MPC&#x60;: [MPC Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction)  - &#x60;SmartContract&#x60;: [Smart Contract Wallets](https://manuals.cobo.com/en/portal/smart-contract-wallets/introduction)  - &#x60;Exchange&#x60;: [Exchange Wallets](https://manuals.cobo.com/en/portal/exchange-wallets/introduction)  | 
 **walletSubtype** | [**WalletSubtype**](WalletSubtype.md) | The wallet subtype.  - &#x60;Asset&#x60;: Custodial Wallets (Asset Wallets)  - &#x60;Web3&#x60;: Custodial Wallets (Web3 Wallets)  - &#x60;Main&#x60;: Exchange Wallets (Main Account)  - &#x60;Sub&#x60;: Exchange Wallets (Sub Account)  - &#x60;Org-Controlled&#x60;: MPC Wallets (Organization-Controlled Wallets)  - &#x60;User-Controlled&#x60;: MPC Wallets (User-Controlled Wallets)  - &#x60;Safe{Wallet}&#x60;: Smart Contract Wallets (Safe{Wallet})  | 
 **projectId** | **string** | The project ID. | 
 **vaultId** | **string** | The MPC vault ID. This parameter is applicable to MPC Wallets only. | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**ListWallets200Response**](ListWallets200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockSpendableList

> LockSpendableList200Response LockSpendableList(ctx, walletId).LockSpendableListRequest(lockSpendableListRequest).Execute()

Lock UTXOs



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.
	lockSpendableListRequest := *CoboWaas2.NewLockSpendableListRequest("BTC", "9bdf8e7ae03c237e115f09543fbdb40f8efa600106e78b67ce4d5adfadda2dbb", int32(0)) // LockSpendableListRequest | The request body of the Lock/Unlock UTXOs operation. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.LockSpendableList(ctx, walletId).LockSpendableListRequest(lockSpendableListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.LockSpendableList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LockSpendableList`: LockSpendableList200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.LockSpendableList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockSpendableListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockSpendableListRequest** | [**LockSpendableListRequest**](LockSpendableListRequest.md) | The request body of the Lock/Unlock UTXOs operation. | 

### Return type

[**LockSpendableList200Response**](LockSpendableList200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlockSpendableList

> LockSpendableList200Response UnlockSpendableList(ctx, walletId).LockSpendableListRequest(lockSpendableListRequest).Execute()

Unlock UTXOs



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.
	lockSpendableListRequest := *CoboWaas2.NewLockSpendableListRequest("BTC", "9bdf8e7ae03c237e115f09543fbdb40f8efa600106e78b67ce4d5adfadda2dbb", int32(0)) // LockSpendableListRequest | The request body of the Lock/Unlock UTXOs operation. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.UnlockSpendableList(ctx, walletId).LockSpendableListRequest(lockSpendableListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.UnlockSpendableList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnlockSpendableList`: LockSpendableList200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.UnlockSpendableList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockSpendableListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockSpendableListRequest** | [**LockSpendableListRequest**](LockSpendableListRequest.md) | The request body of the Lock/Unlock UTXOs operation. | 

### Return type

[**LockSpendableList200Response**](LockSpendableList200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWalletById

> WalletInfo UpdateWalletById(ctx, walletId).UpdateWalletByIdRequest(updateWalletByIdRequest).Execute()

Update wallet



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
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The wallet ID.
	updateWalletByIdRequest := *CoboWaas2.NewUpdateWalletByIdRequest() // UpdateWalletByIdRequest | The request body. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.UpdateWalletById(ctx, walletId).UpdateWalletByIdRequest(updateWalletByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.UpdateWalletById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWalletById`: WalletInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.UpdateWalletById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWalletByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWalletByIdRequest** | [**UpdateWalletByIdRequest**](UpdateWalletByIdRequest.md) | The request body. | 

### Return type

[**WalletInfo**](WalletInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

