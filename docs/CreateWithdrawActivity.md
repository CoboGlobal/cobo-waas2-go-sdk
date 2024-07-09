# CreateWithdrawActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StakingId** | **string** | The id of the related staking. | 
**Amount** | Pointer to **string** | The amount to stake | [optional] 
**Fee** | [**TransactionTransferFee**](TransactionTransferFee.md) |  | 

## Methods

### NewCreateWithdrawActivity

`func NewCreateWithdrawActivity(stakingId string, fee TransactionTransferFee, ) *CreateWithdrawActivity`

NewCreateWithdrawActivity instantiates a new CreateWithdrawActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWithdrawActivityWithDefaults

`func NewCreateWithdrawActivityWithDefaults() *CreateWithdrawActivity`

NewCreateWithdrawActivityWithDefaults instantiates a new CreateWithdrawActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStakingId

`func (o *CreateWithdrawActivity) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *CreateWithdrawActivity) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *CreateWithdrawActivity) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.


### GetAmount

`func (o *CreateWithdrawActivity) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateWithdrawActivity) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateWithdrawActivity) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreateWithdrawActivity) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFee

`func (o *CreateWithdrawActivity) GetFee() TransactionTransferFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateWithdrawActivity) GetFeeOk() (*TransactionTransferFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateWithdrawActivity) SetFee(v TransactionTransferFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


