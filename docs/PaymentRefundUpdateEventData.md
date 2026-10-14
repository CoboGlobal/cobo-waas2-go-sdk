# PaymentRefundUpdateEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBankWithdrawal&#x60;: The payment bank withdrawal event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;PaymentBulkSendItem&#x60;: The payment bulk send item event data. - &#x60;PaymentTopUpUpdate&#x60;: The payment top-up update event data. - &#x60;PaymentOrderUpdate&#x60;: The payment order update event data. - &#x60;PaymentRefundUpdate&#x60;: The payment refund update event data. - &#x60;PaymentPayoutUpdate&#x60;: The payment payout update event data. - &#x60;PaymentBulkSendUpdate&#x60;: The payment bulk send update event data. - &#x60;PaymentBulkSendItemUpdate&#x60;: The payment bulk send item update event data. - &#x60;PaymentAccountBalanceUpdate&#x60;: The Payments account balance updated event data, including account information and balance change details. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. - &#x60;Organization&#x60;: The organization event data. - &#x60;FiatTransaction&#x60;: The fiat transaction event data. | 
**RefundId** | **string** | The refund order ID. | 
**OrderId** | **NullableString** | Unique identifier of a single order | 
**MerchantId** | **NullableString** | The merchant ID associated with the refund. | 
**RequestId** | **NullableString** | The request ID for the refund request. | 
**ToAddress** | **NullableString** | The recipient wallet address for the refund. | 
**Status** | [**RefundStatus**](RefundStatus.md) |  | 
**RequestedAmount** | [**PaymentAssetAmount**](PaymentAssetAmount.md) |  | 
**TxHash** | **NullableString** | The blockchain transaction hash. | 
**RefundType** | [**RefundType**](RefundType.md) |  | 
**Initiator** | **NullableString** | The initiator of the refund order. | 
**CommissionFee** | [**NullablePaymentCommissionFeeTotal**](PaymentCommissionFeeTotal.md) |  | 
**FailedReason** | [**NullablePaymentFailedReason**](PaymentFailedReason.md) |  | 
**CreatedTimestamp** | **int64** | The created time of the refund order, represented as a UNIX timestamp in seconds. | 
**UpdatedTimestamp** | **int64** | The updated time of the refund order, represented as a UNIX timestamp in seconds. | 

## Methods

### NewPaymentRefundUpdateEventData

`func NewPaymentRefundUpdateEventData(dataType string, refundId string, orderId NullableString, merchantId NullableString, requestId NullableString, toAddress NullableString, status RefundStatus, requestedAmount PaymentAssetAmount, txHash NullableString, refundType RefundType, initiator NullableString, commissionFee NullablePaymentCommissionFeeTotal, failedReason NullablePaymentFailedReason, createdTimestamp int64, updatedTimestamp int64, ) *PaymentRefundUpdateEventData`

NewPaymentRefundUpdateEventData instantiates a new PaymentRefundUpdateEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRefundUpdateEventDataWithDefaults

`func NewPaymentRefundUpdateEventDataWithDefaults() *PaymentRefundUpdateEventData`

NewPaymentRefundUpdateEventDataWithDefaults instantiates a new PaymentRefundUpdateEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *PaymentRefundUpdateEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PaymentRefundUpdateEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PaymentRefundUpdateEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetRefundId

`func (o *PaymentRefundUpdateEventData) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *PaymentRefundUpdateEventData) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *PaymentRefundUpdateEventData) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.


### GetOrderId

`func (o *PaymentRefundUpdateEventData) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PaymentRefundUpdateEventData) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PaymentRefundUpdateEventData) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### SetOrderIdNil

`func (o *PaymentRefundUpdateEventData) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *PaymentRefundUpdateEventData) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetMerchantId

`func (o *PaymentRefundUpdateEventData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PaymentRefundUpdateEventData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PaymentRefundUpdateEventData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### SetMerchantIdNil

`func (o *PaymentRefundUpdateEventData) SetMerchantIdNil(b bool)`

 SetMerchantIdNil sets the value for MerchantId to be an explicit nil

### UnsetMerchantId
`func (o *PaymentRefundUpdateEventData) UnsetMerchantId()`

UnsetMerchantId ensures that no value is present for MerchantId, not even an explicit nil
### GetRequestId

`func (o *PaymentRefundUpdateEventData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentRefundUpdateEventData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentRefundUpdateEventData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### SetRequestIdNil

`func (o *PaymentRefundUpdateEventData) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *PaymentRefundUpdateEventData) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetToAddress

`func (o *PaymentRefundUpdateEventData) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *PaymentRefundUpdateEventData) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *PaymentRefundUpdateEventData) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### SetToAddressNil

`func (o *PaymentRefundUpdateEventData) SetToAddressNil(b bool)`

 SetToAddressNil sets the value for ToAddress to be an explicit nil

### UnsetToAddress
`func (o *PaymentRefundUpdateEventData) UnsetToAddress()`

UnsetToAddress ensures that no value is present for ToAddress, not even an explicit nil
### GetStatus

`func (o *PaymentRefundUpdateEventData) GetStatus() RefundStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentRefundUpdateEventData) GetStatusOk() (*RefundStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentRefundUpdateEventData) SetStatus(v RefundStatus)`

SetStatus sets Status field to given value.


### GetRequestedAmount

`func (o *PaymentRefundUpdateEventData) GetRequestedAmount() PaymentAssetAmount`

GetRequestedAmount returns the RequestedAmount field if non-nil, zero value otherwise.

### GetRequestedAmountOk

`func (o *PaymentRefundUpdateEventData) GetRequestedAmountOk() (*PaymentAssetAmount, bool)`

GetRequestedAmountOk returns a tuple with the RequestedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAmount

`func (o *PaymentRefundUpdateEventData) SetRequestedAmount(v PaymentAssetAmount)`

SetRequestedAmount sets RequestedAmount field to given value.


### GetTxHash

`func (o *PaymentRefundUpdateEventData) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *PaymentRefundUpdateEventData) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *PaymentRefundUpdateEventData) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### SetTxHashNil

`func (o *PaymentRefundUpdateEventData) SetTxHashNil(b bool)`

 SetTxHashNil sets the value for TxHash to be an explicit nil

### UnsetTxHash
`func (o *PaymentRefundUpdateEventData) UnsetTxHash()`

UnsetTxHash ensures that no value is present for TxHash, not even an explicit nil
### GetRefundType

`func (o *PaymentRefundUpdateEventData) GetRefundType() RefundType`

GetRefundType returns the RefundType field if non-nil, zero value otherwise.

### GetRefundTypeOk

`func (o *PaymentRefundUpdateEventData) GetRefundTypeOk() (*RefundType, bool)`

GetRefundTypeOk returns a tuple with the RefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundType

`func (o *PaymentRefundUpdateEventData) SetRefundType(v RefundType)`

SetRefundType sets RefundType field to given value.


### GetInitiator

`func (o *PaymentRefundUpdateEventData) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *PaymentRefundUpdateEventData) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *PaymentRefundUpdateEventData) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.


### SetInitiatorNil

`func (o *PaymentRefundUpdateEventData) SetInitiatorNil(b bool)`

 SetInitiatorNil sets the value for Initiator to be an explicit nil

### UnsetInitiator
`func (o *PaymentRefundUpdateEventData) UnsetInitiator()`

UnsetInitiator ensures that no value is present for Initiator, not even an explicit nil
### GetCommissionFee

`func (o *PaymentRefundUpdateEventData) GetCommissionFee() PaymentCommissionFeeTotal`

GetCommissionFee returns the CommissionFee field if non-nil, zero value otherwise.

### GetCommissionFeeOk

`func (o *PaymentRefundUpdateEventData) GetCommissionFeeOk() (*PaymentCommissionFeeTotal, bool)`

GetCommissionFeeOk returns a tuple with the CommissionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionFee

`func (o *PaymentRefundUpdateEventData) SetCommissionFee(v PaymentCommissionFeeTotal)`

SetCommissionFee sets CommissionFee field to given value.


### SetCommissionFeeNil

`func (o *PaymentRefundUpdateEventData) SetCommissionFeeNil(b bool)`

 SetCommissionFeeNil sets the value for CommissionFee to be an explicit nil

### UnsetCommissionFee
`func (o *PaymentRefundUpdateEventData) UnsetCommissionFee()`

UnsetCommissionFee ensures that no value is present for CommissionFee, not even an explicit nil
### GetFailedReason

`func (o *PaymentRefundUpdateEventData) GetFailedReason() PaymentFailedReason`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *PaymentRefundUpdateEventData) GetFailedReasonOk() (*PaymentFailedReason, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *PaymentRefundUpdateEventData) SetFailedReason(v PaymentFailedReason)`

SetFailedReason sets FailedReason field to given value.


### SetFailedReasonNil

`func (o *PaymentRefundUpdateEventData) SetFailedReasonNil(b bool)`

 SetFailedReasonNil sets the value for FailedReason to be an explicit nil

### UnsetFailedReason
`func (o *PaymentRefundUpdateEventData) UnsetFailedReason()`

UnsetFailedReason ensures that no value is present for FailedReason, not even an explicit nil
### GetCreatedTimestamp

`func (o *PaymentRefundUpdateEventData) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentRefundUpdateEventData) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentRefundUpdateEventData) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *PaymentRefundUpdateEventData) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentRefundUpdateEventData) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentRefundUpdateEventData) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


