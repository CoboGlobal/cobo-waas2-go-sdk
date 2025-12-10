# PayoutDestinationDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The destination address. | 
**Amount** | **string** | The amount of cryptocurrency to send to the destination. | 
**Email** | Pointer to **string** | The email of the destination. | [optional] 
**Note** | Pointer to **string** | The note of the destination. | [optional] 
**LoopEnabled** | Pointer to **bool** | Enable loop mode for standard transfers when source is asset wallet. Only applicable when: - &#x60;payout_mode&#x60; is &#x60;Normal&#x60; - &#x60;source_type&#x60; is asset wallet  | [optional] [default to false]
**NetworkFee** | Pointer to [**PayoutFeeData**](PayoutFeeData.md) |  | [optional] 
**Status** | Pointer to [**PayoutTransactionStatus**](PayoutTransactionStatus.md) |  | [optional] 
**TransactionId** | Pointer to **string** | The transaction ID of the destination. | [optional] 

## Methods

### NewPayoutDestinationDetail

`func NewPayoutDestinationDetail(address string, amount string, ) *PayoutDestinationDetail`

NewPayoutDestinationDetail instantiates a new PayoutDestinationDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutDestinationDetailWithDefaults

`func NewPayoutDestinationDetailWithDefaults() *PayoutDestinationDetail`

NewPayoutDestinationDetailWithDefaults instantiates a new PayoutDestinationDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PayoutDestinationDetail) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PayoutDestinationDetail) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PayoutDestinationDetail) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *PayoutDestinationDetail) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayoutDestinationDetail) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayoutDestinationDetail) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetEmail

`func (o *PayoutDestinationDetail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PayoutDestinationDetail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PayoutDestinationDetail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PayoutDestinationDetail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNote

`func (o *PayoutDestinationDetail) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PayoutDestinationDetail) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PayoutDestinationDetail) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *PayoutDestinationDetail) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetLoopEnabled

`func (o *PayoutDestinationDetail) GetLoopEnabled() bool`

GetLoopEnabled returns the LoopEnabled field if non-nil, zero value otherwise.

### GetLoopEnabledOk

`func (o *PayoutDestinationDetail) GetLoopEnabledOk() (*bool, bool)`

GetLoopEnabledOk returns a tuple with the LoopEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopEnabled

`func (o *PayoutDestinationDetail) SetLoopEnabled(v bool)`

SetLoopEnabled sets LoopEnabled field to given value.

### HasLoopEnabled

`func (o *PayoutDestinationDetail) HasLoopEnabled() bool`

HasLoopEnabled returns a boolean if a field has been set.

### GetNetworkFee

`func (o *PayoutDestinationDetail) GetNetworkFee() PayoutFeeData`

GetNetworkFee returns the NetworkFee field if non-nil, zero value otherwise.

### GetNetworkFeeOk

`func (o *PayoutDestinationDetail) GetNetworkFeeOk() (*PayoutFeeData, bool)`

GetNetworkFeeOk returns a tuple with the NetworkFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFee

`func (o *PayoutDestinationDetail) SetNetworkFee(v PayoutFeeData)`

SetNetworkFee sets NetworkFee field to given value.

### HasNetworkFee

`func (o *PayoutDestinationDetail) HasNetworkFee() bool`

HasNetworkFee returns a boolean if a field has been set.

### GetStatus

`func (o *PayoutDestinationDetail) GetStatus() PayoutTransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PayoutDestinationDetail) GetStatusOk() (*PayoutTransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PayoutDestinationDetail) SetStatus(v PayoutTransactionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PayoutDestinationDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransactionId

`func (o *PayoutDestinationDetail) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *PayoutDestinationDetail) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *PayoutDestinationDetail) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *PayoutDestinationDetail) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


