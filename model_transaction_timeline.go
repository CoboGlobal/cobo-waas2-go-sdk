/*
Cobo Wallet as a Service 2.0

Cobo WaaS 2.0 enables you to programmatically access Cobo's full suite of crypto wallet technologies with powerful and flexible access controls.  # Wallet technologies - Custodial Wallet - MPC Wallet - Smart Contract Wallet (Based on Safe{Wallet}) - Exchange Wallet  # Risk Control technologies - Workflow - Access Control List (ACL)  # Risk Control targets - Wallet Management   - User/team and their permission management   - Risk control configurations, e.g. whitelist, blacklist, rate-limiting etc. - Blockchain Interaction   - Crypto transfer   - Smart Contract Invocation  # Important HTTPS only. RESTful, resource oriented  # Get Started Set up your APIs or get authorization  # Authentication and Authorization CoboAuth  # Request and Response application/json  # Error Handling  ### Common error codes | Error Code | Description | | -- | -- |  ### API-specific error codes For error codes that are dedicated to a specific API, see the Error codes section in each API specification, for example, /v3/wallets.  # Rate and Usage Limiting  # Idempotent Request  # Pagination # Support [Developer Hub](https://cobo.com/developers) 

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
)

// checks if the TransactionTimeline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionTimeline{}

// TransactionTimeline The data for transaction timeline information.
type TransactionTimeline struct {
	Status *TransactionStatus `json:"status,omitempty"`
	// Whether the timeline status finished
	Finished *bool `json:"finished,omitempty"`
	// Timeline status finished time
	FinishedTime *float32 `json:"finished_time,omitempty"`
}

// NewTransactionTimeline instantiates a new TransactionTimeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionTimeline() *TransactionTimeline {
	this := TransactionTimeline{}
	return &this
}

// NewTransactionTimelineWithDefaults instantiates a new TransactionTimeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionTimelineWithDefaults() *TransactionTimeline {
	this := TransactionTimeline{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransactionTimeline) GetStatus() TransactionStatus {
	if o == nil || IsNil(o.Status) {
		var ret TransactionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTimeline) GetStatusOk() (*TransactionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransactionTimeline) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TransactionStatus and assigns it to the Status field.
func (o *TransactionTimeline) SetStatus(v TransactionStatus) {
	o.Status = &v
}

// GetFinished returns the Finished field value if set, zero value otherwise.
func (o *TransactionTimeline) GetFinished() bool {
	if o == nil || IsNil(o.Finished) {
		var ret bool
		return ret
	}
	return *o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTimeline) GetFinishedOk() (*bool, bool) {
	if o == nil || IsNil(o.Finished) {
		return nil, false
	}
	return o.Finished, true
}

// HasFinished returns a boolean if a field has been set.
func (o *TransactionTimeline) HasFinished() bool {
	if o != nil && !IsNil(o.Finished) {
		return true
	}

	return false
}

// SetFinished gets a reference to the given bool and assigns it to the Finished field.
func (o *TransactionTimeline) SetFinished(v bool) {
	o.Finished = &v
}

// GetFinishedTime returns the FinishedTime field value if set, zero value otherwise.
func (o *TransactionTimeline) GetFinishedTime() float32 {
	if o == nil || IsNil(o.FinishedTime) {
		var ret float32
		return ret
	}
	return *o.FinishedTime
}

// GetFinishedTimeOk returns a tuple with the FinishedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTimeline) GetFinishedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.FinishedTime) {
		return nil, false
	}
	return o.FinishedTime, true
}

// HasFinishedTime returns a boolean if a field has been set.
func (o *TransactionTimeline) HasFinishedTime() bool {
	if o != nil && !IsNil(o.FinishedTime) {
		return true
	}

	return false
}

// SetFinishedTime gets a reference to the given float32 and assigns it to the FinishedTime field.
func (o *TransactionTimeline) SetFinishedTime(v float32) {
	o.FinishedTime = &v
}

func (o TransactionTimeline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionTimeline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Finished) {
		toSerialize["finished"] = o.Finished
	}
	if !IsNil(o.FinishedTime) {
		toSerialize["finished_time"] = o.FinishedTime
	}
	return toSerialize, nil
}

type NullableTransactionTimeline struct {
	value *TransactionTimeline
	isSet bool
}

func (v NullableTransactionTimeline) Get() *TransactionTimeline {
	return v.value
}

func (v *NullableTransactionTimeline) Set(val *TransactionTimeline) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionTimeline) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionTimeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionTimeline(val *TransactionTimeline) *NullableTransactionTimeline {
	return &NullableTransactionTimeline{value: val, isSet: true}
}

func (v NullableTransactionTimeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionTimeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


