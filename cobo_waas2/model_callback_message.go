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

// checks if the CallbackMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallbackMessage{}

// CallbackMessage The information about a callback message.
type CallbackMessage struct {
	// The callback message ID.
	Id string `json:"id"`
	// The time when the callback message was created, in Unix timestamp format, measured in milliseconds.
	CreatedTimestamp int64 `json:"created_timestamp"`
	// The time when the callback message was updated, in Unix timestamp format, measured in milliseconds.
	UpdatedTimestamp int64 `json:"updated_timestamp"`
	// The request ID of the transaction.
	RequestId string `json:"request_id"`
	// The transaction ID.
	TransactionId string `json:"transaction_id"`
	// The wallet ID.
	WalletId *string `json:"wallet_id,omitempty"`
	// The callback endpoint URL.
	Url string `json:"url"`
	Data Transaction `json:"data"`
	// The callback message status.
	Status string `json:"status"`
	// The callback message result.
	Result *string `json:"result,omitempty"`
}

type _CallbackMessage CallbackMessage

// NewCallbackMessage instantiates a new CallbackMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallbackMessage(id string, createdTimestamp int64, updatedTimestamp int64, requestId string, transactionId string, url string, data Transaction, status string) *CallbackMessage {
	this := CallbackMessage{}
	this.Id = id
	this.CreatedTimestamp = createdTimestamp
	this.UpdatedTimestamp = updatedTimestamp
	this.RequestId = requestId
	this.TransactionId = transactionId
	this.Url = url
	this.Data = data
	this.Status = status
	return &this
}

// NewCallbackMessageWithDefaults instantiates a new CallbackMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallbackMessageWithDefaults() *CallbackMessage {
	this := CallbackMessage{}
	return &this
}

// GetId returns the Id field value
func (o *CallbackMessage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CallbackMessage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CallbackMessage) SetId(v string) {
	o.Id = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *CallbackMessage) GetCreatedTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *CallbackMessage) GetCreatedTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *CallbackMessage) SetCreatedTimestamp(v int64) {
	o.CreatedTimestamp = v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *CallbackMessage) GetUpdatedTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *CallbackMessage) GetUpdatedTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *CallbackMessage) SetUpdatedTimestamp(v int64) {
	o.UpdatedTimestamp = v
}

// GetRequestId returns the RequestId field value
func (o *CallbackMessage) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CallbackMessage) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CallbackMessage) SetRequestId(v string) {
	o.RequestId = v
}

// GetTransactionId returns the TransactionId field value
func (o *CallbackMessage) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *CallbackMessage) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *CallbackMessage) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *CallbackMessage) GetWalletId() string {
	if o == nil || IsNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallbackMessage) GetWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.WalletId) {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *CallbackMessage) HasWalletId() bool {
	if o != nil && !IsNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *CallbackMessage) SetWalletId(v string) {
	o.WalletId = &v
}

// GetUrl returns the Url field value
func (o *CallbackMessage) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CallbackMessage) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CallbackMessage) SetUrl(v string) {
	o.Url = v
}

// GetData returns the Data field value
func (o *CallbackMessage) GetData() Transaction {
	if o == nil {
		var ret Transaction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CallbackMessage) GetDataOk() (*Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CallbackMessage) SetData(v Transaction) {
	o.Data = v
}

// GetStatus returns the Status field value
func (o *CallbackMessage) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CallbackMessage) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CallbackMessage) SetStatus(v string) {
	o.Status = v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CallbackMessage) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallbackMessage) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CallbackMessage) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *CallbackMessage) SetResult(v string) {
	o.Result = &v
}

func (o CallbackMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallbackMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_timestamp"] = o.CreatedTimestamp
	toSerialize["updated_timestamp"] = o.UpdatedTimestamp
	toSerialize["request_id"] = o.RequestId
	toSerialize["transaction_id"] = o.TransactionId
	if !IsNil(o.WalletId) {
		toSerialize["wallet_id"] = o.WalletId
	}
	toSerialize["url"] = o.Url
	toSerialize["data"] = o.Data
	toSerialize["status"] = o.Status
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

func (o *CallbackMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_timestamp",
		"updated_timestamp",
		"request_id",
		"transaction_id",
		"url",
		"data",
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

	varCallbackMessage := _CallbackMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCallbackMessage)

	if err != nil {
		return err
	}

	*o = CallbackMessage(varCallbackMessage)

	return err
}

type NullableCallbackMessage struct {
	value *CallbackMessage
	isSet bool
}

func (v NullableCallbackMessage) Get() *CallbackMessage {
	return v.value
}

func (v *NullableCallbackMessage) Set(val *CallbackMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableCallbackMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableCallbackMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallbackMessage(val *CallbackMessage) *NullableCallbackMessage {
	return &NullableCallbackMessage{value: val, isSet: true}
}

func (v NullableCallbackMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallbackMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


