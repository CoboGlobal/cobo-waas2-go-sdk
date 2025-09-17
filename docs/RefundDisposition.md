# RefundDisposition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The UUID of the transaction whose funds are to be refunded. This identifies the original transaction that requires refund processing. | 
**DestinationAddress** | **string** | The blockchain address to receive the refunded funds. | 
**DispositionAmount** | **string** | The amount to be refunded from the original transaction, specified as a numeric string. This value cannot exceed the total amount of the original transaction.  | 
**CategoryNames** | Pointer to **[]string** | Custom categories to identify and track this refund transaction. Used for transaction classification and reporting. | [optional] 
**Description** | Pointer to **string** | Additional notes or description for the refund. | [optional] 

## Methods

### NewRefundDisposition

`func NewRefundDisposition(transactionId string, destinationAddress string, dispositionAmount string, ) *RefundDisposition`

NewRefundDisposition instantiates a new RefundDisposition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundDispositionWithDefaults

`func NewRefundDispositionWithDefaults() *RefundDisposition`

NewRefundDispositionWithDefaults instantiates a new RefundDisposition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *RefundDisposition) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *RefundDisposition) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *RefundDisposition) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetDestinationAddress

`func (o *RefundDisposition) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *RefundDisposition) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *RefundDisposition) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.


### GetDispositionAmount

`func (o *RefundDisposition) GetDispositionAmount() string`

GetDispositionAmount returns the DispositionAmount field if non-nil, zero value otherwise.

### GetDispositionAmountOk

`func (o *RefundDisposition) GetDispositionAmountOk() (*string, bool)`

GetDispositionAmountOk returns a tuple with the DispositionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionAmount

`func (o *RefundDisposition) SetDispositionAmount(v string)`

SetDispositionAmount sets DispositionAmount field to given value.


### GetCategoryNames

`func (o *RefundDisposition) GetCategoryNames() []string`

GetCategoryNames returns the CategoryNames field if non-nil, zero value otherwise.

### GetCategoryNamesOk

`func (o *RefundDisposition) GetCategoryNamesOk() (*[]string, bool)`

GetCategoryNamesOk returns a tuple with the CategoryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryNames

`func (o *RefundDisposition) SetCategoryNames(v []string)`

SetCategoryNames sets CategoryNames field to given value.

### HasCategoryNames

`func (o *RefundDisposition) HasCategoryNames() bool`

HasCategoryNames returns a boolean if a field has been set.

### GetDescription

`func (o *RefundDisposition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RefundDisposition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RefundDisposition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RefundDisposition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


