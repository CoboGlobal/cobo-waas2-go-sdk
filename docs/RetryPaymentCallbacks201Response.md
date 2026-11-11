# RetryPaymentCallbacks201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retried** | **bool** | Whether the callback retry request was processed. - &#x60;true&#x60;: The request was processed. This does not guarantee that any callback message was resent or approved, or that the business record has continued processing. - &#x60;false&#x60;: The request was not processed.  | 

## Methods

### NewRetryPaymentCallbacks201Response

`func NewRetryPaymentCallbacks201Response(retried bool, ) *RetryPaymentCallbacks201Response`

NewRetryPaymentCallbacks201Response instantiates a new RetryPaymentCallbacks201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryPaymentCallbacks201ResponseWithDefaults

`func NewRetryPaymentCallbacks201ResponseWithDefaults() *RetryPaymentCallbacks201Response`

NewRetryPaymentCallbacks201ResponseWithDefaults instantiates a new RetryPaymentCallbacks201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetried

`func (o *RetryPaymentCallbacks201Response) GetRetried() bool`

GetRetried returns the Retried field if non-nil, zero value otherwise.

### GetRetriedOk

`func (o *RetryPaymentCallbacks201Response) GetRetriedOk() (*bool, bool)`

GetRetriedOk returns a tuple with the Retried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetried

`func (o *RetryPaymentCallbacks201Response) SetRetried(v bool)`

SetRetried sets Retried field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


