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

// WebhookEventType The event type. To learn the trigger condition of each event type, refer to [Webhook event types and event data](https://www.cobo.com/developers/v2/guides/webhooks-callbacks/webhook-event-type).
type WebhookEventType string

// List of WebhookEventType
const (
	WEBHOOKEVENTTYPE_WALLETS_TRANSACTION_CREATED WebhookEventType = "wallets.transaction.created"
	WEBHOOKEVENTTYPE_WALLETS_TRANSACTION_UPDATED WebhookEventType = "wallets.transaction.updated"
	WEBHOOKEVENTTYPE_WALLETS_TRANSACTION_FAILED WebhookEventType = "wallets.transaction.failed"
	WEBHOOKEVENTTYPE_WALLETS_TRANSACTION_SUCCEEDED WebhookEventType = "wallets.transaction.succeeded"
	WEBHOOKEVENTTYPE_WALLETS_MPC_TSS_REQUEST_CREATED WebhookEventType = "wallets.mpc.tss_request.created"
	WEBHOOKEVENTTYPE_WALLETS_MPC_TSS_REQUEST_UPDATED WebhookEventType = "wallets.mpc.tss_request.updated"
	WEBHOOKEVENTTYPE_WALLETS_MPC_TSS_REQUEST_FAILED WebhookEventType = "wallets.mpc.tss_request.failed"
	WEBHOOKEVENTTYPE_WALLETS_MPC_TSS_REQUEST_SUCCEEDED WebhookEventType = "wallets.mpc.tss_request.succeeded"
	WEBHOOKEVENTTYPE_WALLETS_ADDRESSES_CREATED WebhookEventType = "wallets.addresses.created"
	WEBHOOKEVENTTYPE_WALLETS_CREATED WebhookEventType = "wallets.created"
	WEBHOOKEVENTTYPE_MPC_VAULTS_CREATED WebhookEventType = "mpc_vaults.created"
)

// All allowed values of WebhookEventType enum
var AllowedWebhookEventTypeEnumValues = []WebhookEventType{
	"wallets.transaction.created",
	"wallets.transaction.updated",
	"wallets.transaction.failed",
	"wallets.transaction.succeeded",
	"wallets.mpc.tss_request.created",
	"wallets.mpc.tss_request.updated",
	"wallets.mpc.tss_request.failed",
	"wallets.mpc.tss_request.succeeded",
	"wallets.addresses.created",
	"wallets.created",
	"mpc_vaults.created",
}

func (v *WebhookEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookEventType(value)
	for _, existing := range AllowedWebhookEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = WebhookEventType("unknown")
    return nil
}

// NewWebhookEventTypeFromValue returns a pointer to a valid WebhookEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookEventTypeFromValue(v string) (*WebhookEventType, error) {
	ev := WebhookEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookEventType: valid values are %v", v, AllowedWebhookEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookEventType) IsValid() bool {
	for _, existing := range AllowedWebhookEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhookEventType value
func (v WebhookEventType) Ptr() *WebhookEventType {
	return &v
}

type NullableWebhookEventType struct {
	value *WebhookEventType
	isSet bool
}

func (v NullableWebhookEventType) Get() *WebhookEventType {
	return v.value
}

func (v *NullableWebhookEventType) Set(val *WebhookEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEventType(val *WebhookEventType) *NullableWebhookEventType {
	return &NullableWebhookEventType{value: val, isSet: true}
}

func (v NullableWebhookEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

