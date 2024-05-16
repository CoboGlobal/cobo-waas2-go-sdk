/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
)

// checks if the Transaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transaction{}

// Transaction The data for transaction information.
type Transaction struct {
	// Unique transaction ID
	TransactionId *string `json:"transaction_id,omitempty"`
	// Wallet ID
	WalletId *string `json:"wallet_id,omitempty"`
	// Request ID
	RequestId *string `json:"request_id,omitempty"`
	// Cobo ID
	CoboId *string `json:"cobo_id,omitempty"`
	Status *TransactionStatus `json:"status,omitempty"`
	SubStatus *TransactionSubStatus `json:"sub_status,omitempty"`
	Type *TransactionType `json:"type,omitempty"`
	FromType *TransactionAddressType `json:"from_type,omitempty"`
	FromAddress []TransactionAddress `json:"from_address,omitempty"`
	// From wallet info
	FromInfo *string `json:"from_info,omitempty"`
	ToType *TransactionAddressType `json:"to_type,omitempty"`
	ToAddress []TransactionAddress `json:"to_address,omitempty"`
	// To wallet info
	ToInfo *string `json:"to_info,omitempty"`
	Network *Network `json:"network,omitempty"`
	Txid *string `json:"txid,omitempty"`
	Tokens []TransactionToken `json:"tokens,omitempty"`
	Category []string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
	// Transaction creation time
	CreatedTime *float32 `json:"created_time,omitempty"`
	// Transaction update time
	UpdatedTime *float32 `json:"updated_time,omitempty"`
	// Transaction delegate address
	Delegate *string `json:"delegate,omitempty"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction() *Transaction {
	this := Transaction{}
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *Transaction) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *Transaction) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *Transaction) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *Transaction) GetWalletId() string {
	if o == nil || IsNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.WalletId) {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *Transaction) HasWalletId() bool {
	if o != nil && !IsNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *Transaction) SetWalletId(v string) {
	o.WalletId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *Transaction) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *Transaction) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *Transaction) SetRequestId(v string) {
	o.RequestId = &v
}

// GetCoboId returns the CoboId field value if set, zero value otherwise.
func (o *Transaction) GetCoboId() string {
	if o == nil || IsNil(o.CoboId) {
		var ret string
		return ret
	}
	return *o.CoboId
}

// GetCoboIdOk returns a tuple with the CoboId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetCoboIdOk() (*string, bool) {
	if o == nil || IsNil(o.CoboId) {
		return nil, false
	}
	return o.CoboId, true
}

// HasCoboId returns a boolean if a field has been set.
func (o *Transaction) HasCoboId() bool {
	if o != nil && !IsNil(o.CoboId) {
		return true
	}

	return false
}

// SetCoboId gets a reference to the given string and assigns it to the CoboId field.
func (o *Transaction) SetCoboId(v string) {
	o.CoboId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Transaction) GetStatus() TransactionStatus {
	if o == nil || IsNil(o.Status) {
		var ret TransactionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetStatusOk() (*TransactionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Transaction) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TransactionStatus and assigns it to the Status field.
func (o *Transaction) SetStatus(v TransactionStatus) {
	o.Status = &v
}

// GetSubStatus returns the SubStatus field value if set, zero value otherwise.
func (o *Transaction) GetSubStatus() TransactionSubStatus {
	if o == nil || IsNil(o.SubStatus) {
		var ret TransactionSubStatus
		return ret
	}
	return *o.SubStatus
}

// GetSubStatusOk returns a tuple with the SubStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetSubStatusOk() (*TransactionSubStatus, bool) {
	if o == nil || IsNil(o.SubStatus) {
		return nil, false
	}
	return o.SubStatus, true
}

// HasSubStatus returns a boolean if a field has been set.
func (o *Transaction) HasSubStatus() bool {
	if o != nil && !IsNil(o.SubStatus) {
		return true
	}

	return false
}

// SetSubStatus gets a reference to the given TransactionSubStatus and assigns it to the SubStatus field.
func (o *Transaction) SetSubStatus(v TransactionSubStatus) {
	o.SubStatus = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Transaction) GetType() TransactionType {
	if o == nil || IsNil(o.Type) {
		var ret TransactionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTypeOk() (*TransactionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Transaction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TransactionType and assigns it to the Type field.
func (o *Transaction) SetType(v TransactionType) {
	o.Type = &v
}

// GetFromType returns the FromType field value if set, zero value otherwise.
func (o *Transaction) GetFromType() TransactionAddressType {
	if o == nil || IsNil(o.FromType) {
		var ret TransactionAddressType
		return ret
	}
	return *o.FromType
}

// GetFromTypeOk returns a tuple with the FromType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetFromTypeOk() (*TransactionAddressType, bool) {
	if o == nil || IsNil(o.FromType) {
		return nil, false
	}
	return o.FromType, true
}

// HasFromType returns a boolean if a field has been set.
func (o *Transaction) HasFromType() bool {
	if o != nil && !IsNil(o.FromType) {
		return true
	}

	return false
}

// SetFromType gets a reference to the given TransactionAddressType and assigns it to the FromType field.
func (o *Transaction) SetFromType(v TransactionAddressType) {
	o.FromType = &v
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *Transaction) GetFromAddress() []TransactionAddress {
	if o == nil || IsNil(o.FromAddress) {
		var ret []TransactionAddress
		return ret
	}
	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetFromAddressOk() ([]TransactionAddress, bool) {
	if o == nil || IsNil(o.FromAddress) {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *Transaction) HasFromAddress() bool {
	if o != nil && !IsNil(o.FromAddress) {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given []TransactionAddress and assigns it to the FromAddress field.
func (o *Transaction) SetFromAddress(v []TransactionAddress) {
	o.FromAddress = v
}

// GetFromInfo returns the FromInfo field value if set, zero value otherwise.
func (o *Transaction) GetFromInfo() string {
	if o == nil || IsNil(o.FromInfo) {
		var ret string
		return ret
	}
	return *o.FromInfo
}

// GetFromInfoOk returns a tuple with the FromInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetFromInfoOk() (*string, bool) {
	if o == nil || IsNil(o.FromInfo) {
		return nil, false
	}
	return o.FromInfo, true
}

// HasFromInfo returns a boolean if a field has been set.
func (o *Transaction) HasFromInfo() bool {
	if o != nil && !IsNil(o.FromInfo) {
		return true
	}

	return false
}

// SetFromInfo gets a reference to the given string and assigns it to the FromInfo field.
func (o *Transaction) SetFromInfo(v string) {
	o.FromInfo = &v
}

// GetToType returns the ToType field value if set, zero value otherwise.
func (o *Transaction) GetToType() TransactionAddressType {
	if o == nil || IsNil(o.ToType) {
		var ret TransactionAddressType
		return ret
	}
	return *o.ToType
}

// GetToTypeOk returns a tuple with the ToType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetToTypeOk() (*TransactionAddressType, bool) {
	if o == nil || IsNil(o.ToType) {
		return nil, false
	}
	return o.ToType, true
}

// HasToType returns a boolean if a field has been set.
func (o *Transaction) HasToType() bool {
	if o != nil && !IsNil(o.ToType) {
		return true
	}

	return false
}

// SetToType gets a reference to the given TransactionAddressType and assigns it to the ToType field.
func (o *Transaction) SetToType(v TransactionAddressType) {
	o.ToType = &v
}

// GetToAddress returns the ToAddress field value if set, zero value otherwise.
func (o *Transaction) GetToAddress() []TransactionAddress {
	if o == nil || IsNil(o.ToAddress) {
		var ret []TransactionAddress
		return ret
	}
	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetToAddressOk() ([]TransactionAddress, bool) {
	if o == nil || IsNil(o.ToAddress) {
		return nil, false
	}
	return o.ToAddress, true
}

// HasToAddress returns a boolean if a field has been set.
func (o *Transaction) HasToAddress() bool {
	if o != nil && !IsNil(o.ToAddress) {
		return true
	}

	return false
}

// SetToAddress gets a reference to the given []TransactionAddress and assigns it to the ToAddress field.
func (o *Transaction) SetToAddress(v []TransactionAddress) {
	o.ToAddress = v
}

// GetToInfo returns the ToInfo field value if set, zero value otherwise.
func (o *Transaction) GetToInfo() string {
	if o == nil || IsNil(o.ToInfo) {
		var ret string
		return ret
	}
	return *o.ToInfo
}

// GetToInfoOk returns a tuple with the ToInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetToInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ToInfo) {
		return nil, false
	}
	return o.ToInfo, true
}

// HasToInfo returns a boolean if a field has been set.
func (o *Transaction) HasToInfo() bool {
	if o != nil && !IsNil(o.ToInfo) {
		return true
	}

	return false
}

// SetToInfo gets a reference to the given string and assigns it to the ToInfo field.
func (o *Transaction) SetToInfo(v string) {
	o.ToInfo = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *Transaction) GetNetwork() Network {
	if o == nil || IsNil(o.Network) {
		var ret Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetNetworkOk() (*Network, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *Transaction) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given Network and assigns it to the Network field.
func (o *Transaction) SetNetwork(v Network) {
	o.Network = &v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *Transaction) GetTxid() string {
	if o == nil || IsNil(o.Txid) {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTxidOk() (*string, bool) {
	if o == nil || IsNil(o.Txid) {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *Transaction) HasTxid() bool {
	if o != nil && !IsNil(o.Txid) {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *Transaction) SetTxid(v string) {
	o.Txid = &v
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *Transaction) GetTokens() []TransactionToken {
	if o == nil || IsNil(o.Tokens) {
		var ret []TransactionToken
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTokensOk() ([]TransactionToken, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *Transaction) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given []TransactionToken and assigns it to the Tokens field.
func (o *Transaction) SetTokens(v []TransactionToken) {
	o.Tokens = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Transaction) GetCategory() []string {
	if o == nil || IsNil(o.Category) {
		var ret []string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetCategoryOk() ([]string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Transaction) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given []string and assigns it to the Category field.
func (o *Transaction) SetCategory(v []string) {
	o.Category = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Transaction) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Transaction) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Transaction) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Transaction) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Transaction) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *Transaction) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *Transaction) GetUpdatedTime() float32 {
	if o == nil || IsNil(o.UpdatedTime) {
		var ret float32
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetUpdatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *Transaction) HasUpdatedTime() bool {
	if o != nil && !IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given float32 and assigns it to the UpdatedTime field.
func (o *Transaction) SetUpdatedTime(v float32) {
	o.UpdatedTime = &v
}

// GetDelegate returns the Delegate field value if set, zero value otherwise.
func (o *Transaction) GetDelegate() string {
	if o == nil || IsNil(o.Delegate) {
		var ret string
		return ret
	}
	return *o.Delegate
}

// GetDelegateOk returns a tuple with the Delegate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetDelegateOk() (*string, bool) {
	if o == nil || IsNil(o.Delegate) {
		return nil, false
	}
	return o.Delegate, true
}

// HasDelegate returns a boolean if a field has been set.
func (o *Transaction) HasDelegate() bool {
	if o != nil && !IsNil(o.Delegate) {
		return true
	}

	return false
}

// SetDelegate gets a reference to the given string and assigns it to the Delegate field.
func (o *Transaction) SetDelegate(v string) {
	o.Delegate = &v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if !IsNil(o.WalletId) {
		toSerialize["wallet_id"] = o.WalletId
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.CoboId) {
		toSerialize["cobo_id"] = o.CoboId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubStatus) {
		toSerialize["sub_status"] = o.SubStatus
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.FromType) {
		toSerialize["from_type"] = o.FromType
	}
	if !IsNil(o.FromAddress) {
		toSerialize["from_address"] = o.FromAddress
	}
	if !IsNil(o.FromInfo) {
		toSerialize["from_info"] = o.FromInfo
	}
	if !IsNil(o.ToType) {
		toSerialize["to_type"] = o.ToType
	}
	if !IsNil(o.ToAddress) {
		toSerialize["to_address"] = o.ToAddress
	}
	if !IsNil(o.ToInfo) {
		toSerialize["to_info"] = o.ToInfo
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Txid) {
		toSerialize["txid"] = o.Txid
	}
	if !IsNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.UpdatedTime) {
		toSerialize["updated_time"] = o.UpdatedTime
	}
	if !IsNil(o.Delegate) {
		toSerialize["delegate"] = o.Delegate
	}
	return toSerialize, nil
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


