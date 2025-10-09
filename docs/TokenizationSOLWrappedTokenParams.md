# TokenizationSOLWrappedTokenParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Standard** | [**TokenizationTokenStandard**](TokenizationTokenStandard.md) |  | 
**Name** | **string** | The name of the token. | 
**Symbol** | **string** | The symbol of the token. | 
**Permissions** | Pointer to [**TokenizationSolWrappedTokenPermissionParams**](TokenizationSolWrappedTokenPermissionParams.md) |  | [optional] 
**UnderlyingToken** | **string** | The address of the underlying token that this tokenized asset represents. | 

## Methods

### NewTokenizationSOLWrappedTokenParams

`func NewTokenizationSOLWrappedTokenParams(standard TokenizationTokenStandard, name string, symbol string, underlyingToken string, ) *TokenizationSOLWrappedTokenParams`

NewTokenizationSOLWrappedTokenParams instantiates a new TokenizationSOLWrappedTokenParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationSOLWrappedTokenParamsWithDefaults

`func NewTokenizationSOLWrappedTokenParamsWithDefaults() *TokenizationSOLWrappedTokenParams`

NewTokenizationSOLWrappedTokenParamsWithDefaults instantiates a new TokenizationSOLWrappedTokenParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandard

`func (o *TokenizationSOLWrappedTokenParams) GetStandard() TokenizationTokenStandard`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *TokenizationSOLWrappedTokenParams) GetStandardOk() (*TokenizationTokenStandard, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *TokenizationSOLWrappedTokenParams) SetStandard(v TokenizationTokenStandard)`

SetStandard sets Standard field to given value.


### GetName

`func (o *TokenizationSOLWrappedTokenParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenizationSOLWrappedTokenParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenizationSOLWrappedTokenParams) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *TokenizationSOLWrappedTokenParams) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenizationSOLWrappedTokenParams) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenizationSOLWrappedTokenParams) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetPermissions

`func (o *TokenizationSOLWrappedTokenParams) GetPermissions() TokenizationSolWrappedTokenPermissionParams`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TokenizationSOLWrappedTokenParams) GetPermissionsOk() (*TokenizationSolWrappedTokenPermissionParams, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TokenizationSOLWrappedTokenParams) SetPermissions(v TokenizationSolWrappedTokenPermissionParams)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *TokenizationSOLWrappedTokenParams) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetUnderlyingToken

`func (o *TokenizationSOLWrappedTokenParams) GetUnderlyingToken() string`

GetUnderlyingToken returns the UnderlyingToken field if non-nil, zero value otherwise.

### GetUnderlyingTokenOk

`func (o *TokenizationSOLWrappedTokenParams) GetUnderlyingTokenOk() (*string, bool)`

GetUnderlyingTokenOk returns a tuple with the UnderlyingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingToken

`func (o *TokenizationSOLWrappedTokenParams) SetUnderlyingToken(v string)`

SetUnderlyingToken sets UnderlyingToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


