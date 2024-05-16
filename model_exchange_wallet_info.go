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

// checks if the ExchangeWalletInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeWalletInfo{}

// ExchangeWalletInfo struct for ExchangeWalletInfo
type ExchangeWalletInfo struct {
	WalletId string `json:"wallet_id"`
	WalletType WalletType `json:"wallet_type"`
	WalletSubtype WalletSubtype `json:"wallet_subtype"`
	Name string `json:"name"`
	// The owning custody organization id of the custodial wallet.
	OrgId string `json:"org_id"`
	// The API Key for the exchange.
	Apikey string `json:"apikey"`
	ExchangeId ExchangeId `json:"exchange_id"`
	// The parent wallet id for this wallet if presented.
	ParentWalletId *string `json:"parent_wallet_id,omitempty"`
	SubAccounts []ExchangeWalletInfoAllOfSubAccounts `json:"sub_accounts,omitempty"`
}

type _ExchangeWalletInfo ExchangeWalletInfo

// NewExchangeWalletInfo instantiates a new ExchangeWalletInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeWalletInfo(walletId string, walletType WalletType, walletSubtype WalletSubtype, name string, orgId string, apikey string, exchangeId ExchangeId) *ExchangeWalletInfo {
	this := ExchangeWalletInfo{}
	this.WalletId = walletId
	this.WalletType = walletType
	this.WalletSubtype = walletSubtype
	this.Name = name
	this.OrgId = orgId
	this.Apikey = apikey
	this.ExchangeId = exchangeId
	return &this
}

// NewExchangeWalletInfoWithDefaults instantiates a new ExchangeWalletInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeWalletInfoWithDefaults() *ExchangeWalletInfo {
	this := ExchangeWalletInfo{}
	return &this
}

// GetWalletId returns the WalletId field value
func (o *ExchangeWalletInfo) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *ExchangeWalletInfo) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *ExchangeWalletInfo) SetWalletId(v string) {
	o.WalletId = v
}

// GetWalletType returns the WalletType field value
func (o *ExchangeWalletInfo) GetWalletType() WalletType {
	if o == nil {
		var ret WalletType
		return ret
	}

	return o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value
// and a boolean to check if the value has been set.
func (o *ExchangeWalletInfo) GetWalletTypeOk() (*WalletType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletType, true
}

// SetWalletType sets field value
func (o *ExchangeWalletInfo) SetWalletType(v WalletType) {
	o.WalletType = v
}

// GetWalletSubtype returns the WalletSubtype field value
func (o *ExchangeWalletInfo) GetWalletSubtype() WalletSubtype {
	if o == nil {
		var ret WalletSubtype
		return ret
	}

	return o.WalletSubtype
}

// GetWalletSubtypeOk returns a tuple with the WalletSubtype field value
// and a boolean to check if the value has been set.
func (o *ExchangeWalletInfo) GetWalletSubtypeOk() (*WalletSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletSubtype, true
}

// SetWalletSubtype sets field value
func (o *ExchangeWalletInfo) SetWalletSubtype(v WalletSubtype) {
	o.WalletSubtype = v
}

// GetName returns the Name field value
func (o *ExchangeWalletInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExchangeWalletInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExchangeWalletInfo) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value
func (o *ExchangeWalletInfo) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ExchangeWalletInfo) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *ExchangeWalletInfo) SetOrgId(v string) {
	o.OrgId = v
}

// GetApikey returns the Apikey field value
func (o *ExchangeWalletInfo) GetApikey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Apikey
}

// GetApikeyOk returns a tuple with the Apikey field value
// and a boolean to check if the value has been set.
func (o *ExchangeWalletInfo) GetApikeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Apikey, true
}

// SetApikey sets field value
func (o *ExchangeWalletInfo) SetApikey(v string) {
	o.Apikey = v
}

// GetExchangeId returns the ExchangeId field value
func (o *ExchangeWalletInfo) GetExchangeId() ExchangeId {
	if o == nil {
		var ret ExchangeId
		return ret
	}

	return o.ExchangeId
}

// GetExchangeIdOk returns a tuple with the ExchangeId field value
// and a boolean to check if the value has been set.
func (o *ExchangeWalletInfo) GetExchangeIdOk() (*ExchangeId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeId, true
}

// SetExchangeId sets field value
func (o *ExchangeWalletInfo) SetExchangeId(v ExchangeId) {
	o.ExchangeId = v
}

// GetParentWalletId returns the ParentWalletId field value if set, zero value otherwise.
func (o *ExchangeWalletInfo) GetParentWalletId() string {
	if o == nil || IsNil(o.ParentWalletId) {
		var ret string
		return ret
	}
	return *o.ParentWalletId
}

// GetParentWalletIdOk returns a tuple with the ParentWalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeWalletInfo) GetParentWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentWalletId) {
		return nil, false
	}
	return o.ParentWalletId, true
}

// HasParentWalletId returns a boolean if a field has been set.
func (o *ExchangeWalletInfo) HasParentWalletId() bool {
	if o != nil && !IsNil(o.ParentWalletId) {
		return true
	}

	return false
}

// SetParentWalletId gets a reference to the given string and assigns it to the ParentWalletId field.
func (o *ExchangeWalletInfo) SetParentWalletId(v string) {
	o.ParentWalletId = &v
}

// GetSubAccounts returns the SubAccounts field value if set, zero value otherwise.
func (o *ExchangeWalletInfo) GetSubAccounts() []ExchangeWalletInfoAllOfSubAccounts {
	if o == nil || IsNil(o.SubAccounts) {
		var ret []ExchangeWalletInfoAllOfSubAccounts
		return ret
	}
	return o.SubAccounts
}

// GetSubAccountsOk returns a tuple with the SubAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeWalletInfo) GetSubAccountsOk() ([]ExchangeWalletInfoAllOfSubAccounts, bool) {
	if o == nil || IsNil(o.SubAccounts) {
		return nil, false
	}
	return o.SubAccounts, true
}

// HasSubAccounts returns a boolean if a field has been set.
func (o *ExchangeWalletInfo) HasSubAccounts() bool {
	if o != nil && !IsNil(o.SubAccounts) {
		return true
	}

	return false
}

// SetSubAccounts gets a reference to the given []ExchangeWalletInfoAllOfSubAccounts and assigns it to the SubAccounts field.
func (o *ExchangeWalletInfo) SetSubAccounts(v []ExchangeWalletInfoAllOfSubAccounts) {
	o.SubAccounts = v
}

func (o ExchangeWalletInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeWalletInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["wallet_type"] = o.WalletType
	toSerialize["wallet_subtype"] = o.WalletSubtype
	toSerialize["name"] = o.Name
	toSerialize["org_id"] = o.OrgId
	toSerialize["apikey"] = o.Apikey
	toSerialize["exchange_id"] = o.ExchangeId
	if !IsNil(o.ParentWalletId) {
		toSerialize["parent_wallet_id"] = o.ParentWalletId
	}
	if !IsNil(o.SubAccounts) {
		toSerialize["sub_accounts"] = o.SubAccounts
	}
	return toSerialize, nil
}

func (o *ExchangeWalletInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wallet_id",
		"wallet_type",
		"wallet_subtype",
		"name",
		"org_id",
		"apikey",
		"exchange_id",
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

	varExchangeWalletInfo := _ExchangeWalletInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExchangeWalletInfo)

	if err != nil {
		return err
	}

	*o = ExchangeWalletInfo(varExchangeWalletInfo)

	return err
}

type NullableExchangeWalletInfo struct {
	value *ExchangeWalletInfo
	isSet bool
}

func (v NullableExchangeWalletInfo) Get() *ExchangeWalletInfo {
	return v.value
}

func (v *NullableExchangeWalletInfo) Set(val *ExchangeWalletInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeWalletInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeWalletInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeWalletInfo(val *ExchangeWalletInfo) *NullableExchangeWalletInfo {
	return &NullableExchangeWalletInfo{value: val, isSet: true}
}

func (v NullableExchangeWalletInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeWalletInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


