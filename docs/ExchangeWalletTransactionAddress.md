# ExchangeWalletTransactionAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TransactionAddressType**](TransactionAddressType.md) |  | 
**Address** | Pointer to **string** | Address | [optional] 
**Memo** | Pointer to **string** | Address memo | [optional] 
**WalletId** | **string** | Unique id of the wallet. | 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**SubWalletId** | Pointer to **string** | Exchange trading account or any sub wallet info for transfer. | [optional] 

## Methods

### NewExchangeWalletTransactionAddress

`func NewExchangeWalletTransactionAddress(type_ TransactionAddressType, walletId string, exchangeId ExchangeId, ) *ExchangeWalletTransactionAddress`

NewExchangeWalletTransactionAddress instantiates a new ExchangeWalletTransactionAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeWalletTransactionAddressWithDefaults

`func NewExchangeWalletTransactionAddressWithDefaults() *ExchangeWalletTransactionAddress`

NewExchangeWalletTransactionAddressWithDefaults instantiates a new ExchangeWalletTransactionAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExchangeWalletTransactionAddress) GetType() TransactionAddressType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExchangeWalletTransactionAddress) GetTypeOk() (*TransactionAddressType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExchangeWalletTransactionAddress) SetType(v TransactionAddressType)`

SetType sets Type field to given value.


### GetAddress

`func (o *ExchangeWalletTransactionAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ExchangeWalletTransactionAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ExchangeWalletTransactionAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ExchangeWalletTransactionAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMemo

`func (o *ExchangeWalletTransactionAddress) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ExchangeWalletTransactionAddress) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ExchangeWalletTransactionAddress) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *ExchangeWalletTransactionAddress) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetWalletId

`func (o *ExchangeWalletTransactionAddress) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ExchangeWalletTransactionAddress) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ExchangeWalletTransactionAddress) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetExchangeId

`func (o *ExchangeWalletTransactionAddress) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *ExchangeWalletTransactionAddress) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *ExchangeWalletTransactionAddress) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.


### GetSubWalletId

`func (o *ExchangeWalletTransactionAddress) GetSubWalletId() string`

GetSubWalletId returns the SubWalletId field if non-nil, zero value otherwise.

### GetSubWalletIdOk

`func (o *ExchangeWalletTransactionAddress) GetSubWalletIdOk() (*string, bool)`

GetSubWalletIdOk returns a tuple with the SubWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWalletId

`func (o *ExchangeWalletTransactionAddress) SetSubWalletId(v string)`

SetSubWalletId sets SubWalletId field to given value.

### HasSubWalletId

`func (o *ExchangeWalletTransactionAddress) HasSubWalletId() bool`

HasSubWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


