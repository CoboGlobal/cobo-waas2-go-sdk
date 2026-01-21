# CreateBulkSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAccount** | **string** | The source account from which the bulk send will be made. - If the source account is a merchant account, provide the merchant&#39;s ID (e.g., \&quot;M1001\&quot;). - If the source account is the developer account, use the string &#x60;\&quot;developer\&quot;&#x60;.  | 
**ExecutionMode** | [**PaymentBulkSendExecutionMode**](PaymentBulkSendExecutionMode.md) |  | 
**Description** | Pointer to **string** | The description for the entire bulk send batch. | [optional] 
**PayoutParams** | [**[]CreateBulkSendRequestPayoutParamsInner**](CreateBulkSendRequestPayoutParamsInner.md) | The bulk send items. | 

## Methods

### NewCreateBulkSendRequest

`func NewCreateBulkSendRequest(sourceAccount string, executionMode PaymentBulkSendExecutionMode, payoutParams []CreateBulkSendRequestPayoutParamsInner, ) *CreateBulkSendRequest`

NewCreateBulkSendRequest instantiates a new CreateBulkSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBulkSendRequestWithDefaults

`func NewCreateBulkSendRequestWithDefaults() *CreateBulkSendRequest`

NewCreateBulkSendRequestWithDefaults instantiates a new CreateBulkSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAccount

`func (o *CreateBulkSendRequest) GetSourceAccount() string`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *CreateBulkSendRequest) GetSourceAccountOk() (*string, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *CreateBulkSendRequest) SetSourceAccount(v string)`

SetSourceAccount sets SourceAccount field to given value.


### GetExecutionMode

`func (o *CreateBulkSendRequest) GetExecutionMode() PaymentBulkSendExecutionMode`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *CreateBulkSendRequest) GetExecutionModeOk() (*PaymentBulkSendExecutionMode, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *CreateBulkSendRequest) SetExecutionMode(v PaymentBulkSendExecutionMode)`

SetExecutionMode sets ExecutionMode field to given value.


### GetDescription

`func (o *CreateBulkSendRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateBulkSendRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateBulkSendRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateBulkSendRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPayoutParams

`func (o *CreateBulkSendRequest) GetPayoutParams() []CreateBulkSendRequestPayoutParamsInner`

GetPayoutParams returns the PayoutParams field if non-nil, zero value otherwise.

### GetPayoutParamsOk

`func (o *CreateBulkSendRequest) GetPayoutParamsOk() (*[]CreateBulkSendRequestPayoutParamsInner, bool)`

GetPayoutParamsOk returns a tuple with the PayoutParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutParams

`func (o *CreateBulkSendRequest) SetPayoutParams(v []CreateBulkSendRequestPayoutParamsInner)`

SetPayoutParams sets PayoutParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


