/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SmartContractCall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartContractCall{}

// SmartContractCall The data for create smart contract call transaction.
type SmartContractCall struct {
	// Unique id of the request.
	RequestId string `json:"request_id"`
	RequestType string `json:"request_type"`
	// Unique id of the wallet to transfer from.
	FromWalletId string `json:"from_wallet_id"`
	// From address
	FromAddressStr string `json:"from_address_str"`
	// The blockchain on which the token operates.
	ChainId string `json:"chain_id"`
	// To address
	ToAddressStr string `json:"to_address_str"`
	// Transaction value (Note that this is an absolute value. If you trade 1.5 ETH, then the value is 1.5) 
	Value *string `json:"value,omitempty"`
	// calldata for this transaction. Commonly used as part of contract interaction. 
	Calldata string `json:"calldata"`
	MpcUsedKeyGroup *MpcSigningGroup `json:"mpc_used_key_group,omitempty"`
	Fee *TransactionFee `json:"fee,omitempty"`
}

type _SmartContractCall SmartContractCall

// NewSmartContractCall instantiates a new SmartContractCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartContractCall(requestId string, requestType string, fromWalletId string, fromAddressStr string, chainId string, toAddressStr string, calldata string) *SmartContractCall {
	this := SmartContractCall{}
	this.RequestId = requestId
	this.RequestType = requestType
	this.FromWalletId = fromWalletId
	this.FromAddressStr = fromAddressStr
	this.ChainId = chainId
	this.ToAddressStr = toAddressStr
	this.Calldata = calldata
	return &this
}

// NewSmartContractCallWithDefaults instantiates a new SmartContractCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartContractCallWithDefaults() *SmartContractCall {
	this := SmartContractCall{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *SmartContractCall) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SmartContractCall) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SmartContractCall) SetRequestId(v string) {
	o.RequestId = v
}

// GetRequestType returns the RequestType field value
func (o *SmartContractCall) GetRequestType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *SmartContractCall) GetRequestTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value
func (o *SmartContractCall) SetRequestType(v string) {
	o.RequestType = v
}

// GetFromWalletId returns the FromWalletId field value
func (o *SmartContractCall) GetFromWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromWalletId
}

// GetFromWalletIdOk returns a tuple with the FromWalletId field value
// and a boolean to check if the value has been set.
func (o *SmartContractCall) GetFromWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromWalletId, true
}

// SetFromWalletId sets field value
func (o *SmartContractCall) SetFromWalletId(v string) {
	o.FromWalletId = v
}

// GetFromAddressStr returns the FromAddressStr field value
func (o *SmartContractCall) GetFromAddressStr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAddressStr
}

// GetFromAddressStrOk returns a tuple with the FromAddressStr field value
// and a boolean to check if the value has been set.
func (o *SmartContractCall) GetFromAddressStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAddressStr, true
}

// SetFromAddressStr sets field value
func (o *SmartContractCall) SetFromAddressStr(v string) {
	o.FromAddressStr = v
}

// GetChainId returns the ChainId field value
func (o *SmartContractCall) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *SmartContractCall) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *SmartContractCall) SetChainId(v string) {
	o.ChainId = v
}

// GetToAddressStr returns the ToAddressStr field value
func (o *SmartContractCall) GetToAddressStr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAddressStr
}

// GetToAddressStrOk returns a tuple with the ToAddressStr field value
// and a boolean to check if the value has been set.
func (o *SmartContractCall) GetToAddressStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddressStr, true
}

// SetToAddressStr sets field value
func (o *SmartContractCall) SetToAddressStr(v string) {
	o.ToAddressStr = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SmartContractCall) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractCall) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SmartContractCall) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SmartContractCall) SetValue(v string) {
	o.Value = &v
}

// GetCalldata returns the Calldata field value
func (o *SmartContractCall) GetCalldata() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Calldata
}

// GetCalldataOk returns a tuple with the Calldata field value
// and a boolean to check if the value has been set.
func (o *SmartContractCall) GetCalldataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Calldata, true
}

// SetCalldata sets field value
func (o *SmartContractCall) SetCalldata(v string) {
	o.Calldata = v
}

// GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field value if set, zero value otherwise.
func (o *SmartContractCall) GetMpcUsedKeyGroup() MpcSigningGroup {
	if o == nil || IsNil(o.MpcUsedKeyGroup) {
		var ret MpcSigningGroup
		return ret
	}
	return *o.MpcUsedKeyGroup
}

// GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractCall) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool) {
	if o == nil || IsNil(o.MpcUsedKeyGroup) {
		return nil, false
	}
	return o.MpcUsedKeyGroup, true
}

// HasMpcUsedKeyGroup returns a boolean if a field has been set.
func (o *SmartContractCall) HasMpcUsedKeyGroup() bool {
	if o != nil && !IsNil(o.MpcUsedKeyGroup) {
		return true
	}

	return false
}

// SetMpcUsedKeyGroup gets a reference to the given MpcSigningGroup and assigns it to the MpcUsedKeyGroup field.
func (o *SmartContractCall) SetMpcUsedKeyGroup(v MpcSigningGroup) {
	o.MpcUsedKeyGroup = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *SmartContractCall) GetFee() TransactionFee {
	if o == nil || IsNil(o.Fee) {
		var ret TransactionFee
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartContractCall) GetFeeOk() (*TransactionFee, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *SmartContractCall) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given TransactionFee and assigns it to the Fee field.
func (o *SmartContractCall) SetFee(v TransactionFee) {
	o.Fee = &v
}

func (o SmartContractCall) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartContractCall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	toSerialize["request_type"] = o.RequestType
	toSerialize["from_wallet_id"] = o.FromWalletId
	toSerialize["from_address_str"] = o.FromAddressStr
	toSerialize["chain_id"] = o.ChainId
	toSerialize["to_address_str"] = o.ToAddressStr
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	toSerialize["calldata"] = o.Calldata
	if !IsNil(o.MpcUsedKeyGroup) {
		toSerialize["mpc_used_key_group"] = o.MpcUsedKeyGroup
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	return toSerialize, nil
}

func (o *SmartContractCall) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request_id",
		"request_type",
		"from_wallet_id",
		"from_address_str",
		"chain_id",
		"to_address_str",
		"calldata",
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

	varSmartContractCall := _SmartContractCall{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartContractCall)

	if err != nil {
		return err
	}

	*o = SmartContractCall(varSmartContractCall)

	return err
}

type NullableSmartContractCall struct {
	value *SmartContractCall
	isSet bool
}

func (v NullableSmartContractCall) Get() *SmartContractCall {
	return v.value
}

func (v *NullableSmartContractCall) Set(val *SmartContractCall) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContractCall) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContractCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContractCall(val *SmartContractCall) *NullableSmartContractCall {
	return &NullableSmartContractCall{value: val, isSet: true}
}

func (v NullableSmartContractCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContractCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


