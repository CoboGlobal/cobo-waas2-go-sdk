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

// checks if the CreateMpcWalletParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMpcWalletParams{}

// CreateMpcWalletParams struct for CreateMpcWalletParams
type CreateMpcWalletParams struct {
	// The wallet name.
	Name string `json:"name"`
	WalletType WalletType `json:"wallet_type"`
	WalletSubtype WalletSubtype `json:"wallet_subtype"`
	// The ID of the owning vault. You can call [List all vaults](/v2/api-references/wallets--mpc-wallets/list-all-vaults) to retrieve all vault IDs under your organization.
	VaultId string `json:"vault_id"`
}

type _CreateMpcWalletParams CreateMpcWalletParams

// NewCreateMpcWalletParams instantiates a new CreateMpcWalletParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMpcWalletParams(name string, walletType WalletType, walletSubtype WalletSubtype, vaultId string) *CreateMpcWalletParams {
	this := CreateMpcWalletParams{}
	this.Name = name
	this.WalletType = walletType
	this.WalletSubtype = walletSubtype
	this.VaultId = vaultId
	return &this
}

// NewCreateMpcWalletParamsWithDefaults instantiates a new CreateMpcWalletParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMpcWalletParamsWithDefaults() *CreateMpcWalletParams {
	this := CreateMpcWalletParams{}
	return &this
}

// GetName returns the Name field value
func (o *CreateMpcWalletParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateMpcWalletParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateMpcWalletParams) SetName(v string) {
	o.Name = v
}

// GetWalletType returns the WalletType field value
func (o *CreateMpcWalletParams) GetWalletType() WalletType {
	if o == nil {
		var ret WalletType
		return ret
	}

	return o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value
// and a boolean to check if the value has been set.
func (o *CreateMpcWalletParams) GetWalletTypeOk() (*WalletType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletType, true
}

// SetWalletType sets field value
func (o *CreateMpcWalletParams) SetWalletType(v WalletType) {
	o.WalletType = v
}

// GetWalletSubtype returns the WalletSubtype field value
func (o *CreateMpcWalletParams) GetWalletSubtype() WalletSubtype {
	if o == nil {
		var ret WalletSubtype
		return ret
	}

	return o.WalletSubtype
}

// GetWalletSubtypeOk returns a tuple with the WalletSubtype field value
// and a boolean to check if the value has been set.
func (o *CreateMpcWalletParams) GetWalletSubtypeOk() (*WalletSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletSubtype, true
}

// SetWalletSubtype sets field value
func (o *CreateMpcWalletParams) SetWalletSubtype(v WalletSubtype) {
	o.WalletSubtype = v
}

// GetVaultId returns the VaultId field value
func (o *CreateMpcWalletParams) GetVaultId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultId
}

// GetVaultIdOk returns a tuple with the VaultId field value
// and a boolean to check if the value has been set.
func (o *CreateMpcWalletParams) GetVaultIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultId, true
}

// SetVaultId sets field value
func (o *CreateMpcWalletParams) SetVaultId(v string) {
	o.VaultId = v
}

func (o CreateMpcWalletParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMpcWalletParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["wallet_type"] = o.WalletType
	toSerialize["wallet_subtype"] = o.WalletSubtype
	toSerialize["vault_id"] = o.VaultId
	return toSerialize, nil
}

func (o *CreateMpcWalletParams) UnmarshalJSON(data []byte) (err error) {
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

	varCreateMpcWalletParams := _CreateMpcWalletParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateMpcWalletParams)

	if err != nil {
		return err
	}

	*o = CreateMpcWalletParams(varCreateMpcWalletParams)

	return err
}

type NullableCreateMpcWalletParams struct {
	value *CreateMpcWalletParams
	isSet bool
}

func (v NullableCreateMpcWalletParams) Get() *CreateMpcWalletParams {
	return v.value
}

func (v *NullableCreateMpcWalletParams) Set(val *CreateMpcWalletParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMpcWalletParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMpcWalletParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMpcWalletParams(val *CreateMpcWalletParams) *NullableCreateMpcWalletParams {
	return &NullableCreateMpcWalletParams{value: val, isSet: true}
}

func (v NullableCreateMpcWalletParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMpcWalletParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


