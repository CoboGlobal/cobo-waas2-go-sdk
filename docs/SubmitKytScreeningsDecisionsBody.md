# SubmitKytScreeningsDecisionsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The UUID of the transaction for KYT result submission. This identifies the specific transaction that the external KYT result applies to. | 
**Result** | [**KytScreeningsDecisionsType**](KytScreeningsDecisionsType.md) |  | 

## Methods

### NewSubmitKytScreeningsDecisionsBody

`func NewSubmitKytScreeningsDecisionsBody(transactionId string, result KytScreeningsDecisionsType, ) *SubmitKytScreeningsDecisionsBody`

NewSubmitKytScreeningsDecisionsBody instantiates a new SubmitKytScreeningsDecisionsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitKytScreeningsDecisionsBodyWithDefaults

`func NewSubmitKytScreeningsDecisionsBodyWithDefaults() *SubmitKytScreeningsDecisionsBody`

NewSubmitKytScreeningsDecisionsBodyWithDefaults instantiates a new SubmitKytScreeningsDecisionsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *SubmitKytScreeningsDecisionsBody) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *SubmitKytScreeningsDecisionsBody) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *SubmitKytScreeningsDecisionsBody) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetResult

`func (o *SubmitKytScreeningsDecisionsBody) GetResult() KytScreeningsDecisionsType`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SubmitKytScreeningsDecisionsBody) GetResultOk() (*KytScreeningsDecisionsType, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SubmitKytScreeningsDecisionsBody) SetResult(v KytScreeningsDecisionsType)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


