# PaymentRefundEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. | 
**RequestId** | Pointer to **string** | The request ID provided by you when creating the refund request. | [optional] 
**RefundId** | **string** | The refund order ID. | 
**OrderId** | Pointer to **string** | The ID of the pay-in order corresponding to this refund. | [optional] 
**MerchantId** | Pointer to **string** | The merchant ID. | [optional] 
**TokenId** | **string** | The ID of the cryptocurrency used for refund. | 
**ChainId** | **string** | The ID of the blockchain network on which the refund transaction occurs. | 
**Amount** | **string** | The amount in cryptocurrency to be returned for this refund order. | 
**ToAddress** | **string** | The recipient&#39;s wallet address where the refund will be sent. | 
**Status** | [**RefundStatus**](RefundStatus.md) |  | 
**RefundType** | Pointer to [**RefundType**](RefundType.md) |  | [optional] 
**CreatedTimestamp** | Pointer to **int32** | The creation time of the refund order, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The last update time of the refund order, represented as a UNIX timestamp in seconds. | [optional] 
**Initiator** | Pointer to **string** |  The initiator of this settlement request. Can return either an API key or the Payments App&#39;s ID.  - Format &#x60;api_key_&lt;API_KEY&gt;&#x60;: Indicates the settlement request was initiated via the Payments API using the API key. - Format &#x60;app_&lt;APP_ID&gt;&#x60;: Indicates the settlement request was initiated through the Payments App using the App ID.  | [optional] 
**Transactions** | Pointer to [**[]PaymentTransaction**](PaymentTransaction.md) | An array of transactions associated with this refund order. Each transaction represents a separate blockchain operation related to the refund process. | [optional] 
**ChargeMerchantFee** | Pointer to **bool** | Whether to charge developer fee to the merchant for the refund.    - &#x60;true&#x60;: The fee amount (specified in &#x60;merchant_fee_amount&#x60;) will be deducted from the merchant&#39;s balance and added to the developer&#39;s balance    - &#x60;false&#x60;: The merchant is not charged any developer fee.  | [optional] 
**MerchantFeeAmount** | Pointer to **string** | The developer fee amount to charge the merchant, denominated in the cryptocurrency specified by &#x60;merchant_fee_token_id&#x60;. This is only applicable if &#x60;charge_merchant_fee&#x60; is set to &#x60;true&#x60;. | [optional] 
**MerchantFeeTokenId** | Pointer to **string** | The ID of the cryptocurrency used for the developer fee. This is only applicable if &#x60;charge_merchant_fee&#x60; is set to true. | [optional] 
**CommissionFee** | Pointer to [**CommissionFee**](CommissionFee.md) |  | [optional] 

## Methods

### NewPaymentRefundEventData

`func NewPaymentRefundEventData(dataType string, refundId string, tokenId string, chainId string, amount string, toAddress string, status RefundStatus, ) *PaymentRefundEventData`

NewPaymentRefundEventData instantiates a new PaymentRefundEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRefundEventDataWithDefaults

`func NewPaymentRefundEventDataWithDefaults() *PaymentRefundEventData`

NewPaymentRefundEventDataWithDefaults instantiates a new PaymentRefundEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *PaymentRefundEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PaymentRefundEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PaymentRefundEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetRequestId

`func (o *PaymentRefundEventData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentRefundEventData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentRefundEventData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *PaymentRefundEventData) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRefundId

`func (o *PaymentRefundEventData) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *PaymentRefundEventData) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *PaymentRefundEventData) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.


### GetOrderId

`func (o *PaymentRefundEventData) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PaymentRefundEventData) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PaymentRefundEventData) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PaymentRefundEventData) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetMerchantId

`func (o *PaymentRefundEventData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PaymentRefundEventData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PaymentRefundEventData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *PaymentRefundEventData) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetTokenId

`func (o *PaymentRefundEventData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentRefundEventData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentRefundEventData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *PaymentRefundEventData) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *PaymentRefundEventData) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *PaymentRefundEventData) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetAmount

`func (o *PaymentRefundEventData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentRefundEventData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentRefundEventData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetToAddress

`func (o *PaymentRefundEventData) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *PaymentRefundEventData) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *PaymentRefundEventData) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetStatus

`func (o *PaymentRefundEventData) GetStatus() RefundStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentRefundEventData) GetStatusOk() (*RefundStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentRefundEventData) SetStatus(v RefundStatus)`

SetStatus sets Status field to given value.


### GetRefundType

`func (o *PaymentRefundEventData) GetRefundType() RefundType`

GetRefundType returns the RefundType field if non-nil, zero value otherwise.

### GetRefundTypeOk

`func (o *PaymentRefundEventData) GetRefundTypeOk() (*RefundType, bool)`

GetRefundTypeOk returns a tuple with the RefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundType

`func (o *PaymentRefundEventData) SetRefundType(v RefundType)`

SetRefundType sets RefundType field to given value.

### HasRefundType

`func (o *PaymentRefundEventData) HasRefundType() bool`

HasRefundType returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *PaymentRefundEventData) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentRefundEventData) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentRefundEventData) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *PaymentRefundEventData) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *PaymentRefundEventData) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentRefundEventData) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentRefundEventData) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *PaymentRefundEventData) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetInitiator

`func (o *PaymentRefundEventData) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *PaymentRefundEventData) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *PaymentRefundEventData) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *PaymentRefundEventData) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetTransactions

`func (o *PaymentRefundEventData) GetTransactions() []PaymentTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *PaymentRefundEventData) GetTransactionsOk() (*[]PaymentTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *PaymentRefundEventData) SetTransactions(v []PaymentTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *PaymentRefundEventData) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetChargeMerchantFee

`func (o *PaymentRefundEventData) GetChargeMerchantFee() bool`

GetChargeMerchantFee returns the ChargeMerchantFee field if non-nil, zero value otherwise.

### GetChargeMerchantFeeOk

`func (o *PaymentRefundEventData) GetChargeMerchantFeeOk() (*bool, bool)`

GetChargeMerchantFeeOk returns a tuple with the ChargeMerchantFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeMerchantFee

`func (o *PaymentRefundEventData) SetChargeMerchantFee(v bool)`

SetChargeMerchantFee sets ChargeMerchantFee field to given value.

### HasChargeMerchantFee

`func (o *PaymentRefundEventData) HasChargeMerchantFee() bool`

HasChargeMerchantFee returns a boolean if a field has been set.

### GetMerchantFeeAmount

`func (o *PaymentRefundEventData) GetMerchantFeeAmount() string`

GetMerchantFeeAmount returns the MerchantFeeAmount field if non-nil, zero value otherwise.

### GetMerchantFeeAmountOk

`func (o *PaymentRefundEventData) GetMerchantFeeAmountOk() (*string, bool)`

GetMerchantFeeAmountOk returns a tuple with the MerchantFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantFeeAmount

`func (o *PaymentRefundEventData) SetMerchantFeeAmount(v string)`

SetMerchantFeeAmount sets MerchantFeeAmount field to given value.

### HasMerchantFeeAmount

`func (o *PaymentRefundEventData) HasMerchantFeeAmount() bool`

HasMerchantFeeAmount returns a boolean if a field has been set.

### GetMerchantFeeTokenId

`func (o *PaymentRefundEventData) GetMerchantFeeTokenId() string`

GetMerchantFeeTokenId returns the MerchantFeeTokenId field if non-nil, zero value otherwise.

### GetMerchantFeeTokenIdOk

`func (o *PaymentRefundEventData) GetMerchantFeeTokenIdOk() (*string, bool)`

GetMerchantFeeTokenIdOk returns a tuple with the MerchantFeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantFeeTokenId

`func (o *PaymentRefundEventData) SetMerchantFeeTokenId(v string)`

SetMerchantFeeTokenId sets MerchantFeeTokenId field to given value.

### HasMerchantFeeTokenId

`func (o *PaymentRefundEventData) HasMerchantFeeTokenId() bool`

HasMerchantFeeTokenId returns a boolean if a field has been set.

### GetCommissionFee

`func (o *PaymentRefundEventData) GetCommissionFee() CommissionFee`

GetCommissionFee returns the CommissionFee field if non-nil, zero value otherwise.

### GetCommissionFeeOk

`func (o *PaymentRefundEventData) GetCommissionFeeOk() (*CommissionFee, bool)`

GetCommissionFeeOk returns a tuple with the CommissionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionFee

`func (o *PaymentRefundEventData) SetCommissionFee(v CommissionFee)`

SetCommissionFee sets CommissionFee field to given value.

### HasCommissionFee

`func (o *PaymentRefundEventData) HasCommissionFee() bool`

HasCommissionFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


