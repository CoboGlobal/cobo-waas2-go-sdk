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

// checks if the TokenizationPauseEstimateFeeParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenizationPauseEstimateFeeParams{}

// TokenizationPauseEstimateFeeParams struct for TokenizationPauseEstimateFeeParams
type TokenizationPauseEstimateFeeParams struct {
	Source TokenizationTokenOperationSource `json:"source"`
	OperationType TokenizationOperationType `json:"operation_type"`
	// The ID of the token.
	TokenId string `json:"token_id"`
}

type _TokenizationPauseEstimateFeeParams TokenizationPauseEstimateFeeParams

// NewTokenizationPauseEstimateFeeParams instantiates a new TokenizationPauseEstimateFeeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizationPauseEstimateFeeParams(source TokenizationTokenOperationSource, operationType TokenizationOperationType, tokenId string) *TokenizationPauseEstimateFeeParams {
	this := TokenizationPauseEstimateFeeParams{}
	this.Source = source
	this.OperationType = operationType
	this.TokenId = tokenId
	return &this
}

// NewTokenizationPauseEstimateFeeParamsWithDefaults instantiates a new TokenizationPauseEstimateFeeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizationPauseEstimateFeeParamsWithDefaults() *TokenizationPauseEstimateFeeParams {
	this := TokenizationPauseEstimateFeeParams{}
	return &this
}

// GetSource returns the Source field value
func (o *TokenizationPauseEstimateFeeParams) GetSource() TokenizationTokenOperationSource {
	if o == nil {
		var ret TokenizationTokenOperationSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TokenizationPauseEstimateFeeParams) GetSourceOk() (*TokenizationTokenOperationSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *TokenizationPauseEstimateFeeParams) SetSource(v TokenizationTokenOperationSource) {
	o.Source = v
}

// GetOperationType returns the OperationType field value
func (o *TokenizationPauseEstimateFeeParams) GetOperationType() TokenizationOperationType {
	if o == nil {
		var ret TokenizationOperationType
		return ret
	}

	return o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value
// and a boolean to check if the value has been set.
func (o *TokenizationPauseEstimateFeeParams) GetOperationTypeOk() (*TokenizationOperationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationType, true
}

// SetOperationType sets field value
func (o *TokenizationPauseEstimateFeeParams) SetOperationType(v TokenizationOperationType) {
	o.OperationType = v
}

// GetTokenId returns the TokenId field value
func (o *TokenizationPauseEstimateFeeParams) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *TokenizationPauseEstimateFeeParams) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *TokenizationPauseEstimateFeeParams) SetTokenId(v string) {
	o.TokenId = v
}

func (o TokenizationPauseEstimateFeeParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizationPauseEstimateFeeParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["operation_type"] = o.OperationType
	toSerialize["token_id"] = o.TokenId
	return toSerialize, nil
}

func (o *TokenizationPauseEstimateFeeParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"operation_type",
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

	varTokenizationPauseEstimateFeeParams := _TokenizationPauseEstimateFeeParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenizationPauseEstimateFeeParams)

	if err != nil {
		return err
	}

	*o = TokenizationPauseEstimateFeeParams(varTokenizationPauseEstimateFeeParams)

	return err
}

type NullableTokenizationPauseEstimateFeeParams struct {
	value *TokenizationPauseEstimateFeeParams
	isSet bool
}

func (v NullableTokenizationPauseEstimateFeeParams) Get() *TokenizationPauseEstimateFeeParams {
	return v.value
}

func (v *NullableTokenizationPauseEstimateFeeParams) Set(val *TokenizationPauseEstimateFeeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizationPauseEstimateFeeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizationPauseEstimateFeeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizationPauseEstimateFeeParams(val *TokenizationPauseEstimateFeeParams) *NullableTokenizationPauseEstimateFeeParams {
	return &NullableTokenizationPauseEstimateFeeParams{value: val, isSet: true}
}

func (v NullableTokenizationPauseEstimateFeeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizationPauseEstimateFeeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


