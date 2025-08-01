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

// checks if the EvmContractCallDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvmContractCallDestination{}

// EvmContractCallDestination The information about the transaction destination. Refer to [Transaction sources and destinations](https://www.cobo.com/developers/v2/guides/transactions/sources-and-destinations) for a detailed introduction about the supported sources and destinations for each transaction type.
type EvmContractCallDestination struct {
	DestinationType ContractCallDestinationType `json:"destination_type"`
	// The destination address.
	Address string `json:"address"`
	// The transfer amount. For example, if you trade 1.5 ETH, then the value is `1.5`. 
	Value *string `json:"value,omitempty"`
	// The data used to invoke a specific function or method within the specified contract at the destination address, with a maximum length of 65,000 characters. 
	Calldata string `json:"calldata"`
}

type _EvmContractCallDestination EvmContractCallDestination

// NewEvmContractCallDestination instantiates a new EvmContractCallDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvmContractCallDestination(destinationType ContractCallDestinationType, address string, calldata string) *EvmContractCallDestination {
	this := EvmContractCallDestination{}
	this.DestinationType = destinationType
	this.Address = address
	this.Calldata = calldata
	return &this
}

// NewEvmContractCallDestinationWithDefaults instantiates a new EvmContractCallDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvmContractCallDestinationWithDefaults() *EvmContractCallDestination {
	this := EvmContractCallDestination{}
	return &this
}

// GetDestinationType returns the DestinationType field value
func (o *EvmContractCallDestination) GetDestinationType() ContractCallDestinationType {
	if o == nil {
		var ret ContractCallDestinationType
		return ret
	}

	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *EvmContractCallDestination) GetDestinationTypeOk() (*ContractCallDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value
func (o *EvmContractCallDestination) SetDestinationType(v ContractCallDestinationType) {
	o.DestinationType = v
}

// GetAddress returns the Address field value
func (o *EvmContractCallDestination) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *EvmContractCallDestination) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *EvmContractCallDestination) SetAddress(v string) {
	o.Address = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EvmContractCallDestination) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmContractCallDestination) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EvmContractCallDestination) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EvmContractCallDestination) SetValue(v string) {
	o.Value = &v
}

// GetCalldata returns the Calldata field value
func (o *EvmContractCallDestination) GetCalldata() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Calldata
}

// GetCalldataOk returns a tuple with the Calldata field value
// and a boolean to check if the value has been set.
func (o *EvmContractCallDestination) GetCalldataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Calldata, true
}

// SetCalldata sets field value
func (o *EvmContractCallDestination) SetCalldata(v string) {
	o.Calldata = v
}

func (o EvmContractCallDestination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvmContractCallDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination_type"] = o.DestinationType
	toSerialize["address"] = o.Address
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	toSerialize["calldata"] = o.Calldata
	return toSerialize, nil
}

func (o *EvmContractCallDestination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination_type",
		"address",
		"calldata",
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

	varEvmContractCallDestination := _EvmContractCallDestination{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEvmContractCallDestination)

	if err != nil {
		return err
	}

	*o = EvmContractCallDestination(varEvmContractCallDestination)

	return err
}

type NullableEvmContractCallDestination struct {
	value *EvmContractCallDestination
	isSet bool
}

func (v NullableEvmContractCallDestination) Get() *EvmContractCallDestination {
	return v.value
}

func (v *NullableEvmContractCallDestination) Set(val *EvmContractCallDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableEvmContractCallDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableEvmContractCallDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvmContractCallDestination(val *EvmContractCallDestination) *NullableEvmContractCallDestination {
	return &NullableEvmContractCallDestination{value: val, isSet: true}
}

func (v NullableEvmContractCallDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvmContractCallDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


