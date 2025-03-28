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

// checks if the ContractCallParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractCallParams{}

// ContractCallParams The information about a transaction that interacts with a smart contract
type ContractCallParams struct {
	// The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization.
	RequestId string `json:"request_id"`
	// The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains).
	ChainId string `json:"chain_id"`
	Source ContractCallSource `json:"source"`
	Destination ContractCallDestination `json:"destination"`
	// The description of the contract call transaction.
	Description *string `json:"description,omitempty"`
	// The custom category for you to identify your transactions.
	CategoryNames []string `json:"category_names,omitempty"`
	Fee *TransactionRequestFee `json:"fee,omitempty"`
	TransactionProcessType *TransactionProcessType `json:"transaction_process_type,omitempty"`
	AutoFuel *AutoFuelType `json:"auto_fuel,omitempty"`
}

type _ContractCallParams ContractCallParams

// NewContractCallParams instantiates a new ContractCallParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractCallParams(requestId string, chainId string, source ContractCallSource, destination ContractCallDestination) *ContractCallParams {
	this := ContractCallParams{}
	this.RequestId = requestId
	this.ChainId = chainId
	this.Source = source
	this.Destination = destination
	return &this
}

// NewContractCallParamsWithDefaults instantiates a new ContractCallParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractCallParamsWithDefaults() *ContractCallParams {
	this := ContractCallParams{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *ContractCallParams) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ContractCallParams) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ContractCallParams) SetRequestId(v string) {
	o.RequestId = v
}

// GetChainId returns the ChainId field value
func (o *ContractCallParams) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *ContractCallParams) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *ContractCallParams) SetChainId(v string) {
	o.ChainId = v
}

// GetSource returns the Source field value
func (o *ContractCallParams) GetSource() ContractCallSource {
	if o == nil {
		var ret ContractCallSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ContractCallParams) GetSourceOk() (*ContractCallSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ContractCallParams) SetSource(v ContractCallSource) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *ContractCallParams) GetDestination() ContractCallDestination {
	if o == nil {
		var ret ContractCallDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *ContractCallParams) GetDestinationOk() (*ContractCallDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *ContractCallParams) SetDestination(v ContractCallDestination) {
	o.Destination = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ContractCallParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCallParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ContractCallParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ContractCallParams) SetDescription(v string) {
	o.Description = &v
}

// GetCategoryNames returns the CategoryNames field value if set, zero value otherwise.
func (o *ContractCallParams) GetCategoryNames() []string {
	if o == nil || IsNil(o.CategoryNames) {
		var ret []string
		return ret
	}
	return o.CategoryNames
}

// GetCategoryNamesOk returns a tuple with the CategoryNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCallParams) GetCategoryNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.CategoryNames) {
		return nil, false
	}
	return o.CategoryNames, true
}

// HasCategoryNames returns a boolean if a field has been set.
func (o *ContractCallParams) HasCategoryNames() bool {
	if o != nil && !IsNil(o.CategoryNames) {
		return true
	}

	return false
}

// SetCategoryNames gets a reference to the given []string and assigns it to the CategoryNames field.
func (o *ContractCallParams) SetCategoryNames(v []string) {
	o.CategoryNames = v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *ContractCallParams) GetFee() TransactionRequestFee {
	if o == nil || IsNil(o.Fee) {
		var ret TransactionRequestFee
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCallParams) GetFeeOk() (*TransactionRequestFee, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *ContractCallParams) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given TransactionRequestFee and assigns it to the Fee field.
func (o *ContractCallParams) SetFee(v TransactionRequestFee) {
	o.Fee = &v
}

// GetTransactionProcessType returns the TransactionProcessType field value if set, zero value otherwise.
func (o *ContractCallParams) GetTransactionProcessType() TransactionProcessType {
	if o == nil || IsNil(o.TransactionProcessType) {
		var ret TransactionProcessType
		return ret
	}
	return *o.TransactionProcessType
}

// GetTransactionProcessTypeOk returns a tuple with the TransactionProcessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCallParams) GetTransactionProcessTypeOk() (*TransactionProcessType, bool) {
	if o == nil || IsNil(o.TransactionProcessType) {
		return nil, false
	}
	return o.TransactionProcessType, true
}

// HasTransactionProcessType returns a boolean if a field has been set.
func (o *ContractCallParams) HasTransactionProcessType() bool {
	if o != nil && !IsNil(o.TransactionProcessType) {
		return true
	}

	return false
}

// SetTransactionProcessType gets a reference to the given TransactionProcessType and assigns it to the TransactionProcessType field.
func (o *ContractCallParams) SetTransactionProcessType(v TransactionProcessType) {
	o.TransactionProcessType = &v
}

// GetAutoFuel returns the AutoFuel field value if set, zero value otherwise.
func (o *ContractCallParams) GetAutoFuel() AutoFuelType {
	if o == nil || IsNil(o.AutoFuel) {
		var ret AutoFuelType
		return ret
	}
	return *o.AutoFuel
}

// GetAutoFuelOk returns a tuple with the AutoFuel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCallParams) GetAutoFuelOk() (*AutoFuelType, bool) {
	if o == nil || IsNil(o.AutoFuel) {
		return nil, false
	}
	return o.AutoFuel, true
}

// HasAutoFuel returns a boolean if a field has been set.
func (o *ContractCallParams) HasAutoFuel() bool {
	if o != nil && !IsNil(o.AutoFuel) {
		return true
	}

	return false
}

// SetAutoFuel gets a reference to the given AutoFuelType and assigns it to the AutoFuel field.
func (o *ContractCallParams) SetAutoFuel(v AutoFuelType) {
	o.AutoFuel = &v
}

func (o ContractCallParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractCallParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	toSerialize["chain_id"] = o.ChainId
	toSerialize["source"] = o.Source
	toSerialize["destination"] = o.Destination
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CategoryNames) {
		toSerialize["category_names"] = o.CategoryNames
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.TransactionProcessType) {
		toSerialize["transaction_process_type"] = o.TransactionProcessType
	}
	if !IsNil(o.AutoFuel) {
		toSerialize["auto_fuel"] = o.AutoFuel
	}
	return toSerialize, nil
}

func (o *ContractCallParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request_id",
		"chain_id",
		"source",
		"destination",
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

	varContractCallParams := _ContractCallParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContractCallParams)

	if err != nil {
		return err
	}

	*o = ContractCallParams(varContractCallParams)

	return err
}

type NullableContractCallParams struct {
	value *ContractCallParams
	isSet bool
}

func (v NullableContractCallParams) Get() *ContractCallParams {
	return v.value
}

func (v *NullableContractCallParams) Set(val *ContractCallParams) {
	v.value = val
	v.isSet = true
}

func (v NullableContractCallParams) IsSet() bool {
	return v.isSet
}

func (v *NullableContractCallParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractCallParams(val *ContractCallParams) *NullableContractCallParams {
	return &NullableContractCallParams{value: val, isSet: true}
}

func (v NullableContractCallParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractCallParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


