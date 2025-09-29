# SubmitKytResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The UUID of the transaction being processed for screened. | 
**Submitted** | **bool** | Indicates whether the KYT submission was successfully submitted. | 

## Methods

### NewSubmitKytResponse

`func NewSubmitKytResponse(transactionId string, submitted bool, ) *SubmitKytResponse`

NewSubmitKytResponse instantiates a new SubmitKytResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitKytResponseWithDefaults

`func NewSubmitKytResponseWithDefaults() *SubmitKytResponse`

NewSubmitKytResponseWithDefaults instantiates a new SubmitKytResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *SubmitKytResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *SubmitKytResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *SubmitKytResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetSubmitted

`func (o *SubmitKytResponse) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *SubmitKytResponse) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *SubmitKytResponse) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


