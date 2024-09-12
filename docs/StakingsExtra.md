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

## Methods

### NewStakingsExtra

`func NewStakingsExtra(poolType StakingPoolType, posChain string, ) *StakingsExtra`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


