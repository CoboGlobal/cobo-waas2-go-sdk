# TriggerTestWebhookEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**WebhookEventType**](WebhookEventType.md) |  | 
**OverrideData** | Pointer to **map[string]interface{}** | An object for customization of the webhook event payload. You only need to include the fields you want to customize.   The provided fields must match the webhook event data structure, depending on the specified event type. For a complete introduction of the webhook event data structure, refer to the &#x60;data.data&#x60; property in the response of [List all webhook events](/v2/api-references/developers--webhooks/list-all-webhook-events).  If this property is not provided, a default payload will be returned.  | [optional] 

## Methods

### NewTriggerTestWebhookEventRequest

`func NewTriggerTestWebhookEventRequest(eventType WebhookEventType, ) *TriggerTestWebhookEventRequest`

NewTriggerTestWebhookEventRequest instantiates a new TriggerTestWebhookEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerTestWebhookEventRequestWithDefaults

`func NewTriggerTestWebhookEventRequestWithDefaults() *TriggerTestWebhookEventRequest`

NewTriggerTestWebhookEventRequestWithDefaults instantiates a new TriggerTestWebhookEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *TriggerTestWebhookEventRequest) GetEventType() WebhookEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TriggerTestWebhookEventRequest) GetEventTypeOk() (*WebhookEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TriggerTestWebhookEventRequest) SetEventType(v WebhookEventType)`

SetEventType sets EventType field to given value.


### GetOverrideData

`func (o *TriggerTestWebhookEventRequest) GetOverrideData() map[string]interface{}`

GetOverrideData returns the OverrideData field if non-nil, zero value otherwise.

### GetOverrideDataOk

`func (o *TriggerTestWebhookEventRequest) GetOverrideDataOk() (*map[string]interface{}, bool)`

GetOverrideDataOk returns a tuple with the OverrideData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideData

`func (o *TriggerTestWebhookEventRequest) SetOverrideData(v map[string]interface{})`

SetOverrideData sets OverrideData field to given value.

### HasOverrideData

`func (o *TriggerTestWebhookEventRequest) HasOverrideData() bool`

HasOverrideData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


