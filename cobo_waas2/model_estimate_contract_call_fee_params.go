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

// checks if the EstimateContractCallFeeParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EstimateContractCallFeeParams{}

// EstimateContractCallFeeParams The information about a transaction that interacts with a smart contract
type EstimateContractCallFeeParams struct {
	// The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. It is recommended to use the same request ID as the transaction for which you want to estimate the transaction fee.
	RequestId *string `json:"request_id,omitempty"`
	RequestType EstimateFeeRequestType `json:"request_type"`
	// The chain ID of the chain on which the smart contract is deployed. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains).
	ChainId string `json:"chain_id"`
	Source ContractCallSource `json:"source"`
	Destination ContractCallDestination `json:"destination"`
	FeeType *FeeType `json:"fee_type,omitempty"`
}

type _EstimateContractCallFeeParams EstimateContractCallFeeParams

// NewEstimateContractCallFeeParams instantiates a new EstimateContractCallFeeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimateContractCallFeeParams(requestType EstimateFeeRequestType, chainId string, source ContractCallSource, destination ContractCallDestination) *EstimateContractCallFeeParams {
	this := EstimateContractCallFeeParams{}
	this.RequestType = requestType
	this.ChainId = chainId
	this.Source = source
	this.Destination = destination
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = &feeType
	return &this
}

// NewEstimateContractCallFeeParamsWithDefaults instantiates a new EstimateContractCallFeeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimateContractCallFeeParamsWithDefaults() *EstimateContractCallFeeParams {
	this := EstimateContractCallFeeParams{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = &feeType
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *EstimateContractCallFeeParams) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimateContractCallFeeParams) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *EstimateContractCallFeeParams) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *EstimateContractCallFeeParams) SetRequestId(v string) {
	o.RequestId = &v
}

// GetRequestType returns the RequestType field value
func (o *EstimateContractCallFeeParams) GetRequestType() EstimateFeeRequestType {
	if o == nil {
		var ret EstimateFeeRequestType
		return ret
	}

	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *EstimateContractCallFeeParams) GetRequestTypeOk() (*EstimateFeeRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value
func (o *EstimateContractCallFeeParams) SetRequestType(v EstimateFeeRequestType) {
	o.RequestType = v
}

// GetChainId returns the ChainId field value
func (o *EstimateContractCallFeeParams) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *EstimateContractCallFeeParams) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *EstimateContractCallFeeParams) SetChainId(v string) {
	o.ChainId = v
}

// GetSource returns the Source field value
func (o *EstimateContractCallFeeParams) GetSource() ContractCallSource {
	if o == nil {
		var ret ContractCallSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *EstimateContractCallFeeParams) GetSourceOk() (*ContractCallSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *EstimateContractCallFeeParams) SetSource(v ContractCallSource) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *EstimateContractCallFeeParams) GetDestination() ContractCallDestination {
	if o == nil {
		var ret ContractCallDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *EstimateContractCallFeeParams) GetDestinationOk() (*ContractCallDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *EstimateContractCallFeeParams) SetDestination(v ContractCallDestination) {
	o.Destination = v
}

// GetFeeType returns the FeeType field value if set, zero value otherwise.
func (o *EstimateContractCallFeeParams) GetFeeType() FeeType {
	if o == nil || IsNil(o.FeeType) {
		var ret FeeType
		return ret
	}
	return *o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimateContractCallFeeParams) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil || IsNil(o.FeeType) {
		return nil, false
	}
	return o.FeeType, true
}

// HasFeeType returns a boolean if a field has been set.
func (o *EstimateContractCallFeeParams) HasFeeType() bool {
	if o != nil && !IsNil(o.FeeType) {
		return true
	}

	return false
}

// SetFeeType gets a reference to the given FeeType and assigns it to the FeeType field.
func (o *EstimateContractCallFeeParams) SetFeeType(v FeeType) {
	o.FeeType = &v
}

func (o EstimateContractCallFeeParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstimateContractCallFeeParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	toSerialize["request_type"] = o.RequestType
	toSerialize["chain_id"] = o.ChainId
	toSerialize["source"] = o.Source
	toSerialize["destination"] = o.Destination
	if !IsNil(o.FeeType) {
		toSerialize["fee_type"] = o.FeeType
	}
	return toSerialize, nil
}

func (o *EstimateContractCallFeeParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request_type",
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

	varEstimateContractCallFeeParams := _EstimateContractCallFeeParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEstimateContractCallFeeParams)

	if err != nil {
		return err
	}

	*o = EstimateContractCallFeeParams(varEstimateContractCallFeeParams)

	return err
}

type NullableEstimateContractCallFeeParams struct {
	value *EstimateContractCallFeeParams
	isSet bool
}

func (v NullableEstimateContractCallFeeParams) Get() *EstimateContractCallFeeParams {
	return v.value
}

func (v *NullableEstimateContractCallFeeParams) Set(val *EstimateContractCallFeeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimateContractCallFeeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimateContractCallFeeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimateContractCallFeeParams(val *EstimateContractCallFeeParams) *NullableEstimateContractCallFeeParams {
	return &NullableEstimateContractCallFeeParams{value: val, isSet: true}
}

func (v NullableEstimateContractCallFeeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimateContractCallFeeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


