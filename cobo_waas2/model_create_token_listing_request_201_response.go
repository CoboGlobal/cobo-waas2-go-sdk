/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateTokenListingRequest201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTokenListingRequest201Response{}

// CreateTokenListingRequest201Response struct for CreateTokenListingRequest201Response
type CreateTokenListingRequest201Response struct {
	// The unique identifier for the token listing request. You can use it to track the request status with the [Get token listing request](https://www.cobo.com/developers/v2/api-references/wallets/get-token-listing-request) operation.
	RequestId string `json:"request_id"`
}

type _CreateTokenListingRequest201Response CreateTokenListingRequest201Response

// NewCreateTokenListingRequest201Response instantiates a new CreateTokenListingRequest201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTokenListingRequest201Response(requestId string) *CreateTokenListingRequest201Response {
	this := CreateTokenListingRequest201Response{}
	this.RequestId = requestId
	return &this
}

// NewCreateTokenListingRequest201ResponseWithDefaults instantiates a new CreateTokenListingRequest201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTokenListingRequest201ResponseWithDefaults() *CreateTokenListingRequest201Response {
	this := CreateTokenListingRequest201Response{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *CreateTokenListingRequest201Response) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreateTokenListingRequest201Response) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreateTokenListingRequest201Response) SetRequestId(v string) {
	o.RequestId = v
}

func (o CreateTokenListingRequest201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTokenListingRequest201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	return toSerialize, nil
}

func (o *CreateTokenListingRequest201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request_id",
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

	varCreateTokenListingRequest201Response := _CreateTokenListingRequest201Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTokenListingRequest201Response)

	if err != nil {
		return err
	}

	*o = CreateTokenListingRequest201Response(varCreateTokenListingRequest201Response)

	return err
}

type NullableCreateTokenListingRequest201Response struct {
	value *CreateTokenListingRequest201Response
	isSet bool
}

func (v NullableCreateTokenListingRequest201Response) Get() *CreateTokenListingRequest201Response {
	return v.value
}

func (v *NullableCreateTokenListingRequest201Response) Set(val *CreateTokenListingRequest201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTokenListingRequest201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTokenListingRequest201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTokenListingRequest201Response(val *CreateTokenListingRequest201Response) *NullableCreateTokenListingRequest201Response {
	return &NullableCreateTokenListingRequest201Response{value: val, isSet: true}
}

func (v NullableCreateTokenListingRequest201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTokenListingRequest201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


