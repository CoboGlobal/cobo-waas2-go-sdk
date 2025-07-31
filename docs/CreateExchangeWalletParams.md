# CreateExchangeWalletParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The wallet name. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**EnableAutoSweep** | Pointer to **bool** | Enable the auto sweep feature for the wallet. This parameter only applies to MPC and Web3 wallets. | [optional] 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**Apikey** | **string** | The API key of your exchange account. | 
**Secret** | **string** | The API secret of your exchange account. | 
**Passphrase** | Pointer to **string** | The passphrase of your exchange account. | [optional] 
**Memo** | Pointer to **string** | The memo you use when applying for the API key of your exchange account. | [optional] 
**AccountIdentify** | Pointer to **string** | The identifier of your exchange account. - For Binance, this is email address of your exchange account. - For OKX, this is the user name of your exchange account.  | [optional] 
**GaCode** | Pointer to **string** | The GA code for the exchange. | [optional] 
**MainWalletId** | Pointer to **string** | The ID of the Exchange Wallet (Main Account). | [optional] 

## Methods

### NewCreateExchangeWalletParams

`func NewCreateExchangeWalletParams(name string, walletType WalletType, walletSubtype WalletSubtype, exchangeId ExchangeId, apikey string, secret string, ) *CreateExchangeWalletParams`

NewCreateExchangeWalletParams instantiates a new CreateExchangeWalletParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExchangeWalletParamsWithDefaults

`func NewCreateExchangeWalletParamsWithDefaults() *CreateExchangeWalletParams`

NewCreateExchangeWalletParamsWithDefaults instantiates a new CreateExchangeWalletParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateExchangeWalletParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateExchangeWalletParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateExchangeWalletParams) SetName(v string)`

SetName sets Name field to given value.


### GetWalletType

`func (o *CreateExchangeWalletParams) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *CreateExchangeWalletParams) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *CreateExchangeWalletParams) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *CreateExchangeWalletParams) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *CreateExchangeWalletParams) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *CreateExchangeWalletParams) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetEnableAutoSweep

`func (o *CreateExchangeWalletParams) GetEnableAutoSweep() bool`

GetEnableAutoSweep returns the EnableAutoSweep field if non-nil, zero value otherwise.

### GetEnableAutoSweepOk

`func (o *CreateExchangeWalletParams) GetEnableAutoSweepOk() (*bool, bool)`

GetEnableAutoSweepOk returns a tuple with the EnableAutoSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoSweep

`func (o *CreateExchangeWalletParams) SetEnableAutoSweep(v bool)`

SetEnableAutoSweep sets EnableAutoSweep field to given value.

### HasEnableAutoSweep

`func (o *CreateExchangeWalletParams) HasEnableAutoSweep() bool`

HasEnableAutoSweep returns a boolean if a field has been set.

### GetExchangeId

`func (o *CreateExchangeWalletParams) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *CreateExchangeWalletParams) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *CreateExchangeWalletParams) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.


### GetApikey

`func (o *CreateExchangeWalletParams) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *CreateExchangeWalletParams) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *CreateExchangeWalletParams) SetApikey(v string)`

SetApikey sets Apikey field to given value.


### GetSecret

`func (o *CreateExchangeWalletParams) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateExchangeWalletParams) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateExchangeWalletParams) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetPassphrase

`func (o *CreateExchangeWalletParams) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *CreateExchangeWalletParams) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *CreateExchangeWalletParams) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *CreateExchangeWalletParams) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetMemo

`func (o *CreateExchangeWalletParams) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CreateExchangeWalletParams) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CreateExchangeWalletParams) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CreateExchangeWalletParams) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetAccountIdentify

`func (o *CreateExchangeWalletParams) GetAccountIdentify() string`

GetAccountIdentify returns the AccountIdentify field if non-nil, zero value otherwise.

### GetAccountIdentifyOk

`func (o *CreateExchangeWalletParams) GetAccountIdentifyOk() (*string, bool)`

GetAccountIdentifyOk returns a tuple with the AccountIdentify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentify

`func (o *CreateExchangeWalletParams) SetAccountIdentify(v string)`

SetAccountIdentify sets AccountIdentify field to given value.

### HasAccountIdentify

`func (o *CreateExchangeWalletParams) HasAccountIdentify() bool`

HasAccountIdentify returns a boolean if a field has been set.

### GetGaCode

`func (o *CreateExchangeWalletParams) GetGaCode() string`

GetGaCode returns the GaCode field if non-nil, zero value otherwise.

### GetGaCodeOk

`func (o *CreateExchangeWalletParams) GetGaCodeOk() (*string, bool)`

GetGaCodeOk returns a tuple with the GaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaCode

`func (o *CreateExchangeWalletParams) SetGaCode(v string)`

SetGaCode sets GaCode field to given value.

### HasGaCode

`func (o *CreateExchangeWalletParams) HasGaCode() bool`

HasGaCode returns a boolean if a field has been set.

### GetMainWalletId

`func (o *CreateExchangeWalletParams) GetMainWalletId() string`

GetMainWalletId returns the MainWalletId field if non-nil, zero value otherwise.

### GetMainWalletIdOk

`func (o *CreateExchangeWalletParams) GetMainWalletIdOk() (*string, bool)`

GetMainWalletIdOk returns a tuple with the MainWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainWalletId

`func (o *CreateExchangeWalletParams) SetMainWalletId(v string)`

SetMainWalletId sets MainWalletId field to given value.

### HasMainWalletId

`func (o *CreateExchangeWalletParams) HasMainWalletId() bool`

HasMainWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


