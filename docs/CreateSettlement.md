# CreateSettlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | Pointer to **string** | Only used in Merchant settlement type. The merchant ID.  | [optional] 
**TokenId** | **string** | Only used in Crypto payout channel. The ID of the cryptocurrency you want to settle. Supported values:  - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDC&#x60;, &#x60;SOL_USDC&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC&#x60;, &#x60;BSC_USDC&#x60; - USDT: &#x60;TRON_USDT&#x60;, &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;SOL_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**Currency** | Pointer to **string** | Only used in OffRamp payout channel. The fiat currency for settling the cryptocurrency. Currently, only &#x60;USD&#x60; is supported.  | [optional] 
**Amount** | Pointer to **string** | The settlement amount. - In Crypto payout channel, this represents the settlement amount in the specified cryptocurrency. - In OffRamp payout channel, this represents the settlement amount in the specified fiat currency.  | [optional] 
**BankAccountId** | Pointer to **string** | ｜ Only used in OffRamp payout channel. The ID of the bank account where the settled funds will be deposited. | [optional] 
**CryptoAddressId** | Pointer to **string** | Only used in Crypto payout channel. The ID of the pre-approved crypto address used for Crypto settlements. - The value must refer to a valid address that has been pre-configured and approved for the given token.  | [optional] 
**OrderIds** | Pointer to **[]string** | A list of unique order IDs to be included in this settlement.  - This field is only applicable when &#x60;settlement_type&#x60; is set to &#x60;Merchant&#x60;. - If provided, the settlement will only apply to the specified orders. - The settlement &#x60;amount&#x60; must exactly match the total eligible amount from these orders. - This ensures consistency between the declared amount and the actual order-level data being settled.  | [optional] 

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


### GetCurrency

`func (o *CreateSettlement) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateSettlement) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateSettlement) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateSettlement) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

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

### GetBankAccountId

`func (o *CreateSettlement) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *CreateSettlement) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *CreateSettlement) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *CreateSettlement) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

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


