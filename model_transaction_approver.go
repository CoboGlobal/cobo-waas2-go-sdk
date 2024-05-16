/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
)

// checks if the TransactionApprover type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionApprover{}

// TransactionApprover The approver data for transaction.
type TransactionApprover struct {
	// The approver name of the transaction.
	Name *string `json:"name,omitempty"`
	// The approval status.
	Status *string `json:"status,omitempty"`
}

// NewTransactionApprover instantiates a new TransactionApprover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionApprover() *TransactionApprover {
	this := TransactionApprover{}
	return &this
}

// NewTransactionApproverWithDefaults instantiates a new TransactionApprover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionApproverWithDefaults() *TransactionApprover {
	this := TransactionApprover{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TransactionApprover) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionApprover) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TransactionApprover) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TransactionApprover) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransactionApprover) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionApprover) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransactionApprover) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TransactionApprover) SetStatus(v string) {
	o.Status = &v
}

func (o TransactionApprover) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionApprover) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableTransactionApprover struct {
	value *TransactionApprover
	isSet bool
}

func (v NullableTransactionApprover) Get() *TransactionApprover {
	return v.value
}

func (v *NullableTransactionApprover) Set(val *TransactionApprover) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionApprover) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionApprover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionApprover(val *TransactionApprover) *NullableTransactionApprover {
	return &NullableTransactionApprover{value: val, isSet: true}
}

func (v NullableTransactionApprover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionApprover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


