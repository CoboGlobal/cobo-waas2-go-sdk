# BookkeepingRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The wallet ID. | 
**WalletType** | Pointer to [**WalletType**](WalletType.md) |  | [optional] 
**WalletName** | Pointer to **string** | The wallet name. | [optional] 
**TransactionTimestamp** | **int64** | The time when the bookkeeping was created, in Unix timestamp format, measured in milliseconds. | 
**TokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](/v2/api-references/wallets/list-enabled-tokens). | 
**Type** | **string** | The bookkeeping type. | 
**Amount** | **string** | The amount of the bookkeeping. | 
**Balance** | **string** | The post-balance of the tx. | 
**FromAddress** | Pointer to **string** | The from address. | [optional] 
**ToAddress** | Pointer to **string** | The to address. | [optional] 
**TransactionHash** | Pointer to **string** | The transaction hash. | [optional] 

## Methods

### NewBookkeepingRecord

`func NewBookkeepingRecord(walletId string, transactionTimestamp int64, tokenId string, type_ string, amount string, balance string, ) *BookkeepingRecord`

NewBookkeepingRecord instantiates a new BookkeepingRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookkeepingRecordWithDefaults

`func NewBookkeepingRecordWithDefaults() *BookkeepingRecord`

NewBookkeepingRecordWithDefaults instantiates a new BookkeepingRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *BookkeepingRecord) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *BookkeepingRecord) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *BookkeepingRecord) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetWalletType

`func (o *BookkeepingRecord) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *BookkeepingRecord) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *BookkeepingRecord) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *BookkeepingRecord) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.

### GetWalletName

`func (o *BookkeepingRecord) GetWalletName() string`

GetWalletName returns the WalletName field if non-nil, zero value otherwise.

### GetWalletNameOk

`func (o *BookkeepingRecord) GetWalletNameOk() (*string, bool)`

GetWalletNameOk returns a tuple with the WalletName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletName

`func (o *BookkeepingRecord) SetWalletName(v string)`

SetWalletName sets WalletName field to given value.

### HasWalletName

`func (o *BookkeepingRecord) HasWalletName() bool`

HasWalletName returns a boolean if a field has been set.

### GetTransactionTimestamp

`func (o *BookkeepingRecord) GetTransactionTimestamp() int64`

GetTransactionTimestamp returns the TransactionTimestamp field if non-nil, zero value otherwise.

### GetTransactionTimestampOk

`func (o *BookkeepingRecord) GetTransactionTimestampOk() (*int64, bool)`

GetTransactionTimestampOk returns a tuple with the TransactionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTimestamp

`func (o *BookkeepingRecord) SetTransactionTimestamp(v int64)`

SetTransactionTimestamp sets TransactionTimestamp field to given value.


### GetTokenId

`func (o *BookkeepingRecord) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *BookkeepingRecord) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *BookkeepingRecord) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetType

`func (o *BookkeepingRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BookkeepingRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BookkeepingRecord) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *BookkeepingRecord) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BookkeepingRecord) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BookkeepingRecord) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetBalance

`func (o *BookkeepingRecord) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BookkeepingRecord) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BookkeepingRecord) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetFromAddress

`func (o *BookkeepingRecord) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *BookkeepingRecord) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *BookkeepingRecord) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *BookkeepingRecord) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetToAddress

`func (o *BookkeepingRecord) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *BookkeepingRecord) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *BookkeepingRecord) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *BookkeepingRecord) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetTransactionHash

`func (o *BookkeepingRecord) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *BookkeepingRecord) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *BookkeepingRecord) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *BookkeepingRecord) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


