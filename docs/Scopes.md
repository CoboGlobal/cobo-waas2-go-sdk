# Scopes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletTypes** | Pointer to [**[]WalletType**](WalletType.md) | The list of wallet types that this API key can access. Possible values include:   - &#x60;Custodial&#x60;: [Custodial Wallets](https://manuals.cobo.com/en/portal/custodial-wallets/introduction)   - &#x60;MPC&#x60;: [MPC Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction)   - &#x60;SmartContract&#x60;: [Smart Contract Wallets](https://manuals.cobo.com/en/portal/smart-contract-wallets/introduction)   - &#x60;Exchange&#x60;: [Exchange Wallets](https://manuals.cobo.com/en/portal/exchange-wallets/introduction)  | [optional] 
**WalletSubtypes** | Pointer to [**[]WalletSubtype**](WalletSubtype.md) | The list of wallet sub-types that this API key can access. Possible values include:   - &#x60;Asset&#x60;: Custodial Wallets (Asset Wallets).   - &#x60;Org-Controlled&#x60;: MPC Wallets (Organization-Controlled Wallets).   - &#x60;User-Controlled&#x60;: MPC Wallets (User-Controlled Wallets).   - &#x60;Safe{Wallet}&#x60;: Smart Contract Wallets (Safe).   - &#x60;Main&#x60;: Exchange Wallets (Main Account).   - &#x60;Sub&#x60;: Exchange Wallets (Sub Account).  | [optional] 
**WalletIds** | Pointer to **[]string** | The list of wallet IDs that this API key can access. | [optional] 
**VaultIds** | Pointer to **[]string** | (Applicable to MPC Wallets only) The list of vault IDs that this API key can access. | [optional] 
**ProjectIds** | Pointer to **[]string** | (Applicable to MPC Wallets only) The list of project IDs that this API key can access. | [optional] 

## Methods

### NewScopes

`func NewScopes() *Scopes`

NewScopes instantiates a new Scopes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopesWithDefaults

`func NewScopesWithDefaults() *Scopes`

NewScopesWithDefaults instantiates a new Scopes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletTypes

`func (o *Scopes) GetWalletTypes() []WalletType`

GetWalletTypes returns the WalletTypes field if non-nil, zero value otherwise.

### GetWalletTypesOk

`func (o *Scopes) GetWalletTypesOk() (*[]WalletType, bool)`

GetWalletTypesOk returns a tuple with the WalletTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletTypes

`func (o *Scopes) SetWalletTypes(v []WalletType)`

SetWalletTypes sets WalletTypes field to given value.

### HasWalletTypes

`func (o *Scopes) HasWalletTypes() bool`

HasWalletTypes returns a boolean if a field has been set.

### GetWalletSubtypes

`func (o *Scopes) GetWalletSubtypes() []WalletSubtype`

GetWalletSubtypes returns the WalletSubtypes field if non-nil, zero value otherwise.

### GetWalletSubtypesOk

`func (o *Scopes) GetWalletSubtypesOk() (*[]WalletSubtype, bool)`

GetWalletSubtypesOk returns a tuple with the WalletSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtypes

`func (o *Scopes) SetWalletSubtypes(v []WalletSubtype)`

SetWalletSubtypes sets WalletSubtypes field to given value.

### HasWalletSubtypes

`func (o *Scopes) HasWalletSubtypes() bool`

HasWalletSubtypes returns a boolean if a field has been set.

### GetWalletIds

`func (o *Scopes) GetWalletIds() []string`

GetWalletIds returns the WalletIds field if non-nil, zero value otherwise.

### GetWalletIdsOk

`func (o *Scopes) GetWalletIdsOk() (*[]string, bool)`

GetWalletIdsOk returns a tuple with the WalletIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletIds

`func (o *Scopes) SetWalletIds(v []string)`

SetWalletIds sets WalletIds field to given value.

### HasWalletIds

`func (o *Scopes) HasWalletIds() bool`

HasWalletIds returns a boolean if a field has been set.

### GetVaultIds

`func (o *Scopes) GetVaultIds() []string`

GetVaultIds returns the VaultIds field if non-nil, zero value otherwise.

### GetVaultIdsOk

`func (o *Scopes) GetVaultIdsOk() (*[]string, bool)`

GetVaultIdsOk returns a tuple with the VaultIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIds

`func (o *Scopes) SetVaultIds(v []string)`

SetVaultIds sets VaultIds field to given value.

### HasVaultIds

`func (o *Scopes) HasVaultIds() bool`

HasVaultIds returns a boolean if a field has been set.

### GetProjectIds

`func (o *Scopes) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *Scopes) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *Scopes) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *Scopes) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


