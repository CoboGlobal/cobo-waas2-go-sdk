# CreateWithdrawActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization. | [optional] 
**StakingId** | **string** | The ID of the corresponding staking position. | 
**Amount** | Pointer to **string** | The amount to withdraw. | [optional] 
**Address** | Pointer to **string** | The withdrawal address. | [optional] 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 
**AppInitiator** | Pointer to **string** | The initiator of the staking activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator. | [optional] 

## Methods

### NewCreateWithdrawActivityRequest

`func NewCreateWithdrawActivityRequest(stakingId string, fee TransactionRequestFee, ) *CreateWithdrawActivityRequest`

NewCreateWithdrawActivityRequest instantiates a new CreateWithdrawActivityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWithdrawActivityRequestWithDefaults

`func NewCreateWithdrawActivityRequestWithDefaults() *CreateWithdrawActivityRequest`

NewCreateWithdrawActivityRequestWithDefaults instantiates a new CreateWithdrawActivityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateWithdrawActivityRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateWithdrawActivityRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateWithdrawActivityRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateWithdrawActivityRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStakingId

`func (o *CreateWithdrawActivityRequest) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *CreateWithdrawActivityRequest) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *CreateWithdrawActivityRequest) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.


### GetAmount

`func (o *CreateWithdrawActivityRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateWithdrawActivityRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateWithdrawActivityRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreateWithdrawActivityRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAddress

`func (o *CreateWithdrawActivityRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateWithdrawActivityRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateWithdrawActivityRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CreateWithdrawActivityRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetFee

`func (o *CreateWithdrawActivityRequest) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateWithdrawActivityRequest) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateWithdrawActivityRequest) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.


### GetAppInitiator

`func (o *CreateWithdrawActivityRequest) GetAppInitiator() string`

GetAppInitiator returns the AppInitiator field if non-nil, zero value otherwise.

### GetAppInitiatorOk

`func (o *CreateWithdrawActivityRequest) GetAppInitiatorOk() (*string, bool)`

GetAppInitiatorOk returns a tuple with the AppInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInitiator

`func (o *CreateWithdrawActivityRequest) SetAppInitiator(v string)`

SetAppInitiator sets AppInitiator field to given value.

### HasAppInitiator

`func (o *CreateWithdrawActivityRequest) HasAppInitiator() bool`

HasAppInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


