# TokenInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | ID of the token. Unique in all chains scope. | 
**ChainId** | **string** | The blockchain on which the token operates. | 
**Symbol** | Pointer to **string** | Symbol for the token. | [optional] 
**Description** | Pointer to **string** | The description of the token. | [optional] 
**IconUrl** | Pointer to **string** | URL of the icon image. | [optional] 
**TokenAddress** | Pointer to **string** | Address for token, if applicable. | [optional] 
**AssetId** | Pointer to **string** | ID of the asset. Used to group token balance when needed. | [optional] 

## Methods

### NewTokenInfo

`func NewTokenInfo(tokenId string, chainId string, ) *TokenInfo`

NewTokenInfo instantiates a new TokenInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenInfoWithDefaults

`func NewTokenInfoWithDefaults() *TokenInfo`

NewTokenInfoWithDefaults instantiates a new TokenInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *TokenInfo) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenInfo) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenInfo) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *TokenInfo) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TokenInfo) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TokenInfo) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSymbol

`func (o *TokenInfo) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenInfo) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenInfo) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TokenInfo) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *TokenInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TokenInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconUrl

`func (o *TokenInfo) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *TokenInfo) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *TokenInfo) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *TokenInfo) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetTokenAddress

`func (o *TokenInfo) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TokenInfo) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TokenInfo) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *TokenInfo) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetAssetId

`func (o *TokenInfo) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *TokenInfo) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *TokenInfo) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *TokenInfo) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


