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

// checks if the CreateExchangeWallet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateExchangeWallet{}

// CreateExchangeWallet struct for CreateExchangeWallet
type CreateExchangeWallet struct {
	Name string `json:"name"`
	WalletType string `json:"wallet_type"`
	WalletSubtype string `json:"wallet_subtype"`
	ExchangeId ExchangeId `json:"exchange_id"`
	// The API Key for the exchange.
	Apikey string `json:"apikey"`
	// The API Secret for the exchange.
	Secret string `json:"secret"`
	// The passphrase for the exchange.
	Passphrase *string `json:"passphrase,omitempty"`
	// The ga_code for the exchange.
	GaCode *string `json:"ga_code,omitempty"`
	// The unique identifier associated with the exchange sub-account. It can be an email address, username, or a custom account ID.
	SubAccountIds []string `json:"sub_account_ids,omitempty"`
}

type _CreateExchangeWallet CreateExchangeWallet

// NewCreateExchangeWallet instantiates a new CreateExchangeWallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateExchangeWallet(name string, walletType string, walletSubtype string, exchangeId ExchangeId, apikey string, secret string) *CreateExchangeWallet {
	this := CreateExchangeWallet{}
	this.Name = name
	this.WalletType = walletType
	this.WalletSubtype = walletSubtype
	this.ExchangeId = exchangeId
	this.Apikey = apikey
	this.Secret = secret
	return &this
}

// NewCreateExchangeWalletWithDefaults instantiates a new CreateExchangeWallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateExchangeWalletWithDefaults() *CreateExchangeWallet {
	this := CreateExchangeWallet{}
	return &this
}

// GetName returns the Name field value
func (o *CreateExchangeWallet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateExchangeWallet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateExchangeWallet) SetName(v string) {
	o.Name = v
}

// GetWalletType returns the WalletType field value
func (o *CreateExchangeWallet) GetWalletType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value
// and a boolean to check if the value has been set.
func (o *CreateExchangeWallet) GetWalletTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletType, true
}

// SetWalletType sets field value
func (o *CreateExchangeWallet) SetWalletType(v string) {
	o.WalletType = v
}

// GetWalletSubtype returns the WalletSubtype field value
func (o *CreateExchangeWallet) GetWalletSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletSubtype
}

// GetWalletSubtypeOk returns a tuple with the WalletSubtype field value
// and a boolean to check if the value has been set.
func (o *CreateExchangeWallet) GetWalletSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletSubtype, true
}

// SetWalletSubtype sets field value
func (o *CreateExchangeWallet) SetWalletSubtype(v string) {
	o.WalletSubtype = v
}

// GetExchangeId returns the ExchangeId field value
func (o *CreateExchangeWallet) GetExchangeId() ExchangeId {
	if o == nil {
		var ret ExchangeId
		return ret
	}

	return o.ExchangeId
}

// GetExchangeIdOk returns a tuple with the ExchangeId field value
// and a boolean to check if the value has been set.
func (o *CreateExchangeWallet) GetExchangeIdOk() (*ExchangeId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeId, true
}

// SetExchangeId sets field value
func (o *CreateExchangeWallet) SetExchangeId(v ExchangeId) {
	o.ExchangeId = v
}

// GetApikey returns the Apikey field value
func (o *CreateExchangeWallet) GetApikey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Apikey
}

// GetApikeyOk returns a tuple with the Apikey field value
// and a boolean to check if the value has been set.
func (o *CreateExchangeWallet) GetApikeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Apikey, true
}

// SetApikey sets field value
func (o *CreateExchangeWallet) SetApikey(v string) {
	o.Apikey = v
}

// GetSecret returns the Secret field value
func (o *CreateExchangeWallet) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *CreateExchangeWallet) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *CreateExchangeWallet) SetSecret(v string) {
	o.Secret = v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *CreateExchangeWallet) GetPassphrase() string {
	if o == nil || IsNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExchangeWallet) GetPassphraseOk() (*string, bool) {
	if o == nil || IsNil(o.Passphrase) {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *CreateExchangeWallet) HasPassphrase() bool {
	if o != nil && !IsNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *CreateExchangeWallet) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetGaCode returns the GaCode field value if set, zero value otherwise.
func (o *CreateExchangeWallet) GetGaCode() string {
	if o == nil || IsNil(o.GaCode) {
		var ret string
		return ret
	}
	return *o.GaCode
}

// GetGaCodeOk returns a tuple with the GaCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExchangeWallet) GetGaCodeOk() (*string, bool) {
	if o == nil || IsNil(o.GaCode) {
		return nil, false
	}
	return o.GaCode, true
}

// HasGaCode returns a boolean if a field has been set.
func (o *CreateExchangeWallet) HasGaCode() bool {
	if o != nil && !IsNil(o.GaCode) {
		return true
	}

	return false
}

// SetGaCode gets a reference to the given string and assigns it to the GaCode field.
func (o *CreateExchangeWallet) SetGaCode(v string) {
	o.GaCode = &v
}

// GetSubAccountIds returns the SubAccountIds field value if set, zero value otherwise.
func (o *CreateExchangeWallet) GetSubAccountIds() []string {
	if o == nil || IsNil(o.SubAccountIds) {
		var ret []string
		return ret
	}
	return o.SubAccountIds
}

// GetSubAccountIdsOk returns a tuple with the SubAccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExchangeWallet) GetSubAccountIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SubAccountIds) {
		return nil, false
	}
	return o.SubAccountIds, true
}

// HasSubAccountIds returns a boolean if a field has been set.
func (o *CreateExchangeWallet) HasSubAccountIds() bool {
	if o != nil && !IsNil(o.SubAccountIds) {
		return true
	}

	return false
}

// SetSubAccountIds gets a reference to the given []string and assigns it to the SubAccountIds field.
func (o *CreateExchangeWallet) SetSubAccountIds(v []string) {
	o.SubAccountIds = v
}

func (o CreateExchangeWallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateExchangeWallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["wallet_type"] = o.WalletType
	toSerialize["wallet_subtype"] = o.WalletSubtype
	toSerialize["exchange_id"] = o.ExchangeId
	toSerialize["apikey"] = o.Apikey
	toSerialize["secret"] = o.Secret
	if !IsNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	if !IsNil(o.GaCode) {
		toSerialize["ga_code"] = o.GaCode
	}
	if !IsNil(o.SubAccountIds) {
		toSerialize["sub_account_ids"] = o.SubAccountIds
	}
	return toSerialize, nil
}

func (o *CreateExchangeWallet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"wallet_type",
		"wallet_subtype",
		"exchange_id",
		"apikey",
		"secret",
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

	varCreateExchangeWallet := _CreateExchangeWallet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateExchangeWallet)

	if err != nil {
		return err
	}

	*o = CreateExchangeWallet(varCreateExchangeWallet)

	return err
}

type NullableCreateExchangeWallet struct {
	value *CreateExchangeWallet
	isSet bool
}

func (v NullableCreateExchangeWallet) Get() *CreateExchangeWallet {
	return v.value
}

func (v *NullableCreateExchangeWallet) Set(val *CreateExchangeWallet) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateExchangeWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateExchangeWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateExchangeWallet(val *CreateExchangeWallet) *NullableCreateExchangeWallet {
	return &NullableCreateExchangeWallet{value: val, isSet: true}
}

func (v NullableCreateExchangeWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateExchangeWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


