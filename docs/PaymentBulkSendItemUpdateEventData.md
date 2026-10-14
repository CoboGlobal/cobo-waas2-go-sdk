# PaymentBulkSendItemUpdateEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBankWithdrawal&#x60;: The payment bank withdrawal event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;PaymentBulkSendItem&#x60;: The payment bulk send item event data. - &#x60;PaymentTopUpUpdate&#x60;: The payment top-up update event data. - &#x60;PaymentOrderUpdate&#x60;: The payment order update event data. - &#x60;PaymentRefundUpdate&#x60;: The payment refund update event data. - &#x60;PaymentPayoutUpdate&#x60;: The payment payout update event data. - &#x60;PaymentBulkSendUpdate&#x60;: The payment bulk send update event data. - &#x60;PaymentBulkSendItemUpdate&#x60;: The payment bulk send item update event data. - &#x60;PaymentAccountBalanceUpdate&#x60;: The Payments account balance updated event data, including account information and balance change details. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. - &#x60;Organization&#x60;: The organization event data. - &#x60;FiatTransaction&#x60;: The fiat transaction event data. | 
**BulkSendItemId** | **string** | The bulk send item ID. | 
**BulkSendId** | **string** | The ID of the bulk send that contains this item. | 
**RequestId** | **NullableString** | The request ID for the item. | 
**Status** | [**PaymentBulkSendItemStatus**](PaymentBulkSendItemStatus.md) |  | 
**RequestedAmount** | [**PaymentAssetAmount**](PaymentAssetAmount.md) |  | 
**ReceivingAddress** | **string** | The receiving address of the payout item. | 
**TxHash** | **NullableString** | The blockchain transaction hash. | 
**SourceAccount** | **string** | The Payments account from which the funds are debited: the merchant ID for a merchant account, or &#x60;developer&#x60; for the developer account. | 
**Description** | **NullableString** | The description of the payout item. | 
**FailedReason** | [**NullablePaymentFailedReason**](PaymentFailedReason.md) |  | 
**CreatedTimestamp** | **int64** | The created time of the bulk send item, represented as a UNIX timestamp in seconds. | 
**UpdatedTimestamp** | **int64** | The updated time of the bulk send item, represented as a UNIX timestamp in seconds. | 

## Methods

### NewPaymentBulkSendItemUpdateEventData

`func NewPaymentBulkSendItemUpdateEventData(dataType string, bulkSendItemId string, bulkSendId string, requestId NullableString, status PaymentBulkSendItemStatus, requestedAmount PaymentAssetAmount, receivingAddress string, txHash NullableString, sourceAccount string, description NullableString, failedReason NullablePaymentFailedReason, createdTimestamp int64, updatedTimestamp int64, ) *PaymentBulkSendItemUpdateEventData`

NewPaymentBulkSendItemUpdateEventData instantiates a new PaymentBulkSendItemUpdateEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentBulkSendItemUpdateEventDataWithDefaults

`func NewPaymentBulkSendItemUpdateEventDataWithDefaults() *PaymentBulkSendItemUpdateEventData`

NewPaymentBulkSendItemUpdateEventDataWithDefaults instantiates a new PaymentBulkSendItemUpdateEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *PaymentBulkSendItemUpdateEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PaymentBulkSendItemUpdateEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PaymentBulkSendItemUpdateEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetBulkSendItemId

`func (o *PaymentBulkSendItemUpdateEventData) GetBulkSendItemId() string`

GetBulkSendItemId returns the BulkSendItemId field if non-nil, zero value otherwise.

### GetBulkSendItemIdOk

`func (o *PaymentBulkSendItemUpdateEventData) GetBulkSendItemIdOk() (*string, bool)`

GetBulkSendItemIdOk returns a tuple with the BulkSendItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkSendItemId

`func (o *PaymentBulkSendItemUpdateEventData) SetBulkSendItemId(v string)`

SetBulkSendItemId sets BulkSendItemId field to given value.


### GetBulkSendId

`func (o *PaymentBulkSendItemUpdateEventData) GetBulkSendId() string`

GetBulkSendId returns the BulkSendId field if non-nil, zero value otherwise.

### GetBulkSendIdOk

`func (o *PaymentBulkSendItemUpdateEventData) GetBulkSendIdOk() (*string, bool)`

GetBulkSendIdOk returns a tuple with the BulkSendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkSendId

`func (o *PaymentBulkSendItemUpdateEventData) SetBulkSendId(v string)`

SetBulkSendId sets BulkSendId field to given value.


### GetRequestId

`func (o *PaymentBulkSendItemUpdateEventData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentBulkSendItemUpdateEventData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentBulkSendItemUpdateEventData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### SetRequestIdNil

`func (o *PaymentBulkSendItemUpdateEventData) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *PaymentBulkSendItemUpdateEventData) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetStatus

`func (o *PaymentBulkSendItemUpdateEventData) GetStatus() PaymentBulkSendItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentBulkSendItemUpdateEventData) GetStatusOk() (*PaymentBulkSendItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentBulkSendItemUpdateEventData) SetStatus(v PaymentBulkSendItemStatus)`

SetStatus sets Status field to given value.


### GetRequestedAmount

`func (o *PaymentBulkSendItemUpdateEventData) GetRequestedAmount() PaymentAssetAmount`

GetRequestedAmount returns the RequestedAmount field if non-nil, zero value otherwise.

### GetRequestedAmountOk

`func (o *PaymentBulkSendItemUpdateEventData) GetRequestedAmountOk() (*PaymentAssetAmount, bool)`

GetRequestedAmountOk returns a tuple with the RequestedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAmount

`func (o *PaymentBulkSendItemUpdateEventData) SetRequestedAmount(v PaymentAssetAmount)`

SetRequestedAmount sets RequestedAmount field to given value.


### GetReceivingAddress

`func (o *PaymentBulkSendItemUpdateEventData) GetReceivingAddress() string`

GetReceivingAddress returns the ReceivingAddress field if non-nil, zero value otherwise.

### GetReceivingAddressOk

`func (o *PaymentBulkSendItemUpdateEventData) GetReceivingAddressOk() (*string, bool)`

GetReceivingAddressOk returns a tuple with the ReceivingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAddress

`func (o *PaymentBulkSendItemUpdateEventData) SetReceivingAddress(v string)`

SetReceivingAddress sets ReceivingAddress field to given value.


### GetTxHash

`func (o *PaymentBulkSendItemUpdateEventData) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *PaymentBulkSendItemUpdateEventData) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *PaymentBulkSendItemUpdateEventData) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### SetTxHashNil

`func (o *PaymentBulkSendItemUpdateEventData) SetTxHashNil(b bool)`

 SetTxHashNil sets the value for TxHash to be an explicit nil

### UnsetTxHash
`func (o *PaymentBulkSendItemUpdateEventData) UnsetTxHash()`

UnsetTxHash ensures that no value is present for TxHash, not even an explicit nil
### GetSourceAccount

`func (o *PaymentBulkSendItemUpdateEventData) GetSourceAccount() string`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *PaymentBulkSendItemUpdateEventData) GetSourceAccountOk() (*string, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *PaymentBulkSendItemUpdateEventData) SetSourceAccount(v string)`

SetSourceAccount sets SourceAccount field to given value.


### GetDescription

`func (o *PaymentBulkSendItemUpdateEventData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentBulkSendItemUpdateEventData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentBulkSendItemUpdateEventData) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *PaymentBulkSendItemUpdateEventData) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentBulkSendItemUpdateEventData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFailedReason

`func (o *PaymentBulkSendItemUpdateEventData) GetFailedReason() PaymentFailedReason`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *PaymentBulkSendItemUpdateEventData) GetFailedReasonOk() (*PaymentFailedReason, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *PaymentBulkSendItemUpdateEventData) SetFailedReason(v PaymentFailedReason)`

SetFailedReason sets FailedReason field to given value.


### SetFailedReasonNil

`func (o *PaymentBulkSendItemUpdateEventData) SetFailedReasonNil(b bool)`

 SetFailedReasonNil sets the value for FailedReason to be an explicit nil

### UnsetFailedReason
`func (o *PaymentBulkSendItemUpdateEventData) UnsetFailedReason()`

UnsetFailedReason ensures that no value is present for FailedReason, not even an explicit nil
### GetCreatedTimestamp

`func (o *PaymentBulkSendItemUpdateEventData) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentBulkSendItemUpdateEventData) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentBulkSendItemUpdateEventData) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *PaymentBulkSendItemUpdateEventData) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentBulkSendItemUpdateEventData) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentBulkSendItemUpdateEventData) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


