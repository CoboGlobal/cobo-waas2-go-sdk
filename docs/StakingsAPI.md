# \StakingsAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStakeActivity**](StakingsAPI.md#CreateStakeActivity) | **Post** /stakings/activities/stake | Create staking activity
[**CreateUnstakeActivity**](StakingsAPI.md#CreateUnstakeActivity) | **Post** /stakings/activities/unstake | Create unstake activity
[**CreateWithdrawActivity**](StakingsAPI.md#CreateWithdrawActivity) | **Post** /stakings/activities/withdraw | Create withdraw request
[**GetActivityById**](StakingsAPI.md#GetActivityById) | **Get** /stakings/activities/{activity_id} | Get activity details
[**GetStakingById**](StakingsAPI.md#GetStakingById) | **Get** /stakings/{staking_id} | Get staking by id
[**GetStakingEstimationFee**](StakingsAPI.md#GetStakingEstimationFee) | **Post** /stakings/estimate_fee | Estimate staking transaction fee
[**GetStakingPoolById**](StakingsAPI.md#GetStakingPoolById) | **Get** /stakings/pools/{pool_id} | Get staking pool details
[**ListActivities**](StakingsAPI.md#ListActivities) | **Get** /stakings/activities | List activities
[**ListStakingPools**](StakingsAPI.md#ListStakingPools) | **Get** /stakings/pools | List staking pools
[**ListStakings**](StakingsAPI.md#ListStakings) | **Get** /stakings | List all stakings



## CreateStakeActivity

> CreateStakeActivity201Response CreateStakeActivity(ctx).CreateStakeActivityRequest(createStakeActivityRequest).Execute()

Create staking activity



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
	createStakeActivityRequest := *CoboWaas2.NewCreateStakeActivityRequest("0111039d-27fb-49ba-b172-6e0aa80e37ec", "0x0000000000000000000000000000000000000000", "babylon_btc", "100.00", CoboWaas2.TransactionTransferFee{EvmEip1559TransactionFee: CoboWaas2.NewEvmEip1559TransactionFee("0.1", "0.9", "21000", CoboWaas2.FeeType("Fixed"), "ETH")}, CoboWaas2.CreateStakeActivity_extra{BabylonStakeExtra: CoboWaas2.NewBabylonStakeExtra(CoboWaas2.StakingPoolType("Babylon"), "0000000000000000000000000000000000000000000000000000000000000000")}) // CreateStakeActivityRequest | The request body to create a staking activity. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
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
 **createStakeActivityRequest** | [**CreateStakeActivityRequest**](CreateStakeActivityRequest.md) | The request body to create a staking activity. | 

### Return type

[**CreateStakeActivity201Response**](CreateStakeActivity201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

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
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	createUnstakeActivityRequest := *CoboWaas2.NewCreateUnstakeActivityRequest("0011039d-27fb-49ba-b172-6e0aa80e37ec", CoboWaas2.TransactionTransferFee{EvmEip1559TransactionFee: CoboWaas2.NewEvmEip1559TransactionFee("0.1", "0.9", "21000", CoboWaas2.FeeType("Fixed"), "ETH")}) // CreateUnstakeActivityRequest | The request body to create a unstake activity. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
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
 **createUnstakeActivityRequest** | [**CreateUnstakeActivityRequest**](CreateUnstakeActivityRequest.md) | The request body to create a unstake activity. | 

### Return type

[**CreateStakeActivity201Response**](CreateStakeActivity201Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWithdrawActivity

> CreateStakeActivity201Response CreateWithdrawActivity(ctx).CreateWithdrawActivityRequest(createWithdrawActivityRequest).Execute()

Create withdraw request



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
	createWithdrawActivityRequest := *CoboWaas2.NewCreateWithdrawActivityRequest("0011039d-27fb-49ba-b172-6e0aa80e37ec", CoboWaas2.TransactionTransferFee{EvmEip1559TransactionFee: CoboWaas2.NewEvmEip1559TransactionFee("0.1", "0.9", "21000", CoboWaas2.FeeType("Fixed"), "ETH")}) // CreateWithdrawActivityRequest | The request body to create a withdraw activity. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
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

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityById

> Activity GetActivityById(ctx, activityId).Execute()

Get activity details



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
	activityId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | activity id

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.StakingsAPI.GetActivityById(ctx, activityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.GetActivityById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityById`: Activity
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.GetActivityById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**activityId** | **string** | activity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Activity**](Activity.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStakingById

> ListStakings200Response GetStakingById(ctx, stakingId).Execute()

Get staking by id



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
	stakingId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | staking id

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.StakingsAPI.GetStakingById(ctx, stakingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.GetStakingById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStakingById`: ListStakings200Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.GetStakingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for ServerHost/Env, Signer, etc.
**stakingId** | **string** | staking id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStakingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListStakings200Response**](ListStakings200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStakingEstimationFee

> []EstimationFee GetStakingEstimationFee(ctx).GetStakingEstimationFeeRequest(getStakingEstimationFeeRequest).Execute()

Estimate staking transaction fee



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
	getStakingEstimationFeeRequest := CoboWaas2.get_staking_estimation_fee_request{EstimateStakeFee: CoboWaas2.NewEstimateStakeFee("0111039d-27fb-49ba-b172-6e0aa80e37ec", "0x0000000000000000000000000000000000000000", "babylon_btc", "100.00", CoboWaas2.TransactionTransferFee{EvmEip1559TransactionFee: CoboWaas2.NewEvmEip1559TransactionFee("0.1", "0.9", "21000", CoboWaas2.FeeType("Fixed"), "ETH")}, CoboWaas2.CreateStakeActivity_extra{BabylonStakeExtra: CoboWaas2.NewBabylonStakeExtra(CoboWaas2.StakingPoolType("Babylon"), "0000000000000000000000000000000000000000000000000000000000000000")})} // GetStakingEstimationFeeRequest | The request body to create a get estimate fee of a staking activity. (optional)

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})
	resp, r, err := apiClient.StakingsAPI.GetStakingEstimationFee(ctx).GetStakingEstimationFeeRequest(getStakingEstimationFeeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.GetStakingEstimationFee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStakingEstimationFee`: []EstimationFee
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.GetStakingEstimationFee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStakingEstimationFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getStakingEstimationFeeRequest** | [**GetStakingEstimationFeeRequest**](GetStakingEstimationFeeRequest.md) | The request body to create a get estimate fee of a staking activity. | 

### Return type

[**[]EstimationFee**](EstimationFee.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

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
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	poolId := "babylon_btc" // string | staking pool id

	configuration := CoboWaas2.NewConfiguration()
	apiClient := CoboWaas2.NewAPIClient(configuration)
	ctx := context.Background()
	// ctx = context.WithValue(ctx, CoboWaas2.ContextServerHost, "https://api[.xxx].cobo.com/v2")
	// ctx = context.WithValue(ctx, CoboWaas2.ContextEnv, CoboWaas2.DevEnv)
	ctx = context.WithValue(ctx, CoboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
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
**poolId** | **string** | staking pool id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStakingPoolByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoolDetails**](PoolDetails.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActivities

> ListActivities200Response ListActivities(ctx).PoolId(poolId).StakingId(stakingId).ActivityType(activityType).ActivityStatus(activityStatus).MinModifiedTimestamp(minModifiedTimestamp).MaxModifiedTimestamp(maxModifiedTimestamp).Initiator(initiator).Limit(limit).Before(before).After(after).Execute()

List activities



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
	poolId := "babylon_btc" // string | staking pool id (optional)
	stakingId := "f47ac10b-58cc-4372-a567-0e02b2c3d479" // string | staking id (optional)
	activityType := CoboWaas2.ActivityType("Stake") // ActivityType | activity type (optional)
	activityStatus := CoboWaas2.ActivityStatus("Success") // ActivityStatus | activity status (optional)
	minModifiedTimestamp := int64(1635744000000) // int64 | The minimum modified timestamp in Unix epoch seconds (optional)
	maxModifiedTimestamp := int64(1635744000000) // int64 | The maximum modified timestamp in Unix epoch seconds (optional)
	initiator := "vanya@cobo.com" // string | activity initiator, maybe email or api key. (optional)
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
	resp, r, err := apiClient.StakingsAPI.ListActivities(ctx).PoolId(poolId).StakingId(stakingId).ActivityType(activityType).ActivityStatus(activityStatus).MinModifiedTimestamp(minModifiedTimestamp).MaxModifiedTimestamp(maxModifiedTimestamp).Initiator(initiator).Limit(limit).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingsAPI.ListActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListActivities`: ListActivities200Response
	fmt.Fprintf(os.Stdout, "Response from `StakingsAPI.ListActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolId** | **string** | staking pool id | 
 **stakingId** | **string** | staking id | 
 **activityType** | [**ActivityType**](ActivityType.md) | activity type | 
 **activityStatus** | [**ActivityStatus**](ActivityStatus.md) | activity status | 
 **minModifiedTimestamp** | **int64** | The minimum modified timestamp in Unix epoch seconds | 
 **maxModifiedTimestamp** | **int64** | The maximum modified timestamp in Unix epoch seconds | 
 **initiator** | **string** | activity initiator, maybe email or api key. | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**ListActivities200Response**](ListActivities200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

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
	CoboWaas2 "github.com/CoboGlobal/cobo-waas2-go-api"
        "github.com/CoboGlobal/cobo-waas2-go-api/crypto"
)

func main() {
	chainId := "ETH" // string | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). (optional)
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
 **chainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List organization enabled chains](/v2/api-references/wallets/list-organization-enabled-chains). | 
 **tokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**ListStakingPools200Response**](ListStakingPools200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStakings

> ListStakings200Response ListStakings(ctx).PoolId(poolId).Limit(limit).Before(before).After(after).Execute()

List all stakings



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
	poolId := "babylon_btc" // string | staking pool id (optional)
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
	resp, r, err := apiClient.StakingsAPI.ListStakings(ctx).PoolId(poolId).Limit(limit).Before(before).After(after).Execute()
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
 **poolId** | **string** | staking pool id | 
 **limit** | **int32** | The maximum number of objects to return. The default value range is [1, 50] and can be set endpoint specified. | [default to 10]
 **before** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;before&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; and &#x60;after&#x60; are both set to empty, the first slice is returned.  | 
 **after** | **string** | An object ID which serves as a cursor for pagination. For example, if you specify &#x60;after&#x60; as &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;8f2e919a-6a7b-4a9b-8c1a-4c0b3f5b8b1f&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  If you set both &#x60;after&#x60; or &#x60;before&#x60;, only the setting of &#x60;before&#x60; will take effect.  If the &#x60;before&#x60; is set to empty and &#x60;after&#x60; is set to &#x60;last&#x60;, the last slice is returned.  | 

### Return type

[**ListStakings200Response**](ListStakings200Response.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

