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

// checks if the AddressesEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressesEventData{}

// AddressesEventData struct for AddressesEventData
type AddressesEventData struct {
	//  The data type of the event. - `Transaction`: The transaction event data. - `TSSRequest`: The TSS request event data. - `Addresses`: The addresses event data. - `WalletInfo`: The wallet information event data. - `MPCVault`: The MPC vault event data. - `Chains`: The enabled chain event data. - `Tokens`: The enabled token event data. - `TokenListing`: The token listing event data.        - `PaymentOrder`: The payment order event data. - `PaymentRefund`: The payment refund event data. - `PaymentSettlement`: The payment settlement event data. - `PaymentTransaction`: The payment transaction event data. - `PaymentAddressUpdate`: The payment address update event data. - `BalanceUpdateInfo`: The balance update event data. - `SuspendedToken`: The token suspension event data.
	DataType string `json:"data_type"`
	// A list of addresses.
	Addresses []AddressesEventDataAllOfAddresses `json:"addresses,omitempty"`
}

type _AddressesEventData AddressesEventData

// NewAddressesEventData instantiates a new AddressesEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressesEventData(dataType string) *AddressesEventData {
	this := AddressesEventData{}
	this.DataType = dataType
	return &this
}

// NewAddressesEventDataWithDefaults instantiates a new AddressesEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressesEventDataWithDefaults() *AddressesEventData {
	this := AddressesEventData{}
	return &this
}

// GetDataType returns the DataType field value
func (o *AddressesEventData) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *AddressesEventData) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *AddressesEventData) SetDataType(v string) {
	o.DataType = v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *AddressesEventData) GetAddresses() []AddressesEventDataAllOfAddresses {
	if o == nil || IsNil(o.Addresses) {
		var ret []AddressesEventDataAllOfAddresses
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressesEventData) GetAddressesOk() ([]AddressesEventDataAllOfAddresses, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *AddressesEventData) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []AddressesEventDataAllOfAddresses and assigns it to the Addresses field.
func (o *AddressesEventData) SetAddresses(v []AddressesEventDataAllOfAddresses) {
	o.Addresses = v
}

func (o AddressesEventData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressesEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data_type"] = o.DataType
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	return toSerialize, nil
}

func (o *AddressesEventData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data_type",
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

	varAddressesEventData := _AddressesEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressesEventData)

	if err != nil {
		return err
	}

	*o = AddressesEventData(varAddressesEventData)

	return err
}

type NullableAddressesEventData struct {
	value *AddressesEventData
	isSet bool
}

func (v NullableAddressesEventData) Get() *AddressesEventData {
	return v.value
}

func (v *NullableAddressesEventData) Set(val *AddressesEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressesEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressesEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressesEventData(val *AddressesEventData) *NullableAddressesEventData {
	return &NullableAddressesEventData{value: val, isSet: true}
}

func (v NullableAddressesEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressesEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


