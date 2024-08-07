/*
Cobo Wallet as a Service 2.0

Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TransactionTokenAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionTokenAmount{}

// TransactionTokenAmount The data for transaction asset information.
type TransactionTokenAmount struct {
	// The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](/v2/api-references/wallets/list-enabled-tokens).
	TokenId *string `json:"token_id,omitempty"`
	// (This concept applies to Exchange Wallets only) The asset ID. An asset ID is the unique identifier of the asset held within your linked exchange account.
	AssetId string `json:"asset_id"`
	// Transaction value (Note that this is an absolute value. If you trade 1.5 BTC, then the value is 1.5) 
	Amount float32 `json:"amount"`
}

type _TransactionTokenAmount TransactionTokenAmount

// NewTransactionTokenAmount instantiates a new TransactionTokenAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionTokenAmount(assetId string, amount float32) *TransactionTokenAmount {
	this := TransactionTokenAmount{}
	this.AssetId = assetId
	this.Amount = amount
	return &this
}

// NewTransactionTokenAmountWithDefaults instantiates a new TransactionTokenAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionTokenAmountWithDefaults() *TransactionTokenAmount {
	this := TransactionTokenAmount{}
	return &this
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *TransactionTokenAmount) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokenAmount) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *TransactionTokenAmount) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *TransactionTokenAmount) SetTokenId(v string) {
	o.TokenId = &v
}

// GetAssetId returns the AssetId field value
func (o *TransactionTokenAmount) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *TransactionTokenAmount) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *TransactionTokenAmount) SetAssetId(v string) {
	o.AssetId = v
}

// GetAmount returns the Amount field value
func (o *TransactionTokenAmount) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionTokenAmount) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionTokenAmount) SetAmount(v float32) {
	o.Amount = v
}

func (o TransactionTokenAmount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionTokenAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	toSerialize["asset_id"] = o.AssetId
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *TransactionTokenAmount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asset_id",
		"amount",
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

	varTransactionTokenAmount := _TransactionTokenAmount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionTokenAmount)

	if err != nil {
		return err
	}

	*o = TransactionTokenAmount(varTransactionTokenAmount)

	return err
}

type NullableTransactionTokenAmount struct {
	value *TransactionTokenAmount
	isSet bool
}

func (v NullableTransactionTokenAmount) Get() *TransactionTokenAmount {
	return v.value
}

func (v *NullableTransactionTokenAmount) Set(val *TransactionTokenAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionTokenAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionTokenAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionTokenAmount(val *TransactionTokenAmount) *NullableTransactionTokenAmount {
	return &NullableTransactionTokenAmount{value: val, isSet: true}
}

func (v NullableTransactionTokenAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionTokenAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


