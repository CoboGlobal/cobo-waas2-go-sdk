/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the ListEnableTokenPairs200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEnableTokenPairs200Response{}

// ListEnableTokenPairs200Response struct for ListEnableTokenPairs200Response
type ListEnableTokenPairs200Response struct {
	Data []SwapTokenPair `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewListEnableTokenPairs200Response instantiates a new ListEnableTokenPairs200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEnableTokenPairs200Response() *ListEnableTokenPairs200Response {
	this := ListEnableTokenPairs200Response{}
	return &this
}

// NewListEnableTokenPairs200ResponseWithDefaults instantiates a new ListEnableTokenPairs200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEnableTokenPairs200ResponseWithDefaults() *ListEnableTokenPairs200Response {
	this := ListEnableTokenPairs200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListEnableTokenPairs200Response) GetData() []SwapTokenPair {
	if o == nil || IsNil(o.Data) {
		var ret []SwapTokenPair
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEnableTokenPairs200Response) GetDataOk() ([]SwapTokenPair, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListEnableTokenPairs200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SwapTokenPair and assigns it to the Data field.
func (o *ListEnableTokenPairs200Response) SetData(v []SwapTokenPair) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListEnableTokenPairs200Response) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEnableTokenPairs200Response) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListEnableTokenPairs200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListEnableTokenPairs200Response) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListEnableTokenPairs200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEnableTokenPairs200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListEnableTokenPairs200Response struct {
	value *ListEnableTokenPairs200Response
	isSet bool
}

func (v NullableListEnableTokenPairs200Response) Get() *ListEnableTokenPairs200Response {
	return v.value
}

func (v *NullableListEnableTokenPairs200Response) Set(val *ListEnableTokenPairs200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListEnableTokenPairs200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListEnableTokenPairs200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEnableTokenPairs200Response(val *ListEnableTokenPairs200Response) *NullableListEnableTokenPairs200Response {
	return &NullableListEnableTokenPairs200Response{value: val, isSet: true}
}

func (v NullableListEnableTokenPairs200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEnableTokenPairs200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


