/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TransactionDepositFromWalletSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionDepositFromWalletSource{}

// TransactionDepositFromWalletSource Information about the transaction source, which varies depending on whether you are the initiator or the receiver of the transaction.  - As the initiator, you will see detailed information about the transaction source, and the `source` will be displayed as one of the following wallet sub-types: `Asset`, `Org-Controlled`, `User-Controlled`, `Safe{Wallet}`, `Main`, or `Sub`. - As the receiver, you will see the `source` as either `DepositFromAddress`, `DepositFromWallet`, or `DepositFromLoop`. If the transaction is from the Loop transfer network, the source will be `DepositFromLoop`. Otherwise, it will be either `DepositFromWallet` (indicating an Exchange Wallet) or `DepositFromAddress` (indicating other wallet type than an Exchange Wallet or an external address). 
type TransactionDepositFromWalletSource struct {
	SourceType TransactionSourceType `json:"source_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	WalletType WalletType `json:"wallet_type"`
	WalletSubtype WalletSubtype `json:"wallet_subtype"`
	// The exchange trading account or a sub-wallet ID.
	SubWalletId *string `json:"sub_wallet_id,omitempty"`
}

type _TransactionDepositFromWalletSource TransactionDepositFromWalletSource

// NewTransactionDepositFromWalletSource instantiates a new TransactionDepositFromWalletSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionDepositFromWalletSource(sourceType TransactionSourceType, walletId string, walletType WalletType, walletSubtype WalletSubtype) *TransactionDepositFromWalletSource {
	this := TransactionDepositFromWalletSource{}
	this.SourceType = sourceType
	this.WalletId = walletId
	this.WalletType = walletType
	this.WalletSubtype = walletSubtype
	return &this
}

// NewTransactionDepositFromWalletSourceWithDefaults instantiates a new TransactionDepositFromWalletSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDepositFromWalletSourceWithDefaults() *TransactionDepositFromWalletSource {
	this := TransactionDepositFromWalletSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *TransactionDepositFromWalletSource) GetSourceType() TransactionSourceType {
	if o == nil {
		var ret TransactionSourceType
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositFromWalletSource) GetSourceTypeOk() (*TransactionSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *TransactionDepositFromWalletSource) SetSourceType(v TransactionSourceType) {
	o.SourceType = v
}

// GetWalletId returns the WalletId field value
func (o *TransactionDepositFromWalletSource) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositFromWalletSource) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *TransactionDepositFromWalletSource) SetWalletId(v string) {
	o.WalletId = v
}

// GetWalletType returns the WalletType field value
func (o *TransactionDepositFromWalletSource) GetWalletType() WalletType {
	if o == nil {
		var ret WalletType
		return ret
	}

	return o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositFromWalletSource) GetWalletTypeOk() (*WalletType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletType, true
}

// SetWalletType sets field value
func (o *TransactionDepositFromWalletSource) SetWalletType(v WalletType) {
	o.WalletType = v
}

// GetWalletSubtype returns the WalletSubtype field value
func (o *TransactionDepositFromWalletSource) GetWalletSubtype() WalletSubtype {
	if o == nil {
		var ret WalletSubtype
		return ret
	}

	return o.WalletSubtype
}

// GetWalletSubtypeOk returns a tuple with the WalletSubtype field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositFromWalletSource) GetWalletSubtypeOk() (*WalletSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletSubtype, true
}

// SetWalletSubtype sets field value
func (o *TransactionDepositFromWalletSource) SetWalletSubtype(v WalletSubtype) {
	o.WalletSubtype = v
}

// GetSubWalletId returns the SubWalletId field value if set, zero value otherwise.
func (o *TransactionDepositFromWalletSource) GetSubWalletId() string {
	if o == nil || IsNil(o.SubWalletId) {
		var ret string
		return ret
	}
	return *o.SubWalletId
}

// GetSubWalletIdOk returns a tuple with the SubWalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDepositFromWalletSource) GetSubWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubWalletId) {
		return nil, false
	}
	return o.SubWalletId, true
}

// HasSubWalletId returns a boolean if a field has been set.
func (o *TransactionDepositFromWalletSource) HasSubWalletId() bool {
	if o != nil && !IsNil(o.SubWalletId) {
		return true
	}

	return false
}

// SetSubWalletId gets a reference to the given string and assigns it to the SubWalletId field.
func (o *TransactionDepositFromWalletSource) SetSubWalletId(v string) {
	o.SubWalletId = &v
}

func (o TransactionDepositFromWalletSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionDepositFromWalletSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["wallet_type"] = o.WalletType
	toSerialize["wallet_subtype"] = o.WalletSubtype
	if !IsNil(o.SubWalletId) {
		toSerialize["sub_wallet_id"] = o.SubWalletId
	}
	return toSerialize, nil
}

func (o *TransactionDepositFromWalletSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_type",
		"wallet_id",
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

	varTransactionDepositFromWalletSource := _TransactionDepositFromWalletSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionDepositFromWalletSource)

	if err != nil {
		return err
	}

	*o = TransactionDepositFromWalletSource(varTransactionDepositFromWalletSource)

	return err
}

type NullableTransactionDepositFromWalletSource struct {
	value *TransactionDepositFromWalletSource
	isSet bool
}

func (v NullableTransactionDepositFromWalletSource) Get() *TransactionDepositFromWalletSource {
	return v.value
}

func (v *NullableTransactionDepositFromWalletSource) Set(val *TransactionDepositFromWalletSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDepositFromWalletSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDepositFromWalletSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDepositFromWalletSource(val *TransactionDepositFromWalletSource) *NullableTransactionDepositFromWalletSource {
	return &NullableTransactionDepositFromWalletSource{value: val, isSet: true}
}

func (v NullableTransactionDepositFromWalletSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDepositFromWalletSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


