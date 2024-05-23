# WebhookEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The event ID. | 
**Url** | **string** | The URL of the webhook endpoint. | 
**CreatedTimestamp** | **int32** | The time when the event occurred, in Unix timestamp format, measured in milliseconds. | 
**Type** | [**WebhookEventType**](WebhookEventType.md) |  | 
**Data** | **map[string]interface{}** | The data of the webhook event, in JSON format. | 
**Status** | [**WebhookEventStatus**](WebhookEventStatus.md) |  | 
**NextRetryTimestamp** | Pointer to **int32** | The timestamp indicating the next scheduled retry to deliver this event, in Unix timestamp format, measured in milliseconds. This field is only present if the event status is &#x60;Retrying&#x60;.  | [optional] 
**RetriesLeft** | Pointer to **int32** | The number of retries left. This field is only present if the event status is &#x60;Retrying&#x60;. | [optional] 

## Methods

### NewWebhookEvent

`func NewWebhookEvent(id string, url string, createdTimestamp int32, type_ WebhookEventType, data map[string]interface{}, status WebhookEventStatus, ) *WebhookEvent`

NewWebhookEvent instantiates a new WebhookEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventWithDefaults

`func NewWebhookEventWithDefaults() *WebhookEvent`

NewWebhookEventWithDefaults instantiates a new WebhookEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookEvent) SetId(v string)`

SetId sets Id field to given value.


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

`func (o *WebhookEvent) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *WebhookEvent) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *WebhookEvent) SetCreatedTimestamp(v int32)`

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

`func (o *WebhookEvent) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookEvent) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookEvent) SetData(v map[string]interface{})`

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


### GetNextRetryTimestamp

`func (o *WebhookEvent) GetNextRetryTimestamp() int32`

GetNextRetryTimestamp returns the NextRetryTimestamp field if non-nil, zero value otherwise.

### GetNextRetryTimestampOk

`func (o *WebhookEvent) GetNextRetryTimestampOk() (*int32, bool)`

GetNextRetryTimestampOk returns a tuple with the NextRetryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRetryTimestamp

`func (o *WebhookEvent) SetNextRetryTimestamp(v int32)`

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


