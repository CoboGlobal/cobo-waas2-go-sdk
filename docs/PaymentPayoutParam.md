# PaymentPayoutParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | Pointer to **string** | Only used in Merchant source type. The merchant ID.  | [optional] 
**TokenId** | **string** | Only used in Crypto payout channel. The ID of the cryptocurrency you want to settle. Supported values:  - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDC&#x60;, &#x60;SOL_USDC&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC&#x60;, &#x60;BSC_USDC&#x60; - USDT: &#x60;TRON_USDT&#x60;, &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;SOL_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**Amount** | **string** | The payout cryptocurrency amount.  | 
**CryptoAddressId** | Pointer to **string** | Only used in Crypto payout channel. The ID of the pre-approved crypto address used for Crypto settlements. - The value must refer to a valid address that has been pre-configured and approved for the given token.  | [optional] 
**CryptoAddress** | Pointer to **string** | Only used in Crypto payout channel. The actual blockchain address to which funds will be transferred. If enable destination whitelist, this address must be associated with a destination.  | [optional] 

## Methods

### NewPaymentPayoutParam

`func NewPaymentPayoutParam(tokenId string, amount string, ) *PaymentPayoutParam`

NewPaymentPayoutParam instantiates a new PaymentPayoutParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPayoutParamWithDefaults

`func NewPaymentPayoutParamWithDefaults() *PaymentPayoutParam`

NewPaymentPayoutParamWithDefaults instantiates a new PaymentPayoutParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *PaymentPayoutParam) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PaymentPayoutParam) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PaymentPayoutParam) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *PaymentPayoutParam) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetTokenId

`func (o *PaymentPayoutParam) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentPayoutParam) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentPayoutParam) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *PaymentPayoutParam) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentPayoutParam) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentPayoutParam) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCryptoAddressId

`func (o *PaymentPayoutParam) GetCryptoAddressId() string`

GetCryptoAddressId returns the CryptoAddressId field if non-nil, zero value otherwise.

### GetCryptoAddressIdOk

`func (o *PaymentPayoutParam) GetCryptoAddressIdOk() (*string, bool)`

GetCryptoAddressIdOk returns a tuple with the CryptoAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAddressId

`func (o *PaymentPayoutParam) SetCryptoAddressId(v string)`

SetCryptoAddressId sets CryptoAddressId field to given value.

### HasCryptoAddressId

`func (o *PaymentPayoutParam) HasCryptoAddressId() bool`

HasCryptoAddressId returns a boolean if a field has been set.

### GetCryptoAddress

`func (o *PaymentPayoutParam) GetCryptoAddress() string`

GetCryptoAddress returns the CryptoAddress field if non-nil, zero value otherwise.

### GetCryptoAddressOk

`func (o *PaymentPayoutParam) GetCryptoAddressOk() (*string, bool)`

GetCryptoAddressOk returns a tuple with the CryptoAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAddress

`func (o *PaymentPayoutParam) SetCryptoAddress(v string)`

SetCryptoAddress sets CryptoAddress field to given value.

### HasCryptoAddress

`func (o *PaymentPayoutParam) HasCryptoAddress() bool`

HasCryptoAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


