/*
Cobo Wallet as a Service 2.0

Cobo WaaS 2.0 enables you to programmatically access Cobo's full suite of crypto wallet technologies with powerful and flexible access controls.  # Wallet technologies - Custodial Wallet - MPC Wallet - Smart Contract Wallet (Based on Safe{Wallet}) - Exchange Wallet  # Risk Control technologies - Workflow - Access Control List (ACL)  # Risk Control targets - Wallet Management   - User/team and their permission management   - Risk control configurations, e.g. whitelist, blacklist, rate-limiting etc. - Blockchain Interaction   - Crypto transfer   - Smart Contract Invocation  # Important HTTPS only. RESTful, resource oriented  # Get Started Set up your APIs or get authorization  # Authentication and Authorization CoboAuth  # Request and Response application/json  # Error Handling  ### Common error codes | Error Code | Description | | -- | -- |  ### API-specific error codes For error codes that are dedicated to a specific API, see the Error codes section in each API specification, for example, /v3/wallets.  # Rate and Usage Limiting  # Idempotent Request  # Pagination # Support [Developer Hub](https://cobo.com/developers) 

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
)

// checks if the ListAddresses200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAddresses200Response{}

// ListAddresses200Response struct for ListAddresses200Response
type ListAddresses200Response struct {
	Data []AddressInfo `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewListAddresses200Response instantiates a new ListAddresses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAddresses200Response() *ListAddresses200Response {
	this := ListAddresses200Response{}
	return &this
}

// NewListAddresses200ResponseWithDefaults instantiates a new ListAddresses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAddresses200ResponseWithDefaults() *ListAddresses200Response {
	this := ListAddresses200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListAddresses200Response) GetData() []AddressInfo {
	if o == nil || IsNil(o.Data) {
		var ret []AddressInfo
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAddresses200Response) GetDataOk() ([]AddressInfo, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListAddresses200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AddressInfo and assigns it to the Data field.
func (o *ListAddresses200Response) SetData(v []AddressInfo) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListAddresses200Response) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAddresses200Response) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListAddresses200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListAddresses200Response) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListAddresses200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAddresses200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListAddresses200Response struct {
	value *ListAddresses200Response
	isSet bool
}

func (v NullableListAddresses200Response) Get() *ListAddresses200Response {
	return v.value
}

func (v *NullableListAddresses200Response) Set(val *ListAddresses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListAddresses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListAddresses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAddresses200Response(val *ListAddresses200Response) *NullableListAddresses200Response {
	return &NullableListAddresses200Response{value: val, isSet: true}
}

func (v NullableListAddresses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAddresses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


