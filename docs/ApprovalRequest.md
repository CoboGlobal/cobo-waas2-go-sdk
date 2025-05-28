# ApprovalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalId** | **string** |  | 
**RequestId** | **string** |  | 
**Fields** | [**[]AppWorkflowField**](AppWorkflowField.md) |  | 
**Status** | [**ApprovalStatus**](ApprovalStatus.md) |  | 
**InitiatedTimestamp** | **int64** | The time when the approval was initiated, in Unix timestamp format, measured in milliseconds. | 

## Methods

### NewApprovalRequest

`func NewApprovalRequest(approvalId string, requestId string, fields []AppWorkflowField, status ApprovalStatus, initiatedTimestamp int64, ) *ApprovalRequest`

NewApprovalRequest instantiates a new ApprovalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestWithDefaults

`func NewApprovalRequestWithDefaults() *ApprovalRequest`

NewApprovalRequestWithDefaults instantiates a new ApprovalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalId

`func (o *ApprovalRequest) GetApprovalId() string`

GetApprovalId returns the ApprovalId field if non-nil, zero value otherwise.

### GetApprovalIdOk

`func (o *ApprovalRequest) GetApprovalIdOk() (*string, bool)`

GetApprovalIdOk returns a tuple with the ApprovalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalId

`func (o *ApprovalRequest) SetApprovalId(v string)`

SetApprovalId sets ApprovalId field to given value.


### GetRequestId

`func (o *ApprovalRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ApprovalRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ApprovalRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetFields

`func (o *ApprovalRequest) GetFields() []AppWorkflowField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ApprovalRequest) GetFieldsOk() (*[]AppWorkflowField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ApprovalRequest) SetFields(v []AppWorkflowField)`

SetFields sets Fields field to given value.


### GetStatus

`func (o *ApprovalRequest) GetStatus() ApprovalStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalRequest) GetStatusOk() (*ApprovalStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalRequest) SetStatus(v ApprovalStatus)`

SetStatus sets Status field to given value.


### GetInitiatedTimestamp

`func (o *ApprovalRequest) GetInitiatedTimestamp() int64`

GetInitiatedTimestamp returns the InitiatedTimestamp field if non-nil, zero value otherwise.

### GetInitiatedTimestampOk

`func (o *ApprovalRequest) GetInitiatedTimestampOk() (*int64, bool)`

GetInitiatedTimestampOk returns a tuple with the InitiatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedTimestamp

`func (o *ApprovalRequest) SetInitiatedTimestamp(v int64)`

SetInitiatedTimestamp sets InitiatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


