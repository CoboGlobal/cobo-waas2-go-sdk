# TransactionToken

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
**Amount** | Pointer to **float32** | Transaction value (Note that this is an absolute value. If you trade 1.5 BTC, then the value is 1.5)  | [optional] 

## Methods

### NewTransactionToken

`func NewTransactionToken(tokenId string, chainId string, ) *TransactionToken`

NewTransactionToken instantiates a new TransactionToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTokenWithDefaults

`func NewTransactionTokenWithDefaults() *TransactionToken`

NewTransactionTokenWithDefaults instantiates a new TransactionToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *TransactionToken) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionToken) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionToken) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *TransactionToken) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TransactionToken) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TransactionToken) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSymbol

`func (o *TransactionToken) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TransactionToken) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TransactionToken) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TransactionToken) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconUrl

`func (o *TransactionToken) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *TransactionToken) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *TransactionToken) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *TransactionToken) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetTokenAddress

`func (o *TransactionToken) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TransactionToken) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TransactionToken) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *TransactionToken) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetAssetId

`func (o *TransactionToken) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *TransactionToken) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *TransactionToken) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *TransactionToken) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionToken) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionToken) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionToken) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionToken) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


