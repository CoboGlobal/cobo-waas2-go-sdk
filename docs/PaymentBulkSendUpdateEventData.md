# PaymentBulkSendUpdateEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBankWithdrawal&#x60;: The payment bank withdrawal event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;PaymentBulkSendItem&#x60;: The payment bulk send item event data. - &#x60;PaymentTopUpUpdate&#x60;: The payment top-up update event data. - &#x60;PaymentOrderUpdate&#x60;: The payment order update event data. - &#x60;PaymentRefundUpdate&#x60;: The payment refund update event data. - &#x60;PaymentPayoutUpdate&#x60;: The payment payout update event data. - &#x60;PaymentBulkSendUpdate&#x60;: The payment bulk send update event data. - &#x60;PaymentBulkSendItemUpdate&#x60;: The payment bulk send item update event data. - &#x60;PaymentAccountBalanceUpdate&#x60;: The Payments account balance updated event data, including account information and balance change details. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. - &#x60;Organization&#x60;: The organization event data. - &#x60;FiatTransaction&#x60;: The fiat transaction event data. | 
**BulkSendId** | **string** | The bulk send ID. | 
**RequestId** | **NullableString** | The request ID. | 
**SourceAccount** | **string** | The Payments account from which the funds are debited: the merchant ID for a merchant account, or &#x60;developer&#x60; for the developer account. | 
**ExecutionMode** | [**PaymentBulkSendExecutionMode**](PaymentBulkSendExecutionMode.md) |  | 
**Status** | [**PaymentBulkSendStatus**](PaymentBulkSendStatus.md) |  | 
**ItemCount** | **int32** | The total number of items in the bulk send. | 
**FailedItemCount** | **int32** | The number of failed items. | 
**NotExecutedItemCount** | **int32** | The number of items not executed. | 
**RequestedAmounts** | [**[]PaymentAssetAmount**](PaymentAssetAmount.md) | The requested amounts aggregated by asset. | 
**CompletedAmounts** | [**[]PaymentAssetAmount**](PaymentAssetAmount.md) | The amounts of completed items aggregated by asset. | 
**FailedAmounts** | [**[]PaymentAssetAmount**](PaymentAssetAmount.md) | The amounts of failed and not-executed items aggregated by asset. | 
**Description** | **NullableString** | The description for the entire bulk send batch. | 
**CommissionFee** | [**NullablePaymentCommissionFeeTotal**](PaymentCommissionFeeTotal.md) |  | 
**FailedReason** | [**NullablePaymentFailedReason**](PaymentFailedReason.md) |  | 
**CreatedTimestamp** | **int64** | The created time of the bulk send, represented as a UNIX timestamp in seconds. | 
**UpdatedTimestamp** | **int64** | The updated time of the bulk send, represented as a UNIX timestamp in seconds. | 

## Methods

### NewPaymentBulkSendUpdateEventData

`func NewPaymentBulkSendUpdateEventData(dataType string, bulkSendId string, requestId NullableString, sourceAccount string, executionMode PaymentBulkSendExecutionMode, status PaymentBulkSendStatus, itemCount int32, failedItemCount int32, notExecutedItemCount int32, requestedAmounts []PaymentAssetAmount, completedAmounts []PaymentAssetAmount, failedAmounts []PaymentAssetAmount, description NullableString, commissionFee NullablePaymentCommissionFeeTotal, failedReason NullablePaymentFailedReason, createdTimestamp int64, updatedTimestamp int64, ) *PaymentBulkSendUpdateEventData`

NewPaymentBulkSendUpdateEventData instantiates a new PaymentBulkSendUpdateEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentBulkSendUpdateEventDataWithDefaults

`func NewPaymentBulkSendUpdateEventDataWithDefaults() *PaymentBulkSendUpdateEventData`

NewPaymentBulkSendUpdateEventDataWithDefaults instantiates a new PaymentBulkSendUpdateEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *PaymentBulkSendUpdateEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PaymentBulkSendUpdateEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PaymentBulkSendUpdateEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetBulkSendId

`func (o *PaymentBulkSendUpdateEventData) GetBulkSendId() string`

GetBulkSendId returns the BulkSendId field if non-nil, zero value otherwise.

### GetBulkSendIdOk

`func (o *PaymentBulkSendUpdateEventData) GetBulkSendIdOk() (*string, bool)`

GetBulkSendIdOk returns a tuple with the BulkSendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkSendId

`func (o *PaymentBulkSendUpdateEventData) SetBulkSendId(v string)`

SetBulkSendId sets BulkSendId field to given value.


### GetRequestId

`func (o *PaymentBulkSendUpdateEventData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentBulkSendUpdateEventData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentBulkSendUpdateEventData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### SetRequestIdNil

`func (o *PaymentBulkSendUpdateEventData) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *PaymentBulkSendUpdateEventData) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetSourceAccount

`func (o *PaymentBulkSendUpdateEventData) GetSourceAccount() string`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *PaymentBulkSendUpdateEventData) GetSourceAccountOk() (*string, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *PaymentBulkSendUpdateEventData) SetSourceAccount(v string)`

SetSourceAccount sets SourceAccount field to given value.


### GetExecutionMode

`func (o *PaymentBulkSendUpdateEventData) GetExecutionMode() PaymentBulkSendExecutionMode`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *PaymentBulkSendUpdateEventData) GetExecutionModeOk() (*PaymentBulkSendExecutionMode, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *PaymentBulkSendUpdateEventData) SetExecutionMode(v PaymentBulkSendExecutionMode)`

SetExecutionMode sets ExecutionMode field to given value.


### GetStatus

`func (o *PaymentBulkSendUpdateEventData) GetStatus() PaymentBulkSendStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentBulkSendUpdateEventData) GetStatusOk() (*PaymentBulkSendStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentBulkSendUpdateEventData) SetStatus(v PaymentBulkSendStatus)`

SetStatus sets Status field to given value.


### GetItemCount

`func (o *PaymentBulkSendUpdateEventData) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *PaymentBulkSendUpdateEventData) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *PaymentBulkSendUpdateEventData) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.


### GetFailedItemCount

`func (o *PaymentBulkSendUpdateEventData) GetFailedItemCount() int32`

GetFailedItemCount returns the FailedItemCount field if non-nil, zero value otherwise.

### GetFailedItemCountOk

`func (o *PaymentBulkSendUpdateEventData) GetFailedItemCountOk() (*int32, bool)`

GetFailedItemCountOk returns a tuple with the FailedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedItemCount

`func (o *PaymentBulkSendUpdateEventData) SetFailedItemCount(v int32)`

SetFailedItemCount sets FailedItemCount field to given value.


### GetNotExecutedItemCount

`func (o *PaymentBulkSendUpdateEventData) GetNotExecutedItemCount() int32`

GetNotExecutedItemCount returns the NotExecutedItemCount field if non-nil, zero value otherwise.

### GetNotExecutedItemCountOk

`func (o *PaymentBulkSendUpdateEventData) GetNotExecutedItemCountOk() (*int32, bool)`

GetNotExecutedItemCountOk returns a tuple with the NotExecutedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotExecutedItemCount

`func (o *PaymentBulkSendUpdateEventData) SetNotExecutedItemCount(v int32)`

SetNotExecutedItemCount sets NotExecutedItemCount field to given value.


### GetRequestedAmounts

`func (o *PaymentBulkSendUpdateEventData) GetRequestedAmounts() []PaymentAssetAmount`

GetRequestedAmounts returns the RequestedAmounts field if non-nil, zero value otherwise.

### GetRequestedAmountsOk

`func (o *PaymentBulkSendUpdateEventData) GetRequestedAmountsOk() (*[]PaymentAssetAmount, bool)`

GetRequestedAmountsOk returns a tuple with the RequestedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAmounts

`func (o *PaymentBulkSendUpdateEventData) SetRequestedAmounts(v []PaymentAssetAmount)`

SetRequestedAmounts sets RequestedAmounts field to given value.


### GetCompletedAmounts

`func (o *PaymentBulkSendUpdateEventData) GetCompletedAmounts() []PaymentAssetAmount`

GetCompletedAmounts returns the CompletedAmounts field if non-nil, zero value otherwise.

### GetCompletedAmountsOk

`func (o *PaymentBulkSendUpdateEventData) GetCompletedAmountsOk() (*[]PaymentAssetAmount, bool)`

GetCompletedAmountsOk returns a tuple with the CompletedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAmounts

`func (o *PaymentBulkSendUpdateEventData) SetCompletedAmounts(v []PaymentAssetAmount)`

SetCompletedAmounts sets CompletedAmounts field to given value.


### GetFailedAmounts

`func (o *PaymentBulkSendUpdateEventData) GetFailedAmounts() []PaymentAssetAmount`

GetFailedAmounts returns the FailedAmounts field if non-nil, zero value otherwise.

### GetFailedAmountsOk

`func (o *PaymentBulkSendUpdateEventData) GetFailedAmountsOk() (*[]PaymentAssetAmount, bool)`

GetFailedAmountsOk returns a tuple with the FailedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAmounts

`func (o *PaymentBulkSendUpdateEventData) SetFailedAmounts(v []PaymentAssetAmount)`

SetFailedAmounts sets FailedAmounts field to given value.


### GetDescription

`func (o *PaymentBulkSendUpdateEventData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentBulkSendUpdateEventData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentBulkSendUpdateEventData) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *PaymentBulkSendUpdateEventData) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentBulkSendUpdateEventData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCommissionFee

`func (o *PaymentBulkSendUpdateEventData) GetCommissionFee() PaymentCommissionFeeTotal`

GetCommissionFee returns the CommissionFee field if non-nil, zero value otherwise.

### GetCommissionFeeOk

`func (o *PaymentBulkSendUpdateEventData) GetCommissionFeeOk() (*PaymentCommissionFeeTotal, bool)`

GetCommissionFeeOk returns a tuple with the CommissionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionFee

`func (o *PaymentBulkSendUpdateEventData) SetCommissionFee(v PaymentCommissionFeeTotal)`

SetCommissionFee sets CommissionFee field to given value.


### SetCommissionFeeNil

`func (o *PaymentBulkSendUpdateEventData) SetCommissionFeeNil(b bool)`

 SetCommissionFeeNil sets the value for CommissionFee to be an explicit nil

### UnsetCommissionFee
`func (o *PaymentBulkSendUpdateEventData) UnsetCommissionFee()`

UnsetCommissionFee ensures that no value is present for CommissionFee, not even an explicit nil
### GetFailedReason

`func (o *PaymentBulkSendUpdateEventData) GetFailedReason() PaymentFailedReason`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *PaymentBulkSendUpdateEventData) GetFailedReasonOk() (*PaymentFailedReason, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *PaymentBulkSendUpdateEventData) SetFailedReason(v PaymentFailedReason)`

SetFailedReason sets FailedReason field to given value.


### SetFailedReasonNil

`func (o *PaymentBulkSendUpdateEventData) SetFailedReasonNil(b bool)`

 SetFailedReasonNil sets the value for FailedReason to be an explicit nil

### UnsetFailedReason
`func (o *PaymentBulkSendUpdateEventData) UnsetFailedReason()`

UnsetFailedReason ensures that no value is present for FailedReason, not even an explicit nil
### GetCreatedTimestamp

`func (o *PaymentBulkSendUpdateEventData) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentBulkSendUpdateEventData) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentBulkSendUpdateEventData) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *PaymentBulkSendUpdateEventData) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentBulkSendUpdateEventData) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentBulkSendUpdateEventData) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


