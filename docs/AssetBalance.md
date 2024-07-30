# AssetBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** | (This concept applies to Exchange Wallets only) The asset ID. An asset ID is the unique identifier of the asset held within your linked exchange account. | 
**Balance** | [**TokenBalanceBalance**](TokenBalanceBalance.md) |  | 

## Methods

### NewAssetBalance

`func NewAssetBalance(assetId string, balance TokenBalanceBalance, ) *AssetBalance`

NewAssetBalance instantiates a new AssetBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetBalanceWithDefaults

`func NewAssetBalanceWithDefaults() *AssetBalance`

NewAssetBalanceWithDefaults instantiates a new AssetBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *AssetBalance) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetBalance) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetBalance) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetBalance

`func (o *AssetBalance) GetBalance() TokenBalanceBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AssetBalance) GetBalanceOk() (*TokenBalanceBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AssetBalance) SetBalance(v TokenBalanceBalance)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


