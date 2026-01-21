# PaymentBulkSendItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkSendItemId** | **string** | The payout item ID. | 
**TokenId** | **string** | The token ID of the payout item. | 
**ReceivingAddress** | **string** | The receiving address of the payout item. | 
**Amount** | **string** | The amount of the payout item. | 
**Description** | Pointer to **string** | The note of the payout item. | [optional] 
**Status** | [**PaymentBulkSendItemStatus**](PaymentBulkSendItemStatus.md) |  | 
**ValidationStatus** | [**PaymentBulkSendItemValidationStatus**](PaymentBulkSendItemValidationStatus.md) |  | 

## Methods

### NewPaymentBulkSendItem

`func NewPaymentBulkSendItem(bulkSendItemId string, tokenId string, receivingAddress string, amount string, status PaymentBulkSendItemStatus, validationStatus PaymentBulkSendItemValidationStatus, ) *PaymentBulkSendItem`

NewPaymentBulkSendItem instantiates a new PaymentBulkSendItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentBulkSendItemWithDefaults

`func NewPaymentBulkSendItemWithDefaults() *PaymentBulkSendItem`

NewPaymentBulkSendItemWithDefaults instantiates a new PaymentBulkSendItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkSendItemId

`func (o *PaymentBulkSendItem) GetBulkSendItemId() string`

GetBulkSendItemId returns the BulkSendItemId field if non-nil, zero value otherwise.

### GetBulkSendItemIdOk

`func (o *PaymentBulkSendItem) GetBulkSendItemIdOk() (*string, bool)`

GetBulkSendItemIdOk returns a tuple with the BulkSendItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkSendItemId

`func (o *PaymentBulkSendItem) SetBulkSendItemId(v string)`

SetBulkSendItemId sets BulkSendItemId field to given value.


### GetTokenId

`func (o *PaymentBulkSendItem) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentBulkSendItem) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentBulkSendItem) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetReceivingAddress

`func (o *PaymentBulkSendItem) GetReceivingAddress() string`

GetReceivingAddress returns the ReceivingAddress field if non-nil, zero value otherwise.

### GetReceivingAddressOk

`func (o *PaymentBulkSendItem) GetReceivingAddressOk() (*string, bool)`

GetReceivingAddressOk returns a tuple with the ReceivingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAddress

`func (o *PaymentBulkSendItem) SetReceivingAddress(v string)`

SetReceivingAddress sets ReceivingAddress field to given value.


### GetAmount

`func (o *PaymentBulkSendItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentBulkSendItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentBulkSendItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *PaymentBulkSendItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentBulkSendItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentBulkSendItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentBulkSendItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentBulkSendItem) GetStatus() PaymentBulkSendItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentBulkSendItem) GetStatusOk() (*PaymentBulkSendItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentBulkSendItem) SetStatus(v PaymentBulkSendItemStatus)`

SetStatus sets Status field to given value.


### GetValidationStatus

`func (o *PaymentBulkSendItem) GetValidationStatus() PaymentBulkSendItemValidationStatus`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *PaymentBulkSendItem) GetValidationStatusOk() (*PaymentBulkSendItemValidationStatus, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *PaymentBulkSendItem) SetValidationStatus(v PaymentBulkSendItemValidationStatus)`

SetValidationStatus sets ValidationStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


