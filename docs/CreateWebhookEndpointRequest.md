# CreateWebhookEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The webhook endpoint URL. | 
**SubscribedEvents** | [**[]WebhookEventType**](WebhookEventType.md) | The event types you want to subscribe to for this webhook endpoint. You can call [Get webhook event types](https://www.cobo.com/developers/v2/api-references/developers--webhooks/get-webhook-event-types) to retrieve all available event types.  | 
**Description** | Pointer to **string** | The description of the webhook endpoint. | [optional] 

## Methods

### NewCreateWebhookEndpointRequest

`func NewCreateWebhookEndpointRequest(url string, subscribedEvents []WebhookEventType, ) *CreateWebhookEndpointRequest`

NewCreateWebhookEndpointRequest instantiates a new CreateWebhookEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookEndpointRequestWithDefaults

`func NewCreateWebhookEndpointRequestWithDefaults() *CreateWebhookEndpointRequest`

NewCreateWebhookEndpointRequestWithDefaults instantiates a new CreateWebhookEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CreateWebhookEndpointRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateWebhookEndpointRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateWebhookEndpointRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSubscribedEvents

`func (o *CreateWebhookEndpointRequest) GetSubscribedEvents() []WebhookEventType`

GetSubscribedEvents returns the SubscribedEvents field if non-nil, zero value otherwise.

### GetSubscribedEventsOk

`func (o *CreateWebhookEndpointRequest) GetSubscribedEventsOk() (*[]WebhookEventType, bool)`

GetSubscribedEventsOk returns a tuple with the SubscribedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEvents

`func (o *CreateWebhookEndpointRequest) SetSubscribedEvents(v []WebhookEventType)`

SetSubscribedEvents sets SubscribedEvents field to given value.


### GetDescription

`func (o *CreateWebhookEndpointRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateWebhookEndpointRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateWebhookEndpointRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateWebhookEndpointRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


