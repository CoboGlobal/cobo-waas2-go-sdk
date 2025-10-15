# SwapToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID. | 
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. | 
**AssetId** | **string** | The asset ID. | 
**TokenAddress** | Pointer to **string** | The on-chain contract address of the token. | [optional] 
**MinAmount** | Pointer to **string** | The minimum amount allowed for a swap. | [optional] 
**MaxAmount** | Pointer to **string** | The maximum amount allowed for a swap. | [optional] 

## Methods

### NewSwapToken

`func NewSwapToken(tokenId string, chainId string, assetId string, ) *SwapToken`

NewSwapToken instantiates a new SwapToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapTokenWithDefaults

`func NewSwapTokenWithDefaults() *SwapToken`

NewSwapTokenWithDefaults instantiates a new SwapToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *SwapToken) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *SwapToken) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *SwapToken) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *SwapToken) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *SwapToken) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *SwapToken) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetAssetId

`func (o *SwapToken) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *SwapToken) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *SwapToken) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetTokenAddress

`func (o *SwapToken) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *SwapToken) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *SwapToken) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *SwapToken) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetMinAmount

`func (o *SwapToken) GetMinAmount() string`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *SwapToken) GetMinAmountOk() (*string, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *SwapToken) SetMinAmount(v string)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *SwapToken) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### GetMaxAmount

`func (o *SwapToken) GetMaxAmount() string`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *SwapToken) GetMaxAmountOk() (*string, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *SwapToken) SetMaxAmount(v string)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *SwapToken) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


