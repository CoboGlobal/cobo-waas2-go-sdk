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

// checks if the Transaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transaction{}

// Transaction The information about a transaction.
type Transaction struct {
	// The transaction ID.
	TransactionId string `json:"transaction_id"`
	// The Cobo ID, which can be used to track a transaction.
	CoboId *string `json:"cobo_id,omitempty"`
	// The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization.
	RequestId *string `json:"request_id,omitempty"`
	// For deposit transactions, this property represents the wallet ID of the transaction destination. For transactions of other types, this property represents the wallet ID of the transaction source.
	WalletId string `json:"wallet_id"`
	Type *TransactionType `json:"type,omitempty"`
	Status TransactionStatus `json:"status"`
	SubStatus *TransactionSubStatus `json:"sub_status,omitempty"`
	// (This property is applicable to approval failures and signature failures only) The reason why the transaction failed.
	FailedReason *string `json:"failed_reason,omitempty"`
	// The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains).
	ChainId *string `json:"chain_id,omitempty"`
	// The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens).
	TokenId *string `json:"token_id,omitempty"`
	// (This concept applies to Exchange Wallets only) The asset ID. An asset ID is the unique identifier of the asset held within your linked exchange account.
	AssetId *string `json:"asset_id,omitempty"`
	Source TransactionSource `json:"source"`
	Destination TransactionDestination `json:"destination"`
	Result *TransactionResult `json:"result,omitempty"`
	Fee *TransactionFee `json:"fee,omitempty"`
	// The transaction initiator.
	Initiator *string `json:"initiator,omitempty"`
	InitiatorType TransactionInitiatorType `json:"initiator_type"`
	// The number of confirmations this transaction has received.
	ConfirmedNum *int32 `json:"confirmed_num,omitempty"`
	// The minimum number of confirmations required to deem a transaction secure. The common threshold is 6 for a Bitcoin transaction.
	ConfirmingThreshold *int32 `json:"confirming_threshold,omitempty"`
	// The transaction hash.
	TransactionHash *string `json:"transaction_hash,omitempty"`
	BlockInfo *TransactionBlockInfo `json:"block_info,omitempty"`
	RawTxInfo *TransactionRawTxInfo `json:"raw_tx_info,omitempty"`
	Replacement *TransactionReplacement `json:"replacement,omitempty"`
	// A custom transaction category for you to identify your transfers more easily.
	Category []string `json:"category,omitempty"`
	// The description for your transaction.
	Description *string `json:"description,omitempty"`
	// Whether the transaction was executed as a [Cobo Loop](https://manuals.cobo.com/en/portal/custodial-wallets/cobo-loop) transfer. - `true`: The transaction was executed as a Cobo Loop transfer. - `false`: The transaction was not executed as a Cobo Loop transfer. 
	IsLoop *bool `json:"is_loop,omitempty"`
	// The transaction category defined by Cobo. Possible values include:  - `AutoSweep`: An auto-sweep transaction. - `AutoFueling`: A transaction where Fee Station pays transaction fees to an address within your wallet. - `AutoFuelingRefund`: A refund for an auto-fueling transaction. - `SafeTxMessage`: A message signing transaction to authorize a Smart Contract Wallet (Safe\\{Wallet\\}) transaction. - `BillPayment`: A transaction to pay Cobo bills through Fee Station. - `BillRefund`: A refund for a previously made bill payment. - `CommissionFeeCharge`: A transaction to charge commission fees via Fee Station. - `CommissionFeeRefund`: A refund of previously charged commission fees. 
	CoboCategory []string `json:"cobo_category,omitempty"`
	// A list of JSON-encoded strings containing structured, business-specific extra information for the transaction. Each item corresponds to a specific data type, indicated by the `extra_type` field in the JSON object (for example, \"BabylonBusinessInfo\", \"BtcAddressInfo\"). 
	Extra []string `json:"extra,omitempty"`
	FuelingInfo *TransactionFuelingInfo `json:"fueling_info,omitempty"`
	// The time when the transaction was created, in Unix timestamp format, measured in milliseconds.
	CreatedTimestamp int64 `json:"created_timestamp"`
	// The time when the transaction was updated, in Unix timestamp format, measured in milliseconds.
	UpdatedTimestamp int64 `json:"updated_timestamp"`
}

type _Transaction Transaction

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(transactionId string, walletId string, status TransactionStatus, source TransactionSource, destination TransactionDestination, initiatorType TransactionInitiatorType, createdTimestamp int64, updatedTimestamp int64) *Transaction {
	this := Transaction{}
	this.TransactionId = transactionId
	this.WalletId = walletId
	this.Status = status
	this.Source = source
	this.Destination = destination
	this.InitiatorType = initiatorType
	this.CreatedTimestamp = createdTimestamp
	this.UpdatedTimestamp = updatedTimestamp
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *Transaction) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *Transaction) SetTransactionId(v string) {
	o.TransactionId = v
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

// GetWalletId returns the WalletId field value
func (o *Transaction) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *Transaction) SetWalletId(v string) {
	o.WalletId = v
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

// GetStatus returns the Status field value
func (o *Transaction) GetStatus() TransactionStatus {
	if o == nil {
		var ret TransactionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetStatusOk() (*TransactionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Transaction) SetStatus(v TransactionStatus) {
	o.Status = v
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

// GetFailedReason returns the FailedReason field value if set, zero value otherwise.
func (o *Transaction) GetFailedReason() string {
	if o == nil || IsNil(o.FailedReason) {
		var ret string
		return ret
	}
	return *o.FailedReason
}

// GetFailedReasonOk returns a tuple with the FailedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetFailedReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailedReason) {
		return nil, false
	}
	return o.FailedReason, true
}

// HasFailedReason returns a boolean if a field has been set.
func (o *Transaction) HasFailedReason() bool {
	if o != nil && !IsNil(o.FailedReason) {
		return true
	}

	return false
}

// SetFailedReason gets a reference to the given string and assigns it to the FailedReason field.
func (o *Transaction) SetFailedReason(v string) {
	o.FailedReason = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *Transaction) GetChainId() string {
	if o == nil || IsNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetChainIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *Transaction) HasChainId() bool {
	if o != nil && !IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *Transaction) SetChainId(v string) {
	o.ChainId = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *Transaction) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *Transaction) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *Transaction) SetTokenId(v string) {
	o.TokenId = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *Transaction) GetAssetId() string {
	if o == nil || IsNil(o.AssetId) {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *Transaction) HasAssetId() bool {
	if o != nil && !IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *Transaction) SetAssetId(v string) {
	o.AssetId = &v
}

// GetSource returns the Source field value
func (o *Transaction) GetSource() TransactionSource {
	if o == nil {
		var ret TransactionSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetSourceOk() (*TransactionSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *Transaction) SetSource(v TransactionSource) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *Transaction) GetDestination() TransactionDestination {
	if o == nil {
		var ret TransactionDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetDestinationOk() (*TransactionDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *Transaction) SetDestination(v TransactionDestination) {
	o.Destination = v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *Transaction) GetResult() TransactionResult {
	if o == nil || IsNil(o.Result) {
		var ret TransactionResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetResultOk() (*TransactionResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *Transaction) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given TransactionResult and assigns it to the Result field.
func (o *Transaction) SetResult(v TransactionResult) {
	o.Result = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *Transaction) GetFee() TransactionFee {
	if o == nil || IsNil(o.Fee) {
		var ret TransactionFee
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetFeeOk() (*TransactionFee, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *Transaction) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given TransactionFee and assigns it to the Fee field.
func (o *Transaction) SetFee(v TransactionFee) {
	o.Fee = &v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *Transaction) GetInitiator() string {
	if o == nil || IsNil(o.Initiator) {
		var ret string
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetInitiatorOk() (*string, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *Transaction) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given string and assigns it to the Initiator field.
func (o *Transaction) SetInitiator(v string) {
	o.Initiator = &v
}

// GetInitiatorType returns the InitiatorType field value
func (o *Transaction) GetInitiatorType() TransactionInitiatorType {
	if o == nil {
		var ret TransactionInitiatorType
		return ret
	}

	return o.InitiatorType
}

// GetInitiatorTypeOk returns a tuple with the InitiatorType field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetInitiatorTypeOk() (*TransactionInitiatorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitiatorType, true
}

// SetInitiatorType sets field value
func (o *Transaction) SetInitiatorType(v TransactionInitiatorType) {
	o.InitiatorType = v
}

// GetConfirmedNum returns the ConfirmedNum field value if set, zero value otherwise.
func (o *Transaction) GetConfirmedNum() int32 {
	if o == nil || IsNil(o.ConfirmedNum) {
		var ret int32
		return ret
	}
	return *o.ConfirmedNum
}

// GetConfirmedNumOk returns a tuple with the ConfirmedNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetConfirmedNumOk() (*int32, bool) {
	if o == nil || IsNil(o.ConfirmedNum) {
		return nil, false
	}
	return o.ConfirmedNum, true
}

// HasConfirmedNum returns a boolean if a field has been set.
func (o *Transaction) HasConfirmedNum() bool {
	if o != nil && !IsNil(o.ConfirmedNum) {
		return true
	}

	return false
}

// SetConfirmedNum gets a reference to the given int32 and assigns it to the ConfirmedNum field.
func (o *Transaction) SetConfirmedNum(v int32) {
	o.ConfirmedNum = &v
}

// GetConfirmingThreshold returns the ConfirmingThreshold field value if set, zero value otherwise.
func (o *Transaction) GetConfirmingThreshold() int32 {
	if o == nil || IsNil(o.ConfirmingThreshold) {
		var ret int32
		return ret
	}
	return *o.ConfirmingThreshold
}

// GetConfirmingThresholdOk returns a tuple with the ConfirmingThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetConfirmingThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.ConfirmingThreshold) {
		return nil, false
	}
	return o.ConfirmingThreshold, true
}

// HasConfirmingThreshold returns a boolean if a field has been set.
func (o *Transaction) HasConfirmingThreshold() bool {
	if o != nil && !IsNil(o.ConfirmingThreshold) {
		return true
	}

	return false
}

// SetConfirmingThreshold gets a reference to the given int32 and assigns it to the ConfirmingThreshold field.
func (o *Transaction) SetConfirmingThreshold(v int32) {
	o.ConfirmingThreshold = &v
}

// GetTransactionHash returns the TransactionHash field value if set, zero value otherwise.
func (o *Transaction) GetTransactionHash() string {
	if o == nil || IsNil(o.TransactionHash) {
		var ret string
		return ret
	}
	return *o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionHashOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionHash) {
		return nil, false
	}
	return o.TransactionHash, true
}

// HasTransactionHash returns a boolean if a field has been set.
func (o *Transaction) HasTransactionHash() bool {
	if o != nil && !IsNil(o.TransactionHash) {
		return true
	}

	return false
}

// SetTransactionHash gets a reference to the given string and assigns it to the TransactionHash field.
func (o *Transaction) SetTransactionHash(v string) {
	o.TransactionHash = &v
}

// GetBlockInfo returns the BlockInfo field value if set, zero value otherwise.
func (o *Transaction) GetBlockInfo() TransactionBlockInfo {
	if o == nil || IsNil(o.BlockInfo) {
		var ret TransactionBlockInfo
		return ret
	}
	return *o.BlockInfo
}

// GetBlockInfoOk returns a tuple with the BlockInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetBlockInfoOk() (*TransactionBlockInfo, bool) {
	if o == nil || IsNil(o.BlockInfo) {
		return nil, false
	}
	return o.BlockInfo, true
}

// HasBlockInfo returns a boolean if a field has been set.
func (o *Transaction) HasBlockInfo() bool {
	if o != nil && !IsNil(o.BlockInfo) {
		return true
	}

	return false
}

// SetBlockInfo gets a reference to the given TransactionBlockInfo and assigns it to the BlockInfo field.
func (o *Transaction) SetBlockInfo(v TransactionBlockInfo) {
	o.BlockInfo = &v
}

// GetRawTxInfo returns the RawTxInfo field value if set, zero value otherwise.
func (o *Transaction) GetRawTxInfo() TransactionRawTxInfo {
	if o == nil || IsNil(o.RawTxInfo) {
		var ret TransactionRawTxInfo
		return ret
	}
	return *o.RawTxInfo
}

// GetRawTxInfoOk returns a tuple with the RawTxInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetRawTxInfoOk() (*TransactionRawTxInfo, bool) {
	if o == nil || IsNil(o.RawTxInfo) {
		return nil, false
	}
	return o.RawTxInfo, true
}

// HasRawTxInfo returns a boolean if a field has been set.
func (o *Transaction) HasRawTxInfo() bool {
	if o != nil && !IsNil(o.RawTxInfo) {
		return true
	}

	return false
}

// SetRawTxInfo gets a reference to the given TransactionRawTxInfo and assigns it to the RawTxInfo field.
func (o *Transaction) SetRawTxInfo(v TransactionRawTxInfo) {
	o.RawTxInfo = &v
}

// GetReplacement returns the Replacement field value if set, zero value otherwise.
func (o *Transaction) GetReplacement() TransactionReplacement {
	if o == nil || IsNil(o.Replacement) {
		var ret TransactionReplacement
		return ret
	}
	return *o.Replacement
}

// GetReplacementOk returns a tuple with the Replacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetReplacementOk() (*TransactionReplacement, bool) {
	if o == nil || IsNil(o.Replacement) {
		return nil, false
	}
	return o.Replacement, true
}

// HasReplacement returns a boolean if a field has been set.
func (o *Transaction) HasReplacement() bool {
	if o != nil && !IsNil(o.Replacement) {
		return true
	}

	return false
}

// SetReplacement gets a reference to the given TransactionReplacement and assigns it to the Replacement field.
func (o *Transaction) SetReplacement(v TransactionReplacement) {
	o.Replacement = &v
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

// GetIsLoop returns the IsLoop field value if set, zero value otherwise.
func (o *Transaction) GetIsLoop() bool {
	if o == nil || IsNil(o.IsLoop) {
		var ret bool
		return ret
	}
	return *o.IsLoop
}

// GetIsLoopOk returns a tuple with the IsLoop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetIsLoopOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLoop) {
		return nil, false
	}
	return o.IsLoop, true
}

// HasIsLoop returns a boolean if a field has been set.
func (o *Transaction) HasIsLoop() bool {
	if o != nil && !IsNil(o.IsLoop) {
		return true
	}

	return false
}

// SetIsLoop gets a reference to the given bool and assigns it to the IsLoop field.
func (o *Transaction) SetIsLoop(v bool) {
	o.IsLoop = &v
}

// GetCoboCategory returns the CoboCategory field value if set, zero value otherwise.
func (o *Transaction) GetCoboCategory() []string {
	if o == nil || IsNil(o.CoboCategory) {
		var ret []string
		return ret
	}
	return o.CoboCategory
}

// GetCoboCategoryOk returns a tuple with the CoboCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetCoboCategoryOk() ([]string, bool) {
	if o == nil || IsNil(o.CoboCategory) {
		return nil, false
	}
	return o.CoboCategory, true
}

// HasCoboCategory returns a boolean if a field has been set.
func (o *Transaction) HasCoboCategory() bool {
	if o != nil && !IsNil(o.CoboCategory) {
		return true
	}

	return false
}

// SetCoboCategory gets a reference to the given []string and assigns it to the CoboCategory field.
func (o *Transaction) SetCoboCategory(v []string) {
	o.CoboCategory = v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *Transaction) GetExtra() []string {
	if o == nil || IsNil(o.Extra) {
		var ret []string
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetExtraOk() ([]string, bool) {
	if o == nil || IsNil(o.Extra) {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *Transaction) HasExtra() bool {
	if o != nil && !IsNil(o.Extra) {
		return true
	}

	return false
}

// SetExtra gets a reference to the given []string and assigns it to the Extra field.
func (o *Transaction) SetExtra(v []string) {
	o.Extra = v
}

// GetFuelingInfo returns the FuelingInfo field value if set, zero value otherwise.
func (o *Transaction) GetFuelingInfo() TransactionFuelingInfo {
	if o == nil || IsNil(o.FuelingInfo) {
		var ret TransactionFuelingInfo
		return ret
	}
	return *o.FuelingInfo
}

// GetFuelingInfoOk returns a tuple with the FuelingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetFuelingInfoOk() (*TransactionFuelingInfo, bool) {
	if o == nil || IsNil(o.FuelingInfo) {
		return nil, false
	}
	return o.FuelingInfo, true
}

// HasFuelingInfo returns a boolean if a field has been set.
func (o *Transaction) HasFuelingInfo() bool {
	if o != nil && !IsNil(o.FuelingInfo) {
		return true
	}

	return false
}

// SetFuelingInfo gets a reference to the given TransactionFuelingInfo and assigns it to the FuelingInfo field.
func (o *Transaction) SetFuelingInfo(v TransactionFuelingInfo) {
	o.FuelingInfo = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *Transaction) GetCreatedTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetCreatedTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *Transaction) SetCreatedTimestamp(v int64) {
	o.CreatedTimestamp = v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *Transaction) GetUpdatedTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetUpdatedTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *Transaction) SetUpdatedTimestamp(v int64) {
	o.UpdatedTimestamp = v
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
	toSerialize["transaction_id"] = o.TransactionId
	if !IsNil(o.CoboId) {
		toSerialize["cobo_id"] = o.CoboId
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	toSerialize["wallet_id"] = o.WalletId
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.SubStatus) {
		toSerialize["sub_status"] = o.SubStatus
	}
	if !IsNil(o.FailedReason) {
		toSerialize["failed_reason"] = o.FailedReason
	}
	if !IsNil(o.ChainId) {
		toSerialize["chain_id"] = o.ChainId
	}
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !IsNil(o.AssetId) {
		toSerialize["asset_id"] = o.AssetId
	}
	toSerialize["source"] = o.Source
	toSerialize["destination"] = o.Destination
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	toSerialize["initiator_type"] = o.InitiatorType
	if !IsNil(o.ConfirmedNum) {
		toSerialize["confirmed_num"] = o.ConfirmedNum
	}
	if !IsNil(o.ConfirmingThreshold) {
		toSerialize["confirming_threshold"] = o.ConfirmingThreshold
	}
	if !IsNil(o.TransactionHash) {
		toSerialize["transaction_hash"] = o.TransactionHash
	}
	if !IsNil(o.BlockInfo) {
		toSerialize["block_info"] = o.BlockInfo
	}
	if !IsNil(o.RawTxInfo) {
		toSerialize["raw_tx_info"] = o.RawTxInfo
	}
	if !IsNil(o.Replacement) {
		toSerialize["replacement"] = o.Replacement
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsLoop) {
		toSerialize["is_loop"] = o.IsLoop
	}
	if !IsNil(o.CoboCategory) {
		toSerialize["cobo_category"] = o.CoboCategory
	}
	if !IsNil(o.Extra) {
		toSerialize["extra"] = o.Extra
	}
	if !IsNil(o.FuelingInfo) {
		toSerialize["fueling_info"] = o.FuelingInfo
	}
	toSerialize["created_timestamp"] = o.CreatedTimestamp
	toSerialize["updated_timestamp"] = o.UpdatedTimestamp
	return toSerialize, nil
}

func (o *Transaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction_id",
		"wallet_id",
		"status",
		"source",
		"destination",
		"initiator_type",
		"created_timestamp",
		"updated_timestamp",
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

	varTransaction := _Transaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransaction)

	if err != nil {
		return err
	}

	*o = Transaction(varTransaction)

	return err
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


