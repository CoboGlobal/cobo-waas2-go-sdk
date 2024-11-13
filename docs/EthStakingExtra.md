# EthStakingExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**PosChain** | **string** | The Proof-of-Stake (PoS) chain. | 

## Methods

### NewEthStakingExtra

`func NewEthStakingExtra(poolType StakingPoolType, posChain string, ) *EthStakingExtra`

NewEthStakingExtra instantiates a new EthStakingExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthStakingExtraWithDefaults

`func NewEthStakingExtraWithDefaults() *EthStakingExtra`

NewEthStakingExtraWithDefaults instantiates a new EthStakingExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *EthStakingExtra) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *EthStakingExtra) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *EthStakingExtra) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.


### GetPosChain

`func (o *EthStakingExtra) GetPosChain() string`

GetPosChain returns the PosChain field if non-nil, zero value otherwise.

### GetPosChainOk

`func (o *EthStakingExtra) GetPosChainOk() (*string, bool)`

GetPosChainOk returns a tuple with the PosChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosChain

`func (o *EthStakingExtra) SetPosChain(v string)`

SetPosChain sets PosChain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


