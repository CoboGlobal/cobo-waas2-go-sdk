# CreateWithdrawActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization. | [optional] 
**StakingId** | **string** | The ID of the corresponding staking position. | 
**Amount** | Pointer to **string** | The amount to withdraw. | [optional] 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 

## Methods

### NewCreateWithdrawActivity

`func NewCreateWithdrawActivity(stakingId string, fee TransactionRequestFee, ) *CreateWithdrawActivity`

NewCreateWithdrawActivity instantiates a new CreateWithdrawActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWithdrawActivityWithDefaults

`func NewCreateWithdrawActivityWithDefaults() *CreateWithdrawActivity`

NewCreateWithdrawActivityWithDefaults instantiates a new CreateWithdrawActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateWithdrawActivity) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateWithdrawActivity) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateWithdrawActivity) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateWithdrawActivity) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

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

`func (o *CreateWithdrawActivity) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateWithdrawActivity) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateWithdrawActivity) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


