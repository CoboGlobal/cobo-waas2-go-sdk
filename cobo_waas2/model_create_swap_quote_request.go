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

// checks if the CreateSwapQuoteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSwapQuoteRequest{}

// CreateSwapQuoteRequest struct for CreateSwapQuoteRequest
type CreateSwapQuoteRequest struct {
	// The unique identifier of the wallet.
	WalletId string `json:"wallet_id"`
	// Unique id of the token to pay.
	PayTokenId string `json:"pay_token_id"`
	// Unique id of the token to receive.
	ReceiveTokenId string `json:"receive_token_id"`
	// Amount of tokens to pay. For example \"0.5 BTC\". Note: Either pay_amount or receive_amount must be provided, but not both. 
	PayAmount *string `json:"pay_amount,omitempty"`
	// Amount of tokens to receive. For example \"0.5 ETH_WBTC\". Note: Either pay_amount or receive_amount must be provided, but not both. 
	ReceiveAmount *string `json:"receive_amount,omitempty"`
}

type _CreateSwapQuoteRequest CreateSwapQuoteRequest

// NewCreateSwapQuoteRequest instantiates a new CreateSwapQuoteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSwapQuoteRequest(walletId string, payTokenId string, receiveTokenId string) *CreateSwapQuoteRequest {
	this := CreateSwapQuoteRequest{}
	this.WalletId = walletId
	this.PayTokenId = payTokenId
	this.ReceiveTokenId = receiveTokenId
	return &this
}

// NewCreateSwapQuoteRequestWithDefaults instantiates a new CreateSwapQuoteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSwapQuoteRequestWithDefaults() *CreateSwapQuoteRequest {
	this := CreateSwapQuoteRequest{}
	return &this
}

// GetWalletId returns the WalletId field value
func (o *CreateSwapQuoteRequest) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *CreateSwapQuoteRequest) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *CreateSwapQuoteRequest) SetWalletId(v string) {
	o.WalletId = v
}

// GetPayTokenId returns the PayTokenId field value
func (o *CreateSwapQuoteRequest) GetPayTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayTokenId
}

// GetPayTokenIdOk returns a tuple with the PayTokenId field value
// and a boolean to check if the value has been set.
func (o *CreateSwapQuoteRequest) GetPayTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayTokenId, true
}

// SetPayTokenId sets field value
func (o *CreateSwapQuoteRequest) SetPayTokenId(v string) {
	o.PayTokenId = v
}

// GetReceiveTokenId returns the ReceiveTokenId field value
func (o *CreateSwapQuoteRequest) GetReceiveTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiveTokenId
}

// GetReceiveTokenIdOk returns a tuple with the ReceiveTokenId field value
// and a boolean to check if the value has been set.
func (o *CreateSwapQuoteRequest) GetReceiveTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiveTokenId, true
}

// SetReceiveTokenId sets field value
func (o *CreateSwapQuoteRequest) SetReceiveTokenId(v string) {
	o.ReceiveTokenId = v
}

// GetPayAmount returns the PayAmount field value if set, zero value otherwise.
func (o *CreateSwapQuoteRequest) GetPayAmount() string {
	if o == nil || IsNil(o.PayAmount) {
		var ret string
		return ret
	}
	return *o.PayAmount
}

// GetPayAmountOk returns a tuple with the PayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSwapQuoteRequest) GetPayAmountOk() (*string, bool) {
	if o == nil || IsNil(o.PayAmount) {
		return nil, false
	}
	return o.PayAmount, true
}

// HasPayAmount returns a boolean if a field has been set.
func (o *CreateSwapQuoteRequest) HasPayAmount() bool {
	if o != nil && !IsNil(o.PayAmount) {
		return true
	}

	return false
}

// SetPayAmount gets a reference to the given string and assigns it to the PayAmount field.
func (o *CreateSwapQuoteRequest) SetPayAmount(v string) {
	o.PayAmount = &v
}

// GetReceiveAmount returns the ReceiveAmount field value if set, zero value otherwise.
func (o *CreateSwapQuoteRequest) GetReceiveAmount() string {
	if o == nil || IsNil(o.ReceiveAmount) {
		var ret string
		return ret
	}
	return *o.ReceiveAmount
}

// GetReceiveAmountOk returns a tuple with the ReceiveAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSwapQuoteRequest) GetReceiveAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiveAmount) {
		return nil, false
	}
	return o.ReceiveAmount, true
}

// HasReceiveAmount returns a boolean if a field has been set.
func (o *CreateSwapQuoteRequest) HasReceiveAmount() bool {
	if o != nil && !IsNil(o.ReceiveAmount) {
		return true
	}

	return false
}

// SetReceiveAmount gets a reference to the given string and assigns it to the ReceiveAmount field.
func (o *CreateSwapQuoteRequest) SetReceiveAmount(v string) {
	o.ReceiveAmount = &v
}

func (o CreateSwapQuoteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSwapQuoteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["pay_token_id"] = o.PayTokenId
	toSerialize["receive_token_id"] = o.ReceiveTokenId
	if !IsNil(o.PayAmount) {
		toSerialize["pay_amount"] = o.PayAmount
	}
	if !IsNil(o.ReceiveAmount) {
		toSerialize["receive_amount"] = o.ReceiveAmount
	}
	return toSerialize, nil
}

func (o *CreateSwapQuoteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wallet_id",
		"pay_token_id",
		"receive_token_id",
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

	varCreateSwapQuoteRequest := _CreateSwapQuoteRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateSwapQuoteRequest)

	if err != nil {
		return err
	}

	*o = CreateSwapQuoteRequest(varCreateSwapQuoteRequest)

	return err
}

type NullableCreateSwapQuoteRequest struct {
	value *CreateSwapQuoteRequest
	isSet bool
}

func (v NullableCreateSwapQuoteRequest) Get() *CreateSwapQuoteRequest {
	return v.value
}

func (v *NullableCreateSwapQuoteRequest) Set(val *CreateSwapQuoteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSwapQuoteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSwapQuoteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSwapQuoteRequest(val *CreateSwapQuoteRequest) *NullableCreateSwapQuoteRequest {
	return &NullableCreateSwapQuoteRequest{value: val, isSet: true}
}

func (v NullableCreateSwapQuoteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSwapQuoteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


