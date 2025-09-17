# CreateMpcWalletParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The wallet name. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**EnableAutoSweep** | Pointer to **bool** | Enable the auto-sweep feature for the wallet. This parameter only applies to MPC Wallets and Web3 Wallets. | [optional] 
**VaultId** | **string** | The ID of the owning vault. You can call [List all vaults](https://www.cobo.com/developers/v2/api-references/wallets--mpc-wallets/list-all-vaults) to retrieve all vault IDs under your organization. | 

## Methods

### NewCreateMpcWalletParams

`func NewCreateMpcWalletParams(name string, walletType WalletType, walletSubtype WalletSubtype, vaultId string, ) *CreateMpcWalletParams`

NewCreateMpcWalletParams instantiates a new CreateMpcWalletParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMpcWalletParamsWithDefaults

`func NewCreateMpcWalletParamsWithDefaults() *CreateMpcWalletParams`

NewCreateMpcWalletParamsWithDefaults instantiates a new CreateMpcWalletParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMpcWalletParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMpcWalletParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMpcWalletParams) SetName(v string)`

SetName sets Name field to given value.


### GetWalletType

`func (o *CreateMpcWalletParams) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *CreateMpcWalletParams) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *CreateMpcWalletParams) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *CreateMpcWalletParams) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *CreateMpcWalletParams) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *CreateMpcWalletParams) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetEnableAutoSweep

`func (o *CreateMpcWalletParams) GetEnableAutoSweep() bool`

GetEnableAutoSweep returns the EnableAutoSweep field if non-nil, zero value otherwise.

### GetEnableAutoSweepOk

`func (o *CreateMpcWalletParams) GetEnableAutoSweepOk() (*bool, bool)`

GetEnableAutoSweepOk returns a tuple with the EnableAutoSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoSweep

`func (o *CreateMpcWalletParams) SetEnableAutoSweep(v bool)`

SetEnableAutoSweep sets EnableAutoSweep field to given value.

### HasEnableAutoSweep

`func (o *CreateMpcWalletParams) HasEnableAutoSweep() bool`

HasEnableAutoSweep returns a boolean if a field has been set.

### GetVaultId

`func (o *CreateMpcWalletParams) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *CreateMpcWalletParams) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *CreateMpcWalletParams) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


