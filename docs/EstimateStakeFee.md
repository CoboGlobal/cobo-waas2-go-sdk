# EstimateStakeFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityType** | [**ActivityType**](ActivityType.md) |  | 
**RequestId** | Pointer to **string** | The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization. | [optional] 
**Source** | Pointer to [**StakingSource**](StakingSource.md) |  | [optional] 
**PoolId** | [**StakingPoolId**](StakingPoolId.md) |  | 
**Amount** | **string** | The amount to stake. | 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 
**Extra** | Pointer to [**CreateStakeActivityExtra**](CreateStakeActivityExtra.md) |  | [optional] 

## Methods

### NewEstimateStakeFee

`func NewEstimateStakeFee(activityType ActivityType, poolId StakingPoolId, amount string, fee TransactionRequestFee, ) *EstimateStakeFee`

NewEstimateStakeFee instantiates a new EstimateStakeFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateStakeFeeWithDefaults

`func NewEstimateStakeFeeWithDefaults() *EstimateStakeFee`

NewEstimateStakeFeeWithDefaults instantiates a new EstimateStakeFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityType

`func (o *EstimateStakeFee) GetActivityType() ActivityType`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *EstimateStakeFee) GetActivityTypeOk() (*ActivityType, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *EstimateStakeFee) SetActivityType(v ActivityType)`

SetActivityType sets ActivityType field to given value.


### GetRequestId

`func (o *EstimateStakeFee) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EstimateStakeFee) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EstimateStakeFee) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *EstimateStakeFee) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetSource

`func (o *EstimateStakeFee) GetSource() StakingSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EstimateStakeFee) GetSourceOk() (*StakingSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EstimateStakeFee) SetSource(v StakingSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *EstimateStakeFee) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPoolId

`func (o *EstimateStakeFee) GetPoolId() StakingPoolId`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *EstimateStakeFee) GetPoolIdOk() (*StakingPoolId, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *EstimateStakeFee) SetPoolId(v StakingPoolId)`

SetPoolId sets PoolId field to given value.


### GetAmount

`func (o *EstimateStakeFee) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EstimateStakeFee) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EstimateStakeFee) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFee

`func (o *EstimateStakeFee) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *EstimateStakeFee) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *EstimateStakeFee) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.


### GetExtra

`func (o *EstimateStakeFee) GetExtra() CreateStakeActivityExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *EstimateStakeFee) GetExtraOk() (*CreateStakeActivityExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *EstimateStakeFee) SetExtra(v CreateStakeActivityExtra)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *EstimateStakeFee) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


