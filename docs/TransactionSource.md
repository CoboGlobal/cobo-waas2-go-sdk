# TransactionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TransactionAddressType**](TransactionAddressType.md) |  | 
**Address** | Pointer to **string** | Address | [optional] 
**Memo** | Pointer to **string** | Address memo | [optional] 
**WalletId** | **string** | Unique id of the wallet. | 
**MpcUsedKeyGroup** | Pointer to [**MpcSigningGroup**](MpcSigningGroup.md) |  | [optional] 
**Delegate** | [**SafeTransactionAddressAllOfDelegate**](SafeTransactionAddressAllOfDelegate.md) |  | 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**SubWalletId** | Pointer to **string** | Exchange trading account or any sub wallet info for transfer. | [optional] 

## Methods

### NewTransactionSource

`func NewTransactionSource(type_ TransactionAddressType, walletId string, delegate SafeTransactionAddressAllOfDelegate, exchangeId ExchangeId, ) *TransactionSource`

NewTransactionSource instantiates a new TransactionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSourceWithDefaults

`func NewTransactionSourceWithDefaults() *TransactionSource`

NewTransactionSourceWithDefaults instantiates a new TransactionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionSource) GetType() TransactionAddressType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionSource) GetTypeOk() (*TransactionAddressType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionSource) SetType(v TransactionAddressType)`

SetType sets Type field to given value.


### GetAddress

`func (o *TransactionSource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionSource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionSource) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransactionSource) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMemo

`func (o *TransactionSource) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionSource) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionSource) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransactionSource) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetWalletId

`func (o *TransactionSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetMpcUsedKeyGroup

`func (o *TransactionSource) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *TransactionSource) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *TransactionSource) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.

### HasMpcUsedKeyGroup

`func (o *TransactionSource) HasMpcUsedKeyGroup() bool`

HasMpcUsedKeyGroup returns a boolean if a field has been set.

### GetDelegate

`func (o *TransactionSource) GetDelegate() SafeTransactionAddressAllOfDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *TransactionSource) GetDelegateOk() (*SafeTransactionAddressAllOfDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *TransactionSource) SetDelegate(v SafeTransactionAddressAllOfDelegate)`

SetDelegate sets Delegate field to given value.


### GetExchangeId

`func (o *TransactionSource) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *TransactionSource) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *TransactionSource) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.


### GetSubWalletId

`func (o *TransactionSource) GetSubWalletId() string`

GetSubWalletId returns the SubWalletId field if non-nil, zero value otherwise.

### GetSubWalletIdOk

`func (o *TransactionSource) GetSubWalletIdOk() (*string, bool)`

GetSubWalletIdOk returns a tuple with the SubWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWalletId

`func (o *TransactionSource) SetSubWalletId(v string)`

SetSubWalletId sets SubWalletId field to given value.

### HasSubWalletId

`func (o *TransactionSource) HasSubWalletId() bool`

HasSubWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


