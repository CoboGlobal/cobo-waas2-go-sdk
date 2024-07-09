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

// checks if the TransactionMPCWalletSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionMPCWalletSource{}

// TransactionMPCWalletSource The information about the transaction source. Specify either the `account_input` property or the `utxo_inputs` property.
type TransactionMPCWalletSource struct {
	SourceType TransactionSourceType `json:"source_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	MpcUsedKeyGroup MpcSigningGroup `json:"mpc_used_key_group"`
	AccountInput *TransactionMPCWalletSourceAccountInput `json:"account_input,omitempty"`
	UtxoInputs []TransactionMPCWalletSourceUtxoInputsInner `json:"utxo_inputs,omitempty"`
}

type _TransactionMPCWalletSource TransactionMPCWalletSource

// NewTransactionMPCWalletSource instantiates a new TransactionMPCWalletSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionMPCWalletSource(sourceType TransactionSourceType, walletId string, mpcUsedKeyGroup MpcSigningGroup) *TransactionMPCWalletSource {
	this := TransactionMPCWalletSource{}
	this.SourceType = sourceType
	this.WalletId = walletId
	this.MpcUsedKeyGroup = mpcUsedKeyGroup
	return &this
}

// NewTransactionMPCWalletSourceWithDefaults instantiates a new TransactionMPCWalletSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionMPCWalletSourceWithDefaults() *TransactionMPCWalletSource {
	this := TransactionMPCWalletSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *TransactionMPCWalletSource) GetSourceType() TransactionSourceType {
	if o == nil {
		var ret TransactionSourceType
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *TransactionMPCWalletSource) GetSourceTypeOk() (*TransactionSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *TransactionMPCWalletSource) SetSourceType(v TransactionSourceType) {
	o.SourceType = v
}

// GetWalletId returns the WalletId field value
func (o *TransactionMPCWalletSource) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *TransactionMPCWalletSource) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *TransactionMPCWalletSource) SetWalletId(v string) {
	o.WalletId = v
}

// GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field value
func (o *TransactionMPCWalletSource) GetMpcUsedKeyGroup() MpcSigningGroup {
	if o == nil {
		var ret MpcSigningGroup
		return ret
	}

	return o.MpcUsedKeyGroup
}

// GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field value
// and a boolean to check if the value has been set.
func (o *TransactionMPCWalletSource) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MpcUsedKeyGroup, true
}

// SetMpcUsedKeyGroup sets field value
func (o *TransactionMPCWalletSource) SetMpcUsedKeyGroup(v MpcSigningGroup) {
	o.MpcUsedKeyGroup = v
}

// GetAccountInput returns the AccountInput field value if set, zero value otherwise.
func (o *TransactionMPCWalletSource) GetAccountInput() TransactionMPCWalletSourceAccountInput {
	if o == nil || IsNil(o.AccountInput) {
		var ret TransactionMPCWalletSourceAccountInput
		return ret
	}
	return *o.AccountInput
}

// GetAccountInputOk returns a tuple with the AccountInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionMPCWalletSource) GetAccountInputOk() (*TransactionMPCWalletSourceAccountInput, bool) {
	if o == nil || IsNil(o.AccountInput) {
		return nil, false
	}
	return o.AccountInput, true
}

// HasAccountInput returns a boolean if a field has been set.
func (o *TransactionMPCWalletSource) HasAccountInput() bool {
	if o != nil && !IsNil(o.AccountInput) {
		return true
	}

	return false
}

// SetAccountInput gets a reference to the given TransactionMPCWalletSourceAccountInput and assigns it to the AccountInput field.
func (o *TransactionMPCWalletSource) SetAccountInput(v TransactionMPCWalletSourceAccountInput) {
	o.AccountInput = &v
}

// GetUtxoInputs returns the UtxoInputs field value if set, zero value otherwise.
func (o *TransactionMPCWalletSource) GetUtxoInputs() []TransactionMPCWalletSourceUtxoInputsInner {
	if o == nil || IsNil(o.UtxoInputs) {
		var ret []TransactionMPCWalletSourceUtxoInputsInner
		return ret
	}
	return o.UtxoInputs
}

// GetUtxoInputsOk returns a tuple with the UtxoInputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionMPCWalletSource) GetUtxoInputsOk() ([]TransactionMPCWalletSourceUtxoInputsInner, bool) {
	if o == nil || IsNil(o.UtxoInputs) {
		return nil, false
	}
	return o.UtxoInputs, true
}

// HasUtxoInputs returns a boolean if a field has been set.
func (o *TransactionMPCWalletSource) HasUtxoInputs() bool {
	if o != nil && !IsNil(o.UtxoInputs) {
		return true
	}

	return false
}

// SetUtxoInputs gets a reference to the given []TransactionMPCWalletSourceUtxoInputsInner and assigns it to the UtxoInputs field.
func (o *TransactionMPCWalletSource) SetUtxoInputs(v []TransactionMPCWalletSourceUtxoInputsInner) {
	o.UtxoInputs = v
}

func (o TransactionMPCWalletSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionMPCWalletSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["mpc_used_key_group"] = o.MpcUsedKeyGroup
	if !IsNil(o.AccountInput) {
		toSerialize["account_input"] = o.AccountInput
	}
	if !IsNil(o.UtxoInputs) {
		toSerialize["utxo_inputs"] = o.UtxoInputs
	}
	return toSerialize, nil
}

func (o *TransactionMPCWalletSource) UnmarshalJSON(data []byte) (err error) {
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

	varTransactionMPCWalletSource := _TransactionMPCWalletSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionMPCWalletSource)

	if err != nil {
		return err
	}

	*o = TransactionMPCWalletSource(varTransactionMPCWalletSource)

	return err
}

type NullableTransactionMPCWalletSource struct {
	value *TransactionMPCWalletSource
	isSet bool
}

func (v NullableTransactionMPCWalletSource) Get() *TransactionMPCWalletSource {
	return v.value
}

func (v *NullableTransactionMPCWalletSource) Set(val *TransactionMPCWalletSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionMPCWalletSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionMPCWalletSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionMPCWalletSource(val *TransactionMPCWalletSource) *NullableTransactionMPCWalletSource {
	return &NullableTransactionMPCWalletSource{value: val, isSet: true}
}

func (v NullableTransactionMPCWalletSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionMPCWalletSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

