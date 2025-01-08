# RetryCallbackMessage201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retried** | Pointer to **bool** | Whether the callback message has been successfully resent: - &#x60;true&#x60;: The callback message has been successfully resent. - &#x60;false&#x60;: The callback message has not been successfully resent.  | [optional] 

## Methods

### NewRetryCallbackMessage201Response

`func NewRetryCallbackMessage201Response() *RetryCallbackMessage201Response`

NewRetryCallbackMessage201Response instantiates a new RetryCallbackMessage201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryCallbackMessage201ResponseWithDefaults

`func NewRetryCallbackMessage201ResponseWithDefaults() *RetryCallbackMessage201Response`

NewRetryCallbackMessage201ResponseWithDefaults instantiates a new RetryCallbackMessage201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetried

`func (o *RetryCallbackMessage201Response) GetRetried() bool`

GetRetried returns the Retried field if non-nil, zero value otherwise.

### GetRetriedOk

`func (o *RetryCallbackMessage201Response) GetRetriedOk() (*bool, bool)`

GetRetriedOk returns a tuple with the Retried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetried

`func (o *RetryCallbackMessage201Response) SetRetried(v bool)`

SetRetried sets Retried field to given value.

### HasRetried

`func (o *RetryCallbackMessage201Response) HasRetried() bool`

HasRetried returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


