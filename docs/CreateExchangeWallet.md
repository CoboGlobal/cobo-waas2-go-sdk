# CreateExchangeWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The wallet name. | 
**WalletType** | **string** | The wallet type. | 
**WalletSubtype** | **string** | The wallet subtype. | 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**Apikey** | **string** | The API key of your exchange account. | 
**Secret** | **string** | The API secret of your exchange account. | 
**Passphrase** | Pointer to **string** | The passphrase of your exchange account. | [optional] 
**GaCode** | Pointer to **string** | The ga_code for the exchange. | [optional] 
**SubAccountIds** | Pointer to **[]string** | The Sub Account ID. It can be an email address, a user name, or a custom account ID. | [optional] 

## Methods

### NewCreateExchangeWallet

`func NewCreateExchangeWallet(name string, walletType string, walletSubtype string, exchangeId ExchangeId, apikey string, secret string, ) *CreateExchangeWallet`

NewCreateExchangeWallet instantiates a new CreateExchangeWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExchangeWalletWithDefaults

`func NewCreateExchangeWalletWithDefaults() *CreateExchangeWallet`

NewCreateExchangeWalletWithDefaults instantiates a new CreateExchangeWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateExchangeWallet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateExchangeWallet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateExchangeWallet) SetName(v string)`

SetName sets Name field to given value.


### GetWalletType

`func (o *CreateExchangeWallet) GetWalletType() string`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *CreateExchangeWallet) GetWalletTypeOk() (*string, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *CreateExchangeWallet) SetWalletType(v string)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *CreateExchangeWallet) GetWalletSubtype() string`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *CreateExchangeWallet) GetWalletSubtypeOk() (*string, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *CreateExchangeWallet) SetWalletSubtype(v string)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetExchangeId

`func (o *CreateExchangeWallet) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *CreateExchangeWallet) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *CreateExchangeWallet) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.


### GetApikey

`func (o *CreateExchangeWallet) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *CreateExchangeWallet) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *CreateExchangeWallet) SetApikey(v string)`

SetApikey sets Apikey field to given value.


### GetSecret

`func (o *CreateExchangeWallet) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateExchangeWallet) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateExchangeWallet) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetPassphrase

`func (o *CreateExchangeWallet) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *CreateExchangeWallet) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *CreateExchangeWallet) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *CreateExchangeWallet) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetGaCode

`func (o *CreateExchangeWallet) GetGaCode() string`

GetGaCode returns the GaCode field if non-nil, zero value otherwise.

### GetGaCodeOk

`func (o *CreateExchangeWallet) GetGaCodeOk() (*string, bool)`

GetGaCodeOk returns a tuple with the GaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaCode

`func (o *CreateExchangeWallet) SetGaCode(v string)`

SetGaCode sets GaCode field to given value.

### HasGaCode

`func (o *CreateExchangeWallet) HasGaCode() bool`

HasGaCode returns a boolean if a field has been set.

### GetSubAccountIds

`func (o *CreateExchangeWallet) GetSubAccountIds() []string`

GetSubAccountIds returns the SubAccountIds field if non-nil, zero value otherwise.

### GetSubAccountIdsOk

`func (o *CreateExchangeWallet) GetSubAccountIdsOk() (*[]string, bool)`

GetSubAccountIdsOk returns a tuple with the SubAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountIds

`func (o *CreateExchangeWallet) SetSubAccountIds(v []string)`

SetSubAccountIds sets SubAccountIds field to given value.

### HasSubAccountIds

`func (o *CreateExchangeWallet) HasSubAccountIds() bool`

HasSubAccountIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


