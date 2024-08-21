# WalletInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The wallet ID. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**Name** | **string** | The wallet name. | 
**OrgId** | **string** | The ID of the owning organization. | 
**ProjectId** | Pointer to **string** | The project ID. | [optional] 
**ProjectName** | Pointer to **string** | The project name. | [optional] 
**VaultId** | **string** | The ID of the owning vault. | 
**VaultName** | Pointer to **string** | The vault name. | [optional] 
**Apikey** | **string** | The API key of your exchange account. | 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**MainWalletId** | Pointer to **string** | The wallet ID of the Main Account associated with the Sub Account. This property is returned only if you are creating or querying an Exchange Wallet (Sub Account). | [optional] 

## Methods

### NewWalletInfo

`func NewWalletInfo(walletId string, walletType WalletType, walletSubtype WalletSubtype, name string, orgId string, vaultId string, apikey string, exchangeId ExchangeId, ) *WalletInfo`

NewWalletInfo instantiates a new WalletInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletInfoWithDefaults

`func NewWalletInfoWithDefaults() *WalletInfo`

NewWalletInfoWithDefaults instantiates a new WalletInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *WalletInfo) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *WalletInfo) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *WalletInfo) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetWalletType

`func (o *WalletInfo) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *WalletInfo) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *WalletInfo) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *WalletInfo) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *WalletInfo) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *WalletInfo) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetName

`func (o *WalletInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletInfo) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *WalletInfo) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WalletInfo) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WalletInfo) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetProjectId

`func (o *WalletInfo) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WalletInfo) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WalletInfo) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *WalletInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *WalletInfo) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *WalletInfo) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *WalletInfo) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *WalletInfo) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetVaultId

`func (o *WalletInfo) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *WalletInfo) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *WalletInfo) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.


### GetVaultName

`func (o *WalletInfo) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *WalletInfo) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *WalletInfo) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.

### HasVaultName

`func (o *WalletInfo) HasVaultName() bool`

HasVaultName returns a boolean if a field has been set.

### GetApikey

`func (o *WalletInfo) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *WalletInfo) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *WalletInfo) SetApikey(v string)`

SetApikey sets Apikey field to given value.


### GetExchangeId

`func (o *WalletInfo) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *WalletInfo) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *WalletInfo) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.


### GetMainWalletId

`func (o *WalletInfo) GetMainWalletId() string`

GetMainWalletId returns the MainWalletId field if non-nil, zero value otherwise.

### GetMainWalletIdOk

`func (o *WalletInfo) GetMainWalletIdOk() (*string, bool)`

GetMainWalletIdOk returns a tuple with the MainWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainWalletId

`func (o *WalletInfo) SetMainWalletId(v string)`

SetMainWalletId sets MainWalletId field to given value.

### HasMainWalletId

`func (o *WalletInfo) HasMainWalletId() bool`

HasMainWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


