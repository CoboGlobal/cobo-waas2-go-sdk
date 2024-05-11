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

// checks if the CreateSafeWallet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSafeWallet{}

// CreateSafeWallet struct for CreateSafeWallet
type CreateSafeWallet struct {
	Name string `json:"name"`
	WalletType string `json:"wallet_type"`
	WalletSubtype string `json:"wallet_subtype"`
	// The label of the wallet.
	Label *string `json:"label,omitempty"`
	// The chain id the wallet is on.
	ChainId string `json:"chain_id"`
	SmartContractWalletType SmartContractWalletType `json:"smart_contract_wallet_type"`
	// The address of the smart contract wallet. If this is not provided, WaaS 2.0 will create a new safe wallet and setup cobo safe module for user. In this case, threshold, owners is required.
	SafeAddress *string `json:"safe_address,omitempty"`
	// The owners of the smart contract wallet. This MUST be provided when user want to create a new safe wallet.
	Owners []string `json:"owners,omitempty"`
	// The threshold of required confirmations for the smart contract wallet. This MUST be provided when user want to create a new safe wallet.
	Threshold *int32 `json:"threshold,omitempty"`
	// The address of the cobo safe module. Cobo safe module must has been created & enabled when import a existing safe wallet.
	CoboSafeAddress *string `json:"cobo_safe_address,omitempty"`
	Initiator *SafeWalletAllOfInitiator `json:"initiator,omitempty"`
}

type _CreateSafeWallet CreateSafeWallet

// NewCreateSafeWallet instantiates a new CreateSafeWallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSafeWallet(name string, walletType string, walletSubtype string, chainId string, smartContractWalletType SmartContractWalletType) *CreateSafeWallet {
	this := CreateSafeWallet{}
	this.Name = name
	this.WalletType = walletType
	this.WalletSubtype = walletSubtype
	this.ChainId = chainId
	this.SmartContractWalletType = smartContractWalletType
	return &this
}

// NewCreateSafeWalletWithDefaults instantiates a new CreateSafeWallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSafeWalletWithDefaults() *CreateSafeWallet {
	this := CreateSafeWallet{}
	var smartContractWalletType SmartContractWalletType = SMARTCONTRACTWALLETTYPE_SAFE_WALLET
	this.SmartContractWalletType = smartContractWalletType
	return &this
}

// GetName returns the Name field value
func (o *CreateSafeWallet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSafeWallet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSafeWallet) SetName(v string) {
	o.Name = v
}

// GetWalletType returns the WalletType field value
func (o *CreateSafeWallet) GetWalletType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value
// and a boolean to check if the value has been set.
func (o *CreateSafeWallet) GetWalletTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletType, true
}

// SetWalletType sets field value
func (o *CreateSafeWallet) SetWalletType(v string) {
	o.WalletType = v
}

// GetWalletSubtype returns the WalletSubtype field value
func (o *CreateSafeWallet) GetWalletSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletSubtype
}

// GetWalletSubtypeOk returns a tuple with the WalletSubtype field value
// and a boolean to check if the value has been set.
func (o *CreateSafeWallet) GetWalletSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletSubtype, true
}

// SetWalletSubtype sets field value
func (o *CreateSafeWallet) SetWalletSubtype(v string) {
	o.WalletSubtype = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CreateSafeWallet) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSafeWallet) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CreateSafeWallet) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CreateSafeWallet) SetLabel(v string) {
	o.Label = &v
}

// GetChainId returns the ChainId field value
func (o *CreateSafeWallet) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *CreateSafeWallet) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *CreateSafeWallet) SetChainId(v string) {
	o.ChainId = v
}

// GetSmartContractWalletType returns the SmartContractWalletType field value
func (o *CreateSafeWallet) GetSmartContractWalletType() SmartContractWalletType {
	if o == nil {
		var ret SmartContractWalletType
		return ret
	}

	return o.SmartContractWalletType
}

// GetSmartContractWalletTypeOk returns a tuple with the SmartContractWalletType field value
// and a boolean to check if the value has been set.
func (o *CreateSafeWallet) GetSmartContractWalletTypeOk() (*SmartContractWalletType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmartContractWalletType, true
}

// SetSmartContractWalletType sets field value
func (o *CreateSafeWallet) SetSmartContractWalletType(v SmartContractWalletType) {
	o.SmartContractWalletType = v
}

// GetSafeAddress returns the SafeAddress field value if set, zero value otherwise.
func (o *CreateSafeWallet) GetSafeAddress() string {
	if o == nil || IsNil(o.SafeAddress) {
		var ret string
		return ret
	}
	return *o.SafeAddress
}

// GetSafeAddressOk returns a tuple with the SafeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSafeWallet) GetSafeAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SafeAddress) {
		return nil, false
	}
	return o.SafeAddress, true
}

// HasSafeAddress returns a boolean if a field has been set.
func (o *CreateSafeWallet) HasSafeAddress() bool {
	if o != nil && !IsNil(o.SafeAddress) {
		return true
	}

	return false
}

// SetSafeAddress gets a reference to the given string and assigns it to the SafeAddress field.
func (o *CreateSafeWallet) SetSafeAddress(v string) {
	o.SafeAddress = &v
}

// GetOwners returns the Owners field value if set, zero value otherwise.
func (o *CreateSafeWallet) GetOwners() []string {
	if o == nil || IsNil(o.Owners) {
		var ret []string
		return ret
	}
	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSafeWallet) GetOwnersOk() ([]string, bool) {
	if o == nil || IsNil(o.Owners) {
		return nil, false
	}
	return o.Owners, true
}

// HasOwners returns a boolean if a field has been set.
func (o *CreateSafeWallet) HasOwners() bool {
	if o != nil && !IsNil(o.Owners) {
		return true
	}

	return false
}

// SetOwners gets a reference to the given []string and assigns it to the Owners field.
func (o *CreateSafeWallet) SetOwners(v []string) {
	o.Owners = v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *CreateSafeWallet) GetThreshold() int32 {
	if o == nil || IsNil(o.Threshold) {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSafeWallet) GetThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *CreateSafeWallet) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *CreateSafeWallet) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetCoboSafeAddress returns the CoboSafeAddress field value if set, zero value otherwise.
func (o *CreateSafeWallet) GetCoboSafeAddress() string {
	if o == nil || IsNil(o.CoboSafeAddress) {
		var ret string
		return ret
	}
	return *o.CoboSafeAddress
}

// GetCoboSafeAddressOk returns a tuple with the CoboSafeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSafeWallet) GetCoboSafeAddressOk() (*string, bool) {
	if o == nil || IsNil(o.CoboSafeAddress) {
		return nil, false
	}
	return o.CoboSafeAddress, true
}

// HasCoboSafeAddress returns a boolean if a field has been set.
func (o *CreateSafeWallet) HasCoboSafeAddress() bool {
	if o != nil && !IsNil(o.CoboSafeAddress) {
		return true
	}

	return false
}

// SetCoboSafeAddress gets a reference to the given string and assigns it to the CoboSafeAddress field.
func (o *CreateSafeWallet) SetCoboSafeAddress(v string) {
	o.CoboSafeAddress = &v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *CreateSafeWallet) GetInitiator() SafeWalletAllOfInitiator {
	if o == nil || IsNil(o.Initiator) {
		var ret SafeWalletAllOfInitiator
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSafeWallet) GetInitiatorOk() (*SafeWalletAllOfInitiator, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *CreateSafeWallet) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given SafeWalletAllOfInitiator and assigns it to the Initiator field.
func (o *CreateSafeWallet) SetInitiator(v SafeWalletAllOfInitiator) {
	o.Initiator = &v
}

func (o CreateSafeWallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSafeWallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["wallet_type"] = o.WalletType
	toSerialize["wallet_subtype"] = o.WalletSubtype
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["chain_id"] = o.ChainId
	toSerialize["smart_contract_wallet_type"] = o.SmartContractWalletType
	if !IsNil(o.SafeAddress) {
		toSerialize["safe_address"] = o.SafeAddress
	}
	if !IsNil(o.Owners) {
		toSerialize["owners"] = o.Owners
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.CoboSafeAddress) {
		toSerialize["cobo_safe_address"] = o.CoboSafeAddress
	}
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	return toSerialize, nil
}

func (o *CreateSafeWallet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"wallet_type",
		"wallet_subtype",
		"chain_id",
		"smart_contract_wallet_type",
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

	varCreateSafeWallet := _CreateSafeWallet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateSafeWallet)

	if err != nil {
		return err
	}

	*o = CreateSafeWallet(varCreateSafeWallet)

	return err
}

type NullableCreateSafeWallet struct {
	value *CreateSafeWallet
	isSet bool
}

func (v NullableCreateSafeWallet) Get() *CreateSafeWallet {
	return v.value
}

func (v *NullableCreateSafeWallet) Set(val *CreateSafeWallet) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSafeWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSafeWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSafeWallet(val *CreateSafeWallet) *NullableCreateSafeWallet {
	return &NullableCreateSafeWallet{value: val, isSet: true}
}

func (v NullableCreateSafeWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSafeWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


