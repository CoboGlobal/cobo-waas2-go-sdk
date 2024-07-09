# TransactionExchangeWalletDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**SubWalletId** | Pointer to **string** | The exchange trading account or the sub-wallet ID. | [optional] 
**TokenId** | Pointer to **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | [optional] 
**AssetId** | **string** | (This concept applies to Exchange Wallets only) The asset ID. An asset is a digital representation of a valuable resource on a blockchain network. Exchange Wallets group your holdings by asset, even if the same asset exists on different blockchains. For example, if your Exchange Wallet has 1 USDT on Ethereum and 1 USDT on TRON, then your asset balance is 2 USDT. | 
**Address** | Pointer to **string** | The destination address. | [optional] 
**Memo** | Pointer to **string** | The memo that identifies a transaction in order to credit the correct account. For transfers out of Cobo Portal, it is highly recommended to include a memo for the chains such as XRP, EOS, XLM, IOST, BNB_BNB, ATOM, LUNA, and TON. | [optional] 
**Amount** | **string** | The quantity of the token in the transaction. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | 

## Methods

### NewTransactionExchangeWalletDestination

`func NewTransactionExchangeWalletDestination(destinationType TransactionDestinationType, exchangeId ExchangeId, walletId string, assetId string, amount string, ) *TransactionExchangeWalletDestination`

NewTransactionExchangeWalletDestination instantiates a new TransactionExchangeWalletDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionExchangeWalletDestinationWithDefaults

`func NewTransactionExchangeWalletDestinationWithDefaults() *TransactionExchangeWalletDestination`

NewTransactionExchangeWalletDestinationWithDefaults instantiates a new TransactionExchangeWalletDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionExchangeWalletDestination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionExchangeWalletDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionExchangeWalletDestination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetExchangeId

`func (o *TransactionExchangeWalletDestination) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *TransactionExchangeWalletDestination) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *TransactionExchangeWalletDestination) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.


### GetWalletId

`func (o *TransactionExchangeWalletDestination) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionExchangeWalletDestination) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionExchangeWalletDestination) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetSubWalletId

`func (o *TransactionExchangeWalletDestination) GetSubWalletId() string`

GetSubWalletId returns the SubWalletId field if non-nil, zero value otherwise.

### GetSubWalletIdOk

`func (o *TransactionExchangeWalletDestination) GetSubWalletIdOk() (*string, bool)`

GetSubWalletIdOk returns a tuple with the SubWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWalletId

`func (o *TransactionExchangeWalletDestination) SetSubWalletId(v string)`

SetSubWalletId sets SubWalletId field to given value.

### HasSubWalletId

`func (o *TransactionExchangeWalletDestination) HasSubWalletId() bool`

HasSubWalletId returns a boolean if a field has been set.

### GetTokenId

`func (o *TransactionExchangeWalletDestination) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionExchangeWalletDestination) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionExchangeWalletDestination) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TransactionExchangeWalletDestination) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetAssetId

`func (o *TransactionExchangeWalletDestination) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *TransactionExchangeWalletDestination) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *TransactionExchangeWalletDestination) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetAddress

`func (o *TransactionExchangeWalletDestination) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionExchangeWalletDestination) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionExchangeWalletDestination) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransactionExchangeWalletDestination) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMemo

`func (o *TransactionExchangeWalletDestination) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionExchangeWalletDestination) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionExchangeWalletDestination) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransactionExchangeWalletDestination) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionExchangeWalletDestination) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionExchangeWalletDestination) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionExchangeWalletDestination) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


