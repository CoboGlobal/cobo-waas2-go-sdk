# SubmitKytScreeningsReviewBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The UUID of the transaction for external KYT review. This identifies the transaction that requires external compliance review. | 
**Result** | [**KytScreeningsReviewType**](KytScreeningsReviewType.md) |  | 

## Methods

### NewSubmitKytScreeningsReviewBody

`func NewSubmitKytScreeningsReviewBody(transactionId string, result KytScreeningsReviewType, ) *SubmitKytScreeningsReviewBody`

NewSubmitKytScreeningsReviewBody instantiates a new SubmitKytScreeningsReviewBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitKytScreeningsReviewBodyWithDefaults

`func NewSubmitKytScreeningsReviewBodyWithDefaults() *SubmitKytScreeningsReviewBody`

NewSubmitKytScreeningsReviewBodyWithDefaults instantiates a new SubmitKytScreeningsReviewBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *SubmitKytScreeningsReviewBody) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *SubmitKytScreeningsReviewBody) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *SubmitKytScreeningsReviewBody) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetResult

`func (o *SubmitKytScreeningsReviewBody) GetResult() KytScreeningsReviewType`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SubmitKytScreeningsReviewBody) GetResultOk() (*KytScreeningsReviewType, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SubmitKytScreeningsReviewBody) SetResult(v KytScreeningsReviewType)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


