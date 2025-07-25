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

// checks if the TransactionRequestEvmEip1559Fee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionRequestEvmEip1559Fee{}

// TransactionRequestEvmEip1559Fee The preset properties to limit transaction fee.  In the EIP-1559 fee model, the transaction fee is calculated by multiplying the gas price and the gas units used by the transaction. This can be expressed as: Transaction fee = gas price * gas units used. For more information about the EIP-1559 fee model, refer to [Fee models](https://www.cobo.com/developers/v2/guides/transactions/estimate-fees#fee-models).  You can specify the maximum gas fee per gas unit, maximum priority fee per gas unit, and the gas limit to limit the gas price, priority fee per gas unit, gas units used in the transaction.   Switch between the tabs to display the properties for different transaction fee models. 
type TransactionRequestEvmEip1559Fee struct {
	// The maximum gas fee per gas unit used on the chain, in wei.
	MaxFeePerGas string `json:"max_fee_per_gas"`
	// The maximum priority fee per gas unit used, in wei. The maximum priority fee represents the highest amount of miner tips that you are willing to pay for your transaction.
	MaxPriorityFeePerGas string `json:"max_priority_fee_per_gas"`
	FeeType FeeType `json:"fee_type"`
	// The token used to pay the transaction fee.
	TokenId string `json:"token_id"`
	// The gas limit. It represents the maximum number of gas units that you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. The gas unit cost of each operation varies.
	GasLimit *string `json:"gas_limit,omitempty"`
}

type _TransactionRequestEvmEip1559Fee TransactionRequestEvmEip1559Fee

// NewTransactionRequestEvmEip1559Fee instantiates a new TransactionRequestEvmEip1559Fee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRequestEvmEip1559Fee(maxFeePerGas string, maxPriorityFeePerGas string, feeType FeeType, tokenId string) *TransactionRequestEvmEip1559Fee {
	this := TransactionRequestEvmEip1559Fee{}
	this.MaxFeePerGas = maxFeePerGas
	this.MaxPriorityFeePerGas = maxPriorityFeePerGas
	this.FeeType = feeType
	this.TokenId = tokenId
	return &this
}

// NewTransactionRequestEvmEip1559FeeWithDefaults instantiates a new TransactionRequestEvmEip1559Fee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRequestEvmEip1559FeeWithDefaults() *TransactionRequestEvmEip1559Fee {
	this := TransactionRequestEvmEip1559Fee{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = feeType
	return &this
}

// GetMaxFeePerGas returns the MaxFeePerGas field value
func (o *TransactionRequestEvmEip1559Fee) GetMaxFeePerGas() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxFeePerGas
}

// GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestEvmEip1559Fee) GetMaxFeePerGasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxFeePerGas, true
}

// SetMaxFeePerGas sets field value
func (o *TransactionRequestEvmEip1559Fee) SetMaxFeePerGas(v string) {
	o.MaxFeePerGas = v
}

// GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field value
func (o *TransactionRequestEvmEip1559Fee) GetMaxPriorityFeePerGas() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxPriorityFeePerGas
}

// GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestEvmEip1559Fee) GetMaxPriorityFeePerGasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxPriorityFeePerGas, true
}

// SetMaxPriorityFeePerGas sets field value
func (o *TransactionRequestEvmEip1559Fee) SetMaxPriorityFeePerGas(v string) {
	o.MaxPriorityFeePerGas = v
}

// GetFeeType returns the FeeType field value
func (o *TransactionRequestEvmEip1559Fee) GetFeeType() FeeType {
	if o == nil {
		var ret FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestEvmEip1559Fee) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *TransactionRequestEvmEip1559Fee) SetFeeType(v FeeType) {
	o.FeeType = v
}

// GetTokenId returns the TokenId field value
func (o *TransactionRequestEvmEip1559Fee) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestEvmEip1559Fee) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *TransactionRequestEvmEip1559Fee) SetTokenId(v string) {
	o.TokenId = v
}

// GetGasLimit returns the GasLimit field value if set, zero value otherwise.
func (o *TransactionRequestEvmEip1559Fee) GetGasLimit() string {
	if o == nil || IsNil(o.GasLimit) {
		var ret string
		return ret
	}
	return *o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequestEvmEip1559Fee) GetGasLimitOk() (*string, bool) {
	if o == nil || IsNil(o.GasLimit) {
		return nil, false
	}
	return o.GasLimit, true
}

// HasGasLimit returns a boolean if a field has been set.
func (o *TransactionRequestEvmEip1559Fee) HasGasLimit() bool {
	if o != nil && !IsNil(o.GasLimit) {
		return true
	}

	return false
}

// SetGasLimit gets a reference to the given string and assigns it to the GasLimit field.
func (o *TransactionRequestEvmEip1559Fee) SetGasLimit(v string) {
	o.GasLimit = &v
}

func (o TransactionRequestEvmEip1559Fee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRequestEvmEip1559Fee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["max_fee_per_gas"] = o.MaxFeePerGas
	toSerialize["max_priority_fee_per_gas"] = o.MaxPriorityFeePerGas
	toSerialize["fee_type"] = o.FeeType
	toSerialize["token_id"] = o.TokenId
	if !IsNil(o.GasLimit) {
		toSerialize["gas_limit"] = o.GasLimit
	}
	return toSerialize, nil
}

func (o *TransactionRequestEvmEip1559Fee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"max_fee_per_gas",
		"max_priority_fee_per_gas",
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

	varTransactionRequestEvmEip1559Fee := _TransactionRequestEvmEip1559Fee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionRequestEvmEip1559Fee)

	if err != nil {
		return err
	}

	*o = TransactionRequestEvmEip1559Fee(varTransactionRequestEvmEip1559Fee)

	return err
}

type NullableTransactionRequestEvmEip1559Fee struct {
	value *TransactionRequestEvmEip1559Fee
	isSet bool
}

func (v NullableTransactionRequestEvmEip1559Fee) Get() *TransactionRequestEvmEip1559Fee {
	return v.value
}

func (v *NullableTransactionRequestEvmEip1559Fee) Set(val *TransactionRequestEvmEip1559Fee) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRequestEvmEip1559Fee) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRequestEvmEip1559Fee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRequestEvmEip1559Fee(val *TransactionRequestEvmEip1559Fee) *NullableTransactionRequestEvmEip1559Fee {
	return &NullableTransactionRequestEvmEip1559Fee{value: val, isSet: true}
}

func (v NullableTransactionRequestEvmEip1559Fee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRequestEvmEip1559Fee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


