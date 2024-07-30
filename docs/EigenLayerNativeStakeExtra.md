# EigenLayerNativeStakeExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**FeeRecipient** | Pointer to **float32** | The fee recipient address, if not provided the staker address will be used. | [optional] 

## Methods

### NewEigenLayerNativeStakeExtra

`func NewEigenLayerNativeStakeExtra(poolType StakingPoolType, ) *EigenLayerNativeStakeExtra`

NewEigenLayerNativeStakeExtra instantiates a new EigenLayerNativeStakeExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEigenLayerNativeStakeExtraWithDefaults

`func NewEigenLayerNativeStakeExtraWithDefaults() *EigenLayerNativeStakeExtra`

NewEigenLayerNativeStakeExtraWithDefaults instantiates a new EigenLayerNativeStakeExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *EigenLayerNativeStakeExtra) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *EigenLayerNativeStakeExtra) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *EigenLayerNativeStakeExtra) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.


### GetFeeRecipient

`func (o *EigenLayerNativeStakeExtra) GetFeeRecipient() float32`

GetFeeRecipient returns the FeeRecipient field if non-nil, zero value otherwise.

### GetFeeRecipientOk

`func (o *EigenLayerNativeStakeExtra) GetFeeRecipientOk() (*float32, bool)`

GetFeeRecipientOk returns a tuple with the FeeRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRecipient

`func (o *EigenLayerNativeStakeExtra) SetFeeRecipient(v float32)`

SetFeeRecipient sets FeeRecipient field to given value.

### HasFeeRecipient

`func (o *EigenLayerNativeStakeExtra) HasFeeRecipient() bool`

HasFeeRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


