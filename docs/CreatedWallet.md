# CreatedWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**WalletType** | **string** |  | 
**WalletSubtype** | **string** |  | 
**VaultId** | **string** | The owning mpc vault id of the mpc wallet. | 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**Apikey** | **string** | The API Key for the exchange. | 
**Secret** | **string** | The API Secret for the exchange. | 
**Passphrase** | Pointer to **string** | The passphrase for the exchange. | [optional] 
**GaCode** | Pointer to **string** | The ga_code for the exchange. | [optional] 
**SubAccountIds** | Pointer to **[]string** | The unique identifier associated with the exchange sub-account. It can be an email address, username, or a custom account ID. | [optional] 

## Methods

### NewCreatedWallet

`func NewCreatedWallet(name string, walletType string, walletSubtype string, vaultId string, exchangeId ExchangeId, apikey string, secret string, ) *CreatedWallet`

NewCreatedWallet instantiates a new CreatedWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedWalletWithDefaults

`func NewCreatedWalletWithDefaults() *CreatedWallet`

NewCreatedWalletWithDefaults instantiates a new CreatedWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatedWallet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatedWallet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatedWallet) SetName(v string)`

SetName sets Name field to given value.


### GetWalletType

`func (o *CreatedWallet) GetWalletType() string`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *CreatedWallet) GetWalletTypeOk() (*string, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *CreatedWallet) SetWalletType(v string)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *CreatedWallet) GetWalletSubtype() string`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *CreatedWallet) GetWalletSubtypeOk() (*string, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *CreatedWallet) SetWalletSubtype(v string)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetVaultId

`func (o *CreatedWallet) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *CreatedWallet) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *CreatedWallet) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.


### GetExchangeId

`func (o *CreatedWallet) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *CreatedWallet) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *CreatedWallet) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.


### GetApikey

`func (o *CreatedWallet) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *CreatedWallet) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *CreatedWallet) SetApikey(v string)`

SetApikey sets Apikey field to given value.


### GetSecret

`func (o *CreatedWallet) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreatedWallet) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreatedWallet) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetPassphrase

`func (o *CreatedWallet) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *CreatedWallet) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *CreatedWallet) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *CreatedWallet) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetGaCode

`func (o *CreatedWallet) GetGaCode() string`

GetGaCode returns the GaCode field if non-nil, zero value otherwise.

### GetGaCodeOk

`func (o *CreatedWallet) GetGaCodeOk() (*string, bool)`

GetGaCodeOk returns a tuple with the GaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaCode

`func (o *CreatedWallet) SetGaCode(v string)`

SetGaCode sets GaCode field to given value.

### HasGaCode

`func (o *CreatedWallet) HasGaCode() bool`

HasGaCode returns a boolean if a field has been set.

### GetSubAccountIds

`func (o *CreatedWallet) GetSubAccountIds() []string`

GetSubAccountIds returns the SubAccountIds field if non-nil, zero value otherwise.

### GetSubAccountIdsOk

`func (o *CreatedWallet) GetSubAccountIdsOk() (*[]string, bool)`

GetSubAccountIdsOk returns a tuple with the SubAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountIds

`func (o *CreatedWallet) SetSubAccountIds(v []string)`

SetSubAccountIds sets SubAccountIds field to given value.

### HasSubAccountIds

`func (o *CreatedWallet) HasSubAccountIds() bool`

HasSubAccountIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


