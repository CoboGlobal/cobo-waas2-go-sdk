# TokenizationERC20WrappedTokenParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Standard** | [**TokenizationTokenStandard**](TokenizationTokenStandard.md) |  | 
**Name** | **string** | The name of the token. | 
**Symbol** | **string** | The symbol of the token. | 
**Decimals** | **int32** | The number of decimals for the token. | 
**Permissions** | Pointer to [**TokenizationERC20WrappedTokenPermissionParams**](TokenizationERC20WrappedTokenPermissionParams.md) |  | [optional] 
**TokenAccessActivated** | Pointer to **bool** | Whether the allowlist feature is activated for the token. When activated, only addresses in the allowlist can perform token operations. | [optional] [default to false]
**UnderlyingToken** | **string** | The address of the underlying token that this tokenized asset represents. | 

## Methods

### NewTokenizationERC20WrappedTokenParams

`func NewTokenizationERC20WrappedTokenParams(standard TokenizationTokenStandard, name string, symbol string, decimals int32, underlyingToken string, ) *TokenizationERC20WrappedTokenParams`

NewTokenizationERC20WrappedTokenParams instantiates a new TokenizationERC20WrappedTokenParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationERC20WrappedTokenParamsWithDefaults

`func NewTokenizationERC20WrappedTokenParamsWithDefaults() *TokenizationERC20WrappedTokenParams`

NewTokenizationERC20WrappedTokenParamsWithDefaults instantiates a new TokenizationERC20WrappedTokenParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandard

`func (o *TokenizationERC20WrappedTokenParams) GetStandard() TokenizationTokenStandard`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *TokenizationERC20WrappedTokenParams) GetStandardOk() (*TokenizationTokenStandard, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *TokenizationERC20WrappedTokenParams) SetStandard(v TokenizationTokenStandard)`

SetStandard sets Standard field to given value.


### GetName

`func (o *TokenizationERC20WrappedTokenParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenizationERC20WrappedTokenParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenizationERC20WrappedTokenParams) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *TokenizationERC20WrappedTokenParams) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenizationERC20WrappedTokenParams) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenizationERC20WrappedTokenParams) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimals

`func (o *TokenizationERC20WrappedTokenParams) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenizationERC20WrappedTokenParams) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenizationERC20WrappedTokenParams) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetPermissions

`func (o *TokenizationERC20WrappedTokenParams) GetPermissions() TokenizationERC20WrappedTokenPermissionParams`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TokenizationERC20WrappedTokenParams) GetPermissionsOk() (*TokenizationERC20WrappedTokenPermissionParams, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TokenizationERC20WrappedTokenParams) SetPermissions(v TokenizationERC20WrappedTokenPermissionParams)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *TokenizationERC20WrappedTokenParams) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTokenAccessActivated

`func (o *TokenizationERC20WrappedTokenParams) GetTokenAccessActivated() bool`

GetTokenAccessActivated returns the TokenAccessActivated field if non-nil, zero value otherwise.

### GetTokenAccessActivatedOk

`func (o *TokenizationERC20WrappedTokenParams) GetTokenAccessActivatedOk() (*bool, bool)`

GetTokenAccessActivatedOk returns a tuple with the TokenAccessActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAccessActivated

`func (o *TokenizationERC20WrappedTokenParams) SetTokenAccessActivated(v bool)`

SetTokenAccessActivated sets TokenAccessActivated field to given value.

### HasTokenAccessActivated

`func (o *TokenizationERC20WrappedTokenParams) HasTokenAccessActivated() bool`

HasTokenAccessActivated returns a boolean if a field has been set.

### GetUnderlyingToken

`func (o *TokenizationERC20WrappedTokenParams) GetUnderlyingToken() string`

GetUnderlyingToken returns the UnderlyingToken field if non-nil, zero value otherwise.

### GetUnderlyingTokenOk

`func (o *TokenizationERC20WrappedTokenParams) GetUnderlyingTokenOk() (*string, bool)`

GetUnderlyingTokenOk returns a tuple with the UnderlyingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingToken

`func (o *TokenizationERC20WrappedTokenParams) SetUnderlyingToken(v string)`

SetUnderlyingToken sets UnderlyingToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


