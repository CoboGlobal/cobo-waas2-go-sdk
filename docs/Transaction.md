# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The transaction ID. | 
**CoboId** | Pointer to **string** | The Cobo ID, which can be used to track a transaction. | [optional] 
**RequestId** | Pointer to **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | [optional] 
**WalletId** | **string** | For deposit transactions, this property represents the wallet ID of the transaction destination. For transactions of other types, this property represents the wallet ID of the transaction source. | 
**Type** | Pointer to [**TransactionType**](TransactionType.md) |  | [optional] 
**Status** | [**TransactionStatus**](TransactionStatus.md) |  | 
**SubStatus** | Pointer to [**TransactionSubStatus**](TransactionSubStatus.md) |  | [optional] 
**FailedReason** | Pointer to **string** | (This property is applicable to approval failures and signature failures only) The reason why the transaction failed. | [optional] 
**ChainId** | Pointer to **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | [optional] 
**TokenId** | Pointer to **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | [optional] 
**AssetId** | Pointer to **string** | (This concept applies to Exchange Wallets only) The asset ID. An asset ID is the unique identifier of the asset held within your linked exchange account. | [optional] 
**Source** | [**TransactionSource**](TransactionSource.md) |  | 
**Destination** | [**TransactionDestination**](TransactionDestination.md) |  | 
**Result** | Pointer to [**TransactionResult**](TransactionResult.md) |  | [optional] 
**Fee** | Pointer to [**TransactionFee**](TransactionFee.md) |  | [optional] 
**Initiator** | Pointer to **string** | The transaction initiator. | [optional] 
**InitiatorType** | [**TransactionInitiatorType**](TransactionInitiatorType.md) |  | 
**ConfirmedNum** | Pointer to **int32** | The number of confirmations this transaction has received. | [optional] 
**ConfirmingThreshold** | Pointer to **int32** | The minimum number of confirmations required to deem a transaction secure. The common threshold is 6 for a Bitcoin transaction. | [optional] 
**TransactionHash** | Pointer to **string** | The transaction hash. | [optional] 
**BlockInfo** | Pointer to [**TransactionBlockInfo**](TransactionBlockInfo.md) |  | [optional] 
**RawTxInfo** | Pointer to [**TransactionRawTxInfo**](TransactionRawTxInfo.md) |  | [optional] 
**Replacement** | Pointer to [**TransactionReplacement**](TransactionReplacement.md) |  | [optional] 
**Category** | Pointer to **[]string** | A custom transaction category for you to identify your transfers more easily. | [optional] 
**Description** | Pointer to **string** | The description for your transaction. | [optional] 
**IsLoop** | Pointer to **bool** | Whether the transaction was executed as a [Cobo Loop](https://manuals.cobo.com/en/portal/custodial-wallets/cobo-loop) transfer. - &#x60;true&#x60;: The transaction was executed as a Cobo Loop transfer. - &#x60;false&#x60;: The transaction was not executed as a Cobo Loop transfer.  | [optional] 
**CoboCategory** | Pointer to **[]string** | The transaction category defined by Cobo. Possible values include:  - &#x60;AutoSweep&#x60;: An auto-sweep transaction. - &#x60;AutoFueling&#x60;: A transaction where Fee Station pays transaction fees to an address within your wallet. - &#x60;AutoFuelingRefund&#x60;: A refund for an auto-fueling transaction. - &#x60;SafeTxMessage&#x60;: A message signing transaction to authorize a Smart Contract Wallet (Safe\\{Wallet\\}) transaction. - &#x60;BillPayment&#x60;: A transaction to pay Cobo bills through Fee Station. - &#x60;BillRefund&#x60;: A refund for a previously made bill payment. - &#x60;CommissionFeeCharge&#x60;: A transaction to charge commission fees via Fee Station. - &#x60;CommissionFeeRefund&#x60;: A refund of previously charged commission fees.  | [optional] 
**Extra** | Pointer to **[]string** | A list of JSON-encoded strings containing structured, business-specific extra information for the transaction. Each item corresponds to a specific data type, indicated by the &#x60;extra_type&#x60; field in the JSON object (for example, \&quot;BabylonBusinessInfo\&quot;, \&quot;BtcAddressInfo\&quot;).  | [optional] 
**FuelingInfo** | Pointer to [**TransactionFuelingInfo**](TransactionFuelingInfo.md) |  | [optional] 
**CreatedTimestamp** | **int64** | The time when the transaction was created, in Unix timestamp format, measured in milliseconds. | 
**UpdatedTimestamp** | **int64** | The time when the transaction was updated, in Unix timestamp format, measured in milliseconds. | 

## Methods

### NewTransaction

`func NewTransaction(transactionId string, walletId string, status TransactionStatus, source TransactionSource, destination TransactionDestination, initiatorType TransactionInitiatorType, createdTimestamp int64, updatedTimestamp int64, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *Transaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Transaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Transaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetCoboId

`func (o *Transaction) GetCoboId() string`

GetCoboId returns the CoboId field if non-nil, zero value otherwise.

### GetCoboIdOk

`func (o *Transaction) GetCoboIdOk() (*string, bool)`

GetCoboIdOk returns a tuple with the CoboId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboId

`func (o *Transaction) SetCoboId(v string)`

SetCoboId sets CoboId field to given value.

### HasCoboId

`func (o *Transaction) HasCoboId() bool`

HasCoboId returns a boolean if a field has been set.

### GetRequestId

`func (o *Transaction) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Transaction) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Transaction) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Transaction) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetWalletId

`func (o *Transaction) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *Transaction) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *Transaction) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetType

`func (o *Transaction) GetType() TransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*TransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v TransactionType)`

SetType sets Type field to given value.

### HasType

`func (o *Transaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *Transaction) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.


### GetSubStatus

`func (o *Transaction) GetSubStatus() TransactionSubStatus`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *Transaction) GetSubStatusOk() (*TransactionSubStatus, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *Transaction) SetSubStatus(v TransactionSubStatus)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *Transaction) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetFailedReason

`func (o *Transaction) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *Transaction) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *Transaction) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *Transaction) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.

### GetChainId

`func (o *Transaction) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Transaction) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Transaction) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *Transaction) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetTokenId

`func (o *Transaction) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Transaction) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Transaction) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *Transaction) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetAssetId

`func (o *Transaction) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Transaction) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Transaction) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *Transaction) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetSource

`func (o *Transaction) GetSource() TransactionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Transaction) GetSourceOk() (*TransactionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Transaction) SetSource(v TransactionSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *Transaction) GetDestination() TransactionDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Transaction) GetDestinationOk() (*TransactionDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Transaction) SetDestination(v TransactionDestination)`

SetDestination sets Destination field to given value.


### GetResult

`func (o *Transaction) GetResult() TransactionResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Transaction) GetResultOk() (*TransactionResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Transaction) SetResult(v TransactionResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *Transaction) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetFee

`func (o *Transaction) GetFee() TransactionFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *Transaction) GetFeeOk() (*TransactionFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *Transaction) SetFee(v TransactionFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *Transaction) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetInitiator

`func (o *Transaction) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *Transaction) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *Transaction) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *Transaction) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetInitiatorType

`func (o *Transaction) GetInitiatorType() TransactionInitiatorType`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *Transaction) GetInitiatorTypeOk() (*TransactionInitiatorType, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *Transaction) SetInitiatorType(v TransactionInitiatorType)`

SetInitiatorType sets InitiatorType field to given value.


### GetConfirmedNum

`func (o *Transaction) GetConfirmedNum() int32`

GetConfirmedNum returns the ConfirmedNum field if non-nil, zero value otherwise.

### GetConfirmedNumOk

`func (o *Transaction) GetConfirmedNumOk() (*int32, bool)`

GetConfirmedNumOk returns a tuple with the ConfirmedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedNum

`func (o *Transaction) SetConfirmedNum(v int32)`

SetConfirmedNum sets ConfirmedNum field to given value.

### HasConfirmedNum

`func (o *Transaction) HasConfirmedNum() bool`

HasConfirmedNum returns a boolean if a field has been set.

### GetConfirmingThreshold

`func (o *Transaction) GetConfirmingThreshold() int32`

GetConfirmingThreshold returns the ConfirmingThreshold field if non-nil, zero value otherwise.

### GetConfirmingThresholdOk

`func (o *Transaction) GetConfirmingThresholdOk() (*int32, bool)`

GetConfirmingThresholdOk returns a tuple with the ConfirmingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmingThreshold

`func (o *Transaction) SetConfirmingThreshold(v int32)`

SetConfirmingThreshold sets ConfirmingThreshold field to given value.

### HasConfirmingThreshold

`func (o *Transaction) HasConfirmingThreshold() bool`

HasConfirmingThreshold returns a boolean if a field has been set.

### GetTransactionHash

`func (o *Transaction) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *Transaction) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *Transaction) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *Transaction) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetBlockInfo

`func (o *Transaction) GetBlockInfo() TransactionBlockInfo`

GetBlockInfo returns the BlockInfo field if non-nil, zero value otherwise.

### GetBlockInfoOk

`func (o *Transaction) GetBlockInfoOk() (*TransactionBlockInfo, bool)`

GetBlockInfoOk returns a tuple with the BlockInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockInfo

`func (o *Transaction) SetBlockInfo(v TransactionBlockInfo)`

SetBlockInfo sets BlockInfo field to given value.

### HasBlockInfo

`func (o *Transaction) HasBlockInfo() bool`

HasBlockInfo returns a boolean if a field has been set.

### GetRawTxInfo

`func (o *Transaction) GetRawTxInfo() TransactionRawTxInfo`

GetRawTxInfo returns the RawTxInfo field if non-nil, zero value otherwise.

### GetRawTxInfoOk

`func (o *Transaction) GetRawTxInfoOk() (*TransactionRawTxInfo, bool)`

GetRawTxInfoOk returns a tuple with the RawTxInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTxInfo

`func (o *Transaction) SetRawTxInfo(v TransactionRawTxInfo)`

SetRawTxInfo sets RawTxInfo field to given value.

### HasRawTxInfo

`func (o *Transaction) HasRawTxInfo() bool`

HasRawTxInfo returns a boolean if a field has been set.

### GetReplacement

`func (o *Transaction) GetReplacement() TransactionReplacement`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *Transaction) GetReplacementOk() (*TransactionReplacement, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *Transaction) SetReplacement(v TransactionReplacement)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *Transaction) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetCategory

`func (o *Transaction) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Transaction) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Transaction) SetCategory(v []string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Transaction) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *Transaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Transaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsLoop

`func (o *Transaction) GetIsLoop() bool`

GetIsLoop returns the IsLoop field if non-nil, zero value otherwise.

### GetIsLoopOk

`func (o *Transaction) GetIsLoopOk() (*bool, bool)`

GetIsLoopOk returns a tuple with the IsLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoop

`func (o *Transaction) SetIsLoop(v bool)`

SetIsLoop sets IsLoop field to given value.

### HasIsLoop

`func (o *Transaction) HasIsLoop() bool`

HasIsLoop returns a boolean if a field has been set.

### GetCoboCategory

`func (o *Transaction) GetCoboCategory() []string`

GetCoboCategory returns the CoboCategory field if non-nil, zero value otherwise.

### GetCoboCategoryOk

`func (o *Transaction) GetCoboCategoryOk() (*[]string, bool)`

GetCoboCategoryOk returns a tuple with the CoboCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboCategory

`func (o *Transaction) SetCoboCategory(v []string)`

SetCoboCategory sets CoboCategory field to given value.

### HasCoboCategory

`func (o *Transaction) HasCoboCategory() bool`

HasCoboCategory returns a boolean if a field has been set.

### GetExtra

`func (o *Transaction) GetExtra() []string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *Transaction) GetExtraOk() (*[]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *Transaction) SetExtra(v []string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *Transaction) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetFuelingInfo

`func (o *Transaction) GetFuelingInfo() TransactionFuelingInfo`

GetFuelingInfo returns the FuelingInfo field if non-nil, zero value otherwise.

### GetFuelingInfoOk

`func (o *Transaction) GetFuelingInfoOk() (*TransactionFuelingInfo, bool)`

GetFuelingInfoOk returns a tuple with the FuelingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelingInfo

`func (o *Transaction) SetFuelingInfo(v TransactionFuelingInfo)`

SetFuelingInfo sets FuelingInfo field to given value.

### HasFuelingInfo

`func (o *Transaction) HasFuelingInfo() bool`

HasFuelingInfo returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *Transaction) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Transaction) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Transaction) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *Transaction) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Transaction) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Transaction) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


