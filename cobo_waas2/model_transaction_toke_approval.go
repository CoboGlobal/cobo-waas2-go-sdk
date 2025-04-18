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

// checks if the TransactionTokeApproval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionTokeApproval{}

// TransactionTokeApproval struct for TransactionTokeApproval
type TransactionTokeApproval struct {
	// The token ID, which is the unique identifier of a token.
	TokenId string `json:"token_id"`
	// The ID of the chain on which the token operates.
	ChainId string `json:"chain_id"`
	// (This concept applies to Exchange Wallets only) The asset ID. An asset ID is the unique identifier of the asset held within your linked exchange account.
	AssetId *string `json:"asset_id,omitempty"`
	// The token symbol, which is the abbreviated name of a token.
	Symbol *string `json:"symbol,omitempty"`
	// The token name, which is the full name of a token.
	Name *string `json:"name,omitempty"`
	// The token decimal.
	Decimal *int32 `json:"decimal,omitempty"`
	// The URL of the token icon.
	IconUrl *string `json:"icon_url,omitempty"`
	// The token address, if applicable.
	TokenAddress *string `json:"token_address,omitempty"`
	// The fee token ID. A fee token is the token with which you pay transaction fees.
	FeeTokenId *string `json:"fee_token_id,omitempty"`
	// Whether deposits are enabled for this token.
	CanDeposit *bool `json:"can_deposit,omitempty"`
	// Whether withdrawals are enabled for this token.
	CanWithdraw *bool `json:"can_withdraw,omitempty"`
	// The minimum withdrawal amount for Custodial Wallets. If your withdrawal amount is smaller than this threshold, the withdrawal request will receive an error.  Note: [Cobo Loop](https://manuals.cobo.com/en/portal/custodial-wallets/cobo-loop) transfers do not have this limitation. 
	DustThreshold *string `json:"dust_threshold,omitempty"`
	// The minimum deposit amount for Custodial Wallets. If the amount you deposit to a Custodial Wallet is smaller than this threshold, the deposit will not show up on Cobo Portal or trigger any webhook events.  Note: [Cobo Loop](https://manuals.cobo.com/en/portal/custodial-wallets/cobo-loop)transfers do not have this limitation. 
	CustodialMinimumDepositThreshold *string `json:"custodial_minimum_deposit_threshold,omitempty"`
	AssetModelType *TokenAssetModelType `json:"asset_model_type,omitempty"`
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

// GetName returns the Name field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TransactionTokeApproval) SetName(v string) {
	o.Name = &v
}

// GetDecimal returns the Decimal field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetDecimal() int32 {
	if o == nil || IsNil(o.Decimal) {
		var ret int32
		return ret
	}
	return *o.Decimal
}

// GetDecimalOk returns a tuple with the Decimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetDecimalOk() (*int32, bool) {
	if o == nil || IsNil(o.Decimal) {
		return nil, false
	}
	return o.Decimal, true
}

// HasDecimal returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasDecimal() bool {
	if o != nil && !IsNil(o.Decimal) {
		return true
	}

	return false
}

// SetDecimal gets a reference to the given int32 and assigns it to the Decimal field.
func (o *TransactionTokeApproval) SetDecimal(v int32) {
	o.Decimal = &v
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

// GetFeeTokenId returns the FeeTokenId field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetFeeTokenId() string {
	if o == nil || IsNil(o.FeeTokenId) {
		var ret string
		return ret
	}
	return *o.FeeTokenId
}

// GetFeeTokenIdOk returns a tuple with the FeeTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetFeeTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeeTokenId) {
		return nil, false
	}
	return o.FeeTokenId, true
}

// HasFeeTokenId returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasFeeTokenId() bool {
	if o != nil && !IsNil(o.FeeTokenId) {
		return true
	}

	return false
}

// SetFeeTokenId gets a reference to the given string and assigns it to the FeeTokenId field.
func (o *TransactionTokeApproval) SetFeeTokenId(v string) {
	o.FeeTokenId = &v
}

// GetCanDeposit returns the CanDeposit field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetCanDeposit() bool {
	if o == nil || IsNil(o.CanDeposit) {
		var ret bool
		return ret
	}
	return *o.CanDeposit
}

// GetCanDepositOk returns a tuple with the CanDeposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetCanDepositOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeposit) {
		return nil, false
	}
	return o.CanDeposit, true
}

// HasCanDeposit returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasCanDeposit() bool {
	if o != nil && !IsNil(o.CanDeposit) {
		return true
	}

	return false
}

// SetCanDeposit gets a reference to the given bool and assigns it to the CanDeposit field.
func (o *TransactionTokeApproval) SetCanDeposit(v bool) {
	o.CanDeposit = &v
}

// GetCanWithdraw returns the CanWithdraw field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetCanWithdraw() bool {
	if o == nil || IsNil(o.CanWithdraw) {
		var ret bool
		return ret
	}
	return *o.CanWithdraw
}

// GetCanWithdrawOk returns a tuple with the CanWithdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetCanWithdrawOk() (*bool, bool) {
	if o == nil || IsNil(o.CanWithdraw) {
		return nil, false
	}
	return o.CanWithdraw, true
}

// HasCanWithdraw returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasCanWithdraw() bool {
	if o != nil && !IsNil(o.CanWithdraw) {
		return true
	}

	return false
}

// SetCanWithdraw gets a reference to the given bool and assigns it to the CanWithdraw field.
func (o *TransactionTokeApproval) SetCanWithdraw(v bool) {
	o.CanWithdraw = &v
}

// GetDustThreshold returns the DustThreshold field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetDustThreshold() string {
	if o == nil || IsNil(o.DustThreshold) {
		var ret string
		return ret
	}
	return *o.DustThreshold
}

// GetDustThresholdOk returns a tuple with the DustThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetDustThresholdOk() (*string, bool) {
	if o == nil || IsNil(o.DustThreshold) {
		return nil, false
	}
	return o.DustThreshold, true
}

// HasDustThreshold returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasDustThreshold() bool {
	if o != nil && !IsNil(o.DustThreshold) {
		return true
	}

	return false
}

// SetDustThreshold gets a reference to the given string and assigns it to the DustThreshold field.
func (o *TransactionTokeApproval) SetDustThreshold(v string) {
	o.DustThreshold = &v
}

// GetCustodialMinimumDepositThreshold returns the CustodialMinimumDepositThreshold field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetCustodialMinimumDepositThreshold() string {
	if o == nil || IsNil(o.CustodialMinimumDepositThreshold) {
		var ret string
		return ret
	}
	return *o.CustodialMinimumDepositThreshold
}

// GetCustodialMinimumDepositThresholdOk returns a tuple with the CustodialMinimumDepositThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetCustodialMinimumDepositThresholdOk() (*string, bool) {
	if o == nil || IsNil(o.CustodialMinimumDepositThreshold) {
		return nil, false
	}
	return o.CustodialMinimumDepositThreshold, true
}

// HasCustodialMinimumDepositThreshold returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasCustodialMinimumDepositThreshold() bool {
	if o != nil && !IsNil(o.CustodialMinimumDepositThreshold) {
		return true
	}

	return false
}

// SetCustodialMinimumDepositThreshold gets a reference to the given string and assigns it to the CustodialMinimumDepositThreshold field.
func (o *TransactionTokeApproval) SetCustodialMinimumDepositThreshold(v string) {
	o.CustodialMinimumDepositThreshold = &v
}

// GetAssetModelType returns the AssetModelType field value if set, zero value otherwise.
func (o *TransactionTokeApproval) GetAssetModelType() TokenAssetModelType {
	if o == nil || IsNil(o.AssetModelType) {
		var ret TokenAssetModelType
		return ret
	}
	return *o.AssetModelType
}

// GetAssetModelTypeOk returns a tuple with the AssetModelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokeApproval) GetAssetModelTypeOk() (*TokenAssetModelType, bool) {
	if o == nil || IsNil(o.AssetModelType) {
		return nil, false
	}
	return o.AssetModelType, true
}

// HasAssetModelType returns a boolean if a field has been set.
func (o *TransactionTokeApproval) HasAssetModelType() bool {
	if o != nil && !IsNil(o.AssetModelType) {
		return true
	}

	return false
}

// SetAssetModelType gets a reference to the given TokenAssetModelType and assigns it to the AssetModelType field.
func (o *TransactionTokeApproval) SetAssetModelType(v TokenAssetModelType) {
	o.AssetModelType = &v
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
	if !IsNil(o.AssetId) {
		toSerialize["asset_id"] = o.AssetId
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Decimal) {
		toSerialize["decimal"] = o.Decimal
	}
	if !IsNil(o.IconUrl) {
		toSerialize["icon_url"] = o.IconUrl
	}
	if !IsNil(o.TokenAddress) {
		toSerialize["token_address"] = o.TokenAddress
	}
	if !IsNil(o.FeeTokenId) {
		toSerialize["fee_token_id"] = o.FeeTokenId
	}
	if !IsNil(o.CanDeposit) {
		toSerialize["can_deposit"] = o.CanDeposit
	}
	if !IsNil(o.CanWithdraw) {
		toSerialize["can_withdraw"] = o.CanWithdraw
	}
	if !IsNil(o.DustThreshold) {
		toSerialize["dust_threshold"] = o.DustThreshold
	}
	if !IsNil(o.CustodialMinimumDepositThreshold) {
		toSerialize["custodial_minimum_deposit_threshold"] = o.CustodialMinimumDepositThreshold
	}
	if !IsNil(o.AssetModelType) {
		toSerialize["asset_model_type"] = o.AssetModelType
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


