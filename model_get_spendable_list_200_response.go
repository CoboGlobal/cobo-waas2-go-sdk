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

// checks if the GetSpendableList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSpendableList200Response{}

// GetSpendableList200Response struct for GetSpendableList200Response
type GetSpendableList200Response struct {
	Data []UTXO `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewGetSpendableList200Response instantiates a new GetSpendableList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSpendableList200Response() *GetSpendableList200Response {
	this := GetSpendableList200Response{}
	return &this
}

// NewGetSpendableList200ResponseWithDefaults instantiates a new GetSpendableList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSpendableList200ResponseWithDefaults() *GetSpendableList200Response {
	this := GetSpendableList200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetSpendableList200Response) GetData() []UTXO {
	if o == nil || IsNil(o.Data) {
		var ret []UTXO
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpendableList200Response) GetDataOk() ([]UTXO, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetSpendableList200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []UTXO and assigns it to the Data field.
func (o *GetSpendableList200Response) SetData(v []UTXO) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GetSpendableList200Response) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpendableList200Response) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GetSpendableList200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *GetSpendableList200Response) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o GetSpendableList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSpendableList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableGetSpendableList200Response struct {
	value *GetSpendableList200Response
	isSet bool
}

func (v NullableGetSpendableList200Response) Get() *GetSpendableList200Response {
	return v.value
}

func (v *NullableGetSpendableList200Response) Set(val *GetSpendableList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSpendableList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSpendableList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSpendableList200Response(val *GetSpendableList200Response) *NullableGetSpendableList200Response {
	return &NullableGetSpendableList200Response{value: val, isSet: true}
}

func (v NullableGetSpendableList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSpendableList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

