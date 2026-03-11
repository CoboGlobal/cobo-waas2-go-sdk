# TriggerTestPaymentsWebhookEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**WebhookEventType**](WebhookEventType.md) |  | 
**OverrideData** | Pointer to **map[string]interface{}** | An optional object to customize the webhook event payload. Include only the fields you want to override.  The provided fields must match the webhook event data structure for the specified event type. For the full event data structure, refer to the &#x60;data.data&#x60; property in the response of [List all webhook events](https://www.cobo.com/developers/v2/api-references/developers--webhooks/list-all-webhook-events).  If this property is omitted, a default payload is returned.  | [optional] 

## Methods

### NewTriggerTestPaymentsWebhookEventRequest

`func NewTriggerTestPaymentsWebhookEventRequest(eventType WebhookEventType, ) *TriggerTestPaymentsWebhookEventRequest`

NewTriggerTestPaymentsWebhookEventRequest instantiates a new TriggerTestPaymentsWebhookEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerTestPaymentsWebhookEventRequestWithDefaults

`func NewTriggerTestPaymentsWebhookEventRequestWithDefaults() *TriggerTestPaymentsWebhookEventRequest`

NewTriggerTestPaymentsWebhookEventRequestWithDefaults instantiates a new TriggerTestPaymentsWebhookEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *TriggerTestPaymentsWebhookEventRequest) GetEventType() WebhookEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TriggerTestPaymentsWebhookEventRequest) GetEventTypeOk() (*WebhookEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TriggerTestPaymentsWebhookEventRequest) SetEventType(v WebhookEventType)`

SetEventType sets EventType field to given value.


### GetOverrideData

`func (o *TriggerTestPaymentsWebhookEventRequest) GetOverrideData() map[string]interface{}`

GetOverrideData returns the OverrideData field if non-nil, zero value otherwise.

### GetOverrideDataOk

`func (o *TriggerTestPaymentsWebhookEventRequest) GetOverrideDataOk() (*map[string]interface{}, bool)`

GetOverrideDataOk returns a tuple with the OverrideData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideData

`func (o *TriggerTestPaymentsWebhookEventRequest) SetOverrideData(v map[string]interface{})`

SetOverrideData sets OverrideData field to given value.

### HasOverrideData

`func (o *TriggerTestPaymentsWebhookEventRequest) HasOverrideData() bool`

HasOverrideData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


