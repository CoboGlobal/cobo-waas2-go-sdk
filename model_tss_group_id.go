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
)

// checks if the TSSGroupId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TSSGroupId{}

// TSSGroupId The data for tss group id information.
type TSSGroupId struct {
	// the group id of the tss group.
	GroupId *string `json:"group_id,omitempty"`
	Curve *CurveType `json:"curve,omitempty"`
}

// NewTSSGroupId instantiates a new TSSGroupId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTSSGroupId() *TSSGroupId {
	this := TSSGroupId{}
	return &this
}

// NewTSSGroupIdWithDefaults instantiates a new TSSGroupId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTSSGroupIdWithDefaults() *TSSGroupId {
	this := TSSGroupId{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *TSSGroupId) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSGroupId) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *TSSGroupId) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *TSSGroupId) SetGroupId(v string) {
	o.GroupId = &v
}

// GetCurve returns the Curve field value if set, zero value otherwise.
func (o *TSSGroupId) GetCurve() CurveType {
	if o == nil || IsNil(o.Curve) {
		var ret CurveType
		return ret
	}
	return *o.Curve
}

// GetCurveOk returns a tuple with the Curve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSGroupId) GetCurveOk() (*CurveType, bool) {
	if o == nil || IsNil(o.Curve) {
		return nil, false
	}
	return o.Curve, true
}

// HasCurve returns a boolean if a field has been set.
func (o *TSSGroupId) HasCurve() bool {
	if o != nil && !IsNil(o.Curve) {
		return true
	}

	return false
}

// SetCurve gets a reference to the given CurveType and assigns it to the Curve field.
func (o *TSSGroupId) SetCurve(v CurveType) {
	o.Curve = &v
}

func (o TSSGroupId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TSSGroupId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !IsNil(o.Curve) {
		toSerialize["curve"] = o.Curve
	}
	return toSerialize, nil
}

type NullableTSSGroupId struct {
	value *TSSGroupId
	isSet bool
}

func (v NullableTSSGroupId) Get() *TSSGroupId {
	return v.value
}

func (v *NullableTSSGroupId) Set(val *TSSGroupId) {
	v.value = val
	v.isSet = true
}

func (v NullableTSSGroupId) IsSet() bool {
	return v.isSet
}

func (v *NullableTSSGroupId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTSSGroupId(val *TSSGroupId) *NullableTSSGroupId {
	return &NullableTSSGroupId{value: val, isSet: true}
}

func (v NullableTSSGroupId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTSSGroupId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


