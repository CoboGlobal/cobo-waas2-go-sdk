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

// UpdateGroupAction The available actions of key share holder group update. Possible values include: - `UpgradeToMainGroup`: This upgrades an active [Signing Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups) to the [Main Group](https://manuals.cobo.com/en/portal/mpc-wallets/ocw/create-key-share-groups). The original Main Group will be permanently deleted. 
type UpdateGroupAction string

// List of UpdateGroupAction
const (
	UPDATEGROUPACTION_UPGRADE_TO_MAIN_GROUP UpdateGroupAction = "UpgradeToMainGroup"
)

// All allowed values of UpdateGroupAction enum
var AllowedUpdateGroupActionEnumValues = []UpdateGroupAction{
	"UpgradeToMainGroup",
}

func (v *UpdateGroupAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UpdateGroupAction(value)
	for _, existing := range AllowedUpdateGroupActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = UpdateGroupAction("unknown")
	return nil
}

// NewUpdateGroupActionFromValue returns a pointer to a valid UpdateGroupAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUpdateGroupActionFromValue(v string) (*UpdateGroupAction, error) {
	ev := UpdateGroupAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UpdateGroupAction: valid values are %v", v, AllowedUpdateGroupActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UpdateGroupAction) IsValid() bool {
	for _, existing := range AllowedUpdateGroupActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateGroupAction value
func (v UpdateGroupAction) Ptr() *UpdateGroupAction {
	return &v
}

type NullableUpdateGroupAction struct {
	value *UpdateGroupAction
	isSet bool
}

func (v NullableUpdateGroupAction) Get() *UpdateGroupAction {
	return v.value
}

func (v *NullableUpdateGroupAction) Set(val *UpdateGroupAction) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGroupAction) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGroupAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGroupAction(val *UpdateGroupAction) *NullableUpdateGroupAction {
	return &NullableUpdateGroupAction{value: val, isSet: true}
}

func (v NullableUpdateGroupAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGroupAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

