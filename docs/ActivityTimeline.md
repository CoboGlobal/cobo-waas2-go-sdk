# ActivityTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**ActivityAction**](ActivityAction.md) |  | 
**Status** | Pointer to **string** | The status of the action. Possible values include:   - &#x60;Success&#x60;: The action is successfully completed.   - &#x60;Processing&#x60;: The action is being processed.   - &#x60;Failed&#x60;: The action has failed.  | [optional] 
**Timestamp** | Pointer to **int64** | The time when the action took place, in Unix timestamp format, measured in milliseconds.  - For the &#x60;Submitted&#x60; action, &#x60;timestamp&#x60; represents the time the staking, unstaking, or withdrawal request was created.  - For the &#x60;BTCConfirmation&#x60; action, &#x60;timestamp&#x60; represents the time when the request was confirmed on the Bitcoin chain, or when the confirmation failed. - For the &#x60;BabylonConfirmation&#x60; action, &#x60;timestamp&#x60; represents the time when the request was confirmed by the Babylon protocol, or when the confirmation failed.  | [optional] 
**TransactionId** | Pointer to **string** | The ID of the corresponding transaction. | [optional] 

## Methods

### NewActivityTimeline

`func NewActivityTimeline(action ActivityAction, ) *ActivityTimeline`

NewActivityTimeline instantiates a new ActivityTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityTimelineWithDefaults

`func NewActivityTimelineWithDefaults() *ActivityTimeline`

NewActivityTimelineWithDefaults instantiates a new ActivityTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ActivityTimeline) GetAction() ActivityAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActivityTimeline) GetActionOk() (*ActivityAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActivityTimeline) SetAction(v ActivityAction)`

SetAction sets Action field to given value.


### GetStatus

`func (o *ActivityTimeline) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActivityTimeline) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActivityTimeline) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActivityTimeline) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *ActivityTimeline) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityTimeline) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityTimeline) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ActivityTimeline) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTransactionId

`func (o *ActivityTimeline) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ActivityTimeline) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ActivityTimeline) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ActivityTimeline) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


