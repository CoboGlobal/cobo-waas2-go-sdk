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

// checks if the SafeContractCallSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SafeContractCallSource{}

// SafeContractCallSource struct for SafeContractCallSource
type SafeContractCallSource struct {
	// The type of the wallet. Possible values include: - `Org-Controlled`: MPC Wallets (Organization-Controlled). - `User-Controlled`: MPC Wallets (User-Controlled). - `Safe{Wallet}`: Smart Contract Wallets (Safe{Wallet}). 
	SourceType string `json:"source_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	// The wallet address.
	Address string `json:"address"`
	Delegate SafeContractCallSourceAllOfDelegate `json:"delegate"`
}

type _SafeContractCallSource SafeContractCallSource

// NewSafeContractCallSource instantiates a new SafeContractCallSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSafeContractCallSource(sourceType string, walletId string, address string, delegate SafeContractCallSourceAllOfDelegate) *SafeContractCallSource {
	this := SafeContractCallSource{}
	this.SourceType = sourceType
	this.WalletId = walletId
	this.Address = address
	this.Delegate = delegate
	return &this
}

// NewSafeContractCallSourceWithDefaults instantiates a new SafeContractCallSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSafeContractCallSourceWithDefaults() *SafeContractCallSource {
	this := SafeContractCallSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *SafeContractCallSource) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *SafeContractCallSource) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *SafeContractCallSource) SetSourceType(v string) {
	o.SourceType = v
}

// GetWalletId returns the WalletId field value
func (o *SafeContractCallSource) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *SafeContractCallSource) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *SafeContractCallSource) SetWalletId(v string) {
	o.WalletId = v
}

// GetAddress returns the Address field value
func (o *SafeContractCallSource) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *SafeContractCallSource) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *SafeContractCallSource) SetAddress(v string) {
	o.Address = v
}

// GetDelegate returns the Delegate field value
func (o *SafeContractCallSource) GetDelegate() SafeContractCallSourceAllOfDelegate {
	if o == nil {
		var ret SafeContractCallSourceAllOfDelegate
		return ret
	}

	return o.Delegate
}

// GetDelegateOk returns a tuple with the Delegate field value
// and a boolean to check if the value has been set.
func (o *SafeContractCallSource) GetDelegateOk() (*SafeContractCallSourceAllOfDelegate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delegate, true
}

// SetDelegate sets field value
func (o *SafeContractCallSource) SetDelegate(v SafeContractCallSourceAllOfDelegate) {
	o.Delegate = v
}

func (o SafeContractCallSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SafeContractCallSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["address"] = o.Address
	toSerialize["delegate"] = o.Delegate
	return toSerialize, nil
}

func (o *SafeContractCallSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_type",
		"wallet_id",
		"address",
		"delegate",
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

	varSafeContractCallSource := _SafeContractCallSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSafeContractCallSource)

	if err != nil {
		return err
	}

	*o = SafeContractCallSource(varSafeContractCallSource)

	return err
}

type NullableSafeContractCallSource struct {
	value *SafeContractCallSource
	isSet bool
}

func (v NullableSafeContractCallSource) Get() *SafeContractCallSource {
	return v.value
}

func (v *NullableSafeContractCallSource) Set(val *SafeContractCallSource) {
	v.value = val
	v.isSet = true
}

func (v NullableSafeContractCallSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSafeContractCallSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSafeContractCallSource(val *SafeContractCallSource) *NullableSafeContractCallSource {
	return &NullableSafeContractCallSource{value: val, isSet: true}
}

func (v NullableSafeContractCallSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSafeContractCallSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


