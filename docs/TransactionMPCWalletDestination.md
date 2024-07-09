# TransactionMPCWalletDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**TokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
**AssetId** | Pointer to **string** | (This concept applies to Exchange Wallets only) The asset ID. An asset is a digital representation of a valuable resource on a blockchain network. Exchange Wallets group your holdings by asset, even if the same asset exists on different blockchains. For example, if your Exchange Wallet has 1 USDT on Ethereum and 1 USDT on TRON, then your asset balance is 2 USDT. | [optional] 
**Address** | **string** | The destination address. | 
**Memo** | Pointer to **string** | The memo that identifies a transaction in order to credit the correct account. For transfers out of Cobo Portal, it is highly recommended to include a memo for the chains such as XRP, EOS, XLM, IOST, BNB_BNB, ATOM, LUNA, and TON. | [optional] 
**Amount** | **string** | The transaction amount. For example, if you trade 1.5 ETH, then the amount is &#x60;1.5&#x60;.  | 

## Methods

### NewTransactionMPCWalletDestination

`func NewTransactionMPCWalletDestination(destinationType TransactionDestinationType, walletId string, tokenId string, address string, amount string, ) *TransactionMPCWalletDestination`

NewTransactionMPCWalletDestination instantiates a new TransactionMPCWalletDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionMPCWalletDestinationWithDefaults

`func NewTransactionMPCWalletDestinationWithDefaults() *TransactionMPCWalletDestination`

NewTransactionMPCWalletDestinationWithDefaults instantiates a new TransactionMPCWalletDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionMPCWalletDestination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionMPCWalletDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionMPCWalletDestination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetWalletId

`func (o *TransactionMPCWalletDestination) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionMPCWalletDestination) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionMPCWalletDestination) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetTokenId

`func (o *TransactionMPCWalletDestination) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionMPCWalletDestination) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionMPCWalletDestination) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAssetId

`func (o *TransactionMPCWalletDestination) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *TransactionMPCWalletDestination) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *TransactionMPCWalletDestination) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *TransactionMPCWalletDestination) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAddress

`func (o *TransactionMPCWalletDestination) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionMPCWalletDestination) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionMPCWalletDestination) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMemo

`func (o *TransactionMPCWalletDestination) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionMPCWalletDestination) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionMPCWalletDestination) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransactionMPCWalletDestination) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionMPCWalletDestination) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionMPCWalletDestination) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionMPCWalletDestination) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


