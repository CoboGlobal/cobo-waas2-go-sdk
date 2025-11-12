# TokenizationUpdateAddressPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address to manage permissions for. | 
**Action** | [**TokenizationPermissionAction**](TokenizationPermissionAction.md) |  | 
**Permissions** | [**[]TokenizationTokenPermissionType**](TokenizationTokenPermissionType.md) | The list of permissions to be applied. | 

## Methods

### NewTokenizationUpdateAddressPermissions

`func NewTokenizationUpdateAddressPermissions(address string, action TokenizationPermissionAction, permissions []TokenizationTokenPermissionType, ) *TokenizationUpdateAddressPermissions`

NewTokenizationUpdateAddressPermissions instantiates a new TokenizationUpdateAddressPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationUpdateAddressPermissionsWithDefaults

`func NewTokenizationUpdateAddressPermissionsWithDefaults() *TokenizationUpdateAddressPermissions`

NewTokenizationUpdateAddressPermissionsWithDefaults instantiates a new TokenizationUpdateAddressPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TokenizationUpdateAddressPermissions) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenizationUpdateAddressPermissions) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenizationUpdateAddressPermissions) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAction

`func (o *TokenizationUpdateAddressPermissions) GetAction() TokenizationPermissionAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TokenizationUpdateAddressPermissions) GetActionOk() (*TokenizationPermissionAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TokenizationUpdateAddressPermissions) SetAction(v TokenizationPermissionAction)`

SetAction sets Action field to given value.


### GetPermissions

`func (o *TokenizationUpdateAddressPermissions) GetPermissions() []TokenizationTokenPermissionType`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TokenizationUpdateAddressPermissions) GetPermissionsOk() (*[]TokenizationTokenPermissionType, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TokenizationUpdateAddressPermissions) SetPermissions(v []TokenizationTokenPermissionType)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


