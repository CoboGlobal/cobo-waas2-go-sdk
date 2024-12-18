# CoreStakeExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**Timelock** | **int32** | The Unix timestamp (in seconds) when the staking position will be unlocked and available for withdrawal. | 
**ChangeAddress** | Pointer to **string** | The change address on the Bitcoin chain. If not provided, the source wallet&#39;s address will be used as the change address. | [optional] 
**ValidatorAddress** | **string** | The validator&#39;s EVM address. | 
**RewardAddress** | **string** | The EVM address to receive staking rewards. | 

## Methods

### NewCoreStakeExtra

`func NewCoreStakeExtra(poolType StakingPoolType, timelock int32, validatorAddress string, rewardAddress string, ) *CoreStakeExtra`

NewCoreStakeExtra instantiates a new CoreStakeExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreStakeExtraWithDefaults

`func NewCoreStakeExtraWithDefaults() *CoreStakeExtra`

NewCoreStakeExtraWithDefaults instantiates a new CoreStakeExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *CoreStakeExtra) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *CoreStakeExtra) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *CoreStakeExtra) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.


### GetTimelock

`func (o *CoreStakeExtra) GetTimelock() int32`

GetTimelock returns the Timelock field if non-nil, zero value otherwise.

### GetTimelockOk

`func (o *CoreStakeExtra) GetTimelockOk() (*int32, bool)`

GetTimelockOk returns a tuple with the Timelock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelock

`func (o *CoreStakeExtra) SetTimelock(v int32)`

SetTimelock sets Timelock field to given value.


### GetChangeAddress

`func (o *CoreStakeExtra) GetChangeAddress() string`

GetChangeAddress returns the ChangeAddress field if non-nil, zero value otherwise.

### GetChangeAddressOk

`func (o *CoreStakeExtra) GetChangeAddressOk() (*string, bool)`

GetChangeAddressOk returns a tuple with the ChangeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAddress

`func (o *CoreStakeExtra) SetChangeAddress(v string)`

SetChangeAddress sets ChangeAddress field to given value.

### HasChangeAddress

`func (o *CoreStakeExtra) HasChangeAddress() bool`

HasChangeAddress returns a boolean if a field has been set.

### GetValidatorAddress

`func (o *CoreStakeExtra) GetValidatorAddress() string`

GetValidatorAddress returns the ValidatorAddress field if non-nil, zero value otherwise.

### GetValidatorAddressOk

`func (o *CoreStakeExtra) GetValidatorAddressOk() (*string, bool)`

GetValidatorAddressOk returns a tuple with the ValidatorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorAddress

`func (o *CoreStakeExtra) SetValidatorAddress(v string)`

SetValidatorAddress sets ValidatorAddress field to given value.


### GetRewardAddress

`func (o *CoreStakeExtra) GetRewardAddress() string`

GetRewardAddress returns the RewardAddress field if non-nil, zero value otherwise.

### GetRewardAddressOk

`func (o *CoreStakeExtra) GetRewardAddressOk() (*string, bool)`

GetRewardAddressOk returns a tuple with the RewardAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAddress

`func (o *CoreStakeExtra) SetRewardAddress(v string)`

SetRewardAddress sets RewardAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


