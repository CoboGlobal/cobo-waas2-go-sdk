# WalletInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The wallet ID. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**Name** | **string** | The wallet name. | 
**OrgId** | **string** | The ID of the owning organization. | 
**VaultId** | **string** | The ID of the owning vault. | 
**Apikey** | **string** | The API key of your exchange account. | 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**ParentWalletId** | Pointer to **string** | The wallet ID of the Main Account associated with the Sub Account. This property is returned only if you are creating or querying an Exchange Wallet (Sub Account). | [optional] 
**SubAccounts** | Pointer to [**[]ExchangeWalletInfoAllOfSubAccounts**](ExchangeWalletInfoAllOfSubAccounts.md) |  | [optional] 

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


### GetParentWalletId

`func (o *WalletInfo) GetParentWalletId() string`

GetParentWalletId returns the ParentWalletId field if non-nil, zero value otherwise.

### GetParentWalletIdOk

`func (o *WalletInfo) GetParentWalletIdOk() (*string, bool)`

GetParentWalletIdOk returns a tuple with the ParentWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWalletId

`func (o *WalletInfo) SetParentWalletId(v string)`

SetParentWalletId sets ParentWalletId field to given value.

### HasParentWalletId

`func (o *WalletInfo) HasParentWalletId() bool`

HasParentWalletId returns a boolean if a field has been set.

### GetSubAccounts

`func (o *WalletInfo) GetSubAccounts() []ExchangeWalletInfoAllOfSubAccounts`

GetSubAccounts returns the SubAccounts field if non-nil, zero value otherwise.

### GetSubAccountsOk

`func (o *WalletInfo) GetSubAccountsOk() (*[]ExchangeWalletInfoAllOfSubAccounts, bool)`

GetSubAccountsOk returns a tuple with the SubAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccounts

`func (o *WalletInfo) SetSubAccounts(v []ExchangeWalletInfoAllOfSubAccounts)`

SetSubAccounts sets SubAccounts field to given value.

### HasSubAccounts

`func (o *WalletInfo) HasSubAccounts() bool`

HasSubAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


