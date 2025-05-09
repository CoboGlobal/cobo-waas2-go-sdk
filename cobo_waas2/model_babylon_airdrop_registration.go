/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the BabylonAirdropRegistration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BabylonAirdropRegistration{}

// BabylonAirdropRegistration The details of a Babylon airdrop registration.
type BabylonAirdropRegistration struct {
	// The registration ID, a unique identifier for tracking the airdrop registration.
	Id *string `json:"id,omitempty"`
	Status *BabylonRegistrationRequestStatus `json:"status,omitempty"`
	BtcAddress *StakingSource `json:"btc_address,omitempty"`
	BabylonAddress *StakingSource `json:"babylon_address,omitempty"`
	// The actual airdrop amount allocated for this BTC address.
	AirdropAmount *string `json:"airdrop_amount,omitempty"`
	// The detailed error message if the registration failed.
	ErrorMessage *string `json:"error_message,omitempty"`
	// The time when the registration was created, in Unix timestamp format, measured in milliseconds.
	CreatedTimestamp *int64 `json:"created_timestamp,omitempty"`
	// The time when the registration was updated, in Unix timestamp format, measured in milliseconds.
	UpdatedTimestamp *int64 `json:"updated_timestamp,omitempty"`
}

// NewBabylonAirdropRegistration instantiates a new BabylonAirdropRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBabylonAirdropRegistration() *BabylonAirdropRegistration {
	this := BabylonAirdropRegistration{}
	return &this
}

// NewBabylonAirdropRegistrationWithDefaults instantiates a new BabylonAirdropRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBabylonAirdropRegistrationWithDefaults() *BabylonAirdropRegistration {
	this := BabylonAirdropRegistration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BabylonAirdropRegistration) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonAirdropRegistration) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BabylonAirdropRegistration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BabylonAirdropRegistration) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BabylonAirdropRegistration) GetStatus() BabylonRegistrationRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret BabylonRegistrationRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonAirdropRegistration) GetStatusOk() (*BabylonRegistrationRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BabylonAirdropRegistration) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BabylonRegistrationRequestStatus and assigns it to the Status field.
func (o *BabylonAirdropRegistration) SetStatus(v BabylonRegistrationRequestStatus) {
	o.Status = &v
}

// GetBtcAddress returns the BtcAddress field value if set, zero value otherwise.
func (o *BabylonAirdropRegistration) GetBtcAddress() StakingSource {
	if o == nil || IsNil(o.BtcAddress) {
		var ret StakingSource
		return ret
	}
	return *o.BtcAddress
}

// GetBtcAddressOk returns a tuple with the BtcAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonAirdropRegistration) GetBtcAddressOk() (*StakingSource, bool) {
	if o == nil || IsNil(o.BtcAddress) {
		return nil, false
	}
	return o.BtcAddress, true
}

// HasBtcAddress returns a boolean if a field has been set.
func (o *BabylonAirdropRegistration) HasBtcAddress() bool {
	if o != nil && !IsNil(o.BtcAddress) {
		return true
	}

	return false
}

// SetBtcAddress gets a reference to the given StakingSource and assigns it to the BtcAddress field.
func (o *BabylonAirdropRegistration) SetBtcAddress(v StakingSource) {
	o.BtcAddress = &v
}

// GetBabylonAddress returns the BabylonAddress field value if set, zero value otherwise.
func (o *BabylonAirdropRegistration) GetBabylonAddress() StakingSource {
	if o == nil || IsNil(o.BabylonAddress) {
		var ret StakingSource
		return ret
	}
	return *o.BabylonAddress
}

// GetBabylonAddressOk returns a tuple with the BabylonAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonAirdropRegistration) GetBabylonAddressOk() (*StakingSource, bool) {
	if o == nil || IsNil(o.BabylonAddress) {
		return nil, false
	}
	return o.BabylonAddress, true
}

// HasBabylonAddress returns a boolean if a field has been set.
func (o *BabylonAirdropRegistration) HasBabylonAddress() bool {
	if o != nil && !IsNil(o.BabylonAddress) {
		return true
	}

	return false
}

// SetBabylonAddress gets a reference to the given StakingSource and assigns it to the BabylonAddress field.
func (o *BabylonAirdropRegistration) SetBabylonAddress(v StakingSource) {
	o.BabylonAddress = &v
}

// GetAirdropAmount returns the AirdropAmount field value if set, zero value otherwise.
func (o *BabylonAirdropRegistration) GetAirdropAmount() string {
	if o == nil || IsNil(o.AirdropAmount) {
		var ret string
		return ret
	}
	return *o.AirdropAmount
}

// GetAirdropAmountOk returns a tuple with the AirdropAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonAirdropRegistration) GetAirdropAmountOk() (*string, bool) {
	if o == nil || IsNil(o.AirdropAmount) {
		return nil, false
	}
	return o.AirdropAmount, true
}

// HasAirdropAmount returns a boolean if a field has been set.
func (o *BabylonAirdropRegistration) HasAirdropAmount() bool {
	if o != nil && !IsNil(o.AirdropAmount) {
		return true
	}

	return false
}

// SetAirdropAmount gets a reference to the given string and assigns it to the AirdropAmount field.
func (o *BabylonAirdropRegistration) SetAirdropAmount(v string) {
	o.AirdropAmount = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *BabylonAirdropRegistration) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonAirdropRegistration) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *BabylonAirdropRegistration) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *BabylonAirdropRegistration) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value if set, zero value otherwise.
func (o *BabylonAirdropRegistration) GetCreatedTimestamp() int64 {
	if o == nil || IsNil(o.CreatedTimestamp) {
		var ret int64
		return ret
	}
	return *o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonAirdropRegistration) GetCreatedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedTimestamp) {
		return nil, false
	}
	return o.CreatedTimestamp, true
}

// HasCreatedTimestamp returns a boolean if a field has been set.
func (o *BabylonAirdropRegistration) HasCreatedTimestamp() bool {
	if o != nil && !IsNil(o.CreatedTimestamp) {
		return true
	}

	return false
}

// SetCreatedTimestamp gets a reference to the given int64 and assigns it to the CreatedTimestamp field.
func (o *BabylonAirdropRegistration) SetCreatedTimestamp(v int64) {
	o.CreatedTimestamp = &v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value if set, zero value otherwise.
func (o *BabylonAirdropRegistration) GetUpdatedTimestamp() int64 {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		var ret int64
		return ret
	}
	return *o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonAirdropRegistration) GetUpdatedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedTimestamp) {
		return nil, false
	}
	return o.UpdatedTimestamp, true
}

// HasUpdatedTimestamp returns a boolean if a field has been set.
func (o *BabylonAirdropRegistration) HasUpdatedTimestamp() bool {
	if o != nil && !IsNil(o.UpdatedTimestamp) {
		return true
	}

	return false
}

// SetUpdatedTimestamp gets a reference to the given int64 and assigns it to the UpdatedTimestamp field.
func (o *BabylonAirdropRegistration) SetUpdatedTimestamp(v int64) {
	o.UpdatedTimestamp = &v
}

func (o BabylonAirdropRegistration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BabylonAirdropRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.BtcAddress) {
		toSerialize["btc_address"] = o.BtcAddress
	}
	if !IsNil(o.BabylonAddress) {
		toSerialize["babylon_address"] = o.BabylonAddress
	}
	if !IsNil(o.AirdropAmount) {
		toSerialize["airdrop_amount"] = o.AirdropAmount
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if !IsNil(o.CreatedTimestamp) {
		toSerialize["created_timestamp"] = o.CreatedTimestamp
	}
	if !IsNil(o.UpdatedTimestamp) {
		toSerialize["updated_timestamp"] = o.UpdatedTimestamp
	}
	return toSerialize, nil
}

type NullableBabylonAirdropRegistration struct {
	value *BabylonAirdropRegistration
	isSet bool
}

func (v NullableBabylonAirdropRegistration) Get() *BabylonAirdropRegistration {
	return v.value
}

func (v *NullableBabylonAirdropRegistration) Set(val *BabylonAirdropRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableBabylonAirdropRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableBabylonAirdropRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBabylonAirdropRegistration(val *BabylonAirdropRegistration) *NullableBabylonAirdropRegistration {
	return &NullableBabylonAirdropRegistration{value: val, isSet: true}
}

func (v NullableBabylonAirdropRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBabylonAirdropRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


