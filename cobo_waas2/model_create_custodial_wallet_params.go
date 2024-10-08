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

// checks if the CreateCustodialWalletParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustodialWalletParams{}

// CreateCustodialWalletParams The information of Custodial Wallets.
type CreateCustodialWalletParams struct {
	// The wallet name.
	Name string `json:"name"`
	WalletType WalletType `json:"wallet_type"`
	WalletSubtype WalletSubtype `json:"wallet_subtype"`
}

type _CreateCustodialWalletParams CreateCustodialWalletParams

// NewCreateCustodialWalletParams instantiates a new CreateCustodialWalletParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustodialWalletParams(name string, walletType WalletType, walletSubtype WalletSubtype) *CreateCustodialWalletParams {
	this := CreateCustodialWalletParams{}
	this.Name = name
	this.WalletType = walletType
	this.WalletSubtype = walletSubtype
	return &this
}

// NewCreateCustodialWalletParamsWithDefaults instantiates a new CreateCustodialWalletParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustodialWalletParamsWithDefaults() *CreateCustodialWalletParams {
	this := CreateCustodialWalletParams{}
	return &this
}

// GetName returns the Name field value
func (o *CreateCustodialWalletParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCustodialWalletParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCustodialWalletParams) SetName(v string) {
	o.Name = v
}

// GetWalletType returns the WalletType field value
func (o *CreateCustodialWalletParams) GetWalletType() WalletType {
	if o == nil {
		var ret WalletType
		return ret
	}

	return o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value
// and a boolean to check if the value has been set.
func (o *CreateCustodialWalletParams) GetWalletTypeOk() (*WalletType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletType, true
}

// SetWalletType sets field value
func (o *CreateCustodialWalletParams) SetWalletType(v WalletType) {
	o.WalletType = v
}

// GetWalletSubtype returns the WalletSubtype field value
func (o *CreateCustodialWalletParams) GetWalletSubtype() WalletSubtype {
	if o == nil {
		var ret WalletSubtype
		return ret
	}

	return o.WalletSubtype
}

// GetWalletSubtypeOk returns a tuple with the WalletSubtype field value
// and a boolean to check if the value has been set.
func (o *CreateCustodialWalletParams) GetWalletSubtypeOk() (*WalletSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletSubtype, true
}

// SetWalletSubtype sets field value
func (o *CreateCustodialWalletParams) SetWalletSubtype(v WalletSubtype) {
	o.WalletSubtype = v
}

func (o CreateCustodialWalletParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustodialWalletParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["wallet_type"] = o.WalletType
	toSerialize["wallet_subtype"] = o.WalletSubtype
	return toSerialize, nil
}

func (o *CreateCustodialWalletParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"wallet_type",
		"wallet_subtype",
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

	varCreateCustodialWalletParams := _CreateCustodialWalletParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCustodialWalletParams)

	if err != nil {
		return err
	}

	*o = CreateCustodialWalletParams(varCreateCustodialWalletParams)

	return err
}

type NullableCreateCustodialWalletParams struct {
	value *CreateCustodialWalletParams
	isSet bool
}

func (v NullableCreateCustodialWalletParams) Get() *CreateCustodialWalletParams {
	return v.value
}

func (v *NullableCreateCustodialWalletParams) Set(val *CreateCustodialWalletParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustodialWalletParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustodialWalletParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustodialWalletParams(val *CreateCustodialWalletParams) *NullableCreateCustodialWalletParams {
	return &NullableCreateCustodialWalletParams{value: val, isSet: true}
}

func (v NullableCreateCustodialWalletParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustodialWalletParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


