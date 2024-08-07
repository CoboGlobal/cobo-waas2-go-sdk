/*
Cobo Wallet as a Service 2.0

Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TransactionRbf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionRbf{}

// TransactionRbf The information about the request to drop or to speed up transactions.
type TransactionRbf struct {
	// The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization.
	RequestId string `json:"request_id"`
	Fee TransactionRequestFee `json:"fee"`
	Source *TransactionRbfSource `json:"source,omitempty"`
	// The custom category for you to identify your transactions.
	CategoryNames []string `json:"category_names,omitempty"`
	// The description of the RBF transaction.
	Description *string `json:"description,omitempty"`
}

type _TransactionRbf TransactionRbf

// NewTransactionRbf instantiates a new TransactionRbf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRbf(requestId string, fee TransactionRequestFee) *TransactionRbf {
	this := TransactionRbf{}
	this.RequestId = requestId
	this.Fee = fee
	return &this
}

// NewTransactionRbfWithDefaults instantiates a new TransactionRbf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRbfWithDefaults() *TransactionRbf {
	this := TransactionRbf{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *TransactionRbf) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransactionRbf) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransactionRbf) SetRequestId(v string) {
	o.RequestId = v
}

// GetFee returns the Fee field value
func (o *TransactionRbf) GetFee() TransactionRequestFee {
	if o == nil {
		var ret TransactionRequestFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *TransactionRbf) GetFeeOk() (*TransactionRequestFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *TransactionRbf) SetFee(v TransactionRequestFee) {
	o.Fee = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *TransactionRbf) GetSource() TransactionRbfSource {
	if o == nil || IsNil(o.Source) {
		var ret TransactionRbfSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRbf) GetSourceOk() (*TransactionRbfSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *TransactionRbf) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given TransactionRbfSource and assigns it to the Source field.
func (o *TransactionRbf) SetSource(v TransactionRbfSource) {
	o.Source = &v
}

// GetCategoryNames returns the CategoryNames field value if set, zero value otherwise.
func (o *TransactionRbf) GetCategoryNames() []string {
	if o == nil || IsNil(o.CategoryNames) {
		var ret []string
		return ret
	}
	return o.CategoryNames
}

// GetCategoryNamesOk returns a tuple with the CategoryNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRbf) GetCategoryNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.CategoryNames) {
		return nil, false
	}
	return o.CategoryNames, true
}

// HasCategoryNames returns a boolean if a field has been set.
func (o *TransactionRbf) HasCategoryNames() bool {
	if o != nil && !IsNil(o.CategoryNames) {
		return true
	}

	return false
}

// SetCategoryNames gets a reference to the given []string and assigns it to the CategoryNames field.
func (o *TransactionRbf) SetCategoryNames(v []string) {
	o.CategoryNames = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransactionRbf) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRbf) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransactionRbf) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransactionRbf) SetDescription(v string) {
	o.Description = &v
}

func (o TransactionRbf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRbf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	toSerialize["fee"] = o.Fee
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.CategoryNames) {
		toSerialize["category_names"] = o.CategoryNames
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *TransactionRbf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request_id",
		"fee",
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

	varTransactionRbf := _TransactionRbf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionRbf)

	if err != nil {
		return err
	}

	*o = TransactionRbf(varTransactionRbf)

	return err
}

type NullableTransactionRbf struct {
	value *TransactionRbf
	isSet bool
}

func (v NullableTransactionRbf) Get() *TransactionRbf {
	return v.value
}

func (v *NullableTransactionRbf) Set(val *TransactionRbf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRbf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRbf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRbf(val *TransactionRbf) *NullableTransactionRbf {
	return &NullableTransactionRbf{value: val, isSet: true}
}

func (v NullableTransactionRbf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRbf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


