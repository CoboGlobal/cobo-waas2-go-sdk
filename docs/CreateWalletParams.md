# CreateWalletParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The wallet name. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**VaultId** | **string** | The ID of the owning vault. You can call [List all vaults](https://www.cobo.com/developers/v2/api-references/wallets--mpc-wallets/list-all-vaults) to retrieve all vault IDs under your organization. | 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**Apikey** | **string** | The API key of your exchange account. | 
**Secret** | **string** | The API secret of your exchange account. | 
**Passphrase** | Pointer to **string** | The passphrase of your exchange account. | [optional] 
**Memo** | Pointer to **string** | The memo you use when applying for the API key of your exchange account. | [optional] 
**AccountIdentify** | Pointer to **string** | The identifier of your exchange account. - For Binance, this is email address of your exchange account. - For OKX, this is the user name of your exchange account.  | [optional] 
**GaCode** | Pointer to **string** | The GA code for the exchange. | [optional] 
**MainWalletId** | Pointer to **string** | The ID of the Exchange Wallet (Main Account). | [optional] 

## Methods

### NewCreateWalletParams

`func NewCreateWalletParams(name string, walletType WalletType, walletSubtype WalletSubtype, vaultId string, exchangeId ExchangeId, apikey string, secret string, ) *CreateWalletParams`

NewCreateWalletParams instantiates a new CreateWalletParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWalletParamsWithDefaults

`func NewCreateWalletParamsWithDefaults() *CreateWalletParams`

NewCreateWalletParamsWithDefaults instantiates a new CreateWalletParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateWalletParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWalletParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWalletParams) SetName(v string)`

SetName sets Name field to given value.


### GetWalletType

`func (o *CreateWalletParams) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *CreateWalletParams) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *CreateWalletParams) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *CreateWalletParams) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *CreateWalletParams) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *CreateWalletParams) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetVaultId

`func (o *CreateWalletParams) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *CreateWalletParams) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *CreateWalletParams) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.


### GetExchangeId

`func (o *CreateWalletParams) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *CreateWalletParams) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *CreateWalletParams) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.


### GetApikey

`func (o *CreateWalletParams) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *CreateWalletParams) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *CreateWalletParams) SetApikey(v string)`

SetApikey sets Apikey field to given value.


### GetSecret

`func (o *CreateWalletParams) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateWalletParams) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateWalletParams) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetPassphrase

`func (o *CreateWalletParams) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *CreateWalletParams) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *CreateWalletParams) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *CreateWalletParams) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetMemo

`func (o *CreateWalletParams) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CreateWalletParams) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CreateWalletParams) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CreateWalletParams) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetAccountIdentify

`func (o *CreateWalletParams) GetAccountIdentify() string`

GetAccountIdentify returns the AccountIdentify field if non-nil, zero value otherwise.

### GetAccountIdentifyOk

`func (o *CreateWalletParams) GetAccountIdentifyOk() (*string, bool)`

GetAccountIdentifyOk returns a tuple with the AccountIdentify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentify

`func (o *CreateWalletParams) SetAccountIdentify(v string)`

SetAccountIdentify sets AccountIdentify field to given value.

### HasAccountIdentify

`func (o *CreateWalletParams) HasAccountIdentify() bool`

HasAccountIdentify returns a boolean if a field has been set.

### GetGaCode

`func (o *CreateWalletParams) GetGaCode() string`

GetGaCode returns the GaCode field if non-nil, zero value otherwise.

### GetGaCodeOk

`func (o *CreateWalletParams) GetGaCodeOk() (*string, bool)`

GetGaCodeOk returns a tuple with the GaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaCode

`func (o *CreateWalletParams) SetGaCode(v string)`

SetGaCode sets GaCode field to given value.

### HasGaCode

`func (o *CreateWalletParams) HasGaCode() bool`

HasGaCode returns a boolean if a field has been set.

### GetMainWalletId

`func (o *CreateWalletParams) GetMainWalletId() string`

GetMainWalletId returns the MainWalletId field if non-nil, zero value otherwise.

### GetMainWalletIdOk

`func (o *CreateWalletParams) GetMainWalletIdOk() (*string, bool)`

GetMainWalletIdOk returns a tuple with the MainWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainWalletId

`func (o *CreateWalletParams) SetMainWalletId(v string)`

SetMainWalletId sets MainWalletId field to given value.

### HasMainWalletId

`func (o *CreateWalletParams) HasMainWalletId() bool`

HasMainWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


