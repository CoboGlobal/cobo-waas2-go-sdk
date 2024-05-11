/*
Cobo Wallet as a Service 2.0

Cobo WaaS 2.0 enables you to programmatically access Cobo's full suite of crypto wallet technologies with powerful and flexible access controls.  # Wallet technologies - Custodial Wallet - MPC Wallet - Smart Contract Wallet (Based on Safe{Wallet}) - Exchange Wallet  # Risk Control technologies - Workflow - Access Control List (ACL)  # Risk Control targets - Wallet Management   - User/team and their permission management   - Risk control configurations, e.g. whitelist, blacklist, rate-limiting etc. - Blockchain Interaction   - Crypto transfer   - Smart Contract Invocation  # Important HTTPS only. RESTful, resource oriented  # Get Started Set up your APIs or get authorization  # Authentication and Authorization CoboAuth  # Request and Response application/json  # Error Handling  ### Common error codes | Error Code | Description | | -- | -- |  ### API-specific error codes For error codes that are dedicated to a specific API, see the Error codes section in each API specification, for example, /v3/wallets.  # Rate and Usage Limiting  # Idempotent Request  # Pagination # Support [Developer Hub](https://cobo.com/developers) 

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
	"fmt"
)

// CurveType the model 'CurveType'
type CurveType string

// List of CurveType
const (
	CURVETYPE_SECP256_K1 CurveType = "SECP256K1"
	CURVETYPE_ED25519 CurveType = "ED25519"
)

// All allowed values of CurveType enum
var AllowedCurveTypeEnumValues = []CurveType{
	"SECP256K1",
	"ED25519",
}

func (v *CurveType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CurveType(value)
	for _, existing := range AllowedCurveTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CurveType", value)
}

// NewCurveTypeFromValue returns a pointer to a valid CurveType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCurveTypeFromValue(v string) (*CurveType, error) {
	ev := CurveType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CurveType: valid values are %v", v, AllowedCurveTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CurveType) IsValid() bool {
	for _, existing := range AllowedCurveTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CurveType value
func (v CurveType) Ptr() *CurveType {
	return &v
}

type NullableCurveType struct {
	value *CurveType
	isSet bool
}

func (v NullableCurveType) Get() *CurveType {
	return v.value
}

func (v *NullableCurveType) Set(val *CurveType) {
	v.value = val
	v.isSet = true
}

func (v NullableCurveType) IsSet() bool {
	return v.isSet
}

func (v *NullableCurveType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurveType(val *CurveType) *NullableCurveType {
	return &NullableCurveType{value: val, isSet: true}
}

func (v NullableCurveType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurveType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

