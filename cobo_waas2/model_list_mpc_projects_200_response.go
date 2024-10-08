/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the ListMpcProjects200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMpcProjects200Response{}

// ListMpcProjects200Response struct for ListMpcProjects200Response
type ListMpcProjects200Response struct {
	Data []MPCProject `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewListMpcProjects200Response instantiates a new ListMpcProjects200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMpcProjects200Response() *ListMpcProjects200Response {
	this := ListMpcProjects200Response{}
	return &this
}

// NewListMpcProjects200ResponseWithDefaults instantiates a new ListMpcProjects200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMpcProjects200ResponseWithDefaults() *ListMpcProjects200Response {
	this := ListMpcProjects200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListMpcProjects200Response) GetData() []MPCProject {
	if o == nil || IsNil(o.Data) {
		var ret []MPCProject
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMpcProjects200Response) GetDataOk() ([]MPCProject, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListMpcProjects200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []MPCProject and assigns it to the Data field.
func (o *ListMpcProjects200Response) SetData(v []MPCProject) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListMpcProjects200Response) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMpcProjects200Response) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListMpcProjects200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListMpcProjects200Response) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListMpcProjects200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMpcProjects200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListMpcProjects200Response struct {
	value *ListMpcProjects200Response
	isSet bool
}

func (v NullableListMpcProjects200Response) Get() *ListMpcProjects200Response {
	return v.value
}

func (v *NullableListMpcProjects200Response) Set(val *ListMpcProjects200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListMpcProjects200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListMpcProjects200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMpcProjects200Response(val *ListMpcProjects200Response) *NullableListMpcProjects200Response {
	return &NullableListMpcProjects200Response{value: val, isSet: true}
}

func (v NullableListMpcProjects200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMpcProjects200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


