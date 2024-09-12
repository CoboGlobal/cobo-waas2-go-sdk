# WalletBalanceSnapshotRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The wallet ID. | 
**WalletType** | Pointer to [**WalletType**](WalletType.md) |  | [optional] 
**WalletName** | Pointer to **string** | The wallet name. | [optional] 
**TokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](/v2/api-references/wallets/list-enabled-tokens). | 
**Balance** | **string** | The balance of the token. | 

## Methods

### NewWalletBalanceSnapshotRecord

`func NewWalletBalanceSnapshotRecord(walletId string, tokenId string, balance string, ) *WalletBalanceSnapshotRecord`

NewWalletBalanceSnapshotRecord instantiates a new WalletBalanceSnapshotRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletBalanceSnapshotRecordWithDefaults

`func NewWalletBalanceSnapshotRecordWithDefaults() *WalletBalanceSnapshotRecord`

NewWalletBalanceSnapshotRecordWithDefaults instantiates a new WalletBalanceSnapshotRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *WalletBalanceSnapshotRecord) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *WalletBalanceSnapshotRecord) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *WalletBalanceSnapshotRecord) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetWalletType

`func (o *WalletBalanceSnapshotRecord) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *WalletBalanceSnapshotRecord) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *WalletBalanceSnapshotRecord) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *WalletBalanceSnapshotRecord) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.

### GetWalletName

`func (o *WalletBalanceSnapshotRecord) GetWalletName() string`

GetWalletName returns the WalletName field if non-nil, zero value otherwise.

### GetWalletNameOk

`func (o *WalletBalanceSnapshotRecord) GetWalletNameOk() (*string, bool)`

GetWalletNameOk returns a tuple with the WalletName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletName

`func (o *WalletBalanceSnapshotRecord) SetWalletName(v string)`

SetWalletName sets WalletName field to given value.

### HasWalletName

`func (o *WalletBalanceSnapshotRecord) HasWalletName() bool`

HasWalletName returns a boolean if a field has been set.

### GetTokenId

`func (o *WalletBalanceSnapshotRecord) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *WalletBalanceSnapshotRecord) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *WalletBalanceSnapshotRecord) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetBalance

`func (o *WalletBalanceSnapshotRecord) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *WalletBalanceSnapshotRecord) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *WalletBalanceSnapshotRecord) SetBalance(v string)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


