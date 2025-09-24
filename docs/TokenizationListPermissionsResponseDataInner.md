# TokenizationListPermissionsResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address with permissions. | 
**Permissions** | [**[]TokenizationTokenPermissionType**](TokenizationTokenPermissionType.md) | The permissions assigned to this address. | 

## Methods

### NewTokenizationListPermissionsResponseDataInner

`func NewTokenizationListPermissionsResponseDataInner(address string, permissions []TokenizationTokenPermissionType, ) *TokenizationListPermissionsResponseDataInner`

NewTokenizationListPermissionsResponseDataInner instantiates a new TokenizationListPermissionsResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationListPermissionsResponseDataInnerWithDefaults

`func NewTokenizationListPermissionsResponseDataInnerWithDefaults() *TokenizationListPermissionsResponseDataInner`

NewTokenizationListPermissionsResponseDataInnerWithDefaults instantiates a new TokenizationListPermissionsResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TokenizationListPermissionsResponseDataInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenizationListPermissionsResponseDataInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenizationListPermissionsResponseDataInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPermissions

`func (o *TokenizationListPermissionsResponseDataInner) GetPermissions() []TokenizationTokenPermissionType`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TokenizationListPermissionsResponseDataInner) GetPermissionsOk() (*[]TokenizationTokenPermissionType, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TokenizationListPermissionsResponseDataInner) SetPermissions(v []TokenizationTokenPermissionType)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


