# CreateSettlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | Pointer to **string** | The merchant ID. Specify this field when &#x60;settlement_type&#x60; is set to &#x60;Merchant&#x60;. | [optional] 
**TokenId** | **string** | The ID of the cryptocurrency you want to settle. Specify this field when &#x60;payout_channel&#x60; is set to &#x60;Crypto&#x60;. Supported values: - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDCOIN&#x60;, &#x60;SOL_USDC&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC2&#x60;, &#x60;BSC_USDC&#x60; - USDT: &#x60;TRON_USDT&#x60;, &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;SOL_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**Amount** | Pointer to **string** | The amount of cryptocurrency to be settled. When settling merchant balance from orders (&#x60;acquiring_type&#x60; is &#x60;Order&#x60; and &#x60;settlement_type&#x60; is &#x60;Merchant&#x60;), do not specify this field as the amount will be automatically calculated based on the order amounts.  | [optional] 
**CryptoAddressId** | Pointer to **string** | The ID of the crypto address used for crypto payouts. Specify this field when &#x60;payout_channel&#x60; is set to &#x60;Crypto&#x60;.  Call [List crypto addresses](https://www.cobo.com/payments/en/api-references/payment/list-crypto-addresses) to retrieve registered crypto addresses.  | [optional] 
**CryptoAddress** | Pointer to **string** | Only used in Crypto payout channel. The actual blockchain address to which funds will be transferred. If enable destination whitelist, this address must be associated with a destination.  | [optional] 
**OrderIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateSettlement

`func NewCreateSettlement(tokenId string, ) *CreateSettlement`

NewCreateSettlement instantiates a new CreateSettlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSettlementWithDefaults

`func NewCreateSettlementWithDefaults() *CreateSettlement`

NewCreateSettlementWithDefaults instantiates a new CreateSettlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *CreateSettlement) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CreateSettlement) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CreateSettlement) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *CreateSettlement) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetTokenId

`func (o *CreateSettlement) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CreateSettlement) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CreateSettlement) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *CreateSettlement) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateSettlement) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateSettlement) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreateSettlement) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCryptoAddressId

`func (o *CreateSettlement) GetCryptoAddressId() string`

GetCryptoAddressId returns the CryptoAddressId field if non-nil, zero value otherwise.

### GetCryptoAddressIdOk

`func (o *CreateSettlement) GetCryptoAddressIdOk() (*string, bool)`

GetCryptoAddressIdOk returns a tuple with the CryptoAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAddressId

`func (o *CreateSettlement) SetCryptoAddressId(v string)`

SetCryptoAddressId sets CryptoAddressId field to given value.

### HasCryptoAddressId

`func (o *CreateSettlement) HasCryptoAddressId() bool`

HasCryptoAddressId returns a boolean if a field has been set.

### GetCryptoAddress

`func (o *CreateSettlement) GetCryptoAddress() string`

GetCryptoAddress returns the CryptoAddress field if non-nil, zero value otherwise.

### GetCryptoAddressOk

`func (o *CreateSettlement) GetCryptoAddressOk() (*string, bool)`

GetCryptoAddressOk returns a tuple with the CryptoAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAddress

`func (o *CreateSettlement) SetCryptoAddress(v string)`

SetCryptoAddress sets CryptoAddress field to given value.

### HasCryptoAddress

`func (o *CreateSettlement) HasCryptoAddress() bool`

HasCryptoAddress returns a boolean if a field has been set.

### GetOrderIds

`func (o *CreateSettlement) GetOrderIds() []string`

GetOrderIds returns the OrderIds field if non-nil, zero value otherwise.

### GetOrderIdsOk

`func (o *CreateSettlement) GetOrderIdsOk() (*[]string, bool)`

GetOrderIdsOk returns a tuple with the OrderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIds

`func (o *CreateSettlement) SetOrderIds(v []string)`

SetOrderIds sets OrderIds field to given value.

### HasOrderIds

`func (o *CreateSettlement) HasOrderIds() bool`

HasOrderIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


