# PaymentBulkSendItemEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;PaymentBulkSendItem&#x60;: The payment bulk send item event data. - &#x60;PaymentAccountBalanceUpdate&#x60;: The Payments account balance updated event data, including account information and balance change details. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. - &#x60;Organization&#x60;: The organization event data. - &#x60;FiatTransaction&#x60;: The fiat transaction event data. | 
**BulkSendItemId** | **string** | The bulk send item ID. | 
**TokenId** | **string** | The token ID of the cryptocurrency to be sent to the recipient. | 
**ReceivingAddress** | **string** | The receiving address. | 
**Amount** | **string** | The amount of the cryptocurrency to be sent to the recipient. | 
**Description** | Pointer to **string** | A note or comment about the bulk send item. | [optional] 
**TxHash** | Pointer to **string** | The transaction hash of the bulk send item. | [optional] 
**Status** | [**PaymentBulkSendItemStatus**](PaymentBulkSendItemStatus.md) |  | 
**ValidationStatus** | [**PaymentBulkSendItemValidationStatus**](PaymentBulkSendItemValidationStatus.md) |  | 
**FailedReason** | Pointer to **string** | The reason why the bulk send item failed. | [optional] 
**BulkSendId** | **string** | The bulk send ID that this item belongs to. | 
**RequestId** | Pointer to **string** | The request ID of the bulk send batch. | [optional] 
**SourceAccount** | **string** | The source account ID of the bulk send batch. | 
**CreatedTimestamp** | **int32** | The created time of the bulk send item, represented as a UNIX timestamp in seconds. | 
**UpdatedTimestamp** | **int32** | The updated time of the bulk send item, represented as a UNIX timestamp in seconds. | 

## Methods

### NewPaymentBulkSendItemEvent

`func NewPaymentBulkSendItemEvent(dataType string, bulkSendItemId string, tokenId string, receivingAddress string, amount string, status PaymentBulkSendItemStatus, validationStatus PaymentBulkSendItemValidationStatus, bulkSendId string, sourceAccount string, createdTimestamp int32, updatedTimestamp int32, ) *PaymentBulkSendItemEvent`

NewPaymentBulkSendItemEvent instantiates a new PaymentBulkSendItemEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentBulkSendItemEventWithDefaults

`func NewPaymentBulkSendItemEventWithDefaults() *PaymentBulkSendItemEvent`

NewPaymentBulkSendItemEventWithDefaults instantiates a new PaymentBulkSendItemEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *PaymentBulkSendItemEvent) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PaymentBulkSendItemEvent) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PaymentBulkSendItemEvent) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetBulkSendItemId

`func (o *PaymentBulkSendItemEvent) GetBulkSendItemId() string`

GetBulkSendItemId returns the BulkSendItemId field if non-nil, zero value otherwise.

### GetBulkSendItemIdOk

`func (o *PaymentBulkSendItemEvent) GetBulkSendItemIdOk() (*string, bool)`

GetBulkSendItemIdOk returns a tuple with the BulkSendItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkSendItemId

`func (o *PaymentBulkSendItemEvent) SetBulkSendItemId(v string)`

SetBulkSendItemId sets BulkSendItemId field to given value.


### GetTokenId

`func (o *PaymentBulkSendItemEvent) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentBulkSendItemEvent) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentBulkSendItemEvent) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetReceivingAddress

`func (o *PaymentBulkSendItemEvent) GetReceivingAddress() string`

GetReceivingAddress returns the ReceivingAddress field if non-nil, zero value otherwise.

### GetReceivingAddressOk

`func (o *PaymentBulkSendItemEvent) GetReceivingAddressOk() (*string, bool)`

GetReceivingAddressOk returns a tuple with the ReceivingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAddress

`func (o *PaymentBulkSendItemEvent) SetReceivingAddress(v string)`

SetReceivingAddress sets ReceivingAddress field to given value.


### GetAmount

`func (o *PaymentBulkSendItemEvent) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentBulkSendItemEvent) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentBulkSendItemEvent) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *PaymentBulkSendItemEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentBulkSendItemEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentBulkSendItemEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentBulkSendItemEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTxHash

`func (o *PaymentBulkSendItemEvent) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *PaymentBulkSendItemEvent) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *PaymentBulkSendItemEvent) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *PaymentBulkSendItemEvent) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentBulkSendItemEvent) GetStatus() PaymentBulkSendItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentBulkSendItemEvent) GetStatusOk() (*PaymentBulkSendItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentBulkSendItemEvent) SetStatus(v PaymentBulkSendItemStatus)`

SetStatus sets Status field to given value.


### GetValidationStatus

`func (o *PaymentBulkSendItemEvent) GetValidationStatus() PaymentBulkSendItemValidationStatus`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *PaymentBulkSendItemEvent) GetValidationStatusOk() (*PaymentBulkSendItemValidationStatus, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *PaymentBulkSendItemEvent) SetValidationStatus(v PaymentBulkSendItemValidationStatus)`

SetValidationStatus sets ValidationStatus field to given value.


### GetFailedReason

`func (o *PaymentBulkSendItemEvent) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *PaymentBulkSendItemEvent) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *PaymentBulkSendItemEvent) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *PaymentBulkSendItemEvent) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.

### GetBulkSendId

`func (o *PaymentBulkSendItemEvent) GetBulkSendId() string`

GetBulkSendId returns the BulkSendId field if non-nil, zero value otherwise.

### GetBulkSendIdOk

`func (o *PaymentBulkSendItemEvent) GetBulkSendIdOk() (*string, bool)`

GetBulkSendIdOk returns a tuple with the BulkSendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkSendId

`func (o *PaymentBulkSendItemEvent) SetBulkSendId(v string)`

SetBulkSendId sets BulkSendId field to given value.


### GetRequestId

`func (o *PaymentBulkSendItemEvent) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentBulkSendItemEvent) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentBulkSendItemEvent) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *PaymentBulkSendItemEvent) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetSourceAccount

`func (o *PaymentBulkSendItemEvent) GetSourceAccount() string`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *PaymentBulkSendItemEvent) GetSourceAccountOk() (*string, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *PaymentBulkSendItemEvent) SetSourceAccount(v string)`

SetSourceAccount sets SourceAccount field to given value.


### GetCreatedTimestamp

`func (o *PaymentBulkSendItemEvent) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentBulkSendItemEvent) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentBulkSendItemEvent) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *PaymentBulkSendItemEvent) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentBulkSendItemEvent) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentBulkSendItemEvent) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


