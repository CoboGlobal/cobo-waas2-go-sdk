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

// checks if the GetExchangeSupportedAssets200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetExchangeSupportedAssets200Response{}

// GetExchangeSupportedAssets200Response struct for GetExchangeSupportedAssets200Response
type GetExchangeSupportedAssets200Response struct {
	Data []AssetInfo `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewGetExchangeSupportedAssets200Response instantiates a new GetExchangeSupportedAssets200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchangeSupportedAssets200Response() *GetExchangeSupportedAssets200Response {
	this := GetExchangeSupportedAssets200Response{}
	return &this
}

// NewGetExchangeSupportedAssets200ResponseWithDefaults instantiates a new GetExchangeSupportedAssets200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchangeSupportedAssets200ResponseWithDefaults() *GetExchangeSupportedAssets200Response {
	this := GetExchangeSupportedAssets200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetExchangeSupportedAssets200Response) GetData() []AssetInfo {
	if o == nil || IsNil(o.Data) {
		var ret []AssetInfo
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeSupportedAssets200Response) GetDataOk() ([]AssetInfo, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetExchangeSupportedAssets200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AssetInfo and assigns it to the Data field.
func (o *GetExchangeSupportedAssets200Response) SetData(v []AssetInfo) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GetExchangeSupportedAssets200Response) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeSupportedAssets200Response) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GetExchangeSupportedAssets200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *GetExchangeSupportedAssets200Response) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o GetExchangeSupportedAssets200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExchangeSupportedAssets200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableGetExchangeSupportedAssets200Response struct {
	value *GetExchangeSupportedAssets200Response
	isSet bool
}

func (v NullableGetExchangeSupportedAssets200Response) Get() *GetExchangeSupportedAssets200Response {
	return v.value
}

func (v *NullableGetExchangeSupportedAssets200Response) Set(val *GetExchangeSupportedAssets200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeSupportedAssets200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeSupportedAssets200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeSupportedAssets200Response(val *GetExchangeSupportedAssets200Response) *NullableGetExchangeSupportedAssets200Response {
	return &NullableGetExchangeSupportedAssets200Response{value: val, isSet: true}
}

func (v NullableGetExchangeSupportedAssets200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeSupportedAssets200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

