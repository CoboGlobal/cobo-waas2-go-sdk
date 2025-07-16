# TokenizationAddressPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionAddress** | **string** | The execution address. | 
**Permissions** | [**[]TokenizationTokenPermissionType**](TokenizationTokenPermissionType.md) | List of permissions granted to this address. | 
**CreatedTimestamp** | Pointer to **int64** | The time when the permission was created, in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewTokenizationAddressPermission

`func NewTokenizationAddressPermission(executionAddress string, permissions []TokenizationTokenPermissionType, ) *TokenizationAddressPermission`

NewTokenizationAddressPermission instantiates a new TokenizationAddressPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationAddressPermissionWithDefaults

`func NewTokenizationAddressPermissionWithDefaults() *TokenizationAddressPermission`

NewTokenizationAddressPermissionWithDefaults instantiates a new TokenizationAddressPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionAddress

`func (o *TokenizationAddressPermission) GetExecutionAddress() string`

GetExecutionAddress returns the ExecutionAddress field if non-nil, zero value otherwise.

### GetExecutionAddressOk

`func (o *TokenizationAddressPermission) GetExecutionAddressOk() (*string, bool)`

GetExecutionAddressOk returns a tuple with the ExecutionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionAddress

`func (o *TokenizationAddressPermission) SetExecutionAddress(v string)`

SetExecutionAddress sets ExecutionAddress field to given value.


### GetPermissions

`func (o *TokenizationAddressPermission) GetPermissions() []TokenizationTokenPermissionType`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TokenizationAddressPermission) GetPermissionsOk() (*[]TokenizationTokenPermissionType, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TokenizationAddressPermission) SetPermissions(v []TokenizationTokenPermissionType)`

SetPermissions sets Permissions field to given value.


### GetCreatedTimestamp

`func (o *TokenizationAddressPermission) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *TokenizationAddressPermission) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *TokenizationAddressPermission) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *TokenizationAddressPermission) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


