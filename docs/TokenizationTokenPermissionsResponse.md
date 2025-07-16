# TokenizationTokenPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | [**[]TokenizationTokenPermission**](TokenizationTokenPermission.md) | List of available token permissions. | 
**TotalCount** | **int32** | Total number of permissions. | 

## Methods

### NewTokenizationTokenPermissionsResponse

`func NewTokenizationTokenPermissionsResponse(permissions []TokenizationTokenPermission, totalCount int32, ) *TokenizationTokenPermissionsResponse`

NewTokenizationTokenPermissionsResponse instantiates a new TokenizationTokenPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationTokenPermissionsResponseWithDefaults

`func NewTokenizationTokenPermissionsResponseWithDefaults() *TokenizationTokenPermissionsResponse`

NewTokenizationTokenPermissionsResponseWithDefaults instantiates a new TokenizationTokenPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *TokenizationTokenPermissionsResponse) GetPermissions() []TokenizationTokenPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TokenizationTokenPermissionsResponse) GetPermissionsOk() (*[]TokenizationTokenPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TokenizationTokenPermissionsResponse) SetPermissions(v []TokenizationTokenPermission)`

SetPermissions sets Permissions field to given value.


### GetTotalCount

`func (o *TokenizationTokenPermissionsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TokenizationTokenPermissionsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TokenizationTokenPermissionsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


