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

// checks if the TravelRuleDepositRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TravelRuleDepositRequest{}

// TravelRuleDepositRequest struct for TravelRuleDepositRequest
type TravelRuleDepositRequest struct {
	// The transaction ID.
	TransactionId string `json:"transaction_id"`
	TravelRuleInfo TravelRuleDepositRequestTravelRuleInfo `json:"travel_rule_info"`
}

type _TravelRuleDepositRequest TravelRuleDepositRequest

// NewTravelRuleDepositRequest instantiates a new TravelRuleDepositRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTravelRuleDepositRequest(transactionId string, travelRuleInfo TravelRuleDepositRequestTravelRuleInfo) *TravelRuleDepositRequest {
	this := TravelRuleDepositRequest{}
	this.TransactionId = transactionId
	this.TravelRuleInfo = travelRuleInfo
	return &this
}

// NewTravelRuleDepositRequestWithDefaults instantiates a new TravelRuleDepositRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTravelRuleDepositRequestWithDefaults() *TravelRuleDepositRequest {
	this := TravelRuleDepositRequest{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *TravelRuleDepositRequest) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *TravelRuleDepositRequest) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *TravelRuleDepositRequest) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetTravelRuleInfo returns the TravelRuleInfo field value
func (o *TravelRuleDepositRequest) GetTravelRuleInfo() TravelRuleDepositRequestTravelRuleInfo {
	if o == nil {
		var ret TravelRuleDepositRequestTravelRuleInfo
		return ret
	}

	return o.TravelRuleInfo
}

// GetTravelRuleInfoOk returns a tuple with the TravelRuleInfo field value
// and a boolean to check if the value has been set.
func (o *TravelRuleDepositRequest) GetTravelRuleInfoOk() (*TravelRuleDepositRequestTravelRuleInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TravelRuleInfo, true
}

// SetTravelRuleInfo sets field value
func (o *TravelRuleDepositRequest) SetTravelRuleInfo(v TravelRuleDepositRequestTravelRuleInfo) {
	o.TravelRuleInfo = v
}

func (o TravelRuleDepositRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TravelRuleDepositRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction_id"] = o.TransactionId
	toSerialize["travel_rule_info"] = o.TravelRuleInfo
	return toSerialize, nil
}

func (o *TravelRuleDepositRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction_id",
		"travel_rule_info",
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

	varTravelRuleDepositRequest := _TravelRuleDepositRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTravelRuleDepositRequest)

	if err != nil {
		return err
	}

	*o = TravelRuleDepositRequest(varTravelRuleDepositRequest)

	return err
}

type NullableTravelRuleDepositRequest struct {
	value *TravelRuleDepositRequest
	isSet bool
}

func (v NullableTravelRuleDepositRequest) Get() *TravelRuleDepositRequest {
	return v.value
}

func (v *NullableTravelRuleDepositRequest) Set(val *TravelRuleDepositRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTravelRuleDepositRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTravelRuleDepositRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTravelRuleDepositRequest(val *TravelRuleDepositRequest) *NullableTravelRuleDepositRequest {
	return &NullableTravelRuleDepositRequest{value: val, isSet: true}
}

func (v NullableTravelRuleDepositRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTravelRuleDepositRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


