/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the ListWebhookEventDefinitions200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWebhookEventDefinitions200ResponseInner{}

// ListWebhookEventDefinitions200ResponseInner struct for ListWebhookEventDefinitions200ResponseInner
type ListWebhookEventDefinitions200ResponseInner struct {
	EventType *WebhookEventType `json:"event_type,omitempty"`
	// The description of the webhook event type.
	Description *string `json:"description,omitempty"`
}

// NewListWebhookEventDefinitions200ResponseInner instantiates a new ListWebhookEventDefinitions200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWebhookEventDefinitions200ResponseInner() *ListWebhookEventDefinitions200ResponseInner {
	this := ListWebhookEventDefinitions200ResponseInner{}
	return &this
}

// NewListWebhookEventDefinitions200ResponseInnerWithDefaults instantiates a new ListWebhookEventDefinitions200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWebhookEventDefinitions200ResponseInnerWithDefaults() *ListWebhookEventDefinitions200ResponseInner {
	this := ListWebhookEventDefinitions200ResponseInner{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *ListWebhookEventDefinitions200ResponseInner) GetEventType() WebhookEventType {
	if o == nil || IsNil(o.EventType) {
		var ret WebhookEventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhookEventDefinitions200ResponseInner) GetEventTypeOk() (*WebhookEventType, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *ListWebhookEventDefinitions200ResponseInner) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given WebhookEventType and assigns it to the EventType field.
func (o *ListWebhookEventDefinitions200ResponseInner) SetEventType(v WebhookEventType) {
	o.EventType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListWebhookEventDefinitions200ResponseInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhookEventDefinitions200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ListWebhookEventDefinitions200ResponseInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListWebhookEventDefinitions200ResponseInner) SetDescription(v string) {
	o.Description = &v
}

func (o ListWebhookEventDefinitions200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWebhookEventDefinitions200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableListWebhookEventDefinitions200ResponseInner struct {
	value *ListWebhookEventDefinitions200ResponseInner
	isSet bool
}

func (v NullableListWebhookEventDefinitions200ResponseInner) Get() *ListWebhookEventDefinitions200ResponseInner {
	return v.value
}

func (v *NullableListWebhookEventDefinitions200ResponseInner) Set(val *ListWebhookEventDefinitions200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListWebhookEventDefinitions200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListWebhookEventDefinitions200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWebhookEventDefinitions200ResponseInner(val *ListWebhookEventDefinitions200ResponseInner) *NullableListWebhookEventDefinitions200ResponseInner {
	return &NullableListWebhookEventDefinitions200ResponseInner{value: val, isSet: true}
}

func (v NullableListWebhookEventDefinitions200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWebhookEventDefinitions200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


