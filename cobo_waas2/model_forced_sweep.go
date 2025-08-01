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

// checks if the ForcedSweep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForcedSweep{}

// ForcedSweep struct for ForcedSweep
type ForcedSweep struct {
	// The force sweep ID.
	ForcedSweepId string `json:"forced_sweep_id"`
	// The request ID provided by you when creating the force sweep request.
	RequestId string `json:"request_id"`
	// The wallet ID to force sweep, which is the unique identifier of a wallet.
	WalletId *string `json:"wallet_id,omitempty"`
	// The token ID to force sweep, which is the unique identifier of a token.
	TokenId *string `json:"token_id,omitempty"`
	// The amount of needing force sweep.
	Amount *string `json:"amount,omitempty"`
	Status ForcedSweepStatus `json:"status"`
	// The created time of the force sweep request, represented as a UNIX timestamp in seconds.
	CreatedTimestamp *int32 `json:"created_timestamp,omitempty"`
	// The updated time of the force sweep request, represented as a UNIX timestamp in seconds.
	UpdatedTimestamp *int32 `json:"updated_timestamp,omitempty"`
}

type _ForcedSweep ForcedSweep

// NewForcedSweep instantiates a new ForcedSweep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForcedSweep(forcedSweepId string, requestId string, status ForcedSweepStatus) *ForcedSweep {
	this := ForcedSweep{}
	this.ForcedSweepId = forcedSweepId
	this.RequestId = requestId
	this.Status = status
	return &this
}

// NewForcedSweepWithDefaults instantiates a new ForcedSweep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForcedSweepWithDefaults() *ForcedSweep {
	this := ForcedSweep{}
	return &this
}

// GetForcedSweepId returns the ForcedSweepId field value
func (o *ForcedSweep) GetForcedSweepId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForcedSweepId
}

// GetForcedSweepIdOk returns a tuple with the ForcedSweepId field value
// and a boolean to check if the value has been set.
func (o *ForcedSweep) GetForcedSweepIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForcedSweepId, true
}

// SetForcedSweepId sets field value
func (o *ForcedSweep) SetForcedSweepId(v string) {
	o.ForcedSweepId = v
}

// GetRequestId returns the RequestId field value
func (o *ForcedSweep) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ForcedSweep) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ForcedSweep) SetRequestId(v string) {
	o.RequestId = v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *ForcedSweep) GetWalletId() string {
	if o == nil || IsNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForcedSweep) GetWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.WalletId) {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *ForcedSweep) HasWalletId() bool {
	if o != nil && !IsNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *ForcedSweep) SetWalletId(v string) {
	o.WalletId = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *ForcedSweep) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForcedSweep) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *ForcedSweep) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *ForcedSweep) SetTokenId(v string) {
	o.TokenId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ForcedSweep) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForcedSweep) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ForcedSweep) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *ForcedSweep) SetAmount(v string) {
	o.Amount = &v
}

// GetStatus returns the Status field value
func (o *ForcedSweep) GetStatus() ForcedSweepStatus {
	if o == nil {
		var ret ForcedSweepStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ForcedSweep) GetStatusOk() (*ForcedSweepStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ForcedSweep) SetStatus(v ForcedSweepStatus) {
	o.Status = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value if set, zero value otherwise.
func (o *ForcedSweep) GetCreatedTimestamp() int32 {
	if o == nil || IsNil(o.CreatedTimestamp) {
		var ret int32
		return ret
	}
	return *o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForcedSweep) GetCreatedTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedTimestamp) {
		return nil, false
	}
	return o.CreatedTimestamp, true
}

// HasCreatedTimestamp returns a boolean if a field has been set.
func (o *ForcedSweep) HasCreatedTimestamp() bool {
	if o != nil && !IsNil(o.CreatedTimestamp) {
		return true
	}

	return false
}

// SetCreatedTimestamp gets a reference to the given int32 and assigns it to the CreatedTimestamp field.
func (o *ForcedSweep) SetCreatedTimestamp(v int32) {
	o.CreatedTimestamp = &v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value if set, zero value otherwise.
func (o *ForcedSweep) GetUpdatedTimestamp() int32 {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		var ret int32
		return ret
	}
	return *o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForcedSweep) GetUpdatedTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		return nil, false
	}
	return o.UpdatedTimestamp, true
}

// HasUpdatedTimestamp returns a boolean if a field has been set.
func (o *ForcedSweep) HasUpdatedTimestamp() bool {
	if o != nil && !IsNil(o.UpdatedTimestamp) {
		return true
	}

	return false
}

// SetUpdatedTimestamp gets a reference to the given int32 and assigns it to the UpdatedTimestamp field.
func (o *ForcedSweep) SetUpdatedTimestamp(v int32) {
	o.UpdatedTimestamp = &v
}

func (o ForcedSweep) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForcedSweep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["forced_sweep_id"] = o.ForcedSweepId
	toSerialize["request_id"] = o.RequestId
	if !IsNil(o.WalletId) {
		toSerialize["wallet_id"] = o.WalletId
	}
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.CreatedTimestamp) {
		toSerialize["created_timestamp"] = o.CreatedTimestamp
	}
	if !IsNil(o.UpdatedTimestamp) {
		toSerialize["updated_timestamp"] = o.UpdatedTimestamp
	}
	return toSerialize, nil
}

func (o *ForcedSweep) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"forced_sweep_id",
		"request_id",
		"status",
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

	varForcedSweep := _ForcedSweep{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varForcedSweep)

	if err != nil {
		return err
	}

	*o = ForcedSweep(varForcedSweep)

	return err
}

type NullableForcedSweep struct {
	value *ForcedSweep
	isSet bool
}

func (v NullableForcedSweep) Get() *ForcedSweep {
	return v.value
}

func (v *NullableForcedSweep) Set(val *ForcedSweep) {
	v.value = val
	v.isSet = true
}

func (v NullableForcedSweep) IsSet() bool {
	return v.isSet
}

func (v *NullableForcedSweep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForcedSweep(val *ForcedSweep) *NullableForcedSweep {
	return &NullableForcedSweep{value: val, isSet: true}
}

func (v NullableForcedSweep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForcedSweep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


