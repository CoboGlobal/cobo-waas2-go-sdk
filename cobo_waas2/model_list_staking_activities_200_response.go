/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the ListStakingActivities200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListStakingActivities200Response{}

// ListStakingActivities200Response struct for ListStakingActivities200Response
type ListStakingActivities200Response struct {
	Data []Activity `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewListStakingActivities200Response instantiates a new ListStakingActivities200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStakingActivities200Response() *ListStakingActivities200Response {
	this := ListStakingActivities200Response{}
	return &this
}

// NewListStakingActivities200ResponseWithDefaults instantiates a new ListStakingActivities200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStakingActivities200ResponseWithDefaults() *ListStakingActivities200Response {
	this := ListStakingActivities200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListStakingActivities200Response) GetData() []Activity {
	if o == nil || IsNil(o.Data) {
		var ret []Activity
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStakingActivities200Response) GetDataOk() ([]Activity, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListStakingActivities200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Activity and assigns it to the Data field.
func (o *ListStakingActivities200Response) SetData(v []Activity) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListStakingActivities200Response) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStakingActivities200Response) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListStakingActivities200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListStakingActivities200Response) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListStakingActivities200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListStakingActivities200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListStakingActivities200Response struct {
	value *ListStakingActivities200Response
	isSet bool
}

func (v NullableListStakingActivities200Response) Get() *ListStakingActivities200Response {
	return v.value
}

func (v *NullableListStakingActivities200Response) Set(val *ListStakingActivities200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListStakingActivities200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListStakingActivities200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStakingActivities200Response(val *ListStakingActivities200Response) *NullableListStakingActivities200Response {
	return &NullableListStakingActivities200Response{value: val, isSet: true}
}

func (v NullableListStakingActivities200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStakingActivities200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


