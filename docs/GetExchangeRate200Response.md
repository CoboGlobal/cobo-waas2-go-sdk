# GetExchangeRate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID, which is a unique identifier that specifies both the blockchain network and cryptocurrency token in the format &#x60;{CHAIN}_{TOKEN}&#x60;. | 
**Currency** | **string** | The fiat currency. | 
**Rate** | **string** | The current exchange rate between the specified currency pair. Expressed as the amount of fiat currency per one unit of cryptocurrency. For example, if the cryptocurrency is USDT and the fiat currency is USD, a rate of \&quot;0.99\&quot; means 1 USDT &#x3D; 0.99 USD. | 

## Methods

### NewGetExchangeRate200Response

`func NewGetExchangeRate200Response(tokenId string, currency string, rate string, ) *GetExchangeRate200Response`

NewGetExchangeRate200Response instantiates a new GetExchangeRate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExchangeRate200ResponseWithDefaults

`func NewGetExchangeRate200ResponseWithDefaults() *GetExchangeRate200Response`

NewGetExchangeRate200ResponseWithDefaults instantiates a new GetExchangeRate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *GetExchangeRate200Response) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetExchangeRate200Response) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetExchangeRate200Response) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetCurrency

`func (o *GetExchangeRate200Response) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetExchangeRate200Response) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetExchangeRate200Response) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetRate

`func (o *GetExchangeRate200Response) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GetExchangeRate200Response) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GetExchangeRate200Response) SetRate(v string)`

SetRate sets Rate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


