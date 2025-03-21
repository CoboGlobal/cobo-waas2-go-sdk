/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the UpdateWebhookEndpointByIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateWebhookEndpointByIdRequest{}

// UpdateWebhookEndpointByIdRequest struct for UpdateWebhookEndpointByIdRequest
type UpdateWebhookEndpointByIdRequest struct {
	// The new event types you want to subscribe to for this webhook endpoint. You can call [Get webhook event types](https://www.cobo.com/developers/v2/api-references/developers--webhooks/get-webhook-event-types) to retrieve all available event types.
	SubscribedEvents []WebhookEventType `json:"subscribed_events,omitempty"`
	// The new status you want to set the webhook endpoint to. If you set `status` to `STATUS_INACTIVE`, the endpoint will be revoked, meaning it will no longer receive any webhook events.
	Status *string `json:"status,omitempty"`
	// The webhook endpoint description.
	Description *string `json:"description,omitempty"`
}

// NewUpdateWebhookEndpointByIdRequest instantiates a new UpdateWebhookEndpointByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebhookEndpointByIdRequest() *UpdateWebhookEndpointByIdRequest {
	this := UpdateWebhookEndpointByIdRequest{}
	return &this
}

// NewUpdateWebhookEndpointByIdRequestWithDefaults instantiates a new UpdateWebhookEndpointByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebhookEndpointByIdRequestWithDefaults() *UpdateWebhookEndpointByIdRequest {
	this := UpdateWebhookEndpointByIdRequest{}
	return &this
}

// GetSubscribedEvents returns the SubscribedEvents field value if set, zero value otherwise.
func (o *UpdateWebhookEndpointByIdRequest) GetSubscribedEvents() []WebhookEventType {
	if o == nil || IsNil(o.SubscribedEvents) {
		var ret []WebhookEventType
		return ret
	}
	return o.SubscribedEvents
}

// GetSubscribedEventsOk returns a tuple with the SubscribedEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookEndpointByIdRequest) GetSubscribedEventsOk() ([]WebhookEventType, bool) {
	if o == nil || IsNil(o.SubscribedEvents) {
		return nil, false
	}
	return o.SubscribedEvents, true
}

// HasSubscribedEvents returns a boolean if a field has been set.
func (o *UpdateWebhookEndpointByIdRequest) HasSubscribedEvents() bool {
	if o != nil && !IsNil(o.SubscribedEvents) {
		return true
	}

	return false
}

// SetSubscribedEvents gets a reference to the given []WebhookEventType and assigns it to the SubscribedEvents field.
func (o *UpdateWebhookEndpointByIdRequest) SetSubscribedEvents(v []WebhookEventType) {
	o.SubscribedEvents = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateWebhookEndpointByIdRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookEndpointByIdRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateWebhookEndpointByIdRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdateWebhookEndpointByIdRequest) SetStatus(v string) {
	o.Status = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateWebhookEndpointByIdRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookEndpointByIdRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateWebhookEndpointByIdRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateWebhookEndpointByIdRequest) SetDescription(v string) {
	o.Description = &v
}

func (o UpdateWebhookEndpointByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateWebhookEndpointByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscribedEvents) {
		toSerialize["subscribed_events"] = o.SubscribedEvents
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableUpdateWebhookEndpointByIdRequest struct {
	value *UpdateWebhookEndpointByIdRequest
	isSet bool
}

func (v NullableUpdateWebhookEndpointByIdRequest) Get() *UpdateWebhookEndpointByIdRequest {
	return v.value
}

func (v *NullableUpdateWebhookEndpointByIdRequest) Set(val *UpdateWebhookEndpointByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebhookEndpointByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebhookEndpointByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebhookEndpointByIdRequest(val *UpdateWebhookEndpointByIdRequest) *NullableUpdateWebhookEndpointByIdRequest {
	return &NullableUpdateWebhookEndpointByIdRequest{value: val, isSet: true}
}

func (v NullableUpdateWebhookEndpointByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebhookEndpointByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


