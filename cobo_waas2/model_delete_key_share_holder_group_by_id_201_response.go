/*
Cobo Wallet as a Service 2.0

Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DeleteKeyShareHolderGroupById201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteKeyShareHolderGroupById201Response{}

// DeleteKeyShareHolderGroupById201Response struct for DeleteKeyShareHolderGroupById201Response
type DeleteKeyShareHolderGroupById201Response struct {
	// Whether the request to delete the key share holder group has been successfully submitted. - `true`: The request to delete the key share holder group has been successfully submitted. - `false`: The request to delete the key share holder group has not been submitted. 
	Submitted bool `json:"submitted"`
}

type _DeleteKeyShareHolderGroupById201Response DeleteKeyShareHolderGroupById201Response

// NewDeleteKeyShareHolderGroupById201Response instantiates a new DeleteKeyShareHolderGroupById201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteKeyShareHolderGroupById201Response(submitted bool) *DeleteKeyShareHolderGroupById201Response {
	this := DeleteKeyShareHolderGroupById201Response{}
	this.Submitted = submitted
	return &this
}

// NewDeleteKeyShareHolderGroupById201ResponseWithDefaults instantiates a new DeleteKeyShareHolderGroupById201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteKeyShareHolderGroupById201ResponseWithDefaults() *DeleteKeyShareHolderGroupById201Response {
	this := DeleteKeyShareHolderGroupById201Response{}
	return &this
}

// GetSubmitted returns the Submitted field value
func (o *DeleteKeyShareHolderGroupById201Response) GetSubmitted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Submitted
}

// GetSubmittedOk returns a tuple with the Submitted field value
// and a boolean to check if the value has been set.
func (o *DeleteKeyShareHolderGroupById201Response) GetSubmittedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Submitted, true
}

// SetSubmitted sets field value
func (o *DeleteKeyShareHolderGroupById201Response) SetSubmitted(v bool) {
	o.Submitted = v
}

func (o DeleteKeyShareHolderGroupById201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteKeyShareHolderGroupById201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["submitted"] = o.Submitted
	return toSerialize, nil
}

func (o *DeleteKeyShareHolderGroupById201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"submitted",
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

	varDeleteKeyShareHolderGroupById201Response := _DeleteKeyShareHolderGroupById201Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteKeyShareHolderGroupById201Response)

	if err != nil {
		return err
	}

	*o = DeleteKeyShareHolderGroupById201Response(varDeleteKeyShareHolderGroupById201Response)

	return err
}

type NullableDeleteKeyShareHolderGroupById201Response struct {
	value *DeleteKeyShareHolderGroupById201Response
	isSet bool
}

func (v NullableDeleteKeyShareHolderGroupById201Response) Get() *DeleteKeyShareHolderGroupById201Response {
	return v.value
}

func (v *NullableDeleteKeyShareHolderGroupById201Response) Set(val *DeleteKeyShareHolderGroupById201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteKeyShareHolderGroupById201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteKeyShareHolderGroupById201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteKeyShareHolderGroupById201Response(val *DeleteKeyShareHolderGroupById201Response) *NullableDeleteKeyShareHolderGroupById201Response {
	return &NullableDeleteKeyShareHolderGroupById201Response{value: val, isSet: true}
}

func (v NullableDeleteKeyShareHolderGroupById201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteKeyShareHolderGroupById201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


