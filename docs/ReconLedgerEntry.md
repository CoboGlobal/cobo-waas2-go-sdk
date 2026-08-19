# ReconLedgerEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The transaction ID (the Cobo transaction ID you provided in &#x60;transaction_ids&#x60;). | 
**BlockTime** | Pointer to **int64** | The time when the block containing the transaction was created, in Unix timestamp format, measured in milliseconds. | [optional] 
**WalletId** | **string** | The wallet ID. | 
**Address** | **string** | The wallet address involved in this entry. | 
**TokenId** | **string** | The token ID, which is the unique identifier of a token. | 
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. | 
**Amount** | **string** | The transaction amount for this entry, expressed in the token&#39;s main unit (already divided by the token&#39;s decimals). The value is signed - positive for deposits and negative for withdrawals. | 
**BalanceAfter** | **string** | The running balance of the address for this token after this transaction, expressed in the token&#39;s main unit. | 
**TransactionHash** | Pointer to **string** | The transaction hash on the blockchain. | [optional] 
**BlockNumber** | Pointer to **int64** | The number of the block containing the transaction. | [optional] 

## Methods

### NewReconLedgerEntry

`func NewReconLedgerEntry(transactionId string, walletId string, address string, tokenId string, chainId string, amount string, balanceAfter string, ) *ReconLedgerEntry`

NewReconLedgerEntry instantiates a new ReconLedgerEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconLedgerEntryWithDefaults

`func NewReconLedgerEntryWithDefaults() *ReconLedgerEntry`

NewReconLedgerEntryWithDefaults instantiates a new ReconLedgerEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *ReconLedgerEntry) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ReconLedgerEntry) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ReconLedgerEntry) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetBlockTime

`func (o *ReconLedgerEntry) GetBlockTime() int64`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *ReconLedgerEntry) GetBlockTimeOk() (*int64, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *ReconLedgerEntry) SetBlockTime(v int64)`

SetBlockTime sets BlockTime field to given value.

### HasBlockTime

`func (o *ReconLedgerEntry) HasBlockTime() bool`

HasBlockTime returns a boolean if a field has been set.

### GetWalletId

`func (o *ReconLedgerEntry) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ReconLedgerEntry) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ReconLedgerEntry) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *ReconLedgerEntry) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ReconLedgerEntry) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ReconLedgerEntry) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTokenId

`func (o *ReconLedgerEntry) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ReconLedgerEntry) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ReconLedgerEntry) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *ReconLedgerEntry) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ReconLedgerEntry) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ReconLedgerEntry) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetAmount

`func (o *ReconLedgerEntry) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ReconLedgerEntry) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ReconLedgerEntry) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetBalanceAfter

`func (o *ReconLedgerEntry) GetBalanceAfter() string`

GetBalanceAfter returns the BalanceAfter field if non-nil, zero value otherwise.

### GetBalanceAfterOk

`func (o *ReconLedgerEntry) GetBalanceAfterOk() (*string, bool)`

GetBalanceAfterOk returns a tuple with the BalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfter

`func (o *ReconLedgerEntry) SetBalanceAfter(v string)`

SetBalanceAfter sets BalanceAfter field to given value.


### GetTransactionHash

`func (o *ReconLedgerEntry) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ReconLedgerEntry) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ReconLedgerEntry) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *ReconLedgerEntry) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetBlockNumber

`func (o *ReconLedgerEntry) GetBlockNumber() int64`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *ReconLedgerEntry) GetBlockNumberOk() (*int64, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *ReconLedgerEntry) SetBlockNumber(v int64)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *ReconLedgerEntry) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


