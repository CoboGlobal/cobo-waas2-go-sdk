# OrderLinkBusinessInfoPayableAmountsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The ID of the cryptocurrency used for payment. Supported values: - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDC&#x60;, &#x60;SOL_USDC&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC&#x60;, &#x60;BSC_USDC&#x60; - USDT: &#x60;TRON_USDT&#x60;, &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;SOL_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**Amount** | **string** | The payable amount in the specified currency. | 

## Methods

### NewOrderLinkBusinessInfoPayableAmountsInner

`func NewOrderLinkBusinessInfoPayableAmountsInner(tokenId string, amount string, ) *OrderLinkBusinessInfoPayableAmountsInner`

NewOrderLinkBusinessInfoPayableAmountsInner instantiates a new OrderLinkBusinessInfoPayableAmountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderLinkBusinessInfoPayableAmountsInnerWithDefaults

`func NewOrderLinkBusinessInfoPayableAmountsInnerWithDefaults() *OrderLinkBusinessInfoPayableAmountsInner`

NewOrderLinkBusinessInfoPayableAmountsInnerWithDefaults instantiates a new OrderLinkBusinessInfoPayableAmountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *OrderLinkBusinessInfoPayableAmountsInner) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *OrderLinkBusinessInfoPayableAmountsInner) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *OrderLinkBusinessInfoPayableAmountsInner) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *OrderLinkBusinessInfoPayableAmountsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrderLinkBusinessInfoPayableAmountsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrderLinkBusinessInfoPayableAmountsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


