# \TokenizationAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BurnTokenization**](TokenizationAPI.md#BurnTokenization) | **Post** /tokenization/tokens/{token_id}/burn | Burn tokens
[**EstimateTokenizationFee**](TokenizationAPI.md#EstimateTokenizationFee) | **Post** /tokenization/estimate_fee | Estimate tokenization operation fee
[**GetTokenizationActivity**](TokenizationAPI.md#GetTokenizationActivity) | **Get** /tokenization/activities/{activity_id} | Get tokenization activity details
[**GetTokenizationAllowlistActivation**](TokenizationAPI.md#GetTokenizationAllowlistActivation) | **Get** /tokenization/tokens/{token_id}/allowlist/activation | Get allowlist activation status
[**GetTokenizationInfo**](TokenizationAPI.md#GetTokenizationInfo) | **Get** /tokenization/tokens/{token_id} | Get token details
[**IssueToken**](TokenizationAPI.md#IssueToken) | **Post** /tokenization/tokens | Issue token
[**ListIssuedTokens**](TokenizationAPI.md#ListIssuedTokens) | **Get** /tokenization/tokens | List issued tokens
[**ListTokenizationActivities**](TokenizationAPI.md#ListTokenizationActivities) | **Get** /tokenization/activities | List tokenization activities
[**ListTokenizationAllowlistAddresses**](TokenizationAPI.md#ListTokenizationAllowlistAddresses) | **Get** /tokenization/tokens/{token_id}/allowlist/addresses | List addresses on allowlist
[**ListTokenizationBlocklistAddresses**](TokenizationAPI.md#ListTokenizationBlocklistAddresses) | **Get** /tokenization/tokens/{token_id}/blocklist/addresses | Lists addresses on blocklist
[**ListTokenizationHoldings**](TokenizationAPI.md#ListTokenizationHoldings) | **Get** /tokenization/tokens/{token_id}/holdings | Get token holdings information
[**ListTokenizationSupportedChains**](TokenizationAPI.md#ListTokenizationSupportedChains) | **Get** /tokenization/enabled_chains | List supported chains for tokenization
[**MintTokenization**](TokenizationAPI.md#MintTokenization) | **Post** /tokenization/tokens/{token_id}/mint | Mint tokens
[**PauseTokenization**](TokenizationAPI.md#PauseTokenization) | **Post** /tokenization/tokens/{token_id}/pause | Pause token contract
[**TokenizationContractCall**](TokenizationAPI.md#TokenizationContractCall) | **Post** /tokenization/tokens/{token_id}/contract_call | Call token contract
[**UnpauseTokenization**](TokenizationAPI.md#UnpauseTokenization) | **Post** /tokenization/tokens/{token_id}/unpause | Unpause token contract
[**UpdateTokenizationAllowlistActivation**](TokenizationAPI.md#UpdateTokenizationAllowlistActivation) | **Post** /tokenization/tokens/{token_id}/allowlist/activation | Activate or deactivate allowlist
[**UpdateTokenizationAllowlistAddresses**](TokenizationAPI.md#UpdateTokenizationAllowlistAddresses) | **Post** /tokenization/tokens/{token_id}/allowlist/addresses | Update addresses on allowlist
[**UpdateTokenizationBlocklistAddresses**](TokenizationAPI.md#UpdateTokenizationBlocklistAddresses) | **Post** /tokenization/tokens/{token_id}/blocklist/addresses | Update addresses on blocklist



## BurnTokenization

> TokenizationOperationResponse BurnTokenization(ctx, tokenId).TokenizationBurnTokenRequest(tokenizationBurnTokenRequest).Execute()

Burn tokens



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
	tokenId := "ETH_USDT"
	tokenizationBurnTokenRequest := *coboWaas2.NewTokenizationBurnTokenRequest(coboWaas2.TokenizationTokenOperationSource{TokenizationMpcOperationSource: coboWaas2.NewTokenizationMpcOperationSource(coboWaas2.TokenizationOperationSourceType("Org-Controlled"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")}, []coboWaas2.TokenizationBurnTokenParamsBurnsInner{*coboWaas2.NewTokenizationBurnTokenParamsBurnsInner("0.99", "0x051A924H4dCb264226d7B036C2893a0D344")}, coboWaas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: coboWaas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", coboWaas2.FeeType("Fixed"), "ETH")})

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
	resp, r, err := apiClient.TokenizationAPI.BurnTokenization(ctx, tokenId).TokenizationBurnTokenRequest(tokenizationBurnTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.BurnTokenization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BurnTokenization`: TokenizationOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.BurnTokenization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**tokenId** | **string** | The token ID, which is the unique identifier of a token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBurnTokenizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenizationBurnTokenRequest** | [**TokenizationBurnTokenRequest**](TokenizationBurnTokenRequest.md) | The request body for burning tokens. | 

### Return type

[**TokenizationOperationResponse**](TokenizationOperationResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EstimateTokenizationFee

> EstimatedFee EstimateTokenizationFee(ctx).TokenizationEstimateFeeRequest(tokenizationEstimateFeeRequest).Execute()

Estimate tokenization operation fee



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
	tokenizationEstimateFeeRequest := *coboWaas2.NewTokenizationEstimateFeeRequest(coboWaas2.TokenizationEstimateFeeRequest_operation_params{TokenizationBurnEstimateFeeParams: coboWaas2.NewTokenizationBurnEstimateFeeParams(coboWaas2.TokenizationTokenOperationSource{TokenizationMpcOperationSource: coboWaas2.NewTokenizationMpcOperationSource(coboWaas2.TokenizationOperationSourceType("Org-Controlled"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")}, []coboWaas2.TokenizationBurnTokenParamsBurnsInner{*coboWaas2.NewTokenizationBurnTokenParamsBurnsInner("0.99", "0x051A924H4dCb264226d7B036C2893a0D344")}, coboWaas2.TokenizationOperationType("Issue"), "8a4f9324-ef2a-43cf-9f0e-d7f99999d3e8")})

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
	resp, r, err := apiClient.TokenizationAPI.EstimateTokenizationFee(ctx).TokenizationEstimateFeeRequest(tokenizationEstimateFeeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.EstimateTokenizationFee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EstimateTokenizationFee`: EstimatedFee
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.EstimateTokenizationFee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEstimateTokenizationFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenizationEstimateFeeRequest** | [**TokenizationEstimateFeeRequest**](TokenizationEstimateFeeRequest.md) | The request body to estimate tokenization operation fee. | 

### Return type

[**EstimatedFee**](EstimatedFee.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenizationActivity

> TokenizationActivityInfo GetTokenizationActivity(ctx, activityId).Execute()

Get tokenization activity details



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
	activityId := "b7c8e9d0-f1a2-3b4c-5d6e-7f8a9b0c1d2e"

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
	resp, r, err := apiClient.TokenizationAPI.GetTokenizationActivity(ctx, activityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.GetTokenizationActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenizationActivity`: TokenizationActivityInfo
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.GetTokenizationActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**activityId** | **string** | The ID of the activity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenizationActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenizationActivityInfo**](TokenizationActivityInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenizationAllowlistActivation

> GetTokenizationAllowlistActivation200Response GetTokenizationAllowlistActivation(ctx, tokenId).Execute()

Get allowlist activation status



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
	tokenId := "ETH_USDT"

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
	resp, r, err := apiClient.TokenizationAPI.GetTokenizationAllowlistActivation(ctx, tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.GetTokenizationAllowlistActivation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenizationAllowlistActivation`: GetTokenizationAllowlistActivation200Response
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.GetTokenizationAllowlistActivation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**tokenId** | **string** | The token ID, which is the unique identifier of a token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenizationAllowlistActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTokenizationAllowlistActivation200Response**](GetTokenizationAllowlistActivation200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenizationInfo

> TokenizationTokenDetailInfo GetTokenizationInfo(ctx, tokenId).Execute()

Get token details



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
	tokenId := "ETH_USDT"

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
	resp, r, err := apiClient.TokenizationAPI.GetTokenizationInfo(ctx, tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.GetTokenizationInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenizationInfo`: TokenizationTokenDetailInfo
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.GetTokenizationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**tokenId** | **string** | The token ID, which is the unique identifier of a token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenizationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenizationTokenDetailInfo**](TokenizationTokenDetailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueToken

> TokenizationOperationResponse IssueToken(ctx).TokenizationIssuedTokenRequest(tokenizationIssuedTokenRequest).Execute()

Issue token



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
	tokenizationIssuedTokenRequest := *coboWaas2.NewTokenizationIssuedTokenRequest("ETH", coboWaas2.TokenizationTokenOperationSource{TokenizationMpcOperationSource: coboWaas2.NewTokenizationMpcOperationSource(coboWaas2.TokenizationOperationSourceType("Org-Controlled"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")}, coboWaas2.TokenizationIssueTokenParams_token_params{TokenizationERC20TokenParams: coboWaas2.NewTokenizationERC20TokenParams(coboWaas2.TokenizationTokenStandard("ERC20"), "My Awesome Token", "MAT", int32(18))}, coboWaas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: coboWaas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", coboWaas2.FeeType("Fixed"), "ETH")})

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
	resp, r, err := apiClient.TokenizationAPI.IssueToken(ctx).TokenizationIssuedTokenRequest(tokenizationIssuedTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.IssueToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueToken`: TokenizationOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.IssueToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenizationIssuedTokenRequest** | [**TokenizationIssuedTokenRequest**](TokenizationIssuedTokenRequest.md) | The request body to issue a new token. | 

### Return type

[**TokenizationOperationResponse**](TokenizationOperationResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIssuedTokens

> TokenizationListTokenInfoResponse ListIssuedTokens(ctx).ChainId(chainId).TokenId(tokenId).TokenStandard(tokenStandard).Status(status).Limit(limit).Before(before).After(after).Execute()

List issued tokens



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
	chainId := "ETH"
	tokenId := "ETH_USDT"
	tokenStandard := coboWaas2.TokenizationTokenStandard("ERC20")
	status := coboWaas2.TokenizationStatus("Processing")
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
	resp, r, err := apiClient.TokenizationAPI.ListIssuedTokens(ctx).ChainId(chainId).TokenId(tokenId).TokenStandard(tokenStandard).Status(status).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.ListIssuedTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIssuedTokens`: TokenizationListTokenInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.ListIssuedTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIssuedTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | 
 **tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
 **tokenStandard** | [**TokenizationTokenStandard**](TokenizationTokenStandard.md) | Filter by token standard. | 
 **status** | [**TokenizationStatus**](TokenizationStatus.md) | Filter by token status. | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 

### Return type

[**TokenizationListTokenInfoResponse**](TokenizationListTokenInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokenizationActivities

> TokenizationListActivitiesResponse ListTokenizationActivities(ctx).TokenId(tokenId).ActivityType(activityType).ActivityStatus(activityStatus).Limit(limit).After(after).Before(before).Direction(direction).Execute()

List tokenization activities



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
	tokenId := "ETH_USDT"
	activityType := coboWaas2.TokenizationOperationType("Issue")
	activityStatus := coboWaas2.TokenizationActivityStatus("Processing")
	limit := int32(10)
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk"
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1"
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
	resp, r, err := apiClient.TokenizationAPI.ListTokenizationActivities(ctx).TokenId(tokenId).ActivityType(activityType).ActivityStatus(activityStatus).Limit(limit).After(after).Before(before).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.ListTokenizationActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTokenizationActivities`: TokenizationListActivitiesResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.ListTokenizationActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTokenizationActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
 **activityType** | [**TokenizationOperationType**](TokenizationOperationType.md) | Filter by tokenization activity type. | 
 **activityStatus** | [**TokenizationActivityStatus**](TokenizationActivityStatus.md) | Filter by tokenization activity status. | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **direction** | **string** | The sort direction. Possible values include:   - &#x60;ASC&#x60;: Sort the results in ascending order.   - &#x60;DESC&#x60;: Sort the results in descending order.  | [default to &quot;ASC&quot;]

### Return type

[**TokenizationListActivitiesResponse**](TokenizationListActivitiesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokenizationAllowlistAddresses

> TokenizationAllowlistAddressesResponse ListTokenizationAllowlistAddresses(ctx, tokenId).Limit(limit).After(after).Before(before).Direction(direction).Execute()

List addresses on allowlist



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
	tokenId := "ETH_USDT"
	limit := int32(10)
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk"
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1"
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
	resp, r, err := apiClient.TokenizationAPI.ListTokenizationAllowlistAddresses(ctx, tokenId).Limit(limit).After(after).Before(before).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.ListTokenizationAllowlistAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTokenizationAllowlistAddresses`: TokenizationAllowlistAddressesResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.ListTokenizationAllowlistAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**tokenId** | **string** | The token ID, which is the unique identifier of a token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTokenizationAllowlistAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **direction** | **string** | The sort direction. Possible values include:   - &#x60;ASC&#x60;: Sort the results in ascending order.   - &#x60;DESC&#x60;: Sort the results in descending order.  | [default to &quot;ASC&quot;]

### Return type

[**TokenizationAllowlistAddressesResponse**](TokenizationAllowlistAddressesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokenizationBlocklistAddresses

> ListTokenizationBlocklistAddresses200Response ListTokenizationBlocklistAddresses(ctx, tokenId).Limit(limit).After(after).Before(before).Direction(direction).Execute()

Lists addresses on blocklist



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
	tokenId := "ETH_USDT"
	limit := int32(10)
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk"
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1"
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
	resp, r, err := apiClient.TokenizationAPI.ListTokenizationBlocklistAddresses(ctx, tokenId).Limit(limit).After(after).Before(before).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.ListTokenizationBlocklistAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTokenizationBlocklistAddresses`: ListTokenizationBlocklistAddresses200Response
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.ListTokenizationBlocklistAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**tokenId** | **string** | The token ID, which is the unique identifier of a token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTokenizationBlocklistAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **direction** | **string** | The sort direction. Possible values include:   - &#x60;ASC&#x60;: Sort the results in ascending order.   - &#x60;DESC&#x60;: Sort the results in descending order.  | [default to &quot;ASC&quot;]

### Return type

[**ListTokenizationBlocklistAddresses200Response**](ListTokenizationBlocklistAddresses200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokenizationHoldings

> TokenizationListHoldingsResponse ListTokenizationHoldings(ctx, tokenId).Limit(limit).Before(before).After(after).Execute()

Get token holdings information



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
	tokenId := "ETH_USDT"
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
	resp, r, err := apiClient.TokenizationAPI.ListTokenizationHoldings(ctx, tokenId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.ListTokenizationHoldings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTokenizationHoldings`: TokenizationListHoldingsResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.ListTokenizationHoldings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**tokenId** | **string** | The token ID, which is the unique identifier of a token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTokenizationHoldingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 

### Return type

[**TokenizationListHoldingsResponse**](TokenizationListHoldingsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokenizationSupportedChains

> TokenizationListEnabledChainsResponse ListTokenizationSupportedChains(ctx).Limit(limit).After(after).Before(before).Execute()

List supported chains for tokenization



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
	after := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk"
	before := "RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1"

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
	resp, r, err := apiClient.TokenizationAPI.ListTokenizationSupportedChains(ctx).Limit(limit).After(after).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.ListTokenizationSupportedChains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTokenizationSupportedChains`: TokenizationListEnabledChainsResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.ListTokenizationSupportedChains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTokenizationSupportedChainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 

### Return type

[**TokenizationListEnabledChainsResponse**](TokenizationListEnabledChainsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MintTokenization

> TokenizationOperationResponse MintTokenization(ctx, tokenId).TokenizationMintTokenRequest(tokenizationMintTokenRequest).Execute()

Mint tokens



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
	tokenId := "ETH_USDT"
	tokenizationMintTokenRequest := *coboWaas2.NewTokenizationMintTokenRequest(coboWaas2.TokenizationTokenOperationSource{TokenizationMpcOperationSource: coboWaas2.NewTokenizationMpcOperationSource(coboWaas2.TokenizationOperationSourceType("Org-Controlled"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")}, []coboWaas2.TokenizationMintTokenParamsMintsInner{*coboWaas2.NewTokenizationMintTokenParamsMintsInner("0.99", "0x051A924H4dCb264226d7B036C2893a0D344")}, coboWaas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: coboWaas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", coboWaas2.FeeType("Fixed"), "ETH")})

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
	resp, r, err := apiClient.TokenizationAPI.MintTokenization(ctx, tokenId).TokenizationMintTokenRequest(tokenizationMintTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.MintTokenization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MintTokenization`: TokenizationOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.MintTokenization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**tokenId** | **string** | The token ID, which is the unique identifier of a token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMintTokenizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenizationMintTokenRequest** | [**TokenizationMintTokenRequest**](TokenizationMintTokenRequest.md) | The request body for minting tokens. | 

### Return type

[**TokenizationOperationResponse**](TokenizationOperationResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseTokenization

> TokenizationOperationResponse PauseTokenization(ctx, tokenId).TokenizationPauseTokenRequest(tokenizationPauseTokenRequest).Execute()

Pause token contract



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
	tokenId := "ETH_USDT"
	tokenizationPauseTokenRequest := *coboWaas2.NewTokenizationPauseTokenRequest(coboWaas2.TokenizationTokenOperationSource{TokenizationMpcOperationSource: coboWaas2.NewTokenizationMpcOperationSource(coboWaas2.TokenizationOperationSourceType("Org-Controlled"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")}, coboWaas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: coboWaas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", coboWaas2.FeeType("Fixed"), "ETH")})

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
	resp, r, err := apiClient.TokenizationAPI.PauseTokenization(ctx, tokenId).TokenizationPauseTokenRequest(tokenizationPauseTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.PauseTokenization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseTokenization`: TokenizationOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.PauseTokenization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**tokenId** | **string** | The token ID, which is the unique identifier of a token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseTokenizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenizationPauseTokenRequest** | [**TokenizationPauseTokenRequest**](TokenizationPauseTokenRequest.md) | The request body for pausing tokens. | 

### Return type

[**TokenizationOperationResponse**](TokenizationOperationResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenizationContractCall

> TokenizationOperationResponse TokenizationContractCall(ctx, tokenId).TokenizationContractCallRequest(tokenizationContractCallRequest).Execute()

Call token contract



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
	tokenId := "ETH_USDT"
	tokenizationContractCallRequest := *coboWaas2.NewTokenizationContractCallRequest(coboWaas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: coboWaas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", coboWaas2.FeeType("Fixed"), "ETH")})

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
	resp, r, err := apiClient.TokenizationAPI.TokenizationContractCall(ctx, tokenId).TokenizationContractCallRequest(tokenizationContractCallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.TokenizationContractCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenizationContractCall`: TokenizationOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.TokenizationContractCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**tokenId** | **string** | The token ID, which is the unique identifier of a token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenizationContractCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenizationContractCallRequest** | [**TokenizationContractCallRequest**](TokenizationContractCallRequest.md) | The request body for contract call. | 

### Return type

[**TokenizationOperationResponse**](TokenizationOperationResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpauseTokenization

> TokenizationOperationResponse UnpauseTokenization(ctx, tokenId).TokenizationUnpauseTokenRequest(tokenizationUnpauseTokenRequest).Execute()

Unpause token contract



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
	tokenId := "ETH_USDT"
	tokenizationUnpauseTokenRequest := *coboWaas2.NewTokenizationUnpauseTokenRequest(coboWaas2.TokenizationTokenOperationSource{TokenizationMpcOperationSource: coboWaas2.NewTokenizationMpcOperationSource(coboWaas2.TokenizationOperationSourceType("Org-Controlled"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")}, coboWaas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: coboWaas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", coboWaas2.FeeType("Fixed"), "ETH")})

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
	resp, r, err := apiClient.TokenizationAPI.UnpauseTokenization(ctx, tokenId).TokenizationUnpauseTokenRequest(tokenizationUnpauseTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.UnpauseTokenization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnpauseTokenization`: TokenizationOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.UnpauseTokenization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**tokenId** | **string** | The token ID, which is the unique identifier of a token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpauseTokenizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenizationUnpauseTokenRequest** | [**TokenizationUnpauseTokenRequest**](TokenizationUnpauseTokenRequest.md) | The request body for unpausing tokens. | 

### Return type

[**TokenizationOperationResponse**](TokenizationOperationResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTokenizationAllowlistActivation

> TokenizationOperationResponse UpdateTokenizationAllowlistActivation(ctx, tokenId).TokenizationAllowlistActivationRequest(tokenizationAllowlistActivationRequest).Execute()

Activate or deactivate allowlist



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
	tokenId := "ETH_USDT"
	tokenizationAllowlistActivationRequest := *coboWaas2.NewTokenizationAllowlistActivationRequest(coboWaas2.TokenizationTokenOperationSource{TokenizationMpcOperationSource: coboWaas2.NewTokenizationMpcOperationSource(coboWaas2.TokenizationOperationSourceType("Org-Controlled"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")}, true, coboWaas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: coboWaas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", coboWaas2.FeeType("Fixed"), "ETH")})

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
	resp, r, err := apiClient.TokenizationAPI.UpdateTokenizationAllowlistActivation(ctx, tokenId).TokenizationAllowlistActivationRequest(tokenizationAllowlistActivationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.UpdateTokenizationAllowlistActivation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTokenizationAllowlistActivation`: TokenizationOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.UpdateTokenizationAllowlistActivation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**tokenId** | **string** | The token ID, which is the unique identifier of a token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenizationAllowlistActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenizationAllowlistActivationRequest** | [**TokenizationAllowlistActivationRequest**](TokenizationAllowlistActivationRequest.md) | The request body for activating or deactivating the allowlist. | 

### Return type

[**TokenizationOperationResponse**](TokenizationOperationResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTokenizationAllowlistAddresses

> TokenizationOperationResponse UpdateTokenizationAllowlistAddresses(ctx, tokenId).TokenizationUpdateAllowlistAddressesRequest(tokenizationUpdateAllowlistAddressesRequest).Execute()

Update addresses on allowlist



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
	tokenId := "ETH_USDT"
	tokenizationUpdateAllowlistAddressesRequest := *coboWaas2.NewTokenizationUpdateAllowlistAddressesRequest(coboWaas2.TokenizationUpdateAddressAction("Grant"), coboWaas2.TokenizationTokenOperationSource{TokenizationMpcOperationSource: coboWaas2.NewTokenizationMpcOperationSource(coboWaas2.TokenizationOperationSourceType("Org-Controlled"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")}, []coboWaas2.TokenizationUpdateAllowlistAddressesParamsAddressesInner{*coboWaas2.NewTokenizationUpdateAllowlistAddressesParamsAddressesInner("0x789abc...")}, coboWaas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: coboWaas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", coboWaas2.FeeType("Fixed"), "ETH")})

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
	resp, r, err := apiClient.TokenizationAPI.UpdateTokenizationAllowlistAddresses(ctx, tokenId).TokenizationUpdateAllowlistAddressesRequest(tokenizationUpdateAllowlistAddressesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.UpdateTokenizationAllowlistAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTokenizationAllowlistAddresses`: TokenizationOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.UpdateTokenizationAllowlistAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**tokenId** | **string** | The token ID, which is the unique identifier of a token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenizationAllowlistAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenizationUpdateAllowlistAddressesRequest** | [**TokenizationUpdateAllowlistAddressesRequest**](TokenizationUpdateAllowlistAddressesRequest.md) | The request body for adding or removing multiple addresses on the allowlist. | 

### Return type

[**TokenizationOperationResponse**](TokenizationOperationResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTokenizationBlocklistAddresses

> TokenizationOperationResponse UpdateTokenizationBlocklistAddresses(ctx, tokenId).TokenizationUpdateBlocklistAddressesRequest(tokenizationUpdateBlocklistAddressesRequest).Execute()

Update addresses on blocklist



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
	tokenId := "ETH_USDT"
	tokenizationUpdateBlocklistAddressesRequest := *coboWaas2.NewTokenizationUpdateBlocklistAddressesRequest(coboWaas2.TokenizationUpdateAddressAction("Grant"), coboWaas2.TokenizationTokenOperationSource{TokenizationMpcOperationSource: coboWaas2.NewTokenizationMpcOperationSource(coboWaas2.TokenizationOperationSourceType("Org-Controlled"), "f47ac10b-58cc-4372-a567-0e02b2c3d479", "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")}, []coboWaas2.TokenizationUpdateBlocklistAddressesParamsAddressesInner{*coboWaas2.NewTokenizationUpdateBlocklistAddressesParamsAddressesInner("0x789abc...")}, coboWaas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: coboWaas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", coboWaas2.FeeType("Fixed"), "ETH")})

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
	resp, r, err := apiClient.TokenizationAPI.UpdateTokenizationBlocklistAddresses(ctx, tokenId).TokenizationUpdateBlocklistAddressesRequest(tokenizationUpdateBlocklistAddressesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenizationAPI.UpdateTokenizationBlocklistAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTokenizationBlocklistAddresses`: TokenizationOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenizationAPI.UpdateTokenizationBlocklistAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**tokenId** | **string** | The token ID, which is the unique identifier of a token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenizationBlocklistAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenizationUpdateBlocklistAddressesRequest** | [**TokenizationUpdateBlocklistAddressesRequest**](TokenizationUpdateBlocklistAddressesRequest.md) | The request body for adding or removing multiple addresses on the blocklist. | 

### Return type

[**TokenizationOperationResponse**](TokenizationOperationResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

