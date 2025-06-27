# TokenizationTokenPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionType** | [**TokenizationTokenPermissionType**](TokenizationTokenPermissionType.md) |  | 
**Name** | **string** | The display name of the permission. | 
**Description** | **string** | Detailed description of what this permission allows. | 
**Enabled** | **bool** | Whether this permission is currently enabled. | 

## Methods

### NewTokenizationTokenPermission

`func NewTokenizationTokenPermission(permissionType TokenizationTokenPermissionType, name string, description string, enabled bool, ) *TokenizationTokenPermission`

NewTokenizationTokenPermission instantiates a new TokenizationTokenPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationTokenPermissionWithDefaults

`func NewTokenizationTokenPermissionWithDefaults() *TokenizationTokenPermission`

NewTokenizationTokenPermissionWithDefaults instantiates a new TokenizationTokenPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionType

`func (o *TokenizationTokenPermission) GetPermissionType() TokenizationTokenPermissionType`

GetPermissionType returns the PermissionType field if non-nil, zero value otherwise.

### GetPermissionTypeOk

`func (o *TokenizationTokenPermission) GetPermissionTypeOk() (*TokenizationTokenPermissionType, bool)`

GetPermissionTypeOk returns a tuple with the PermissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionType

`func (o *TokenizationTokenPermission) SetPermissionType(v TokenizationTokenPermissionType)`

SetPermissionType sets PermissionType field to given value.


### GetName

`func (o *TokenizationTokenPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenizationTokenPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenizationTokenPermission) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TokenizationTokenPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenizationTokenPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenizationTokenPermission) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *TokenizationTokenPermission) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TokenizationTokenPermission) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TokenizationTokenPermission) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


