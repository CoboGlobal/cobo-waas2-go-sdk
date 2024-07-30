/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TransactionRequestEvmLegacyFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionRequestEvmLegacyFee{}

// TransactionRequestEvmLegacyFee In the legacy fee model, the transaction fee is calculated by multiplying the gas price by the gas units used by the transaction. This can be expressed as: Transaction fee =  (gas price * gas units used). If the gas units used exceed the gas limit, the transaction will fail. 
type TransactionRequestEvmLegacyFee struct {
	// The gas price, in wei. The gas price represents the amount of ETH that must be paid to validators for processing transactions per gas unit used.
	GasPrice string `json:"gas_price"`
	FeeType FeeType `json:"fee_type"`
	// The token ID of the transaction fee.
	TokenId string `json:"token_id"`
	// The gas limit. It represents the maximum number of gas units that you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. The gas unit cost of each operation varies.
	GasLimit *string `json:"gas_limit,omitempty"`
}

type _TransactionRequestEvmLegacyFee TransactionRequestEvmLegacyFee

// NewTransactionRequestEvmLegacyFee instantiates a new TransactionRequestEvmLegacyFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRequestEvmLegacyFee(gasPrice string, feeType FeeType, tokenId string) *TransactionRequestEvmLegacyFee {
	this := TransactionRequestEvmLegacyFee{}
	this.GasPrice = gasPrice
	this.FeeType = feeType
	this.TokenId = tokenId
	var gasLimit string = "21000"
	this.GasLimit = &gasLimit
	return &this
}

// NewTransactionRequestEvmLegacyFeeWithDefaults instantiates a new TransactionRequestEvmLegacyFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRequestEvmLegacyFeeWithDefaults() *TransactionRequestEvmLegacyFee {
	this := TransactionRequestEvmLegacyFee{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = feeType
	var gasLimit string = "21000"
	this.GasLimit = &gasLimit
	return &this
}

// GetGasPrice returns the GasPrice field value
func (o *TransactionRequestEvmLegacyFee) GetGasPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestEvmLegacyFee) GetGasPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *TransactionRequestEvmLegacyFee) SetGasPrice(v string) {
	o.GasPrice = v
}

// GetFeeType returns the FeeType field value
func (o *TransactionRequestEvmLegacyFee) GetFeeType() FeeType {
	if o == nil {
		var ret FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestEvmLegacyFee) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *TransactionRequestEvmLegacyFee) SetFeeType(v FeeType) {
	o.FeeType = v
}

// GetTokenId returns the TokenId field value
func (o *TransactionRequestEvmLegacyFee) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestEvmLegacyFee) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *TransactionRequestEvmLegacyFee) SetTokenId(v string) {
	o.TokenId = v
}

// GetGasLimit returns the GasLimit field value if set, zero value otherwise.
func (o *TransactionRequestEvmLegacyFee) GetGasLimit() string {
	if o == nil || IsNil(o.GasLimit) {
		var ret string
		return ret
	}
	return *o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequestEvmLegacyFee) GetGasLimitOk() (*string, bool) {
	if o == nil || IsNil(o.GasLimit) {
		return nil, false
	}
	return o.GasLimit, true
}

// HasGasLimit returns a boolean if a field has been set.
func (o *TransactionRequestEvmLegacyFee) HasGasLimit() bool {
	if o != nil && !IsNil(o.GasLimit) {
		return true
	}

	return false
}

// SetGasLimit gets a reference to the given string and assigns it to the GasLimit field.
func (o *TransactionRequestEvmLegacyFee) SetGasLimit(v string) {
	o.GasLimit = &v
}

func (o TransactionRequestEvmLegacyFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRequestEvmLegacyFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gas_price"] = o.GasPrice
	toSerialize["fee_type"] = o.FeeType
	toSerialize["token_id"] = o.TokenId
	if !IsNil(o.GasLimit) {
		toSerialize["gas_limit"] = o.GasLimit
	}
	return toSerialize, nil
}

func (o *TransactionRequestEvmLegacyFee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gas_price",
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

	varTransactionRequestEvmLegacyFee := _TransactionRequestEvmLegacyFee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionRequestEvmLegacyFee)

	if err != nil {
		return err
	}

	*o = TransactionRequestEvmLegacyFee(varTransactionRequestEvmLegacyFee)

	return err
}

type NullableTransactionRequestEvmLegacyFee struct {
	value *TransactionRequestEvmLegacyFee
	isSet bool
}

func (v NullableTransactionRequestEvmLegacyFee) Get() *TransactionRequestEvmLegacyFee {
	return v.value
}

func (v *NullableTransactionRequestEvmLegacyFee) Set(val *TransactionRequestEvmLegacyFee) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRequestEvmLegacyFee) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRequestEvmLegacyFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRequestEvmLegacyFee(val *TransactionRequestEvmLegacyFee) *NullableTransactionRequestEvmLegacyFee {
	return &NullableTransactionRequestEvmLegacyFee{value: val, isSet: true}
}

func (v NullableTransactionRequestEvmLegacyFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRequestEvmLegacyFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


