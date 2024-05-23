# GetTokens200ResponseDataInner

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
**CanDeposit** | Pointer to **bool** | Ture if the token can be deposited, False otherwise | [optional] [default to false]
**CanWithdraw** | Pointer to **bool** | Ture if the token can be withdrawn, False otherwise | [optional] [default to false]

## Methods

### NewGetTokens200ResponseDataInner

`func NewGetTokens200ResponseDataInner(tokenId string, chainId string, ) *GetTokens200ResponseDataInner`

NewGetTokens200ResponseDataInner instantiates a new GetTokens200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTokens200ResponseDataInnerWithDefaults

`func NewGetTokens200ResponseDataInnerWithDefaults() *GetTokens200ResponseDataInner`

NewGetTokens200ResponseDataInnerWithDefaults instantiates a new GetTokens200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *GetTokens200ResponseDataInner) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetTokens200ResponseDataInner) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetTokens200ResponseDataInner) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *GetTokens200ResponseDataInner) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *GetTokens200ResponseDataInner) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *GetTokens200ResponseDataInner) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetSymbol

`func (o *GetTokens200ResponseDataInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetTokens200ResponseDataInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetTokens200ResponseDataInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetTokens200ResponseDataInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *GetTokens200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetTokens200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetTokens200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetTokens200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconUrl

`func (o *GetTokens200ResponseDataInner) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *GetTokens200ResponseDataInner) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *GetTokens200ResponseDataInner) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *GetTokens200ResponseDataInner) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetTokenAddress

`func (o *GetTokens200ResponseDataInner) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *GetTokens200ResponseDataInner) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *GetTokens200ResponseDataInner) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *GetTokens200ResponseDataInner) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetAssetId

`func (o *GetTokens200ResponseDataInner) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *GetTokens200ResponseDataInner) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *GetTokens200ResponseDataInner) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *GetTokens200ResponseDataInner) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetCanDeposit

`func (o *GetTokens200ResponseDataInner) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *GetTokens200ResponseDataInner) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *GetTokens200ResponseDataInner) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *GetTokens200ResponseDataInner) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *GetTokens200ResponseDataInner) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *GetTokens200ResponseDataInner) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *GetTokens200ResponseDataInner) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *GetTokens200ResponseDataInner) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


