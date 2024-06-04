# \WalletsAPI

All URIs are relative to *https://api.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWalletAddress**](WalletsAPI.md#AddWalletAddress) | **Post** /wallets/{wallet_id}/addresses | Add address to wallet
[**CreateWallet**](WalletsAPI.md#CreateWallet) | **Post** /wallets | Create wallet
[**DeleteWalletById**](WalletsAPI.md#DeleteWalletById) | **Delete** /wallets/{wallet_id} | Delete wallet by ID
[**GetAddressValidity**](WalletsAPI.md#GetAddressValidity) | **Get** /wallets/address/validity | Check address validity
[**GetChains**](WalletsAPI.md#GetChains) | **Get** /wallets/chains | List chain metadata
[**GetEnabledChains**](WalletsAPI.md#GetEnabledChains) | **Get** /wallets/enabled_chains | List enabled chains
[**GetEnabledTokens**](WalletsAPI.md#GetEnabledTokens) | **Get** /wallets/enabled_tokens | List enabled tokens
[**GetMaxTransferableValue**](WalletsAPI.md#GetMaxTransferableValue) | **Get** /wallets/{wallet_id}/max_transferable_value | Get max transferable value
[**GetSpendableList**](WalletsAPI.md#GetSpendableList) | **Get** /wallets/{wallet_id}/spendables | List spendable UTXOs
[**GetSupportedChains**](WalletsAPI.md#GetSupportedChains) | **Get** /wallets/supported_chains | List supported chains
[**GetSupportedTokens**](WalletsAPI.md#GetSupportedTokens) | **Get** /wallets/supported_tokens | List supported tokens
[**GetTokens**](WalletsAPI.md#GetTokens) | **Get** /wallets/tokens | List token metadata
[**GetWalletAddressTokenBalances**](WalletsAPI.md#GetWalletAddressTokenBalances) | **Get** /wallets/{wallet_id}/addresses/{address_id}/tokens | List token balances by address
[**GetWalletById**](WalletsAPI.md#GetWalletById) | **Get** /wallets/{wallet_id} | Retrieve wallet information by ID
[**GetWalletTokenBalances**](WalletsAPI.md#GetWalletTokenBalances) | **Get** /wallets/{wallet_id}/tokens | List token balances by wallet
[**ListAddresses**](WalletsAPI.md#ListAddresses) | **Get** /wallets/{wallet_id}/addresses | List wallet addresses by wallet ID
[**ListWallets**](WalletsAPI.md#ListWallets) | **Get** /wallets | List all wallets
[**LockSpendableList**](WalletsAPI.md#LockSpendableList) | **Post** /wallets/{wallet_id}/spendables/lock | Lock/Unlock the UTXOs in tx hash list
[**UpdateWalletById**](WalletsAPI.md#UpdateWalletById) | **Put** /wallets/{wallet_id} | Update wallet by ID



## AddWalletAddress

> []AddressInfo AddWalletAddress(ctx, walletId).AddWalletAddressRequest(addWalletAddressRequest).Execute()

Add address to wallet



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
	addWalletAddressRequest := *CoboWaas2.NewAddWalletAddressRequest("ETH_USDT", int32(1)) // AddWalletAddressRequest | The request body to add address to a wallet. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.AddWalletAddress(ctx, walletId).AddWalletAddressRequest(addWalletAddressRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.AddWalletAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddWalletAddress`: []AddressInfo
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.AddWalletAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**walletId** | **string** | The wallet ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWalletAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addWalletAddressRequest** | [**AddWalletAddressRequest**](AddWalletAddressRequest.md) | The request body to add address to a wallet. | 

### Return type

[**[]AddressInfo**](AddressInfo.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
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

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWalletById

> DeleteWalletById(ctx, walletId).Execute()

Delete wallet by ID



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
	r, err := apiClient.WalletsAPI.DeleteWalletById(ctx, walletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.DeleteWalletById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

 (empty response body)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressValidity

> GetAddressValidity200Response GetAddressValidity(ctx).TokenId(tokenId).AddressStr(addressStr).Execute()

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
	tokenId := "ETH_USDT" // string | The token ID.
	addressStr := "0x0000000000000000000000000000000000000000" // string | The wallet address.

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetAddressValidity(ctx).TokenId(tokenId).AddressStr(addressStr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetAddressValidity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressValidity`: GetAddressValidity200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetAddressValidity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressValidityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenId** | **string** | The token ID. | 
 **addressStr** | **string** | The wallet address. | 

### Return type

[**GetAddressValidity200Response**](GetAddressValidity200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChains

> GetChains200Response GetChains(ctx).ChainId(chainId).Limit(limit).Before(before).After(after).Execute()

List chain metadata



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
	chainId := "ETH" // string | The chain ID. (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The value range is [1, 50]. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `foo`, the request will retrieve a list of data objects that end before the object with the object ID `foo`. You can set this parameter to the value of `pagination.after` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `bar`, the request will retrieve a list of data objects that start after the object with the object ID `bar`. You can set this parameter to the value of `pagination.before` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetChains(ctx).ChainId(chainId).Limit(limit).Before(before).After(after).Execute()
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
 **chainId** | **string** | The chain ID. | 
 **limit** | **int32** | The maximum number of objects to return. The value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;foo&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;foo&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;bar&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;bar&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 

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

List enabled chains



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
	walletType := CoboWaas2.WalletType("Custodial") // WalletType | The wallet type.  - `Custodial`: Custodial Wallets  - `MPC`: MPC Wallets  - `SmartContract`: Smart Contract Wallets  - `Exchange`: Exchange Wallets  (optional)
	walletSubtype := CoboWaas2.WalletSubtype("Asset") // WalletSubtype | The wallet subtype.  - `Asset`: Custodial Wallets (Asset Wallets)  - `Web3`: Custodial Wallets (Web3 Wallets)  - `Main`: Exchange Wallets (Main Account)  - `Sub`: Exchange Wallets (Sub Account)  - `Org-Controlled`: MPC Wallets (Organization-Controlled Wallets)  - `User-Controlled`: MPC Wallets (User-Controlled Wallets)  - `Safe{Wallet}`: Smart Contract Wallets (Safe{Wallet})  (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The value range is [1, 50]. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `foo`, the request will retrieve a list of data objects that end before the object with the object ID `foo`. You can set this parameter to the value of `pagination.after` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `bar`, the request will retrieve a list of data objects that start after the object with the object ID `bar`. You can set this parameter to the value of `pagination.before` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)

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
 **walletType** | [**WalletType**](WalletType.md) | The wallet type.  - &#x60;Custodial&#x60;: Custodial Wallets  - &#x60;MPC&#x60;: MPC Wallets  - &#x60;SmartContract&#x60;: Smart Contract Wallets  - &#x60;Exchange&#x60;: Exchange Wallets  | 
 **walletSubtype** | [**WalletSubtype**](WalletSubtype.md) | The wallet subtype.  - &#x60;Asset&#x60;: Custodial Wallets (Asset Wallets)  - &#x60;Web3&#x60;: Custodial Wallets (Web3 Wallets)  - &#x60;Main&#x60;: Exchange Wallets (Main Account)  - &#x60;Sub&#x60;: Exchange Wallets (Sub Account)  - &#x60;Org-Controlled&#x60;: MPC Wallets (Organization-Controlled Wallets)  - &#x60;User-Controlled&#x60;: MPC Wallets (User-Controlled Wallets)  - &#x60;Safe{Wallet}&#x60;: Smart Contract Wallets (Safe{Wallet})  | 
 **limit** | **int32** | The maximum number of objects to return. The value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;foo&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;foo&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;bar&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;bar&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 

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

List enabled tokens



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
	walletType := CoboWaas2.WalletType("Custodial") // WalletType | The wallet type.  - `Custodial`: Custodial Wallets  - `MPC`: MPC Wallets  - `SmartContract`: Smart Contract Wallets  - `Exchange`: Exchange Wallets  (optional)
	walletSubtype := CoboWaas2.WalletSubtype("Asset") // WalletSubtype | The wallet subtype.  - `Asset`: Custodial Wallets (Asset Wallets)  - `Web3`: Custodial Wallets (Web3 Wallets)  - `Main`: Exchange Wallets (Main Account)  - `Sub`: Exchange Wallets (Sub Account)  - `Org-Controlled`: MPC Wallets (Organization-Controlled Wallets)  - `User-Controlled`: MPC Wallets (User-Controlled Wallets)  - `Safe{Wallet}`: Smart Contract Wallets (Safe{Wallet})  (optional)
	chainId := "ETH" // string | The chain ID. (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The value range is [1, 50]. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `foo`, the request will retrieve a list of data objects that end before the object with the object ID `foo`. You can set this parameter to the value of `pagination.after` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `bar`, the request will retrieve a list of data objects that start after the object with the object ID `bar`. You can set this parameter to the value of `pagination.before` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)

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
 **walletType** | [**WalletType**](WalletType.md) | The wallet type.  - &#x60;Custodial&#x60;: Custodial Wallets  - &#x60;MPC&#x60;: MPC Wallets  - &#x60;SmartContract&#x60;: Smart Contract Wallets  - &#x60;Exchange&#x60;: Exchange Wallets  | 
 **walletSubtype** | [**WalletSubtype**](WalletSubtype.md) | The wallet subtype.  - &#x60;Asset&#x60;: Custodial Wallets (Asset Wallets)  - &#x60;Web3&#x60;: Custodial Wallets (Web3 Wallets)  - &#x60;Main&#x60;: Exchange Wallets (Main Account)  - &#x60;Sub&#x60;: Exchange Wallets (Sub Account)  - &#x60;Org-Controlled&#x60;: MPC Wallets (Organization-Controlled Wallets)  - &#x60;User-Controlled&#x60;: MPC Wallets (User-Controlled Wallets)  - &#x60;Safe{Wallet}&#x60;: Smart Contract Wallets (Safe{Wallet})  | 
 **chainId** | **string** | The chain ID. | 
 **limit** | **int32** | The maximum number of objects to return. The value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;foo&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;foo&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;bar&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;bar&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 

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

> MaxTransferableValue GetMaxTransferableValue(ctx, walletId).ToAddress(toAddress).FromAddress(fromAddress).Execute()

Get max transferable value



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
	resp, r, err := apiClient.WalletsAPI.GetMaxTransferableValue(ctx, walletId).ToAddress(toAddress).FromAddress(fromAddress).Execute()
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

> []UTXO GetSpendableList(ctx, walletId).TxHash(txHash).TokenId(tokenId).AddressStr(addressStr).Execute()

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
	txHash := "dd7e1cecf6bbde1844ee1815b780711a1e306a718bcd23cd64401b48ef88eb83" // string | Transaction hash of the UTXO.
	tokenId := "ETH_USDT" // string | The token ID. (optional)
	addressStr := "2N2xFZtbCFB6Nb3Pj9Sxsx5mX2fxX3yEgkE" // string | The wallet address. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetSpendableList(ctx, walletId).TxHash(txHash).TokenId(tokenId).AddressStr(addressStr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetSpendableList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpendableList`: []UTXO
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

 **txHash** | **string** | Transaction hash of the UTXO. | 
 **tokenId** | **string** | The token ID. | 
 **addressStr** | **string** | The wallet address. | 

### Return type

[**[]UTXO**](UTXO.md)

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

List supported chains



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
	walletType := CoboWaas2.WalletType("Custodial") // WalletType | The wallet type.  - `Custodial`: Custodial Wallets  - `MPC`: MPC Wallets  - `SmartContract`: Smart Contract Wallets  - `Exchange`: Exchange Wallets  (optional)
	walletSubtype := CoboWaas2.WalletSubtype("Asset") // WalletSubtype | The wallet subtype.  - `Asset`: Custodial Wallets (Asset Wallets)  - `Web3`: Custodial Wallets (Web3 Wallets)  - `Main`: Exchange Wallets (Main Account)  - `Sub`: Exchange Wallets (Sub Account)  - `Org-Controlled`: MPC Wallets (Organization-Controlled Wallets)  - `User-Controlled`: MPC Wallets (User-Controlled Wallets)  - `Safe{Wallet}`: Smart Contract Wallets (Safe{Wallet})  (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The value range is [1, 50]. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `foo`, the request will retrieve a list of data objects that end before the object with the object ID `foo`. You can set this parameter to the value of `pagination.after` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `bar`, the request will retrieve a list of data objects that start after the object with the object ID `bar`. You can set this parameter to the value of `pagination.before` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)

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
 **walletType** | [**WalletType**](WalletType.md) | The wallet type.  - &#x60;Custodial&#x60;: Custodial Wallets  - &#x60;MPC&#x60;: MPC Wallets  - &#x60;SmartContract&#x60;: Smart Contract Wallets  - &#x60;Exchange&#x60;: Exchange Wallets  | 
 **walletSubtype** | [**WalletSubtype**](WalletSubtype.md) | The wallet subtype.  - &#x60;Asset&#x60;: Custodial Wallets (Asset Wallets)  - &#x60;Web3&#x60;: Custodial Wallets (Web3 Wallets)  - &#x60;Main&#x60;: Exchange Wallets (Main Account)  - &#x60;Sub&#x60;: Exchange Wallets (Sub Account)  - &#x60;Org-Controlled&#x60;: MPC Wallets (Organization-Controlled Wallets)  - &#x60;User-Controlled&#x60;: MPC Wallets (User-Controlled Wallets)  - &#x60;Safe{Wallet}&#x60;: Smart Contract Wallets (Safe{Wallet})  | 
 **limit** | **int32** | The maximum number of objects to return. The value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;foo&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;foo&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;bar&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;bar&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 

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

List supported tokens



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
	walletType := CoboWaas2.WalletType("Custodial") // WalletType | The wallet type.  - `Custodial`: Custodial Wallets  - `MPC`: MPC Wallets  - `SmartContract`: Smart Contract Wallets  - `Exchange`: Exchange Wallets  (optional)
	walletSubtype := CoboWaas2.WalletSubtype("Asset") // WalletSubtype | The wallet subtype.  - `Asset`: Custodial Wallets (Asset Wallets)  - `Web3`: Custodial Wallets (Web3 Wallets)  - `Main`: Exchange Wallets (Main Account)  - `Sub`: Exchange Wallets (Sub Account)  - `Org-Controlled`: MPC Wallets (Organization-Controlled Wallets)  - `User-Controlled`: MPC Wallets (User-Controlled Wallets)  - `Safe{Wallet}`: Smart Contract Wallets (Safe{Wallet})  (optional)
	chainId := "ETH" // string | The chain ID. (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The value range is [1, 50]. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `foo`, the request will retrieve a list of data objects that end before the object with the object ID `foo`. You can set this parameter to the value of `pagination.after` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `bar`, the request will retrieve a list of data objects that start after the object with the object ID `bar`. You can set this parameter to the value of `pagination.before` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)

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
 **walletType** | [**WalletType**](WalletType.md) | The wallet type.  - &#x60;Custodial&#x60;: Custodial Wallets  - &#x60;MPC&#x60;: MPC Wallets  - &#x60;SmartContract&#x60;: Smart Contract Wallets  - &#x60;Exchange&#x60;: Exchange Wallets  | 
 **walletSubtype** | [**WalletSubtype**](WalletSubtype.md) | The wallet subtype.  - &#x60;Asset&#x60;: Custodial Wallets (Asset Wallets)  - &#x60;Web3&#x60;: Custodial Wallets (Web3 Wallets)  - &#x60;Main&#x60;: Exchange Wallets (Main Account)  - &#x60;Sub&#x60;: Exchange Wallets (Sub Account)  - &#x60;Org-Controlled&#x60;: MPC Wallets (Organization-Controlled Wallets)  - &#x60;User-Controlled&#x60;: MPC Wallets (User-Controlled Wallets)  - &#x60;Safe{Wallet}&#x60;: Smart Contract Wallets (Safe{Wallet})  | 
 **chainId** | **string** | The chain ID. | 
 **limit** | **int32** | The maximum number of objects to return. The value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;foo&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;foo&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;bar&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;bar&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 

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

> GetTokens200Response GetTokens(ctx).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()

List token metadata



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
	tokenId := "ETH_USDT" // string | The token ID. (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The value range is [1, 50]. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `foo`, the request will retrieve a list of data objects that end before the object with the object ID `foo`. You can set this parameter to the value of `pagination.after` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `bar`, the request will retrieve a list of data objects that start after the object with the object ID `bar`. You can set this parameter to the value of `pagination.before` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetTokens(ctx).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()
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
 **tokenId** | **string** | The token ID. | 
 **limit** | **int32** | The maximum number of objects to return. The value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;foo&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;foo&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;bar&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;bar&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 

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

> GetWalletTokenBalances200Response GetWalletAddressTokenBalances(ctx, walletId, addressId).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()

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
	addressId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The address ID.
	tokenId := "ETH_USDT" // string | The token ID. (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The value range is [1, 50]. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `foo`, the request will retrieve a list of data objects that end before the object with the object ID `foo`. You can set this parameter to the value of `pagination.after` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `bar`, the request will retrieve a list of data objects that start after the object with the object ID `bar`. You can set this parameter to the value of `pagination.before` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.GetWalletAddressTokenBalances(ctx, walletId, addressId).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()
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
**addressId** | **string** | The address ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletAddressTokenBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenId** | **string** | The token ID. | 
 **limit** | **int32** | The maximum number of objects to return. The value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;foo&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;foo&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;bar&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;bar&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 

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

Retrieve wallet information by ID



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
	tokenId := "ETH_USDT" // string | The token ID. (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The value range is [1, 50]. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `foo`, the request will retrieve a list of data objects that end before the object with the object ID `foo`. You can set this parameter to the value of `pagination.after` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `bar`, the request will retrieve a list of data objects that start after the object with the object ID `bar`. You can set this parameter to the value of `pagination.before` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)

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

 **tokenId** | **string** | The token ID. | 
 **limit** | **int32** | The maximum number of objects to return. The value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;foo&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;foo&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;bar&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;bar&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 

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

> ListAddresses200Response ListAddresses(ctx, walletId).TokenId(tokenId).AddressStr(addressStr).Limit(limit).Before(before).After(after).Execute()

List wallet addresses by wallet ID



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
	tokenId := "ETH_USDT" // string | The token ID. (optional)
	addressStr := "2N2xFZtbCFB6Nb3Pj9Sxsx5mX2fxX3yEgkE" // string | The wallet address. (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The value range is [1, 50]. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `foo`, the request will retrieve a list of data objects that end before the object with the object ID `foo`. You can set this parameter to the value of `pagination.after` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `bar`, the request will retrieve a list of data objects that start after the object with the object ID `bar`. You can set this parameter to the value of `pagination.before` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.ListAddresses(ctx, walletId).TokenId(tokenId).AddressStr(addressStr).Limit(limit).Before(before).After(after).Execute()
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

 **tokenId** | **string** | The token ID. | 
 **addressStr** | **string** | The wallet address. | 
 **limit** | **int32** | The maximum number of objects to return. The value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;foo&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;foo&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;bar&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;bar&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 

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

> ListWallets200Response ListWallets(ctx).WalletType(walletType).WalletSubtype(walletSubtype).VaultId(vaultId).Limit(limit).Before(before).After(after).Execute()

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
	walletType := CoboWaas2.WalletType("Custodial") // WalletType | The wallet type.  - `Custodial`: Custodial Wallets  - `MPC`: MPC Wallets  - `SmartContract`: Smart Contract Wallets  - `Exchange`: Exchange Wallets  (optional)
	walletSubtype := CoboWaas2.WalletSubtype("Asset") // WalletSubtype | The wallet subtype.  - `Asset`: Custodial Wallets (Asset Wallets)  - `Web3`: Custodial Wallets (Web3 Wallets)  - `Main`: Exchange Wallets (Main Account)  - `Sub`: Exchange Wallets (Sub Account)  - `Org-Controlled`: MPC Wallets (Organization-Controlled Wallets)  - `User-Controlled`: MPC Wallets (User-Controlled Wallets)  - `Safe{Wallet}`: Smart Contract Wallets (Safe{Wallet})  (optional)
	vaultId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | The MPC Vault ID. (optional)
	limit := int32(10) // int32 | The maximum number of objects to return. The value range is [1, 50]. (optional) (default to 10)
	before := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `before` as `foo`, the request will retrieve a list of data objects that end before the object with the object ID `foo`. You can set this parameter to the value of `pagination.after` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)
	after := "8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f" // string | An object ID which serves as a cursor for pagination. For example, if you specify `after` as `bar`, the request will retrieve a list of data objects that start after the object with the object ID `bar`. You can set this parameter to the value of `pagination.before` in the response of the previous request. If you set both `after` or `before`, only the setting of `before` will take effect. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.WalletsAPI.ListWallets(ctx).WalletType(walletType).WalletSubtype(walletSubtype).VaultId(vaultId).Limit(limit).Before(before).After(after).Execute()
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
 **walletType** | [**WalletType**](WalletType.md) | The wallet type.  - &#x60;Custodial&#x60;: Custodial Wallets  - &#x60;MPC&#x60;: MPC Wallets  - &#x60;SmartContract&#x60;: Smart Contract Wallets  - &#x60;Exchange&#x60;: Exchange Wallets  | 
 **walletSubtype** | [**WalletSubtype**](WalletSubtype.md) | The wallet subtype.  - &#x60;Asset&#x60;: Custodial Wallets (Asset Wallets)  - &#x60;Web3&#x60;: Custodial Wallets (Web3 Wallets)  - &#x60;Main&#x60;: Exchange Wallets (Main Account)  - &#x60;Sub&#x60;: Exchange Wallets (Sub Account)  - &#x60;Org-Controlled&#x60;: MPC Wallets (Organization-Controlled Wallets)  - &#x60;User-Controlled&#x60;: MPC Wallets (User-Controlled Wallets)  - &#x60;Safe{Wallet}&#x60;: Smart Contract Wallets (Safe{Wallet})  | 
 **vaultId** | **string** | The MPC Vault ID. | 
 **limit** | **int32** | The maximum number of objects to return. The value range is [1, 50]. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;foo&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;foo&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;bar&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;bar&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request. If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect. | 

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

Lock/Unlock the UTXOs in tx hash list



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
	lockSpendableListRequest := *CoboWaas2.NewLockSpendableListRequest() // LockSpendableListRequest | The request body to lock/unlock spendable (optional)

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

 **lockSpendableListRequest** | [**LockSpendableListRequest**](LockSpendableListRequest.md) | The request body to lock/unlock spendable | 

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

Update wallet by ID



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

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

