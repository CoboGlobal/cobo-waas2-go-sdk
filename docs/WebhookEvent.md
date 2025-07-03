# WebhookEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** | The event ID. | [optional] 
**Url** | **string** | The webhook endpoint URL. | 
**CreatedTimestamp** | **int64** | The time when the event type was triggered, in Unix timestamp format (milliseconds). - The value remains unchanged across retries. - The default webhook timeout is 2 seconds.  | 
**Type** | [**WebhookEventType**](WebhookEventType.md) |  | 
**Data** | [**WebhookEventData**](WebhookEventData.md) |  | 
**Status** | Pointer to [**WebhookEventStatus**](WebhookEventStatus.md) |  | [optional] 
**NextRetryTimestamp** | Pointer to **int64** | The timestamp indicating the next scheduled retry to deliver this event, in Unix timestamp format, measured in milliseconds. This field is only present if the event status is &#x60;Retrying&#x60;.  | [optional] 
**RetriesLeft** | Pointer to **int32** | The number of retries left. This field is only present if the event status is &#x60;Retrying&#x60;. | [optional] 

## Methods

### NewWebhookEvent

`func NewWebhookEvent(url string, createdTimestamp int64, type_ WebhookEventType, data WebhookEventData, ) *WebhookEvent`

NewWebhookEvent instantiates a new WebhookEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventWithDefaults

`func NewWebhookEventWithDefaults() *WebhookEvent`

NewWebhookEventWithDefaults instantiates a new WebhookEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *WebhookEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *WebhookEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *WebhookEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *WebhookEvent) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookEvent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCreatedTimestamp

`func (o *WebhookEvent) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *WebhookEvent) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *WebhookEvent) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetType

`func (o *WebhookEvent) GetType() WebhookEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookEvent) GetTypeOk() (*WebhookEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookEvent) SetType(v WebhookEventType)`

SetType sets Type field to given value.


### GetData

`func (o *WebhookEvent) GetData() WebhookEventData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookEvent) GetDataOk() (*WebhookEventData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookEvent) SetData(v WebhookEventData)`

SetData sets Data field to given value.


### GetStatus

`func (o *WebhookEvent) GetStatus() WebhookEventStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookEvent) GetStatusOk() (*WebhookEventStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookEvent) SetStatus(v WebhookEventStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNextRetryTimestamp

`func (o *WebhookEvent) GetNextRetryTimestamp() int64`

GetNextRetryTimestamp returns the NextRetryTimestamp field if non-nil, zero value otherwise.

### GetNextRetryTimestampOk

`func (o *WebhookEvent) GetNextRetryTimestampOk() (*int64, bool)`

GetNextRetryTimestampOk returns a tuple with the NextRetryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRetryTimestamp

`func (o *WebhookEvent) SetNextRetryTimestamp(v int64)`

SetNextRetryTimestamp sets NextRetryTimestamp field to given value.

### HasNextRetryTimestamp

`func (o *WebhookEvent) HasNextRetryTimestamp() bool`

HasNextRetryTimestamp returns a boolean if a field has been set.

### GetRetriesLeft

`func (o *WebhookEvent) GetRetriesLeft() int32`

GetRetriesLeft returns the RetriesLeft field if non-nil, zero value otherwise.

### GetRetriesLeftOk

`func (o *WebhookEvent) GetRetriesLeftOk() (*int32, bool)`

GetRetriesLeftOk returns a tuple with the RetriesLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetriesLeft

`func (o *WebhookEvent) SetRetriesLeft(v int32)`

SetRetriesLeft sets RetriesLeft field to given value.

### HasRetriesLeft

`func (o *WebhookEvent) HasRetriesLeft() bool`

HasRetriesLeft returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


