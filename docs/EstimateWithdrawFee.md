# EstimateWithdrawFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityType** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 
**StakingId** | **string** | The id of the related staking. | 
**Amount** | Pointer to **string** | The amount to stake | [optional] 
**Fee** | [**TransactionTransferFee**](TransactionTransferFee.md) |  | 

## Methods

### NewEstimateWithdrawFee

`func NewEstimateWithdrawFee(stakingId string, fee TransactionTransferFee, ) *EstimateWithdrawFee`

NewEstimateWithdrawFee instantiates a new EstimateWithdrawFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateWithdrawFeeWithDefaults

`func NewEstimateWithdrawFeeWithDefaults() *EstimateWithdrawFee`

NewEstimateWithdrawFeeWithDefaults instantiates a new EstimateWithdrawFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityType

`func (o *EstimateWithdrawFee) GetActivityType() ActivityType`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *EstimateWithdrawFee) GetActivityTypeOk() (*ActivityType, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *EstimateWithdrawFee) SetActivityType(v ActivityType)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *EstimateWithdrawFee) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetStakingId

`func (o *EstimateWithdrawFee) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *EstimateWithdrawFee) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *EstimateWithdrawFee) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.


### GetAmount

`func (o *EstimateWithdrawFee) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EstimateWithdrawFee) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EstimateWithdrawFee) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EstimateWithdrawFee) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFee

`func (o *EstimateWithdrawFee) GetFee() TransactionTransferFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *EstimateWithdrawFee) GetFeeOk() (*TransactionTransferFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *EstimateWithdrawFee) SetFee(v TransactionTransferFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


