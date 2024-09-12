# GetStakingEstimationFeeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityType** | [**ActivityType**](ActivityType.md) |  | 
**RequestId** | Pointer to **string** | The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization. | [optional] 
**Source** | Pointer to [**StakingSource**](StakingSource.md) |  | [optional] 
**PoolId** | **string** | The ID of the staking pool. | 
**Amount** | **string** | The amount to withdraw. | 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 
**Extra** | [**CreateStakeActivityExtra**](CreateStakeActivityExtra.md) |  | 
**StakingId** | **string** | The ID of the corresponding staking position. | 
**Address** | **string** | The withdrawal address. | 

## Methods

### NewGetStakingEstimationFeeRequest

`func NewGetStakingEstimationFeeRequest(activityType ActivityType, poolId string, amount string, fee TransactionRequestFee, extra CreateStakeActivityExtra, stakingId string, address string, ) *GetStakingEstimationFeeRequest`

NewGetStakingEstimationFeeRequest instantiates a new GetStakingEstimationFeeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStakingEstimationFeeRequestWithDefaults

`func NewGetStakingEstimationFeeRequestWithDefaults() *GetStakingEstimationFeeRequest`

NewGetStakingEstimationFeeRequestWithDefaults instantiates a new GetStakingEstimationFeeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityType

`func (o *GetStakingEstimationFeeRequest) GetActivityType() ActivityType`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *GetStakingEstimationFeeRequest) GetActivityTypeOk() (*ActivityType, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *GetStakingEstimationFeeRequest) SetActivityType(v ActivityType)`

SetActivityType sets ActivityType field to given value.


### GetRequestId

`func (o *GetStakingEstimationFeeRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetStakingEstimationFeeRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetStakingEstimationFeeRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GetStakingEstimationFeeRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetSource

`func (o *GetStakingEstimationFeeRequest) GetSource() StakingSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetStakingEstimationFeeRequest) GetSourceOk() (*StakingSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetStakingEstimationFeeRequest) SetSource(v StakingSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetStakingEstimationFeeRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPoolId

`func (o *GetStakingEstimationFeeRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *GetStakingEstimationFeeRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *GetStakingEstimationFeeRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetAmount

`func (o *GetStakingEstimationFeeRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetStakingEstimationFeeRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetStakingEstimationFeeRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFee

`func (o *GetStakingEstimationFeeRequest) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetStakingEstimationFeeRequest) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetStakingEstimationFeeRequest) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.


### GetExtra

`func (o *GetStakingEstimationFeeRequest) GetExtra() CreateStakeActivityExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *GetStakingEstimationFeeRequest) GetExtraOk() (*CreateStakeActivityExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *GetStakingEstimationFeeRequest) SetExtra(v CreateStakeActivityExtra)`

SetExtra sets Extra field to given value.


### GetStakingId

`func (o *GetStakingEstimationFeeRequest) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *GetStakingEstimationFeeRequest) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *GetStakingEstimationFeeRequest) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.


### GetAddress

`func (o *GetStakingEstimationFeeRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetStakingEstimationFeeRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetStakingEstimationFeeRequest) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


