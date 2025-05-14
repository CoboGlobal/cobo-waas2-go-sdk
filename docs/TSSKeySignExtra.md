# TSSKeySignExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to [**OrgInfo**](OrgInfo.md) |  | [optional] 
**Project** | Pointer to [**MPCProject**](MPCProject.md) |  | [optional] 
**Vault** | Pointer to [**MPCVault**](MPCVault.md) |  | [optional] 
**Wallet** | Pointer to [**MPCWalletInfo**](MPCWalletInfo.md) |  | [optional] 
**SignerKeyShareHolderGroup** | Pointer to [**KeyShareHolderGroup**](KeyShareHolderGroup.md) |  | [optional] 
**SourceAddresses** | Pointer to [**[]AddressInfo**](AddressInfo.md) |  | [optional] 
**Transaction** | Pointer to [**Transaction**](Transaction.md) |  | [optional] 

## Methods

### NewTSSKeySignExtra

`func NewTSSKeySignExtra() *TSSKeySignExtra`

NewTSSKeySignExtra instantiates a new TSSKeySignExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSKeySignExtraWithDefaults

`func NewTSSKeySignExtraWithDefaults() *TSSKeySignExtra`

NewTSSKeySignExtraWithDefaults instantiates a new TSSKeySignExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *TSSKeySignExtra) GetOrg() OrgInfo`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *TSSKeySignExtra) GetOrgOk() (*OrgInfo, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *TSSKeySignExtra) SetOrg(v OrgInfo)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *TSSKeySignExtra) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *TSSKeySignExtra) GetProject() MPCProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TSSKeySignExtra) GetProjectOk() (*MPCProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TSSKeySignExtra) SetProject(v MPCProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *TSSKeySignExtra) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetVault

`func (o *TSSKeySignExtra) GetVault() MPCVault`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *TSSKeySignExtra) GetVaultOk() (*MPCVault, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *TSSKeySignExtra) SetVault(v MPCVault)`

SetVault sets Vault field to given value.

### HasVault

`func (o *TSSKeySignExtra) HasVault() bool`

HasVault returns a boolean if a field has been set.

### GetWallet

`func (o *TSSKeySignExtra) GetWallet() MPCWalletInfo`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *TSSKeySignExtra) GetWalletOk() (*MPCWalletInfo, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *TSSKeySignExtra) SetWallet(v MPCWalletInfo)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *TSSKeySignExtra) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetSignerKeyShareHolderGroup

`func (o *TSSKeySignExtra) GetSignerKeyShareHolderGroup() KeyShareHolderGroup`

GetSignerKeyShareHolderGroup returns the SignerKeyShareHolderGroup field if non-nil, zero value otherwise.

### GetSignerKeyShareHolderGroupOk

`func (o *TSSKeySignExtra) GetSignerKeyShareHolderGroupOk() (*KeyShareHolderGroup, bool)`

GetSignerKeyShareHolderGroupOk returns a tuple with the SignerKeyShareHolderGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerKeyShareHolderGroup

`func (o *TSSKeySignExtra) SetSignerKeyShareHolderGroup(v KeyShareHolderGroup)`

SetSignerKeyShareHolderGroup sets SignerKeyShareHolderGroup field to given value.

### HasSignerKeyShareHolderGroup

`func (o *TSSKeySignExtra) HasSignerKeyShareHolderGroup() bool`

HasSignerKeyShareHolderGroup returns a boolean if a field has been set.

### GetSourceAddresses

`func (o *TSSKeySignExtra) GetSourceAddresses() []AddressInfo`

GetSourceAddresses returns the SourceAddresses field if non-nil, zero value otherwise.

### GetSourceAddressesOk

`func (o *TSSKeySignExtra) GetSourceAddressesOk() (*[]AddressInfo, bool)`

GetSourceAddressesOk returns a tuple with the SourceAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddresses

`func (o *TSSKeySignExtra) SetSourceAddresses(v []AddressInfo)`

SetSourceAddresses sets SourceAddresses field to given value.

### HasSourceAddresses

`func (o *TSSKeySignExtra) HasSourceAddresses() bool`

HasSourceAddresses returns a boolean if a field has been set.

### GetTransaction

`func (o *TSSKeySignExtra) GetTransaction() Transaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *TSSKeySignExtra) GetTransactionOk() (*Transaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *TSSKeySignExtra) SetTransaction(v Transaction)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *TSSKeySignExtra) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


