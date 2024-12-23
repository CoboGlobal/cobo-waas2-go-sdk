/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the SwapActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwapActivity{}

// SwapActivity struct for SwapActivity
type SwapActivity struct {
	// The unique identifier of the swap activity.
	ActivityId *string `json:"activity_id,omitempty"`
	// The status of the swap activity.
	Status *string `json:"status,omitempty"`
	// The unique identifier of the wallet.
	WalletId *string `json:"wallet_id,omitempty"`
	// The token symbol to swap from.
	PayTokenId *string `json:"pay_token_id,omitempty"`
	// The token symbol to swap to.
	ReceiveTokenId *string `json:"receive_token_id,omitempty"`
	// The amount of tokens to bridge.
	PayAmount *string `json:"pay_amount,omitempty"`
	// The amount of tokens to receive.
	ReceiveAmount *string `json:"receive_amount,omitempty"`
	// The amount of fee.
	FeeAmount *string `json:"fee_amount,omitempty"`
	// The initiator of the swap activity.
	Initiator NullableString `json:"initiator,omitempty"`
	InitiatorType *TransactionInitiatorType `json:"initiator_type,omitempty"`
	// The time when the swap activity was created, in Unix timestamp format, measured in milliseconds.
	CreatedTimestamp *int32 `json:"created_timestamp,omitempty"`
	// The time when the swap activity was last updated, in Unix timestamp format, measured in milliseconds.
	UpdatedTimestamp *int32 `json:"updated_timestamp,omitempty"`
}

// NewSwapActivity instantiates a new SwapActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwapActivity() *SwapActivity {
	this := SwapActivity{}
	return &this
}

// NewSwapActivityWithDefaults instantiates a new SwapActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwapActivityWithDefaults() *SwapActivity {
	this := SwapActivity{}
	return &this
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *SwapActivity) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapActivity) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *SwapActivity) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *SwapActivity) SetActivityId(v string) {
	o.ActivityId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SwapActivity) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapActivity) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SwapActivity) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SwapActivity) SetStatus(v string) {
	o.Status = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *SwapActivity) GetWalletId() string {
	if o == nil || IsNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapActivity) GetWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.WalletId) {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *SwapActivity) HasWalletId() bool {
	if o != nil && !IsNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *SwapActivity) SetWalletId(v string) {
	o.WalletId = &v
}

// GetPayTokenId returns the PayTokenId field value if set, zero value otherwise.
func (o *SwapActivity) GetPayTokenId() string {
	if o == nil || IsNil(o.PayTokenId) {
		var ret string
		return ret
	}
	return *o.PayTokenId
}

// GetPayTokenIdOk returns a tuple with the PayTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapActivity) GetPayTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayTokenId) {
		return nil, false
	}
	return o.PayTokenId, true
}

// HasPayTokenId returns a boolean if a field has been set.
func (o *SwapActivity) HasPayTokenId() bool {
	if o != nil && !IsNil(o.PayTokenId) {
		return true
	}

	return false
}

// SetPayTokenId gets a reference to the given string and assigns it to the PayTokenId field.
func (o *SwapActivity) SetPayTokenId(v string) {
	o.PayTokenId = &v
}

// GetReceiveTokenId returns the ReceiveTokenId field value if set, zero value otherwise.
func (o *SwapActivity) GetReceiveTokenId() string {
	if o == nil || IsNil(o.ReceiveTokenId) {
		var ret string
		return ret
	}
	return *o.ReceiveTokenId
}

// GetReceiveTokenIdOk returns a tuple with the ReceiveTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapActivity) GetReceiveTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiveTokenId) {
		return nil, false
	}
	return o.ReceiveTokenId, true
}

// HasReceiveTokenId returns a boolean if a field has been set.
func (o *SwapActivity) HasReceiveTokenId() bool {
	if o != nil && !IsNil(o.ReceiveTokenId) {
		return true
	}

	return false
}

// SetReceiveTokenId gets a reference to the given string and assigns it to the ReceiveTokenId field.
func (o *SwapActivity) SetReceiveTokenId(v string) {
	o.ReceiveTokenId = &v
}

// GetPayAmount returns the PayAmount field value if set, zero value otherwise.
func (o *SwapActivity) GetPayAmount() string {
	if o == nil || IsNil(o.PayAmount) {
		var ret string
		return ret
	}
	return *o.PayAmount
}

// GetPayAmountOk returns a tuple with the PayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapActivity) GetPayAmountOk() (*string, bool) {
	if o == nil || IsNil(o.PayAmount) {
		return nil, false
	}
	return o.PayAmount, true
}

// HasPayAmount returns a boolean if a field has been set.
func (o *SwapActivity) HasPayAmount() bool {
	if o != nil && !IsNil(o.PayAmount) {
		return true
	}

	return false
}

// SetPayAmount gets a reference to the given string and assigns it to the PayAmount field.
func (o *SwapActivity) SetPayAmount(v string) {
	o.PayAmount = &v
}

// GetReceiveAmount returns the ReceiveAmount field value if set, zero value otherwise.
func (o *SwapActivity) GetReceiveAmount() string {
	if o == nil || IsNil(o.ReceiveAmount) {
		var ret string
		return ret
	}
	return *o.ReceiveAmount
}

// GetReceiveAmountOk returns a tuple with the ReceiveAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapActivity) GetReceiveAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiveAmount) {
		return nil, false
	}
	return o.ReceiveAmount, true
}

// HasReceiveAmount returns a boolean if a field has been set.
func (o *SwapActivity) HasReceiveAmount() bool {
	if o != nil && !IsNil(o.ReceiveAmount) {
		return true
	}

	return false
}

// SetReceiveAmount gets a reference to the given string and assigns it to the ReceiveAmount field.
func (o *SwapActivity) SetReceiveAmount(v string) {
	o.ReceiveAmount = &v
}

// GetFeeAmount returns the FeeAmount field value if set, zero value otherwise.
func (o *SwapActivity) GetFeeAmount() string {
	if o == nil || IsNil(o.FeeAmount) {
		var ret string
		return ret
	}
	return *o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapActivity) GetFeeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.FeeAmount) {
		return nil, false
	}
	return o.FeeAmount, true
}

// HasFeeAmount returns a boolean if a field has been set.
func (o *SwapActivity) HasFeeAmount() bool {
	if o != nil && !IsNil(o.FeeAmount) {
		return true
	}

	return false
}

// SetFeeAmount gets a reference to the given string and assigns it to the FeeAmount field.
func (o *SwapActivity) SetFeeAmount(v string) {
	o.FeeAmount = &v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SwapActivity) GetInitiator() string {
	if o == nil || IsNil(o.Initiator.Get()) {
		var ret string
		return ret
	}
	return *o.Initiator.Get()
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SwapActivity) GetInitiatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Initiator.Get(), o.Initiator.IsSet()
}

// HasInitiator returns a boolean if a field has been set.
func (o *SwapActivity) HasInitiator() bool {
	if o != nil && o.Initiator.IsSet() {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given NullableString and assigns it to the Initiator field.
func (o *SwapActivity) SetInitiator(v string) {
	o.Initiator.Set(&v)
}
// SetInitiatorNil sets the value for Initiator to be an explicit nil
func (o *SwapActivity) SetInitiatorNil() {
	o.Initiator.Set(nil)
}

// UnsetInitiator ensures that no value is present for Initiator, not even an explicit nil
func (o *SwapActivity) UnsetInitiator() {
	o.Initiator.Unset()
}

// GetInitiatorType returns the InitiatorType field value if set, zero value otherwise.
func (o *SwapActivity) GetInitiatorType() TransactionInitiatorType {
	if o == nil || IsNil(o.InitiatorType) {
		var ret TransactionInitiatorType
		return ret
	}
	return *o.InitiatorType
}

// GetInitiatorTypeOk returns a tuple with the InitiatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapActivity) GetInitiatorTypeOk() (*TransactionInitiatorType, bool) {
	if o == nil || IsNil(o.InitiatorType) {
		return nil, false
	}
	return o.InitiatorType, true
}

// HasInitiatorType returns a boolean if a field has been set.
func (o *SwapActivity) HasInitiatorType() bool {
	if o != nil && !IsNil(o.InitiatorType) {
		return true
	}

	return false
}

// SetInitiatorType gets a reference to the given TransactionInitiatorType and assigns it to the InitiatorType field.
func (o *SwapActivity) SetInitiatorType(v TransactionInitiatorType) {
	o.InitiatorType = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value if set, zero value otherwise.
func (o *SwapActivity) GetCreatedTimestamp() int32 {
	if o == nil || IsNil(o.CreatedTimestamp) {
		var ret int32
		return ret
	}
	return *o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapActivity) GetCreatedTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedTimestamp) {
		return nil, false
	}
	return o.CreatedTimestamp, true
}

// HasCreatedTimestamp returns a boolean if a field has been set.
func (o *SwapActivity) HasCreatedTimestamp() bool {
	if o != nil && !IsNil(o.CreatedTimestamp) {
		return true
	}

	return false
}

// SetCreatedTimestamp gets a reference to the given int32 and assigns it to the CreatedTimestamp field.
func (o *SwapActivity) SetCreatedTimestamp(v int32) {
	o.CreatedTimestamp = &v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value if set, zero value otherwise.
func (o *SwapActivity) GetUpdatedTimestamp() int32 {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		var ret int32
		return ret
	}
	return *o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapActivity) GetUpdatedTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		return nil, false
	}
	return o.UpdatedTimestamp, true
}

// HasUpdatedTimestamp returns a boolean if a field has been set.
func (o *SwapActivity) HasUpdatedTimestamp() bool {
	if o != nil && !IsNil(o.UpdatedTimestamp) {
		return true
	}

	return false
}

// SetUpdatedTimestamp gets a reference to the given int32 and assigns it to the UpdatedTimestamp field.
func (o *SwapActivity) SetUpdatedTimestamp(v int32) {
	o.UpdatedTimestamp = &v
}

func (o SwapActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwapActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.WalletId) {
		toSerialize["wallet_id"] = o.WalletId
	}
	if !IsNil(o.PayTokenId) {
		toSerialize["pay_token_id"] = o.PayTokenId
	}
	if !IsNil(o.ReceiveTokenId) {
		toSerialize["receive_token_id"] = o.ReceiveTokenId
	}
	if !IsNil(o.PayAmount) {
		toSerialize["pay_amount"] = o.PayAmount
	}
	if !IsNil(o.ReceiveAmount) {
		toSerialize["receive_amount"] = o.ReceiveAmount
	}
	if !IsNil(o.FeeAmount) {
		toSerialize["fee_amount"] = o.FeeAmount
	}
	if o.Initiator.IsSet() {
		toSerialize["initiator"] = o.Initiator.Get()
	}
	if !IsNil(o.InitiatorType) {
		toSerialize["initiator_type"] = o.InitiatorType
	}
	if !IsNil(o.CreatedTimestamp) {
		toSerialize["created_timestamp"] = o.CreatedTimestamp
	}
	if !IsNil(o.UpdatedTimestamp) {
		toSerialize["updated_timestamp"] = o.UpdatedTimestamp
	}
	return toSerialize, nil
}

type NullableSwapActivity struct {
	value *SwapActivity
	isSet bool
}

func (v NullableSwapActivity) Get() *SwapActivity {
	return v.value
}

func (v *NullableSwapActivity) Set(val *SwapActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableSwapActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableSwapActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwapActivity(val *SwapActivity) *NullableSwapActivity {
	return &NullableSwapActivity{value: val, isSet: true}
}

func (v NullableSwapActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwapActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


