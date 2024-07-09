# ListWebhookEventDefinitions200ResponseResultInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to [**WebhookEventType**](WebhookEventType.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the webhook event type. | [optional] 

## Methods

### NewListWebhookEventDefinitions200ResponseResultInner

`func NewListWebhookEventDefinitions200ResponseResultInner() *ListWebhookEventDefinitions200ResponseResultInner`

NewListWebhookEventDefinitions200ResponseResultInner instantiates a new ListWebhookEventDefinitions200ResponseResultInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebhookEventDefinitions200ResponseResultInnerWithDefaults

`func NewListWebhookEventDefinitions200ResponseResultInnerWithDefaults() *ListWebhookEventDefinitions200ResponseResultInner`

NewListWebhookEventDefinitions200ResponseResultInnerWithDefaults instantiates a new ListWebhookEventDefinitions200ResponseResultInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *ListWebhookEventDefinitions200ResponseResultInner) GetEventType() WebhookEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ListWebhookEventDefinitions200ResponseResultInner) GetEventTypeOk() (*WebhookEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ListWebhookEventDefinitions200ResponseResultInner) SetEventType(v WebhookEventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ListWebhookEventDefinitions200ResponseResultInner) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetDescription

`func (o *ListWebhookEventDefinitions200ResponseResultInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListWebhookEventDefinitions200ResponseResultInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListWebhookEventDefinitions200ResponseResultInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListWebhookEventDefinitions200ResponseResultInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


