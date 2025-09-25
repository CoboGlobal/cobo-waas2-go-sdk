# TokenizationTokenDetailInfo

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
**Permissions** | Pointer to [**[]TokenizationAddressPermission**](TokenizationAddressPermission.md) | List of execution addresses and their permissions. | [optional] 

## Methods

### NewTokenizationTokenDetailInfo

`func NewTokenizationTokenDetailInfo(tokenId string, chainId string, tokenSymbol string, tokenStandard TokenizationTokenStandard, decimals int32, status TokenizationStatus, ) *TokenizationTokenDetailInfo`

NewTokenizationTokenDetailInfo instantiates a new TokenizationTokenDetailInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationTokenDetailInfoWithDefaults

`func NewTokenizationTokenDetailInfoWithDefaults() *TokenizationTokenDetailInfo`

NewTokenizationTokenDetailInfoWithDefaults instantiates a new TokenizationTokenDetailInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *TokenizationTokenDetailInfo) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenizationTokenDetailInfo) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenizationTokenDetailInfo) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *TokenizationTokenDetailInfo) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TokenizationTokenDetailInfo) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TokenizationTokenDetailInfo) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetTokenAddress

`func (o *TokenizationTokenDetailInfo) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TokenizationTokenDetailInfo) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TokenizationTokenDetailInfo) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *TokenizationTokenDetailInfo) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetTokenName

`func (o *TokenizationTokenDetailInfo) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *TokenizationTokenDetailInfo) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *TokenizationTokenDetailInfo) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *TokenizationTokenDetailInfo) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### GetTokenSymbol

`func (o *TokenizationTokenDetailInfo) GetTokenSymbol() string`

GetTokenSymbol returns the TokenSymbol field if non-nil, zero value otherwise.

### GetTokenSymbolOk

`func (o *TokenizationTokenDetailInfo) GetTokenSymbolOk() (*string, bool)`

GetTokenSymbolOk returns a tuple with the TokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSymbol

`func (o *TokenizationTokenDetailInfo) SetTokenSymbol(v string)`

SetTokenSymbol sets TokenSymbol field to given value.


### GetTokenStandard

`func (o *TokenizationTokenDetailInfo) GetTokenStandard() TokenizationTokenStandard`

GetTokenStandard returns the TokenStandard field if non-nil, zero value otherwise.

### GetTokenStandardOk

`func (o *TokenizationTokenDetailInfo) GetTokenStandardOk() (*TokenizationTokenStandard, bool)`

GetTokenStandardOk returns a tuple with the TokenStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenStandard

`func (o *TokenizationTokenDetailInfo) SetTokenStandard(v TokenizationTokenStandard)`

SetTokenStandard sets TokenStandard field to given value.


### GetDecimals

`func (o *TokenizationTokenDetailInfo) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenizationTokenDetailInfo) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenizationTokenDetailInfo) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetTokenAccessActivated

`func (o *TokenizationTokenDetailInfo) GetTokenAccessActivated() bool`

GetTokenAccessActivated returns the TokenAccessActivated field if non-nil, zero value otherwise.

### GetTokenAccessActivatedOk

`func (o *TokenizationTokenDetailInfo) GetTokenAccessActivatedOk() (*bool, bool)`

GetTokenAccessActivatedOk returns a tuple with the TokenAccessActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAccessActivated

`func (o *TokenizationTokenDetailInfo) SetTokenAccessActivated(v bool)`

SetTokenAccessActivated sets TokenAccessActivated field to given value.

### HasTokenAccessActivated

`func (o *TokenizationTokenDetailInfo) HasTokenAccessActivated() bool`

HasTokenAccessActivated returns a boolean if a field has been set.

### GetStatus

`func (o *TokenizationTokenDetailInfo) GetStatus() TokenizationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenizationTokenDetailInfo) GetStatusOk() (*TokenizationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenizationTokenDetailInfo) SetStatus(v TokenizationStatus)`

SetStatus sets Status field to given value.


### GetTotalSupply

`func (o *TokenizationTokenDetailInfo) GetTotalSupply() string`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *TokenizationTokenDetailInfo) GetTotalSupplyOk() (*string, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *TokenizationTokenDetailInfo) SetTotalSupply(v string)`

SetTotalSupply sets TotalSupply field to given value.

### HasTotalSupply

`func (o *TokenizationTokenDetailInfo) HasTotalSupply() bool`

HasTotalSupply returns a boolean if a field has been set.

### GetHoldings

`func (o *TokenizationTokenDetailInfo) GetHoldings() string`

GetHoldings returns the Holdings field if non-nil, zero value otherwise.

### GetHoldingsOk

`func (o *TokenizationTokenDetailInfo) GetHoldingsOk() (*string, bool)`

GetHoldingsOk returns a tuple with the Holdings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldings

`func (o *TokenizationTokenDetailInfo) SetHoldings(v string)`

SetHoldings sets Holdings field to given value.

### HasHoldings

`func (o *TokenizationTokenDetailInfo) HasHoldings() bool`

HasHoldings returns a boolean if a field has been set.

### GetPermissions

`func (o *TokenizationTokenDetailInfo) GetPermissions() []TokenizationAddressPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TokenizationTokenDetailInfo) GetPermissionsOk() (*[]TokenizationAddressPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TokenizationTokenDetailInfo) SetPermissions(v []TokenizationAddressPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *TokenizationTokenDetailInfo) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


