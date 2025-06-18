# AppWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The unique ID distinguishing the approval workflow instance among organizations. | 
**OperationId** | **string** | The unique ID of the approval workflow. | 
**OperationName** | **string** | The name of the approval workflow. | 
**CurrentPolicies** | [**[]AppWorkflowPolicy**](AppWorkflowPolicy.md) |  | 

## Methods

### NewAppWorkflow

`func NewAppWorkflow(workflowId string, operationId string, operationName string, currentPolicies []AppWorkflowPolicy, ) *AppWorkflow`

NewAppWorkflow instantiates a new AppWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWorkflowWithDefaults

`func NewAppWorkflowWithDefaults() *AppWorkflow`

NewAppWorkflowWithDefaults instantiates a new AppWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *AppWorkflow) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *AppWorkflow) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *AppWorkflow) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetOperationId

`func (o *AppWorkflow) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *AppWorkflow) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *AppWorkflow) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.


### GetOperationName

`func (o *AppWorkflow) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *AppWorkflow) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *AppWorkflow) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.


### GetCurrentPolicies

`func (o *AppWorkflow) GetCurrentPolicies() []AppWorkflowPolicy`

GetCurrentPolicies returns the CurrentPolicies field if non-nil, zero value otherwise.

### GetCurrentPoliciesOk

`func (o *AppWorkflow) GetCurrentPoliciesOk() (*[]AppWorkflowPolicy, bool)`

GetCurrentPoliciesOk returns a tuple with the CurrentPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPolicies

`func (o *AppWorkflow) SetCurrentPolicies(v []AppWorkflowPolicy)`

SetCurrentPolicies sets CurrentPolicies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


