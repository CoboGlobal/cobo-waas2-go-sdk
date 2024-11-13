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

// checks if the MpcTransferSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MpcTransferSource{}

// MpcTransferSource The information about the transaction source types `Org-Controlled` and `User-Controlled`. Refer to [Transaction sources and destinations](/v2/guides/transactions/sources-and-destinations) for a detailed introduction about the supported sources and destinations for each transaction type.  You need to provide either the `address` or `included_utxos` property. If neither property is provided, the transfer will fail.  Switch between the tabs to display the properties for different transaction sources. 
type MpcTransferSource struct {
	SourceType WalletSubtype `json:"source_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	// The wallet address. If you want to specify the UTXOs to be used, please provide the `included_utxos` property. If you specify both the `address` and `included_utxos` properties, the specified included UTXOs must belong to the address.  You need to provide either the `address` or `included_utxos` property. If neither property is provided, the transfer will fail. 
	Address *string `json:"address,omitempty"`
	IncludedUtxos []TransactionUtxo `json:"included_utxos,omitempty"`
	ExcludedUtxos []TransactionUtxo `json:"excluded_utxos,omitempty"`
}

type _MpcTransferSource MpcTransferSource

// NewMpcTransferSource instantiates a new MpcTransferSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMpcTransferSource(sourceType WalletSubtype, walletId string) *MpcTransferSource {
	this := MpcTransferSource{}
	this.SourceType = sourceType
	this.WalletId = walletId
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

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *MpcTransferSource) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpcTransferSource) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *MpcTransferSource) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *MpcTransferSource) SetAddress(v string) {
	o.Address = &v
}

// GetIncludedUtxos returns the IncludedUtxos field value if set, zero value otherwise.
func (o *MpcTransferSource) GetIncludedUtxos() []TransactionUtxo {
	if o == nil || IsNil(o.IncludedUtxos) {
		var ret []TransactionUtxo
		return ret
	}
	return o.IncludedUtxos
}

// GetIncludedUtxosOk returns a tuple with the IncludedUtxos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpcTransferSource) GetIncludedUtxosOk() ([]TransactionUtxo, bool) {
	if o == nil || IsNil(o.IncludedUtxos) {
		return nil, false
	}
	return o.IncludedUtxos, true
}

// HasIncludedUtxos returns a boolean if a field has been set.
func (o *MpcTransferSource) HasIncludedUtxos() bool {
	if o != nil && !IsNil(o.IncludedUtxos) {
		return true
	}

	return false
}

// SetIncludedUtxos gets a reference to the given []TransactionUtxo and assigns it to the IncludedUtxos field.
func (o *MpcTransferSource) SetIncludedUtxos(v []TransactionUtxo) {
	o.IncludedUtxos = v
}

// GetExcludedUtxos returns the ExcludedUtxos field value if set, zero value otherwise.
func (o *MpcTransferSource) GetExcludedUtxos() []TransactionUtxo {
	if o == nil || IsNil(o.ExcludedUtxos) {
		var ret []TransactionUtxo
		return ret
	}
	return o.ExcludedUtxos
}

// GetExcludedUtxosOk returns a tuple with the ExcludedUtxos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpcTransferSource) GetExcludedUtxosOk() ([]TransactionUtxo, bool) {
	if o == nil || IsNil(o.ExcludedUtxos) {
		return nil, false
	}
	return o.ExcludedUtxos, true
}

// HasExcludedUtxos returns a boolean if a field has been set.
func (o *MpcTransferSource) HasExcludedUtxos() bool {
	if o != nil && !IsNil(o.ExcludedUtxos) {
		return true
	}

	return false
}

// SetExcludedUtxos gets a reference to the given []TransactionUtxo and assigns it to the ExcludedUtxos field.
func (o *MpcTransferSource) SetExcludedUtxos(v []TransactionUtxo) {
	o.ExcludedUtxos = v
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
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.IncludedUtxos) {
		toSerialize["included_utxos"] = o.IncludedUtxos
	}
	if !IsNil(o.ExcludedUtxos) {
		toSerialize["excluded_utxos"] = o.ExcludedUtxos
	}
	return toSerialize, nil
}

func (o *MpcTransferSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_type",
		"wallet_id",
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


