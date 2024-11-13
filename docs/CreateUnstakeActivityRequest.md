# CreateUnstakeActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization. | [optional] 
**StakingId** | **string** | The ID of the corresponding staking position. | 
**Amount** | Pointer to **string** | The amount to unstake. For the Babylon protocol, this property is ignored. | [optional] 
**Fee** | Pointer to [**TransactionRequestFee**](TransactionRequestFee.md) |  | [optional] 
**Extra** | Pointer to [**CreateUnstakeActivityExtra**](CreateUnstakeActivityExtra.md) |  | [optional] 
**AppInitiator** | Pointer to **string** | The initiator of the staking activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator. | [optional] 

## Methods

### NewCreateUnstakeActivityRequest

`func NewCreateUnstakeActivityRequest(stakingId string, ) *CreateUnstakeActivityRequest`

NewCreateUnstakeActivityRequest instantiates a new CreateUnstakeActivityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUnstakeActivityRequestWithDefaults

`func NewCreateUnstakeActivityRequestWithDefaults() *CreateUnstakeActivityRequest`

NewCreateUnstakeActivityRequestWithDefaults instantiates a new CreateUnstakeActivityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateUnstakeActivityRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateUnstakeActivityRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateUnstakeActivityRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateUnstakeActivityRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

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

`func (o *CreateUnstakeActivityRequest) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateUnstakeActivityRequest) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateUnstakeActivityRequest) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *CreateUnstakeActivityRequest) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetExtra

`func (o *CreateUnstakeActivityRequest) GetExtra() CreateUnstakeActivityExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *CreateUnstakeActivityRequest) GetExtraOk() (*CreateUnstakeActivityExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *CreateUnstakeActivityRequest) SetExtra(v CreateUnstakeActivityExtra)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *CreateUnstakeActivityRequest) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetAppInitiator

`func (o *CreateUnstakeActivityRequest) GetAppInitiator() string`

GetAppInitiator returns the AppInitiator field if non-nil, zero value otherwise.

### GetAppInitiatorOk

`func (o *CreateUnstakeActivityRequest) GetAppInitiatorOk() (*string, bool)`

GetAppInitiatorOk returns a tuple with the AppInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInitiator

`func (o *CreateUnstakeActivityRequest) SetAppInitiator(v string)`

SetAppInitiator sets AppInitiator field to given value.

### HasAppInitiator

`func (o *CreateUnstakeActivityRequest) HasAppInitiator() bool`

HasAppInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


