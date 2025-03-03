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

// checks if the SwapQuote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwapQuote{}

// SwapQuote struct for SwapQuote
type SwapQuote struct {
	// The amount of tokens to pay.
	PayAmount string `json:"pay_amount"`
	// The amount of tokens to receive.
	ReceiveAmount string `json:"receive_amount"`
	// The amount of tokens to pay for fee.
	FeeAmount string `json:"fee_amount"`
	// The minimum amount of tokens to pay.
	MinPayAmount *string `json:"min_pay_amount,omitempty"`
	// The maximum amount of tokens to pay.
	MaxPayAmount *string `json:"max_pay_amount,omitempty"`
	// The minimum amount of tokens to receive.
	MinReceiveAmount *string `json:"min_receive_amount,omitempty"`
	// The maximum amount of tokens to receive.
	MaxReceiveAmount *string `json:"max_receive_amount,omitempty"`
	// The time when the quote will expire, in Unix timestamp format, measured in milliseconds.
	QuoteExpiredTimestamp int32 `json:"quote_expired_timestamp"`
}

type _SwapQuote SwapQuote

// NewSwapQuote instantiates a new SwapQuote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwapQuote(payAmount string, receiveAmount string, feeAmount string, quoteExpiredTimestamp int32) *SwapQuote {
	this := SwapQuote{}
	this.PayAmount = payAmount
	this.ReceiveAmount = receiveAmount
	this.FeeAmount = feeAmount
	this.QuoteExpiredTimestamp = quoteExpiredTimestamp
	return &this
}

// NewSwapQuoteWithDefaults instantiates a new SwapQuote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwapQuoteWithDefaults() *SwapQuote {
	this := SwapQuote{}
	return &this
}

// GetPayAmount returns the PayAmount field value
func (o *SwapQuote) GetPayAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayAmount
}

// GetPayAmountOk returns a tuple with the PayAmount field value
// and a boolean to check if the value has been set.
func (o *SwapQuote) GetPayAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayAmount, true
}

// SetPayAmount sets field value
func (o *SwapQuote) SetPayAmount(v string) {
	o.PayAmount = v
}

// GetReceiveAmount returns the ReceiveAmount field value
func (o *SwapQuote) GetReceiveAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiveAmount
}

// GetReceiveAmountOk returns a tuple with the ReceiveAmount field value
// and a boolean to check if the value has been set.
func (o *SwapQuote) GetReceiveAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiveAmount, true
}

// SetReceiveAmount sets field value
func (o *SwapQuote) SetReceiveAmount(v string) {
	o.ReceiveAmount = v
}

// GetFeeAmount returns the FeeAmount field value
func (o *SwapQuote) GetFeeAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value
// and a boolean to check if the value has been set.
func (o *SwapQuote) GetFeeAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeAmount, true
}

// SetFeeAmount sets field value
func (o *SwapQuote) SetFeeAmount(v string) {
	o.FeeAmount = v
}

// GetMinPayAmount returns the MinPayAmount field value if set, zero value otherwise.
func (o *SwapQuote) GetMinPayAmount() string {
	if o == nil || IsNil(o.MinPayAmount) {
		var ret string
		return ret
	}
	return *o.MinPayAmount
}

// GetMinPayAmountOk returns a tuple with the MinPayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapQuote) GetMinPayAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MinPayAmount) {
		return nil, false
	}
	return o.MinPayAmount, true
}

// HasMinPayAmount returns a boolean if a field has been set.
func (o *SwapQuote) HasMinPayAmount() bool {
	if o != nil && !IsNil(o.MinPayAmount) {
		return true
	}

	return false
}

// SetMinPayAmount gets a reference to the given string and assigns it to the MinPayAmount field.
func (o *SwapQuote) SetMinPayAmount(v string) {
	o.MinPayAmount = &v
}

// GetMaxPayAmount returns the MaxPayAmount field value if set, zero value otherwise.
func (o *SwapQuote) GetMaxPayAmount() string {
	if o == nil || IsNil(o.MaxPayAmount) {
		var ret string
		return ret
	}
	return *o.MaxPayAmount
}

// GetMaxPayAmountOk returns a tuple with the MaxPayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapQuote) GetMaxPayAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MaxPayAmount) {
		return nil, false
	}
	return o.MaxPayAmount, true
}

// HasMaxPayAmount returns a boolean if a field has been set.
func (o *SwapQuote) HasMaxPayAmount() bool {
	if o != nil && !IsNil(o.MaxPayAmount) {
		return true
	}

	return false
}

// SetMaxPayAmount gets a reference to the given string and assigns it to the MaxPayAmount field.
func (o *SwapQuote) SetMaxPayAmount(v string) {
	o.MaxPayAmount = &v
}

// GetMinReceiveAmount returns the MinReceiveAmount field value if set, zero value otherwise.
func (o *SwapQuote) GetMinReceiveAmount() string {
	if o == nil || IsNil(o.MinReceiveAmount) {
		var ret string
		return ret
	}
	return *o.MinReceiveAmount
}

// GetMinReceiveAmountOk returns a tuple with the MinReceiveAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapQuote) GetMinReceiveAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MinReceiveAmount) {
		return nil, false
	}
	return o.MinReceiveAmount, true
}

// HasMinReceiveAmount returns a boolean if a field has been set.
func (o *SwapQuote) HasMinReceiveAmount() bool {
	if o != nil && !IsNil(o.MinReceiveAmount) {
		return true
	}

	return false
}

// SetMinReceiveAmount gets a reference to the given string and assigns it to the MinReceiveAmount field.
func (o *SwapQuote) SetMinReceiveAmount(v string) {
	o.MinReceiveAmount = &v
}

// GetMaxReceiveAmount returns the MaxReceiveAmount field value if set, zero value otherwise.
func (o *SwapQuote) GetMaxReceiveAmount() string {
	if o == nil || IsNil(o.MaxReceiveAmount) {
		var ret string
		return ret
	}
	return *o.MaxReceiveAmount
}

// GetMaxReceiveAmountOk returns a tuple with the MaxReceiveAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapQuote) GetMaxReceiveAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MaxReceiveAmount) {
		return nil, false
	}
	return o.MaxReceiveAmount, true
}

// HasMaxReceiveAmount returns a boolean if a field has been set.
func (o *SwapQuote) HasMaxReceiveAmount() bool {
	if o != nil && !IsNil(o.MaxReceiveAmount) {
		return true
	}

	return false
}

// SetMaxReceiveAmount gets a reference to the given string and assigns it to the MaxReceiveAmount field.
func (o *SwapQuote) SetMaxReceiveAmount(v string) {
	o.MaxReceiveAmount = &v
}

// GetQuoteExpiredTimestamp returns the QuoteExpiredTimestamp field value
func (o *SwapQuote) GetQuoteExpiredTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QuoteExpiredTimestamp
}

// GetQuoteExpiredTimestampOk returns a tuple with the QuoteExpiredTimestamp field value
// and a boolean to check if the value has been set.
func (o *SwapQuote) GetQuoteExpiredTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteExpiredTimestamp, true
}

// SetQuoteExpiredTimestamp sets field value
func (o *SwapQuote) SetQuoteExpiredTimestamp(v int32) {
	o.QuoteExpiredTimestamp = v
}

func (o SwapQuote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwapQuote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pay_amount"] = o.PayAmount
	toSerialize["receive_amount"] = o.ReceiveAmount
	toSerialize["fee_amount"] = o.FeeAmount
	if !IsNil(o.MinPayAmount) {
		toSerialize["min_pay_amount"] = o.MinPayAmount
	}
	if !IsNil(o.MaxPayAmount) {
		toSerialize["max_pay_amount"] = o.MaxPayAmount
	}
	if !IsNil(o.MinReceiveAmount) {
		toSerialize["min_receive_amount"] = o.MinReceiveAmount
	}
	if !IsNil(o.MaxReceiveAmount) {
		toSerialize["max_receive_amount"] = o.MaxReceiveAmount
	}
	toSerialize["quote_expired_timestamp"] = o.QuoteExpiredTimestamp
	return toSerialize, nil
}

func (o *SwapQuote) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pay_amount",
		"receive_amount",
		"fee_amount",
		"quote_expired_timestamp",
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

	varSwapQuote := _SwapQuote{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSwapQuote)

	if err != nil {
		return err
	}

	*o = SwapQuote(varSwapQuote)

	return err
}

type NullableSwapQuote struct {
	value *SwapQuote
	isSet bool
}

func (v NullableSwapQuote) Get() *SwapQuote {
	return v.value
}

func (v *NullableSwapQuote) Set(val *SwapQuote) {
	v.value = val
	v.isSet = true
}

func (v NullableSwapQuote) IsSet() bool {
	return v.isSet
}

func (v *NullableSwapQuote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwapQuote(val *SwapQuote) *NullableSwapQuote {
	return &NullableSwapQuote{value: val, isSet: true}
}

func (v NullableSwapQuote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwapQuote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


