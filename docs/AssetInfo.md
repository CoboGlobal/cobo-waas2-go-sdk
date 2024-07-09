# AssetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** | (This concept applies to Exchange Wallets only) The asset ID. An asset is a digital representation of a valuable resource on a blockchain network. Exchange Wallets group your holdings by asset, even if the same asset exists on different blockchains. For example, if your Exchange Wallet has 1 USDT on Ethereum and 1 USDT on TRON, then your asset balance is 2 USDT. | 
**DisplayCode** | Pointer to **string** | The asset symbol. You can use the value for display purposes. | [optional] 
**Description** | Pointer to **string** | The description of the asset. | [optional] 
**IconUrl** | Pointer to **string** | The URL of the asset icon. | [optional] 

## Methods

### NewAssetInfo

`func NewAssetInfo(assetId string, ) *AssetInfo`

NewAssetInfo instantiates a new AssetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetInfoWithDefaults

`func NewAssetInfoWithDefaults() *AssetInfo`

NewAssetInfoWithDefaults instantiates a new AssetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *AssetInfo) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetInfo) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetInfo) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetDisplayCode

`func (o *AssetInfo) GetDisplayCode() string`

GetDisplayCode returns the DisplayCode field if non-nil, zero value otherwise.

### GetDisplayCodeOk

`func (o *AssetInfo) GetDisplayCodeOk() (*string, bool)`

GetDisplayCodeOk returns a tuple with the DisplayCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCode

`func (o *AssetInfo) SetDisplayCode(v string)`

SetDisplayCode sets DisplayCode field to given value.

### HasDisplayCode

`func (o *AssetInfo) HasDisplayCode() bool`

HasDisplayCode returns a boolean if a field has been set.

### GetDescription

`func (o *AssetInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssetInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconUrl

`func (o *AssetInfo) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *AssetInfo) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *AssetInfo) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *AssetInfo) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


