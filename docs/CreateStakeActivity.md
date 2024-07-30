# CreateStakeActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**StakingSource**](StakingSource.md) |  | [optional] 
**PoolId** | **string** | The id of the staking pool | 
**Amount** | **string** | The amount to stake | 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 
**Extra** | [**CreateStakeActivityExtra**](CreateStakeActivityExtra.md) |  | 

## Methods

### NewCreateStakeActivity

`func NewCreateStakeActivity(poolId string, amount string, fee TransactionRequestFee, extra CreateStakeActivityExtra, ) *CreateStakeActivity`

NewCreateStakeActivity instantiates a new CreateStakeActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStakeActivityWithDefaults

`func NewCreateStakeActivityWithDefaults() *CreateStakeActivity`

NewCreateStakeActivityWithDefaults instantiates a new CreateStakeActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *CreateStakeActivity) GetSource() StakingSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateStakeActivity) GetSourceOk() (*StakingSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateStakeActivity) SetSource(v StakingSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *CreateStakeActivity) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPoolId

`func (o *CreateStakeActivity) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *CreateStakeActivity) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *CreateStakeActivity) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetAmount

`func (o *CreateStakeActivity) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateStakeActivity) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateStakeActivity) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFee

`func (o *CreateStakeActivity) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateStakeActivity) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateStakeActivity) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.


### GetExtra

`func (o *CreateStakeActivity) GetExtra() CreateStakeActivityExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *CreateStakeActivity) GetExtraOk() (*CreateStakeActivityExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *CreateStakeActivity) SetExtra(v CreateStakeActivityExtra)`

SetExtra sets Extra field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


