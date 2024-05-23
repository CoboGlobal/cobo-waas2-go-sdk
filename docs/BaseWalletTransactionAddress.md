# BaseWalletTransactionAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TransactionAddressType**](TransactionAddressType.md) |  | 
**Address** | Pointer to **string** | Address | [optional] 
**Memo** | Pointer to **string** | Address memo | [optional] 
**WalletId** | **string** | Unique id of the wallet. | 

## Methods

### NewBaseWalletTransactionAddress

`func NewBaseWalletTransactionAddress(type_ TransactionAddressType, walletId string, ) *BaseWalletTransactionAddress`

NewBaseWalletTransactionAddress instantiates a new BaseWalletTransactionAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseWalletTransactionAddressWithDefaults

`func NewBaseWalletTransactionAddressWithDefaults() *BaseWalletTransactionAddress`

NewBaseWalletTransactionAddressWithDefaults instantiates a new BaseWalletTransactionAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BaseWalletTransactionAddress) GetType() TransactionAddressType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseWalletTransactionAddress) GetTypeOk() (*TransactionAddressType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseWalletTransactionAddress) SetType(v TransactionAddressType)`

SetType sets Type field to given value.


### GetAddress

`func (o *BaseWalletTransactionAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BaseWalletTransactionAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BaseWalletTransactionAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BaseWalletTransactionAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMemo

`func (o *BaseWalletTransactionAddress) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *BaseWalletTransactionAddress) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *BaseWalletTransactionAddress) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *BaseWalletTransactionAddress) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetWalletId

`func (o *BaseWalletTransactionAddress) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *BaseWalletTransactionAddress) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *BaseWalletTransactionAddress) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


