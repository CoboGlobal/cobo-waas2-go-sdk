# TokenizationSolTokenPermissionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermanentDelegate** | Pointer to **string** | Solana wallet address that acts as a permanent delegate authority for the token. This authority can perform delegated operations on behalf of token holders. | [optional] 
**Minter** | Pointer to **string** | Solana wallet addres that acts as a minter authority for the token. This authority can mint new tokens. | [optional] 
**Freezer** | Pointer to **string** | Solana wallet address that acts as a freezer authority for the token. This authority can freeze token accounts. | [optional] 
**Updater** | Pointer to **string** | Solana wallet address that acts as an updater authority for the token. This authority can update token metadata. | [optional] 
**Pauser** | Pointer to **string** | Solana wallet address that acts as a pauser authority for the token. This authority can pause token transfers. | [optional] 

## Methods

### NewTokenizationSolTokenPermissionParams

`func NewTokenizationSolTokenPermissionParams() *TokenizationSolTokenPermissionParams`

NewTokenizationSolTokenPermissionParams instantiates a new TokenizationSolTokenPermissionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationSolTokenPermissionParamsWithDefaults

`func NewTokenizationSolTokenPermissionParamsWithDefaults() *TokenizationSolTokenPermissionParams`

NewTokenizationSolTokenPermissionParamsWithDefaults instantiates a new TokenizationSolTokenPermissionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermanentDelegate

`func (o *TokenizationSolTokenPermissionParams) GetPermanentDelegate() string`

GetPermanentDelegate returns the PermanentDelegate field if non-nil, zero value otherwise.

### GetPermanentDelegateOk

`func (o *TokenizationSolTokenPermissionParams) GetPermanentDelegateOk() (*string, bool)`

GetPermanentDelegateOk returns a tuple with the PermanentDelegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentDelegate

`func (o *TokenizationSolTokenPermissionParams) SetPermanentDelegate(v string)`

SetPermanentDelegate sets PermanentDelegate field to given value.

### HasPermanentDelegate

`func (o *TokenizationSolTokenPermissionParams) HasPermanentDelegate() bool`

HasPermanentDelegate returns a boolean if a field has been set.

### GetMinter

`func (o *TokenizationSolTokenPermissionParams) GetMinter() string`

GetMinter returns the Minter field if non-nil, zero value otherwise.

### GetMinterOk

`func (o *TokenizationSolTokenPermissionParams) GetMinterOk() (*string, bool)`

GetMinterOk returns a tuple with the Minter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinter

`func (o *TokenizationSolTokenPermissionParams) SetMinter(v string)`

SetMinter sets Minter field to given value.

### HasMinter

`func (o *TokenizationSolTokenPermissionParams) HasMinter() bool`

HasMinter returns a boolean if a field has been set.

### GetFreezer

`func (o *TokenizationSolTokenPermissionParams) GetFreezer() string`

GetFreezer returns the Freezer field if non-nil, zero value otherwise.

### GetFreezerOk

`func (o *TokenizationSolTokenPermissionParams) GetFreezerOk() (*string, bool)`

GetFreezerOk returns a tuple with the Freezer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreezer

`func (o *TokenizationSolTokenPermissionParams) SetFreezer(v string)`

SetFreezer sets Freezer field to given value.

### HasFreezer

`func (o *TokenizationSolTokenPermissionParams) HasFreezer() bool`

HasFreezer returns a boolean if a field has been set.

### GetUpdater

`func (o *TokenizationSolTokenPermissionParams) GetUpdater() string`

GetUpdater returns the Updater field if non-nil, zero value otherwise.

### GetUpdaterOk

`func (o *TokenizationSolTokenPermissionParams) GetUpdaterOk() (*string, bool)`

GetUpdaterOk returns a tuple with the Updater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdater

`func (o *TokenizationSolTokenPermissionParams) SetUpdater(v string)`

SetUpdater sets Updater field to given value.

### HasUpdater

`func (o *TokenizationSolTokenPermissionParams) HasUpdater() bool`

HasUpdater returns a boolean if a field has been set.

### GetPauser

`func (o *TokenizationSolTokenPermissionParams) GetPauser() string`

GetPauser returns the Pauser field if non-nil, zero value otherwise.

### GetPauserOk

`func (o *TokenizationSolTokenPermissionParams) GetPauserOk() (*string, bool)`

GetPauserOk returns a tuple with the Pauser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauser

`func (o *TokenizationSolTokenPermissionParams) SetPauser(v string)`

SetPauser sets Pauser field to given value.

### HasPauser

`func (o *TokenizationSolTokenPermissionParams) HasPauser() bool`

HasPauser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


