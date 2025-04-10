/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the ListAddressBooks200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAddressBooks200Response{}

// ListAddressBooks200Response struct for ListAddressBooks200Response
type ListAddressBooks200Response struct {
	Data []AddressBook `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewListAddressBooks200Response instantiates a new ListAddressBooks200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAddressBooks200Response() *ListAddressBooks200Response {
	this := ListAddressBooks200Response{}
	return &this
}

// NewListAddressBooks200ResponseWithDefaults instantiates a new ListAddressBooks200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAddressBooks200ResponseWithDefaults() *ListAddressBooks200Response {
	this := ListAddressBooks200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListAddressBooks200Response) GetData() []AddressBook {
	if o == nil || IsNil(o.Data) {
		var ret []AddressBook
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAddressBooks200Response) GetDataOk() ([]AddressBook, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListAddressBooks200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AddressBook and assigns it to the Data field.
func (o *ListAddressBooks200Response) SetData(v []AddressBook) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListAddressBooks200Response) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAddressBooks200Response) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListAddressBooks200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListAddressBooks200Response) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListAddressBooks200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAddressBooks200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListAddressBooks200Response struct {
	value *ListAddressBooks200Response
	isSet bool
}

func (v NullableListAddressBooks200Response) Get() *ListAddressBooks200Response {
	return v.value
}

func (v *NullableListAddressBooks200Response) Set(val *ListAddressBooks200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListAddressBooks200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListAddressBooks200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAddressBooks200Response(val *ListAddressBooks200Response) *NullableListAddressBooks200Response {
	return &NullableListAddressBooks200Response{value: val, isSet: true}
}

func (v NullableListAddressBooks200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAddressBooks200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


