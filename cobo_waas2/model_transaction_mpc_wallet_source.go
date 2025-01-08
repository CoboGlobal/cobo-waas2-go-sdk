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

// checks if the TransactionMPCWalletSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionMPCWalletSource{}

// TransactionMPCWalletSource Information about the transaction source type `Org-Controlled` and `User-Controlled`. Refer to [Transaction sources and destinations](https://www.cobo.com/developers/v2/guides/transactions/sources-and-destinations) for a detailed introduction about the supported sources and destinations for each transaction type.  Switch between the tabs to display the properties for different transaction sources. 
type TransactionMPCWalletSource struct {
	SourceType TransactionSourceType `json:"source_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	// The wallet address.
	Address *string `json:"address,omitempty"`
	IncludedUtxos []TransactionUtxo `json:"included_utxos,omitempty"`
	ExcludedUtxos []TransactionUtxo `json:"excluded_utxos,omitempty"`
	// The ID of the key share holder group that is selected to sign the transaction.
	SignerKeyShareHolderGroupId *string `json:"signer_key_share_holder_group_id,omitempty"`
}

type _TransactionMPCWalletSource TransactionMPCWalletSource

// NewTransactionMPCWalletSource instantiates a new TransactionMPCWalletSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionMPCWalletSource(sourceType TransactionSourceType, walletId string) *TransactionMPCWalletSource {
	this := TransactionMPCWalletSource{}
	this.SourceType = sourceType
	this.WalletId = walletId
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

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TransactionMPCWalletSource) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionMPCWalletSource) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TransactionMPCWalletSource) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *TransactionMPCWalletSource) SetAddress(v string) {
	o.Address = &v
}

// GetIncludedUtxos returns the IncludedUtxos field value if set, zero value otherwise.
func (o *TransactionMPCWalletSource) GetIncludedUtxos() []TransactionUtxo {
	if o == nil || IsNil(o.IncludedUtxos) {
		var ret []TransactionUtxo
		return ret
	}
	return o.IncludedUtxos
}

// GetIncludedUtxosOk returns a tuple with the IncludedUtxos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionMPCWalletSource) GetIncludedUtxosOk() ([]TransactionUtxo, bool) {
	if o == nil || IsNil(o.IncludedUtxos) {
		return nil, false
	}
	return o.IncludedUtxos, true
}

// HasIncludedUtxos returns a boolean if a field has been set.
func (o *TransactionMPCWalletSource) HasIncludedUtxos() bool {
	if o != nil && !IsNil(o.IncludedUtxos) {
		return true
	}

	return false
}

// SetIncludedUtxos gets a reference to the given []TransactionUtxo and assigns it to the IncludedUtxos field.
func (o *TransactionMPCWalletSource) SetIncludedUtxos(v []TransactionUtxo) {
	o.IncludedUtxos = v
}

// GetExcludedUtxos returns the ExcludedUtxos field value if set, zero value otherwise.
func (o *TransactionMPCWalletSource) GetExcludedUtxos() []TransactionUtxo {
	if o == nil || IsNil(o.ExcludedUtxos) {
		var ret []TransactionUtxo
		return ret
	}
	return o.ExcludedUtxos
}

// GetExcludedUtxosOk returns a tuple with the ExcludedUtxos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionMPCWalletSource) GetExcludedUtxosOk() ([]TransactionUtxo, bool) {
	if o == nil || IsNil(o.ExcludedUtxos) {
		return nil, false
	}
	return o.ExcludedUtxos, true
}

// HasExcludedUtxos returns a boolean if a field has been set.
func (o *TransactionMPCWalletSource) HasExcludedUtxos() bool {
	if o != nil && !IsNil(o.ExcludedUtxos) {
		return true
	}

	return false
}

// SetExcludedUtxos gets a reference to the given []TransactionUtxo and assigns it to the ExcludedUtxos field.
func (o *TransactionMPCWalletSource) SetExcludedUtxos(v []TransactionUtxo) {
	o.ExcludedUtxos = v
}

// GetSignerKeyShareHolderGroupId returns the SignerKeyShareHolderGroupId field value if set, zero value otherwise.
func (o *TransactionMPCWalletSource) GetSignerKeyShareHolderGroupId() string {
	if o == nil || IsNil(o.SignerKeyShareHolderGroupId) {
		var ret string
		return ret
	}
	return *o.SignerKeyShareHolderGroupId
}

// GetSignerKeyShareHolderGroupIdOk returns a tuple with the SignerKeyShareHolderGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionMPCWalletSource) GetSignerKeyShareHolderGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SignerKeyShareHolderGroupId) {
		return nil, false
	}
	return o.SignerKeyShareHolderGroupId, true
}

// HasSignerKeyShareHolderGroupId returns a boolean if a field has been set.
func (o *TransactionMPCWalletSource) HasSignerKeyShareHolderGroupId() bool {
	if o != nil && !IsNil(o.SignerKeyShareHolderGroupId) {
		return true
	}

	return false
}

// SetSignerKeyShareHolderGroupId gets a reference to the given string and assigns it to the SignerKeyShareHolderGroupId field.
func (o *TransactionMPCWalletSource) SetSignerKeyShareHolderGroupId(v string) {
	o.SignerKeyShareHolderGroupId = &v
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
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.IncludedUtxos) {
		toSerialize["included_utxos"] = o.IncludedUtxos
	}
	if !IsNil(o.ExcludedUtxos) {
		toSerialize["excluded_utxos"] = o.ExcludedUtxos
	}
	if !IsNil(o.SignerKeyShareHolderGroupId) {
		toSerialize["signer_key_share_holder_group_id"] = o.SignerKeyShareHolderGroupId
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


