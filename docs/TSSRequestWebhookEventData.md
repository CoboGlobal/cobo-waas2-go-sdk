# TSSRequestWebhookEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. - &#x60;PaymentTransaction&#x60;: The payment transaction event data. - &#x60;PaymentAddressUpdate&#x60;: The top-up address update event data. - &#x60;PaymentPayout&#x60;: The payment payout event data. - &#x60;PaymentBulkSend&#x60;: The payment bulk send event data. - &#x60;BalanceUpdateInfo&#x60;: The balance update event data. - &#x60;SuspendedToken&#x60;: The token suspension event data. - &#x60;ComplianceDisposition&#x60;: The compliance disposition event data. - &#x60;ComplianceKytScreenings&#x60;: The compliance KYT screenings event data. - &#x60;ComplianceKyaScreenings&#x60;: The compliance KYA screenings event data. | 
**TssRequestId** | Pointer to **string** | The TSS request ID. | [optional] 
**SourceKeyShareHolderGroup** | Pointer to [**SourceGroup**](SourceGroup.md) |  | [optional] 
**TargetKeyShareHolderGroupId** | Pointer to **string** | The target key share holder group ID. | [optional] 
**Type** | Pointer to [**TSSRequestType**](TSSRequestType.md) |  | [optional] 
**Status** | Pointer to [**TSSRequestStatus**](TSSRequestStatus.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the TSS request. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The TSS request&#39;s creation time in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewTSSRequestWebhookEventData

`func NewTSSRequestWebhookEventData(dataType string, ) *TSSRequestWebhookEventData`

NewTSSRequestWebhookEventData instantiates a new TSSRequestWebhookEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSRequestWebhookEventDataWithDefaults

`func NewTSSRequestWebhookEventDataWithDefaults() *TSSRequestWebhookEventData`

NewTSSRequestWebhookEventDataWithDefaults instantiates a new TSSRequestWebhookEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *TSSRequestWebhookEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *TSSRequestWebhookEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *TSSRequestWebhookEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetTssRequestId

`func (o *TSSRequestWebhookEventData) GetTssRequestId() string`

GetTssRequestId returns the TssRequestId field if non-nil, zero value otherwise.

### GetTssRequestIdOk

`func (o *TSSRequestWebhookEventData) GetTssRequestIdOk() (*string, bool)`

GetTssRequestIdOk returns a tuple with the TssRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTssRequestId

`func (o *TSSRequestWebhookEventData) SetTssRequestId(v string)`

SetTssRequestId sets TssRequestId field to given value.

### HasTssRequestId

`func (o *TSSRequestWebhookEventData) HasTssRequestId() bool`

HasTssRequestId returns a boolean if a field has been set.

### GetSourceKeyShareHolderGroup

`func (o *TSSRequestWebhookEventData) GetSourceKeyShareHolderGroup() SourceGroup`

GetSourceKeyShareHolderGroup returns the SourceKeyShareHolderGroup field if non-nil, zero value otherwise.

### GetSourceKeyShareHolderGroupOk

`func (o *TSSRequestWebhookEventData) GetSourceKeyShareHolderGroupOk() (*SourceGroup, bool)`

GetSourceKeyShareHolderGroupOk returns a tuple with the SourceKeyShareHolderGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKeyShareHolderGroup

`func (o *TSSRequestWebhookEventData) SetSourceKeyShareHolderGroup(v SourceGroup)`

SetSourceKeyShareHolderGroup sets SourceKeyShareHolderGroup field to given value.

### HasSourceKeyShareHolderGroup

`func (o *TSSRequestWebhookEventData) HasSourceKeyShareHolderGroup() bool`

HasSourceKeyShareHolderGroup returns a boolean if a field has been set.

### GetTargetKeyShareHolderGroupId

`func (o *TSSRequestWebhookEventData) GetTargetKeyShareHolderGroupId() string`

GetTargetKeyShareHolderGroupId returns the TargetKeyShareHolderGroupId field if non-nil, zero value otherwise.

### GetTargetKeyShareHolderGroupIdOk

`func (o *TSSRequestWebhookEventData) GetTargetKeyShareHolderGroupIdOk() (*string, bool)`

GetTargetKeyShareHolderGroupIdOk returns a tuple with the TargetKeyShareHolderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKeyShareHolderGroupId

`func (o *TSSRequestWebhookEventData) SetTargetKeyShareHolderGroupId(v string)`

SetTargetKeyShareHolderGroupId sets TargetKeyShareHolderGroupId field to given value.

### HasTargetKeyShareHolderGroupId

`func (o *TSSRequestWebhookEventData) HasTargetKeyShareHolderGroupId() bool`

HasTargetKeyShareHolderGroupId returns a boolean if a field has been set.

### GetType

`func (o *TSSRequestWebhookEventData) GetType() TSSRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TSSRequestWebhookEventData) GetTypeOk() (*TSSRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TSSRequestWebhookEventData) SetType(v TSSRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *TSSRequestWebhookEventData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *TSSRequestWebhookEventData) GetStatus() TSSRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TSSRequestWebhookEventData) GetStatusOk() (*TSSRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TSSRequestWebhookEventData) SetStatus(v TSSRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TSSRequestWebhookEventData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *TSSRequestWebhookEventData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TSSRequestWebhookEventData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TSSRequestWebhookEventData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TSSRequestWebhookEventData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *TSSRequestWebhookEventData) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *TSSRequestWebhookEventData) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *TSSRequestWebhookEventData) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *TSSRequestWebhookEventData) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


