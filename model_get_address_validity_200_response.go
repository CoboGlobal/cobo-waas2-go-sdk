/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
)

// checks if the GetAddressValidity200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAddressValidity200Response{}

// GetAddressValidity200Response struct for GetAddressValidity200Response
type GetAddressValidity200Response struct {
	// Whether the address is valid. - `true`: The address is valid. - `false`: the address is invalid. 
	Validity *bool `json:"validity,omitempty"`
}

// NewGetAddressValidity200Response instantiates a new GetAddressValidity200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAddressValidity200Response() *GetAddressValidity200Response {
	this := GetAddressValidity200Response{}
	return &this
}

// NewGetAddressValidity200ResponseWithDefaults instantiates a new GetAddressValidity200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAddressValidity200ResponseWithDefaults() *GetAddressValidity200Response {
	this := GetAddressValidity200Response{}
	return &this
}

// GetValidity returns the Validity field value if set, zero value otherwise.
func (o *GetAddressValidity200Response) GetValidity() bool {
	if o == nil || IsNil(o.Validity) {
		var ret bool
		return ret
	}
	return *o.Validity
}

// GetValidityOk returns a tuple with the Validity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAddressValidity200Response) GetValidityOk() (*bool, bool) {
	if o == nil || IsNil(o.Validity) {
		return nil, false
	}
	return o.Validity, true
}

// HasValidity returns a boolean if a field has been set.
func (o *GetAddressValidity200Response) HasValidity() bool {
	if o != nil && !IsNil(o.Validity) {
		return true
	}

	return false
}

// SetValidity gets a reference to the given bool and assigns it to the Validity field.
func (o *GetAddressValidity200Response) SetValidity(v bool) {
	o.Validity = &v
}

func (o GetAddressValidity200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAddressValidity200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Validity) {
		toSerialize["validity"] = o.Validity
	}
	return toSerialize, nil
}

type NullableGetAddressValidity200Response struct {
	value *GetAddressValidity200Response
	isSet bool
}

func (v NullableGetAddressValidity200Response) Get() *GetAddressValidity200Response {
	return v.value
}

func (v *NullableGetAddressValidity200Response) Set(val *GetAddressValidity200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAddressValidity200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAddressValidity200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAddressValidity200Response(val *GetAddressValidity200Response) *NullableGetAddressValidity200Response {
	return &NullableGetAddressValidity200Response{value: val, isSet: true}
}

func (v NullableGetAddressValidity200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAddressValidity200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


