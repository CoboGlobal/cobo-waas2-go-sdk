# EthUnstakeExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**ValidatorPubkeys** | **[]string** | A list of public keys identifying the validators to unstake from the Ethereum Beacon Chain. | 

## Methods

### NewEthUnstakeExtra

`func NewEthUnstakeExtra(poolType StakingPoolType, validatorPubkeys []string, ) *EthUnstakeExtra`

NewEthUnstakeExtra instantiates a new EthUnstakeExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthUnstakeExtraWithDefaults

`func NewEthUnstakeExtraWithDefaults() *EthUnstakeExtra`

NewEthUnstakeExtraWithDefaults instantiates a new EthUnstakeExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *EthUnstakeExtra) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *EthUnstakeExtra) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *EthUnstakeExtra) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.


### GetValidatorPubkeys

`func (o *EthUnstakeExtra) GetValidatorPubkeys() []string`

GetValidatorPubkeys returns the ValidatorPubkeys field if non-nil, zero value otherwise.

### GetValidatorPubkeysOk

`func (o *EthUnstakeExtra) GetValidatorPubkeysOk() (*[]string, bool)`

GetValidatorPubkeysOk returns a tuple with the ValidatorPubkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorPubkeys

`func (o *EthUnstakeExtra) SetValidatorPubkeys(v []string)`

SetValidatorPubkeys sets ValidatorPubkeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


