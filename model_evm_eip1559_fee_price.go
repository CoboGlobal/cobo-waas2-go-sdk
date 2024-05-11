/*
Cobo Wallet as a Service 2.0

Cobo WaaS 2.0 enables you to programmatically access Cobo's full suite of crypto wallet technologies with powerful and flexible access controls.  # Wallet technologies - Custodial Wallet - MPC Wallet - Smart Contract Wallet (Based on Safe{Wallet}) - Exchange Wallet  # Risk Control technologies - Workflow - Access Control List (ACL)  # Risk Control targets - Wallet Management   - User/team and their permission management   - Risk control configurations, e.g. whitelist, blacklist, rate-limiting etc. - Blockchain Interaction   - Crypto transfer   - Smart Contract Invocation  # Important HTTPS only. RESTful, resource oriented  # Get Started Set up your APIs or get authorization  # Authentication and Authorization CoboAuth  # Request and Response application/json  # Error Handling  ### Common error codes | Error Code | Description | | -- | -- |  ### API-specific error codes For error codes that are dedicated to a specific API, see the Error codes section in each API specification, for example, /v3/wallets.  # Rate and Usage Limiting  # Idempotent Request  # Pagination # Support [Developer Hub](https://cobo.com/developers) 

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

// checks if the EvmEip1559FeePrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvmEip1559FeePrice{}

// EvmEip1559FeePrice The base eip1559 fee data for estimate fees or transfer.
type EvmEip1559FeePrice struct {
	// ID of the fee token. Unique in all chains scope.
	FeeTokenId *string `json:"fee_token_id,omitempty"`
	// The highest Gas price paid for the transfer, unit GWei.
	MaxFee string `json:"max_fee"`
	// The maximum Gas price paid to miners, the higher it is, the faster it is likely to be packaged into the block, unit GWei.
	MaxPriorityFee int32 `json:"max_priority_fee"`
	// The Base Fee of chain.
	BaseFee int32 `json:"base_fee"`
	FeeType FeeType `json:"fee_type"`
}

type _EvmEip1559FeePrice EvmEip1559FeePrice

// NewEvmEip1559FeePrice instantiates a new EvmEip1559FeePrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvmEip1559FeePrice(maxFee string, maxPriorityFee int32, baseFee int32, feeType FeeType) *EvmEip1559FeePrice {
	this := EvmEip1559FeePrice{}
	this.MaxFee = maxFee
	this.MaxPriorityFee = maxPriorityFee
	this.BaseFee = baseFee
	this.FeeType = feeType
	return &this
}

// NewEvmEip1559FeePriceWithDefaults instantiates a new EvmEip1559FeePrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvmEip1559FeePriceWithDefaults() *EvmEip1559FeePrice {
	this := EvmEip1559FeePrice{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = feeType
	return &this
}

// GetFeeTokenId returns the FeeTokenId field value if set, zero value otherwise.
func (o *EvmEip1559FeePrice) GetFeeTokenId() string {
	if o == nil || IsNil(o.FeeTokenId) {
		var ret string
		return ret
	}
	return *o.FeeTokenId
}

// GetFeeTokenIdOk returns a tuple with the FeeTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeePrice) GetFeeTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeeTokenId) {
		return nil, false
	}
	return o.FeeTokenId, true
}

// HasFeeTokenId returns a boolean if a field has been set.
func (o *EvmEip1559FeePrice) HasFeeTokenId() bool {
	if o != nil && !IsNil(o.FeeTokenId) {
		return true
	}

	return false
}

// SetFeeTokenId gets a reference to the given string and assigns it to the FeeTokenId field.
func (o *EvmEip1559FeePrice) SetFeeTokenId(v string) {
	o.FeeTokenId = &v
}

// GetMaxFee returns the MaxFee field value
func (o *EvmEip1559FeePrice) GetMaxFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxFee
}

// GetMaxFeeOk returns a tuple with the MaxFee field value
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeePrice) GetMaxFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxFee, true
}

// SetMaxFee sets field value
func (o *EvmEip1559FeePrice) SetMaxFee(v string) {
	o.MaxFee = v
}

// GetMaxPriorityFee returns the MaxPriorityFee field value
func (o *EvmEip1559FeePrice) GetMaxPriorityFee() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxPriorityFee
}

// GetMaxPriorityFeeOk returns a tuple with the MaxPriorityFee field value
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeePrice) GetMaxPriorityFeeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxPriorityFee, true
}

// SetMaxPriorityFee sets field value
func (o *EvmEip1559FeePrice) SetMaxPriorityFee(v int32) {
	o.MaxPriorityFee = v
}

// GetBaseFee returns the BaseFee field value
func (o *EvmEip1559FeePrice) GetBaseFee() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BaseFee
}

// GetBaseFeeOk returns a tuple with the BaseFee field value
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeePrice) GetBaseFeeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseFee, true
}

// SetBaseFee sets field value
func (o *EvmEip1559FeePrice) SetBaseFee(v int32) {
	o.BaseFee = v
}

// GetFeeType returns the FeeType field value
func (o *EvmEip1559FeePrice) GetFeeType() FeeType {
	if o == nil {
		var ret FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeePrice) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *EvmEip1559FeePrice) SetFeeType(v FeeType) {
	o.FeeType = v
}

func (o EvmEip1559FeePrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvmEip1559FeePrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeeTokenId) {
		toSerialize["fee_token_id"] = o.FeeTokenId
	}
	toSerialize["max_fee"] = o.MaxFee
	toSerialize["max_priority_fee"] = o.MaxPriorityFee
	toSerialize["base_fee"] = o.BaseFee
	toSerialize["fee_type"] = o.FeeType
	return toSerialize, nil
}

func (o *EvmEip1559FeePrice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"max_fee",
		"max_priority_fee",
		"base_fee",
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

	varEvmEip1559FeePrice := _EvmEip1559FeePrice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEvmEip1559FeePrice)

	if err != nil {
		return err
	}

	*o = EvmEip1559FeePrice(varEvmEip1559FeePrice)

	return err
}

type NullableEvmEip1559FeePrice struct {
	value *EvmEip1559FeePrice
	isSet bool
}

func (v NullableEvmEip1559FeePrice) Get() *EvmEip1559FeePrice {
	return v.value
}

func (v *NullableEvmEip1559FeePrice) Set(val *EvmEip1559FeePrice) {
	v.value = val
	v.isSet = true
}

func (v NullableEvmEip1559FeePrice) IsSet() bool {
	return v.isSet
}

func (v *NullableEvmEip1559FeePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvmEip1559FeePrice(val *EvmEip1559FeePrice) *NullableEvmEip1559FeePrice {
	return &NullableEvmEip1559FeePrice{value: val, isSet: true}
}

func (v NullableEvmEip1559FeePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvmEip1559FeePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


