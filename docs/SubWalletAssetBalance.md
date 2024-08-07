# SubWalletAssetBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubWalletId** | Pointer to **string** | The exchange trading account or a sub-wallet ID. | [optional] 
**AssetId** | **string** | (This concept applies to Exchange Wallets only) The asset ID. An asset ID is the unique identifier of the asset held within your linked exchange account. | 
**Balance** | [**TokenBalanceBalance**](TokenBalanceBalance.md) |  | 

## Methods

### NewSubWalletAssetBalance

`func NewSubWalletAssetBalance(assetId string, balance TokenBalanceBalance, ) *SubWalletAssetBalance`

NewSubWalletAssetBalance instantiates a new SubWalletAssetBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubWalletAssetBalanceWithDefaults

`func NewSubWalletAssetBalanceWithDefaults() *SubWalletAssetBalance`

NewSubWalletAssetBalanceWithDefaults instantiates a new SubWalletAssetBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubWalletId

`func (o *SubWalletAssetBalance) GetSubWalletId() string`

GetSubWalletId returns the SubWalletId field if non-nil, zero value otherwise.

### GetSubWalletIdOk

`func (o *SubWalletAssetBalance) GetSubWalletIdOk() (*string, bool)`

GetSubWalletIdOk returns a tuple with the SubWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWalletId

`func (o *SubWalletAssetBalance) SetSubWalletId(v string)`

SetSubWalletId sets SubWalletId field to given value.

### HasSubWalletId

`func (o *SubWalletAssetBalance) HasSubWalletId() bool`

HasSubWalletId returns a boolean if a field has been set.

### GetAssetId

`func (o *SubWalletAssetBalance) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *SubWalletAssetBalance) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *SubWalletAssetBalance) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetBalance

`func (o *SubWalletAssetBalance) GetBalance() TokenBalanceBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *SubWalletAssetBalance) GetBalanceOk() (*TokenBalanceBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *SubWalletAssetBalance) SetBalance(v TokenBalanceBalance)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


