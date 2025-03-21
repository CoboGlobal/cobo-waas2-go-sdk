# SubWalletAssetBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountType** | Pointer to **string** | The trading account type. | [optional] 
**AssetId** | **string** | The asset ID. An asset ID is the unique identifier of the asset held within your linked exchange account. | 
**Balance** | [**Balance**](Balance.md) |  | 

## Methods

### NewSubWalletAssetBalance

`func NewSubWalletAssetBalance(assetId string, balance Balance, ) *SubWalletAssetBalance`

NewSubWalletAssetBalance instantiates a new SubWalletAssetBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubWalletAssetBalanceWithDefaults

`func NewSubWalletAssetBalanceWithDefaults() *SubWalletAssetBalance`

NewSubWalletAssetBalanceWithDefaults instantiates a new SubWalletAssetBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountType

`func (o *SubWalletAssetBalance) GetTradingAccountType() string`

GetTradingAccountType returns the TradingAccountType field if non-nil, zero value otherwise.

### GetTradingAccountTypeOk

`func (o *SubWalletAssetBalance) GetTradingAccountTypeOk() (*string, bool)`

GetTradingAccountTypeOk returns a tuple with the TradingAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountType

`func (o *SubWalletAssetBalance) SetTradingAccountType(v string)`

SetTradingAccountType sets TradingAccountType field to given value.

### HasTradingAccountType

`func (o *SubWalletAssetBalance) HasTradingAccountType() bool`

HasTradingAccountType returns a boolean if a field has been set.

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

`func (o *SubWalletAssetBalance) GetBalance() Balance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *SubWalletAssetBalance) GetBalanceOk() (*Balance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *SubWalletAssetBalance) SetBalance(v Balance)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


