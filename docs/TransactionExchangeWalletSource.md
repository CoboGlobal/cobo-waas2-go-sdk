# TransactionExchangeWalletSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**TransactionSourceType**](TransactionSourceType.md) |  | 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**TradingAccountType** | Pointer to **string** | The exchange trading account or a sub-wallet ID. | [optional] 

## Methods

### NewTransactionExchangeWalletSource

`func NewTransactionExchangeWalletSource(sourceType TransactionSourceType, exchangeId ExchangeId, walletId string, ) *TransactionExchangeWalletSource`

NewTransactionExchangeWalletSource instantiates a new TransactionExchangeWalletSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionExchangeWalletSourceWithDefaults

`func NewTransactionExchangeWalletSourceWithDefaults() *TransactionExchangeWalletSource`

NewTransactionExchangeWalletSourceWithDefaults instantiates a new TransactionExchangeWalletSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TransactionExchangeWalletSource) GetSourceType() TransactionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransactionExchangeWalletSource) GetSourceTypeOk() (*TransactionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransactionExchangeWalletSource) SetSourceType(v TransactionSourceType)`

SetSourceType sets SourceType field to given value.


### GetExchangeId

`func (o *TransactionExchangeWalletSource) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *TransactionExchangeWalletSource) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *TransactionExchangeWalletSource) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.


### GetWalletId

`func (o *TransactionExchangeWalletSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionExchangeWalletSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionExchangeWalletSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetTradingAccountType

`func (o *TransactionExchangeWalletSource) GetTradingAccountType() string`

GetTradingAccountType returns the TradingAccountType field if non-nil, zero value otherwise.

### GetTradingAccountTypeOk

`func (o *TransactionExchangeWalletSource) GetTradingAccountTypeOk() (*string, bool)`

GetTradingAccountTypeOk returns a tuple with the TradingAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountType

`func (o *TransactionExchangeWalletSource) SetTradingAccountType(v string)`

SetTradingAccountType sets TradingAccountType field to given value.

### HasTradingAccountType

`func (o *TransactionExchangeWalletSource) HasTradingAccountType() bool`

HasTradingAccountType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


