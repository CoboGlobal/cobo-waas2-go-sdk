# TokenizationSolWrappedTokenPermissionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wrapper** | Pointer to **[]string** | List of Solana wallet addresses that can perform wrap/unwrap operations. Multiple addresses can be assigned this role. | [optional] 
**Pauser** | Pointer to **string** | Solana wallet address that acts as a pauser authority for the token. This authority can pause token transfers. | [optional] 
**Freezer** | Pointer to **string** | Solana wallet address that acts as a freezer authority for the token. This authority can freeze token accounts. | [optional] 
**Updater** | Pointer to **string** | Solana wallet address that acts as an updater authority for the token. This authority can update token metadata. | [optional] 

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

`func (o *TokenizationSolWrappedTokenPermissionParams) GetPauser() string`

GetPauser returns the Pauser field if non-nil, zero value otherwise.

### GetPauserOk

`func (o *TokenizationSolWrappedTokenPermissionParams) GetPauserOk() (*string, bool)`

GetPauserOk returns a tuple with the Pauser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauser

`func (o *TokenizationSolWrappedTokenPermissionParams) SetPauser(v string)`

SetPauser sets Pauser field to given value.

### HasPauser

`func (o *TokenizationSolWrappedTokenPermissionParams) HasPauser() bool`

HasPauser returns a boolean if a field has been set.

### GetFreezer

`func (o *TokenizationSolWrappedTokenPermissionParams) GetFreezer() string`

GetFreezer returns the Freezer field if non-nil, zero value otherwise.

### GetFreezerOk

`func (o *TokenizationSolWrappedTokenPermissionParams) GetFreezerOk() (*string, bool)`

GetFreezerOk returns a tuple with the Freezer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreezer

`func (o *TokenizationSolWrappedTokenPermissionParams) SetFreezer(v string)`

SetFreezer sets Freezer field to given value.

### HasFreezer

`func (o *TokenizationSolWrappedTokenPermissionParams) HasFreezer() bool`

HasFreezer returns a boolean if a field has been set.

### GetUpdater

`func (o *TokenizationSolWrappedTokenPermissionParams) GetUpdater() string`

GetUpdater returns the Updater field if non-nil, zero value otherwise.

### GetUpdaterOk

`func (o *TokenizationSolWrappedTokenPermissionParams) GetUpdaterOk() (*string, bool)`

GetUpdaterOk returns a tuple with the Updater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdater

`func (o *TokenizationSolWrappedTokenPermissionParams) SetUpdater(v string)`

SetUpdater sets Updater field to given value.

### HasUpdater

`func (o *TokenizationSolWrappedTokenPermissionParams) HasUpdater() bool`

HasUpdater returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


