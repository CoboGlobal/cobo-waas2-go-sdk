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

// checks if the EvmEip1559FeeSlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvmEip1559FeeSlow{}

// EvmEip1559FeeSlow struct for EvmEip1559FeeSlow
type EvmEip1559FeeSlow struct {
	// The token ID of the transaction fee. Unique in all chains scope.
	FeeTokenId *string `json:"fee_token_id,omitempty"`
	// The max priority fee, in gwei. The max priority fee represents the highest amount of miner tips you are willing to pay for your transaction.
	MaxPriorityFee string `json:"max_priority_fee"`
	// The base fee of chain.
	BaseFee string `json:"base_fee"`
	// The gas limit, which represents the max number of gas units you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. Different operations require varying quantities of gas units.
	GasLimit *string `json:"gas_limit,omitempty"`
}

type _EvmEip1559FeeSlow EvmEip1559FeeSlow

// NewEvmEip1559FeeSlow instantiates a new EvmEip1559FeeSlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvmEip1559FeeSlow(maxPriorityFee string, baseFee string) *EvmEip1559FeeSlow {
	this := EvmEip1559FeeSlow{}
	this.MaxPriorityFee = maxPriorityFee
	this.BaseFee = baseFee
	var gasLimit string = "21000"
	this.GasLimit = &gasLimit
	return &this
}

// NewEvmEip1559FeeSlowWithDefaults instantiates a new EvmEip1559FeeSlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvmEip1559FeeSlowWithDefaults() *EvmEip1559FeeSlow {
	this := EvmEip1559FeeSlow{}
	var gasLimit string = "21000"
	this.GasLimit = &gasLimit
	return &this
}

// GetFeeTokenId returns the FeeTokenId field value if set, zero value otherwise.
func (o *EvmEip1559FeeSlow) GetFeeTokenId() string {
	if o == nil || IsNil(o.FeeTokenId) {
		var ret string
		return ret
	}
	return *o.FeeTokenId
}

// GetFeeTokenIdOk returns a tuple with the FeeTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeeSlow) GetFeeTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeeTokenId) {
		return nil, false
	}
	return o.FeeTokenId, true
}

// HasFeeTokenId returns a boolean if a field has been set.
func (o *EvmEip1559FeeSlow) HasFeeTokenId() bool {
	if o != nil && !IsNil(o.FeeTokenId) {
		return true
	}

	return false
}

// SetFeeTokenId gets a reference to the given string and assigns it to the FeeTokenId field.
func (o *EvmEip1559FeeSlow) SetFeeTokenId(v string) {
	o.FeeTokenId = &v
}

// GetMaxPriorityFee returns the MaxPriorityFee field value
func (o *EvmEip1559FeeSlow) GetMaxPriorityFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxPriorityFee
}

// GetMaxPriorityFeeOk returns a tuple with the MaxPriorityFee field value
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeeSlow) GetMaxPriorityFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxPriorityFee, true
}

// SetMaxPriorityFee sets field value
func (o *EvmEip1559FeeSlow) SetMaxPriorityFee(v string) {
	o.MaxPriorityFee = v
}

// GetBaseFee returns the BaseFee field value
func (o *EvmEip1559FeeSlow) GetBaseFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseFee
}

// GetBaseFeeOk returns a tuple with the BaseFee field value
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeeSlow) GetBaseFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseFee, true
}

// SetBaseFee sets field value
func (o *EvmEip1559FeeSlow) SetBaseFee(v string) {
	o.BaseFee = v
}

// GetGasLimit returns the GasLimit field value if set, zero value otherwise.
func (o *EvmEip1559FeeSlow) GetGasLimit() string {
	if o == nil || IsNil(o.GasLimit) {
		var ret string
		return ret
	}
	return *o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeeSlow) GetGasLimitOk() (*string, bool) {
	if o == nil || IsNil(o.GasLimit) {
		return nil, false
	}
	return o.GasLimit, true
}

// HasGasLimit returns a boolean if a field has been set.
func (o *EvmEip1559FeeSlow) HasGasLimit() bool {
	if o != nil && !IsNil(o.GasLimit) {
		return true
	}

	return false
}

// SetGasLimit gets a reference to the given string and assigns it to the GasLimit field.
func (o *EvmEip1559FeeSlow) SetGasLimit(v string) {
	o.GasLimit = &v
}

func (o EvmEip1559FeeSlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvmEip1559FeeSlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeeTokenId) {
		toSerialize["fee_token_id"] = o.FeeTokenId
	}
	toSerialize["max_priority_fee"] = o.MaxPriorityFee
	toSerialize["base_fee"] = o.BaseFee
	if !IsNil(o.GasLimit) {
		toSerialize["gas_limit"] = o.GasLimit
	}
	return toSerialize, nil
}

func (o *EvmEip1559FeeSlow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"max_priority_fee",
		"base_fee",
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

	varEvmEip1559FeeSlow := _EvmEip1559FeeSlow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEvmEip1559FeeSlow)

	if err != nil {
		return err
	}

	*o = EvmEip1559FeeSlow(varEvmEip1559FeeSlow)

	return err
}

type NullableEvmEip1559FeeSlow struct {
	value *EvmEip1559FeeSlow
	isSet bool
}

func (v NullableEvmEip1559FeeSlow) Get() *EvmEip1559FeeSlow {
	return v.value
}

func (v *NullableEvmEip1559FeeSlow) Set(val *EvmEip1559FeeSlow) {
	v.value = val
	v.isSet = true
}

func (v NullableEvmEip1559FeeSlow) IsSet() bool {
	return v.isSet
}

func (v *NullableEvmEip1559FeeSlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvmEip1559FeeSlow(val *EvmEip1559FeeSlow) *NullableEvmEip1559FeeSlow {
	return &NullableEvmEip1559FeeSlow{value: val, isSet: true}
}

func (v NullableEvmEip1559FeeSlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvmEip1559FeeSlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


