# TokenizationSolWrappedTokenPermissionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wrapper** | Pointer to **[]string** | List of Solana wallet addresses that can perform wrap/unwrap operations. Multiple addresses can be assigned this role. | [optional] 
**Pauser** | Pointer to **[]string** | List of Solana wallet addresses that can pause/unpause the contract. Multiple addresses can be assigned this role. | [optional] 

## Methods

### NewTokenizationSolWrappedTokenPermissionParams

`func NewTokenizationSolWrappedTokenPermissionParams() *TokenizationSolWrappedTokenPermissionParams`

NewTokenizationSolWrappedTokenPermissionParams instantiates a new TokenizationSolWrappedTokenPermissionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationSolWrappedTokenPermissionParamsWithDefaults

`func NewTokenizationSolWrappedTokenPermissionParamsWithDefaults() *TokenizationSolWrappedTokenPermissionParams`

NewTokenizationSolWrappedTokenPermissionParamsWithDefaults instantiates a new TokenizationSolWrappedTokenPermissionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWrapper

`func (o *TokenizationSolWrappedTokenPermissionParams) GetWrapper() []string`

GetWrapper returns the Wrapper field if non-nil, zero value otherwise.

### GetWrapperOk

`func (o *TokenizationSolWrappedTokenPermissionParams) GetWrapperOk() (*[]string, bool)`

GetWrapperOk returns a tuple with the Wrapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapper

`func (o *TokenizationSolWrappedTokenPermissionParams) SetWrapper(v []string)`

SetWrapper sets Wrapper field to given value.

### HasWrapper

`func (o *TokenizationSolWrappedTokenPermissionParams) HasWrapper() bool`

HasWrapper returns a boolean if a field has been set.

### GetPauser

`func (o *TokenizationSolWrappedTokenPermissionParams) GetPauser() []string`

GetPauser returns the Pauser field if non-nil, zero value otherwise.

### GetPauserOk

`func (o *TokenizationSolWrappedTokenPermissionParams) GetPauserOk() (*[]string, bool)`

GetPauserOk returns a tuple with the Pauser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauser

`func (o *TokenizationSolWrappedTokenPermissionParams) SetPauser(v []string)`

SetPauser sets Pauser field to given value.

### HasPauser

`func (o *TokenizationSolWrappedTokenPermissionParams) HasPauser() bool`

HasPauser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


