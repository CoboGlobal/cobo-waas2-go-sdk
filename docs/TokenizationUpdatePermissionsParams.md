# TokenizationUpdatePermissionsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Addresses** | [**[]TokenizationUpdateAddressPermissions**](TokenizationUpdateAddressPermissions.md) |  | 

## Methods

### NewTokenizationUpdatePermissionsParams

`func NewTokenizationUpdatePermissionsParams(source TokenizationTokenOperationSource, addresses []TokenizationUpdateAddressPermissions, ) *TokenizationUpdatePermissionsParams`

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


### GetAddresses

`func (o *TokenizationUpdatePermissionsParams) GetAddresses() []TokenizationUpdateAddressPermissions`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *TokenizationUpdatePermissionsParams) GetAddressesOk() (*[]TokenizationUpdateAddressPermissions, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *TokenizationUpdatePermissionsParams) SetAddresses(v []TokenizationUpdateAddressPermissions)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


