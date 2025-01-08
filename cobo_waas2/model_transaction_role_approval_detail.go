/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the TransactionRoleApprovalDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionRoleApprovalDetail{}

// TransactionRoleApprovalDetail The role approval data for transaction.
type TransactionRoleApprovalDetail struct {
	Result *TransactionApprovalResult `json:"result,omitempty"`
	// The threshold for passing the review of this role.
	ReviewThreshold *int32 `json:"review_threshold,omitempty"`
	// The initiator of the transaction.
	Initiator *string `json:"initiator,omitempty"`
	// Time to complete the review.
	CompleteTime *string `json:"complete_time,omitempty"`
	UserDetails []TransactionUserApprovalDetail `json:"user_details,omitempty"`
}

// NewTransactionRoleApprovalDetail instantiates a new TransactionRoleApprovalDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRoleApprovalDetail() *TransactionRoleApprovalDetail {
	this := TransactionRoleApprovalDetail{}
	return &this
}

// NewTransactionRoleApprovalDetailWithDefaults instantiates a new TransactionRoleApprovalDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRoleApprovalDetailWithDefaults() *TransactionRoleApprovalDetail {
	this := TransactionRoleApprovalDetail{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *TransactionRoleApprovalDetail) GetResult() TransactionApprovalResult {
	if o == nil || IsNil(o.Result) {
		var ret TransactionApprovalResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRoleApprovalDetail) GetResultOk() (*TransactionApprovalResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *TransactionRoleApprovalDetail) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given TransactionApprovalResult and assigns it to the Result field.
func (o *TransactionRoleApprovalDetail) SetResult(v TransactionApprovalResult) {
	o.Result = &v
}

// GetReviewThreshold returns the ReviewThreshold field value if set, zero value otherwise.
func (o *TransactionRoleApprovalDetail) GetReviewThreshold() int32 {
	if o == nil || IsNil(o.ReviewThreshold) {
		var ret int32
		return ret
	}
	return *o.ReviewThreshold
}

// GetReviewThresholdOk returns a tuple with the ReviewThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRoleApprovalDetail) GetReviewThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.ReviewThreshold) {
		return nil, false
	}
	return o.ReviewThreshold, true
}

// HasReviewThreshold returns a boolean if a field has been set.
func (o *TransactionRoleApprovalDetail) HasReviewThreshold() bool {
	if o != nil && !IsNil(o.ReviewThreshold) {
		return true
	}

	return false
}

// SetReviewThreshold gets a reference to the given int32 and assigns it to the ReviewThreshold field.
func (o *TransactionRoleApprovalDetail) SetReviewThreshold(v int32) {
	o.ReviewThreshold = &v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *TransactionRoleApprovalDetail) GetInitiator() string {
	if o == nil || IsNil(o.Initiator) {
		var ret string
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRoleApprovalDetail) GetInitiatorOk() (*string, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *TransactionRoleApprovalDetail) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given string and assigns it to the Initiator field.
func (o *TransactionRoleApprovalDetail) SetInitiator(v string) {
	o.Initiator = &v
}

// GetCompleteTime returns the CompleteTime field value if set, zero value otherwise.
func (o *TransactionRoleApprovalDetail) GetCompleteTime() string {
	if o == nil || IsNil(o.CompleteTime) {
		var ret string
		return ret
	}
	return *o.CompleteTime
}

// GetCompleteTimeOk returns a tuple with the CompleteTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRoleApprovalDetail) GetCompleteTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CompleteTime) {
		return nil, false
	}
	return o.CompleteTime, true
}

// HasCompleteTime returns a boolean if a field has been set.
func (o *TransactionRoleApprovalDetail) HasCompleteTime() bool {
	if o != nil && !IsNil(o.CompleteTime) {
		return true
	}

	return false
}

// SetCompleteTime gets a reference to the given string and assigns it to the CompleteTime field.
func (o *TransactionRoleApprovalDetail) SetCompleteTime(v string) {
	o.CompleteTime = &v
}

// GetUserDetails returns the UserDetails field value if set, zero value otherwise.
func (o *TransactionRoleApprovalDetail) GetUserDetails() []TransactionUserApprovalDetail {
	if o == nil || IsNil(o.UserDetails) {
		var ret []TransactionUserApprovalDetail
		return ret
	}
	return o.UserDetails
}

// GetUserDetailsOk returns a tuple with the UserDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRoleApprovalDetail) GetUserDetailsOk() ([]TransactionUserApprovalDetail, bool) {
	if o == nil || IsNil(o.UserDetails) {
		return nil, false
	}
	return o.UserDetails, true
}

// HasUserDetails returns a boolean if a field has been set.
func (o *TransactionRoleApprovalDetail) HasUserDetails() bool {
	if o != nil && !IsNil(o.UserDetails) {
		return true
	}

	return false
}

// SetUserDetails gets a reference to the given []TransactionUserApprovalDetail and assigns it to the UserDetails field.
func (o *TransactionRoleApprovalDetail) SetUserDetails(v []TransactionUserApprovalDetail) {
	o.UserDetails = v
}

func (o TransactionRoleApprovalDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRoleApprovalDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.ReviewThreshold) {
		toSerialize["review_threshold"] = o.ReviewThreshold
	}
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	if !IsNil(o.CompleteTime) {
		toSerialize["complete_time"] = o.CompleteTime
	}
	if !IsNil(o.UserDetails) {
		toSerialize["user_details"] = o.UserDetails
	}
	return toSerialize, nil
}

type NullableTransactionRoleApprovalDetail struct {
	value *TransactionRoleApprovalDetail
	isSet bool
}

func (v NullableTransactionRoleApprovalDetail) Get() *TransactionRoleApprovalDetail {
	return v.value
}

func (v *NullableTransactionRoleApprovalDetail) Set(val *TransactionRoleApprovalDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRoleApprovalDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRoleApprovalDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRoleApprovalDetail(val *TransactionRoleApprovalDetail) *NullableTransactionRoleApprovalDetail {
	return &NullableTransactionRoleApprovalDetail{value: val, isSet: true}
}

func (v NullableTransactionRoleApprovalDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRoleApprovalDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


