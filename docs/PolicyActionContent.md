# PolicyActionContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The quorum type. Possible values include:    - &#x60;FULL_APPROVAL&#x60;: Requires approval from all participants.   - &#x60;PART_APPROVAL&#x60;: Requires approval from a specified number of participants.  | 
**Roles** | Pointer to **[]string** | The roles included in the quorum. Possible values include &#x60;admin&#x60;, &#x60;spender&#x60;, &#x60;operator&#x60;, and &#x60;approver&#x60;. | [optional] 
**UserIds** | Pointer to **[]string** | The ID of the users included in the quorum. | [optional] 
**Threshold** | Pointer to **int32** | The number of approvers required to meet the quorum. | [optional] 

## Methods

### NewPolicyActionContent

`func NewPolicyActionContent(type_ string, ) *PolicyActionContent`

NewPolicyActionContent instantiates a new PolicyActionContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyActionContentWithDefaults

`func NewPolicyActionContentWithDefaults() *PolicyActionContent`

NewPolicyActionContentWithDefaults instantiates a new PolicyActionContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PolicyActionContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyActionContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyActionContent) SetType(v string)`

SetType sets Type field to given value.


### GetRoles

`func (o *PolicyActionContent) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PolicyActionContent) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PolicyActionContent) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PolicyActionContent) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUserIds

`func (o *PolicyActionContent) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *PolicyActionContent) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *PolicyActionContent) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *PolicyActionContent) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetThreshold

`func (o *PolicyActionContent) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *PolicyActionContent) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *PolicyActionContent) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *PolicyActionContent) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


