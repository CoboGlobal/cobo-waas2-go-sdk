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

// checks if the TransactionRequestFILFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionRequestFILFee{}

// TransactionRequestFILFee The preset properties to limit transaction fee.  In the Filecoin fee model, the transaction fee is calculated using the minimum of your specified gas fee cap and the sum of the base fee and gas premium, then multiplied by the gas limit. This can be expressed as: Transaction fee = min(gas fee cap, base fee + gas premium) * gas limit. For more information about the Filecoin fee model, refer to [Fee models](https://www.cobo.com/developers/v2/guides/transactions/estimate-fees#fee-models).  You can specify the gas fee cap, gas premium, and gas limit to control fee behavior and prioritization.  Switch between the tabs to display the properties for different transaction fee models. 
type TransactionRequestFILFee struct {
	// An optional tip you can include to prioritize your transaction. The gas premium incentivizes miners to include your transaction sooner than those offering only the base fee.
	GasPremium string `json:"gas_premium"`
	// The maximum gas price you are willing to pay per unit of gas.
	GasFeeCap string `json:"gas_fee_cap"`
	// The maximum amount of gas your transaction is allowed to consume.
	GasLimit *string `json:"gas_limit,omitempty"`
	FeeType FeeType `json:"fee_type"`
	// The token used to pay the transaction fee.
	TokenId string `json:"token_id"`
}

type _TransactionRequestFILFee TransactionRequestFILFee

// NewTransactionRequestFILFee instantiates a new TransactionRequestFILFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRequestFILFee(gasPremium string, gasFeeCap string, feeType FeeType, tokenId string) *TransactionRequestFILFee {
	this := TransactionRequestFILFee{}
	this.GasPremium = gasPremium
	this.GasFeeCap = gasFeeCap
	this.FeeType = feeType
	this.TokenId = tokenId
	return &this
}

// NewTransactionRequestFILFeeWithDefaults instantiates a new TransactionRequestFILFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRequestFILFeeWithDefaults() *TransactionRequestFILFee {
	this := TransactionRequestFILFee{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = feeType
	return &this
}

// GetGasPremium returns the GasPremium field value
func (o *TransactionRequestFILFee) GetGasPremium() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasPremium
}

// GetGasPremiumOk returns a tuple with the GasPremium field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestFILFee) GetGasPremiumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasPremium, true
}

// SetGasPremium sets field value
func (o *TransactionRequestFILFee) SetGasPremium(v string) {
	o.GasPremium = v
}

// GetGasFeeCap returns the GasFeeCap field value
func (o *TransactionRequestFILFee) GetGasFeeCap() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasFeeCap
}

// GetGasFeeCapOk returns a tuple with the GasFeeCap field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestFILFee) GetGasFeeCapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasFeeCap, true
}

// SetGasFeeCap sets field value
func (o *TransactionRequestFILFee) SetGasFeeCap(v string) {
	o.GasFeeCap = v
}

// GetGasLimit returns the GasLimit field value if set, zero value otherwise.
func (o *TransactionRequestFILFee) GetGasLimit() string {
	if o == nil || IsNil(o.GasLimit) {
		var ret string
		return ret
	}
	return *o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequestFILFee) GetGasLimitOk() (*string, bool) {
	if o == nil || IsNil(o.GasLimit) {
		return nil, false
	}
	return o.GasLimit, true
}

// HasGasLimit returns a boolean if a field has been set.
func (o *TransactionRequestFILFee) HasGasLimit() bool {
	if o != nil && !IsNil(o.GasLimit) {
		return true
	}

	return false
}

// SetGasLimit gets a reference to the given string and assigns it to the GasLimit field.
func (o *TransactionRequestFILFee) SetGasLimit(v string) {
	o.GasLimit = &v
}

// GetFeeType returns the FeeType field value
func (o *TransactionRequestFILFee) GetFeeType() FeeType {
	if o == nil {
		var ret FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestFILFee) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *TransactionRequestFILFee) SetFeeType(v FeeType) {
	o.FeeType = v
}

// GetTokenId returns the TokenId field value
func (o *TransactionRequestFILFee) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestFILFee) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *TransactionRequestFILFee) SetTokenId(v string) {
	o.TokenId = v
}

func (o TransactionRequestFILFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRequestFILFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gas_premium"] = o.GasPremium
	toSerialize["gas_fee_cap"] = o.GasFeeCap
	if !IsNil(o.GasLimit) {
		toSerialize["gas_limit"] = o.GasLimit
	}
	toSerialize["fee_type"] = o.FeeType
	toSerialize["token_id"] = o.TokenId
	return toSerialize, nil
}

func (o *TransactionRequestFILFee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gas_premium",
		"gas_fee_cap",
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

	varTransactionRequestFILFee := _TransactionRequestFILFee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionRequestFILFee)

	if err != nil {
		return err
	}

	*o = TransactionRequestFILFee(varTransactionRequestFILFee)

	return err
}

type NullableTransactionRequestFILFee struct {
	value *TransactionRequestFILFee
	isSet bool
}

func (v NullableTransactionRequestFILFee) Get() *TransactionRequestFILFee {
	return v.value
}

func (v *NullableTransactionRequestFILFee) Set(val *TransactionRequestFILFee) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRequestFILFee) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRequestFILFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRequestFILFee(val *TransactionRequestFILFee) *NullableTransactionRequestFILFee {
	return &NullableTransactionRequestFILFee{value: val, isSet: true}
}

func (v NullableTransactionRequestFILFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRequestFILFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


