# TokenizationIssueTokenParamsTokenParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Standard** | [**TokenizationTokenStandard**](TokenizationTokenStandard.md) |  | 
**Name** | **string** | The name of the token. | 
**Symbol** | **string** | The symbol of the token. | 
**Decimals** | **int32** | The number of decimals for the token (0-18). | 
**TokenAccessActivated** | Pointer to **bool** | Whether the allowlist feature is activated for the token. When activated, only addresses in the allowlist can perform token operations. | [optional] [default to false]
**Permissions** | Pointer to [**TokenizationSolTokenPermissionParams**](TokenizationSolTokenPermissionParams.md) |  | [optional] 
**UnderlyingToken** | **string** | The address of the underlying token that this tokenized asset represents. | 

## Methods

### NewTokenizationIssueTokenParamsTokenParams

`func NewTokenizationIssueTokenParamsTokenParams(standard TokenizationTokenStandard, name string, symbol string, decimals int32, underlyingToken string, ) *TokenizationIssueTokenParamsTokenParams`

NewTokenizationIssueTokenParamsTokenParams instantiates a new TokenizationIssueTokenParamsTokenParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationIssueTokenParamsTokenParamsWithDefaults

`func NewTokenizationIssueTokenParamsTokenParamsWithDefaults() *TokenizationIssueTokenParamsTokenParams`

NewTokenizationIssueTokenParamsTokenParamsWithDefaults instantiates a new TokenizationIssueTokenParamsTokenParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandard

`func (o *TokenizationIssueTokenParamsTokenParams) GetStandard() TokenizationTokenStandard`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetStandardOk() (*TokenizationTokenStandard, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *TokenizationIssueTokenParamsTokenParams) SetStandard(v TokenizationTokenStandard)`

SetStandard sets Standard field to given value.


### GetName

`func (o *TokenizationIssueTokenParamsTokenParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenizationIssueTokenParamsTokenParams) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *TokenizationIssueTokenParamsTokenParams) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenizationIssueTokenParamsTokenParams) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimals

`func (o *TokenizationIssueTokenParamsTokenParams) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenizationIssueTokenParamsTokenParams) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetTokenAccessActivated

`func (o *TokenizationIssueTokenParamsTokenParams) GetTokenAccessActivated() bool`

GetTokenAccessActivated returns the TokenAccessActivated field if non-nil, zero value otherwise.

### GetTokenAccessActivatedOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetTokenAccessActivatedOk() (*bool, bool)`

GetTokenAccessActivatedOk returns a tuple with the TokenAccessActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAccessActivated

`func (o *TokenizationIssueTokenParamsTokenParams) SetTokenAccessActivated(v bool)`

SetTokenAccessActivated sets TokenAccessActivated field to given value.

### HasTokenAccessActivated

`func (o *TokenizationIssueTokenParamsTokenParams) HasTokenAccessActivated() bool`

HasTokenAccessActivated returns a boolean if a field has been set.

### GetPermissions

`func (o *TokenizationIssueTokenParamsTokenParams) GetPermissions() TokenizationSolTokenPermissionParams`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetPermissionsOk() (*TokenizationSolTokenPermissionParams, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TokenizationIssueTokenParamsTokenParams) SetPermissions(v TokenizationSolTokenPermissionParams)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *TokenizationIssueTokenParamsTokenParams) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetUnderlyingToken

`func (o *TokenizationIssueTokenParamsTokenParams) GetUnderlyingToken() string`

GetUnderlyingToken returns the UnderlyingToken field if non-nil, zero value otherwise.

### GetUnderlyingTokenOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetUnderlyingTokenOk() (*string, bool)`

GetUnderlyingTokenOk returns a tuple with the UnderlyingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingToken

`func (o *TokenizationIssueTokenParamsTokenParams) SetUnderlyingToken(v string)`

SetUnderlyingToken sets UnderlyingToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


