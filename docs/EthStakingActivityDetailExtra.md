# EthStakingActivityDetailExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**ProviderName** | Pointer to **string** | The name of the provider. | [optional] 
**ValidatorPubkeys** | Pointer to **[]string** | A list of public keys associated with the Ethereum validators for this unstaking operation. | [optional] 

## Methods

### NewEthStakingActivityDetailExtra

`func NewEthStakingActivityDetailExtra(poolType StakingPoolType, ) *EthStakingActivityDetailExtra`

NewEthStakingActivityDetailExtra instantiates a new EthStakingActivityDetailExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthStakingActivityDetailExtraWithDefaults

`func NewEthStakingActivityDetailExtraWithDefaults() *EthStakingActivityDetailExtra`

NewEthStakingActivityDetailExtraWithDefaults instantiates a new EthStakingActivityDetailExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *EthStakingActivityDetailExtra) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *EthStakingActivityDetailExtra) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *EthStakingActivityDetailExtra) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.


### GetProviderName

`func (o *EthStakingActivityDetailExtra) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *EthStakingActivityDetailExtra) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *EthStakingActivityDetailExtra) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *EthStakingActivityDetailExtra) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetValidatorPubkeys

`func (o *EthStakingActivityDetailExtra) GetValidatorPubkeys() []string`

GetValidatorPubkeys returns the ValidatorPubkeys field if non-nil, zero value otherwise.

### GetValidatorPubkeysOk

`func (o *EthStakingActivityDetailExtra) GetValidatorPubkeysOk() (*[]string, bool)`

GetValidatorPubkeysOk returns a tuple with the ValidatorPubkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorPubkeys

`func (o *EthStakingActivityDetailExtra) SetValidatorPubkeys(v []string)`

SetValidatorPubkeys sets ValidatorPubkeys field to given value.

### HasValidatorPubkeys

`func (o *EthStakingActivityDetailExtra) HasValidatorPubkeys() bool`

HasValidatorPubkeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


