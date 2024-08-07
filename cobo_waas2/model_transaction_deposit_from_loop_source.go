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

// checks if the TransactionDepositFromLoopSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionDepositFromLoopSource{}

// TransactionDepositFromLoopSource Information about the transaction source type `DepositFromLoop`. 
type TransactionDepositFromLoopSource struct {
	SourceType TransactionSourceType `json:"source_type"`
}

type _TransactionDepositFromLoopSource TransactionDepositFromLoopSource

// NewTransactionDepositFromLoopSource instantiates a new TransactionDepositFromLoopSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionDepositFromLoopSource(sourceType TransactionSourceType) *TransactionDepositFromLoopSource {
	this := TransactionDepositFromLoopSource{}
	this.SourceType = sourceType
	return &this
}

// NewTransactionDepositFromLoopSourceWithDefaults instantiates a new TransactionDepositFromLoopSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDepositFromLoopSourceWithDefaults() *TransactionDepositFromLoopSource {
	this := TransactionDepositFromLoopSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *TransactionDepositFromLoopSource) GetSourceType() TransactionSourceType {
	if o == nil {
		var ret TransactionSourceType
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositFromLoopSource) GetSourceTypeOk() (*TransactionSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *TransactionDepositFromLoopSource) SetSourceType(v TransactionSourceType) {
	o.SourceType = v
}

func (o TransactionDepositFromLoopSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionDepositFromLoopSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	return toSerialize, nil
}

func (o *TransactionDepositFromLoopSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_type",
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

	varTransactionDepositFromLoopSource := _TransactionDepositFromLoopSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionDepositFromLoopSource)

	if err != nil {
		return err
	}

	*o = TransactionDepositFromLoopSource(varTransactionDepositFromLoopSource)

	return err
}

type NullableTransactionDepositFromLoopSource struct {
	value *TransactionDepositFromLoopSource
	isSet bool
}

func (v NullableTransactionDepositFromLoopSource) Get() *TransactionDepositFromLoopSource {
	return v.value
}

func (v *NullableTransactionDepositFromLoopSource) Set(val *TransactionDepositFromLoopSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDepositFromLoopSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDepositFromLoopSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDepositFromLoopSource(val *TransactionDepositFromLoopSource) *NullableTransactionDepositFromLoopSource {
	return &NullableTransactionDepositFromLoopSource{value: val, isSet: true}
}

func (v NullableTransactionDepositFromLoopSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDepositFromLoopSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


