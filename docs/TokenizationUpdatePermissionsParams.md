# TokenizationUpdatePermissionsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Action** | [**TokenizationPermissionAction**](TokenizationPermissionAction.md) |  | 
**Address** | **string** | The address to manage permissions for. | 
**Permissions** | [**[]TokenizationTokenPermissionType**](TokenizationTokenPermissionType.md) | The list of permissions to operate on. | 

## Methods

### NewTokenizationUpdatePermissionsParams

`func NewTokenizationUpdatePermissionsParams(source TokenizationTokenOperationSource, action TokenizationPermissionAction, address string, permissions []TokenizationTokenPermissionType, ) *TokenizationUpdatePermissionsParams`

NewTokenizationUpdatePermissionsParams instantiates a new TokenizationUpdatePermissionsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationUpdatePermissionsParamsWithDefaults

`func NewTokenizationUpdatePermissionsParamsWithDefaults() *TokenizationUpdatePermissionsParams`

NewTokenizationUpdatePermissionsParamsWithDefaults instantiates a new TokenizationUpdatePermissionsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationUpdatePermissionsParams) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationUpdatePermissionsParams) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationUpdatePermissionsParams) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetAction

`func (o *TokenizationUpdatePermissionsParams) GetAction() TokenizationPermissionAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TokenizationUpdatePermissionsParams) GetActionOk() (*TokenizationPermissionAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TokenizationUpdatePermissionsParams) SetAction(v TokenizationPermissionAction)`

SetAction sets Action field to given value.


### GetAddress

`func (o *TokenizationUpdatePermissionsParams) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenizationUpdatePermissionsParams) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenizationUpdatePermissionsParams) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPermissions

`func (o *TokenizationUpdatePermissionsParams) GetPermissions() []TokenizationTokenPermissionType`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TokenizationUpdatePermissionsParams) GetPermissionsOk() (*[]TokenizationTokenPermissionType, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TokenizationUpdatePermissionsParams) SetPermissions(v []TokenizationTokenPermissionType)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


