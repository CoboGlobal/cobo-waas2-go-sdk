# PaymentPayoutParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAccount** | **string** | The source account from which the payout will be made. - If the source account is a merchant account, provide the merchant&#39;s ID (e.g., \&quot;M1001\&quot;). - If the source account is the developer account, use the string &#x60;\&quot;developer\&quot;&#x60;.  | 
**TokenId** | **string** | The ID of the cryptocurrency you want to pay out. Specify this field when &#x60;payout_channel&#x60; is set to &#x60;Crypto&#x60;. Supported values: - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDC&#x60;, &#x60;SOL_USDC&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC&#x60;, &#x60;BSC_USDC&#x60; - USDT: &#x60;TRON_USDT&#x60;, &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;SOL_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**Amount** | **string** | The amount of the cryptocurrency to pay out.  | 
**CryptoAddressId** | Pointer to **string** | The ID of the crypto address used for crypto payouts. Specify this field when &#x60;payout_channel&#x60; is set to &#x60;Crypto&#x60;.  Call [List crypto addresses](https://www.cobo.com/payments/en/api-references/payment/list-crypto-addresses) to retrieve registered crypto addresses.  | [optional] 
**CryptoAddress** | Pointer to **string** | The actual blockchain address to which funds will be transferred. Specify this field when &#x60;payout_channel&#x60; is set to &#x60;Crypto&#x60;. &lt;Note&gt;   If you have enabled the *Use Destinations as Payout Whitelist* toggle in *Destinations*, you can only transfer to registered destinations. For more details, see [Destinations](https://www.cobo.com/payments/en/guides/destinations). &lt;/Note&gt;  | [optional] 

## Methods

### NewPaymentPayoutParam

`func NewPaymentPayoutParam(sourceAccount string, tokenId string, amount string, ) *PaymentPayoutParam`

NewPaymentPayoutParam instantiates a new PaymentPayoutParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPayoutParamWithDefaults

`func NewPaymentPayoutParamWithDefaults() *PaymentPayoutParam`

NewPaymentPayoutParamWithDefaults instantiates a new PaymentPayoutParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAccount

`func (o *PaymentPayoutParam) GetSourceAccount() string`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *PaymentPayoutParam) GetSourceAccountOk() (*string, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *PaymentPayoutParam) SetSourceAccount(v string)`

SetSourceAccount sets SourceAccount field to given value.


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


