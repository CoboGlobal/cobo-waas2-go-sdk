# WebhookEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. | 
**TransactionId** | **string** | The transaction ID. | 
**CoboId** | Pointer to **string** | The Cobo ID, which can be used to track a transaction. | [optional] 
**RequestId** | **string** | The request ID provided by you when creating the settlement request. | 
**WalletId** | **string** | For deposit transactions, this property represents the wallet ID of the transaction destination. For transactions of other types, this property represents the wallet ID of the transaction source. | 
**Type** | Pointer to [**MPCVaultType**](MPCVaultType.md) |  | [optional] 
**Status** | [**SettleRequestStatus**](SettleRequestStatus.md) |  | 
**SubStatus** | Pointer to [**TransactionSubStatus**](TransactionSubStatus.md) |  | [optional] 
**FailedReason** | Pointer to **string** | (This property is applicable to approval failures and signature failures only) The reason why the transaction failed. | [optional] 
**ChainId** | **string** | The ID of the blockchain network on which the refund transaction occurs. | 
**TokenId** | **string** | The ID of the cryptocurrency used for refund. | 
**AssetId** | Pointer to **string** | (This concept applies to Exchange Wallets only) The asset ID. An asset ID is the unique identifier of the asset held within your linked exchange account. | [optional] 
**Source** | [**TokenListingRequestSource**](TokenListingRequestSource.md) |  | 
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
**Description** | Pointer to **string** | The description of the TSS request. | [optional] 
**IsLoop** | Pointer to **bool** | Whether the transaction was executed as a [Cobo Loop](https://manuals.cobo.com/en/portal/custodial-wallets/cobo-loop) transfer. - &#x60;true&#x60;: The transaction was executed as a Cobo Loop transfer. - &#x60;false&#x60;: The transaction was not executed as a Cobo Loop transfer.  | [optional] 
**CoboCategory** | Pointer to **[]string** | The transaction category defined by Cobo. Possible values include:  - &#x60;AutoSweep&#x60;: An auto-sweep transaction. - &#x60;AutoFueling&#x60;: A transaction where Fee Station pays transaction fees to an address within your wallet. - &#x60;AutoFuelingRefund&#x60;: A refund for an auto-fueling transaction. - &#x60;SafeTxMessage&#x60;: A message signing transaction to authorize a Smart Contract Wallet (Safe\\{Wallet\\}) transaction. - &#x60;BillPayment&#x60;: A transaction to pay Cobo bills through Fee Station. - &#x60;BillRefund&#x60;: A refund for a previously made bill payment. - &#x60;CommissionFeeCharge&#x60;: A transaction to charge commission fees via Fee Station. - &#x60;CommissionFeeRefund&#x60;: A refund of previously charged commission fees.  | [optional] 
**Extra** | Pointer to **[]string** | A list of JSON-encoded strings containing structured, business-specific extra information for the transaction. Each item corresponds to a specific data type, indicated by the &#x60;extra_type&#x60; field in the JSON object (for example, \&quot;BabylonBusinessInfo\&quot;, \&quot;BtcAddressInfo\&quot;).  | [optional] 
**FuelingInfo** | Pointer to [**TransactionFuelingInfo**](TransactionFuelingInfo.md) |  | [optional] 
**CreatedTimestamp** | **int64** | Timestamp when the request was created (in milliseconds since Unix epoch) | 
**UpdatedTimestamp** | **int64** | Timestamp when the request was last updated (in milliseconds since Unix epoch) | 
**TssRequestId** | Pointer to **string** | The TSS request ID. | [optional] 
**SourceKeyShareHolderGroup** | Pointer to [**SourceGroup**](SourceGroup.md) |  | [optional] 
**TargetKeyShareHolderGroupId** | Pointer to **string** | The target key share holder group ID. | [optional] 
**Addresses** | Pointer to [**[]AddressesEventDataAllOfAddresses**](AddressesEventDataAllOfAddresses.md) | A list of addresses. | [optional] 
**Wallet** | Pointer to [**WalletInfo**](WalletInfo.md) |  | [optional] 
**VaultId** | Pointer to **string** | The vault ID. | [optional] 
**ProjectId** | Pointer to **string** | The project ID. | [optional] 
**Name** | Pointer to **string** | The vault name. | [optional] 
**RootPubkeys** | Pointer to [**[]RootPubkey**](RootPubkey.md) |  | [optional] 
**Chains** | [**[]ChainInfo**](ChainInfo.md) | The enabled chains. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtypes** | Pointer to [**[]WalletSubtype**](WalletSubtype.md) |  | [optional] 
**Tokens** | [**[]TokenInfo**](TokenInfo.md) | The enabled tokens. | 
**ContractAddress** | **string** | Contract address of the token | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**Token** | Pointer to [**TokenInfo**](TokenInfo.md) |  | [optional] 
**Feedback** | Pointer to **string** | Feedback provided by the admin for rejected requests | [optional] 
**OrderId** | **string** | The order ID. | 
**MerchantId** | Pointer to **string** | The merchant ID. | [optional] 
**PayableAmount** | **string** | The cryptocurrency amount to be paid for this order. | 
**ReceiveAddress** | **string** | The recipient wallet address to be used for the payment transaction. | 
**Currency** | **string** | The fiat currency of the order. | 
**OrderAmount** | **string** | The base amount of the order in fiat currency, excluding the developer fee (specified in &#x60;fee_amount&#x60;). | 
**FeeAmount** | **string** | The developer fee for the order in fiat currency. It is added to the base amount (&#x60;order_amount&#x60;) to determine the final charge. | 
**ExchangeRate** | **string** | The exchange rate between a currency pair. Expressed as the amount of fiat currency per one unit of cryptocurrency. For example, if the cryptocurrency is USDT and the fiat currency is USD, a rate of \&quot;0.99\&quot; means 1 USDT &#x3D; 0.99 USD. | 
**ExpiredAt** | Pointer to **int32** | The expiration time of the pay-in order, represented as a UNIX timestamp in seconds. | [optional] 
**MerchantOrderCode** | Pointer to **string** | A unique reference code assigned by the merchant to identify this order in their system. | [optional] 
**PspOrderCode** | **string** | A unique reference code assigned by the developer to identify this order in their system. | 
**ReceivedTokenAmount** | **string** | The total cryptocurrency amount received for this order. Updates until the expiration time. Precision matches the token standard (e.g., 6 decimals for USDT). | 
**RefundId** | **string** | The refund order ID. | 
**Amount** | **string** | The amount in cryptocurrency to be returned for this refund order. | 
**ToAddress** | **string** | The recipient&#39;s wallet address where the refund will be sent. | 
**Transactions** | Pointer to [**[]PaymentTransaction**](PaymentTransaction.md) | An array of transactions associated with this refund order. Each transaction represents a separate blockchain operation related to the refund process. | [optional] 
**SettlementRequestId** | **string** | The settlement request ID generated by Cobo. | 
**Settlements** | [**[]SettlementDetail**](SettlementDetail.md) |  | 

## Methods

### NewWebhookEventData

`func NewWebhookEventData(dataType string, transactionId string, requestId string, walletId string, status SettleRequestStatus, chainId string, tokenId string, source TokenListingRequestSource, destination TransactionDestination, initiatorType TransactionInitiatorType, createdTimestamp int64, updatedTimestamp int64, chains []ChainInfo, walletType WalletType, tokens []TokenInfo, contractAddress string, walletSubtype WalletSubtype, orderId string, payableAmount string, receiveAddress string, currency string, orderAmount string, feeAmount string, exchangeRate string, pspOrderCode string, receivedTokenAmount string, refundId string, amount string, toAddress string, settlementRequestId string, settlements []SettlementDetail, ) *WebhookEventData`

NewWebhookEventData instantiates a new WebhookEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventDataWithDefaults

`func NewWebhookEventDataWithDefaults() *WebhookEventData`

NewWebhookEventDataWithDefaults instantiates a new WebhookEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *WebhookEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *WebhookEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *WebhookEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetTransactionId

`func (o *WebhookEventData) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *WebhookEventData) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *WebhookEventData) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetCoboId

`func (o *WebhookEventData) GetCoboId() string`

GetCoboId returns the CoboId field if non-nil, zero value otherwise.

### GetCoboIdOk

`func (o *WebhookEventData) GetCoboIdOk() (*string, bool)`

GetCoboIdOk returns a tuple with the CoboId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboId

`func (o *WebhookEventData) SetCoboId(v string)`

SetCoboId sets CoboId field to given value.

### HasCoboId

`func (o *WebhookEventData) HasCoboId() bool`

HasCoboId returns a boolean if a field has been set.

### GetRequestId

`func (o *WebhookEventData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *WebhookEventData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *WebhookEventData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetWalletId

`func (o *WebhookEventData) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *WebhookEventData) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *WebhookEventData) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetType

`func (o *WebhookEventData) GetType() MPCVaultType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookEventData) GetTypeOk() (*MPCVaultType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookEventData) SetType(v MPCVaultType)`

SetType sets Type field to given value.

### HasType

`func (o *WebhookEventData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookEventData) GetStatus() SettleRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookEventData) GetStatusOk() (*SettleRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookEventData) SetStatus(v SettleRequestStatus)`

SetStatus sets Status field to given value.


### GetSubStatus

`func (o *WebhookEventData) GetSubStatus() TransactionSubStatus`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *WebhookEventData) GetSubStatusOk() (*TransactionSubStatus, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *WebhookEventData) SetSubStatus(v TransactionSubStatus)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *WebhookEventData) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetFailedReason

`func (o *WebhookEventData) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *WebhookEventData) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *WebhookEventData) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *WebhookEventData) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.

### GetChainId

`func (o *WebhookEventData) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *WebhookEventData) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *WebhookEventData) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetTokenId

`func (o *WebhookEventData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *WebhookEventData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *WebhookEventData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAssetId

`func (o *WebhookEventData) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *WebhookEventData) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *WebhookEventData) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *WebhookEventData) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetSource

`func (o *WebhookEventData) GetSource() TokenListingRequestSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WebhookEventData) GetSourceOk() (*TokenListingRequestSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WebhookEventData) SetSource(v TokenListingRequestSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *WebhookEventData) GetDestination() TransactionDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *WebhookEventData) GetDestinationOk() (*TransactionDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *WebhookEventData) SetDestination(v TransactionDestination)`

SetDestination sets Destination field to given value.


### GetResult

`func (o *WebhookEventData) GetResult() TransactionResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WebhookEventData) GetResultOk() (*TransactionResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WebhookEventData) SetResult(v TransactionResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *WebhookEventData) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetFee

`func (o *WebhookEventData) GetFee() TransactionFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *WebhookEventData) GetFeeOk() (*TransactionFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *WebhookEventData) SetFee(v TransactionFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *WebhookEventData) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetInitiator

`func (o *WebhookEventData) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *WebhookEventData) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *WebhookEventData) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *WebhookEventData) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetInitiatorType

`func (o *WebhookEventData) GetInitiatorType() TransactionInitiatorType`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *WebhookEventData) GetInitiatorTypeOk() (*TransactionInitiatorType, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *WebhookEventData) SetInitiatorType(v TransactionInitiatorType)`

SetInitiatorType sets InitiatorType field to given value.


### GetConfirmedNum

`func (o *WebhookEventData) GetConfirmedNum() int32`

GetConfirmedNum returns the ConfirmedNum field if non-nil, zero value otherwise.

### GetConfirmedNumOk

`func (o *WebhookEventData) GetConfirmedNumOk() (*int32, bool)`

GetConfirmedNumOk returns a tuple with the ConfirmedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedNum

`func (o *WebhookEventData) SetConfirmedNum(v int32)`

SetConfirmedNum sets ConfirmedNum field to given value.

### HasConfirmedNum

`func (o *WebhookEventData) HasConfirmedNum() bool`

HasConfirmedNum returns a boolean if a field has been set.

### GetConfirmingThreshold

`func (o *WebhookEventData) GetConfirmingThreshold() int32`

GetConfirmingThreshold returns the ConfirmingThreshold field if non-nil, zero value otherwise.

### GetConfirmingThresholdOk

`func (o *WebhookEventData) GetConfirmingThresholdOk() (*int32, bool)`

GetConfirmingThresholdOk returns a tuple with the ConfirmingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmingThreshold

`func (o *WebhookEventData) SetConfirmingThreshold(v int32)`

SetConfirmingThreshold sets ConfirmingThreshold field to given value.

### HasConfirmingThreshold

`func (o *WebhookEventData) HasConfirmingThreshold() bool`

HasConfirmingThreshold returns a boolean if a field has been set.

### GetTransactionHash

`func (o *WebhookEventData) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *WebhookEventData) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *WebhookEventData) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *WebhookEventData) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetBlockInfo

`func (o *WebhookEventData) GetBlockInfo() TransactionBlockInfo`

GetBlockInfo returns the BlockInfo field if non-nil, zero value otherwise.

### GetBlockInfoOk

`func (o *WebhookEventData) GetBlockInfoOk() (*TransactionBlockInfo, bool)`

GetBlockInfoOk returns a tuple with the BlockInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockInfo

`func (o *WebhookEventData) SetBlockInfo(v TransactionBlockInfo)`

SetBlockInfo sets BlockInfo field to given value.

### HasBlockInfo

`func (o *WebhookEventData) HasBlockInfo() bool`

HasBlockInfo returns a boolean if a field has been set.

### GetRawTxInfo

`func (o *WebhookEventData) GetRawTxInfo() TransactionRawTxInfo`

GetRawTxInfo returns the RawTxInfo field if non-nil, zero value otherwise.

### GetRawTxInfoOk

`func (o *WebhookEventData) GetRawTxInfoOk() (*TransactionRawTxInfo, bool)`

GetRawTxInfoOk returns a tuple with the RawTxInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTxInfo

`func (o *WebhookEventData) SetRawTxInfo(v TransactionRawTxInfo)`

SetRawTxInfo sets RawTxInfo field to given value.

### HasRawTxInfo

`func (o *WebhookEventData) HasRawTxInfo() bool`

HasRawTxInfo returns a boolean if a field has been set.

### GetReplacement

`func (o *WebhookEventData) GetReplacement() TransactionReplacement`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *WebhookEventData) GetReplacementOk() (*TransactionReplacement, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *WebhookEventData) SetReplacement(v TransactionReplacement)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *WebhookEventData) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetCategory

`func (o *WebhookEventData) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *WebhookEventData) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *WebhookEventData) SetCategory(v []string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *WebhookEventData) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *WebhookEventData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookEventData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookEventData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookEventData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsLoop

`func (o *WebhookEventData) GetIsLoop() bool`

GetIsLoop returns the IsLoop field if non-nil, zero value otherwise.

### GetIsLoopOk

`func (o *WebhookEventData) GetIsLoopOk() (*bool, bool)`

GetIsLoopOk returns a tuple with the IsLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoop

`func (o *WebhookEventData) SetIsLoop(v bool)`

SetIsLoop sets IsLoop field to given value.

### HasIsLoop

`func (o *WebhookEventData) HasIsLoop() bool`

HasIsLoop returns a boolean if a field has been set.

### GetCoboCategory

`func (o *WebhookEventData) GetCoboCategory() []string`

GetCoboCategory returns the CoboCategory field if non-nil, zero value otherwise.

### GetCoboCategoryOk

`func (o *WebhookEventData) GetCoboCategoryOk() (*[]string, bool)`

GetCoboCategoryOk returns a tuple with the CoboCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoboCategory

`func (o *WebhookEventData) SetCoboCategory(v []string)`

SetCoboCategory sets CoboCategory field to given value.

### HasCoboCategory

`func (o *WebhookEventData) HasCoboCategory() bool`

HasCoboCategory returns a boolean if a field has been set.

### GetExtra

`func (o *WebhookEventData) GetExtra() []string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *WebhookEventData) GetExtraOk() (*[]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *WebhookEventData) SetExtra(v []string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *WebhookEventData) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetFuelingInfo

`func (o *WebhookEventData) GetFuelingInfo() TransactionFuelingInfo`

GetFuelingInfo returns the FuelingInfo field if non-nil, zero value otherwise.

### GetFuelingInfoOk

`func (o *WebhookEventData) GetFuelingInfoOk() (*TransactionFuelingInfo, bool)`

GetFuelingInfoOk returns a tuple with the FuelingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelingInfo

`func (o *WebhookEventData) SetFuelingInfo(v TransactionFuelingInfo)`

SetFuelingInfo sets FuelingInfo field to given value.

### HasFuelingInfo

`func (o *WebhookEventData) HasFuelingInfo() bool`

HasFuelingInfo returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *WebhookEventData) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *WebhookEventData) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *WebhookEventData) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *WebhookEventData) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *WebhookEventData) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *WebhookEventData) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetTssRequestId

`func (o *WebhookEventData) GetTssRequestId() string`

GetTssRequestId returns the TssRequestId field if non-nil, zero value otherwise.

### GetTssRequestIdOk

`func (o *WebhookEventData) GetTssRequestIdOk() (*string, bool)`

GetTssRequestIdOk returns a tuple with the TssRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTssRequestId

`func (o *WebhookEventData) SetTssRequestId(v string)`

SetTssRequestId sets TssRequestId field to given value.

### HasTssRequestId

`func (o *WebhookEventData) HasTssRequestId() bool`

HasTssRequestId returns a boolean if a field has been set.

### GetSourceKeyShareHolderGroup

`func (o *WebhookEventData) GetSourceKeyShareHolderGroup() SourceGroup`

GetSourceKeyShareHolderGroup returns the SourceKeyShareHolderGroup field if non-nil, zero value otherwise.

### GetSourceKeyShareHolderGroupOk

`func (o *WebhookEventData) GetSourceKeyShareHolderGroupOk() (*SourceGroup, bool)`

GetSourceKeyShareHolderGroupOk returns a tuple with the SourceKeyShareHolderGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKeyShareHolderGroup

`func (o *WebhookEventData) SetSourceKeyShareHolderGroup(v SourceGroup)`

SetSourceKeyShareHolderGroup sets SourceKeyShareHolderGroup field to given value.

### HasSourceKeyShareHolderGroup

`func (o *WebhookEventData) HasSourceKeyShareHolderGroup() bool`

HasSourceKeyShareHolderGroup returns a boolean if a field has been set.

### GetTargetKeyShareHolderGroupId

`func (o *WebhookEventData) GetTargetKeyShareHolderGroupId() string`

GetTargetKeyShareHolderGroupId returns the TargetKeyShareHolderGroupId field if non-nil, zero value otherwise.

### GetTargetKeyShareHolderGroupIdOk

`func (o *WebhookEventData) GetTargetKeyShareHolderGroupIdOk() (*string, bool)`

GetTargetKeyShareHolderGroupIdOk returns a tuple with the TargetKeyShareHolderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKeyShareHolderGroupId

`func (o *WebhookEventData) SetTargetKeyShareHolderGroupId(v string)`

SetTargetKeyShareHolderGroupId sets TargetKeyShareHolderGroupId field to given value.

### HasTargetKeyShareHolderGroupId

`func (o *WebhookEventData) HasTargetKeyShareHolderGroupId() bool`

HasTargetKeyShareHolderGroupId returns a boolean if a field has been set.

### GetAddresses

`func (o *WebhookEventData) GetAddresses() []AddressesEventDataAllOfAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *WebhookEventData) GetAddressesOk() (*[]AddressesEventDataAllOfAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *WebhookEventData) SetAddresses(v []AddressesEventDataAllOfAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *WebhookEventData) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetWallet

`func (o *WebhookEventData) GetWallet() WalletInfo`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *WebhookEventData) GetWalletOk() (*WalletInfo, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *WebhookEventData) SetWallet(v WalletInfo)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *WebhookEventData) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetVaultId

`func (o *WebhookEventData) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *WebhookEventData) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *WebhookEventData) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *WebhookEventData) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### GetProjectId

`func (o *WebhookEventData) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WebhookEventData) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WebhookEventData) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *WebhookEventData) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetName

`func (o *WebhookEventData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookEventData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookEventData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookEventData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRootPubkeys

`func (o *WebhookEventData) GetRootPubkeys() []RootPubkey`

GetRootPubkeys returns the RootPubkeys field if non-nil, zero value otherwise.

### GetRootPubkeysOk

`func (o *WebhookEventData) GetRootPubkeysOk() (*[]RootPubkey, bool)`

GetRootPubkeysOk returns a tuple with the RootPubkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPubkeys

`func (o *WebhookEventData) SetRootPubkeys(v []RootPubkey)`

SetRootPubkeys sets RootPubkeys field to given value.

### HasRootPubkeys

`func (o *WebhookEventData) HasRootPubkeys() bool`

HasRootPubkeys returns a boolean if a field has been set.

### GetChains

`func (o *WebhookEventData) GetChains() []ChainInfo`

GetChains returns the Chains field if non-nil, zero value otherwise.

### GetChainsOk

`func (o *WebhookEventData) GetChainsOk() (*[]ChainInfo, bool)`

GetChainsOk returns a tuple with the Chains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChains

`func (o *WebhookEventData) SetChains(v []ChainInfo)`

SetChains sets Chains field to given value.


### GetWalletType

`func (o *WebhookEventData) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *WebhookEventData) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *WebhookEventData) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtypes

`func (o *WebhookEventData) GetWalletSubtypes() []WalletSubtype`

GetWalletSubtypes returns the WalletSubtypes field if non-nil, zero value otherwise.

### GetWalletSubtypesOk

`func (o *WebhookEventData) GetWalletSubtypesOk() (*[]WalletSubtype, bool)`

GetWalletSubtypesOk returns a tuple with the WalletSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtypes

`func (o *WebhookEventData) SetWalletSubtypes(v []WalletSubtype)`

SetWalletSubtypes sets WalletSubtypes field to given value.

### HasWalletSubtypes

`func (o *WebhookEventData) HasWalletSubtypes() bool`

HasWalletSubtypes returns a boolean if a field has been set.

### GetTokens

`func (o *WebhookEventData) GetTokens() []TokenInfo`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *WebhookEventData) GetTokensOk() (*[]TokenInfo, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *WebhookEventData) SetTokens(v []TokenInfo)`

SetTokens sets Tokens field to given value.


### GetContractAddress

`func (o *WebhookEventData) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *WebhookEventData) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *WebhookEventData) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetWalletSubtype

`func (o *WebhookEventData) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *WebhookEventData) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *WebhookEventData) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetToken

`func (o *WebhookEventData) GetToken() TokenInfo`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *WebhookEventData) GetTokenOk() (*TokenInfo, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *WebhookEventData) SetToken(v TokenInfo)`

SetToken sets Token field to given value.

### HasToken

`func (o *WebhookEventData) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetFeedback

`func (o *WebhookEventData) GetFeedback() string`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *WebhookEventData) GetFeedbackOk() (*string, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *WebhookEventData) SetFeedback(v string)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *WebhookEventData) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetOrderId

`func (o *WebhookEventData) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *WebhookEventData) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *WebhookEventData) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetMerchantId

`func (o *WebhookEventData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *WebhookEventData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *WebhookEventData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *WebhookEventData) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPayableAmount

`func (o *WebhookEventData) GetPayableAmount() string`

GetPayableAmount returns the PayableAmount field if non-nil, zero value otherwise.

### GetPayableAmountOk

`func (o *WebhookEventData) GetPayableAmountOk() (*string, bool)`

GetPayableAmountOk returns a tuple with the PayableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableAmount

`func (o *WebhookEventData) SetPayableAmount(v string)`

SetPayableAmount sets PayableAmount field to given value.


### GetReceiveAddress

`func (o *WebhookEventData) GetReceiveAddress() string`

GetReceiveAddress returns the ReceiveAddress field if non-nil, zero value otherwise.

### GetReceiveAddressOk

`func (o *WebhookEventData) GetReceiveAddressOk() (*string, bool)`

GetReceiveAddressOk returns a tuple with the ReceiveAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveAddress

`func (o *WebhookEventData) SetReceiveAddress(v string)`

SetReceiveAddress sets ReceiveAddress field to given value.


### GetCurrency

`func (o *WebhookEventData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WebhookEventData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WebhookEventData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetOrderAmount

`func (o *WebhookEventData) GetOrderAmount() string`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *WebhookEventData) GetOrderAmountOk() (*string, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *WebhookEventData) SetOrderAmount(v string)`

SetOrderAmount sets OrderAmount field to given value.


### GetFeeAmount

`func (o *WebhookEventData) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *WebhookEventData) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *WebhookEventData) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetExchangeRate

`func (o *WebhookEventData) GetExchangeRate() string`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *WebhookEventData) GetExchangeRateOk() (*string, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *WebhookEventData) SetExchangeRate(v string)`

SetExchangeRate sets ExchangeRate field to given value.


### GetExpiredAt

`func (o *WebhookEventData) GetExpiredAt() int32`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *WebhookEventData) GetExpiredAtOk() (*int32, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *WebhookEventData) SetExpiredAt(v int32)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *WebhookEventData) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetMerchantOrderCode

`func (o *WebhookEventData) GetMerchantOrderCode() string`

GetMerchantOrderCode returns the MerchantOrderCode field if non-nil, zero value otherwise.

### GetMerchantOrderCodeOk

`func (o *WebhookEventData) GetMerchantOrderCodeOk() (*string, bool)`

GetMerchantOrderCodeOk returns a tuple with the MerchantOrderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderCode

`func (o *WebhookEventData) SetMerchantOrderCode(v string)`

SetMerchantOrderCode sets MerchantOrderCode field to given value.

### HasMerchantOrderCode

`func (o *WebhookEventData) HasMerchantOrderCode() bool`

HasMerchantOrderCode returns a boolean if a field has been set.

### GetPspOrderCode

`func (o *WebhookEventData) GetPspOrderCode() string`

GetPspOrderCode returns the PspOrderCode field if non-nil, zero value otherwise.

### GetPspOrderCodeOk

`func (o *WebhookEventData) GetPspOrderCodeOk() (*string, bool)`

GetPspOrderCodeOk returns a tuple with the PspOrderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspOrderCode

`func (o *WebhookEventData) SetPspOrderCode(v string)`

SetPspOrderCode sets PspOrderCode field to given value.


### GetReceivedTokenAmount

`func (o *WebhookEventData) GetReceivedTokenAmount() string`

GetReceivedTokenAmount returns the ReceivedTokenAmount field if non-nil, zero value otherwise.

### GetReceivedTokenAmountOk

`func (o *WebhookEventData) GetReceivedTokenAmountOk() (*string, bool)`

GetReceivedTokenAmountOk returns a tuple with the ReceivedTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTokenAmount

`func (o *WebhookEventData) SetReceivedTokenAmount(v string)`

SetReceivedTokenAmount sets ReceivedTokenAmount field to given value.


### GetRefundId

`func (o *WebhookEventData) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *WebhookEventData) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *WebhookEventData) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.


### GetAmount

`func (o *WebhookEventData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WebhookEventData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WebhookEventData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetToAddress

`func (o *WebhookEventData) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *WebhookEventData) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *WebhookEventData) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetTransactions

`func (o *WebhookEventData) GetTransactions() []PaymentTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *WebhookEventData) GetTransactionsOk() (*[]PaymentTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *WebhookEventData) SetTransactions(v []PaymentTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *WebhookEventData) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetSettlementRequestId

`func (o *WebhookEventData) GetSettlementRequestId() string`

GetSettlementRequestId returns the SettlementRequestId field if non-nil, zero value otherwise.

### GetSettlementRequestIdOk

`func (o *WebhookEventData) GetSettlementRequestIdOk() (*string, bool)`

GetSettlementRequestIdOk returns a tuple with the SettlementRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementRequestId

`func (o *WebhookEventData) SetSettlementRequestId(v string)`

SetSettlementRequestId sets SettlementRequestId field to given value.


### GetSettlements

`func (o *WebhookEventData) GetSettlements() []SettlementDetail`

GetSettlements returns the Settlements field if non-nil, zero value otherwise.

### GetSettlementsOk

`func (o *WebhookEventData) GetSettlementsOk() (*[]SettlementDetail, bool)`

GetSettlementsOk returns a tuple with the Settlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlements

`func (o *WebhookEventData) SetSettlements(v []SettlementDetail)`

SetSettlements sets Settlements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


