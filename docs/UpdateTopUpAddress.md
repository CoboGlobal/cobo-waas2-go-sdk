# UpdateTopUpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | The merchant ID. | 
**TokenId** | **string** | The token ID, which is a unique identifier that specifies both the blockchain network and cryptocurrency token in the format &#x60;{CHAIN}_{TOKEN}&#x60;. Supported values include:   - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDC&#x60;, &#x60;SOL_USDC&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC&#x60;, &#x60;BSC_USDC&#x60;   - USDT: &#x60;TRON_USDT&#x60;, &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;SOL_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**CustomPayerId** | **string** | Unique customer identifier on the merchant side, used to allocate a dedicated top-up address  | 

## Methods

### NewUpdateTopUpAddress

`func NewUpdateTopUpAddress(merchantId string, tokenId string, customPayerId string, ) *UpdateTopUpAddress`

NewUpdateTopUpAddress instantiates a new UpdateTopUpAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTopUpAddressWithDefaults

`func NewUpdateTopUpAddressWithDefaults() *UpdateTopUpAddress`

NewUpdateTopUpAddressWithDefaults instantiates a new UpdateTopUpAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *UpdateTopUpAddress) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UpdateTopUpAddress) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UpdateTopUpAddress) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetTokenId

`func (o *UpdateTopUpAddress) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *UpdateTopUpAddress) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *UpdateTopUpAddress) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetCustomPayerId

`func (o *UpdateTopUpAddress) GetCustomPayerId() string`

GetCustomPayerId returns the CustomPayerId field if non-nil, zero value otherwise.

### GetCustomPayerIdOk

`func (o *UpdateTopUpAddress) GetCustomPayerIdOk() (*string, bool)`

GetCustomPayerIdOk returns a tuple with the CustomPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayerId

`func (o *UpdateTopUpAddress) SetCustomPayerId(v string)`

SetCustomPayerId sets CustomPayerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


