# UpdateExchangeWalletParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletType** | [**WalletType**](WalletType.md) |  | 
**Name** | Pointer to **string** | The wallet name. | [optional] 
**Apikey** | Pointer to **string** | The API key of your exchange account. This property is required when updating the information of an Exchange Wallet. | [optional] 
**Secret** | Pointer to **string** | The API secret of your exchange account. This property is required when updating the information of an Exchange Wallet. | [optional] 
**Passphrase** | Pointer to **string** | The passphrase of your exchange account. | [optional] 
**Memo** | Pointer to **string** | The memo you use when applying for the API key of your exchange account. | [optional] 
**AccountIdentify** | Pointer to **string** | The identifier of your exchange account. - For Binance, this is email address of your exchange account. - For OKX, this is the user name of your exchange account.  | [optional] 
**GaCode** | Pointer to **string** | The GA code for the exchange. | [optional] 
**MainWalletId** | Pointer to **string** | The ID of the Exchange Wallet (Main Account). | [optional] 

## Methods

### NewUpdateExchangeWalletParams

`func NewUpdateExchangeWalletParams(walletType WalletType, ) *UpdateExchangeWalletParams`

NewUpdateExchangeWalletParams instantiates a new UpdateExchangeWalletParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExchangeWalletParamsWithDefaults

`func NewUpdateExchangeWalletParamsWithDefaults() *UpdateExchangeWalletParams`

NewUpdateExchangeWalletParamsWithDefaults instantiates a new UpdateExchangeWalletParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletType

`func (o *UpdateExchangeWalletParams) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *UpdateExchangeWalletParams) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *UpdateExchangeWalletParams) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetName

`func (o *UpdateExchangeWalletParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateExchangeWalletParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateExchangeWalletParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateExchangeWalletParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApikey

`func (o *UpdateExchangeWalletParams) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *UpdateExchangeWalletParams) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *UpdateExchangeWalletParams) SetApikey(v string)`

SetApikey sets Apikey field to given value.

### HasApikey

`func (o *UpdateExchangeWalletParams) HasApikey() bool`

HasApikey returns a boolean if a field has been set.

### GetSecret

`func (o *UpdateExchangeWalletParams) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateExchangeWalletParams) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateExchangeWalletParams) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *UpdateExchangeWalletParams) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPassphrase

`func (o *UpdateExchangeWalletParams) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *UpdateExchangeWalletParams) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *UpdateExchangeWalletParams) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *UpdateExchangeWalletParams) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetMemo

`func (o *UpdateExchangeWalletParams) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *UpdateExchangeWalletParams) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *UpdateExchangeWalletParams) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *UpdateExchangeWalletParams) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetAccountIdentify

`func (o *UpdateExchangeWalletParams) GetAccountIdentify() string`

GetAccountIdentify returns the AccountIdentify field if non-nil, zero value otherwise.

### GetAccountIdentifyOk

`func (o *UpdateExchangeWalletParams) GetAccountIdentifyOk() (*string, bool)`

GetAccountIdentifyOk returns a tuple with the AccountIdentify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentify

`func (o *UpdateExchangeWalletParams) SetAccountIdentify(v string)`

SetAccountIdentify sets AccountIdentify field to given value.

### HasAccountIdentify

`func (o *UpdateExchangeWalletParams) HasAccountIdentify() bool`

HasAccountIdentify returns a boolean if a field has been set.

### GetGaCode

`func (o *UpdateExchangeWalletParams) GetGaCode() string`

GetGaCode returns the GaCode field if non-nil, zero value otherwise.

### GetGaCodeOk

`func (o *UpdateExchangeWalletParams) GetGaCodeOk() (*string, bool)`

GetGaCodeOk returns a tuple with the GaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaCode

`func (o *UpdateExchangeWalletParams) SetGaCode(v string)`

SetGaCode sets GaCode field to given value.

### HasGaCode

`func (o *UpdateExchangeWalletParams) HasGaCode() bool`

HasGaCode returns a boolean if a field has been set.

### GetMainWalletId

`func (o *UpdateExchangeWalletParams) GetMainWalletId() string`

GetMainWalletId returns the MainWalletId field if non-nil, zero value otherwise.

### GetMainWalletIdOk

`func (o *UpdateExchangeWalletParams) GetMainWalletIdOk() (*string, bool)`

GetMainWalletIdOk returns a tuple with the MainWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainWalletId

`func (o *UpdateExchangeWalletParams) SetMainWalletId(v string)`

SetMainWalletId sets MainWalletId field to given value.

### HasMainWalletId

`func (o *UpdateExchangeWalletParams) HasMainWalletId() bool`

HasMainWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


