# TokenizationERC20TokenParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Standard** | [**TokenizationTokenStandard**](TokenizationTokenStandard.md) |  | 
**Name** | **string** | The name of the token. | 
**Symbol** | **string** | The symbol of the token. | 
**Decimals** | **int32** | The number of decimals for the token (0-18). | 
**AllowlistActivated** | Pointer to **bool** | Whether the allowlist feature is activated for the token. When activated, only addresses in the allowlist can perform token operations. | [optional] [default to false]
**Permissions** | Pointer to [**TokenizationTokenPermissionParams**](TokenizationTokenPermissionParams.md) |  | [optional] 

## Methods

### NewTokenizationERC20TokenParams

`func NewTokenizationERC20TokenParams(standard TokenizationTokenStandard, name string, symbol string, decimals int32, ) *TokenizationERC20TokenParams`

NewTokenizationERC20TokenParams instantiates a new TokenizationERC20TokenParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationERC20TokenParamsWithDefaults

`func NewTokenizationERC20TokenParamsWithDefaults() *TokenizationERC20TokenParams`

NewTokenizationERC20TokenParamsWithDefaults instantiates a new TokenizationERC20TokenParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandard

`func (o *TokenizationERC20TokenParams) GetStandard() TokenizationTokenStandard`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *TokenizationERC20TokenParams) GetStandardOk() (*TokenizationTokenStandard, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *TokenizationERC20TokenParams) SetStandard(v TokenizationTokenStandard)`

SetStandard sets Standard field to given value.


### GetName

`func (o *TokenizationERC20TokenParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenizationERC20TokenParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenizationERC20TokenParams) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *TokenizationERC20TokenParams) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenizationERC20TokenParams) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenizationERC20TokenParams) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimals

`func (o *TokenizationERC20TokenParams) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenizationERC20TokenParams) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenizationERC20TokenParams) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetAllowlistActivated

`func (o *TokenizationERC20TokenParams) GetAllowlistActivated() bool`

GetAllowlistActivated returns the AllowlistActivated field if non-nil, zero value otherwise.

### GetAllowlistActivatedOk

`func (o *TokenizationERC20TokenParams) GetAllowlistActivatedOk() (*bool, bool)`

GetAllowlistActivatedOk returns a tuple with the AllowlistActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistActivated

`func (o *TokenizationERC20TokenParams) SetAllowlistActivated(v bool)`

SetAllowlistActivated sets AllowlistActivated field to given value.

### HasAllowlistActivated

`func (o *TokenizationERC20TokenParams) HasAllowlistActivated() bool`

HasAllowlistActivated returns a boolean if a field has been set.

### GetPermissions

`func (o *TokenizationERC20TokenParams) GetPermissions() TokenizationTokenPermissionParams`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TokenizationERC20TokenParams) GetPermissionsOk() (*TokenizationTokenPermissionParams, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TokenizationERC20TokenParams) SetPermissions(v TokenizationTokenPermissionParams)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *TokenizationERC20TokenParams) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


