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
	"bytes"
	"fmt"
)

// checks if the UpdateMpcProjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMpcProjectRequest{}

// UpdateMpcProjectRequest struct for UpdateMpcProjectRequest
type UpdateMpcProjectRequest struct {
	// The name of the mpc project.
	Name string `json:"name"`
}

type _UpdateMpcProjectRequest UpdateMpcProjectRequest

// NewUpdateMpcProjectRequest instantiates a new UpdateMpcProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMpcProjectRequest(name string) *UpdateMpcProjectRequest {
	this := UpdateMpcProjectRequest{}
	this.Name = name
	return &this
}

// NewUpdateMpcProjectRequestWithDefaults instantiates a new UpdateMpcProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMpcProjectRequestWithDefaults() *UpdateMpcProjectRequest {
	this := UpdateMpcProjectRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateMpcProjectRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateMpcProjectRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateMpcProjectRequest) SetName(v string) {
	o.Name = v
}

func (o UpdateMpcProjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMpcProjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *UpdateMpcProjectRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateMpcProjectRequest := _UpdateMpcProjectRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateMpcProjectRequest)

	if err != nil {
		return err
	}

	*o = UpdateMpcProjectRequest(varUpdateMpcProjectRequest)

	return err
}

type NullableUpdateMpcProjectRequest struct {
	value *UpdateMpcProjectRequest
	isSet bool
}

func (v NullableUpdateMpcProjectRequest) Get() *UpdateMpcProjectRequest {
	return v.value
}

func (v *NullableUpdateMpcProjectRequest) Set(val *UpdateMpcProjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMpcProjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMpcProjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMpcProjectRequest(val *UpdateMpcProjectRequest) *NullableUpdateMpcProjectRequest {
	return &NullableUpdateMpcProjectRequest{value: val, isSet: true}
}

func (v NullableUpdateMpcProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMpcProjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


