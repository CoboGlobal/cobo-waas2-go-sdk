# DispositionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The UUID of the transaction being processed for disposition. | 
**Status** | [**DispositionStatus**](DispositionStatus.md) |  | 

## Methods

### NewDispositionResponse

`func NewDispositionResponse(transactionId string, status DispositionStatus, ) *DispositionResponse`

NewDispositionResponse instantiates a new DispositionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDispositionResponseWithDefaults

`func NewDispositionResponseWithDefaults() *DispositionResponse`

NewDispositionResponseWithDefaults instantiates a new DispositionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *DispositionResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DispositionResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DispositionResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetStatus

`func (o *DispositionResponse) GetStatus() DispositionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DispositionResponse) GetStatusOk() (*DispositionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DispositionResponse) SetStatus(v DispositionStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


