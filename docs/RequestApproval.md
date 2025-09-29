# RequestApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | **string** | The operation id of workflow approval request. | 
**RequestId** | **string** | The request id of workflow approval request. | 
**InitiatorEmail** | **string** | The initiator email of workflow approval request. | 
**Fields** | [**[]AppWorkflowField**](AppWorkflowField.md) |  | 
**GuardTemplate** | **string** | The guard template content of workflow approval request, need to connect cobo. | 

## Methods

### NewRequestApproval

`func NewRequestApproval(operationId string, requestId string, initiatorEmail string, fields []AppWorkflowField, guardTemplate string, ) *RequestApproval`

NewRequestApproval instantiates a new RequestApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestApprovalWithDefaults

`func NewRequestApprovalWithDefaults() *RequestApproval`

NewRequestApprovalWithDefaults instantiates a new RequestApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationId

`func (o *RequestApproval) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *RequestApproval) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *RequestApproval) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.


### GetRequestId

`func (o *RequestApproval) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RequestApproval) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RequestApproval) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetInitiatorEmail

`func (o *RequestApproval) GetInitiatorEmail() string`

GetInitiatorEmail returns the InitiatorEmail field if non-nil, zero value otherwise.

### GetInitiatorEmailOk

`func (o *RequestApproval) GetInitiatorEmailOk() (*string, bool)`

GetInitiatorEmailOk returns a tuple with the InitiatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorEmail

`func (o *RequestApproval) SetInitiatorEmail(v string)`

SetInitiatorEmail sets InitiatorEmail field to given value.


### GetFields

`func (o *RequestApproval) GetFields() []AppWorkflowField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RequestApproval) GetFieldsOk() (*[]AppWorkflowField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RequestApproval) SetFields(v []AppWorkflowField)`

SetFields sets Fields field to given value.


### GetGuardTemplate

`func (o *RequestApproval) GetGuardTemplate() string`

GetGuardTemplate returns the GuardTemplate field if non-nil, zero value otherwise.

### GetGuardTemplateOk

`func (o *RequestApproval) GetGuardTemplateOk() (*string, bool)`

GetGuardTemplateOk returns a tuple with the GuardTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardTemplate

`func (o *RequestApproval) SetGuardTemplate(v string)`

SetGuardTemplate sets GuardTemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


