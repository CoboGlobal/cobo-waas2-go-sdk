/*
Cobo Wallet as a Service 2.0

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

// checks if the CreateMpcWallet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMpcWallet{}

// CreateMpcWallet struct for CreateMpcWallet
type CreateMpcWallet struct {
	Name string `json:"name"`
	WalletType string `json:"wallet_type"`
	WalletSubtype string `json:"wallet_subtype"`
	// The owning mpc vault id of the mpc wallet.
	VaultId string `json:"vault_id"`
}

type _CreateMpcWallet CreateMpcWallet

// NewCreateMpcWallet instantiates a new CreateMpcWallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMpcWallet(name string, walletType string, walletSubtype string, vaultId string) *CreateMpcWallet {
	this := CreateMpcWallet{}
	this.Name = name
	this.WalletType = walletType
	this.WalletSubtype = walletSubtype
	this.VaultId = vaultId
	return &this
}

// NewCreateMpcWalletWithDefaults instantiates a new CreateMpcWallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMpcWalletWithDefaults() *CreateMpcWallet {
	this := CreateMpcWallet{}
	return &this
}

// GetName returns the Name field value
func (o *CreateMpcWallet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateMpcWallet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateMpcWallet) SetName(v string) {
	o.Name = v
}

// GetWalletType returns the WalletType field value
func (o *CreateMpcWallet) GetWalletType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value
// and a boolean to check if the value has been set.
func (o *CreateMpcWallet) GetWalletTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletType, true
}

// SetWalletType sets field value
func (o *CreateMpcWallet) SetWalletType(v string) {
	o.WalletType = v
}

// GetWalletSubtype returns the WalletSubtype field value
func (o *CreateMpcWallet) GetWalletSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletSubtype
}

// GetWalletSubtypeOk returns a tuple with the WalletSubtype field value
// and a boolean to check if the value has been set.
func (o *CreateMpcWallet) GetWalletSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletSubtype, true
}

// SetWalletSubtype sets field value
func (o *CreateMpcWallet) SetWalletSubtype(v string) {
	o.WalletSubtype = v
}

// GetVaultId returns the VaultId field value
func (o *CreateMpcWallet) GetVaultId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultId
}

// GetVaultIdOk returns a tuple with the VaultId field value
// and a boolean to check if the value has been set.
func (o *CreateMpcWallet) GetVaultIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultId, true
}

// SetVaultId sets field value
func (o *CreateMpcWallet) SetVaultId(v string) {
	o.VaultId = v
}

func (o CreateMpcWallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMpcWallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["wallet_type"] = o.WalletType
	toSerialize["wallet_subtype"] = o.WalletSubtype
	toSerialize["vault_id"] = o.VaultId
	return toSerialize, nil
}

func (o *CreateMpcWallet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"wallet_type",
		"wallet_subtype",
		"vault_id",
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

	varCreateMpcWallet := _CreateMpcWallet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateMpcWallet)

	if err != nil {
		return err
	}

	*o = CreateMpcWallet(varCreateMpcWallet)

	return err
}

type NullableCreateMpcWallet struct {
	value *CreateMpcWallet
	isSet bool
}

func (v NullableCreateMpcWallet) Get() *CreateMpcWallet {
	return v.value
}

func (v *NullableCreateMpcWallet) Set(val *CreateMpcWallet) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMpcWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMpcWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMpcWallet(val *CreateMpcWallet) *NullableCreateMpcWallet {
	return &NullableCreateMpcWallet{value: val, isSet: true}
}

func (v NullableCreateMpcWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMpcWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


