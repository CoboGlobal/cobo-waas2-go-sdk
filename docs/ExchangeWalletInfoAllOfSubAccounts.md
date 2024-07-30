# ExchangeWalletInfoAllOfSubAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The wallet ID of the Sub Account. This property is returned only if you are creating or querying an Exchange Wallet (Main Account). | 
**AccountId** | **string** | The Sub Account ID. It can be an email address, a user name, or a custom account ID. This property is returned only if you are creating or querying an Exchange Wallet (Main Account). | 

## Methods

### NewExchangeWalletInfoAllOfSubAccounts

`func NewExchangeWalletInfoAllOfSubAccounts(walletId string, accountId string, ) *ExchangeWalletInfoAllOfSubAccounts`

NewExchangeWalletInfoAllOfSubAccounts instantiates a new ExchangeWalletInfoAllOfSubAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeWalletInfoAllOfSubAccountsWithDefaults

`func NewExchangeWalletInfoAllOfSubAccountsWithDefaults() *ExchangeWalletInfoAllOfSubAccounts`

NewExchangeWalletInfoAllOfSubAccountsWithDefaults instantiates a new ExchangeWalletInfoAllOfSubAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *ExchangeWalletInfoAllOfSubAccounts) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ExchangeWalletInfoAllOfSubAccounts) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ExchangeWalletInfoAllOfSubAccounts) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAccountId

`func (o *ExchangeWalletInfoAllOfSubAccounts) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ExchangeWalletInfoAllOfSubAccounts) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ExchangeWalletInfoAllOfSubAccounts) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


