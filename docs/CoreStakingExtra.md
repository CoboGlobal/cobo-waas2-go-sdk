# CoreStakingExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**PosChain** | **string** | The Proof-of-Stake (PoS) chain. | 
**StakerAddress** | **string** | The staker&#39;s Bitcoin address. | 
**ValidatorAddress** | **string** | The validator&#39;s EVM address. | 
**RewardAddress** | **string** | The EVM address to receive staking rewards. | 
**Timelock** | **int32** | The Unix timestamp (in seconds) when the staking position will be unlocked and available for withdrawal. | 

## Methods

### NewCoreStakingExtra

`func NewCoreStakingExtra(poolType StakingPoolType, posChain string, stakerAddress string, validatorAddress string, rewardAddress string, timelock int32, ) *CoreStakingExtra`

NewCoreStakingExtra instantiates a new CoreStakingExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreStakingExtraWithDefaults

`func NewCoreStakingExtraWithDefaults() *CoreStakingExtra`

NewCoreStakingExtraWithDefaults instantiates a new CoreStakingExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *CoreStakingExtra) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *CoreStakingExtra) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *CoreStakingExtra) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.


### GetPosChain

`func (o *CoreStakingExtra) GetPosChain() string`

GetPosChain returns the PosChain field if non-nil, zero value otherwise.

### GetPosChainOk

`func (o *CoreStakingExtra) GetPosChainOk() (*string, bool)`

GetPosChainOk returns a tuple with the PosChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosChain

`func (o *CoreStakingExtra) SetPosChain(v string)`

SetPosChain sets PosChain field to given value.


### GetStakerAddress

`func (o *CoreStakingExtra) GetStakerAddress() string`

GetStakerAddress returns the StakerAddress field if non-nil, zero value otherwise.

### GetStakerAddressOk

`func (o *CoreStakingExtra) GetStakerAddressOk() (*string, bool)`

GetStakerAddressOk returns a tuple with the StakerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakerAddress

`func (o *CoreStakingExtra) SetStakerAddress(v string)`

SetStakerAddress sets StakerAddress field to given value.


### GetValidatorAddress

`func (o *CoreStakingExtra) GetValidatorAddress() string`

GetValidatorAddress returns the ValidatorAddress field if non-nil, zero value otherwise.

### GetValidatorAddressOk

`func (o *CoreStakingExtra) GetValidatorAddressOk() (*string, bool)`

GetValidatorAddressOk returns a tuple with the ValidatorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorAddress

`func (o *CoreStakingExtra) SetValidatorAddress(v string)`

SetValidatorAddress sets ValidatorAddress field to given value.


### GetRewardAddress

`func (o *CoreStakingExtra) GetRewardAddress() string`

GetRewardAddress returns the RewardAddress field if non-nil, zero value otherwise.

### GetRewardAddressOk

`func (o *CoreStakingExtra) GetRewardAddressOk() (*string, bool)`

GetRewardAddressOk returns a tuple with the RewardAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAddress

`func (o *CoreStakingExtra) SetRewardAddress(v string)`

SetRewardAddress sets RewardAddress field to given value.


### GetTimelock

`func (o *CoreStakingExtra) GetTimelock() int32`

GetTimelock returns the Timelock field if non-nil, zero value otherwise.

### GetTimelockOk

`func (o *CoreStakingExtra) GetTimelockOk() (*int32, bool)`

GetTimelockOk returns a tuple with the Timelock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelock

`func (o *CoreStakingExtra) SetTimelock(v int32)`

SetTimelock sets Timelock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


