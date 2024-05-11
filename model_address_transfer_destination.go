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
	"bytes"
	"fmt"
)

// checks if the AddressTransferDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressTransferDestination{}

// AddressTransferDestination The data for address destination.
type AddressTransferDestination struct {
	DestinationType TransferDestinationType `json:"destination_type"`
	// Destination address
	AddressStr string `json:"address_str"`
	// Destination address memo
	Memo *string `json:"memo,omitempty"`
}

type _AddressTransferDestination AddressTransferDestination

// NewAddressTransferDestination instantiates a new AddressTransferDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransferDestination(destinationType TransferDestinationType, addressStr string) *AddressTransferDestination {
	this := AddressTransferDestination{}
	this.DestinationType = destinationType
	this.AddressStr = addressStr
	return &this
}

// NewAddressTransferDestinationWithDefaults instantiates a new AddressTransferDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransferDestinationWithDefaults() *AddressTransferDestination {
	this := AddressTransferDestination{}
	return &this
}

// GetDestinationType returns the DestinationType field value
func (o *AddressTransferDestination) GetDestinationType() TransferDestinationType {
	if o == nil {
		var ret TransferDestinationType
		return ret
	}

	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *AddressTransferDestination) GetDestinationTypeOk() (*TransferDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value
func (o *AddressTransferDestination) SetDestinationType(v TransferDestinationType) {
	o.DestinationType = v
}

// GetAddressStr returns the AddressStr field value
func (o *AddressTransferDestination) GetAddressStr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressStr
}

// GetAddressStrOk returns a tuple with the AddressStr field value
// and a boolean to check if the value has been set.
func (o *AddressTransferDestination) GetAddressStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressStr, true
}

// SetAddressStr sets field value
func (o *AddressTransferDestination) SetAddressStr(v string) {
	o.AddressStr = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *AddressTransferDestination) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransferDestination) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *AddressTransferDestination) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *AddressTransferDestination) SetMemo(v string) {
	o.Memo = &v
}

func (o AddressTransferDestination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransferDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination_type"] = o.DestinationType
	toSerialize["address_str"] = o.AddressStr
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	return toSerialize, nil
}

func (o *AddressTransferDestination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination_type",
		"address_str",
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

	varAddressTransferDestination := _AddressTransferDestination{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressTransferDestination)

	if err != nil {
		return err
	}

	*o = AddressTransferDestination(varAddressTransferDestination)

	return err
}

type NullableAddressTransferDestination struct {
	value *AddressTransferDestination
	isSet bool
}

func (v NullableAddressTransferDestination) Get() *AddressTransferDestination {
	return v.value
}

func (v *NullableAddressTransferDestination) Set(val *AddressTransferDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransferDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransferDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransferDestination(val *AddressTransferDestination) *NullableAddressTransferDestination {
	return &NullableAddressTransferDestination{value: val, isSet: true}
}

func (v NullableAddressTransferDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransferDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


