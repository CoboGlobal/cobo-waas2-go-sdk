/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
	"fmt"
)

// WebhookEventStatus Possible values for the webhook event status: - `Success`: The event has been delivered, meaning that the webhook endpoint has received the event and responded successfully. - `Retrying`: The event has been delivered, but the webhook endpoint has not responded successfully. The event will be retried. - `Failed`: The event has been delivered, but the webhook endpoint has responded with an error. The event will not be retried.
type WebhookEventStatus string

// List of WebhookEventStatus
const (
	WEBHOOKEVENTSTATUS_SUCCESS WebhookEventStatus = "Success"
	WEBHOOKEVENTSTATUS_RETRYING WebhookEventStatus = "Retrying"
	WEBHOOKEVENTSTATUS_FAILED WebhookEventStatus = "Failed"
)

// All allowed values of WebhookEventStatus enum
var AllowedWebhookEventStatusEnumValues = []WebhookEventStatus{
	"Success",
	"Retrying",
	"Failed",
}

func (v *WebhookEventStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookEventStatus(value)
	for _, existing := range AllowedWebhookEventStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookEventStatus", value)
}

// NewWebhookEventStatusFromValue returns a pointer to a valid WebhookEventStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookEventStatusFromValue(v string) (*WebhookEventStatus, error) {
	ev := WebhookEventStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookEventStatus: valid values are %v", v, AllowedWebhookEventStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookEventStatus) IsValid() bool {
	for _, existing := range AllowedWebhookEventStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhookEventStatus value
func (v WebhookEventStatus) Ptr() *WebhookEventStatus {
	return &v
}

type NullableWebhookEventStatus struct {
	value *WebhookEventStatus
	isSet bool
}

func (v NullableWebhookEventStatus) Get() *WebhookEventStatus {
	return v.value
}

func (v *NullableWebhookEventStatus) Set(val *WebhookEventStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEventStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEventStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEventStatus(val *WebhookEventStatus) *NullableWebhookEventStatus {
	return &NullableWebhookEventStatus{value: val, isSet: true}
}

func (v NullableWebhookEventStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEventStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

