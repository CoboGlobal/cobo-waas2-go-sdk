# AppWorkflowPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**PolicyAction**](PolicyAction.md) |  | 
**Conditions** | Pointer to [**[]PolicyCondition**](PolicyCondition.md) |  | [optional] 

## Methods

### NewAppWorkflowPolicy

`func NewAppWorkflowPolicy(action PolicyAction, ) *AppWorkflowPolicy`

NewAppWorkflowPolicy instantiates a new AppWorkflowPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWorkflowPolicyWithDefaults

`func NewAppWorkflowPolicyWithDefaults() *AppWorkflowPolicy`

NewAppWorkflowPolicyWithDefaults instantiates a new AppWorkflowPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AppWorkflowPolicy) GetAction() PolicyAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AppWorkflowPolicy) GetActionOk() (*PolicyAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AppWorkflowPolicy) SetAction(v PolicyAction)`

SetAction sets Action field to given value.


### GetConditions

`func (o *AppWorkflowPolicy) GetConditions() []PolicyCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AppWorkflowPolicy) GetConditionsOk() (*[]PolicyCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AppWorkflowPolicy) SetConditions(v []PolicyCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AppWorkflowPolicy) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


