/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the ListAddressBalancesByToken200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAddressBalancesByToken200Response{}

// ListAddressBalancesByToken200Response struct for ListAddressBalancesByToken200Response
type ListAddressBalancesByToken200Response struct {
	Data []AddressBalance `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewListAddressBalancesByToken200Response instantiates a new ListAddressBalancesByToken200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAddressBalancesByToken200Response() *ListAddressBalancesByToken200Response {
	this := ListAddressBalancesByToken200Response{}
	return &this
}

// NewListAddressBalancesByToken200ResponseWithDefaults instantiates a new ListAddressBalancesByToken200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAddressBalancesByToken200ResponseWithDefaults() *ListAddressBalancesByToken200Response {
	this := ListAddressBalancesByToken200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListAddressBalancesByToken200Response) GetData() []AddressBalance {
	if o == nil || IsNil(o.Data) {
		var ret []AddressBalance
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAddressBalancesByToken200Response) GetDataOk() ([]AddressBalance, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListAddressBalancesByToken200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AddressBalance and assigns it to the Data field.
func (o *ListAddressBalancesByToken200Response) SetData(v []AddressBalance) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListAddressBalancesByToken200Response) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAddressBalancesByToken200Response) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListAddressBalancesByToken200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListAddressBalancesByToken200Response) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListAddressBalancesByToken200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAddressBalancesByToken200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListAddressBalancesByToken200Response struct {
	value *ListAddressBalancesByToken200Response
	isSet bool
}

func (v NullableListAddressBalancesByToken200Response) Get() *ListAddressBalancesByToken200Response {
	return v.value
}

func (v *NullableListAddressBalancesByToken200Response) Set(val *ListAddressBalancesByToken200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListAddressBalancesByToken200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListAddressBalancesByToken200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAddressBalancesByToken200Response(val *ListAddressBalancesByToken200Response) *NullableListAddressBalancesByToken200Response {
	return &NullableListAddressBalancesByToken200Response{value: val, isSet: true}
}

func (v NullableListAddressBalancesByToken200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAddressBalancesByToken200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


