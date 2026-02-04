# WebhookEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. | 
**TransactionId** | **string** | The transaction ID. | 
**CoboId** | Pointer to **string** | The Cobo ID, which can be used to track a transaction. | [optional] 
**RequestId** | **string** | The request ID provided by you when creating the payout. | 
**WalletId** | **string** | For deposit transactions, this property represents the wallet ID of the transaction destination. For transactions of other types, this property represents the wallet ID of the transaction source. | 
**Type** | Pointer to [**TransactionType**](TransactionType.md) |  | [optional] 
**Status** | [**KyaScreeningStatus**](KyaScreeningStatus.md) |  | 
**SubStatus** | Pointer to [**TransactionSubStatus**](TransactionSubStatus.md) |  | [optional] 
**FailedReason** | Pointer to **string** | (This property is applicable to approval failures and signature failures only) The reason why the transaction failed. | [optional] 
**ChainId** | **string** | The chain identifier. | 
**TokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
**AssetId** | Pointer to **string** | (This concept applies to Exchange Wallets only) The asset ID. An asset ID is the unique identifier of the asset held within your linked exchange account. | [optional] 
**Source** | [**TransactionSource**](TransactionSource.md) |  | 
**Destination** | [**TransactionDestination**](TransactionDestination.md) |  | 
**Result** | Pointer to [**TransactionResult**](TransactionResult.md) |  | [optional] 
**Fee** | Pointer to [**TransactionFee**](TransactionFee.md) |  | [optional] 
**Initiator** | Pointer to **string** | The initiator of this payout, usually the user&#39;s API key. | [optional] 
**InitiatorType** | [**TransactionInitiatorType**](TransactionInitiatorType.md) |  | 
**ConfirmedNum** | Pointer to **int32** | The number of confirmations this transaction has received. | [optional] 
**ConfirmingThreshold** | Pointer to **int32** | The minimum number of confirmations required to deem a transaction secure. The common threshold is 6 for a Bitcoin transaction. | [optional] 
**TransactionHash** | Pointer to **string** | The transaction hash. | [optional] 
**BlockInfo** | Pointer to [**TransactionBlockInfo**](TransactionBlockInfo.md) |  | [optional] 
**RawTxInfo** | Pointer to [**TransactionRawTxInfo**](TransactionRawTxInfo.md) |  | [optional] 
**Replacement** | Pointer to [**TransactionReplacement**](TransactionReplacement.md) |  | [optional] 
**Category** | Pointer to **[]string** | A custom transaction category for you to identify your transfers more easily. | [optional] 
**Description** | Pointer to **string** | The description for the entire bulk send batch. | [optional] 
**IsLoop** | Pointer to **bool** | Whether the transaction was executed as a [Cobo Loop](https://manuals.cobo.com/en/portal/custodial-wallets/cobo-loop) transfer. - &#x60;true&#x60;: The transaction was executed as a Cobo Loop transfer. - &#x60;false&#x60;: The transaction was not executed as a Cobo Loop transfer.  | [optional] 
**CoboCategory** | Pointer to **[]string** | The transaction category defined by Cobo. For more details, refer to [Cobo-defined categories](/v2/guides/transactions/manage-transactions#cobo-defined-categories).  | [optional] 
**Extra** | Pointer to **[]string** | A list of JSON-encoded strings containing structured, business-specific extra information for the transaction. Each item corresponds to a specific data type, indicated by the &#x60;extra_type&#x60; field in the JSON object (for example, \&quot;BabylonBusinessInfo\&quot;, \&quot;BtcAddressInfo\&quot;).  | [optional] 
**FuelingInfo** | Pointer to [**TransactionFuelingInfo**](TransactionFuelingInfo.md) |  | [optional] 
**CreatedTimestamp** | **int64** | The time when the screening request was created, in Unix timestamp format, measured in milliseconds. | 
**UpdatedTimestamp** | **int64** | The time when the screening status was updated, in Unix timestamp format, measured in milliseconds. | 
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
**ContractAddress** | **string** | The token&#39;s contract address on the specified blockchain. | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**Token** | Pointer to [**TokenInfo**](TokenInfo.md) |  | [optional] 
**Feedback** | Pointer to **string** | The feedback provided by Cobo when a token listing request is rejected. | [optional] 
**Address** | **string** | The screened blockchain address. | 
**WalletUuid** | **string** | The wallet ID. | 
**Balance** | [**Balance**](Balance.md) |  | 
**TokenIds** | **string** | A list of token IDs, separated by comma. | 
**OperationType** | [**SuspendedTokenOperationType**](SuspendedTokenOperationType.md) |  | 
**OrderId** | **string** | The pay-in order ID. | 
**MerchantId** | Pointer to **string** | The merchant ID. | [optional] 
**MerchantOrderCode** | Pointer to **string** | A unique reference code assigned by the merchant to identify this order in their system. | [optional] 
**PspOrderCode** | **string** | A unique reference code assigned by the developer to identify this order in their system. | 
**PricingCurrency** | Pointer to **string** | The pricing currency of the order. | [optional] 
**PricingAmount** | Pointer to **string** | The base amount of the order, excluding the developer fee (specified in &#x60;fee_amount&#x60;). | [optional] 
**FeeAmount** | **string** | The developer fee for the order. It is added to the base amount to determine the final charge. | 
**PayableCurrency** | Pointer to **string** | The ID of the cryptocurrency used for payment. | [optional] 
**PayableAmount** | **string** | The cryptocurrency amount to be paid for this order. | 
**ExchangeRate** | **string** | The exchange rate between &#x60;payable_currency&#x60; and &#x60;pricing_currency&#x60;, calculated as (&#x60;pricing_amount&#x60; + &#x60;fee_amount&#x60;) / &#x60;payable_amount&#x60;.    &lt;Note&gt;This field is only returned when &#x60;payable_amount&#x60; was not provided in the order creation request. &lt;/Note&gt;  | 
**AmountTolerance** | Pointer to **string** | The allowed amount deviation, with precision up to 1 decimal place.  For example, if &#x60;payable_amount&#x60; is &#x60;100.00&#x60; and &#x60;amount_tolerance&#x60; is &#x60;0.50&#x60;: - Payer pays 99.55 → Success (difference of 0.45 ≤ 0.5) - Payer pays 99.40 → Underpaid (difference of 0.60 &gt; 0.5)  | [optional] 
**ReceiveAddress** | **string** | The recipient wallet address to be used for the payment transaction. | 
**ReceivedTokenAmount** | **string** | The total cryptocurrency amount received for this order. Updates until the expiration time. Precision matches the token standard (e.g., 6 decimals for USDT). | 
**ExpiredAt** | Pointer to **int32** | The expiration time of the pay-in order, represented as a UNIX timestamp in seconds. | [optional] 
**Transactions** | Pointer to [**[]PaymentTransaction**](PaymentTransaction.md) | An array of payout transactions. | [optional] 
**Currency** | Pointer to **string** | The fiat currency for the off-ramp. | [optional] 
**OrderAmount** | Pointer to **string** | This field has been deprecated. Please use &#x60;pricing_amount&#x60; instead. | [optional] 
**SettlementStatus** | Pointer to [**SettleStatus**](SettleStatus.md) |  | [optional] 
**RefundId** | **string** | The refund order ID. | 
**Amount** | **string** | The amount in cryptocurrency to be returned for this refund order. | 
**ToAddress** | **string** | The recipient&#39;s wallet address where the refund will be sent. | 
**RefundType** | Pointer to [**RefundType**](RefundType.md) |  | [optional] 
**ChargeMerchantFee** | Pointer to **bool** | Whether to charge developer fee to the merchant for the refund.    - &#x60;true&#x60;: The fee amount (specified in &#x60;merchant_fee_amount&#x60;) will be deducted from the merchant&#39;s balance and added to the developer&#39;s balance    - &#x60;false&#x60;: The merchant is not charged any developer fee.  | [optional] 
**MerchantFeeAmount** | Pointer to **string** | The developer fee amount to charge the merchant, denominated in the cryptocurrency specified by &#x60;merchant_fee_token_id&#x60;. This is only applicable if &#x60;charge_merchant_fee&#x60; is set to &#x60;true&#x60;. | [optional] 
**MerchantFeeTokenId** | Pointer to **string** | The ID of the cryptocurrency used for the developer fee. This is only applicable if &#x60;charge_merchant_fee&#x60; is set to true. | [optional] 
**CommissionFee** | Pointer to [**CommissionFee**](CommissionFee.md) |  | [optional] 
**SettlementRequestId** | **string** | The settlement request ID generated by Cobo. | 
**Settlements** | [**[]SettlementDetail**](SettlementDetail.md) |  | 
**AcquiringType** | [**AcquiringType**](AcquiringType.md) |  | 
**PayoutChannel** | [**PayoutChannel**](PayoutChannel.md) |  | 
**SettlementType** | Pointer to [**SettlementType**](SettlementType.md) |  | [optional] 
**ReceivedAmountFiat** | Pointer to **string** | The estimated amount of the fiat currency to receive after off-ramping. This amount is subject to change due to bank transfer fees. | [optional] 
**BankAccount** | Pointer to [**BankAccount**](BankAccount.md) |  | [optional] 
**PayerId** | **string** | A unique identifier assigned by Cobo to track and identify individual payers. | 
**CustomPayerId** | **string** | A unique identifier assigned by the developer to track and identify individual payers in their system. | 
**Chain** | **string** | The chain ID. | 
**PreviousAddress** | **string** | The previous top-up address that was assigned to the payer. | 
**UpdatedAddress** | **string** | The new top-up address that has been assigned to the payer. | 
**PayoutId** | **string** | The payout ID generated by Cobo. | 
**SourceAccount** | **string** | The source account from which the bulk send will be made. - If the source account is a merchant account, provide the merchant&#39;s ID (e.g., \&quot;M1001\&quot;). - If the source account is the developer account, use the string &#x60;\&quot;developer\&quot;&#x60;.  | 
**PayoutItems** | Pointer to [**[]PaymentPayoutItem**](PaymentPayoutItem.md) | required | [optional] 
**RecipientInfo** | Pointer to [**PaymentPayoutRecipientInfo**](PaymentPayoutRecipientInfo.md) |  | [optional] 
**ActualPayoutAmount** | Pointer to **string** | - For &#x60;Crypto&#x60; payouts: The amount of cryptocurrency sent to the recipient&#39;s address, denominated in the token specified in &#x60;recipient_info.token_id&#x60;. - For &#x60;OffRamp&#x60; payouts: The amount of fiat currency sent to the recipient&#39;s bank account, denominated in the currency specified in &#x60;recipient_info.currency&#x60;. (Note: The actual amount received may be lower due to additional bank transfer fees.)  | [optional] 
**CommissionFees** | Pointer to [**[]CommissionFee**](CommissionFee.md) | The commission fees of the payout. | [optional] 
**Remark** | Pointer to **string** | A note or comment about the payout. | [optional] 
**BulkSendId** | **string** | The bulk send ID. | 
**ExecutionMode** | [**PaymentBulkSendExecutionMode**](PaymentBulkSendExecutionMode.md) |  | 
**DispositionType** | [**DispositionType**](DispositionType.md) |  | 
**DispositionStatus** | [**DispositionStatus**](DispositionStatus.md) |  | 
**DestinationAddress** | Pointer to **string** | The blockchain address to receive the refunded/isolated funds. | [optional] 
**DispositionAmount** | Pointer to **string** | The amount to be refunded/isolated from the original transaction, specified as a numeric string. This value cannot exceed the total amount of the original transaction.  | [optional] 
**TransactionType** | [**KytScreeningsTransactionType**](KytScreeningsTransactionType.md) |  | 
**ReviewStatus** | [**ReviewStatusType**](ReviewStatusType.md) |  | 
**FundsStatus** | [**FundsStatusType**](FundsStatusType.md) |  | 
**ScreeningId** | **string** | The unique system-generated identifier for this screening request (UUID format, fixed 36 characters). | 

## Methods

### NewWebhookEventData

`func NewWebhookEventData(dataType string, transactionId string, requestId string, walletId string, status KyaScreeningStatus, chainId string, tokenId string, source TransactionSource, destination TransactionDestination, initiatorType TransactionInitiatorType, createdTimestamp int64, updatedTimestamp int64, chains []ChainInfo, walletType WalletType, tokens []TokenInfo, contractAddress string, walletSubtype WalletSubtype, address string, walletUuid string, balance Balance, tokenIds string, operationType SuspendedTokenOperationType, orderId string, pspOrderCode string, feeAmount string, payableAmount string, exchangeRate string, receiveAddress string, receivedTokenAmount string, refundId string, amount string, toAddress string, settlementRequestId string, settlements []SettlementDetail, acquiringType AcquiringType, payoutChannel PayoutChannel, payerId string, customPayerId string, chain string, previousAddress string, updatedAddress string, payoutId string, sourceAccount string, bulkSendId string, executionMode PaymentBulkSendExecutionMode, dispositionType DispositionType, dispositionStatus DispositionStatus, transactionType KytScreeningsTransactionType, reviewStatus ReviewStatusType, fundsStatus FundsStatusType, screeningId string, ) *WebhookEventData`

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

`func (o *WebhookEventData) GetType() TransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookEventData) GetTypeOk() (*TransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookEventData) SetType(v TransactionType)`

SetType sets Type field to given value.

### HasType

`func (o *WebhookEventData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookEventData) GetStatus() KyaScreeningStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookEventData) GetStatusOk() (*KyaScreeningStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookEventData) SetStatus(v KyaScreeningStatus)`

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

`func (o *WebhookEventData) GetSource() TransactionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WebhookEventData) GetSourceOk() (*TransactionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WebhookEventData) SetSource(v TransactionSource)`

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

### GetAddress

`func (o *WebhookEventData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WebhookEventData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WebhookEventData) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetWalletUuid

`func (o *WebhookEventData) GetWalletUuid() string`

GetWalletUuid returns the WalletUuid field if non-nil, zero value otherwise.

### GetWalletUuidOk

`func (o *WebhookEventData) GetWalletUuidOk() (*string, bool)`

GetWalletUuidOk returns a tuple with the WalletUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletUuid

`func (o *WebhookEventData) SetWalletUuid(v string)`

SetWalletUuid sets WalletUuid field to given value.


### GetBalance

`func (o *WebhookEventData) GetBalance() Balance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *WebhookEventData) GetBalanceOk() (*Balance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *WebhookEventData) SetBalance(v Balance)`

SetBalance sets Balance field to given value.


### GetTokenIds

`func (o *WebhookEventData) GetTokenIds() string`

GetTokenIds returns the TokenIds field if non-nil, zero value otherwise.

### GetTokenIdsOk

`func (o *WebhookEventData) GetTokenIdsOk() (*string, bool)`

GetTokenIdsOk returns a tuple with the TokenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIds

`func (o *WebhookEventData) SetTokenIds(v string)`

SetTokenIds sets TokenIds field to given value.


### GetOperationType

`func (o *WebhookEventData) GetOperationType() SuspendedTokenOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *WebhookEventData) GetOperationTypeOk() (*SuspendedTokenOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *WebhookEventData) SetOperationType(v SuspendedTokenOperationType)`

SetOperationType sets OperationType field to given value.


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


### GetPricingCurrency

`func (o *WebhookEventData) GetPricingCurrency() string`

GetPricingCurrency returns the PricingCurrency field if non-nil, zero value otherwise.

### GetPricingCurrencyOk

`func (o *WebhookEventData) GetPricingCurrencyOk() (*string, bool)`

GetPricingCurrencyOk returns a tuple with the PricingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingCurrency

`func (o *WebhookEventData) SetPricingCurrency(v string)`

SetPricingCurrency sets PricingCurrency field to given value.

### HasPricingCurrency

`func (o *WebhookEventData) HasPricingCurrency() bool`

HasPricingCurrency returns a boolean if a field has been set.

### GetPricingAmount

`func (o *WebhookEventData) GetPricingAmount() string`

GetPricingAmount returns the PricingAmount field if non-nil, zero value otherwise.

### GetPricingAmountOk

`func (o *WebhookEventData) GetPricingAmountOk() (*string, bool)`

GetPricingAmountOk returns a tuple with the PricingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingAmount

`func (o *WebhookEventData) SetPricingAmount(v string)`

SetPricingAmount sets PricingAmount field to given value.

### HasPricingAmount

`func (o *WebhookEventData) HasPricingAmount() bool`

HasPricingAmount returns a boolean if a field has been set.

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


### GetPayableCurrency

`func (o *WebhookEventData) GetPayableCurrency() string`

GetPayableCurrency returns the PayableCurrency field if non-nil, zero value otherwise.

### GetPayableCurrencyOk

`func (o *WebhookEventData) GetPayableCurrencyOk() (*string, bool)`

GetPayableCurrencyOk returns a tuple with the PayableCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableCurrency

`func (o *WebhookEventData) SetPayableCurrency(v string)`

SetPayableCurrency sets PayableCurrency field to given value.

### HasPayableCurrency

`func (o *WebhookEventData) HasPayableCurrency() bool`

HasPayableCurrency returns a boolean if a field has been set.

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


### GetAmountTolerance

`func (o *WebhookEventData) GetAmountTolerance() string`

GetAmountTolerance returns the AmountTolerance field if non-nil, zero value otherwise.

### GetAmountToleranceOk

`func (o *WebhookEventData) GetAmountToleranceOk() (*string, bool)`

GetAmountToleranceOk returns a tuple with the AmountTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountTolerance

`func (o *WebhookEventData) SetAmountTolerance(v string)`

SetAmountTolerance sets AmountTolerance field to given value.

### HasAmountTolerance

`func (o *WebhookEventData) HasAmountTolerance() bool`

HasAmountTolerance returns a boolean if a field has been set.

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

### HasCurrency

`func (o *WebhookEventData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

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

### HasOrderAmount

`func (o *WebhookEventData) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetSettlementStatus

`func (o *WebhookEventData) GetSettlementStatus() SettleStatus`

GetSettlementStatus returns the SettlementStatus field if non-nil, zero value otherwise.

### GetSettlementStatusOk

`func (o *WebhookEventData) GetSettlementStatusOk() (*SettleStatus, bool)`

GetSettlementStatusOk returns a tuple with the SettlementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementStatus

`func (o *WebhookEventData) SetSettlementStatus(v SettleStatus)`

SetSettlementStatus sets SettlementStatus field to given value.

### HasSettlementStatus

`func (o *WebhookEventData) HasSettlementStatus() bool`

HasSettlementStatus returns a boolean if a field has been set.

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


### GetRefundType

`func (o *WebhookEventData) GetRefundType() RefundType`

GetRefundType returns the RefundType field if non-nil, zero value otherwise.

### GetRefundTypeOk

`func (o *WebhookEventData) GetRefundTypeOk() (*RefundType, bool)`

GetRefundTypeOk returns a tuple with the RefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundType

`func (o *WebhookEventData) SetRefundType(v RefundType)`

SetRefundType sets RefundType field to given value.

### HasRefundType

`func (o *WebhookEventData) HasRefundType() bool`

HasRefundType returns a boolean if a field has been set.

### GetChargeMerchantFee

`func (o *WebhookEventData) GetChargeMerchantFee() bool`

GetChargeMerchantFee returns the ChargeMerchantFee field if non-nil, zero value otherwise.

### GetChargeMerchantFeeOk

`func (o *WebhookEventData) GetChargeMerchantFeeOk() (*bool, bool)`

GetChargeMerchantFeeOk returns a tuple with the ChargeMerchantFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeMerchantFee

`func (o *WebhookEventData) SetChargeMerchantFee(v bool)`

SetChargeMerchantFee sets ChargeMerchantFee field to given value.

### HasChargeMerchantFee

`func (o *WebhookEventData) HasChargeMerchantFee() bool`

HasChargeMerchantFee returns a boolean if a field has been set.

### GetMerchantFeeAmount

`func (o *WebhookEventData) GetMerchantFeeAmount() string`

GetMerchantFeeAmount returns the MerchantFeeAmount field if non-nil, zero value otherwise.

### GetMerchantFeeAmountOk

`func (o *WebhookEventData) GetMerchantFeeAmountOk() (*string, bool)`

GetMerchantFeeAmountOk returns a tuple with the MerchantFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantFeeAmount

`func (o *WebhookEventData) SetMerchantFeeAmount(v string)`

SetMerchantFeeAmount sets MerchantFeeAmount field to given value.

### HasMerchantFeeAmount

`func (o *WebhookEventData) HasMerchantFeeAmount() bool`

HasMerchantFeeAmount returns a boolean if a field has been set.

### GetMerchantFeeTokenId

`func (o *WebhookEventData) GetMerchantFeeTokenId() string`

GetMerchantFeeTokenId returns the MerchantFeeTokenId field if non-nil, zero value otherwise.

### GetMerchantFeeTokenIdOk

`func (o *WebhookEventData) GetMerchantFeeTokenIdOk() (*string, bool)`

GetMerchantFeeTokenIdOk returns a tuple with the MerchantFeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantFeeTokenId

`func (o *WebhookEventData) SetMerchantFeeTokenId(v string)`

SetMerchantFeeTokenId sets MerchantFeeTokenId field to given value.

### HasMerchantFeeTokenId

`func (o *WebhookEventData) HasMerchantFeeTokenId() bool`

HasMerchantFeeTokenId returns a boolean if a field has been set.

### GetCommissionFee

`func (o *WebhookEventData) GetCommissionFee() CommissionFee`

GetCommissionFee returns the CommissionFee field if non-nil, zero value otherwise.

### GetCommissionFeeOk

`func (o *WebhookEventData) GetCommissionFeeOk() (*CommissionFee, bool)`

GetCommissionFeeOk returns a tuple with the CommissionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionFee

`func (o *WebhookEventData) SetCommissionFee(v CommissionFee)`

SetCommissionFee sets CommissionFee field to given value.

### HasCommissionFee

`func (o *WebhookEventData) HasCommissionFee() bool`

HasCommissionFee returns a boolean if a field has been set.

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


### GetAcquiringType

`func (o *WebhookEventData) GetAcquiringType() AcquiringType`

GetAcquiringType returns the AcquiringType field if non-nil, zero value otherwise.

### GetAcquiringTypeOk

`func (o *WebhookEventData) GetAcquiringTypeOk() (*AcquiringType, bool)`

GetAcquiringTypeOk returns a tuple with the AcquiringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquiringType

`func (o *WebhookEventData) SetAcquiringType(v AcquiringType)`

SetAcquiringType sets AcquiringType field to given value.


### GetPayoutChannel

`func (o *WebhookEventData) GetPayoutChannel() PayoutChannel`

GetPayoutChannel returns the PayoutChannel field if non-nil, zero value otherwise.

### GetPayoutChannelOk

`func (o *WebhookEventData) GetPayoutChannelOk() (*PayoutChannel, bool)`

GetPayoutChannelOk returns a tuple with the PayoutChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutChannel

`func (o *WebhookEventData) SetPayoutChannel(v PayoutChannel)`

SetPayoutChannel sets PayoutChannel field to given value.


### GetSettlementType

`func (o *WebhookEventData) GetSettlementType() SettlementType`

GetSettlementType returns the SettlementType field if non-nil, zero value otherwise.

### GetSettlementTypeOk

`func (o *WebhookEventData) GetSettlementTypeOk() (*SettlementType, bool)`

GetSettlementTypeOk returns a tuple with the SettlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementType

`func (o *WebhookEventData) SetSettlementType(v SettlementType)`

SetSettlementType sets SettlementType field to given value.

### HasSettlementType

`func (o *WebhookEventData) HasSettlementType() bool`

HasSettlementType returns a boolean if a field has been set.

### GetReceivedAmountFiat

`func (o *WebhookEventData) GetReceivedAmountFiat() string`

GetReceivedAmountFiat returns the ReceivedAmountFiat field if non-nil, zero value otherwise.

### GetReceivedAmountFiatOk

`func (o *WebhookEventData) GetReceivedAmountFiatOk() (*string, bool)`

GetReceivedAmountFiatOk returns a tuple with the ReceivedAmountFiat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAmountFiat

`func (o *WebhookEventData) SetReceivedAmountFiat(v string)`

SetReceivedAmountFiat sets ReceivedAmountFiat field to given value.

### HasReceivedAmountFiat

`func (o *WebhookEventData) HasReceivedAmountFiat() bool`

HasReceivedAmountFiat returns a boolean if a field has been set.

### GetBankAccount

`func (o *WebhookEventData) GetBankAccount() BankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *WebhookEventData) GetBankAccountOk() (*BankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *WebhookEventData) SetBankAccount(v BankAccount)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *WebhookEventData) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### GetPayerId

`func (o *WebhookEventData) GetPayerId() string`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *WebhookEventData) GetPayerIdOk() (*string, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *WebhookEventData) SetPayerId(v string)`

SetPayerId sets PayerId field to given value.


### GetCustomPayerId

`func (o *WebhookEventData) GetCustomPayerId() string`

GetCustomPayerId returns the CustomPayerId field if non-nil, zero value otherwise.

### GetCustomPayerIdOk

`func (o *WebhookEventData) GetCustomPayerIdOk() (*string, bool)`

GetCustomPayerIdOk returns a tuple with the CustomPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayerId

`func (o *WebhookEventData) SetCustomPayerId(v string)`

SetCustomPayerId sets CustomPayerId field to given value.


### GetChain

`func (o *WebhookEventData) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *WebhookEventData) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *WebhookEventData) SetChain(v string)`

SetChain sets Chain field to given value.


### GetPreviousAddress

`func (o *WebhookEventData) GetPreviousAddress() string`

GetPreviousAddress returns the PreviousAddress field if non-nil, zero value otherwise.

### GetPreviousAddressOk

`func (o *WebhookEventData) GetPreviousAddressOk() (*string, bool)`

GetPreviousAddressOk returns a tuple with the PreviousAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAddress

`func (o *WebhookEventData) SetPreviousAddress(v string)`

SetPreviousAddress sets PreviousAddress field to given value.


### GetUpdatedAddress

`func (o *WebhookEventData) GetUpdatedAddress() string`

GetUpdatedAddress returns the UpdatedAddress field if non-nil, zero value otherwise.

### GetUpdatedAddressOk

`func (o *WebhookEventData) GetUpdatedAddressOk() (*string, bool)`

GetUpdatedAddressOk returns a tuple with the UpdatedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAddress

`func (o *WebhookEventData) SetUpdatedAddress(v string)`

SetUpdatedAddress sets UpdatedAddress field to given value.


### GetPayoutId

`func (o *WebhookEventData) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *WebhookEventData) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *WebhookEventData) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.


### GetSourceAccount

`func (o *WebhookEventData) GetSourceAccount() string`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *WebhookEventData) GetSourceAccountOk() (*string, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *WebhookEventData) SetSourceAccount(v string)`

SetSourceAccount sets SourceAccount field to given value.


### GetPayoutItems

`func (o *WebhookEventData) GetPayoutItems() []PaymentPayoutItem`

GetPayoutItems returns the PayoutItems field if non-nil, zero value otherwise.

### GetPayoutItemsOk

`func (o *WebhookEventData) GetPayoutItemsOk() (*[]PaymentPayoutItem, bool)`

GetPayoutItemsOk returns a tuple with the PayoutItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutItems

`func (o *WebhookEventData) SetPayoutItems(v []PaymentPayoutItem)`

SetPayoutItems sets PayoutItems field to given value.

### HasPayoutItems

`func (o *WebhookEventData) HasPayoutItems() bool`

HasPayoutItems returns a boolean if a field has been set.

### GetRecipientInfo

`func (o *WebhookEventData) GetRecipientInfo() PaymentPayoutRecipientInfo`

GetRecipientInfo returns the RecipientInfo field if non-nil, zero value otherwise.

### GetRecipientInfoOk

`func (o *WebhookEventData) GetRecipientInfoOk() (*PaymentPayoutRecipientInfo, bool)`

GetRecipientInfoOk returns a tuple with the RecipientInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientInfo

`func (o *WebhookEventData) SetRecipientInfo(v PaymentPayoutRecipientInfo)`

SetRecipientInfo sets RecipientInfo field to given value.

### HasRecipientInfo

`func (o *WebhookEventData) HasRecipientInfo() bool`

HasRecipientInfo returns a boolean if a field has been set.

### GetActualPayoutAmount

`func (o *WebhookEventData) GetActualPayoutAmount() string`

GetActualPayoutAmount returns the ActualPayoutAmount field if non-nil, zero value otherwise.

### GetActualPayoutAmountOk

`func (o *WebhookEventData) GetActualPayoutAmountOk() (*string, bool)`

GetActualPayoutAmountOk returns a tuple with the ActualPayoutAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPayoutAmount

`func (o *WebhookEventData) SetActualPayoutAmount(v string)`

SetActualPayoutAmount sets ActualPayoutAmount field to given value.

### HasActualPayoutAmount

`func (o *WebhookEventData) HasActualPayoutAmount() bool`

HasActualPayoutAmount returns a boolean if a field has been set.

### GetCommissionFees

`func (o *WebhookEventData) GetCommissionFees() []CommissionFee`

GetCommissionFees returns the CommissionFees field if non-nil, zero value otherwise.

### GetCommissionFeesOk

`func (o *WebhookEventData) GetCommissionFeesOk() (*[]CommissionFee, bool)`

GetCommissionFeesOk returns a tuple with the CommissionFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionFees

`func (o *WebhookEventData) SetCommissionFees(v []CommissionFee)`

SetCommissionFees sets CommissionFees field to given value.

### HasCommissionFees

`func (o *WebhookEventData) HasCommissionFees() bool`

HasCommissionFees returns a boolean if a field has been set.

### GetRemark

`func (o *WebhookEventData) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *WebhookEventData) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *WebhookEventData) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *WebhookEventData) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### GetBulkSendId

`func (o *WebhookEventData) GetBulkSendId() string`

GetBulkSendId returns the BulkSendId field if non-nil, zero value otherwise.

### GetBulkSendIdOk

`func (o *WebhookEventData) GetBulkSendIdOk() (*string, bool)`

GetBulkSendIdOk returns a tuple with the BulkSendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkSendId

`func (o *WebhookEventData) SetBulkSendId(v string)`

SetBulkSendId sets BulkSendId field to given value.


### GetExecutionMode

`func (o *WebhookEventData) GetExecutionMode() PaymentBulkSendExecutionMode`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *WebhookEventData) GetExecutionModeOk() (*PaymentBulkSendExecutionMode, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *WebhookEventData) SetExecutionMode(v PaymentBulkSendExecutionMode)`

SetExecutionMode sets ExecutionMode field to given value.


### GetDispositionType

`func (o *WebhookEventData) GetDispositionType() DispositionType`

GetDispositionType returns the DispositionType field if non-nil, zero value otherwise.

### GetDispositionTypeOk

`func (o *WebhookEventData) GetDispositionTypeOk() (*DispositionType, bool)`

GetDispositionTypeOk returns a tuple with the DispositionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionType

`func (o *WebhookEventData) SetDispositionType(v DispositionType)`

SetDispositionType sets DispositionType field to given value.


### GetDispositionStatus

`func (o *WebhookEventData) GetDispositionStatus() DispositionStatus`

GetDispositionStatus returns the DispositionStatus field if non-nil, zero value otherwise.

### GetDispositionStatusOk

`func (o *WebhookEventData) GetDispositionStatusOk() (*DispositionStatus, bool)`

GetDispositionStatusOk returns a tuple with the DispositionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionStatus

`func (o *WebhookEventData) SetDispositionStatus(v DispositionStatus)`

SetDispositionStatus sets DispositionStatus field to given value.


### GetDestinationAddress

`func (o *WebhookEventData) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *WebhookEventData) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *WebhookEventData) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *WebhookEventData) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetDispositionAmount

`func (o *WebhookEventData) GetDispositionAmount() string`

GetDispositionAmount returns the DispositionAmount field if non-nil, zero value otherwise.

### GetDispositionAmountOk

`func (o *WebhookEventData) GetDispositionAmountOk() (*string, bool)`

GetDispositionAmountOk returns a tuple with the DispositionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionAmount

`func (o *WebhookEventData) SetDispositionAmount(v string)`

SetDispositionAmount sets DispositionAmount field to given value.

### HasDispositionAmount

`func (o *WebhookEventData) HasDispositionAmount() bool`

HasDispositionAmount returns a boolean if a field has been set.

### GetTransactionType

`func (o *WebhookEventData) GetTransactionType() KytScreeningsTransactionType`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *WebhookEventData) GetTransactionTypeOk() (*KytScreeningsTransactionType, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *WebhookEventData) SetTransactionType(v KytScreeningsTransactionType)`

SetTransactionType sets TransactionType field to given value.


### GetReviewStatus

`func (o *WebhookEventData) GetReviewStatus() ReviewStatusType`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *WebhookEventData) GetReviewStatusOk() (*ReviewStatusType, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *WebhookEventData) SetReviewStatus(v ReviewStatusType)`

SetReviewStatus sets ReviewStatus field to given value.


### GetFundsStatus

`func (o *WebhookEventData) GetFundsStatus() FundsStatusType`

GetFundsStatus returns the FundsStatus field if non-nil, zero value otherwise.

### GetFundsStatusOk

`func (o *WebhookEventData) GetFundsStatusOk() (*FundsStatusType, bool)`

GetFundsStatusOk returns a tuple with the FundsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsStatus

`func (o *WebhookEventData) SetFundsStatus(v FundsStatusType)`

SetFundsStatus sets FundsStatus field to given value.


### GetScreeningId

`func (o *WebhookEventData) GetScreeningId() string`

GetScreeningId returns the ScreeningId field if non-nil, zero value otherwise.

### GetScreeningIdOk

`func (o *WebhookEventData) GetScreeningIdOk() (*string, bool)`

GetScreeningIdOk returns a tuple with the ScreeningId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreeningId

`func (o *WebhookEventData) SetScreeningId(v string)`

SetScreeningId sets ScreeningId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


