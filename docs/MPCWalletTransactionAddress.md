# MPCWalletTransactionAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TransactionAddressType**](TransactionAddressType.md) |  | 
**Address** | Pointer to **string** | Address | [optional] 
**Memo** | Pointer to **string** | Address memo | [optional] 
**WalletId** | **string** | Unique id of the wallet. | 
**MpcUsedKeyGroup** | Pointer to [**MpcSigningGroup**](MpcSigningGroup.md) |  | [optional] 

## Methods

### NewMPCWalletTransactionAddress

`func NewMPCWalletTransactionAddress(type_ TransactionAddressType, walletId string, ) *MPCWalletTransactionAddress`

NewMPCWalletTransactionAddress instantiates a new MPCWalletTransactionAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMPCWalletTransactionAddressWithDefaults

`func NewMPCWalletTransactionAddressWithDefaults() *MPCWalletTransactionAddress`

NewMPCWalletTransactionAddressWithDefaults instantiates a new MPCWalletTransactionAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MPCWalletTransactionAddress) GetType() TransactionAddressType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MPCWalletTransactionAddress) GetTypeOk() (*TransactionAddressType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MPCWalletTransactionAddress) SetType(v TransactionAddressType)`

SetType sets Type field to given value.


### GetAddress

`func (o *MPCWalletTransactionAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MPCWalletTransactionAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MPCWalletTransactionAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MPCWalletTransactionAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMemo

`func (o *MPCWalletTransactionAddress) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *MPCWalletTransactionAddress) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *MPCWalletTransactionAddress) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *MPCWalletTransactionAddress) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetWalletId

`func (o *MPCWalletTransactionAddress) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *MPCWalletTransactionAddress) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *MPCWalletTransactionAddress) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetMpcUsedKeyGroup

`func (o *MPCWalletTransactionAddress) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *MPCWalletTransactionAddress) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *MPCWalletTransactionAddress) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.

### HasMpcUsedKeyGroup

`func (o *MPCWalletTransactionAddress) HasMpcUsedKeyGroup() bool`

HasMpcUsedKeyGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


