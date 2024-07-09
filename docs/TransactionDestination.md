# TransactionDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**TokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
**AssetId** | **string** | (This concept applies to Exchange Wallets only) The asset ID. An asset is a digital representation of a valuable resource on a blockchain network. Exchange Wallets group your holdings by asset, even if the same asset exists on different blockchains. For example, if your Exchange Wallet has 1 USDT on Ethereum and 1 USDT on TRON, then your asset balance is 2 USDT. | 
**AccountOutput** | Pointer to [**TransactionAddressDestinationAccountOutput**](TransactionAddressDestinationAccountOutput.md) |  | [optional] 
**UtxoOutputs** | Pointer to [**TransactionAddressDestinationUtxoOutputs**](TransactionAddressDestinationUtxoOutputs.md) |  | [optional] 
**Address** | **string** | The destination address. | 
**Value** | **string** | The quantity of the token in the transaction. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | 
**Calldata** | **string** | The data that is used to invoke a specific function or method within the specified contract at the destination address.  | 
**Message** | Pointer to **string** | The raw data to be signed that are encoded in Base64. | [optional] 
**StructuredData** | Pointer to **string** | The structured data to be signed that are encoded in JSON. | [optional] 
**WalletId** | **string** | The wallet ID. | 
**Memo** | Pointer to **string** | The memo that identifies a transaction in order to credit the correct account. For transfers out of Cobo Portal, it is highly recommended to include a memo for the chains such as XRP, EOS, XLM, IOST, BNB_BNB, ATOM, LUNA, and TON. | [optional] 
**Amount** | **string** | The quantity of the token in the transaction. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**SubWalletId** | Pointer to **string** | The exchange trading account or the sub-wallet ID. | [optional] 

## Methods

### NewTransactionDestination

`func NewTransactionDestination(destinationType TransactionDestinationType, tokenId string, assetId string, address string, value string, calldata string, walletId string, amount string, exchangeId ExchangeId, ) *TransactionDestination`

NewTransactionDestination instantiates a new TransactionDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDestinationWithDefaults

`func NewTransactionDestinationWithDefaults() *TransactionDestination`

NewTransactionDestinationWithDefaults instantiates a new TransactionDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionDestination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionDestination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetTokenId

`func (o *TransactionDestination) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionDestination) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionDestination) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAssetId

`func (o *TransactionDestination) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *TransactionDestination) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *TransactionDestination) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetAccountOutput

`func (o *TransactionDestination) GetAccountOutput() TransactionAddressDestinationAccountOutput`

GetAccountOutput returns the AccountOutput field if non-nil, zero value otherwise.

### GetAccountOutputOk

`func (o *TransactionDestination) GetAccountOutputOk() (*TransactionAddressDestinationAccountOutput, bool)`

GetAccountOutputOk returns a tuple with the AccountOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOutput

`func (o *TransactionDestination) SetAccountOutput(v TransactionAddressDestinationAccountOutput)`

SetAccountOutput sets AccountOutput field to given value.

### HasAccountOutput

`func (o *TransactionDestination) HasAccountOutput() bool`

HasAccountOutput returns a boolean if a field has been set.

### GetUtxoOutputs

`func (o *TransactionDestination) GetUtxoOutputs() TransactionAddressDestinationUtxoOutputs`

GetUtxoOutputs returns the UtxoOutputs field if non-nil, zero value otherwise.

### GetUtxoOutputsOk

`func (o *TransactionDestination) GetUtxoOutputsOk() (*TransactionAddressDestinationUtxoOutputs, bool)`

GetUtxoOutputsOk returns a tuple with the UtxoOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoOutputs

`func (o *TransactionDestination) SetUtxoOutputs(v TransactionAddressDestinationUtxoOutputs)`

SetUtxoOutputs sets UtxoOutputs field to given value.

### HasUtxoOutputs

`func (o *TransactionDestination) HasUtxoOutputs() bool`

HasUtxoOutputs returns a boolean if a field has been set.

### GetAddress

`func (o *TransactionDestination) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionDestination) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionDestination) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetValue

`func (o *TransactionDestination) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TransactionDestination) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TransactionDestination) SetValue(v string)`

SetValue sets Value field to given value.


### GetCalldata

`func (o *TransactionDestination) GetCalldata() string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *TransactionDestination) GetCalldataOk() (*string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *TransactionDestination) SetCalldata(v string)`

SetCalldata sets Calldata field to given value.


### GetMessage

`func (o *TransactionDestination) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransactionDestination) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransactionDestination) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TransactionDestination) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStructuredData

`func (o *TransactionDestination) GetStructuredData() string`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *TransactionDestination) GetStructuredDataOk() (*string, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *TransactionDestination) SetStructuredData(v string)`

SetStructuredData sets StructuredData field to given value.

### HasStructuredData

`func (o *TransactionDestination) HasStructuredData() bool`

HasStructuredData returns a boolean if a field has been set.

### GetWalletId

`func (o *TransactionDestination) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionDestination) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionDestination) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetMemo

`func (o *TransactionDestination) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionDestination) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionDestination) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransactionDestination) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionDestination) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionDestination) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionDestination) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetExchangeId

`func (o *TransactionDestination) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *TransactionDestination) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *TransactionDestination) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.


### GetSubWalletId

`func (o *TransactionDestination) GetSubWalletId() string`

GetSubWalletId returns the SubWalletId field if non-nil, zero value otherwise.

### GetSubWalletIdOk

`func (o *TransactionDestination) GetSubWalletIdOk() (*string, bool)`

GetSubWalletIdOk returns a tuple with the SubWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWalletId

`func (o *TransactionDestination) SetSubWalletId(v string)`

SetSubWalletId sets SubWalletId field to given value.

### HasSubWalletId

`func (o *TransactionDestination) HasSubWalletId() bool`

HasSubWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


