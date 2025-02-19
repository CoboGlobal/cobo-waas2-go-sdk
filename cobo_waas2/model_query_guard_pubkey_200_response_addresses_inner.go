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

// checks if the QueryGuardPubkey200ResponseAddressesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryGuardPubkey200ResponseAddressesInner{}

// QueryGuardPubkey200ResponseAddressesInner struct for QueryGuardPubkey200ResponseAddressesInner
type QueryGuardPubkey200ResponseAddressesInner struct {
	// The wallet address.
	Address string `json:"address"`
	// The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](/v2/api-references/wallets/list-enabled-chains).
	ChainId string `json:"chain_id"`
}

type _QueryGuardPubkey200ResponseAddressesInner QueryGuardPubkey200ResponseAddressesInner

// NewQueryGuardPubkey200ResponseAddressesInner instantiates a new QueryGuardPubkey200ResponseAddressesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryGuardPubkey200ResponseAddressesInner(address string, chainId string) *QueryGuardPubkey200ResponseAddressesInner {
	this := QueryGuardPubkey200ResponseAddressesInner{}
	this.Address = address
	this.ChainId = chainId
	return &this
}

// NewQueryGuardPubkey200ResponseAddressesInnerWithDefaults instantiates a new QueryGuardPubkey200ResponseAddressesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryGuardPubkey200ResponseAddressesInnerWithDefaults() *QueryGuardPubkey200ResponseAddressesInner {
	this := QueryGuardPubkey200ResponseAddressesInner{}
	return &this
}

// GetAddress returns the Address field value
func (o *QueryGuardPubkey200ResponseAddressesInner) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *QueryGuardPubkey200ResponseAddressesInner) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *QueryGuardPubkey200ResponseAddressesInner) SetAddress(v string) {
	o.Address = v
}

// GetChainId returns the ChainId field value
func (o *QueryGuardPubkey200ResponseAddressesInner) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *QueryGuardPubkey200ResponseAddressesInner) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *QueryGuardPubkey200ResponseAddressesInner) SetChainId(v string) {
	o.ChainId = v
}

func (o QueryGuardPubkey200ResponseAddressesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryGuardPubkey200ResponseAddressesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["chain_id"] = o.ChainId
	return toSerialize, nil
}

func (o *QueryGuardPubkey200ResponseAddressesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"chain_id",
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

	varQueryGuardPubkey200ResponseAddressesInner := _QueryGuardPubkey200ResponseAddressesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryGuardPubkey200ResponseAddressesInner)

	if err != nil {
		return err
	}

	*o = QueryGuardPubkey200ResponseAddressesInner(varQueryGuardPubkey200ResponseAddressesInner)

	return err
}

type NullableQueryGuardPubkey200ResponseAddressesInner struct {
	value *QueryGuardPubkey200ResponseAddressesInner
	isSet bool
}

func (v NullableQueryGuardPubkey200ResponseAddressesInner) Get() *QueryGuardPubkey200ResponseAddressesInner {
	return v.value
}

func (v *NullableQueryGuardPubkey200ResponseAddressesInner) Set(val *QueryGuardPubkey200ResponseAddressesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryGuardPubkey200ResponseAddressesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryGuardPubkey200ResponseAddressesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryGuardPubkey200ResponseAddressesInner(val *QueryGuardPubkey200ResponseAddressesInner) *NullableQueryGuardPubkey200ResponseAddressesInner {
	return &NullableQueryGuardPubkey200ResponseAddressesInner{value: val, isSet: true}
}

func (v NullableQueryGuardPubkey200ResponseAddressesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryGuardPubkey200ResponseAddressesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


