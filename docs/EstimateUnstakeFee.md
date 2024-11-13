# EstimateUnstakeFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityType** | [**ActivityType**](ActivityType.md) |  | 
**RequestId** | Pointer to **string** | The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization. | [optional] 
**StakingId** | **string** | The ID of the corresponding staking position. | 
**Amount** | Pointer to **string** | The amount to unstake. For the Babylon protocol, this property is ignored. | [optional] 
**Fee** | Pointer to [**TransactionRequestFee**](TransactionRequestFee.md) |  | [optional] 
**Extra** | Pointer to [**CreateUnstakeActivityExtra**](CreateUnstakeActivityExtra.md) |  | [optional] 

## Methods

### NewEstimateUnstakeFee

`func NewEstimateUnstakeFee(activityType ActivityType, stakingId string, ) *EstimateUnstakeFee`

NewEstimateUnstakeFee instantiates a new EstimateUnstakeFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateUnstakeFeeWithDefaults

`func NewEstimateUnstakeFeeWithDefaults() *EstimateUnstakeFee`

NewEstimateUnstakeFeeWithDefaults instantiates a new EstimateUnstakeFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityType

`func (o *EstimateUnstakeFee) GetActivityType() ActivityType`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *EstimateUnstakeFee) GetActivityTypeOk() (*ActivityType, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *EstimateUnstakeFee) SetActivityType(v ActivityType)`

SetActivityType sets ActivityType field to given value.


### GetRequestId

`func (o *EstimateUnstakeFee) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EstimateUnstakeFee) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EstimateUnstakeFee) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *EstimateUnstakeFee) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStakingId

`func (o *EstimateUnstakeFee) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *EstimateUnstakeFee) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *EstimateUnstakeFee) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.


### GetAmount

`func (o *EstimateUnstakeFee) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EstimateUnstakeFee) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EstimateUnstakeFee) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EstimateUnstakeFee) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFee

`func (o *EstimateUnstakeFee) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *EstimateUnstakeFee) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *EstimateUnstakeFee) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *EstimateUnstakeFee) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetExtra

`func (o *EstimateUnstakeFee) GetExtra() CreateUnstakeActivityExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *EstimateUnstakeFee) GetExtraOk() (*CreateUnstakeActivityExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *EstimateUnstakeFee) SetExtra(v CreateUnstakeActivityExtra)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *EstimateUnstakeFee) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


