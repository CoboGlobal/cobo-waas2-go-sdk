# PolicyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | [**PolicyActionType**](PolicyActionType.md) |  | 
**Content** | Pointer to [**PolicyActionContent**](PolicyActionContent.md) |  | [optional] 

## Methods

### NewPolicyAction

`func NewPolicyAction(actionType PolicyActionType, ) *PolicyAction`

NewPolicyAction instantiates a new PolicyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyActionWithDefaults

`func NewPolicyActionWithDefaults() *PolicyAction`

NewPolicyActionWithDefaults instantiates a new PolicyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *PolicyAction) GetActionType() PolicyActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *PolicyAction) GetActionTypeOk() (*PolicyActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *PolicyAction) SetActionType(v PolicyActionType)`

SetActionType sets ActionType field to given value.


### GetContent

`func (o *PolicyAction) GetContent() PolicyActionContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PolicyAction) GetContentOk() (*PolicyActionContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PolicyAction) SetContent(v PolicyActionContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *PolicyAction) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


