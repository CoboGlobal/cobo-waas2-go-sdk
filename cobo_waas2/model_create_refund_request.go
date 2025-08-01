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

// checks if the CreateRefundRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRefundRequest{}

// CreateRefundRequest struct for CreateRefundRequest
type CreateRefundRequest struct {
	// The request ID that is used to track a refund request. The request ID is provided by you and must be unique.
	RequestId string `json:"request_id"`
	// The merchant ID.
	MerchantId *string `json:"merchant_id,omitempty"`
	// The amount to refund in cryptocurrency. The amount must be a positive integer with up to two decimal places.
	PayableAmount string `json:"payable_amount"`
	// The address where the refunded cryptocurrency will be sent.
	ToAddress *string `json:"to_address,omitempty"`
	// The ID of the cryptocurrency used for refund. Supported values:    - USDC: `ETH_USDC`, `ARBITRUM_USDCOIN`, `SOL_USDC`, `BASE_USDC`, `MATIC_USDC2`, `BSC_USDC`   - USDT: `TRON_USDT`, `ETH_USDT`, `ARBITRUM_USDT`, `SOL_USDT`, `BASE_USDT`, `MATIC_USDT`, `BSC_USDT` 
	TokenId string `json:"token_id"`
	RefundType RefundType `json:"refund_type"`
	// The ID of the original pay-in order associated with this refund. Use this to track refunds against specific payments.
	OrderId *string `json:"order_id,omitempty"`
	// Whether to charge developer fee to the merchant.     - `true`: The fee amount (specified in `merchant_fee_amount`) will be deducted from the merchant's balance and added to the developer's balance    - `false`: The merchant is not charged any developer fee  When enabled, ensure both `merchant_fee_amount` and `merchant_fee_token_id` are properly specified. 
	ChargeMerchantFee *bool `json:"charge_merchant_fee,omitempty"`
	// The developer fee amount to charge the merchant, denominated in the cryptocurrency specified by `merchant_fee_token_id`. Required when `charge_merchant_fee` is `true`. Must be:   - A positive integer with up to two decimal places.   - Less than the refund amount 
	MerchantFeeAmount *string `json:"merchant_fee_amount,omitempty"`
	// The ID of the cryptocurrency used for the developer fee. It must be the same as `token_id`. Supported values:   - USDC: `ETH_USDC`, `ARBITRUM_USDCOIN`, `SOL_USDC`, `BASE_USDC`, `MATIC_USDC2`, `BSC_USDC`   - USDT: `TRON_USDT`, `ETH_USDT`, `ARBITRUM_USDT`, `SOL_USDT`, `BASE_USDT`, `MATIC_USDT`, `BSC_USDT` 
	MerchantFeeTokenId *string `json:"merchant_fee_token_id,omitempty"`
}

type _CreateRefundRequest CreateRefundRequest

// NewCreateRefundRequest instantiates a new CreateRefundRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRefundRequest(requestId string, payableAmount string, tokenId string, refundType RefundType) *CreateRefundRequest {
	this := CreateRefundRequest{}
	this.RequestId = requestId
	this.PayableAmount = payableAmount
	this.TokenId = tokenId
	this.RefundType = refundType
	return &this
}

// NewCreateRefundRequestWithDefaults instantiates a new CreateRefundRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRefundRequestWithDefaults() *CreateRefundRequest {
	this := CreateRefundRequest{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *CreateRefundRequest) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreateRefundRequest) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreateRefundRequest) SetRequestId(v string) {
	o.RequestId = v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *CreateRefundRequest) GetMerchantId() string {
	if o == nil || IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefundRequest) GetMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *CreateRefundRequest) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *CreateRefundRequest) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetPayableAmount returns the PayableAmount field value
func (o *CreateRefundRequest) GetPayableAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayableAmount
}

// GetPayableAmountOk returns a tuple with the PayableAmount field value
// and a boolean to check if the value has been set.
func (o *CreateRefundRequest) GetPayableAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayableAmount, true
}

// SetPayableAmount sets field value
func (o *CreateRefundRequest) SetPayableAmount(v string) {
	o.PayableAmount = v
}

// GetToAddress returns the ToAddress field value if set, zero value otherwise.
func (o *CreateRefundRequest) GetToAddress() string {
	if o == nil || IsNil(o.ToAddress) {
		var ret string
		return ret
	}
	return *o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefundRequest) GetToAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ToAddress) {
		return nil, false
	}
	return o.ToAddress, true
}

// HasToAddress returns a boolean if a field has been set.
func (o *CreateRefundRequest) HasToAddress() bool {
	if o != nil && !IsNil(o.ToAddress) {
		return true
	}

	return false
}

// SetToAddress gets a reference to the given string and assigns it to the ToAddress field.
func (o *CreateRefundRequest) SetToAddress(v string) {
	o.ToAddress = &v
}

// GetTokenId returns the TokenId field value
func (o *CreateRefundRequest) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *CreateRefundRequest) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *CreateRefundRequest) SetTokenId(v string) {
	o.TokenId = v
}

// GetRefundType returns the RefundType field value
func (o *CreateRefundRequest) GetRefundType() RefundType {
	if o == nil {
		var ret RefundType
		return ret
	}

	return o.RefundType
}

// GetRefundTypeOk returns a tuple with the RefundType field value
// and a boolean to check if the value has been set.
func (o *CreateRefundRequest) GetRefundTypeOk() (*RefundType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefundType, true
}

// SetRefundType sets field value
func (o *CreateRefundRequest) SetRefundType(v RefundType) {
	o.RefundType = v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CreateRefundRequest) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefundRequest) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CreateRefundRequest) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *CreateRefundRequest) SetOrderId(v string) {
	o.OrderId = &v
}

// GetChargeMerchantFee returns the ChargeMerchantFee field value if set, zero value otherwise.
func (o *CreateRefundRequest) GetChargeMerchantFee() bool {
	if o == nil || IsNil(o.ChargeMerchantFee) {
		var ret bool
		return ret
	}
	return *o.ChargeMerchantFee
}

// GetChargeMerchantFeeOk returns a tuple with the ChargeMerchantFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefundRequest) GetChargeMerchantFeeOk() (*bool, bool) {
	if o == nil || IsNil(o.ChargeMerchantFee) {
		return nil, false
	}
	return o.ChargeMerchantFee, true
}

// HasChargeMerchantFee returns a boolean if a field has been set.
func (o *CreateRefundRequest) HasChargeMerchantFee() bool {
	if o != nil && !IsNil(o.ChargeMerchantFee) {
		return true
	}

	return false
}

// SetChargeMerchantFee gets a reference to the given bool and assigns it to the ChargeMerchantFee field.
func (o *CreateRefundRequest) SetChargeMerchantFee(v bool) {
	o.ChargeMerchantFee = &v
}

// GetMerchantFeeAmount returns the MerchantFeeAmount field value if set, zero value otherwise.
func (o *CreateRefundRequest) GetMerchantFeeAmount() string {
	if o == nil || IsNil(o.MerchantFeeAmount) {
		var ret string
		return ret
	}
	return *o.MerchantFeeAmount
}

// GetMerchantFeeAmountOk returns a tuple with the MerchantFeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefundRequest) GetMerchantFeeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantFeeAmount) {
		return nil, false
	}
	return o.MerchantFeeAmount, true
}

// HasMerchantFeeAmount returns a boolean if a field has been set.
func (o *CreateRefundRequest) HasMerchantFeeAmount() bool {
	if o != nil && !IsNil(o.MerchantFeeAmount) {
		return true
	}

	return false
}

// SetMerchantFeeAmount gets a reference to the given string and assigns it to the MerchantFeeAmount field.
func (o *CreateRefundRequest) SetMerchantFeeAmount(v string) {
	o.MerchantFeeAmount = &v
}

// GetMerchantFeeTokenId returns the MerchantFeeTokenId field value if set, zero value otherwise.
func (o *CreateRefundRequest) GetMerchantFeeTokenId() string {
	if o == nil || IsNil(o.MerchantFeeTokenId) {
		var ret string
		return ret
	}
	return *o.MerchantFeeTokenId
}

// GetMerchantFeeTokenIdOk returns a tuple with the MerchantFeeTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefundRequest) GetMerchantFeeTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantFeeTokenId) {
		return nil, false
	}
	return o.MerchantFeeTokenId, true
}

// HasMerchantFeeTokenId returns a boolean if a field has been set.
func (o *CreateRefundRequest) HasMerchantFeeTokenId() bool {
	if o != nil && !IsNil(o.MerchantFeeTokenId) {
		return true
	}

	return false
}

// SetMerchantFeeTokenId gets a reference to the given string and assigns it to the MerchantFeeTokenId field.
func (o *CreateRefundRequest) SetMerchantFeeTokenId(v string) {
	o.MerchantFeeTokenId = &v
}

func (o CreateRefundRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRefundRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	if !IsNil(o.MerchantId) {
		toSerialize["merchant_id"] = o.MerchantId
	}
	toSerialize["payable_amount"] = o.PayableAmount
	if !IsNil(o.ToAddress) {
		toSerialize["to_address"] = o.ToAddress
	}
	toSerialize["token_id"] = o.TokenId
	toSerialize["refund_type"] = o.RefundType
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.ChargeMerchantFee) {
		toSerialize["charge_merchant_fee"] = o.ChargeMerchantFee
	}
	if !IsNil(o.MerchantFeeAmount) {
		toSerialize["merchant_fee_amount"] = o.MerchantFeeAmount
	}
	if !IsNil(o.MerchantFeeTokenId) {
		toSerialize["merchant_fee_token_id"] = o.MerchantFeeTokenId
	}
	return toSerialize, nil
}

func (o *CreateRefundRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request_id",
		"payable_amount",
		"token_id",
		"refund_type",
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

	varCreateRefundRequest := _CreateRefundRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRefundRequest)

	if err != nil {
		return err
	}

	*o = CreateRefundRequest(varCreateRefundRequest)

	return err
}

type NullableCreateRefundRequest struct {
	value *CreateRefundRequest
	isSet bool
}

func (v NullableCreateRefundRequest) Get() *CreateRefundRequest {
	return v.value
}

func (v *NullableCreateRefundRequest) Set(val *CreateRefundRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRefundRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRefundRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRefundRequest(val *CreateRefundRequest) *NullableCreateRefundRequest {
	return &NullableCreateRefundRequest{value: val, isSet: true}
}

func (v NullableCreateRefundRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRefundRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


