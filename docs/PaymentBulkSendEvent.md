# PaymentBulkSendEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. | 
**BulkSendId** | **string** | The bulk send ID. | 
**SourceAccount** | **string** | The source account from which the bulk send will be made. - If the source account is a merchant account, provide the merchant&#39;s ID (e.g., \&quot;M1001\&quot;). - If the source account is the developer account, use the string &#x60;\&quot;developer\&quot;&#x60;.  | 
**Description** | Pointer to **string** | The description for the entire bulk send batch. | [optional] 
**ExecutionMode** | [**PaymentBulkSendExecutionMode**](PaymentBulkSendExecutionMode.md) |  | 
**Status** | [**PaymentBulkSendStatus**](PaymentBulkSendStatus.md) |  | 
**CreatedTimestamp** | **int32** | The created time of the bulk send, represented as a UNIX timestamp in seconds. | 
**UpdatedTimestamp** | **int32** | The updated time of the bulk send, represented as a UNIX timestamp in seconds. | 

## Methods

### NewPaymentBulkSendEvent

`func NewPaymentBulkSendEvent(dataType string, bulkSendId string, sourceAccount string, executionMode PaymentBulkSendExecutionMode, status PaymentBulkSendStatus, createdTimestamp int32, updatedTimestamp int32, ) *PaymentBulkSendEvent`

NewPaymentBulkSendEvent instantiates a new PaymentBulkSendEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentBulkSendEventWithDefaults

`func NewPaymentBulkSendEventWithDefaults() *PaymentBulkSendEvent`

NewPaymentBulkSendEventWithDefaults instantiates a new PaymentBulkSendEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *PaymentBulkSendEvent) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PaymentBulkSendEvent) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PaymentBulkSendEvent) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetBulkSendId

`func (o *PaymentBulkSendEvent) GetBulkSendId() string`

GetBulkSendId returns the BulkSendId field if non-nil, zero value otherwise.

### GetBulkSendIdOk

`func (o *PaymentBulkSendEvent) GetBulkSendIdOk() (*string, bool)`

GetBulkSendIdOk returns a tuple with the BulkSendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkSendId

`func (o *PaymentBulkSendEvent) SetBulkSendId(v string)`

SetBulkSendId sets BulkSendId field to given value.


### GetSourceAccount

`func (o *PaymentBulkSendEvent) GetSourceAccount() string`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *PaymentBulkSendEvent) GetSourceAccountOk() (*string, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *PaymentBulkSendEvent) SetSourceAccount(v string)`

SetSourceAccount sets SourceAccount field to given value.


### GetDescription

`func (o *PaymentBulkSendEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentBulkSendEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentBulkSendEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentBulkSendEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionMode

`func (o *PaymentBulkSendEvent) GetExecutionMode() PaymentBulkSendExecutionMode`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *PaymentBulkSendEvent) GetExecutionModeOk() (*PaymentBulkSendExecutionMode, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *PaymentBulkSendEvent) SetExecutionMode(v PaymentBulkSendExecutionMode)`

SetExecutionMode sets ExecutionMode field to given value.


### GetStatus

`func (o *PaymentBulkSendEvent) GetStatus() PaymentBulkSendStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentBulkSendEvent) GetStatusOk() (*PaymentBulkSendStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentBulkSendEvent) SetStatus(v PaymentBulkSendStatus)`

SetStatus sets Status field to given value.


### GetCreatedTimestamp

`func (o *PaymentBulkSendEvent) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentBulkSendEvent) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentBulkSendEvent) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *PaymentBulkSendEvent) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentBulkSendEvent) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentBulkSendEvent) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


