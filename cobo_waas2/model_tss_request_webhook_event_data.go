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

// checks if the TSSRequestWebhookEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TSSRequestWebhookEventData{}

// TSSRequestWebhookEventData struct for TSSRequestWebhookEventData
type TSSRequestWebhookEventData struct {
	//  The data type of the event. - `Transaction`: The transaction event data. - `TSSRequest`: The TSS request event data. - `Addresses`: The addresses event data. - `WalletInfo`: The wallet information event data. - `MPCVault`: The MPC vault event data.
	DataType string `json:"data_type"`
	// The TSS request ID.
	TssRequestId *string `json:"tss_request_id,omitempty"`
	SourceKeyShareHolderGroup *SourceGroup `json:"source_key_share_holder_group,omitempty"`
	// The target key share holder group ID.
	TargetKeyShareHolderGroupId *string `json:"target_key_share_holder_group_id,omitempty"`
	Type *TSSRequestType `json:"type,omitempty"`
	Status *TSSRequestStatus `json:"status,omitempty"`
	// The description of the TSS request.
	Description *string `json:"description,omitempty"`
	// The TSS request's creation time in Unix timestamp format, measured in milliseconds.
	CreatedTimestamp *int64 `json:"created_timestamp,omitempty"`
}

type _TSSRequestWebhookEventData TSSRequestWebhookEventData

// NewTSSRequestWebhookEventData instantiates a new TSSRequestWebhookEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTSSRequestWebhookEventData(dataType string) *TSSRequestWebhookEventData {
	this := TSSRequestWebhookEventData{}
	this.DataType = dataType
	return &this
}

// NewTSSRequestWebhookEventDataWithDefaults instantiates a new TSSRequestWebhookEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTSSRequestWebhookEventDataWithDefaults() *TSSRequestWebhookEventData {
	this := TSSRequestWebhookEventData{}
	return &this
}

// GetDataType returns the DataType field value
func (o *TSSRequestWebhookEventData) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *TSSRequestWebhookEventData) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *TSSRequestWebhookEventData) SetDataType(v string) {
	o.DataType = v
}

// GetTssRequestId returns the TssRequestId field value if set, zero value otherwise.
func (o *TSSRequestWebhookEventData) GetTssRequestId() string {
	if o == nil || IsNil(o.TssRequestId) {
		var ret string
		return ret
	}
	return *o.TssRequestId
}

// GetTssRequestIdOk returns a tuple with the TssRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSRequestWebhookEventData) GetTssRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.TssRequestId) {
		return nil, false
	}
	return o.TssRequestId, true
}

// HasTssRequestId returns a boolean if a field has been set.
func (o *TSSRequestWebhookEventData) HasTssRequestId() bool {
	if o != nil && !IsNil(o.TssRequestId) {
		return true
	}

	return false
}

// SetTssRequestId gets a reference to the given string and assigns it to the TssRequestId field.
func (o *TSSRequestWebhookEventData) SetTssRequestId(v string) {
	o.TssRequestId = &v
}

// GetSourceKeyShareHolderGroup returns the SourceKeyShareHolderGroup field value if set, zero value otherwise.
func (o *TSSRequestWebhookEventData) GetSourceKeyShareHolderGroup() SourceGroup {
	if o == nil || IsNil(o.SourceKeyShareHolderGroup) {
		var ret SourceGroup
		return ret
	}
	return *o.SourceKeyShareHolderGroup
}

// GetSourceKeyShareHolderGroupOk returns a tuple with the SourceKeyShareHolderGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSRequestWebhookEventData) GetSourceKeyShareHolderGroupOk() (*SourceGroup, bool) {
	if o == nil || IsNil(o.SourceKeyShareHolderGroup) {
		return nil, false
	}
	return o.SourceKeyShareHolderGroup, true
}

// HasSourceKeyShareHolderGroup returns a boolean if a field has been set.
func (o *TSSRequestWebhookEventData) HasSourceKeyShareHolderGroup() bool {
	if o != nil && !IsNil(o.SourceKeyShareHolderGroup) {
		return true
	}

	return false
}

// SetSourceKeyShareHolderGroup gets a reference to the given SourceGroup and assigns it to the SourceKeyShareHolderGroup field.
func (o *TSSRequestWebhookEventData) SetSourceKeyShareHolderGroup(v SourceGroup) {
	o.SourceKeyShareHolderGroup = &v
}

// GetTargetKeyShareHolderGroupId returns the TargetKeyShareHolderGroupId field value if set, zero value otherwise.
func (o *TSSRequestWebhookEventData) GetTargetKeyShareHolderGroupId() string {
	if o == nil || IsNil(o.TargetKeyShareHolderGroupId) {
		var ret string
		return ret
	}
	return *o.TargetKeyShareHolderGroupId
}

// GetTargetKeyShareHolderGroupIdOk returns a tuple with the TargetKeyShareHolderGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSRequestWebhookEventData) GetTargetKeyShareHolderGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetKeyShareHolderGroupId) {
		return nil, false
	}
	return o.TargetKeyShareHolderGroupId, true
}

// HasTargetKeyShareHolderGroupId returns a boolean if a field has been set.
func (o *TSSRequestWebhookEventData) HasTargetKeyShareHolderGroupId() bool {
	if o != nil && !IsNil(o.TargetKeyShareHolderGroupId) {
		return true
	}

	return false
}

// SetTargetKeyShareHolderGroupId gets a reference to the given string and assigns it to the TargetKeyShareHolderGroupId field.
func (o *TSSRequestWebhookEventData) SetTargetKeyShareHolderGroupId(v string) {
	o.TargetKeyShareHolderGroupId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TSSRequestWebhookEventData) GetType() TSSRequestType {
	if o == nil || IsNil(o.Type) {
		var ret TSSRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSRequestWebhookEventData) GetTypeOk() (*TSSRequestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TSSRequestWebhookEventData) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TSSRequestType and assigns it to the Type field.
func (o *TSSRequestWebhookEventData) SetType(v TSSRequestType) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TSSRequestWebhookEventData) GetStatus() TSSRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret TSSRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSRequestWebhookEventData) GetStatusOk() (*TSSRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TSSRequestWebhookEventData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TSSRequestStatus and assigns it to the Status field.
func (o *TSSRequestWebhookEventData) SetStatus(v TSSRequestStatus) {
	o.Status = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TSSRequestWebhookEventData) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSRequestWebhookEventData) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TSSRequestWebhookEventData) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TSSRequestWebhookEventData) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value if set, zero value otherwise.
func (o *TSSRequestWebhookEventData) GetCreatedTimestamp() int64 {
	if o == nil || IsNil(o.CreatedTimestamp) {
		var ret int64
		return ret
	}
	return *o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSRequestWebhookEventData) GetCreatedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedTimestamp) {
		return nil, false
	}
	return o.CreatedTimestamp, true
}

// HasCreatedTimestamp returns a boolean if a field has been set.
func (o *TSSRequestWebhookEventData) HasCreatedTimestamp() bool {
	if o != nil && !IsNil(o.CreatedTimestamp) {
		return true
	}

	return false
}

// SetCreatedTimestamp gets a reference to the given int64 and assigns it to the CreatedTimestamp field.
func (o *TSSRequestWebhookEventData) SetCreatedTimestamp(v int64) {
	o.CreatedTimestamp = &v
}

func (o TSSRequestWebhookEventData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TSSRequestWebhookEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data_type"] = o.DataType
	if !IsNil(o.TssRequestId) {
		toSerialize["tss_request_id"] = o.TssRequestId
	}
	if !IsNil(o.SourceKeyShareHolderGroup) {
		toSerialize["source_key_share_holder_group"] = o.SourceKeyShareHolderGroup
	}
	if !IsNil(o.TargetKeyShareHolderGroupId) {
		toSerialize["target_key_share_holder_group_id"] = o.TargetKeyShareHolderGroupId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CreatedTimestamp) {
		toSerialize["created_timestamp"] = o.CreatedTimestamp
	}
	return toSerialize, nil
}

func (o *TSSRequestWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data_type",
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

	varTSSRequestWebhookEventData := _TSSRequestWebhookEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTSSRequestWebhookEventData)

	if err != nil {
		return err
	}

	*o = TSSRequestWebhookEventData(varTSSRequestWebhookEventData)

	return err
}

type NullableTSSRequestWebhookEventData struct {
	value *TSSRequestWebhookEventData
	isSet bool
}

func (v NullableTSSRequestWebhookEventData) Get() *TSSRequestWebhookEventData {
	return v.value
}

func (v *NullableTSSRequestWebhookEventData) Set(val *TSSRequestWebhookEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableTSSRequestWebhookEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableTSSRequestWebhookEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTSSRequestWebhookEventData(val *TSSRequestWebhookEventData) *NullableTSSRequestWebhookEventData {
	return &NullableTSSRequestWebhookEventData{value: val, isSet: true}
}

func (v NullableTSSRequestWebhookEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTSSRequestWebhookEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


