# SupportedToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The unique identifier of the token, in the format of &#x60;{chain_id}_{token_symbol}&#x60;. | 
**Name** | **string** | The full name of the token. | 
**Symbol** | **string** | The symbol of the token. | 
**Decimal** | **int32** | The number of decimal places for the token. This value is used to convert  between the token&#39;s smallest unit and its display value.  | 
**TokenAddress** | **NullableString** | The contract address of the token. This is &#x60;null&#x60; for native coins (e.g., ETH on Ethereum).  | 
**ChainId** | **string** | The ID of the chain on which the token exists. | 
**ChainSymbol** | **NullableString** | The symbol of the chain on which the token exists. | 
**ChainIconUrl** | Pointer to **NullableString** | The URL of the chain icon image. | [optional] 
**TokenIconUrl** | Pointer to **NullableString** | The URL of the token icon image. | [optional] 
**CanOffRamp** | Pointer to **bool** | Whether the token supports fiat off-ramp. - &#x60;true&#x60;: The token can be used for fiat off-ramp. - &#x60;false&#x60;: The token cannot be used for fiat off-ramp.  | [optional] 

## Methods

### NewSupportedToken

`func NewSupportedToken(tokenId string, name string, symbol string, decimal int32, tokenAddress NullableString, chainId string, chainSymbol NullableString, ) *SupportedToken`

NewSupportedToken instantiates a new SupportedToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedTokenWithDefaults

`func NewSupportedTokenWithDefaults() *SupportedToken`

NewSupportedTokenWithDefaults instantiates a new SupportedToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *SupportedToken) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *SupportedToken) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *SupportedToken) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetName

`func (o *SupportedToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupportedToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupportedToken) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *SupportedToken) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SupportedToken) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SupportedToken) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimal

`func (o *SupportedToken) GetDecimal() int32`

GetDecimal returns the Decimal field if non-nil, zero value otherwise.

### GetDecimalOk

`func (o *SupportedToken) GetDecimalOk() (*int32, bool)`

GetDecimalOk returns a tuple with the Decimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimal

`func (o *SupportedToken) SetDecimal(v int32)`

SetDecimal sets Decimal field to given value.


### GetTokenAddress

`func (o *SupportedToken) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *SupportedToken) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *SupportedToken) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### SetTokenAddressNil

`func (o *SupportedToken) SetTokenAddressNil(b bool)`

 SetTokenAddressNil sets the value for TokenAddress to be an explicit nil

### UnsetTokenAddress
`func (o *SupportedToken) UnsetTokenAddress()`

UnsetTokenAddress ensures that no value is present for TokenAddress, not even an explicit nil
### GetChainId

`func (o *SupportedToken) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *SupportedToken) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *SupportedToken) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetChainSymbol

`func (o *SupportedToken) GetChainSymbol() string`

GetChainSymbol returns the ChainSymbol field if non-nil, zero value otherwise.

### GetChainSymbolOk

`func (o *SupportedToken) GetChainSymbolOk() (*string, bool)`

GetChainSymbolOk returns a tuple with the ChainSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainSymbol

`func (o *SupportedToken) SetChainSymbol(v string)`

SetChainSymbol sets ChainSymbol field to given value.


### SetChainSymbolNil

`func (o *SupportedToken) SetChainSymbolNil(b bool)`

 SetChainSymbolNil sets the value for ChainSymbol to be an explicit nil

### UnsetChainSymbol
`func (o *SupportedToken) UnsetChainSymbol()`

UnsetChainSymbol ensures that no value is present for ChainSymbol, not even an explicit nil
### GetChainIconUrl

`func (o *SupportedToken) GetChainIconUrl() string`

GetChainIconUrl returns the ChainIconUrl field if non-nil, zero value otherwise.

### GetChainIconUrlOk

`func (o *SupportedToken) GetChainIconUrlOk() (*string, bool)`

GetChainIconUrlOk returns a tuple with the ChainIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainIconUrl

`func (o *SupportedToken) SetChainIconUrl(v string)`

SetChainIconUrl sets ChainIconUrl field to given value.

### HasChainIconUrl

`func (o *SupportedToken) HasChainIconUrl() bool`

HasChainIconUrl returns a boolean if a field has been set.

### SetChainIconUrlNil

`func (o *SupportedToken) SetChainIconUrlNil(b bool)`

 SetChainIconUrlNil sets the value for ChainIconUrl to be an explicit nil

### UnsetChainIconUrl
`func (o *SupportedToken) UnsetChainIconUrl()`

UnsetChainIconUrl ensures that no value is present for ChainIconUrl, not even an explicit nil
### GetTokenIconUrl

`func (o *SupportedToken) GetTokenIconUrl() string`

GetTokenIconUrl returns the TokenIconUrl field if non-nil, zero value otherwise.

### GetTokenIconUrlOk

`func (o *SupportedToken) GetTokenIconUrlOk() (*string, bool)`

GetTokenIconUrlOk returns a tuple with the TokenIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIconUrl

`func (o *SupportedToken) SetTokenIconUrl(v string)`

SetTokenIconUrl sets TokenIconUrl field to given value.

### HasTokenIconUrl

`func (o *SupportedToken) HasTokenIconUrl() bool`

HasTokenIconUrl returns a boolean if a field has been set.

### SetTokenIconUrlNil

`func (o *SupportedToken) SetTokenIconUrlNil(b bool)`

 SetTokenIconUrlNil sets the value for TokenIconUrl to be an explicit nil

### UnsetTokenIconUrl
`func (o *SupportedToken) UnsetTokenIconUrl()`

UnsetTokenIconUrl ensures that no value is present for TokenIconUrl, not even an explicit nil
### GetCanOffRamp

`func (o *SupportedToken) GetCanOffRamp() bool`

GetCanOffRamp returns the CanOffRamp field if non-nil, zero value otherwise.

### GetCanOffRampOk

`func (o *SupportedToken) GetCanOffRampOk() (*bool, bool)`

GetCanOffRampOk returns a tuple with the CanOffRamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanOffRamp

`func (o *SupportedToken) SetCanOffRamp(v bool)`

SetCanOffRamp sets CanOffRamp field to given value.

### HasCanOffRamp

`func (o *SupportedToken) HasCanOffRamp() bool`

HasCanOffRamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


