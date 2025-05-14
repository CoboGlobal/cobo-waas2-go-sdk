# TSSKeyShareSignExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to [**OrgInfo**](OrgInfo.md) |  | [optional] 
**Project** | Pointer to [**MPCProject**](MPCProject.md) |  | [optional] 
**Vault** | Pointer to [**MPCVault**](MPCVault.md) |  | [optional] 
**Wallet** | Pointer to [**MPCWalletInfo**](MPCWalletInfo.md) |  | [optional] 
**ValidityKeyShareHolderGroups** | Pointer to [**[]KeyShareHolderGroup**](KeyShareHolderGroup.md) |  | [optional] 

## Methods

### NewTSSKeyShareSignExtra

`func NewTSSKeyShareSignExtra() *TSSKeyShareSignExtra`

NewTSSKeyShareSignExtra instantiates a new TSSKeyShareSignExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSKeyShareSignExtraWithDefaults

`func NewTSSKeyShareSignExtraWithDefaults() *TSSKeyShareSignExtra`

NewTSSKeyShareSignExtraWithDefaults instantiates a new TSSKeyShareSignExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *TSSKeyShareSignExtra) GetOrg() OrgInfo`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *TSSKeyShareSignExtra) GetOrgOk() (*OrgInfo, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *TSSKeyShareSignExtra) SetOrg(v OrgInfo)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *TSSKeyShareSignExtra) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *TSSKeyShareSignExtra) GetProject() MPCProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TSSKeyShareSignExtra) GetProjectOk() (*MPCProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TSSKeyShareSignExtra) SetProject(v MPCProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *TSSKeyShareSignExtra) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetVault

`func (o *TSSKeyShareSignExtra) GetVault() MPCVault`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *TSSKeyShareSignExtra) GetVaultOk() (*MPCVault, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *TSSKeyShareSignExtra) SetVault(v MPCVault)`

SetVault sets Vault field to given value.

### HasVault

`func (o *TSSKeyShareSignExtra) HasVault() bool`

HasVault returns a boolean if a field has been set.

### GetWallet

`func (o *TSSKeyShareSignExtra) GetWallet() MPCWalletInfo`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *TSSKeyShareSignExtra) GetWalletOk() (*MPCWalletInfo, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *TSSKeyShareSignExtra) SetWallet(v MPCWalletInfo)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *TSSKeyShareSignExtra) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetValidityKeyShareHolderGroups

`func (o *TSSKeyShareSignExtra) GetValidityKeyShareHolderGroups() []KeyShareHolderGroup`

GetValidityKeyShareHolderGroups returns the ValidityKeyShareHolderGroups field if non-nil, zero value otherwise.

### GetValidityKeyShareHolderGroupsOk

`func (o *TSSKeyShareSignExtra) GetValidityKeyShareHolderGroupsOk() (*[]KeyShareHolderGroup, bool)`

GetValidityKeyShareHolderGroupsOk returns a tuple with the ValidityKeyShareHolderGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityKeyShareHolderGroups

`func (o *TSSKeyShareSignExtra) SetValidityKeyShareHolderGroups(v []KeyShareHolderGroup)`

SetValidityKeyShareHolderGroups sets ValidityKeyShareHolderGroups field to given value.

### HasValidityKeyShareHolderGroups

`func (o *TSSKeyShareSignExtra) HasValidityKeyShareHolderGroups() bool`

HasValidityKeyShareHolderGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


