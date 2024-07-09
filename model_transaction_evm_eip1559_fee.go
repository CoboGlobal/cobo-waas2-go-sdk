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

// checks if the TransactionEvmEip1559Fee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionEvmEip1559Fee{}

// TransactionEvmEip1559Fee The transaction fee actually charged by the chain that uses the EIP-1559 fee model.  The transaction fee is calculated by multiplying the sum of the base fee price and the priority fee by the used gas. This can be expressed as: Transaction fee = (base fee price + maximum priority fee) * used gas. 
type TransactionEvmEip1559Fee struct {
	// The maximum priority fee, in gwei. The maximum priority fee represents the highest amount of miner tips that you are willing to pay for your transaction.
	MaxPriorityFee string `json:"max_priority_fee"`
	// The base fee price of the chain, in gwei.
	BaseFee string `json:"base_fee"`
	// The gas limit. It represents the maximum number of gas units that you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. The gas unit cost of each operation varies.
	GasLimit string `json:"gas_limit"`
	FeeType FeeType `json:"fee_type"`
	// The transaction fee.
	FeeUsed *string `json:"fee_used,omitempty"`
	// The number of gas units used in the transaction.
	GasUsed *string `json:"gas_used,omitempty"`
}

type _TransactionEvmEip1559Fee TransactionEvmEip1559Fee

// NewTransactionEvmEip1559Fee instantiates a new TransactionEvmEip1559Fee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEvmEip1559Fee(maxPriorityFee string, baseFee string, gasLimit string, feeType FeeType) *TransactionEvmEip1559Fee {
	this := TransactionEvmEip1559Fee{}
	this.MaxPriorityFee = maxPriorityFee
	this.BaseFee = baseFee
	this.GasLimit = gasLimit
	this.FeeType = feeType
	return &this
}

// NewTransactionEvmEip1559FeeWithDefaults instantiates a new TransactionEvmEip1559Fee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEvmEip1559FeeWithDefaults() *TransactionEvmEip1559Fee {
	this := TransactionEvmEip1559Fee{}
	var gasLimit string = "21000"
	this.GasLimit = gasLimit
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = feeType
	return &this
}

// GetMaxPriorityFee returns the MaxPriorityFee field value
func (o *TransactionEvmEip1559Fee) GetMaxPriorityFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxPriorityFee
}

// GetMaxPriorityFeeOk returns a tuple with the MaxPriorityFee field value
// and a boolean to check if the value has been set.
func (o *TransactionEvmEip1559Fee) GetMaxPriorityFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxPriorityFee, true
}

// SetMaxPriorityFee sets field value
func (o *TransactionEvmEip1559Fee) SetMaxPriorityFee(v string) {
	o.MaxPriorityFee = v
}

// GetBaseFee returns the BaseFee field value
func (o *TransactionEvmEip1559Fee) GetBaseFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseFee
}

// GetBaseFeeOk returns a tuple with the BaseFee field value
// and a boolean to check if the value has been set.
func (o *TransactionEvmEip1559Fee) GetBaseFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseFee, true
}

// SetBaseFee sets field value
func (o *TransactionEvmEip1559Fee) SetBaseFee(v string) {
	o.BaseFee = v
}

// GetGasLimit returns the GasLimit field value
func (o *TransactionEvmEip1559Fee) GetGasLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *TransactionEvmEip1559Fee) GetGasLimitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *TransactionEvmEip1559Fee) SetGasLimit(v string) {
	o.GasLimit = v
}

// GetFeeType returns the FeeType field value
func (o *TransactionEvmEip1559Fee) GetFeeType() FeeType {
	if o == nil {
		var ret FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *TransactionEvmEip1559Fee) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *TransactionEvmEip1559Fee) SetFeeType(v FeeType) {
	o.FeeType = v
}

// GetFeeUsed returns the FeeUsed field value if set, zero value otherwise.
func (o *TransactionEvmEip1559Fee) GetFeeUsed() string {
	if o == nil || IsNil(o.FeeUsed) {
		var ret string
		return ret
	}
	return *o.FeeUsed
}

// GetFeeUsedOk returns a tuple with the FeeUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionEvmEip1559Fee) GetFeeUsedOk() (*string, bool) {
	if o == nil || IsNil(o.FeeUsed) {
		return nil, false
	}
	return o.FeeUsed, true
}

// HasFeeUsed returns a boolean if a field has been set.
func (o *TransactionEvmEip1559Fee) HasFeeUsed() bool {
	if o != nil && !IsNil(o.FeeUsed) {
		return true
	}

	return false
}

// SetFeeUsed gets a reference to the given string and assigns it to the FeeUsed field.
func (o *TransactionEvmEip1559Fee) SetFeeUsed(v string) {
	o.FeeUsed = &v
}

// GetGasUsed returns the GasUsed field value if set, zero value otherwise.
func (o *TransactionEvmEip1559Fee) GetGasUsed() string {
	if o == nil || IsNil(o.GasUsed) {
		var ret string
		return ret
	}
	return *o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionEvmEip1559Fee) GetGasUsedOk() (*string, bool) {
	if o == nil || IsNil(o.GasUsed) {
		return nil, false
	}
	return o.GasUsed, true
}

// HasGasUsed returns a boolean if a field has been set.
func (o *TransactionEvmEip1559Fee) HasGasUsed() bool {
	if o != nil && !IsNil(o.GasUsed) {
		return true
	}

	return false
}

// SetGasUsed gets a reference to the given string and assigns it to the GasUsed field.
func (o *TransactionEvmEip1559Fee) SetGasUsed(v string) {
	o.GasUsed = &v
}

func (o TransactionEvmEip1559Fee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionEvmEip1559Fee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["max_priority_fee"] = o.MaxPriorityFee
	toSerialize["base_fee"] = o.BaseFee
	toSerialize["gas_limit"] = o.GasLimit
	toSerialize["fee_type"] = o.FeeType
	if !IsNil(o.FeeUsed) {
		toSerialize["fee_used"] = o.FeeUsed
	}
	if !IsNil(o.GasUsed) {
		toSerialize["gas_used"] = o.GasUsed
	}
	return toSerialize, nil
}

func (o *TransactionEvmEip1559Fee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"max_priority_fee",
		"base_fee",
		"gas_limit",
		"fee_type",
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

	varTransactionEvmEip1559Fee := _TransactionEvmEip1559Fee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionEvmEip1559Fee)

	if err != nil {
		return err
	}

	*o = TransactionEvmEip1559Fee(varTransactionEvmEip1559Fee)

	return err
}

type NullableTransactionEvmEip1559Fee struct {
	value *TransactionEvmEip1559Fee
	isSet bool
}

func (v NullableTransactionEvmEip1559Fee) Get() *TransactionEvmEip1559Fee {
	return v.value
}

func (v *NullableTransactionEvmEip1559Fee) Set(val *TransactionEvmEip1559Fee) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEvmEip1559Fee) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEvmEip1559Fee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEvmEip1559Fee(val *TransactionEvmEip1559Fee) *NullableTransactionEvmEip1559Fee {
	return &NullableTransactionEvmEip1559Fee{value: val, isSet: true}
}

func (v NullableTransactionEvmEip1559Fee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEvmEip1559Fee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


