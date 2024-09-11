/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"fmt"
)

// WebhookEventStatus The event status. Possible values include: - `Success`: The event has been delivered, and the webhook endpoint has responded to the event. - `Retrying`: The event has been delivered, but the webhook endpoint has not responded. In this case, Cobo will retry delivering the event. - `Failed`: The event cannot be delivered and Cobo will stop retrying. This may occur if the number of retries reaches 10, or if the event has been delivered but the webhook endpoint responded with an error.
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
    *v = WebhookEventStatus("unknown")
    return nil
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

