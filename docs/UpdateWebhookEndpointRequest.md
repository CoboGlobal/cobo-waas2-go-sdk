# UpdateWebhookEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscribedEvents** | Pointer to [**[]WebhookEventType**](WebhookEventType.md) | The new event types you want to subscribe to for this webhook endpoint. You can call [Get webhook event types](/v2/api-references/developers--webhooks/get-webhook-event-types) to retrieve all available event types. | [optional] 
**Status** | Pointer to **string** | The new status you want to set the webhook endpoint to. If you set &#x60;status&#x60; to &#x60;STATUS_INACTIVE&#x60;, the endpoint will be revoked, meaning it will no longer receive any webhook events. | [optional] 
**Description** | Pointer to **string** | The webhook endpoint description. | [optional] 

## Methods

### NewUpdateWebhookEndpointRequest

`func NewUpdateWebhookEndpointRequest() *UpdateWebhookEndpointRequest`

NewUpdateWebhookEndpointRequest instantiates a new UpdateWebhookEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWebhookEndpointRequestWithDefaults

`func NewUpdateWebhookEndpointRequestWithDefaults() *UpdateWebhookEndpointRequest`

NewUpdateWebhookEndpointRequestWithDefaults instantiates a new UpdateWebhookEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscribedEvents

`func (o *UpdateWebhookEndpointRequest) GetSubscribedEvents() []WebhookEventType`

GetSubscribedEvents returns the SubscribedEvents field if non-nil, zero value otherwise.

### GetSubscribedEventsOk

`func (o *UpdateWebhookEndpointRequest) GetSubscribedEventsOk() (*[]WebhookEventType, bool)`

GetSubscribedEventsOk returns a tuple with the SubscribedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEvents

`func (o *UpdateWebhookEndpointRequest) SetSubscribedEvents(v []WebhookEventType)`

SetSubscribedEvents sets SubscribedEvents field to given value.

### HasSubscribedEvents

`func (o *UpdateWebhookEndpointRequest) HasSubscribedEvents() bool`

HasSubscribedEvents returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateWebhookEndpointRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateWebhookEndpointRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateWebhookEndpointRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateWebhookEndpointRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateWebhookEndpointRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateWebhookEndpointRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateWebhookEndpointRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateWebhookEndpointRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


