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

// checks if the TransactionDepositToAddressDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionDepositToAddressDestination{}

// TransactionDepositToAddressDestination Information about the transaction destination type `DepositToAddress`. Switch between the tabs to display the properties for different transaction destinations. 
type TransactionDepositToAddressDestination struct {
	DestinationType TransactionDestinationType `json:"destination_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	WalletType WalletType `json:"wallet_type"`
	WalletSubtype *WalletSubtype `json:"wallet_subtype,omitempty"`
	// The destination address.
	Address string `json:"address"`
	// The memo that identifies a transaction in order to credit the correct account. For transfers out of Cobo Portal, it is highly recommended to include a memo for the chains such as XRP, EOS, XLM, IOST, BNB_BNB, ATOM, LUNA, and TON.
	Memo *string `json:"memo,omitempty"`
	// The transfer amount. For example, if you trade 1.5 BTC, then the value is `1.5`. 
	Amount string `json:"amount"`
}

type _TransactionDepositToAddressDestination TransactionDepositToAddressDestination

// NewTransactionDepositToAddressDestination instantiates a new TransactionDepositToAddressDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionDepositToAddressDestination(destinationType TransactionDestinationType, walletId string, walletType WalletType, address string, amount string) *TransactionDepositToAddressDestination {
	this := TransactionDepositToAddressDestination{}
	this.DestinationType = destinationType
	this.WalletId = walletId
	this.WalletType = walletType
	this.Address = address
	this.Amount = amount
	return &this
}

// NewTransactionDepositToAddressDestinationWithDefaults instantiates a new TransactionDepositToAddressDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDepositToAddressDestinationWithDefaults() *TransactionDepositToAddressDestination {
	this := TransactionDepositToAddressDestination{}
	return &this
}

// GetDestinationType returns the DestinationType field value
func (o *TransactionDepositToAddressDestination) GetDestinationType() TransactionDestinationType {
	if o == nil {
		var ret TransactionDestinationType
		return ret
	}

	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositToAddressDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value
func (o *TransactionDepositToAddressDestination) SetDestinationType(v TransactionDestinationType) {
	o.DestinationType = v
}

// GetWalletId returns the WalletId field value
func (o *TransactionDepositToAddressDestination) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositToAddressDestination) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *TransactionDepositToAddressDestination) SetWalletId(v string) {
	o.WalletId = v
}

// GetWalletType returns the WalletType field value
func (o *TransactionDepositToAddressDestination) GetWalletType() WalletType {
	if o == nil {
		var ret WalletType
		return ret
	}

	return o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositToAddressDestination) GetWalletTypeOk() (*WalletType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletType, true
}

// SetWalletType sets field value
func (o *TransactionDepositToAddressDestination) SetWalletType(v WalletType) {
	o.WalletType = v
}

// GetWalletSubtype returns the WalletSubtype field value if set, zero value otherwise.
func (o *TransactionDepositToAddressDestination) GetWalletSubtype() WalletSubtype {
	if o == nil || IsNil(o.WalletSubtype) {
		var ret WalletSubtype
		return ret
	}
	return *o.WalletSubtype
}

// GetWalletSubtypeOk returns a tuple with the WalletSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDepositToAddressDestination) GetWalletSubtypeOk() (*WalletSubtype, bool) {
	if o == nil || IsNil(o.WalletSubtype) {
		return nil, false
	}
	return o.WalletSubtype, true
}

// HasWalletSubtype returns a boolean if a field has been set.
func (o *TransactionDepositToAddressDestination) HasWalletSubtype() bool {
	if o != nil && !IsNil(o.WalletSubtype) {
		return true
	}

	return false
}

// SetWalletSubtype gets a reference to the given WalletSubtype and assigns it to the WalletSubtype field.
func (o *TransactionDepositToAddressDestination) SetWalletSubtype(v WalletSubtype) {
	o.WalletSubtype = &v
}

// GetAddress returns the Address field value
func (o *TransactionDepositToAddressDestination) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositToAddressDestination) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *TransactionDepositToAddressDestination) SetAddress(v string) {
	o.Address = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *TransactionDepositToAddressDestination) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDepositToAddressDestination) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *TransactionDepositToAddressDestination) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *TransactionDepositToAddressDestination) SetMemo(v string) {
	o.Memo = &v
}

// GetAmount returns the Amount field value
func (o *TransactionDepositToAddressDestination) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositToAddressDestination) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionDepositToAddressDestination) SetAmount(v string) {
	o.Amount = v
}

func (o TransactionDepositToAddressDestination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionDepositToAddressDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination_type"] = o.DestinationType
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["wallet_type"] = o.WalletType
	if !IsNil(o.WalletSubtype) {
		toSerialize["wallet_subtype"] = o.WalletSubtype
	}
	toSerialize["address"] = o.Address
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *TransactionDepositToAddressDestination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination_type",
		"wallet_id",
		"wallet_type",
		"address",
		"amount",
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

	varTransactionDepositToAddressDestination := _TransactionDepositToAddressDestination{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionDepositToAddressDestination)

	if err != nil {
		return err
	}

	*o = TransactionDepositToAddressDestination(varTransactionDepositToAddressDestination)

	return err
}

type NullableTransactionDepositToAddressDestination struct {
	value *TransactionDepositToAddressDestination
	isSet bool
}

func (v NullableTransactionDepositToAddressDestination) Get() *TransactionDepositToAddressDestination {
	return v.value
}

func (v *NullableTransactionDepositToAddressDestination) Set(val *TransactionDepositToAddressDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDepositToAddressDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDepositToAddressDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDepositToAddressDestination(val *TransactionDepositToAddressDestination) *NullableTransactionDepositToAddressDestination {
	return &NullableTransactionDepositToAddressDestination{value: val, isSet: true}
}

func (v NullableTransactionDepositToAddressDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDepositToAddressDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


