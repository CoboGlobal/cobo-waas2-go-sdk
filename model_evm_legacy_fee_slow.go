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

// checks if the EvmLegacyFeeSlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvmLegacyFeeSlow{}

// EvmLegacyFeeSlow struct for EvmLegacyFeeSlow
type EvmLegacyFeeSlow struct {
	// ID of the fee token. Unique in all chains scope.
	FeeTokenId *string `json:"fee_token_id,omitempty"`
	// The Price of Gas, unit GWei.
	GasPrice string `json:"gas_price"`
	// The Limit of gas.
	GasLimit *string `json:"gas_limit,omitempty"`
	// The estimated fee amount in fee_coin.
	FeeAmount *string `json:"fee_amount,omitempty"`
}

type _EvmLegacyFeeSlow EvmLegacyFeeSlow

// NewEvmLegacyFeeSlow instantiates a new EvmLegacyFeeSlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvmLegacyFeeSlow(gasPrice string) *EvmLegacyFeeSlow {
	this := EvmLegacyFeeSlow{}
	this.GasPrice = gasPrice
	var gasLimit string = "21000"
	this.GasLimit = &gasLimit
	return &this
}

// NewEvmLegacyFeeSlowWithDefaults instantiates a new EvmLegacyFeeSlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvmLegacyFeeSlowWithDefaults() *EvmLegacyFeeSlow {
	this := EvmLegacyFeeSlow{}
	var gasLimit string = "21000"
	this.GasLimit = &gasLimit
	return &this
}

// GetFeeTokenId returns the FeeTokenId field value if set, zero value otherwise.
func (o *EvmLegacyFeeSlow) GetFeeTokenId() string {
	if o == nil || IsNil(o.FeeTokenId) {
		var ret string
		return ret
	}
	return *o.FeeTokenId
}

// GetFeeTokenIdOk returns a tuple with the FeeTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmLegacyFeeSlow) GetFeeTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeeTokenId) {
		return nil, false
	}
	return o.FeeTokenId, true
}

// HasFeeTokenId returns a boolean if a field has been set.
func (o *EvmLegacyFeeSlow) HasFeeTokenId() bool {
	if o != nil && !IsNil(o.FeeTokenId) {
		return true
	}

	return false
}

// SetFeeTokenId gets a reference to the given string and assigns it to the FeeTokenId field.
func (o *EvmLegacyFeeSlow) SetFeeTokenId(v string) {
	o.FeeTokenId = &v
}

// GetGasPrice returns the GasPrice field value
func (o *EvmLegacyFeeSlow) GetGasPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *EvmLegacyFeeSlow) GetGasPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *EvmLegacyFeeSlow) SetGasPrice(v string) {
	o.GasPrice = v
}

// GetGasLimit returns the GasLimit field value if set, zero value otherwise.
func (o *EvmLegacyFeeSlow) GetGasLimit() string {
	if o == nil || IsNil(o.GasLimit) {
		var ret string
		return ret
	}
	return *o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmLegacyFeeSlow) GetGasLimitOk() (*string, bool) {
	if o == nil || IsNil(o.GasLimit) {
		return nil, false
	}
	return o.GasLimit, true
}

// HasGasLimit returns a boolean if a field has been set.
func (o *EvmLegacyFeeSlow) HasGasLimit() bool {
	if o != nil && !IsNil(o.GasLimit) {
		return true
	}

	return false
}

// SetGasLimit gets a reference to the given string and assigns it to the GasLimit field.
func (o *EvmLegacyFeeSlow) SetGasLimit(v string) {
	o.GasLimit = &v
}

// GetFeeAmount returns the FeeAmount field value if set, zero value otherwise.
func (o *EvmLegacyFeeSlow) GetFeeAmount() string {
	if o == nil || IsNil(o.FeeAmount) {
		var ret string
		return ret
	}
	return *o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmLegacyFeeSlow) GetFeeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.FeeAmount) {
		return nil, false
	}
	return o.FeeAmount, true
}

// HasFeeAmount returns a boolean if a field has been set.
func (o *EvmLegacyFeeSlow) HasFeeAmount() bool {
	if o != nil && !IsNil(o.FeeAmount) {
		return true
	}

	return false
}

// SetFeeAmount gets a reference to the given string and assigns it to the FeeAmount field.
func (o *EvmLegacyFeeSlow) SetFeeAmount(v string) {
	o.FeeAmount = &v
}

func (o EvmLegacyFeeSlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvmLegacyFeeSlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeeTokenId) {
		toSerialize["fee_token_id"] = o.FeeTokenId
	}
	toSerialize["gas_price"] = o.GasPrice
	if !IsNil(o.GasLimit) {
		toSerialize["gas_limit"] = o.GasLimit
	}
	if !IsNil(o.FeeAmount) {
		toSerialize["fee_amount"] = o.FeeAmount
	}
	return toSerialize, nil
}

func (o *EvmLegacyFeeSlow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gas_price",
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

	varEvmLegacyFeeSlow := _EvmLegacyFeeSlow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEvmLegacyFeeSlow)

	if err != nil {
		return err
	}

	*o = EvmLegacyFeeSlow(varEvmLegacyFeeSlow)

	return err
}

type NullableEvmLegacyFeeSlow struct {
	value *EvmLegacyFeeSlow
	isSet bool
}

func (v NullableEvmLegacyFeeSlow) Get() *EvmLegacyFeeSlow {
	return v.value
}

func (v *NullableEvmLegacyFeeSlow) Set(val *EvmLegacyFeeSlow) {
	v.value = val
	v.isSet = true
}

func (v NullableEvmLegacyFeeSlow) IsSet() bool {
	return v.isSet
}

func (v *NullableEvmLegacyFeeSlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvmLegacyFeeSlow(val *EvmLegacyFeeSlow) *NullableEvmLegacyFeeSlow {
	return &NullableEvmLegacyFeeSlow{value: val, isSet: true}
}

func (v NullableEvmLegacyFeeSlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvmLegacyFeeSlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


