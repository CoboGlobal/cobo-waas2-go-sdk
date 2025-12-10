# SwapActivityTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action of the swap activity. Possible values include:   - &#x60;Submitted&#x60;: The swap request is submitted.   - &#x60;Pending Authorization&#x60;: The swap request is pending authorization.   - &#x60;Bridge {Token}&#x60;: The token is being bridged to the target chain.   - &#x60;Swap {Token}&#x60;: The token is being swapped on the target chain.   - &#x60;Cobo Confirmation&#x60;: The swap result is waiting for Cobo confirmation.  | 
**Status** | **string** | The status of the action. Possible values include:   - &#x60;Success&#x60;: The action is successfully completed.   - &#x60;Processing&#x60;: The action is being processed.   - &#x60;Failed&#x60;: The action has failed.  | 
**Timestamp** | Pointer to **int64** | The time when the action took place, in Unix timestamp format, measured in milliseconds.   | [optional] 

## Methods

### NewSwapActivityTimeline

`func NewSwapActivityTimeline(action string, status string, ) *SwapActivityTimeline`

NewSwapActivityTimeline instantiates a new SwapActivityTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapActivityTimelineWithDefaults

`func NewSwapActivityTimelineWithDefaults() *SwapActivityTimeline`

NewSwapActivityTimelineWithDefaults instantiates a new SwapActivityTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SwapActivityTimeline) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SwapActivityTimeline) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SwapActivityTimeline) SetAction(v string)`

SetAction sets Action field to given value.


### GetStatus

`func (o *SwapActivityTimeline) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwapActivityTimeline) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwapActivityTimeline) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *SwapActivityTimeline) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SwapActivityTimeline) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SwapActivityTimeline) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SwapActivityTimeline) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


