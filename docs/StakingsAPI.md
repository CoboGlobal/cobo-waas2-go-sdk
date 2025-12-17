# \StakingsAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBabylonAirdropRegistration**](StakingsAPI.md#CreateBabylonAirdropRegistration) | **Post** /stakings/protocols/babylon/airdrops/registrations | Register for Babylon airdrop
[**CreateBabylonStakingExpansion**](StakingsAPI.md#CreateBabylonStakingExpansion) | **Post** /stakings/protocols/babylon/stakings/expansions | Expand Babylon BTC staking
[**CreateBabylonStakingRegistration**](StakingsAPI.md#CreateBabylonStakingRegistration) | **Post** /stakings/protocols/babylon/stakings/registrations | Register for Babylon Phase-2
[**CreateClaimActivity**](StakingsAPI.md#CreateClaimActivity) | **Post** /stakings/activities/claim | Create claim activity
[**CreateStakeActivity**](StakingsAPI.md#CreateStakeActivity) | **Post** /stakings/activities/stake | Create stake activity
[**CreateUnstakeActivity**](StakingsAPI.md#CreateUnstakeActivity) | **Post** /stakings/activities/unstake | Create unstake activity
[**CreateWithdrawActivity**](StakingsAPI.md#CreateWithdrawActivity) | **Post** /stakings/activities/withdraw | Create withdraw activity
[**GetBabylonAirdropRegistrationById**](StakingsAPI.md#GetBabylonAirdropRegistrationById) | **Get** /stakings/protocols/babylon/airdrops/registrations/{registration_id} | Get Babylon airdrop registration details
[**GetBabylonStakingRegistrationById**](StakingsAPI.md#GetBabylonStakingRegistrationById) | **Get** /stakings/protocols/babylon/stakings/registrations/{registration_id} | Get Babylon Phase-2 registration details
[**GetStakingActivityById**](StakingsAPI.md#GetStakingActivityById) | **Get** /stakings/activities/{activity_id} | Get staking activity details
[**GetStakingById**](StakingsAPI.md#GetStakingById) | **Get** /stakings/{staking_id} | Get staking position details
[**GetStakingEstimationFee**](StakingsAPI.md#GetStakingEstimationFee) | **Post** /stakings/estimate_fee | Estimate staking fees
[**GetStakingEstimationFeeV2**](StakingsAPI.md#GetStakingEstimationFeeV2) | **Post** /stakings/estimate_fee_v2 | Estimate staking fees v2
[**GetStakingPoolById**](StakingsAPI.md#GetStakingPoolById) | **Get** /stakings/pools/{pool_id} | Get staking pool details
[**ListBabylonAirdropRegistrations**](StakingsAPI.md#ListBabylonAirdropRegistrations) | **Get** /stakings/protocols/babylon/airdrops/registrations | List Babylon airdrop registrations
[**ListBabylonEligibleAirdrops**](StakingsAPI.md#ListBabylonEligibleAirdrops) | **Get** /stakings/protocols/babylon/airdrops/eligibles | List wallets eligible for Babylon airdrop
[**ListBabylonEligibleStakings**](StakingsAPI.md#ListBabylonEligibleStakings) | **Get** /stakings/protocols/babylon/stakings/eligibles | List staking positions eligible for Babylon Phase-2
[**ListBabylonStakingRegistrations**](StakingsAPI.md#ListBabylonStakingRegistrations) | **Get** /stakings/protocols/babylon/stakings/registrations | List Babylon Phase-2 registrations
[**ListStakingActivities**](StakingsAPI.md#ListStakingActivities) | **Get** /stakings/activities | List staking activities
[**ListStakingPools**](StakingsAPI.md#ListStakingPools) | **Get** /stakings/pools | List staking pools
[**ListStakings**](StakingsAPI.md#ListStakings) | **Get** /stakings | List staking positions



## CreateBabylonAirdropRegistration

> CreateBabylonAirdropRegistration201Response CreateBabylonAirdropRegistration(ctx).CreateBabylonAirdropRegistrationRequest(createBabylonAirdropRegistrationRequest).Execute()

Register for Babylon airdrop



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
	createBabylonAirdropRegistrationRequest := *coboWaas2.NewCreateBabylonAirdropRegistrationRequest()

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
	resp, r, err := apiClient.StakingsAPI.CreateBabylonAirdropRegistration(ctx).CreateBabylonAirdropRegistrationRequest(createBabylonAirdropRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.CreateBabylonAirdropRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBabylonAirdropRegistration`: CreateBabylonAirdropRegistration201Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.CreateBabylonAirdropRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBabylonAirdropRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBabylonAirdropRegistrationRequest** | [**CreateBabylonAirdropRegistrationRequest**](CreateBabylonAirdropRegistrationRequest.md) | The request body to register for the Babylon airdrop. | 

### Return type

[**CreateBabylonAirdropRegistration201Response**](CreateBabylonAirdropRegistration201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBabylonStakingExpansion

> Stakings CreateBabylonStakingExpansion(ctx).BabylonCreateStakingExpansion(babylonCreateStakingExpansion).Execute()

Expand Babylon BTC staking



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
	babylonCreateStakingExpansion := *coboWaas2.NewBabylonCreateStakingExpansion("3f2840ce-44eb-450b-aa81-d3f84b772efb", []string{"eca1b104dce16c30705f4147a9c4a373ac88646c5d1bcda6a89c018940cb96a0"}, coboWaas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: coboWaas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", coboWaas2.FeeType("Fixed"), "ETH")})

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
	resp, r, err := apiClient.StakingsAPI.CreateBabylonStakingExpansion(ctx).BabylonCreateStakingExpansion(babylonCreateStakingExpansion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.CreateBabylonStakingExpansion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBabylonStakingExpansion`: Stakings
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.CreateBabylonStakingExpansion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBabylonStakingExpansionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **babylonCreateStakingExpansion** | [**BabylonCreateStakingExpansion**](BabylonCreateStakingExpansion.md) | The request body to expand Babylon BTC staking to Phase-3 | 

### Return type

[**Stakings**](Stakings.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBabylonStakingRegistration

> CreateBabylonStakingRegistration201Response CreateBabylonStakingRegistration(ctx).CreateBabylonStakingRegistrationRequest(createBabylonStakingRegistrationRequest).Execute()

Register for Babylon Phase-2



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
	createBabylonStakingRegistrationRequest := *coboWaas2.NewCreateBabylonStakingRegistrationRequest()

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
	resp, r, err := apiClient.StakingsAPI.CreateBabylonStakingRegistration(ctx).CreateBabylonStakingRegistrationRequest(createBabylonStakingRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.CreateBabylonStakingRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBabylonStakingRegistration`: CreateBabylonStakingRegistration201Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.CreateBabylonStakingRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBabylonStakingRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBabylonStakingRegistrationRequest** | [**CreateBabylonStakingRegistrationRequest**](CreateBabylonStakingRegistrationRequest.md) | The request body to transit Babylon BTC staking to Phase-2 | 

### Return type

[**CreateBabylonStakingRegistration201Response**](CreateBabylonStakingRegistration201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClaimActivity

> CreateStakeActivity201Response CreateClaimActivity(ctx).CreateClaimActivityRequest(createClaimActivityRequest).Execute()

Create claim activity



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
	createClaimActivityRequest := *coboWaas2.NewCreateClaimActivityRequest("f47ac10b-58cc-4372-a567-0e02b2c3d479")

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
	resp, r, err := apiClient.StakingsAPI.CreateClaimActivity(ctx).CreateClaimActivityRequest(createClaimActivityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.CreateClaimActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClaimActivity`: CreateStakeActivity201Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.CreateClaimActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClaimActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createClaimActivityRequest** | [**CreateClaimActivityRequest**](CreateClaimActivityRequest.md) | The request body to create a staking request. | 

### Return type

[**CreateStakeActivity201Response**](CreateStakeActivity201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStakeActivity

> CreateStakeActivity201Response CreateStakeActivity(ctx).CreateStakeActivityRequest(createStakeActivityRequest).Execute()

Create stake activity



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
	createStakeActivityRequest := *coboWaas2.NewCreateStakeActivityRequest(coboWaas2.StakingPoolId("babylon_btc_signet"), "100.00", coboWaas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: coboWaas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", coboWaas2.FeeType("Fixed"), "ETH")})

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
	resp, r, err := apiClient.StakingsAPI.CreateStakeActivity(ctx).CreateStakeActivityRequest(createStakeActivityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.CreateStakeActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStakeActivity`: CreateStakeActivity201Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.CreateStakeActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStakeActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStakeActivityRequest** | [**CreateStakeActivityRequest**](CreateStakeActivityRequest.md) | The request body to create a staking request. | 

### Return type

[**CreateStakeActivity201Response**](CreateStakeActivity201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUnstakeActivity

> CreateStakeActivity201Response CreateUnstakeActivity(ctx).CreateUnstakeActivityRequest(createUnstakeActivityRequest).Execute()

Create unstake activity



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
	createUnstakeActivityRequest := *coboWaas2.NewCreateUnstakeActivityRequest("0011039d-27fb-49ba-b172-6e0aa80e37ec")

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
	resp, r, err := apiClient.StakingsAPI.CreateUnstakeActivity(ctx).CreateUnstakeActivityRequest(createUnstakeActivityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.CreateUnstakeActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUnstakeActivity`: CreateStakeActivity201Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.CreateUnstakeActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUnstakeActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUnstakeActivityRequest** | [**CreateUnstakeActivityRequest**](CreateUnstakeActivityRequest.md) | The request body to create a unstaking request. | 

### Return type

[**CreateStakeActivity201Response**](CreateStakeActivity201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWithdrawActivity

> CreateStakeActivity201Response CreateWithdrawActivity(ctx).CreateWithdrawActivityRequest(createWithdrawActivityRequest).Execute()

Create withdraw activity



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
	createWithdrawActivityRequest := *coboWaas2.NewCreateWithdrawActivityRequest("0011039d-27fb-49ba-b172-6e0aa80e37ec", coboWaas2.TransactionRequestFee{TransactionRequestEvmEip1559Fee: coboWaas2.NewTransactionRequestEvmEip1559Fee("9000000000000", "1000000000000", coboWaas2.FeeType("Fixed"), "ETH")})

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
	resp, r, err := apiClient.StakingsAPI.CreateWithdrawActivity(ctx).CreateWithdrawActivityRequest(createWithdrawActivityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.CreateWithdrawActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWithdrawActivity`: CreateStakeActivity201Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.CreateWithdrawActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWithdrawActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWithdrawActivityRequest** | [**CreateWithdrawActivityRequest**](CreateWithdrawActivityRequest.md) | The request body to create a withdraw activity. | 

### Return type

[**CreateStakeActivity201Response**](CreateStakeActivity201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBabylonAirdropRegistrationById

> BabylonAirdropRegistration GetBabylonAirdropRegistrationById(ctx, registrationId).Execute()

Get Babylon airdrop registration details



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
	registrationId := "registrationId_example"

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
	resp, r, err := apiClient.StakingsAPI.GetBabylonAirdropRegistrationById(ctx, registrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.GetBabylonAirdropRegistrationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBabylonAirdropRegistrationById`: BabylonAirdropRegistration
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.GetBabylonAirdropRegistrationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**registrationId** | **string** | The Babylon airdrop or Babylon Phase-2 registration ID. You can use the [Register for Babylon airdrop](https://www.cobo.com/developers/v2/api-references/stakings/register-for-babylon-airdrop) or the [Register for Babylon Phase-2](https://www.cobo.com/developers/v2/api-references/stakings/register-for-babylon-phase-2) operation to get this information. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBabylonAirdropRegistrationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BabylonAirdropRegistration**](BabylonAirdropRegistration.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBabylonStakingRegistrationById

> BabylonStakingRegistration GetBabylonStakingRegistrationById(ctx, registrationId).Execute()

Get Babylon Phase-2 registration details



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
	registrationId := "registrationId_example"

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
	resp, r, err := apiClient.StakingsAPI.GetBabylonStakingRegistrationById(ctx, registrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.GetBabylonStakingRegistrationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBabylonStakingRegistrationById`: BabylonStakingRegistration
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.GetBabylonStakingRegistrationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**registrationId** | **string** | The Babylon airdrop or Babylon Phase-2 registration ID. You can use the [Register for Babylon airdrop](https://www.cobo.com/developers/v2/api-references/stakings/register-for-babylon-airdrop) or the [Register for Babylon Phase-2](https://www.cobo.com/developers/v2/api-references/stakings/register-for-babylon-phase-2) operation to get this information. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBabylonStakingRegistrationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BabylonStakingRegistration**](BabylonStakingRegistration.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStakingActivityById

> Activity GetStakingActivityById(ctx, activityId).Execute()

Get staking activity details



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
	activityId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"

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
	resp, r, err := apiClient.StakingsAPI.GetStakingActivityById(ctx, activityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.GetStakingActivityById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStakingActivityById`: Activity
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.GetStakingActivityById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**activityId** | **string** | The activity ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStakingActivityByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Activity**](Activity.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStakingById

> Stakings GetStakingById(ctx, stakingId).Execute()

Get staking position details



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
	stakingId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"

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
	resp, r, err := apiClient.StakingsAPI.GetStakingById(ctx, stakingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.GetStakingById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStakingById`: Stakings
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.GetStakingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**stakingId** | **string** | The ID of the staking position. You can retrieve a list of staking positions by calling [List staking positions](https://www.cobo.com/developers/v2/api-references/stakings/list-staking-positions). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStakingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Stakings**](Stakings.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStakingEstimationFee

> GetStakingEstimationFee201Response GetStakingEstimationFee(ctx).GetStakingEstimationFeeRequest(getStakingEstimationFeeRequest).Execute()

Estimate staking fees



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
	getStakingEstimationFeeRequest := coboWaas2.get_staking_estimation_fee_request{EstimateClaimFee: coboWaas2.NewEstimateClaimFee(coboWaas2.ActivityType("Stake"))}

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
	resp, r, err := apiClient.StakingsAPI.GetStakingEstimationFee(ctx).GetStakingEstimationFeeRequest(getStakingEstimationFeeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.GetStakingEstimationFee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStakingEstimationFee`: GetStakingEstimationFee201Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.GetStakingEstimationFee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStakingEstimationFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getStakingEstimationFeeRequest** | [**GetStakingEstimationFeeRequest**](GetStakingEstimationFeeRequest.md) | The request body to get the estimated fee of a staking activity. | 

### Return type

[**GetStakingEstimationFee201Response**](GetStakingEstimationFee201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStakingEstimationFeeV2

> EthStakeEstimatedFee GetStakingEstimationFeeV2(ctx).GetStakingEstimationFeeRequest(getStakingEstimationFeeRequest).Execute()

Estimate staking fees v2



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
	getStakingEstimationFeeRequest := coboWaas2.get_staking_estimation_fee_request{EstimateClaimFee: coboWaas2.NewEstimateClaimFee(coboWaas2.ActivityType("Stake"))}

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
	resp, r, err := apiClient.StakingsAPI.GetStakingEstimationFeeV2(ctx).GetStakingEstimationFeeRequest(getStakingEstimationFeeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.GetStakingEstimationFeeV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStakingEstimationFeeV2`: EthStakeEstimatedFee
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.GetStakingEstimationFeeV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStakingEstimationFeeV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getStakingEstimationFeeRequest** | [**GetStakingEstimationFeeRequest**](GetStakingEstimationFeeRequest.md) | The request body to get the estimated fee of a staking activity. | 

### Return type

[**EthStakeEstimatedFee**](EthStakeEstimatedFee.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStakingPoolById

> PoolDetails GetStakingPoolById(ctx, poolId).Execute()

Get staking pool details



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
	poolId := "babylon_btc"

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
	resp, r, err := apiClient.StakingsAPI.GetStakingPoolById(ctx, poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.GetStakingPoolById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStakingPoolById`: PoolDetails
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.GetStakingPoolById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**poolId** | **string** | The ID of the staking pool. A staking pool is a pairing of a staking protocol and a specific type of token. You can call [List staking pools](https://www.cobo.com/developers/v2/api-references/stakings/list-staking-pools) to retrieve a list of staking pools. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStakingPoolByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoolDetails**](PoolDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBabylonAirdropRegistrations

> ListBabylonAirdropRegistrations200Response ListBabylonAirdropRegistrations(ctx).Status(status).BtcAddress(btcAddress).Limit(limit).Before(before).After(after).Execute()

List Babylon airdrop registrations



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
	status := "Processing"
	btcAddress := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
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
	resp, r, err := apiClient.StakingsAPI.ListBabylonAirdropRegistrations(ctx).Status(status).BtcAddress(btcAddress).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.ListBabylonAirdropRegistrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBabylonAirdropRegistrations`: ListBabylonAirdropRegistrations200Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.ListBabylonAirdropRegistrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBabylonAirdropRegistrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | The registration request status. | 
 **btcAddress** | **string** | The Bitcoin (BTC) address used for staking. | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 

### Return type

[**ListBabylonAirdropRegistrations200Response**](ListBabylonAirdropRegistrations200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBabylonEligibleAirdrops

> ListBabylonEligibleAirdrops200Response ListBabylonEligibleAirdrops(ctx).Status(status).Limit(limit).Before(before).After(after).Execute()

List wallets eligible for Babylon airdrop



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
	status := "Registered"
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
	resp, r, err := apiClient.StakingsAPI.ListBabylonEligibleAirdrops(ctx).Status(status).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.ListBabylonEligibleAirdrops``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBabylonEligibleAirdrops`: ListBabylonEligibleAirdrops200Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.ListBabylonEligibleAirdrops`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBabylonEligibleAirdropsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | The status of Babylon airdrop or Phase-2 registration. Possible values are: - &#x60;Registered&#x60;: Registered for Babylon airdrop or Phase-2. - &#x60;Unregistered&#x60;: Not registered for any Babylon airdrop or Phase-2. - &#x60;Registering&#x60;: The Babylon airdrop or Phase-2 registration is in progress but not yet completed.  | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 

### Return type

[**ListBabylonEligibleAirdrops200Response**](ListBabylonEligibleAirdrops200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBabylonEligibleStakings

> ListBabylonEligibleStakings200Response ListBabylonEligibleStakings(ctx).Status(status).Limit(limit).Before(before).After(after).Execute()

List staking positions eligible for Babylon Phase-2



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
	status := "Registered"
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
	resp, r, err := apiClient.StakingsAPI.ListBabylonEligibleStakings(ctx).Status(status).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.ListBabylonEligibleStakings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBabylonEligibleStakings`: ListBabylonEligibleStakings200Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.ListBabylonEligibleStakings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBabylonEligibleStakingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | The status of Babylon airdrop or Phase-2 registration. Possible values are: - &#x60;Registered&#x60;: Registered for Babylon airdrop or Phase-2. - &#x60;Unregistered&#x60;: Not registered for any Babylon airdrop or Phase-2. - &#x60;Registering&#x60;: The Babylon airdrop or Phase-2 registration is in progress but not yet completed.  | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 

### Return type

[**ListBabylonEligibleStakings200Response**](ListBabylonEligibleStakings200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBabylonStakingRegistrations

> ListBabylonStakingRegistrations200Response ListBabylonStakingRegistrations(ctx).Status(status).StakingId(stakingId).Limit(limit).Before(before).After(after).Execute()

List Babylon Phase-2 registrations



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
	status := "Processing"
	stakingId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
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
	resp, r, err := apiClient.StakingsAPI.ListBabylonStakingRegistrations(ctx).Status(status).StakingId(stakingId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.ListBabylonStakingRegistrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBabylonStakingRegistrations`: ListBabylonStakingRegistrations200Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.ListBabylonStakingRegistrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBabylonStakingRegistrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | The registration request status. | 
 **stakingId** | **string** | The ID of the Phase-1 BTC staking position. | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 

### Return type

[**ListBabylonStakingRegistrations200Response**](ListBabylonStakingRegistrations200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStakingActivities

> ListStakingActivities200Response ListStakingActivities(ctx).PoolId(poolId).StakingId(stakingId).ActivityType(activityType).ActivityStatus(activityStatus).MinModifiedTimestamp(minModifiedTimestamp).MaxModifiedTimestamp(maxModifiedTimestamp).Initiator(initiator).RequestId(requestId).Limit(limit).Before(before).After(after).Execute()

List staking activities



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
	poolId := "babylon_btc"
	stakingId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	activityType := coboWaas2.ActivityType("Stake")
	activityStatus := coboWaas2.ActivityStatus("Success")
	minModifiedTimestamp := int64(1635744000000)
	maxModifiedTimestamp := int64(1635744000000)
	initiator := "steve@example.com"
	requestId := "web_send_by_user_327_1610444045047"
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
	resp, r, err := apiClient.StakingsAPI.ListStakingActivities(ctx).PoolId(poolId).StakingId(stakingId).ActivityType(activityType).ActivityStatus(activityStatus).MinModifiedTimestamp(minModifiedTimestamp).MaxModifiedTimestamp(maxModifiedTimestamp).Initiator(initiator).RequestId(requestId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.ListStakingActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStakingActivities`: ListStakingActivities200Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.ListStakingActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStakingActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolId** | **string** | The ID of the staking pool. A staking pool is a pairing of a staking protocol and a specific type of token. You can call [List staking pools](https://www.cobo.com/developers/v2/api-references/stakings/list-staking-pools) to retrieve a list of staking pools. | 
 **stakingId** | **string** | The ID of the Phase-1 BTC staking position. | 
 **activityType** | [**ActivityType**](ActivityType.md) |  | 
 **activityStatus** | [**ActivityStatus**](ActivityStatus.md) |  | 
 **minModifiedTimestamp** | **int64** | The start time of the query. All staking activities updated after the specified time will be retrieved. The time is in Unix timestamp format, measured in milliseconds. | 
 **maxModifiedTimestamp** | **int64** | The end time of the query. All staking activities updated before the specified time will be retrieved. The time is in Unix timestamp format, measured in milliseconds. | 
 **initiator** | **string** | The activity initiator, which is your API key by default. You can also specify the initiator when creating the activity. | 
 **requestId** | **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 

### Return type

[**ListStakingActivities200Response**](ListStakingActivities200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStakingPools

> ListStakingPools200Response ListStakingPools(ctx).ChainId(chainId).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()

List staking pools



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
	resp, r, err := apiClient.StakingsAPI.ListStakingPools(ctx).ChainId(chainId).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.ListStakingPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStakingPools`: ListStakingPools200Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.ListStakingPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStakingPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | 
 **tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 

### Return type

[**ListStakingPools200Response**](ListStakingPools200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStakings

> ListStakings200Response ListStakings(ctx).PoolId(poolId).Statuses(statuses).WalletId(walletId).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()

List staking positions



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
	poolId := "babylon_btc"
	statuses := "Active,StakeInProgress,"
	walletId := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
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
	resp, r, err := apiClient.StakingsAPI.ListStakings(ctx).PoolId(poolId).Statuses(statuses).WalletId(walletId).TokenId(tokenId).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.ListStakings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStakings`: ListStakings200Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.ListStakings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStakingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolId** | **string** | The ID of the staking pool. A staking pool is a pairing of a staking protocol and a specific type of token. You can call [List staking pools](https://www.cobo.com/developers/v2/api-references/stakings/list-staking-pools) to retrieve a list of staking pools. | 
 **statuses** | **string** | The statuses of the staking amounts, separated by comma. Possible values include:  - &#x60;StakeInProgress&#x60;: The staking request is submitted and is waiting to be confirmed by the staking protocol. - &#x60;Active&#x60;: The amount has been staked. - &#x60;Rejected&#x60;: The staking request has been rejected because the signer refuses to sign the transaction. - &#x60;LimitExceeded&#x60;: The total staking cap of the staking protocol has been reached. - &#x60;Invalid&#x60;: The staking request is invalid. This is often due to the failure to broadcast the transaction. - &#x60;UnstakeInProgress&#x60;: The unstaking request is submitted and is waiting to be confirmed by the staking protocol. - &#x60;Withdrawable&#x60;: The tokens have been unstaked and are ready to be withdrawn. - &#x60;WithdrawInProgress&#x60;: The withdrawal request is submitted and is waiting to be confirmed on the chain network. - &#x60;Closed&#x60;: The staking position is closed.  | 
 **walletId** | **string** | The wallet ID. | 
 **tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
 **limit** | **int32** | The maximum number of objects to return. For most operations, the value range is [1, 50]. | [default to 10]
 **before** | **string** | A cursor indicating the position before the current page. This value is generated by Cobo and returned in the response. If you are paginating forward from the beginning, you do not need to provide it on the first request. When paginating backward (to the previous page), you should pass the before value returned from the last response.  | 
 **after** | **string** | A cursor indicating the position after the current page. This value is generated by Cobo and returned in the response. You do not need to provide it on the first request. When paginating forward (to the next page), you should pass the after value returned from the last response.  | 

### Return type

[**ListStakings200Response**](ListStakings200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

