# IsolateDisposition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The UUID of the transaction whose funds are to be isolated. This identifies the original transaction that requires fund isolation. | 
**DestinationAddress** | **string** | The blockchain address to receive the isolated funds. | 
**DispositionAmount** | **string** | The amount to be isolated from the original transaction, specified as a numeric string. This value cannot exceed the total amount of the original transaction.  | 
**CategoryNames** | Pointer to **[]string** | Custom categories to identify and track this isolation transaction. Used for transaction classification and reporting. | [optional] 
**Description** | Pointer to **string** | Additional notes or description for the isolation. | [optional] 

## Methods

### NewIsolateDisposition

`func NewIsolateDisposition(transactionId string, destinationAddress string, dispositionAmount string, ) *IsolateDisposition`

NewIsolateDisposition instantiates a new IsolateDisposition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsolateDispositionWithDefaults

`func NewIsolateDispositionWithDefaults() *IsolateDisposition`

NewIsolateDispositionWithDefaults instantiates a new IsolateDisposition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *IsolateDisposition) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *IsolateDisposition) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *IsolateDisposition) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetDestinationAddress

`func (o *IsolateDisposition) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *IsolateDisposition) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *IsolateDisposition) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.


### GetDispositionAmount

`func (o *IsolateDisposition) GetDispositionAmount() string`

GetDispositionAmount returns the DispositionAmount field if non-nil, zero value otherwise.

### GetDispositionAmountOk

`func (o *IsolateDisposition) GetDispositionAmountOk() (*string, bool)`

GetDispositionAmountOk returns a tuple with the DispositionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionAmount

`func (o *IsolateDisposition) SetDispositionAmount(v string)`

SetDispositionAmount sets DispositionAmount field to given value.


### GetCategoryNames

`func (o *IsolateDisposition) GetCategoryNames() []string`

GetCategoryNames returns the CategoryNames field if non-nil, zero value otherwise.

### GetCategoryNamesOk

`func (o *IsolateDisposition) GetCategoryNamesOk() (*[]string, bool)`

GetCategoryNamesOk returns a tuple with the CategoryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryNames

`func (o *IsolateDisposition) SetCategoryNames(v []string)`

SetCategoryNames sets CategoryNames field to given value.

### HasCategoryNames

`func (o *IsolateDisposition) HasCategoryNames() bool`

HasCategoryNames returns a boolean if a field has been set.

### GetDescription

`func (o *IsolateDisposition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IsolateDisposition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IsolateDisposition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IsolateDisposition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


