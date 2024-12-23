# CoreStakingActivityDetailExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**Timelock** | Pointer to **int32** | The Unix timestamp (in seconds) when the staking position will be unlocked and available for withdrawal. | [optional] 
**ChangeAddress** | Pointer to **string** | The change bitcoin address. If not provided, the source wallet&#39;s address will be used as the change address. | [optional] 
**ValidatorAddress** | Pointer to **string** | The validator evm address. | [optional] 
**RewardAddress** | Pointer to **string** | The reward evm address. | [optional] 

## Methods

### NewCoreStakingActivityDetailExtra

`func NewCoreStakingActivityDetailExtra(poolType StakingPoolType, ) *CoreStakingActivityDetailExtra`

NewCoreStakingActivityDetailExtra instantiates a new CoreStakingActivityDetailExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreStakingActivityDetailExtraWithDefaults

`func NewCoreStakingActivityDetailExtraWithDefaults() *CoreStakingActivityDetailExtra`

NewCoreStakingActivityDetailExtraWithDefaults instantiates a new CoreStakingActivityDetailExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *CoreStakingActivityDetailExtra) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *CoreStakingActivityDetailExtra) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *CoreStakingActivityDetailExtra) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.


### GetTimelock

`func (o *CoreStakingActivityDetailExtra) GetTimelock() int32`

GetTimelock returns the Timelock field if non-nil, zero value otherwise.

### GetTimelockOk

`func (o *CoreStakingActivityDetailExtra) GetTimelockOk() (*int32, bool)`

GetTimelockOk returns a tuple with the Timelock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelock

`func (o *CoreStakingActivityDetailExtra) SetTimelock(v int32)`

SetTimelock sets Timelock field to given value.

### HasTimelock

`func (o *CoreStakingActivityDetailExtra) HasTimelock() bool`

HasTimelock returns a boolean if a field has been set.

### GetChangeAddress

`func (o *CoreStakingActivityDetailExtra) GetChangeAddress() string`

GetChangeAddress returns the ChangeAddress field if non-nil, zero value otherwise.

### GetChangeAddressOk

`func (o *CoreStakingActivityDetailExtra) GetChangeAddressOk() (*string, bool)`

GetChangeAddressOk returns a tuple with the ChangeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAddress

`func (o *CoreStakingActivityDetailExtra) SetChangeAddress(v string)`

SetChangeAddress sets ChangeAddress field to given value.

### HasChangeAddress

`func (o *CoreStakingActivityDetailExtra) HasChangeAddress() bool`

HasChangeAddress returns a boolean if a field has been set.

### GetValidatorAddress

`func (o *CoreStakingActivityDetailExtra) GetValidatorAddress() string`

GetValidatorAddress returns the ValidatorAddress field if non-nil, zero value otherwise.

### GetValidatorAddressOk

`func (o *CoreStakingActivityDetailExtra) GetValidatorAddressOk() (*string, bool)`

GetValidatorAddressOk returns a tuple with the ValidatorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorAddress

`func (o *CoreStakingActivityDetailExtra) SetValidatorAddress(v string)`

SetValidatorAddress sets ValidatorAddress field to given value.

### HasValidatorAddress

`func (o *CoreStakingActivityDetailExtra) HasValidatorAddress() bool`

HasValidatorAddress returns a boolean if a field has been set.

### GetRewardAddress

`func (o *CoreStakingActivityDetailExtra) GetRewardAddress() string`

GetRewardAddress returns the RewardAddress field if non-nil, zero value otherwise.

### GetRewardAddressOk

`func (o *CoreStakingActivityDetailExtra) GetRewardAddressOk() (*string, bool)`

GetRewardAddressOk returns a tuple with the RewardAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAddress

`func (o *CoreStakingActivityDetailExtra) SetRewardAddress(v string)`

SetRewardAddress sets RewardAddress field to given value.

### HasRewardAddress

`func (o *CoreStakingActivityDetailExtra) HasRewardAddress() bool`

HasRewardAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


