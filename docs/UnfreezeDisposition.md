# UnfreezeDisposition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The UUID of the transaction to be unfrozen. This identifies the frozen transaction that needs to be released. | 

## Methods

### NewUnfreezeDisposition

`func NewUnfreezeDisposition(transactionId string, ) *UnfreezeDisposition`

NewUnfreezeDisposition instantiates a new UnfreezeDisposition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnfreezeDispositionWithDefaults

`func NewUnfreezeDispositionWithDefaults() *UnfreezeDisposition`

NewUnfreezeDispositionWithDefaults instantiates a new UnfreezeDisposition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *UnfreezeDisposition) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *UnfreezeDisposition) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *UnfreezeDisposition) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


