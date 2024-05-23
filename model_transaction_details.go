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

// checks if the TransactionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionDetails{}

// TransactionDetails struct for TransactionDetails
type TransactionDetails struct {
	// Unique transaction ID
	TransactionId string `json:"transaction_id"`
	// Wallet ID
	WalletId string `json:"wallet_id"`
	// Request ID
	RequestId *string `json:"request_id,omitempty"`
	// Cobo ID
	CoboId string `json:"cobo_id"`
	// Transaction initiator
	Initiator *string `json:"initiator,omitempty"`
	// Transaction hash.
	TransactionHash *string `json:"transaction_hash,omitempty"`
	Status TransactionStatus `json:"status"`
	SubStatus *TransactionSubStatus `json:"sub_status,omitempty"`
	Type TransactionType `json:"type"`
	Source TransactionSource `json:"source"`
	Destination TransactionDestination `json:"destination"`
	// The blockchain on which the token operates.
	ChainId *string `json:"chain_id,omitempty"`
	ExchangeId *ExchangeId `json:"exchange_id,omitempty"`
	Tokens []TransactionToken `json:"tokens,omitempty"`
	Fee *TransactionFee `json:"fee,omitempty"`
	Category []string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
	// Transaction confirmed number
	ConfirmedNum *float32 `json:"confirmed_num,omitempty"`
	// Number of confirmations required for a transaction, such as 15 for ETH chain.
	ConfirmingThreshold *int32 `json:"confirming_threshold,omitempty"`
	// Transaction creation time
	CreatedTime float32 `json:"created_time"`
	// Transaction update time
	UpdatedTime float32 `json:"updated_time"`
	Approvers []TransactionApprover `json:"approvers,omitempty"`
	Signers []TransactionSigner `json:"signers,omitempty"`
	// Transaction nonce
	Nonce *int32 `json:"nonce,omitempty"`
	// Replace by transaction hash
	ReplacedBy *string `json:"replaced_by,omitempty"`
	// Fueled by address
	FueledBy *string `json:"fueled_by,omitempty"`
	TokenApproval *TransactionTokeApproval `json:"token_approval,omitempty"`
	// Transaction raw message
	Message *string `json:"message,omitempty"`
	// Transaction message signing algorithm
	Algorithm *string `json:"algorithm,omitempty"`
	Timeline []TransactionTimeline `json:"timeline,omitempty"`
}

type _TransactionDetails TransactionDetails

// NewTransactionDetails instantiates a new TransactionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionDetails(transactionId string, walletId string, coboId string, status TransactionStatus, type_ TransactionType, source TransactionSource, destination TransactionDestination, createdTime float32, updatedTime float32) *TransactionDetails {
	this := TransactionDetails{}
	this.TransactionId = transactionId
	this.WalletId = walletId
	this.CoboId = coboId
	this.Status = status
	this.Type = type_
	this.Source = source
	this.Destination = destination
	this.CreatedTime = createdTime
	this.UpdatedTime = updatedTime
	return &this
}

// NewTransactionDetailsWithDefaults instantiates a new TransactionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDetailsWithDefaults() *TransactionDetails {
	this := TransactionDetails{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *TransactionDetails) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *TransactionDetails) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetWalletId returns the WalletId field value
func (o *TransactionDetails) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *TransactionDetails) SetWalletId(v string) {
	o.WalletId = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *TransactionDetails) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *TransactionDetails) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *TransactionDetails) SetRequestId(v string) {
	o.RequestId = &v
}

// GetCoboId returns the CoboId field value
func (o *TransactionDetails) GetCoboId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CoboId
}

// GetCoboIdOk returns a tuple with the CoboId field value
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetCoboIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CoboId, true
}

// SetCoboId sets field value
func (o *TransactionDetails) SetCoboId(v string) {
	o.CoboId = v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *TransactionDetails) GetInitiator() string {
	if o == nil || IsNil(o.Initiator) {
		var ret string
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetInitiatorOk() (*string, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *TransactionDetails) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given string and assigns it to the Initiator field.
func (o *TransactionDetails) SetInitiator(v string) {
	o.Initiator = &v
}

// GetTransactionHash returns the TransactionHash field value if set, zero value otherwise.
func (o *TransactionDetails) GetTransactionHash() string {
	if o == nil || IsNil(o.TransactionHash) {
		var ret string
		return ret
	}
	return *o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetTransactionHashOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionHash) {
		return nil, false
	}
	return o.TransactionHash, true
}

// HasTransactionHash returns a boolean if a field has been set.
func (o *TransactionDetails) HasTransactionHash() bool {
	if o != nil && !IsNil(o.TransactionHash) {
		return true
	}

	return false
}

// SetTransactionHash gets a reference to the given string and assigns it to the TransactionHash field.
func (o *TransactionDetails) SetTransactionHash(v string) {
	o.TransactionHash = &v
}

// GetStatus returns the Status field value
func (o *TransactionDetails) GetStatus() TransactionStatus {
	if o == nil {
		var ret TransactionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetStatusOk() (*TransactionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransactionDetails) SetStatus(v TransactionStatus) {
	o.Status = v
}

// GetSubStatus returns the SubStatus field value if set, zero value otherwise.
func (o *TransactionDetails) GetSubStatus() TransactionSubStatus {
	if o == nil || IsNil(o.SubStatus) {
		var ret TransactionSubStatus
		return ret
	}
	return *o.SubStatus
}

// GetSubStatusOk returns a tuple with the SubStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetSubStatusOk() (*TransactionSubStatus, bool) {
	if o == nil || IsNil(o.SubStatus) {
		return nil, false
	}
	return o.SubStatus, true
}

// HasSubStatus returns a boolean if a field has been set.
func (o *TransactionDetails) HasSubStatus() bool {
	if o != nil && !IsNil(o.SubStatus) {
		return true
	}

	return false
}

// SetSubStatus gets a reference to the given TransactionSubStatus and assigns it to the SubStatus field.
func (o *TransactionDetails) SetSubStatus(v TransactionSubStatus) {
	o.SubStatus = &v
}

// GetType returns the Type field value
func (o *TransactionDetails) GetType() TransactionType {
	if o == nil {
		var ret TransactionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetTypeOk() (*TransactionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionDetails) SetType(v TransactionType) {
	o.Type = v
}

// GetSource returns the Source field value
func (o *TransactionDetails) GetSource() TransactionSource {
	if o == nil {
		var ret TransactionSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetSourceOk() (*TransactionSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *TransactionDetails) SetSource(v TransactionSource) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *TransactionDetails) GetDestination() TransactionDestination {
	if o == nil {
		var ret TransactionDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetDestinationOk() (*TransactionDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *TransactionDetails) SetDestination(v TransactionDestination) {
	o.Destination = v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *TransactionDetails) GetChainId() string {
	if o == nil || IsNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetChainIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *TransactionDetails) HasChainId() bool {
	if o != nil && !IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *TransactionDetails) SetChainId(v string) {
	o.ChainId = &v
}

// GetExchangeId returns the ExchangeId field value if set, zero value otherwise.
func (o *TransactionDetails) GetExchangeId() ExchangeId {
	if o == nil || IsNil(o.ExchangeId) {
		var ret ExchangeId
		return ret
	}
	return *o.ExchangeId
}

// GetExchangeIdOk returns a tuple with the ExchangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetExchangeIdOk() (*ExchangeId, bool) {
	if o == nil || IsNil(o.ExchangeId) {
		return nil, false
	}
	return o.ExchangeId, true
}

// HasExchangeId returns a boolean if a field has been set.
func (o *TransactionDetails) HasExchangeId() bool {
	if o != nil && !IsNil(o.ExchangeId) {
		return true
	}

	return false
}

// SetExchangeId gets a reference to the given ExchangeId and assigns it to the ExchangeId field.
func (o *TransactionDetails) SetExchangeId(v ExchangeId) {
	o.ExchangeId = &v
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *TransactionDetails) GetTokens() []TransactionToken {
	if o == nil || IsNil(o.Tokens) {
		var ret []TransactionToken
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetTokensOk() ([]TransactionToken, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *TransactionDetails) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given []TransactionToken and assigns it to the Tokens field.
func (o *TransactionDetails) SetTokens(v []TransactionToken) {
	o.Tokens = v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *TransactionDetails) GetFee() TransactionFee {
	if o == nil || IsNil(o.Fee) {
		var ret TransactionFee
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetFeeOk() (*TransactionFee, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *TransactionDetails) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given TransactionFee and assigns it to the Fee field.
func (o *TransactionDetails) SetFee(v TransactionFee) {
	o.Fee = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *TransactionDetails) GetCategory() []string {
	if o == nil || IsNil(o.Category) {
		var ret []string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetCategoryOk() ([]string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *TransactionDetails) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given []string and assigns it to the Category field.
func (o *TransactionDetails) SetCategory(v []string) {
	o.Category = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransactionDetails) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransactionDetails) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransactionDetails) SetDescription(v string) {
	o.Description = &v
}

// GetConfirmedNum returns the ConfirmedNum field value if set, zero value otherwise.
func (o *TransactionDetails) GetConfirmedNum() float32 {
	if o == nil || IsNil(o.ConfirmedNum) {
		var ret float32
		return ret
	}
	return *o.ConfirmedNum
}

// GetConfirmedNumOk returns a tuple with the ConfirmedNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetConfirmedNumOk() (*float32, bool) {
	if o == nil || IsNil(o.ConfirmedNum) {
		return nil, false
	}
	return o.ConfirmedNum, true
}

// HasConfirmedNum returns a boolean if a field has been set.
func (o *TransactionDetails) HasConfirmedNum() bool {
	if o != nil && !IsNil(o.ConfirmedNum) {
		return true
	}

	return false
}

// SetConfirmedNum gets a reference to the given float32 and assigns it to the ConfirmedNum field.
func (o *TransactionDetails) SetConfirmedNum(v float32) {
	o.ConfirmedNum = &v
}

// GetConfirmingThreshold returns the ConfirmingThreshold field value if set, zero value otherwise.
func (o *TransactionDetails) GetConfirmingThreshold() int32 {
	if o == nil || IsNil(o.ConfirmingThreshold) {
		var ret int32
		return ret
	}
	return *o.ConfirmingThreshold
}

// GetConfirmingThresholdOk returns a tuple with the ConfirmingThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetConfirmingThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.ConfirmingThreshold) {
		return nil, false
	}
	return o.ConfirmingThreshold, true
}

// HasConfirmingThreshold returns a boolean if a field has been set.
func (o *TransactionDetails) HasConfirmingThreshold() bool {
	if o != nil && !IsNil(o.ConfirmingThreshold) {
		return true
	}

	return false
}

// SetConfirmingThreshold gets a reference to the given int32 and assigns it to the ConfirmingThreshold field.
func (o *TransactionDetails) SetConfirmingThreshold(v int32) {
	o.ConfirmingThreshold = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *TransactionDetails) GetCreatedTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetCreatedTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *TransactionDetails) SetCreatedTime(v float32) {
	o.CreatedTime = v
}

// GetUpdatedTime returns the UpdatedTime field value
func (o *TransactionDetails) GetUpdatedTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetUpdatedTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTime, true
}

// SetUpdatedTime sets field value
func (o *TransactionDetails) SetUpdatedTime(v float32) {
	o.UpdatedTime = v
}

// GetApprovers returns the Approvers field value if set, zero value otherwise.
func (o *TransactionDetails) GetApprovers() []TransactionApprover {
	if o == nil || IsNil(o.Approvers) {
		var ret []TransactionApprover
		return ret
	}
	return o.Approvers
}

// GetApproversOk returns a tuple with the Approvers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetApproversOk() ([]TransactionApprover, bool) {
	if o == nil || IsNil(o.Approvers) {
		return nil, false
	}
	return o.Approvers, true
}

// HasApprovers returns a boolean if a field has been set.
func (o *TransactionDetails) HasApprovers() bool {
	if o != nil && !IsNil(o.Approvers) {
		return true
	}

	return false
}

// SetApprovers gets a reference to the given []TransactionApprover and assigns it to the Approvers field.
func (o *TransactionDetails) SetApprovers(v []TransactionApprover) {
	o.Approvers = v
}

// GetSigners returns the Signers field value if set, zero value otherwise.
func (o *TransactionDetails) GetSigners() []TransactionSigner {
	if o == nil || IsNil(o.Signers) {
		var ret []TransactionSigner
		return ret
	}
	return o.Signers
}

// GetSignersOk returns a tuple with the Signers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetSignersOk() ([]TransactionSigner, bool) {
	if o == nil || IsNil(o.Signers) {
		return nil, false
	}
	return o.Signers, true
}

// HasSigners returns a boolean if a field has been set.
func (o *TransactionDetails) HasSigners() bool {
	if o != nil && !IsNil(o.Signers) {
		return true
	}

	return false
}

// SetSigners gets a reference to the given []TransactionSigner and assigns it to the Signers field.
func (o *TransactionDetails) SetSigners(v []TransactionSigner) {
	o.Signers = v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *TransactionDetails) GetNonce() int32 {
	if o == nil || IsNil(o.Nonce) {
		var ret int32
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetNonceOk() (*int32, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *TransactionDetails) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given int32 and assigns it to the Nonce field.
func (o *TransactionDetails) SetNonce(v int32) {
	o.Nonce = &v
}

// GetReplacedBy returns the ReplacedBy field value if set, zero value otherwise.
func (o *TransactionDetails) GetReplacedBy() string {
	if o == nil || IsNil(o.ReplacedBy) {
		var ret string
		return ret
	}
	return *o.ReplacedBy
}

// GetReplacedByOk returns a tuple with the ReplacedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetReplacedByOk() (*string, bool) {
	if o == nil || IsNil(o.ReplacedBy) {
		return nil, false
	}
	return o.ReplacedBy, true
}

// HasReplacedBy returns a boolean if a field has been set.
func (o *TransactionDetails) HasReplacedBy() bool {
	if o != nil && !IsNil(o.ReplacedBy) {
		return true
	}

	return false
}

// SetReplacedBy gets a reference to the given string and assigns it to the ReplacedBy field.
func (o *TransactionDetails) SetReplacedBy(v string) {
	o.ReplacedBy = &v
}

// GetFueledBy returns the FueledBy field value if set, zero value otherwise.
func (o *TransactionDetails) GetFueledBy() string {
	if o == nil || IsNil(o.FueledBy) {
		var ret string
		return ret
	}
	return *o.FueledBy
}

// GetFueledByOk returns a tuple with the FueledBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetFueledByOk() (*string, bool) {
	if o == nil || IsNil(o.FueledBy) {
		return nil, false
	}
	return o.FueledBy, true
}

// HasFueledBy returns a boolean if a field has been set.
func (o *TransactionDetails) HasFueledBy() bool {
	if o != nil && !IsNil(o.FueledBy) {
		return true
	}

	return false
}

// SetFueledBy gets a reference to the given string and assigns it to the FueledBy field.
func (o *TransactionDetails) SetFueledBy(v string) {
	o.FueledBy = &v
}

// GetTokenApproval returns the TokenApproval field value if set, zero value otherwise.
func (o *TransactionDetails) GetTokenApproval() TransactionTokeApproval {
	if o == nil || IsNil(o.TokenApproval) {
		var ret TransactionTokeApproval
		return ret
	}
	return *o.TokenApproval
}

// GetTokenApprovalOk returns a tuple with the TokenApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetTokenApprovalOk() (*TransactionTokeApproval, bool) {
	if o == nil || IsNil(o.TokenApproval) {
		return nil, false
	}
	return o.TokenApproval, true
}

// HasTokenApproval returns a boolean if a field has been set.
func (o *TransactionDetails) HasTokenApproval() bool {
	if o != nil && !IsNil(o.TokenApproval) {
		return true
	}

	return false
}

// SetTokenApproval gets a reference to the given TransactionTokeApproval and assigns it to the TokenApproval field.
func (o *TransactionDetails) SetTokenApproval(v TransactionTokeApproval) {
	o.TokenApproval = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TransactionDetails) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TransactionDetails) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TransactionDetails) SetMessage(v string) {
	o.Message = &v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *TransactionDetails) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *TransactionDetails) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *TransactionDetails) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetTimeline returns the Timeline field value if set, zero value otherwise.
func (o *TransactionDetails) GetTimeline() []TransactionTimeline {
	if o == nil || IsNil(o.Timeline) {
		var ret []TransactionTimeline
		return ret
	}
	return o.Timeline
}

// GetTimelineOk returns a tuple with the Timeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDetails) GetTimelineOk() ([]TransactionTimeline, bool) {
	if o == nil || IsNil(o.Timeline) {
		return nil, false
	}
	return o.Timeline, true
}

// HasTimeline returns a boolean if a field has been set.
func (o *TransactionDetails) HasTimeline() bool {
	if o != nil && !IsNil(o.Timeline) {
		return true
	}

	return false
}

// SetTimeline gets a reference to the given []TransactionTimeline and assigns it to the Timeline field.
func (o *TransactionDetails) SetTimeline(v []TransactionTimeline) {
	o.Timeline = v
}

func (o TransactionDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction_id"] = o.TransactionId
	toSerialize["wallet_id"] = o.WalletId
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	toSerialize["cobo_id"] = o.CoboId
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	if !IsNil(o.TransactionHash) {
		toSerialize["transaction_hash"] = o.TransactionHash
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.SubStatus) {
		toSerialize["sub_status"] = o.SubStatus
	}
	toSerialize["type"] = o.Type
	toSerialize["source"] = o.Source
	toSerialize["destination"] = o.Destination
	if !IsNil(o.ChainId) {
		toSerialize["chain_id"] = o.ChainId
	}
	if !IsNil(o.ExchangeId) {
		toSerialize["exchange_id"] = o.ExchangeId
	}
	if !IsNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ConfirmedNum) {
		toSerialize["confirmed_num"] = o.ConfirmedNum
	}
	if !IsNil(o.ConfirmingThreshold) {
		toSerialize["confirming_threshold"] = o.ConfirmingThreshold
	}
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["updated_time"] = o.UpdatedTime
	if !IsNil(o.Approvers) {
		toSerialize["approvers"] = o.Approvers
	}
	if !IsNil(o.Signers) {
		toSerialize["signers"] = o.Signers
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.ReplacedBy) {
		toSerialize["replaced_by"] = o.ReplacedBy
	}
	if !IsNil(o.FueledBy) {
		toSerialize["fueled_by"] = o.FueledBy
	}
	if !IsNil(o.TokenApproval) {
		toSerialize["token_approval"] = o.TokenApproval
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.Timeline) {
		toSerialize["timeline"] = o.Timeline
	}
	return toSerialize, nil
}

func (o *TransactionDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction_id",
		"wallet_id",
		"cobo_id",
		"status",
		"type",
		"source",
		"destination",
		"created_time",
		"updated_time",
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

	varTransactionDetails := _TransactionDetails{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionDetails)

	if err != nil {
		return err
	}

	*o = TransactionDetails(varTransactionDetails)

	return err
}

type NullableTransactionDetails struct {
	value *TransactionDetails
	isSet bool
}

func (v NullableTransactionDetails) Get() *TransactionDetails {
	return v.value
}

func (v *NullableTransactionDetails) Set(val *TransactionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDetails(val *TransactionDetails) *NullableTransactionDetails {
	return &NullableTransactionDetails{value: val, isSet: true}
}

func (v NullableTransactionDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


