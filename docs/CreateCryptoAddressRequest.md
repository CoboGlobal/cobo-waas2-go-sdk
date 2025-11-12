# CreateCryptoAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID, which is a unique identifier that specifies both the blockchain network and cryptocurrency token in the format &#x60;{CHAIN}_{TOKEN}&#x60;. Supported values include:   - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDCOIN&#x60;, &#x60;SOL_USDC&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC2&#x60;, &#x60;BSC_USDC&#x60;   - USDT: &#x60;TRON_USDT&#x60;, &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;SOL_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**Address** | **string** | The blockchain address in its native format. This is the actual destination address where funds will be sent. The address must match the format required by the specified blockchain. For example:   - For &#x60;SOL_USDC&#x60;: Provide a Solana address   - For &#x60;ETH_USDT&#x60;: Provide an Ethereum address  | 
**Label** | Pointer to **string** | A label to help identify the address&#39;s purpose. Can contain up to 128 characters. | [optional] 

## Methods

### NewCreateCryptoAddressRequest

`func NewCreateCryptoAddressRequest(tokenId string, address string, ) *CreateCryptoAddressRequest`

NewCreateCryptoAddressRequest instantiates a new CreateCryptoAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCryptoAddressRequestWithDefaults

`func NewCreateCryptoAddressRequestWithDefaults() *CreateCryptoAddressRequest`

NewCreateCryptoAddressRequestWithDefaults instantiates a new CreateCryptoAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *CreateCryptoAddressRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CreateCryptoAddressRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CreateCryptoAddressRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAddress

`func (o *CreateCryptoAddressRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateCryptoAddressRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateCryptoAddressRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetLabel

`func (o *CreateCryptoAddressRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateCryptoAddressRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateCryptoAddressRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateCryptoAddressRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


