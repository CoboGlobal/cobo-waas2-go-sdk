# EthStakeEstimatedFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | Pointer to [**StakingPoolType**](StakingPoolType.md) |  | [optional] 
**Fee** | Pointer to [**EstimatedFee**](EstimatedFee.md) |  | [optional] 
**ValidatorPubkeys** | Pointer to **[]string** | A list of public keys associated with the Ethereum validators for this staking operation. | [optional] 

## Methods

### NewEthStakeEstimatedFee

`func NewEthStakeEstimatedFee() *EthStakeEstimatedFee`

NewEthStakeEstimatedFee instantiates a new EthStakeEstimatedFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthStakeEstimatedFeeWithDefaults

`func NewEthStakeEstimatedFeeWithDefaults() *EthStakeEstimatedFee`

NewEthStakeEstimatedFeeWithDefaults instantiates a new EthStakeEstimatedFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *EthStakeEstimatedFee) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *EthStakeEstimatedFee) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *EthStakeEstimatedFee) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *EthStakeEstimatedFee) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetFee

`func (o *EthStakeEstimatedFee) GetFee() EstimatedFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *EthStakeEstimatedFee) GetFeeOk() (*EstimatedFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *EthStakeEstimatedFee) SetFee(v EstimatedFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *EthStakeEstimatedFee) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetValidatorPubkeys

`func (o *EthStakeEstimatedFee) GetValidatorPubkeys() []string`

GetValidatorPubkeys returns the ValidatorPubkeys field if non-nil, zero value otherwise.

### GetValidatorPubkeysOk

`func (o *EthStakeEstimatedFee) GetValidatorPubkeysOk() (*[]string, bool)`

GetValidatorPubkeysOk returns a tuple with the ValidatorPubkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorPubkeys

`func (o *EthStakeEstimatedFee) SetValidatorPubkeys(v []string)`

SetValidatorPubkeys sets ValidatorPubkeys field to given value.

### HasValidatorPubkeys

`func (o *EthStakeEstimatedFee) HasValidatorPubkeys() bool`

HasValidatorPubkeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


