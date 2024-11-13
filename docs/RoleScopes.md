# RoleScopes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleName** | **string** | The user role associated with this API key. | 
**Scopes** | [**Scopes**](Scopes.md) |  | 

## Methods

### NewRoleScopes

`func NewRoleScopes(roleName string, scopes Scopes, ) *RoleScopes`

NewRoleScopes instantiates a new RoleScopes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleScopesWithDefaults

`func NewRoleScopesWithDefaults() *RoleScopes`

NewRoleScopesWithDefaults instantiates a new RoleScopes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleName

`func (o *RoleScopes) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *RoleScopes) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *RoleScopes) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetScopes

`func (o *RoleScopes) GetScopes() Scopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *RoleScopes) GetScopesOk() (*Scopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *RoleScopes) SetScopes(v Scopes)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


