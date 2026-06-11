# PaymentBulkSend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkSendId** | **string** | The bulk send ID. | 
**RequestId** | Pointer to **string** | The request ID. | [optional] 
**SourceAccount** | **string** | The source account from which the bulk send will be made. - If the source account is a merchant account, provide the merchant&#39;s ID (e.g., \&quot;M1001\&quot;). - If the source account is the developer account, use the string &#x60;\&quot;developer\&quot;&#x60;.  | 
**Description** | Pointer to **string** | The description for the entire bulk send batch. | [optional] 
**ExecutionMode** | [**PaymentBulkSendExecutionMode**](PaymentBulkSendExecutionMode.md) |  | 
**Status** | [**PaymentBulkSendStatus**](PaymentBulkSendStatus.md) |  | 
**CreatedTimestamp** | **int32** | The created time of the bulk send, represented as a UNIX timestamp in seconds. | 
**UpdatedTimestamp** | **int32** | The updated time of the bulk send, represented as a UNIX timestamp in seconds. | 
**CommissionFee** | Pointer to [**CommissionFee**](CommissionFee.md) | The commission fee. Not returned when no fee has been incurred, the actual charged amount once incurred, or &#x60;0&#x60; if refunded. | [optional] 

## Methods

### NewPaymentBulkSend

`func NewPaymentBulkSend(bulkSendId string, sourceAccount string, executionMode PaymentBulkSendExecutionMode, status PaymentBulkSendStatus, createdTimestamp int32, updatedTimestamp int32, ) *PaymentBulkSend`

NewPaymentBulkSend instantiates a new PaymentBulkSend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentBulkSendWithDefaults

`func NewPaymentBulkSendWithDefaults() *PaymentBulkSend`

NewPaymentBulkSendWithDefaults instantiates a new PaymentBulkSend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkSendId

`func (o *PaymentBulkSend) GetBulkSendId() string`

GetBulkSendId returns the BulkSendId field if non-nil, zero value otherwise.

### GetBulkSendIdOk

`func (o *PaymentBulkSend) GetBulkSendIdOk() (*string, bool)`

GetBulkSendIdOk returns a tuple with the BulkSendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkSendId

`func (o *PaymentBulkSend) SetBulkSendId(v string)`

SetBulkSendId sets BulkSendId field to given value.


### GetRequestId

`func (o *PaymentBulkSend) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentBulkSend) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentBulkSend) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *PaymentBulkSend) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetSourceAccount

`func (o *PaymentBulkSend) GetSourceAccount() string`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *PaymentBulkSend) GetSourceAccountOk() (*string, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *PaymentBulkSend) SetSourceAccount(v string)`

SetSourceAccount sets SourceAccount field to given value.


### GetDescription

`func (o *PaymentBulkSend) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentBulkSend) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentBulkSend) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentBulkSend) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionMode

`func (o *PaymentBulkSend) GetExecutionMode() PaymentBulkSendExecutionMode`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *PaymentBulkSend) GetExecutionModeOk() (*PaymentBulkSendExecutionMode, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *PaymentBulkSend) SetExecutionMode(v PaymentBulkSendExecutionMode)`

SetExecutionMode sets ExecutionMode field to given value.


### GetStatus

`func (o *PaymentBulkSend) GetStatus() PaymentBulkSendStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentBulkSend) GetStatusOk() (*PaymentBulkSendStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentBulkSend) SetStatus(v PaymentBulkSendStatus)`

SetStatus sets Status field to given value.


### GetCreatedTimestamp

`func (o *PaymentBulkSend) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentBulkSend) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentBulkSend) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *PaymentBulkSend) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentBulkSend) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentBulkSend) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetCommissionFee

`func (o *PaymentBulkSend) GetCommissionFee() CommissionFee`

GetCommissionFee returns the CommissionFee field if non-nil, zero value otherwise.

### GetCommissionFeeOk

`func (o *PaymentBulkSend) GetCommissionFeeOk() (*CommissionFee, bool)`

GetCommissionFeeOk returns a tuple with the CommissionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionFee

`func (o *PaymentBulkSend) SetCommissionFee(v CommissionFee)`

SetCommissionFee sets CommissionFee field to given value.

### HasCommissionFee

`func (o *PaymentBulkSend) HasCommissionFee() bool`

HasCommissionFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


