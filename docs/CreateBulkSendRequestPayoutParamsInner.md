# CreateBulkSendRequestPayoutParamsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID of the cryptocurrency to be sent to the recipient. | 
**ReceivingAddress** | **string** | The receiving address. | 
**Amount** | **string** | The amount of the cryptocurrency to be sent to the recipient. | 
**Description** | Pointer to **string** | A note or comment about the bulk send item. | [optional] 

## Methods

### NewCreateBulkSendRequestPayoutParamsInner

`func NewCreateBulkSendRequestPayoutParamsInner(tokenId string, receivingAddress string, amount string, ) *CreateBulkSendRequestPayoutParamsInner`

NewCreateBulkSendRequestPayoutParamsInner instantiates a new CreateBulkSendRequestPayoutParamsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBulkSendRequestPayoutParamsInnerWithDefaults

`func NewCreateBulkSendRequestPayoutParamsInnerWithDefaults() *CreateBulkSendRequestPayoutParamsInner`

NewCreateBulkSendRequestPayoutParamsInnerWithDefaults instantiates a new CreateBulkSendRequestPayoutParamsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *CreateBulkSendRequestPayoutParamsInner) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CreateBulkSendRequestPayoutParamsInner) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CreateBulkSendRequestPayoutParamsInner) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetReceivingAddress

`func (o *CreateBulkSendRequestPayoutParamsInner) GetReceivingAddress() string`

GetReceivingAddress returns the ReceivingAddress field if non-nil, zero value otherwise.

### GetReceivingAddressOk

`func (o *CreateBulkSendRequestPayoutParamsInner) GetReceivingAddressOk() (*string, bool)`

GetReceivingAddressOk returns a tuple with the ReceivingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAddress

`func (o *CreateBulkSendRequestPayoutParamsInner) SetReceivingAddress(v string)`

SetReceivingAddress sets ReceivingAddress field to given value.


### GetAmount

`func (o *CreateBulkSendRequestPayoutParamsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateBulkSendRequestPayoutParamsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateBulkSendRequestPayoutParamsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *CreateBulkSendRequestPayoutParamsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateBulkSendRequestPayoutParamsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateBulkSendRequestPayoutParamsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateBulkSendRequestPayoutParamsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


