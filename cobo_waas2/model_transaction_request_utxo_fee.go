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

// checks if the TransactionRequestUtxoFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionRequestUtxoFee{}

// TransactionRequestUtxoFee The preset properties to limit transaction fee.  In the UTXO fee model, the transaction fee is calculated by multiplying the fee rate by the transaction size. This can be expressed as: Transaction fee = fee rate * transaction size. For more information about the UTXO fee model, see [Fee models](https://www.cobo.com/developers/v2/guides/transactions/estimate-fees#fee-models).  You can specify the maximum fee amount to limit the transaction fee. The transaction will fail if the transaction fee exceeds the specified maximum fee amount.  Switch between the tabs to display the properties for different transaction fee models. 
type TransactionRequestUtxoFee struct {
	// The fee rate in sat/vByte. The fee rate represents the satoshis you are willing to pay for each byte of data that your transaction will consume on the blockchain.
	FeeRate *string `json:"fee_rate,omitempty"`
	FeeType FeeType `json:"fee_type"`
	// The token used to pay the transaction fee.
	TokenId string `json:"token_id"`
	// The maximum fee that you are willing to pay for the transaction. Provide the value without applying precision. The transaction will fail if the transaction fee exceeds the maximum fee.
	MaxFeeAmount *string `json:"max_fee_amount,omitempty"`
}

type _TransactionRequestUtxoFee TransactionRequestUtxoFee

// NewTransactionRequestUtxoFee instantiates a new TransactionRequestUtxoFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRequestUtxoFee(feeType FeeType, tokenId string) *TransactionRequestUtxoFee {
	this := TransactionRequestUtxoFee{}
	this.FeeType = feeType
	this.TokenId = tokenId
	return &this
}

// NewTransactionRequestUtxoFeeWithDefaults instantiates a new TransactionRequestUtxoFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRequestUtxoFeeWithDefaults() *TransactionRequestUtxoFee {
	this := TransactionRequestUtxoFee{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = feeType
	return &this
}

// GetFeeRate returns the FeeRate field value if set, zero value otherwise.
func (o *TransactionRequestUtxoFee) GetFeeRate() string {
	if o == nil || IsNil(o.FeeRate) {
		var ret string
		return ret
	}
	return *o.FeeRate
}

// GetFeeRateOk returns a tuple with the FeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequestUtxoFee) GetFeeRateOk() (*string, bool) {
	if o == nil || IsNil(o.FeeRate) {
		return nil, false
	}
	return o.FeeRate, true
}

// HasFeeRate returns a boolean if a field has been set.
func (o *TransactionRequestUtxoFee) HasFeeRate() bool {
	if o != nil && !IsNil(o.FeeRate) {
		return true
	}

	return false
}

// SetFeeRate gets a reference to the given string and assigns it to the FeeRate field.
func (o *TransactionRequestUtxoFee) SetFeeRate(v string) {
	o.FeeRate = &v
}

// GetFeeType returns the FeeType field value
func (o *TransactionRequestUtxoFee) GetFeeType() FeeType {
	if o == nil {
		var ret FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestUtxoFee) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *TransactionRequestUtxoFee) SetFeeType(v FeeType) {
	o.FeeType = v
}

// GetTokenId returns the TokenId field value
func (o *TransactionRequestUtxoFee) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestUtxoFee) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *TransactionRequestUtxoFee) SetTokenId(v string) {
	o.TokenId = v
}

// GetMaxFeeAmount returns the MaxFeeAmount field value if set, zero value otherwise.
func (o *TransactionRequestUtxoFee) GetMaxFeeAmount() string {
	if o == nil || IsNil(o.MaxFeeAmount) {
		var ret string
		return ret
	}
	return *o.MaxFeeAmount
}

// GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequestUtxoFee) GetMaxFeeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MaxFeeAmount) {
		return nil, false
	}
	return o.MaxFeeAmount, true
}

// HasMaxFeeAmount returns a boolean if a field has been set.
func (o *TransactionRequestUtxoFee) HasMaxFeeAmount() bool {
	if o != nil && !IsNil(o.MaxFeeAmount) {
		return true
	}

	return false
}

// SetMaxFeeAmount gets a reference to the given string and assigns it to the MaxFeeAmount field.
func (o *TransactionRequestUtxoFee) SetMaxFeeAmount(v string) {
	o.MaxFeeAmount = &v
}

func (o TransactionRequestUtxoFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRequestUtxoFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeeRate) {
		toSerialize["fee_rate"] = o.FeeRate
	}
	toSerialize["fee_type"] = o.FeeType
	toSerialize["token_id"] = o.TokenId
	if !IsNil(o.MaxFeeAmount) {
		toSerialize["max_fee_amount"] = o.MaxFeeAmount
	}
	return toSerialize, nil
}

func (o *TransactionRequestUtxoFee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fee_type",
		"token_id",
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

	varTransactionRequestUtxoFee := _TransactionRequestUtxoFee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionRequestUtxoFee)

	if err != nil {
		return err
	}

	*o = TransactionRequestUtxoFee(varTransactionRequestUtxoFee)

	return err
}

type NullableTransactionRequestUtxoFee struct {
	value *TransactionRequestUtxoFee
	isSet bool
}

func (v NullableTransactionRequestUtxoFee) Get() *TransactionRequestUtxoFee {
	return v.value
}

func (v *NullableTransactionRequestUtxoFee) Set(val *TransactionRequestUtxoFee) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRequestUtxoFee) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRequestUtxoFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRequestUtxoFee(val *TransactionRequestUtxoFee) *NullableTransactionRequestUtxoFee {
	return &NullableTransactionRequestUtxoFee{value: val, isSet: true}
}

func (v NullableTransactionRequestUtxoFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRequestUtxoFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


