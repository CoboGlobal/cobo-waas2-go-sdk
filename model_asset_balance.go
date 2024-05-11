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
	"bytes"
	"fmt"
)

// checks if the AssetBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetBalance{}

// AssetBalance The data for token balance information.
type AssetBalance struct {
	// ID of the asset. Unique in all assets scope.
	AssetId string `json:"asset_id"`
	Balance TokenBalanceBalance `json:"balance"`
}

type _AssetBalance AssetBalance

// NewAssetBalance instantiates a new AssetBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetBalance(assetId string, balance TokenBalanceBalance) *AssetBalance {
	this := AssetBalance{}
	this.AssetId = assetId
	this.Balance = balance
	return &this
}

// NewAssetBalanceWithDefaults instantiates a new AssetBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetBalanceWithDefaults() *AssetBalance {
	this := AssetBalance{}
	return &this
}

// GetAssetId returns the AssetId field value
func (o *AssetBalance) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *AssetBalance) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *AssetBalance) SetAssetId(v string) {
	o.AssetId = v
}

// GetBalance returns the Balance field value
func (o *AssetBalance) GetBalance() TokenBalanceBalance {
	if o == nil {
		var ret TokenBalanceBalance
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *AssetBalance) GetBalanceOk() (*TokenBalanceBalance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *AssetBalance) SetBalance(v TokenBalanceBalance) {
	o.Balance = v
}

func (o AssetBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asset_id"] = o.AssetId
	toSerialize["balance"] = o.Balance
	return toSerialize, nil
}

func (o *AssetBalance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asset_id",
		"balance",
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

	varAssetBalance := _AssetBalance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssetBalance)

	if err != nil {
		return err
	}

	*o = AssetBalance(varAssetBalance)

	return err
}

type NullableAssetBalance struct {
	value *AssetBalance
	isSet bool
}

func (v NullableAssetBalance) Get() *AssetBalance {
	return v.value
}

func (v *NullableAssetBalance) Set(val *AssetBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetBalance(val *AssetBalance) *NullableAssetBalance {
	return &NullableAssetBalance{value: val, isSet: true}
}

func (v NullableAssetBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


