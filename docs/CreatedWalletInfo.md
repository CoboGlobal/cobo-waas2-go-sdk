# CreatedWalletInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The wallet ID. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**Name** | **string** | The wallet name. | 
**OrgId** | **string** | The ID of the owning organization. | 
**ProjectId** | Pointer to **string** | The project ID. | [optional] 
**VaultId** | **string** | The ID of the owning vault. | 
**Apikey** | **string** | The API key of your exchange account. | 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**ParentWalletId** | Pointer to **string** | The wallet ID of the Main Account associated with the Sub Account. This property is returned only if you are creating or querying an Exchange Wallet (Sub Account). | [optional] 
**SubAccounts** | Pointer to [**[]ExchangeWalletInfoAllOfSubAccounts**](ExchangeWalletInfoAllOfSubAccounts.md) |  | [optional] 

## Methods

### NewCreatedWalletInfo

`func NewCreatedWalletInfo(walletId string, walletType WalletType, walletSubtype WalletSubtype, name string, orgId string, vaultId string, apikey string, exchangeId ExchangeId, ) *CreatedWalletInfo`

NewCreatedWalletInfo instantiates a new CreatedWalletInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedWalletInfoWithDefaults

`func NewCreatedWalletInfoWithDefaults() *CreatedWalletInfo`

NewCreatedWalletInfoWithDefaults instantiates a new CreatedWalletInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *CreatedWalletInfo) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CreatedWalletInfo) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CreatedWalletInfo) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetWalletType

`func (o *CreatedWalletInfo) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *CreatedWalletInfo) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *CreatedWalletInfo) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *CreatedWalletInfo) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *CreatedWalletInfo) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *CreatedWalletInfo) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetName

`func (o *CreatedWalletInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatedWalletInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatedWalletInfo) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *CreatedWalletInfo) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CreatedWalletInfo) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CreatedWalletInfo) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetProjectId

`func (o *CreatedWalletInfo) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreatedWalletInfo) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreatedWalletInfo) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreatedWalletInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetVaultId

`func (o *CreatedWalletInfo) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *CreatedWalletInfo) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *CreatedWalletInfo) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.


### GetApikey

`func (o *CreatedWalletInfo) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *CreatedWalletInfo) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *CreatedWalletInfo) SetApikey(v string)`

SetApikey sets Apikey field to given value.


### GetExchangeId

`func (o *CreatedWalletInfo) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *CreatedWalletInfo) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *CreatedWalletInfo) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.


### GetParentWalletId

`func (o *CreatedWalletInfo) GetParentWalletId() string`

GetParentWalletId returns the ParentWalletId field if non-nil, zero value otherwise.

### GetParentWalletIdOk

`func (o *CreatedWalletInfo) GetParentWalletIdOk() (*string, bool)`

GetParentWalletIdOk returns a tuple with the ParentWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWalletId

`func (o *CreatedWalletInfo) SetParentWalletId(v string)`

SetParentWalletId sets ParentWalletId field to given value.

### HasParentWalletId

`func (o *CreatedWalletInfo) HasParentWalletId() bool`

HasParentWalletId returns a boolean if a field has been set.

### GetSubAccounts

`func (o *CreatedWalletInfo) GetSubAccounts() []ExchangeWalletInfoAllOfSubAccounts`

GetSubAccounts returns the SubAccounts field if non-nil, zero value otherwise.

### GetSubAccountsOk

`func (o *CreatedWalletInfo) GetSubAccountsOk() (*[]ExchangeWalletInfoAllOfSubAccounts, bool)`

GetSubAccountsOk returns a tuple with the SubAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccounts

`func (o *CreatedWalletInfo) SetSubAccounts(v []ExchangeWalletInfoAllOfSubAccounts)`

SetSubAccounts sets SubAccounts field to given value.

### HasSubAccounts

`func (o *CreatedWalletInfo) HasSubAccounts() bool`

HasSubAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


