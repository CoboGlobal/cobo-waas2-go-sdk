# EigenLayerLstStakeExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolType** | [**StakingPoolType**](StakingPoolType.md) |  | 
**Operator** | Pointer to **string** | The operator address. | [optional] 

## Methods

### NewEigenLayerLstStakeExtra

`func NewEigenLayerLstStakeExtra(poolType StakingPoolType, ) *EigenLayerLstStakeExtra`

NewEigenLayerLstStakeExtra instantiates a new EigenLayerLstStakeExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEigenLayerLstStakeExtraWithDefaults

`func NewEigenLayerLstStakeExtraWithDefaults() *EigenLayerLstStakeExtra`

NewEigenLayerLstStakeExtraWithDefaults instantiates a new EigenLayerLstStakeExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolType

`func (o *EigenLayerLstStakeExtra) GetPoolType() StakingPoolType`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *EigenLayerLstStakeExtra) GetPoolTypeOk() (*StakingPoolType, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *EigenLayerLstStakeExtra) SetPoolType(v StakingPoolType)`

SetPoolType sets PoolType field to given value.


### GetOperator

`func (o *EigenLayerLstStakeExtra) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *EigenLayerLstStakeExtra) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *EigenLayerLstStakeExtra) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *EigenLayerLstStakeExtra) HasOperator() bool`

HasOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


