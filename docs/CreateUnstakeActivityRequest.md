# CreateUnstakeActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StakingId** | **string** | The id of the related staking. | 
**Amount** | Pointer to **string** | The amount to stake | [optional] 
**Fee** | [**TransactionTransferFee**](TransactionTransferFee.md) |  | 
**Initiator** | Pointer to **string** | The initiator of the staking activity. | [optional] 

## Methods

### NewCreateUnstakeActivityRequest

`func NewCreateUnstakeActivityRequest(stakingId string, fee TransactionTransferFee, ) *CreateUnstakeActivityRequest`

NewCreateUnstakeActivityRequest instantiates a new CreateUnstakeActivityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUnstakeActivityRequestWithDefaults

`func NewCreateUnstakeActivityRequestWithDefaults() *CreateUnstakeActivityRequest`

NewCreateUnstakeActivityRequestWithDefaults instantiates a new CreateUnstakeActivityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStakingId

`func (o *CreateUnstakeActivityRequest) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *CreateUnstakeActivityRequest) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *CreateUnstakeActivityRequest) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.


### GetAmount

`func (o *CreateUnstakeActivityRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateUnstakeActivityRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateUnstakeActivityRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreateUnstakeActivityRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFee

`func (o *CreateUnstakeActivityRequest) GetFee() TransactionTransferFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateUnstakeActivityRequest) GetFeeOk() (*TransactionTransferFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateUnstakeActivityRequest) SetFee(v TransactionTransferFee)`

SetFee sets Fee field to given value.


### GetInitiator

`func (o *CreateUnstakeActivityRequest) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *CreateUnstakeActivityRequest) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *CreateUnstakeActivityRequest) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *CreateUnstakeActivityRequest) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


