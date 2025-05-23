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

// checks if the TSSKeyGenEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TSSKeyGenEventData{}

// TSSKeyGenEventData struct for TSSKeyGenEventData
type TSSKeyGenEventData struct {
	DataType TSSEventDataType `json:"data_type"`
	// The request ID.
	RequestId *string `json:"request_id,omitempty"`
	RequestType *TSSRequestTypeEenum `json:"request_type,omitempty"`
	RequestStatus *TSSStatus `json:"request_status,omitempty"`
	// The extra info.
	ExtraInfo *string `json:"extra_info,omitempty"`
	// The failed reason.
	FailedReason *string `json:"failed_reason,omitempty"`
	RequestDetail *TSSKeyGenRequest `json:"request_detail,omitempty"`
	Result *TSSGroup `json:"result,omitempty"`
}

type _TSSKeyGenEventData TSSKeyGenEventData

// NewTSSKeyGenEventData instantiates a new TSSKeyGenEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTSSKeyGenEventData(dataType TSSEventDataType) *TSSKeyGenEventData {
	this := TSSKeyGenEventData{}
	this.DataType = dataType
	return &this
}

// NewTSSKeyGenEventDataWithDefaults instantiates a new TSSKeyGenEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTSSKeyGenEventDataWithDefaults() *TSSKeyGenEventData {
	this := TSSKeyGenEventData{}
	return &this
}

// GetDataType returns the DataType field value
func (o *TSSKeyGenEventData) GetDataType() TSSEventDataType {
	if o == nil {
		var ret TSSEventDataType
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *TSSKeyGenEventData) GetDataTypeOk() (*TSSEventDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *TSSKeyGenEventData) SetDataType(v TSSEventDataType) {
	o.DataType = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *TSSKeyGenEventData) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSKeyGenEventData) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *TSSKeyGenEventData) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *TSSKeyGenEventData) SetRequestId(v string) {
	o.RequestId = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *TSSKeyGenEventData) GetRequestType() TSSRequestTypeEenum {
	if o == nil || IsNil(o.RequestType) {
		var ret TSSRequestTypeEenum
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSKeyGenEventData) GetRequestTypeOk() (*TSSRequestTypeEenum, bool) {
	if o == nil || IsNil(o.RequestType) {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *TSSKeyGenEventData) HasRequestType() bool {
	if o != nil && !IsNil(o.RequestType) {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given TSSRequestTypeEenum and assigns it to the RequestType field.
func (o *TSSKeyGenEventData) SetRequestType(v TSSRequestTypeEenum) {
	o.RequestType = &v
}

// GetRequestStatus returns the RequestStatus field value if set, zero value otherwise.
func (o *TSSKeyGenEventData) GetRequestStatus() TSSStatus {
	if o == nil || IsNil(o.RequestStatus) {
		var ret TSSStatus
		return ret
	}
	return *o.RequestStatus
}

// GetRequestStatusOk returns a tuple with the RequestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSKeyGenEventData) GetRequestStatusOk() (*TSSStatus, bool) {
	if o == nil || IsNil(o.RequestStatus) {
		return nil, false
	}
	return o.RequestStatus, true
}

// HasRequestStatus returns a boolean if a field has been set.
func (o *TSSKeyGenEventData) HasRequestStatus() bool {
	if o != nil && !IsNil(o.RequestStatus) {
		return true
	}

	return false
}

// SetRequestStatus gets a reference to the given TSSStatus and assigns it to the RequestStatus field.
func (o *TSSKeyGenEventData) SetRequestStatus(v TSSStatus) {
	o.RequestStatus = &v
}

// GetExtraInfo returns the ExtraInfo field value if set, zero value otherwise.
func (o *TSSKeyGenEventData) GetExtraInfo() string {
	if o == nil || IsNil(o.ExtraInfo) {
		var ret string
		return ret
	}
	return *o.ExtraInfo
}

// GetExtraInfoOk returns a tuple with the ExtraInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSKeyGenEventData) GetExtraInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ExtraInfo) {
		return nil, false
	}
	return o.ExtraInfo, true
}

// HasExtraInfo returns a boolean if a field has been set.
func (o *TSSKeyGenEventData) HasExtraInfo() bool {
	if o != nil && !IsNil(o.ExtraInfo) {
		return true
	}

	return false
}

// SetExtraInfo gets a reference to the given string and assigns it to the ExtraInfo field.
func (o *TSSKeyGenEventData) SetExtraInfo(v string) {
	o.ExtraInfo = &v
}

// GetFailedReason returns the FailedReason field value if set, zero value otherwise.
func (o *TSSKeyGenEventData) GetFailedReason() string {
	if o == nil || IsNil(o.FailedReason) {
		var ret string
		return ret
	}
	return *o.FailedReason
}

// GetFailedReasonOk returns a tuple with the FailedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSKeyGenEventData) GetFailedReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailedReason) {
		return nil, false
	}
	return o.FailedReason, true
}

// HasFailedReason returns a boolean if a field has been set.
func (o *TSSKeyGenEventData) HasFailedReason() bool {
	if o != nil && !IsNil(o.FailedReason) {
		return true
	}

	return false
}

// SetFailedReason gets a reference to the given string and assigns it to the FailedReason field.
func (o *TSSKeyGenEventData) SetFailedReason(v string) {
	o.FailedReason = &v
}

// GetRequestDetail returns the RequestDetail field value if set, zero value otherwise.
func (o *TSSKeyGenEventData) GetRequestDetail() TSSKeyGenRequest {
	if o == nil || IsNil(o.RequestDetail) {
		var ret TSSKeyGenRequest
		return ret
	}
	return *o.RequestDetail
}

// GetRequestDetailOk returns a tuple with the RequestDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSKeyGenEventData) GetRequestDetailOk() (*TSSKeyGenRequest, bool) {
	if o == nil || IsNil(o.RequestDetail) {
		return nil, false
	}
	return o.RequestDetail, true
}

// HasRequestDetail returns a boolean if a field has been set.
func (o *TSSKeyGenEventData) HasRequestDetail() bool {
	if o != nil && !IsNil(o.RequestDetail) {
		return true
	}

	return false
}

// SetRequestDetail gets a reference to the given TSSKeyGenRequest and assigns it to the RequestDetail field.
func (o *TSSKeyGenEventData) SetRequestDetail(v TSSKeyGenRequest) {
	o.RequestDetail = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *TSSKeyGenEventData) GetResult() TSSGroup {
	if o == nil || IsNil(o.Result) {
		var ret TSSGroup
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSKeyGenEventData) GetResultOk() (*TSSGroup, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *TSSKeyGenEventData) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given TSSGroup and assigns it to the Result field.
func (o *TSSKeyGenEventData) SetResult(v TSSGroup) {
	o.Result = &v
}

func (o TSSKeyGenEventData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TSSKeyGenEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data_type"] = o.DataType
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.RequestType) {
		toSerialize["request_type"] = o.RequestType
	}
	if !IsNil(o.RequestStatus) {
		toSerialize["request_status"] = o.RequestStatus
	}
	if !IsNil(o.ExtraInfo) {
		toSerialize["extra_info"] = o.ExtraInfo
	}
	if !IsNil(o.FailedReason) {
		toSerialize["failed_reason"] = o.FailedReason
	}
	if !IsNil(o.RequestDetail) {
		toSerialize["request_detail"] = o.RequestDetail
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

func (o *TSSKeyGenEventData) UnmarshalJSON(data []byte) (err error) {
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

	varTSSKeyGenEventData := _TSSKeyGenEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTSSKeyGenEventData)

	if err != nil {
		return err
	}

	*o = TSSKeyGenEventData(varTSSKeyGenEventData)

	return err
}

type NullableTSSKeyGenEventData struct {
	value *TSSKeyGenEventData
	isSet bool
}

func (v NullableTSSKeyGenEventData) Get() *TSSKeyGenEventData {
	return v.value
}

func (v *NullableTSSKeyGenEventData) Set(val *TSSKeyGenEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableTSSKeyGenEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableTSSKeyGenEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTSSKeyGenEventData(val *TSSKeyGenEventData) *NullableTSSKeyGenEventData {
	return &NullableTSSKeyGenEventData{value: val, isSet: true}
}

func (v NullableTSSKeyGenEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTSSKeyGenEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


