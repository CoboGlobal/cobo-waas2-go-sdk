# TransactionTransferToWalletDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**TradingAccountType** | Pointer to **string** | The trading account type. | [optional] 
**ExchangeId** | Pointer to [**ExchangeId**](ExchangeId.md) |  | [optional] 
**Amount** | **string** | The transfer amount. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | 

## Methods

### NewTransactionTransferToWalletDestination

`func NewTransactionTransferToWalletDestination(destinationType TransactionDestinationType, walletId string, amount string, ) *TransactionTransferToWalletDestination`

NewTransactionTransferToWalletDestination instantiates a new TransactionTransferToWalletDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTransferToWalletDestinationWithDefaults

`func NewTransactionTransferToWalletDestinationWithDefaults() *TransactionTransferToWalletDestination`

NewTransactionTransferToWalletDestinationWithDefaults instantiates a new TransactionTransferToWalletDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionTransferToWalletDestination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionTransferToWalletDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionTransferToWalletDestination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetWalletId

`func (o *TransactionTransferToWalletDestination) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionTransferToWalletDestination) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionTransferToWalletDestination) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetTradingAccountType

`func (o *TransactionTransferToWalletDestination) GetTradingAccountType() string`

GetTradingAccountType returns the TradingAccountType field if non-nil, zero value otherwise.

### GetTradingAccountTypeOk

`func (o *TransactionTransferToWalletDestination) GetTradingAccountTypeOk() (*string, bool)`

GetTradingAccountTypeOk returns a tuple with the TradingAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountType

`func (o *TransactionTransferToWalletDestination) SetTradingAccountType(v string)`

SetTradingAccountType sets TradingAccountType field to given value.

### HasTradingAccountType

`func (o *TransactionTransferToWalletDestination) HasTradingAccountType() bool`

HasTradingAccountType returns a boolean if a field has been set.

### GetExchangeId

`func (o *TransactionTransferToWalletDestination) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *TransactionTransferToWalletDestination) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *TransactionTransferToWalletDestination) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.

### HasExchangeId

`func (o *TransactionTransferToWalletDestination) HasExchangeId() bool`

HasExchangeId returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionTransferToWalletDestination) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionTransferToWalletDestination) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionTransferToWalletDestination) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


