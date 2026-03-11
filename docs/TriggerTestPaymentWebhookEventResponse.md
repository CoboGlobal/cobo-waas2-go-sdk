# TriggerTestPaymentWebhookEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Triggered** | Pointer to **bool** | Whether a test webhook event was successfully triggered. - &#x60;true&#x60;: The test webhook event was successfully triggered. - &#x60;false&#x60;: The test webhook event could not be triggered.  | [optional] 

## Methods

### NewTriggerTestPaymentWebhookEventResponse

`func NewTriggerTestPaymentWebhookEventResponse() *TriggerTestPaymentWebhookEventResponse`

NewTriggerTestPaymentWebhookEventResponse instantiates a new TriggerTestPaymentWebhookEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerTestPaymentWebhookEventResponseWithDefaults

`func NewTriggerTestPaymentWebhookEventResponseWithDefaults() *TriggerTestPaymentWebhookEventResponse`

NewTriggerTestPaymentWebhookEventResponseWithDefaults instantiates a new TriggerTestPaymentWebhookEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggered

`func (o *TriggerTestPaymentWebhookEventResponse) GetTriggered() bool`

GetTriggered returns the Triggered field if non-nil, zero value otherwise.

### GetTriggeredOk

`func (o *TriggerTestPaymentWebhookEventResponse) GetTriggeredOk() (*bool, bool)`

GetTriggeredOk returns a tuple with the Triggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggered

`func (o *TriggerTestPaymentWebhookEventResponse) SetTriggered(v bool)`

SetTriggered sets Triggered field to given value.

### HasTriggered

`func (o *TriggerTestPaymentWebhookEventResponse) HasTriggered() bool`

HasTriggered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


