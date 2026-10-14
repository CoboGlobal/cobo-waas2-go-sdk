# PaymentOrderUpdateEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBankWithdrawal&#x60;: The payment bank withdrawal event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;PaymentBulkSendItem&#x60;: The payment bulk send item event data. - &#x60;PaymentTopUpUpdate&#x60;: The payment top-up update event data. - &#x60;PaymentOrderUpdate&#x60;: The payment order update event data. - &#x60;PaymentRefundUpdate&#x60;: The payment refund update event data. - &#x60;PaymentPayoutUpdate&#x60;: The payment payout update event data. - &#x60;PaymentBulkSendUpdate&#x60;: The payment bulk send update event data. - &#x60;PaymentBulkSendItemUpdate&#x60;: The payment bulk send item update event data. - &#x60;PaymentAccountBalanceUpdate&#x60;: The Payments account balance updated event data, including account information and balance change details. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. - &#x60;Organization&#x60;: The organization event data. - &#x60;FiatTransaction&#x60;: The fiat transaction event data. | 
**OrderId** | **string** |  The order ID to be included in this settlement.  - If provided: Settles merchant funds received from this specific order - If empty: Settles merchant funds received from all unsettled orders of the specified merchant  **Usage conditions** - Only applicable when &#x60;settlement_type&#x60; is set to &#x60;Merchant&#x60; - Cannot be used together with &#x60;amount&#x60; parameter        | 
**PspOrderCode** | **string** | A unique reference code assigned by the developer to identify this order in their system. | 
**MerchantId** | **string** | The merchant ID associated with the order. | 
**MerchantOrderCode** | **NullableString** | A unique reference code assigned by the merchant to identify this order in their system. | 
**Status** | [**OrderStatus**](OrderStatus.md) |  | 
**NoticeType** | [**NullablePaymentOrderNoticeType**](PaymentOrderNoticeType.md) |  | 
**ReceiveAddress** | **string** | The recipient wallet address for the payment transaction. | 
**ExpiredAt** | **int64** | The expiration time of the pay-in order, represented as a UNIX timestamp in seconds. | 
**RequestedAmount** | [**PaymentAssetAmount**](PaymentAssetAmount.md) | The requested amount. | 
**CreditedAmount** | [**PaymentAssetAmount**](PaymentAssetAmount.md) | The credited amount. | 
**OutstandingAmount** | [**PaymentAssetAmount**](PaymentAssetAmount.md) | The outstanding amount. | 
**Transactions** | [**[]PaymentOrderDeposit**](PaymentOrderDeposit.md) | The incoming transactions associated with the order. | 
**CreatedTimestamp** | **int64** | The created time of the order, represented as a UNIX timestamp in seconds. | 
**UpdatedTimestamp** | **int64** | The updated time of the order, represented as a UNIX timestamp in seconds. | 

## Methods

### NewPaymentOrderUpdateEventData

`func NewPaymentOrderUpdateEventData(dataType string, orderId string, pspOrderCode string, merchantId string, merchantOrderCode NullableString, status OrderStatus, noticeType NullablePaymentOrderNoticeType, receiveAddress string, expiredAt int64, requestedAmount PaymentAssetAmount, creditedAmount PaymentAssetAmount, outstandingAmount PaymentAssetAmount, transactions []PaymentOrderDeposit, createdTimestamp int64, updatedTimestamp int64, ) *PaymentOrderUpdateEventData`

NewPaymentOrderUpdateEventData instantiates a new PaymentOrderUpdateEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOrderUpdateEventDataWithDefaults

`func NewPaymentOrderUpdateEventDataWithDefaults() *PaymentOrderUpdateEventData`

NewPaymentOrderUpdateEventDataWithDefaults instantiates a new PaymentOrderUpdateEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *PaymentOrderUpdateEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PaymentOrderUpdateEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PaymentOrderUpdateEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetOrderId

`func (o *PaymentOrderUpdateEventData) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PaymentOrderUpdateEventData) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PaymentOrderUpdateEventData) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetPspOrderCode

`func (o *PaymentOrderUpdateEventData) GetPspOrderCode() string`

GetPspOrderCode returns the PspOrderCode field if non-nil, zero value otherwise.

### GetPspOrderCodeOk

`func (o *PaymentOrderUpdateEventData) GetPspOrderCodeOk() (*string, bool)`

GetPspOrderCodeOk returns a tuple with the PspOrderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspOrderCode

`func (o *PaymentOrderUpdateEventData) SetPspOrderCode(v string)`

SetPspOrderCode sets PspOrderCode field to given value.


### GetMerchantId

`func (o *PaymentOrderUpdateEventData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PaymentOrderUpdateEventData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PaymentOrderUpdateEventData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetMerchantOrderCode

`func (o *PaymentOrderUpdateEventData) GetMerchantOrderCode() string`

GetMerchantOrderCode returns the MerchantOrderCode field if non-nil, zero value otherwise.

### GetMerchantOrderCodeOk

`func (o *PaymentOrderUpdateEventData) GetMerchantOrderCodeOk() (*string, bool)`

GetMerchantOrderCodeOk returns a tuple with the MerchantOrderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderCode

`func (o *PaymentOrderUpdateEventData) SetMerchantOrderCode(v string)`

SetMerchantOrderCode sets MerchantOrderCode field to given value.


### SetMerchantOrderCodeNil

`func (o *PaymentOrderUpdateEventData) SetMerchantOrderCodeNil(b bool)`

 SetMerchantOrderCodeNil sets the value for MerchantOrderCode to be an explicit nil

### UnsetMerchantOrderCode
`func (o *PaymentOrderUpdateEventData) UnsetMerchantOrderCode()`

UnsetMerchantOrderCode ensures that no value is present for MerchantOrderCode, not even an explicit nil
### GetStatus

`func (o *PaymentOrderUpdateEventData) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentOrderUpdateEventData) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentOrderUpdateEventData) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.


### GetNoticeType

`func (o *PaymentOrderUpdateEventData) GetNoticeType() PaymentOrderNoticeType`

GetNoticeType returns the NoticeType field if non-nil, zero value otherwise.

### GetNoticeTypeOk

`func (o *PaymentOrderUpdateEventData) GetNoticeTypeOk() (*PaymentOrderNoticeType, bool)`

GetNoticeTypeOk returns a tuple with the NoticeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeType

`func (o *PaymentOrderUpdateEventData) SetNoticeType(v PaymentOrderNoticeType)`

SetNoticeType sets NoticeType field to given value.


### SetNoticeTypeNil

`func (o *PaymentOrderUpdateEventData) SetNoticeTypeNil(b bool)`

 SetNoticeTypeNil sets the value for NoticeType to be an explicit nil

### UnsetNoticeType
`func (o *PaymentOrderUpdateEventData) UnsetNoticeType()`

UnsetNoticeType ensures that no value is present for NoticeType, not even an explicit nil
### GetReceiveAddress

`func (o *PaymentOrderUpdateEventData) GetReceiveAddress() string`

GetReceiveAddress returns the ReceiveAddress field if non-nil, zero value otherwise.

### GetReceiveAddressOk

`func (o *PaymentOrderUpdateEventData) GetReceiveAddressOk() (*string, bool)`

GetReceiveAddressOk returns a tuple with the ReceiveAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveAddress

`func (o *PaymentOrderUpdateEventData) SetReceiveAddress(v string)`

SetReceiveAddress sets ReceiveAddress field to given value.


### GetExpiredAt

`func (o *PaymentOrderUpdateEventData) GetExpiredAt() int64`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *PaymentOrderUpdateEventData) GetExpiredAtOk() (*int64, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *PaymentOrderUpdateEventData) SetExpiredAt(v int64)`

SetExpiredAt sets ExpiredAt field to given value.


### GetRequestedAmount

`func (o *PaymentOrderUpdateEventData) GetRequestedAmount() PaymentAssetAmount`

GetRequestedAmount returns the RequestedAmount field if non-nil, zero value otherwise.

### GetRequestedAmountOk

`func (o *PaymentOrderUpdateEventData) GetRequestedAmountOk() (*PaymentAssetAmount, bool)`

GetRequestedAmountOk returns a tuple with the RequestedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAmount

`func (o *PaymentOrderUpdateEventData) SetRequestedAmount(v PaymentAssetAmount)`

SetRequestedAmount sets RequestedAmount field to given value.


### GetCreditedAmount

`func (o *PaymentOrderUpdateEventData) GetCreditedAmount() PaymentAssetAmount`

GetCreditedAmount returns the CreditedAmount field if non-nil, zero value otherwise.

### GetCreditedAmountOk

`func (o *PaymentOrderUpdateEventData) GetCreditedAmountOk() (*PaymentAssetAmount, bool)`

GetCreditedAmountOk returns a tuple with the CreditedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditedAmount

`func (o *PaymentOrderUpdateEventData) SetCreditedAmount(v PaymentAssetAmount)`

SetCreditedAmount sets CreditedAmount field to given value.


### GetOutstandingAmount

`func (o *PaymentOrderUpdateEventData) GetOutstandingAmount() PaymentAssetAmount`

GetOutstandingAmount returns the OutstandingAmount field if non-nil, zero value otherwise.

### GetOutstandingAmountOk

`func (o *PaymentOrderUpdateEventData) GetOutstandingAmountOk() (*PaymentAssetAmount, bool)`

GetOutstandingAmountOk returns a tuple with the OutstandingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstandingAmount

`func (o *PaymentOrderUpdateEventData) SetOutstandingAmount(v PaymentAssetAmount)`

SetOutstandingAmount sets OutstandingAmount field to given value.


### GetTransactions

`func (o *PaymentOrderUpdateEventData) GetTransactions() []PaymentOrderDeposit`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *PaymentOrderUpdateEventData) GetTransactionsOk() (*[]PaymentOrderDeposit, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *PaymentOrderUpdateEventData) SetTransactions(v []PaymentOrderDeposit)`

SetTransactions sets Transactions field to given value.


### GetCreatedTimestamp

`func (o *PaymentOrderUpdateEventData) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentOrderUpdateEventData) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentOrderUpdateEventData) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *PaymentOrderUpdateEventData) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentOrderUpdateEventData) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentOrderUpdateEventData) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


