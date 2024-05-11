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
)

// checks if the FeeData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeeData{}

// FeeData The estimated fee amount in fee_coin.
type FeeData struct {
	// The Limit of gas.
	GasLimit *int32 `json:"gas_limit,omitempty"`
	// The estimated fee amount in fee_coin.
	FeeAmount *string `json:"fee_amount,omitempty"`
}

// NewFeeData instantiates a new FeeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeData() *FeeData {
	this := FeeData{}
	var gasLimit int32 = 21000
	this.GasLimit = &gasLimit
	return &this
}

// NewFeeDataWithDefaults instantiates a new FeeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeDataWithDefaults() *FeeData {
	this := FeeData{}
	var gasLimit int32 = 21000
	this.GasLimit = &gasLimit
	return &this
}

// GetGasLimit returns the GasLimit field value if set, zero value otherwise.
func (o *FeeData) GetGasLimit() int32 {
	if o == nil || IsNil(o.GasLimit) {
		var ret int32
		return ret
	}
	return *o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeData) GetGasLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.GasLimit) {
		return nil, false
	}
	return o.GasLimit, true
}

// HasGasLimit returns a boolean if a field has been set.
func (o *FeeData) HasGasLimit() bool {
	if o != nil && !IsNil(o.GasLimit) {
		return true
	}

	return false
}

// SetGasLimit gets a reference to the given int32 and assigns it to the GasLimit field.
func (o *FeeData) SetGasLimit(v int32) {
	o.GasLimit = &v
}

// GetFeeAmount returns the FeeAmount field value if set, zero value otherwise.
func (o *FeeData) GetFeeAmount() string {
	if o == nil || IsNil(o.FeeAmount) {
		var ret string
		return ret
	}
	return *o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeData) GetFeeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.FeeAmount) {
		return nil, false
	}
	return o.FeeAmount, true
}

// HasFeeAmount returns a boolean if a field has been set.
func (o *FeeData) HasFeeAmount() bool {
	if o != nil && !IsNil(o.FeeAmount) {
		return true
	}

	return false
}

// SetFeeAmount gets a reference to the given string and assigns it to the FeeAmount field.
func (o *FeeData) SetFeeAmount(v string) {
	o.FeeAmount = &v
}

func (o FeeData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeeData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GasLimit) {
		toSerialize["gas_limit"] = o.GasLimit
	}
	if !IsNil(o.FeeAmount) {
		toSerialize["fee_amount"] = o.FeeAmount
	}
	return toSerialize, nil
}

type NullableFeeData struct {
	value *FeeData
	isSet bool
}

func (v NullableFeeData) Get() *FeeData {
	return v.value
}

func (v *NullableFeeData) Set(val *FeeData) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeData) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeData(val *FeeData) *NullableFeeData {
	return &NullableFeeData{value: val, isSet: true}
}

func (v NullableFeeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


