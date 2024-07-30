# TransactionTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**TransactionStatus**](TransactionStatus.md) |  | [optional] 
**Finished** | Pointer to **bool** | Whether the transaction status is completed: - &#x60;true&#x60;: The transaction status is completed. - &#x60;false&#x60;: The transaction is currently in the status.  | [optional] 
**FinishedTime** | Pointer to **int64** | The time when the transaction status is completed in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewTransactionTimeline

`func NewTransactionTimeline() *TransactionTimeline`

NewTransactionTimeline instantiates a new TransactionTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTimelineWithDefaults

`func NewTransactionTimelineWithDefaults() *TransactionTimeline`

NewTransactionTimelineWithDefaults instantiates a new TransactionTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TransactionTimeline) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionTimeline) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionTimeline) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionTimeline) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFinished

`func (o *TransactionTimeline) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *TransactionTimeline) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *TransactionTimeline) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *TransactionTimeline) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetFinishedTime

`func (o *TransactionTimeline) GetFinishedTime() int64`

GetFinishedTime returns the FinishedTime field if non-nil, zero value otherwise.

### GetFinishedTimeOk

`func (o *TransactionTimeline) GetFinishedTimeOk() (*int64, bool)`

GetFinishedTimeOk returns a tuple with the FinishedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedTime

`func (o *TransactionTimeline) SetFinishedTime(v int64)`

SetFinishedTime sets FinishedTime field to given value.

### HasFinishedTime

`func (o *TransactionTimeline) HasFinishedTime() bool`

HasFinishedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


