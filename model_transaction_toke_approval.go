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

// checks if the TransactionTokeApproval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionTokeApproval{}

// TransactionTokeApproval struct for TransactionTokeApproval
type TransactionTokeApproval struct {
	// ID of the token. Unique in all chains scope.
	TokenId string `json:"token_id"`
	// The blockchain on which the token operates.
	ChainId string `json:"chain_id"`
	// Symbol for the token.
	Symbol *string `json:"symbol,omitempty"`
	// The description of the token.
	Description *string `json:"description,omitempty"`
	// URL of the icon image.
	IconUrl *string `json:"icon_url,omitempty"`
	// Address for token, if applicable.
	TokenAddress *string `json:"token_address,omitempty"`
	// ID of the asset. Used to group token balance when needed.
	AssetId *string `json:"asset_id,omitempty"`
	// Transaction value (Note that this is an absolute value. If you trade 1.5 BTC, then the value is 1.5) 
	Amount *float32 `json:"amount,omitempty"`
	// Spender address
	Spender *string `json:"spender,omitempty"`
}

type _TransactionTokeApproval TransactionTokeApproval

// NewTransactionTokeApproval instantiates a new TransactionTokeApproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionTokeApproval(tokenId string, chainId string) *TransactionTokeApproval {
	this := TransactionTokeApproval{}
	this.TokenId = tokenId
	this.ChainId = chainId
	return &this
}

// NewTransactionTokeApprovalWithDefaults instantiates a new TransactionTokeApproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionTokeApprovalWithDefaults() *TransactionTokeApproval {
	this := TransactionTokeApproval{}
	return &this
}

// GetTokenId returns the TokenId field value
func (o *TransactionTokeApproval) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *TransactionTokeApproval) SetTokenId(v string) {
	o.TokenId = v
}

// GetChainId returns the ChainId field value
func (o *TransactionTokeApproval) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *TransactionTokeApproval) SetChainId(v string) {
	o.ChainId = v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TransactionTokeApproval) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransactionTokeApproval) SetDescription(v string) {
	o.Description = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasIconUrl() bool {
	if o != nil && !IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *TransactionTokeApproval) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetTokenAddress returns the TokenAddress field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetTokenAddress() string {
	if o == nil || IsNil(o.TokenAddress) {
		var ret string
		return ret
	}
	return *o.TokenAddress
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetTokenAddressOk() (*string, bool) {
	if o == nil || IsNil(o.TokenAddress) {
		return nil, false
	}
	return o.TokenAddress, true
}

// HasTokenAddress returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasTokenAddress() bool {
	if o != nil && !IsNil(o.TokenAddress) {
		return true
	}

	return false
}

// SetTokenAddress gets a reference to the given string and assigns it to the TokenAddress field.
func (o *TransactionTokeApproval) SetTokenAddress(v string) {
	o.TokenAddress = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetAssetId() string {
	if o == nil || IsNil(o.AssetId) {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasAssetId() bool {
	if o != nil && !IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *TransactionTokeApproval) SetAssetId(v string) {
	o.AssetId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *TransactionTokeApproval) SetAmount(v float32) {
	o.Amount = &v
}

// GetSpender returns the Spender field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetSpender() string {
	if o == nil || IsNil(o.Spender) {
		var ret string
		return ret
	}
	return *o.Spender
}

// GetSpenderOk returns a tuple with the Spender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetSpenderOk() (*string, bool) {
	if o == nil || IsNil(o.Spender) {
		return nil, false
	}
	return o.Spender, true
}

// HasSpender returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasSpender() bool {
	if o != nil && !IsNil(o.Spender) {
		return true
	}

	return false
}

// SetSpender gets a reference to the given string and assigns it to the Spender field.
func (o *TransactionTokeApproval) SetSpender(v string) {
	o.Spender = &v
}

func (o TransactionTokeApproval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionTokeApproval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token_id"] = o.TokenId
	toSerialize["chain_id"] = o.ChainId
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IconUrl) {
		toSerialize["icon_url"] = o.IconUrl
	}
	if !IsNil(o.TokenAddress) {
		toSerialize["token_address"] = o.TokenAddress
	}
	if !IsNil(o.AssetId) {
		toSerialize["asset_id"] = o.AssetId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Spender) {
		toSerialize["spender"] = o.Spender
	}
	return toSerialize, nil
}

func (o *TransactionTokeApproval) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token_id",
		"chain_id",
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

	varTransactionTokeApproval := _TransactionTokeApproval{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionTokeApproval)

	if err != nil {
		return err
	}

	*o = TransactionTokeApproval(varTransactionTokeApproval)

	return err
}

type NullableTransactionTokeApproval struct {
	value *TransactionTokeApproval
	isSet bool
}

func (v NullableTransactionTokeApproval) Get() *TransactionTokeApproval {
	return v.value
}

func (v *NullableTransactionTokeApproval) Set(val *TransactionTokeApproval) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionTokeApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionTokeApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionTokeApproval(val *TransactionTokeApproval) *NullableTransactionTokeApproval {
	return &NullableTransactionTokeApproval{value: val, isSet: true}
}

func (v NullableTransactionTokeApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionTokeApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


