/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the ListTokenBalancesForFeeStation200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTokenBalancesForFeeStation200Response{}

// ListTokenBalancesForFeeStation200Response struct for ListTokenBalancesForFeeStation200Response
type ListTokenBalancesForFeeStation200Response struct {
	Data []ListTokenBalancesForFeeStation200ResponseDataInner `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewListTokenBalancesForFeeStation200Response instantiates a new ListTokenBalancesForFeeStation200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTokenBalancesForFeeStation200Response() *ListTokenBalancesForFeeStation200Response {
	this := ListTokenBalancesForFeeStation200Response{}
	return &this
}

// NewListTokenBalancesForFeeStation200ResponseWithDefaults instantiates a new ListTokenBalancesForFeeStation200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTokenBalancesForFeeStation200ResponseWithDefaults() *ListTokenBalancesForFeeStation200Response {
	this := ListTokenBalancesForFeeStation200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListTokenBalancesForFeeStation200Response) GetData() []ListTokenBalancesForFeeStation200ResponseDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []ListTokenBalancesForFeeStation200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTokenBalancesForFeeStation200Response) GetDataOk() ([]ListTokenBalancesForFeeStation200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListTokenBalancesForFeeStation200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ListTokenBalancesForFeeStation200ResponseDataInner and assigns it to the Data field.
func (o *ListTokenBalancesForFeeStation200Response) SetData(v []ListTokenBalancesForFeeStation200ResponseDataInner) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListTokenBalancesForFeeStation200Response) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTokenBalancesForFeeStation200Response) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListTokenBalancesForFeeStation200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListTokenBalancesForFeeStation200Response) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListTokenBalancesForFeeStation200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTokenBalancesForFeeStation200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListTokenBalancesForFeeStation200Response struct {
	value *ListTokenBalancesForFeeStation200Response
	isSet bool
}

func (v NullableListTokenBalancesForFeeStation200Response) Get() *ListTokenBalancesForFeeStation200Response {
	return v.value
}

func (v *NullableListTokenBalancesForFeeStation200Response) Set(val *ListTokenBalancesForFeeStation200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListTokenBalancesForFeeStation200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListTokenBalancesForFeeStation200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTokenBalancesForFeeStation200Response(val *ListTokenBalancesForFeeStation200Response) *NullableListTokenBalancesForFeeStation200Response {
	return &NullableListTokenBalancesForFeeStation200Response{value: val, isSet: true}
}

func (v NullableListTokenBalancesForFeeStation200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTokenBalancesForFeeStation200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


