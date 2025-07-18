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

// checks if the TokenizationTokenInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenizationTokenInfo{}

// TokenizationTokenInfo struct for TokenizationTokenInfo
type TokenizationTokenInfo struct {
	// The unique token identifier.
	TokenId string `json:"token_id"`
	// The chain ID of the tokenization contract.
	ChainId string `json:"chain_id"`
	// The address of the token contract.
	TokenAddress *string `json:"token_address,omitempty"`
	// The name of the token.
	TokenName *string `json:"token_name,omitempty"`
	// The unique token symbol.
	TokenSymbol string `json:"token_symbol"`
	// The number of decimals of the token.
	Decimals int32 `json:"decimals"`
	// Whether the allowlist feature is activated for the token.
	AllowlistActivated *bool `json:"allowlist_activated,omitempty"`
	Status TokenizationStatus `json:"status"`
	// The total supply of the token.
	TotalSupply *string `json:"total_supply,omitempty"`
	// The amount of tokens held by the organization.
	Holdings *string `json:"holdings,omitempty"`
}

type _TokenizationTokenInfo TokenizationTokenInfo

// NewTokenizationTokenInfo instantiates a new TokenizationTokenInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizationTokenInfo(tokenId string, chainId string, tokenSymbol string, decimals int32, status TokenizationStatus) *TokenizationTokenInfo {
	this := TokenizationTokenInfo{}
	this.TokenId = tokenId
	this.ChainId = chainId
	this.TokenSymbol = tokenSymbol
	this.Decimals = decimals
	this.Status = status
	return &this
}

// NewTokenizationTokenInfoWithDefaults instantiates a new TokenizationTokenInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizationTokenInfoWithDefaults() *TokenizationTokenInfo {
	this := TokenizationTokenInfo{}
	return &this
}

// GetTokenId returns the TokenId field value
func (o *TokenizationTokenInfo) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *TokenizationTokenInfo) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *TokenizationTokenInfo) SetTokenId(v string) {
	o.TokenId = v
}

// GetChainId returns the ChainId field value
func (o *TokenizationTokenInfo) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *TokenizationTokenInfo) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *TokenizationTokenInfo) SetChainId(v string) {
	o.ChainId = v
}

// GetTokenAddress returns the TokenAddress field value if set, zero value otherwise.
func (o *TokenizationTokenInfo) GetTokenAddress() string {
	if o == nil || IsNil(o.TokenAddress) {
		var ret string
		return ret
	}
	return *o.TokenAddress
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizationTokenInfo) GetTokenAddressOk() (*string, bool) {
	if o == nil || IsNil(o.TokenAddress) {
		return nil, false
	}
	return o.TokenAddress, true
}

// HasTokenAddress returns a boolean if a field has been set.
func (o *TokenizationTokenInfo) HasTokenAddress() bool {
	if o != nil && !IsNil(o.TokenAddress) {
		return true
	}

	return false
}

// SetTokenAddress gets a reference to the given string and assigns it to the TokenAddress field.
func (o *TokenizationTokenInfo) SetTokenAddress(v string) {
	o.TokenAddress = &v
}

// GetTokenName returns the TokenName field value if set, zero value otherwise.
func (o *TokenizationTokenInfo) GetTokenName() string {
	if o == nil || IsNil(o.TokenName) {
		var ret string
		return ret
	}
	return *o.TokenName
}

// GetTokenNameOk returns a tuple with the TokenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizationTokenInfo) GetTokenNameOk() (*string, bool) {
	if o == nil || IsNil(o.TokenName) {
		return nil, false
	}
	return o.TokenName, true
}

// HasTokenName returns a boolean if a field has been set.
func (o *TokenizationTokenInfo) HasTokenName() bool {
	if o != nil && !IsNil(o.TokenName) {
		return true
	}

	return false
}

// SetTokenName gets a reference to the given string and assigns it to the TokenName field.
func (o *TokenizationTokenInfo) SetTokenName(v string) {
	o.TokenName = &v
}

// GetTokenSymbol returns the TokenSymbol field value
func (o *TokenizationTokenInfo) GetTokenSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenSymbol
}

// GetTokenSymbolOk returns a tuple with the TokenSymbol field value
// and a boolean to check if the value has been set.
func (o *TokenizationTokenInfo) GetTokenSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenSymbol, true
}

// SetTokenSymbol sets field value
func (o *TokenizationTokenInfo) SetTokenSymbol(v string) {
	o.TokenSymbol = v
}

// GetDecimals returns the Decimals field value
func (o *TokenizationTokenInfo) GetDecimals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value
// and a boolean to check if the value has been set.
func (o *TokenizationTokenInfo) GetDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decimals, true
}

// SetDecimals sets field value
func (o *TokenizationTokenInfo) SetDecimals(v int32) {
	o.Decimals = v
}

// GetAllowlistActivated returns the AllowlistActivated field value if set, zero value otherwise.
func (o *TokenizationTokenInfo) GetAllowlistActivated() bool {
	if o == nil || IsNil(o.AllowlistActivated) {
		var ret bool
		return ret
	}
	return *o.AllowlistActivated
}

// GetAllowlistActivatedOk returns a tuple with the AllowlistActivated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizationTokenInfo) GetAllowlistActivatedOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowlistActivated) {
		return nil, false
	}
	return o.AllowlistActivated, true
}

// HasAllowlistActivated returns a boolean if a field has been set.
func (o *TokenizationTokenInfo) HasAllowlistActivated() bool {
	if o != nil && !IsNil(o.AllowlistActivated) {
		return true
	}

	return false
}

// SetAllowlistActivated gets a reference to the given bool and assigns it to the AllowlistActivated field.
func (o *TokenizationTokenInfo) SetAllowlistActivated(v bool) {
	o.AllowlistActivated = &v
}

// GetStatus returns the Status field value
func (o *TokenizationTokenInfo) GetStatus() TokenizationStatus {
	if o == nil {
		var ret TokenizationStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TokenizationTokenInfo) GetStatusOk() (*TokenizationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TokenizationTokenInfo) SetStatus(v TokenizationStatus) {
	o.Status = v
}

// GetTotalSupply returns the TotalSupply field value if set, zero value otherwise.
func (o *TokenizationTokenInfo) GetTotalSupply() string {
	if o == nil || IsNil(o.TotalSupply) {
		var ret string
		return ret
	}
	return *o.TotalSupply
}

// GetTotalSupplyOk returns a tuple with the TotalSupply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizationTokenInfo) GetTotalSupplyOk() (*string, bool) {
	if o == nil || IsNil(o.TotalSupply) {
		return nil, false
	}
	return o.TotalSupply, true
}

// HasTotalSupply returns a boolean if a field has been set.
func (o *TokenizationTokenInfo) HasTotalSupply() bool {
	if o != nil && !IsNil(o.TotalSupply) {
		return true
	}

	return false
}

// SetTotalSupply gets a reference to the given string and assigns it to the TotalSupply field.
func (o *TokenizationTokenInfo) SetTotalSupply(v string) {
	o.TotalSupply = &v
}

// GetHoldings returns the Holdings field value if set, zero value otherwise.
func (o *TokenizationTokenInfo) GetHoldings() string {
	if o == nil || IsNil(o.Holdings) {
		var ret string
		return ret
	}
	return *o.Holdings
}

// GetHoldingsOk returns a tuple with the Holdings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizationTokenInfo) GetHoldingsOk() (*string, bool) {
	if o == nil || IsNil(o.Holdings) {
		return nil, false
	}
	return o.Holdings, true
}

// HasHoldings returns a boolean if a field has been set.
func (o *TokenizationTokenInfo) HasHoldings() bool {
	if o != nil && !IsNil(o.Holdings) {
		return true
	}

	return false
}

// SetHoldings gets a reference to the given string and assigns it to the Holdings field.
func (o *TokenizationTokenInfo) SetHoldings(v string) {
	o.Holdings = &v
}

func (o TokenizationTokenInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizationTokenInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token_id"] = o.TokenId
	toSerialize["chain_id"] = o.ChainId
	if !IsNil(o.TokenAddress) {
		toSerialize["token_address"] = o.TokenAddress
	}
	if !IsNil(o.TokenName) {
		toSerialize["token_name"] = o.TokenName
	}
	toSerialize["token_symbol"] = o.TokenSymbol
	toSerialize["decimals"] = o.Decimals
	if !IsNil(o.AllowlistActivated) {
		toSerialize["allowlist_activated"] = o.AllowlistActivated
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.TotalSupply) {
		toSerialize["total_supply"] = o.TotalSupply
	}
	if !IsNil(o.Holdings) {
		toSerialize["holdings"] = o.Holdings
	}
	return toSerialize, nil
}

func (o *TokenizationTokenInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token_id",
		"chain_id",
		"token_symbol",
		"decimals",
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

	varTokenizationTokenInfo := _TokenizationTokenInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenizationTokenInfo)

	if err != nil {
		return err
	}

	*o = TokenizationTokenInfo(varTokenizationTokenInfo)

	return err
}

type NullableTokenizationTokenInfo struct {
	value *TokenizationTokenInfo
	isSet bool
}

func (v NullableTokenizationTokenInfo) Get() *TokenizationTokenInfo {
	return v.value
}

func (v *NullableTokenizationTokenInfo) Set(val *TokenizationTokenInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizationTokenInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizationTokenInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizationTokenInfo(val *TokenizationTokenInfo) *NullableTokenizationTokenInfo {
	return &NullableTokenizationTokenInfo{value: val, isSet: true}
}

func (v NullableTokenizationTokenInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizationTokenInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


