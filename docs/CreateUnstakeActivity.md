# CreateUnstakeActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization. | [optional] 
**StakingId** | **string** | The ID of the corresponding staking position. | 
**Amount** | Pointer to **string** | The amount to unstake. For the Babylon protocol, this property is ignored. | [optional] 
**Fee** | Pointer to [**TransactionRequestFee**](TransactionRequestFee.md) |  | [optional] 

## Methods

### NewCreateUnstakeActivity

`func NewCreateUnstakeActivity(stakingId string, ) *CreateUnstakeActivity`

NewCreateUnstakeActivity instantiates a new CreateUnstakeActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUnstakeActivityWithDefaults

`func NewCreateUnstakeActivityWithDefaults() *CreateUnstakeActivity`

NewCreateUnstakeActivityWithDefaults instantiates a new CreateUnstakeActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateUnstakeActivity) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateUnstakeActivity) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateUnstakeActivity) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateUnstakeActivity) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStakingId

`func (o *CreateUnstakeActivity) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *CreateUnstakeActivity) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *CreateUnstakeActivity) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.


### GetAmount

`func (o *CreateUnstakeActivity) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateUnstakeActivity) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateUnstakeActivity) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreateUnstakeActivity) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFee

`func (o *CreateUnstakeActivity) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateUnstakeActivity) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateUnstakeActivity) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *CreateUnstakeActivity) HasFee() bool`

HasFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


