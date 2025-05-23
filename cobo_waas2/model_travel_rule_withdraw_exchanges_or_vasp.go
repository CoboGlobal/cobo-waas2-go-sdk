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

// checks if the TravelRuleWithdrawExchangesOrVASP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TravelRuleWithdrawExchangesOrVASP{}

// TravelRuleWithdrawExchangesOrVASP Required information when withdrawing tokens to an exchange or other virtual asset service providers (VASP).
type TravelRuleWithdrawExchangesOrVASP struct {
	DestinationWalletType DestinationWalletType `json:"destination_wallet_type"`
	// The vendor code of the exchange or virtual asset service provider (VASP).
	VendorCode string `json:"vendor_code"`
	// The unique identifier of the VASP.
	VendorVaspId string `json:"vendor_vasp_id"`
	EntityInfo TravelRuleWithdrawExchangesOrVASPEntityInfo `json:"entity_info"`
}

type _TravelRuleWithdrawExchangesOrVASP TravelRuleWithdrawExchangesOrVASP

// NewTravelRuleWithdrawExchangesOrVASP instantiates a new TravelRuleWithdrawExchangesOrVASP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTravelRuleWithdrawExchangesOrVASP(destinationWalletType DestinationWalletType, vendorCode string, vendorVaspId string, entityInfo TravelRuleWithdrawExchangesOrVASPEntityInfo) *TravelRuleWithdrawExchangesOrVASP {
	this := TravelRuleWithdrawExchangesOrVASP{}
	this.DestinationWalletType = destinationWalletType
	this.VendorCode = vendorCode
	this.VendorVaspId = vendorVaspId
	this.EntityInfo = entityInfo
	return &this
}

// NewTravelRuleWithdrawExchangesOrVASPWithDefaults instantiates a new TravelRuleWithdrawExchangesOrVASP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTravelRuleWithdrawExchangesOrVASPWithDefaults() *TravelRuleWithdrawExchangesOrVASP {
	this := TravelRuleWithdrawExchangesOrVASP{}
	return &this
}

// GetDestinationWalletType returns the DestinationWalletType field value
func (o *TravelRuleWithdrawExchangesOrVASP) GetDestinationWalletType() DestinationWalletType {
	if o == nil {
		var ret DestinationWalletType
		return ret
	}

	return o.DestinationWalletType
}

// GetDestinationWalletTypeOk returns a tuple with the DestinationWalletType field value
// and a boolean to check if the value has been set.
func (o *TravelRuleWithdrawExchangesOrVASP) GetDestinationWalletTypeOk() (*DestinationWalletType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationWalletType, true
}

// SetDestinationWalletType sets field value
func (o *TravelRuleWithdrawExchangesOrVASP) SetDestinationWalletType(v DestinationWalletType) {
	o.DestinationWalletType = v
}

// GetVendorCode returns the VendorCode field value
func (o *TravelRuleWithdrawExchangesOrVASP) GetVendorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorCode
}

// GetVendorCodeOk returns a tuple with the VendorCode field value
// and a boolean to check if the value has been set.
func (o *TravelRuleWithdrawExchangesOrVASP) GetVendorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorCode, true
}

// SetVendorCode sets field value
func (o *TravelRuleWithdrawExchangesOrVASP) SetVendorCode(v string) {
	o.VendorCode = v
}

// GetVendorVaspId returns the VendorVaspId field value
func (o *TravelRuleWithdrawExchangesOrVASP) GetVendorVaspId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorVaspId
}

// GetVendorVaspIdOk returns a tuple with the VendorVaspId field value
// and a boolean to check if the value has been set.
func (o *TravelRuleWithdrawExchangesOrVASP) GetVendorVaspIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorVaspId, true
}

// SetVendorVaspId sets field value
func (o *TravelRuleWithdrawExchangesOrVASP) SetVendorVaspId(v string) {
	o.VendorVaspId = v
}

// GetEntityInfo returns the EntityInfo field value
func (o *TravelRuleWithdrawExchangesOrVASP) GetEntityInfo() TravelRuleWithdrawExchangesOrVASPEntityInfo {
	if o == nil {
		var ret TravelRuleWithdrawExchangesOrVASPEntityInfo
		return ret
	}

	return o.EntityInfo
}

// GetEntityInfoOk returns a tuple with the EntityInfo field value
// and a boolean to check if the value has been set.
func (o *TravelRuleWithdrawExchangesOrVASP) GetEntityInfoOk() (*TravelRuleWithdrawExchangesOrVASPEntityInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityInfo, true
}

// SetEntityInfo sets field value
func (o *TravelRuleWithdrawExchangesOrVASP) SetEntityInfo(v TravelRuleWithdrawExchangesOrVASPEntityInfo) {
	o.EntityInfo = v
}

func (o TravelRuleWithdrawExchangesOrVASP) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TravelRuleWithdrawExchangesOrVASP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination_wallet_type"] = o.DestinationWalletType
	toSerialize["vendor_code"] = o.VendorCode
	toSerialize["vendor_vasp_id"] = o.VendorVaspId
	toSerialize["entity_info"] = o.EntityInfo
	return toSerialize, nil
}

func (o *TravelRuleWithdrawExchangesOrVASP) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination_wallet_type",
		"vendor_code",
		"vendor_vasp_id",
		"entity_info",
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

	varTravelRuleWithdrawExchangesOrVASP := _TravelRuleWithdrawExchangesOrVASP{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTravelRuleWithdrawExchangesOrVASP)

	if err != nil {
		return err
	}

	*o = TravelRuleWithdrawExchangesOrVASP(varTravelRuleWithdrawExchangesOrVASP)

	return err
}

type NullableTravelRuleWithdrawExchangesOrVASP struct {
	value *TravelRuleWithdrawExchangesOrVASP
	isSet bool
}

func (v NullableTravelRuleWithdrawExchangesOrVASP) Get() *TravelRuleWithdrawExchangesOrVASP {
	return v.value
}

func (v *NullableTravelRuleWithdrawExchangesOrVASP) Set(val *TravelRuleWithdrawExchangesOrVASP) {
	v.value = val
	v.isSet = true
}

func (v NullableTravelRuleWithdrawExchangesOrVASP) IsSet() bool {
	return v.isSet
}

func (v *NullableTravelRuleWithdrawExchangesOrVASP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTravelRuleWithdrawExchangesOrVASP(val *TravelRuleWithdrawExchangesOrVASP) *NullableTravelRuleWithdrawExchangesOrVASP {
	return &NullableTravelRuleWithdrawExchangesOrVASP{value: val, isSet: true}
}

func (v NullableTravelRuleWithdrawExchangesOrVASP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTravelRuleWithdrawExchangesOrVASP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


