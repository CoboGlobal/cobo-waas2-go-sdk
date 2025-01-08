# \TravelRuleAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransactionLimitation**](TravelRuleAPI.md#GetTransactionLimitation) | **Get** /travel_rule/transaction/limitation | Retrieve transaction limitations
[**ListSupportedCountries**](TravelRuleAPI.md#ListSupportedCountries) | **Get** /travel_rule/transaction/countries | List supported countries
[**SubmitDepositTravelRuleInfo**](TravelRuleAPI.md#SubmitDepositTravelRuleInfo) | **Post** /travel_rule/transaction/deposit/travel_rule_info | Submit Deposit Transaction Travel Rule information
[**SubmitWithdrawTravelRuleInfo**](TravelRuleAPI.md#SubmitWithdrawTravelRuleInfo) | **Post** /travel_rule/transaction/withdraw/travel_rule_info | Submit Withdraw Transaction Travel Rule information



## GetTransactionLimitation

> GetTransactionLimitation200Response GetTransactionLimitation(ctx).TransactionType(transactionType).TransactionId(transactionId).Execute()

Retrieve transaction limitations



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
	transactionType := "DEPOSIT"
	transactionId := "123e4567-e89b-12d3-a456-426614174000"

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
	resp, r, err := apiClient.TravelRuleAPI.GetTransactionLimitation(ctx).TransactionType(transactionType).TransactionId(transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TravelRuleAPI.GetTransactionLimitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionLimitation`: GetTransactionLimitation200Response
	fmt.Fprintf(os.Stdout, "Response from `TravelRuleAPI.GetTransactionLimitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionLimitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionType** | **string** | The transaction type. Possible values include:    - &#x60;DEPOSIT&#x60;: A deposit transaction.   - &#x60;WITHDRAW&#x60;: A withdrawal transaction.  | 
 **transactionId** | **string** | The transaction ID | 

### Return type

[**GetTransactionLimitation200Response**](GetTransactionLimitation200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportedCountries

> []ListSupportedCountries200ResponseInner ListSupportedCountries(ctx).Execute()

List supported countries



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
	resp, r, err := apiClient.TravelRuleAPI.ListSupportedCountries(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TravelRuleAPI.ListSupportedCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSupportedCountries`: []ListSupportedCountries200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TravelRuleAPI.ListSupportedCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportedCountriesRequest struct via the builder pattern


### Return type

[**[]ListSupportedCountries200ResponseInner**](ListSupportedCountries200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitDepositTravelRuleInfo

> SubmitDepositTravelRuleInfo201Response SubmitDepositTravelRuleInfo(ctx).TravelRuleDepositRequest(travelRuleDepositRequest).Execute()

Submit Deposit Transaction Travel Rule information



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
	travelRuleDepositRequest := *coboWaas2.NewTravelRuleDepositRequest("f47ac10b-58cc-4372-a567-0e02b2c3d479", coboWaas2.TravelRuleDepositRequest_travel_rule_info{SelfCustodyWallet: coboWaas2.NewSelfCustodyWallet(coboWaas2.DestinationWalletType("EXCHANGES_OR_VASP"), "challenge_token_abc123", "0x1234567890abcdef1234567890abcdef12345678", "0xf0a0ca69dd3afc57235c72aba3ff1f1144ee5409aeec013a9b17cdb58d0185a66a525945bfbd66e87bf0503eb0b83bf90cb973a8cbb730d19dc032e00dfe393a1c")})

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
	resp, r, err := apiClient.TravelRuleAPI.SubmitDepositTravelRuleInfo(ctx).TravelRuleDepositRequest(travelRuleDepositRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TravelRuleAPI.SubmitDepositTravelRuleInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitDepositTravelRuleInfo`: SubmitDepositTravelRuleInfo201Response
	fmt.Fprintf(os.Stdout, "Response from `TravelRuleAPI.SubmitDepositTravelRuleInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitDepositTravelRuleInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **travelRuleDepositRequest** | [**TravelRuleDepositRequest**](TravelRuleDepositRequest.md) |  | 

### Return type

[**SubmitDepositTravelRuleInfo201Response**](SubmitDepositTravelRuleInfo201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitWithdrawTravelRuleInfo

> SubmitDepositTravelRuleInfo201Response SubmitWithdrawTravelRuleInfo(ctx).TravelRuleWithdrawRequest(travelRuleWithdrawRequest).Execute()

Submit Withdraw Transaction Travel Rule information



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
	travelRuleWithdrawRequest := *coboWaas2.NewTravelRuleWithdrawRequest("f47ac10b-58cc-4372-a567-0e02b2c3d479", coboWaas2.TravelRuleWithdrawRequest_travel_rule_info{SelfCustodyWallet: coboWaas2.NewSelfCustodyWallet(coboWaas2.DestinationWalletType("EXCHANGES_OR_VASP"), "challenge_token_abc123", "0x1234567890abcdef1234567890abcdef12345678", "0xf0a0ca69dd3afc57235c72aba3ff1f1144ee5409aeec013a9b17cdb58d0185a66a525945bfbd66e87bf0503eb0b83bf90cb973a8cbb730d19dc032e00dfe393a1c")})

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
	resp, r, err := apiClient.TravelRuleAPI.SubmitWithdrawTravelRuleInfo(ctx).TravelRuleWithdrawRequest(travelRuleWithdrawRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TravelRuleAPI.SubmitWithdrawTravelRuleInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitWithdrawTravelRuleInfo`: SubmitDepositTravelRuleInfo201Response
	fmt.Fprintf(os.Stdout, "Response from `TravelRuleAPI.SubmitWithdrawTravelRuleInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitWithdrawTravelRuleInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **travelRuleWithdrawRequest** | [**TravelRuleWithdrawRequest**](TravelRuleWithdrawRequest.md) |  | 

### Return type

[**SubmitDepositTravelRuleInfo201Response**](SubmitDepositTravelRuleInfo201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

