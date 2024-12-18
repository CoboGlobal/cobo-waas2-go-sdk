# StakingsExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**PosChain** | **string** | The Proof-of-Stake (PoS) chain. | 
**UnlockTimestamp** | Pointer to **int64** | The estimated time when the bitcoins will be unlocked, in Unix timestamp format, measured in milliseconds. | [optional] 
**UnlockBlockHeight** | Pointer to **int64** | The block height at which the bitcoins will be unlocked. | [optional] 
**StakeAddress** | Pointer to **string** | The address receiving the staked bitcoins. | [optional] 
**UnbondAddress** | Pointer to **string** | The address receiving the unlocked bitcoins. | [optional] 
**BeaconValidators** | Pointer to [**[]EthStakingExtraAllOfBeaconValidators**](EthStakingExtraAllOfBeaconValidators.md) | The list of validator information. | [optional] 
**StakerAddress** | **string** | The staker&#39;s Bitcoin address. | 
**ValidatorAddress** | **string** | The validator&#39;s EVM address. | 
**RewardAddress** | **string** | The EVM address to receive staking rewards. | 
**Timelock** | **int32** | The Unix timestamp (in seconds) when the staking position will be unlocked and available for withdrawal. | 

## Methods

### NewStakingsExtra

`func NewStakingsExtra(poolType StakingPoolType, posChain string, stakerAddress string, validatorAddress string, rewardAddress string, timelock int32, ) *StakingsExtra`

NewStakingsExtra instantiates a new StakingsExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStakingsExtraWithDefaults

`func NewStakingsExtraWithDefaults() *StakingsExtra`

NewStakingsExtraWithDefaults instantiates a new StakingsExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *StakingsExtra) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *StakingsExtra) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *StakingsExtra) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.


### GetPosChain

`func (o *StakingsExtra) GetPosChain() string`

GetPosChain returns the PosChain field if non-nil, zero value otherwise.

### GetPosChainOk

`func (o *StakingsExtra) GetPosChainOk() (*string, bool)`

GetPosChainOk returns a tuple with the PosChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosChain

`func (o *StakingsExtra) SetPosChain(v string)`

SetPosChain sets PosChain field to given value.


### GetUnlockTimestamp

`func (o *StakingsExtra) GetUnlockTimestamp() int64`

GetUnlockTimestamp returns the UnlockTimestamp field if non-nil, zero value otherwise.

### GetUnlockTimestampOk

`func (o *StakingsExtra) GetUnlockTimestampOk() (*int64, bool)`

GetUnlockTimestampOk returns a tuple with the UnlockTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockTimestamp

`func (o *StakingsExtra) SetUnlockTimestamp(v int64)`

SetUnlockTimestamp sets UnlockTimestamp field to given value.

### HasUnlockTimestamp

`func (o *StakingsExtra) HasUnlockTimestamp() bool`

HasUnlockTimestamp returns a boolean if a field has been set.

### GetUnlockBlockHeight

`func (o *StakingsExtra) GetUnlockBlockHeight() int64`

GetUnlockBlockHeight returns the UnlockBlockHeight field if non-nil, zero value otherwise.

### GetUnlockBlockHeightOk

`func (o *StakingsExtra) GetUnlockBlockHeightOk() (*int64, bool)`

GetUnlockBlockHeightOk returns a tuple with the UnlockBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockBlockHeight

`func (o *StakingsExtra) SetUnlockBlockHeight(v int64)`

SetUnlockBlockHeight sets UnlockBlockHeight field to given value.

### HasUnlockBlockHeight

`func (o *StakingsExtra) HasUnlockBlockHeight() bool`

HasUnlockBlockHeight returns a boolean if a field has been set.

### GetStakeAddress

`func (o *StakingsExtra) GetStakeAddress() string`

GetStakeAddress returns the StakeAddress field if non-nil, zero value otherwise.

### GetStakeAddressOk

`func (o *StakingsExtra) GetStakeAddressOk() (*string, bool)`

GetStakeAddressOk returns a tuple with the StakeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeAddress

`func (o *StakingsExtra) SetStakeAddress(v string)`

SetStakeAddress sets StakeAddress field to given value.

### HasStakeAddress

`func (o *StakingsExtra) HasStakeAddress() bool`

HasStakeAddress returns a boolean if a field has been set.

### GetUnbondAddress

`func (o *StakingsExtra) GetUnbondAddress() string`

GetUnbondAddress returns the UnbondAddress field if non-nil, zero value otherwise.

### GetUnbondAddressOk

`func (o *StakingsExtra) GetUnbondAddressOk() (*string, bool)`

GetUnbondAddressOk returns a tuple with the UnbondAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbondAddress

`func (o *StakingsExtra) SetUnbondAddress(v string)`

SetUnbondAddress sets UnbondAddress field to given value.

### HasUnbondAddress

`func (o *StakingsExtra) HasUnbondAddress() bool`

HasUnbondAddress returns a boolean if a field has been set.

### GetBeaconValidators

`func (o *StakingsExtra) GetBeaconValidators() []EthStakingExtraAllOfBeaconValidators`

GetBeaconValidators returns the BeaconValidators field if non-nil, zero value otherwise.

### GetBeaconValidatorsOk

`func (o *StakingsExtra) GetBeaconValidatorsOk() (*[]EthStakingExtraAllOfBeaconValidators, bool)`

GetBeaconValidatorsOk returns a tuple with the BeaconValidators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconValidators

`func (o *StakingsExtra) SetBeaconValidators(v []EthStakingExtraAllOfBeaconValidators)`

SetBeaconValidators sets BeaconValidators field to given value.

### HasBeaconValidators

`func (o *StakingsExtra) HasBeaconValidators() bool`

HasBeaconValidators returns a boolean if a field has been set.

### GetStakerAddress

`func (o *StakingsExtra) GetStakerAddress() string`

GetStakerAddress returns the StakerAddress field if non-nil, zero value otherwise.

### GetStakerAddressOk

`func (o *StakingsExtra) GetStakerAddressOk() (*string, bool)`

GetStakerAddressOk returns a tuple with the StakerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakerAddress

`func (o *StakingsExtra) SetStakerAddress(v string)`

SetStakerAddress sets StakerAddress field to given value.


### GetValidatorAddress

`func (o *StakingsExtra) GetValidatorAddress() string`

GetValidatorAddress returns the ValidatorAddress field if non-nil, zero value otherwise.

### GetValidatorAddressOk

`func (o *StakingsExtra) GetValidatorAddressOk() (*string, bool)`

GetValidatorAddressOk returns a tuple with the ValidatorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorAddress

`func (o *StakingsExtra) SetValidatorAddress(v string)`

SetValidatorAddress sets ValidatorAddress field to given value.


### GetRewardAddress

`func (o *StakingsExtra) GetRewardAddress() string`

GetRewardAddress returns the RewardAddress field if non-nil, zero value otherwise.

### GetRewardAddressOk

`func (o *StakingsExtra) GetRewardAddressOk() (*string, bool)`

GetRewardAddressOk returns a tuple with the RewardAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAddress

`func (o *StakingsExtra) SetRewardAddress(v string)`

SetRewardAddress sets RewardAddress field to given value.


### GetTimelock

`func (o *StakingsExtra) GetTimelock() int32`

GetTimelock returns the Timelock field if non-nil, zero value otherwise.

### GetTimelockOk

`func (o *StakingsExtra) GetTimelockOk() (*int32, bool)`

GetTimelockOk returns a tuple with the Timelock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelock

`func (o *StakingsExtra) SetTimelock(v int32)`

SetTimelock sets Timelock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


