# TransactionApprover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The approver name of the transaction. | [optional] 
**Status** | Pointer to **string** | The approval status. | [optional] 

## Methods

### NewTransactionApprover

`func NewTransactionApprover() *TransactionApprover`

NewTransactionApprover instantiates a new TransactionApprover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionApproverWithDefaults

`func NewTransactionApproverWithDefaults() *TransactionApprover`

NewTransactionApproverWithDefaults instantiates a new TransactionApprover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TransactionApprover) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransactionApprover) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransactionApprover) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransactionApprover) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionApprover) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionApprover) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionApprover) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionApprover) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


