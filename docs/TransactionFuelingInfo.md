# TransactionFuelingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The request ID of the transaction. | [optional] 
**TransactionId** | Pointer to **string** | The transaction ID. | [optional] 
**MainTransactionId** | Pointer to **string** | The UUID of the parent (main) transaction that this record is associated with. Set only when the current record is a gas/fee transaction initiated by FeeStation; omit for main transactions. | [optional] 

## Methods

### NewTransactionFuelingInfo

`func NewTransactionFuelingInfo() *TransactionFuelingInfo`

NewTransactionFuelingInfo instantiates a new TransactionFuelingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFuelingInfoWithDefaults

`func NewTransactionFuelingInfoWithDefaults() *TransactionFuelingInfo`

NewTransactionFuelingInfoWithDefaults instantiates a new TransactionFuelingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *TransactionFuelingInfo) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransactionFuelingInfo) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransactionFuelingInfo) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TransactionFuelingInfo) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTransactionId

`func (o *TransactionFuelingInfo) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransactionFuelingInfo) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransactionFuelingInfo) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *TransactionFuelingInfo) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetMainTransactionId

`func (o *TransactionFuelingInfo) GetMainTransactionId() string`

GetMainTransactionId returns the MainTransactionId field if non-nil, zero value otherwise.

### GetMainTransactionIdOk

`func (o *TransactionFuelingInfo) GetMainTransactionIdOk() (*string, bool)`

GetMainTransactionIdOk returns a tuple with the MainTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainTransactionId

`func (o *TransactionFuelingInfo) SetMainTransactionId(v string)`

SetMainTransactionId sets MainTransactionId field to given value.

### HasMainTransactionId

`func (o *TransactionFuelingInfo) HasMainTransactionId() bool`

HasMainTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


