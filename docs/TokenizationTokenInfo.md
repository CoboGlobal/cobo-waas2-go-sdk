# TokenizationTokenInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The unique token identifier. | 
**ChainId** | **string** | The chain ID of the tokenization contract. | 
**TokenAddress** | Pointer to **string** | The address of the token contract. | [optional] 
**TokenName** | Pointer to **string** | The name of the token. | [optional] 
**TokenSymbol** | **string** | The unique token symbol. | 
**TokenStandard** | [**TokenizationTokenStandard**](TokenizationTokenStandard.md) |  | 
**Decimals** | **int32** | The number of decimals of the token. | 
**TokenAccessActivated** | Pointer to **bool** | Whether the allowlist feature is activated for the token. | [optional] 
**Status** | [**TokenizationStatus**](TokenizationStatus.md) |  | 
**TotalSupply** | Pointer to **string** | The total supply of the token. | [optional] 
**Holdings** | Pointer to **string** | The amount of tokens held by the organization. | [optional] 
**Archived** | **bool** | Whether the token is archived. If the token is archived, no operations can be initiated on it. | 

## Methods

### NewTokenizationTokenInfo

`func NewTokenizationTokenInfo(tokenId string, chainId string, tokenSymbol string, tokenStandard TokenizationTokenStandard, decimals int32, status TokenizationStatus, archived bool, ) *TokenizationTokenInfo`

NewTokenizationTokenInfo instantiates a new TokenizationTokenInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationTokenInfoWithDefaults

`func NewTokenizationTokenInfoWithDefaults() *TokenizationTokenInfo`

NewTokenizationTokenInfoWithDefaults instantiates a new TokenizationTokenInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *TokenizationTokenInfo) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenizationTokenInfo) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenizationTokenInfo) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *TokenizationTokenInfo) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TokenizationTokenInfo) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TokenizationTokenInfo) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetTokenAddress

`func (o *TokenizationTokenInfo) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TokenizationTokenInfo) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TokenizationTokenInfo) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *TokenizationTokenInfo) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetTokenName

`func (o *TokenizationTokenInfo) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *TokenizationTokenInfo) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *TokenizationTokenInfo) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *TokenizationTokenInfo) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### GetTokenSymbol

`func (o *TokenizationTokenInfo) GetTokenSymbol() string`

GetTokenSymbol returns the TokenSymbol field if non-nil, zero value otherwise.

### GetTokenSymbolOk

`func (o *TokenizationTokenInfo) GetTokenSymbolOk() (*string, bool)`

GetTokenSymbolOk returns a tuple with the TokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSymbol

`func (o *TokenizationTokenInfo) SetTokenSymbol(v string)`

SetTokenSymbol sets TokenSymbol field to given value.


### GetTokenStandard

`func (o *TokenizationTokenInfo) GetTokenStandard() TokenizationTokenStandard`

GetTokenStandard returns the TokenStandard field if non-nil, zero value otherwise.

### GetTokenStandardOk

`func (o *TokenizationTokenInfo) GetTokenStandardOk() (*TokenizationTokenStandard, bool)`

GetTokenStandardOk returns a tuple with the TokenStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenStandard

`func (o *TokenizationTokenInfo) SetTokenStandard(v TokenizationTokenStandard)`

SetTokenStandard sets TokenStandard field to given value.


### GetDecimals

`func (o *TokenizationTokenInfo) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenizationTokenInfo) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenizationTokenInfo) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetTokenAccessActivated

`func (o *TokenizationTokenInfo) GetTokenAccessActivated() bool`

GetTokenAccessActivated returns the TokenAccessActivated field if non-nil, zero value otherwise.

### GetTokenAccessActivatedOk

`func (o *TokenizationTokenInfo) GetTokenAccessActivatedOk() (*bool, bool)`

GetTokenAccessActivatedOk returns a tuple with the TokenAccessActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAccessActivated

`func (o *TokenizationTokenInfo) SetTokenAccessActivated(v bool)`

SetTokenAccessActivated sets TokenAccessActivated field to given value.

### HasTokenAccessActivated

`func (o *TokenizationTokenInfo) HasTokenAccessActivated() bool`

HasTokenAccessActivated returns a boolean if a field has been set.

### GetStatus

`func (o *TokenizationTokenInfo) GetStatus() TokenizationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenizationTokenInfo) GetStatusOk() (*TokenizationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenizationTokenInfo) SetStatus(v TokenizationStatus)`

SetStatus sets Status field to given value.


### GetTotalSupply

`func (o *TokenizationTokenInfo) GetTotalSupply() string`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *TokenizationTokenInfo) GetTotalSupplyOk() (*string, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *TokenizationTokenInfo) SetTotalSupply(v string)`

SetTotalSupply sets TotalSupply field to given value.

### HasTotalSupply

`func (o *TokenizationTokenInfo) HasTotalSupply() bool`

HasTotalSupply returns a boolean if a field has been set.

### GetHoldings

`func (o *TokenizationTokenInfo) GetHoldings() string`

GetHoldings returns the Holdings field if non-nil, zero value otherwise.

### GetHoldingsOk

`func (o *TokenizationTokenInfo) GetHoldingsOk() (*string, bool)`

GetHoldingsOk returns a tuple with the Holdings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldings

`func (o *TokenizationTokenInfo) SetHoldings(v string)`

SetHoldings sets Holdings field to given value.

### HasHoldings

`func (o *TokenizationTokenInfo) HasHoldings() bool`

HasHoldings returns a boolean if a field has been set.

### GetArchived

`func (o *TokenizationTokenInfo) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TokenizationTokenInfo) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TokenizationTokenInfo) SetArchived(v bool)`

SetArchived sets Archived field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


