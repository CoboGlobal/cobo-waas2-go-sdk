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

// checks if the MpcTransferSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MpcTransferSource{}

// MpcTransferSource struct for MpcTransferSource
type MpcTransferSource struct {
	SourceType WalletSubtype `json:"source_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	AccountInput *TransactionMPCWalletSourceAccountInput `json:"account_input,omitempty"`
	UtxoInputs *MpcTransferSourceAllOfUtxoInputs `json:"utxo_inputs,omitempty"`
	MpcUsedKeyGroup MpcSigningGroup `json:"mpc_used_key_group"`
}

type _MpcTransferSource MpcTransferSource

// NewMpcTransferSource instantiates a new MpcTransferSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMpcTransferSource(sourceType WalletSubtype, walletId string, mpcUsedKeyGroup MpcSigningGroup) *MpcTransferSource {
	this := MpcTransferSource{}
	this.SourceType = sourceType
	this.WalletId = walletId
	this.MpcUsedKeyGroup = mpcUsedKeyGroup
	return &this
}

// NewMpcTransferSourceWithDefaults instantiates a new MpcTransferSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMpcTransferSourceWithDefaults() *MpcTransferSource {
	this := MpcTransferSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *MpcTransferSource) GetSourceType() WalletSubtype {
	if o == nil {
		var ret WalletSubtype
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *MpcTransferSource) GetSourceTypeOk() (*WalletSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *MpcTransferSource) SetSourceType(v WalletSubtype) {
	o.SourceType = v
}

// GetWalletId returns the WalletId field value
func (o *MpcTransferSource) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *MpcTransferSource) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *MpcTransferSource) SetWalletId(v string) {
	o.WalletId = v
}

// GetAccountInput returns the AccountInput field value if set, zero value otherwise.
func (o *MpcTransferSource) GetAccountInput() TransactionMPCWalletSourceAccountInput {
	if o == nil || IsNil(o.AccountInput) {
		var ret TransactionMPCWalletSourceAccountInput
		return ret
	}
	return *o.AccountInput
}

// GetAccountInputOk returns a tuple with the AccountInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpcTransferSource) GetAccountInputOk() (*TransactionMPCWalletSourceAccountInput, bool) {
	if o == nil || IsNil(o.AccountInput) {
		return nil, false
	}
	return o.AccountInput, true
}

// HasAccountInput returns a boolean if a field has been set.
func (o *MpcTransferSource) HasAccountInput() bool {
	if o != nil && !IsNil(o.AccountInput) {
		return true
	}

	return false
}

// SetAccountInput gets a reference to the given TransactionMPCWalletSourceAccountInput and assigns it to the AccountInput field.
func (o *MpcTransferSource) SetAccountInput(v TransactionMPCWalletSourceAccountInput) {
	o.AccountInput = &v
}

// GetUtxoInputs returns the UtxoInputs field value if set, zero value otherwise.
func (o *MpcTransferSource) GetUtxoInputs() MpcTransferSourceAllOfUtxoInputs {
	if o == nil || IsNil(o.UtxoInputs) {
		var ret MpcTransferSourceAllOfUtxoInputs
		return ret
	}
	return *o.UtxoInputs
}

// GetUtxoInputsOk returns a tuple with the UtxoInputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpcTransferSource) GetUtxoInputsOk() (*MpcTransferSourceAllOfUtxoInputs, bool) {
	if o == nil || IsNil(o.UtxoInputs) {
		return nil, false
	}
	return o.UtxoInputs, true
}

// HasUtxoInputs returns a boolean if a field has been set.
func (o *MpcTransferSource) HasUtxoInputs() bool {
	if o != nil && !IsNil(o.UtxoInputs) {
		return true
	}

	return false
}

// SetUtxoInputs gets a reference to the given MpcTransferSourceAllOfUtxoInputs and assigns it to the UtxoInputs field.
func (o *MpcTransferSource) SetUtxoInputs(v MpcTransferSourceAllOfUtxoInputs) {
	o.UtxoInputs = &v
}

// GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field value
func (o *MpcTransferSource) GetMpcUsedKeyGroup() MpcSigningGroup {
	if o == nil {
		var ret MpcSigningGroup
		return ret
	}

	return o.MpcUsedKeyGroup
}

// GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field value
// and a boolean to check if the value has been set.
func (o *MpcTransferSource) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MpcUsedKeyGroup, true
}

// SetMpcUsedKeyGroup sets field value
func (o *MpcTransferSource) SetMpcUsedKeyGroup(v MpcSigningGroup) {
	o.MpcUsedKeyGroup = v
}

func (o MpcTransferSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MpcTransferSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	toSerialize["wallet_id"] = o.WalletId
	if !IsNil(o.AccountInput) {
		toSerialize["account_input"] = o.AccountInput
	}
	if !IsNil(o.UtxoInputs) {
		toSerialize["utxo_inputs"] = o.UtxoInputs
	}
	toSerialize["mpc_used_key_group"] = o.MpcUsedKeyGroup
	return toSerialize, nil
}

func (o *MpcTransferSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_type",
		"wallet_id",
		"mpc_used_key_group",
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

	varMpcTransferSource := _MpcTransferSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMpcTransferSource)

	if err != nil {
		return err
	}

	*o = MpcTransferSource(varMpcTransferSource)

	return err
}

type NullableMpcTransferSource struct {
	value *MpcTransferSource
	isSet bool
}

func (v NullableMpcTransferSource) Get() *MpcTransferSource {
	return v.value
}

func (v *NullableMpcTransferSource) Set(val *MpcTransferSource) {
	v.value = val
	v.isSet = true
}

func (v NullableMpcTransferSource) IsSet() bool {
	return v.isSet
}

func (v *NullableMpcTransferSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMpcTransferSource(val *MpcTransferSource) *NullableMpcTransferSource {
	return &NullableMpcTransferSource{value: val, isSet: true}
}

func (v NullableMpcTransferSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMpcTransferSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


