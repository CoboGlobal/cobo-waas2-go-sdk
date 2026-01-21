# OrderLinkBusinessInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | The merchant ID. | 
**MerchantOrderCode** | Pointer to **string** | A unique reference code assigned by the merchant to identify this order in their system. The code should have a maximum length of 128 characters. | [optional] 
**PspOrderCode** | **string** | A unique reference code assigned by you as a developer to identify this order in your system. This code must be unique across all orders in your system. The code should have a maximum length of 128 characters.  | 
**PricingCurrency** | **string** | The pricing currency that denominates &#x60;pricing_amount&#x60; and &#x60;fee_amount&#x60;. Currently, only &#x60;USD&#x60;/&#x60;USDT&#x60;/&#x60;USDC&#x60; are supported. This field is required.  | 
**PricingAmount** | **string** | The base amount of the order, excluding the developer fee (specified in &#x60;fee_amount&#x60;). Values must be greater than &#x60;0&#x60; and contain two decimal places.  | 
**FeeAmount** | **string** | The developer fee for the order. It is added to the base amount (&#x60;pricing_amount&#x60;) to determine the final charge. For example, if &#x60;pricing_amount&#x60; is \&quot;100.00\&quot; and &#x60;fee_amount&#x60; is \&quot;2.00\&quot;, the payer will be charged \&quot;102.00\&quot; in total, with \&quot;100.00\&quot; being settled to the merchant account and \&quot;2.00\&quot; settled to the developer account. Values must be greater than 0 and contain two decimal places.  | 
**PayableCurrencies** | **[]string** | The IDs of the cryptocurrencies used for payment. Supported values:  - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDC&#x60;, &#x60;SOL_USDC&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC&#x60;, &#x60;BSC_USDC&#x60;  - USDT: &#x60;TRON_USDT&#x60;, &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;SOL_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**PayableAmounts** | Pointer to [**[]OrderLinkBusinessInfoPayableAmountsInner**](OrderLinkBusinessInfoPayableAmountsInner.md) | The total amounts the payer needs to pay for each currency in &#x60;payable_currencies&#x60;. If this field is left blank, the system will automatically calculate the amounts at order creation using the following formula: (&#x60;pricing_amount&#x60; + &#x60;fee_amount&#x60;) / current exchange rate.  Values must be greater than 0 and contain two decimal places.  | [optional] 
**ExpiredIn** | Pointer to **int32** | The number of seconds until the pay-in order expires, counted from when the request is sent. For example, if set to &#x60;1800&#x60;, the order will expire in 30 minutes. Must be greater than zero and cannot exceed 3 hours (10800 seconds). After expiration:  - The order status becomes final and cannot be changed - The &#x60;received_token_amount&#x60; field will no longer be updated - Funds received after expiration will be categorized as late payments and can only be settled from the developer balance. - A late payment will trigger a &#x60;transactionLate&#x60; webhook event.  | [optional] [default to 1800]
**AmountTolerance** | Pointer to **string** | The allowed amount deviation, with precision up to 1 decimal place.  For example, if &#x60;payable_amount&#x60; is &#x60;100.00&#x60; and &#x60;amount_tolerance&#x60; is &#x60;0.50&#x60;: - Payer pays 99.55 → Success (difference of 0.45 ≤ 0.5) - Payer pays 99.40 → Underpaid (difference of 0.60 &gt; 0.5)  | [optional] 
**Currency** | Pointer to **string** | This field has been deprecated. Please use &#x60;pricing_currency&#x60; instead. | [optional] 
**OrderAmount** | Pointer to **string** | This field has been deprecated. Please use &#x60;pricing_amount&#x60; instead. | [optional] 
**TokenIds** | Pointer to **[]string** | This field has been deprecated. Please use &#x60;payable_currencies&#x60; instead.  | [optional] 
**CustomExchangeRates** | Pointer to [**[]OrderLinkBusinessInfoCustomExchangeRatesInner**](OrderLinkBusinessInfoCustomExchangeRatesInner.md) | This field has been deprecated.  | [optional] 
**UseDedicatedAddress** | Pointer to **bool** | This field has been deprecated. | [optional] 

## Methods

### NewOrderLinkBusinessInfo

`func NewOrderLinkBusinessInfo(merchantId string, pspOrderCode string, pricingCurrency string, pricingAmount string, feeAmount string, payableCurrencies []string, ) *OrderLinkBusinessInfo`

NewOrderLinkBusinessInfo instantiates a new OrderLinkBusinessInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderLinkBusinessInfoWithDefaults

`func NewOrderLinkBusinessInfoWithDefaults() *OrderLinkBusinessInfo`

NewOrderLinkBusinessInfoWithDefaults instantiates a new OrderLinkBusinessInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *OrderLinkBusinessInfo) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *OrderLinkBusinessInfo) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *OrderLinkBusinessInfo) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetMerchantOrderCode

`func (o *OrderLinkBusinessInfo) GetMerchantOrderCode() string`

GetMerchantOrderCode returns the MerchantOrderCode field if non-nil, zero value otherwise.

### GetMerchantOrderCodeOk

`func (o *OrderLinkBusinessInfo) GetMerchantOrderCodeOk() (*string, bool)`

GetMerchantOrderCodeOk returns a tuple with the MerchantOrderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderCode

`func (o *OrderLinkBusinessInfo) SetMerchantOrderCode(v string)`

SetMerchantOrderCode sets MerchantOrderCode field to given value.

### HasMerchantOrderCode

`func (o *OrderLinkBusinessInfo) HasMerchantOrderCode() bool`

HasMerchantOrderCode returns a boolean if a field has been set.

### GetPspOrderCode

`func (o *OrderLinkBusinessInfo) GetPspOrderCode() string`

GetPspOrderCode returns the PspOrderCode field if non-nil, zero value otherwise.

### GetPspOrderCodeOk

`func (o *OrderLinkBusinessInfo) GetPspOrderCodeOk() (*string, bool)`

GetPspOrderCodeOk returns a tuple with the PspOrderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspOrderCode

`func (o *OrderLinkBusinessInfo) SetPspOrderCode(v string)`

SetPspOrderCode sets PspOrderCode field to given value.


### GetPricingCurrency

`func (o *OrderLinkBusinessInfo) GetPricingCurrency() string`

GetPricingCurrency returns the PricingCurrency field if non-nil, zero value otherwise.

### GetPricingCurrencyOk

`func (o *OrderLinkBusinessInfo) GetPricingCurrencyOk() (*string, bool)`

GetPricingCurrencyOk returns a tuple with the PricingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingCurrency

`func (o *OrderLinkBusinessInfo) SetPricingCurrency(v string)`

SetPricingCurrency sets PricingCurrency field to given value.


### GetPricingAmount

`func (o *OrderLinkBusinessInfo) GetPricingAmount() string`

GetPricingAmount returns the PricingAmount field if non-nil, zero value otherwise.

### GetPricingAmountOk

`func (o *OrderLinkBusinessInfo) GetPricingAmountOk() (*string, bool)`

GetPricingAmountOk returns a tuple with the PricingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingAmount

`func (o *OrderLinkBusinessInfo) SetPricingAmount(v string)`

SetPricingAmount sets PricingAmount field to given value.


### GetFeeAmount

`func (o *OrderLinkBusinessInfo) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *OrderLinkBusinessInfo) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *OrderLinkBusinessInfo) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetPayableCurrencies

`func (o *OrderLinkBusinessInfo) GetPayableCurrencies() []string`

GetPayableCurrencies returns the PayableCurrencies field if non-nil, zero value otherwise.

### GetPayableCurrenciesOk

`func (o *OrderLinkBusinessInfo) GetPayableCurrenciesOk() (*[]string, bool)`

GetPayableCurrenciesOk returns a tuple with the PayableCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableCurrencies

`func (o *OrderLinkBusinessInfo) SetPayableCurrencies(v []string)`

SetPayableCurrencies sets PayableCurrencies field to given value.


### GetPayableAmounts

`func (o *OrderLinkBusinessInfo) GetPayableAmounts() []OrderLinkBusinessInfoPayableAmountsInner`

GetPayableAmounts returns the PayableAmounts field if non-nil, zero value otherwise.

### GetPayableAmountsOk

`func (o *OrderLinkBusinessInfo) GetPayableAmountsOk() (*[]OrderLinkBusinessInfoPayableAmountsInner, bool)`

GetPayableAmountsOk returns a tuple with the PayableAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableAmounts

`func (o *OrderLinkBusinessInfo) SetPayableAmounts(v []OrderLinkBusinessInfoPayableAmountsInner)`

SetPayableAmounts sets PayableAmounts field to given value.

### HasPayableAmounts

`func (o *OrderLinkBusinessInfo) HasPayableAmounts() bool`

HasPayableAmounts returns a boolean if a field has been set.

### GetExpiredIn

`func (o *OrderLinkBusinessInfo) GetExpiredIn() int32`

GetExpiredIn returns the ExpiredIn field if non-nil, zero value otherwise.

### GetExpiredInOk

`func (o *OrderLinkBusinessInfo) GetExpiredInOk() (*int32, bool)`

GetExpiredInOk returns a tuple with the ExpiredIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredIn

`func (o *OrderLinkBusinessInfo) SetExpiredIn(v int32)`

SetExpiredIn sets ExpiredIn field to given value.

### HasExpiredIn

`func (o *OrderLinkBusinessInfo) HasExpiredIn() bool`

HasExpiredIn returns a boolean if a field has been set.

### GetAmountTolerance

`func (o *OrderLinkBusinessInfo) GetAmountTolerance() string`

GetAmountTolerance returns the AmountTolerance field if non-nil, zero value otherwise.

### GetAmountToleranceOk

`func (o *OrderLinkBusinessInfo) GetAmountToleranceOk() (*string, bool)`

GetAmountToleranceOk returns a tuple with the AmountTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountTolerance

`func (o *OrderLinkBusinessInfo) SetAmountTolerance(v string)`

SetAmountTolerance sets AmountTolerance field to given value.

### HasAmountTolerance

`func (o *OrderLinkBusinessInfo) HasAmountTolerance() bool`

HasAmountTolerance returns a boolean if a field has been set.

### GetCurrency

`func (o *OrderLinkBusinessInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderLinkBusinessInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderLinkBusinessInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrderLinkBusinessInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetOrderAmount

`func (o *OrderLinkBusinessInfo) GetOrderAmount() string`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *OrderLinkBusinessInfo) GetOrderAmountOk() (*string, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *OrderLinkBusinessInfo) SetOrderAmount(v string)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *OrderLinkBusinessInfo) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetTokenIds

`func (o *OrderLinkBusinessInfo) GetTokenIds() []string`

GetTokenIds returns the TokenIds field if non-nil, zero value otherwise.

### GetTokenIdsOk

`func (o *OrderLinkBusinessInfo) GetTokenIdsOk() (*[]string, bool)`

GetTokenIdsOk returns a tuple with the TokenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIds

`func (o *OrderLinkBusinessInfo) SetTokenIds(v []string)`

SetTokenIds sets TokenIds field to given value.

### HasTokenIds

`func (o *OrderLinkBusinessInfo) HasTokenIds() bool`

HasTokenIds returns a boolean if a field has been set.

### GetCustomExchangeRates

`func (o *OrderLinkBusinessInfo) GetCustomExchangeRates() []OrderLinkBusinessInfoCustomExchangeRatesInner`

GetCustomExchangeRates returns the CustomExchangeRates field if non-nil, zero value otherwise.

### GetCustomExchangeRatesOk

`func (o *OrderLinkBusinessInfo) GetCustomExchangeRatesOk() (*[]OrderLinkBusinessInfoCustomExchangeRatesInner, bool)`

GetCustomExchangeRatesOk returns a tuple with the CustomExchangeRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomExchangeRates

`func (o *OrderLinkBusinessInfo) SetCustomExchangeRates(v []OrderLinkBusinessInfoCustomExchangeRatesInner)`

SetCustomExchangeRates sets CustomExchangeRates field to given value.

### HasCustomExchangeRates

`func (o *OrderLinkBusinessInfo) HasCustomExchangeRates() bool`

HasCustomExchangeRates returns a boolean if a field has been set.

### GetUseDedicatedAddress

`func (o *OrderLinkBusinessInfo) GetUseDedicatedAddress() bool`

GetUseDedicatedAddress returns the UseDedicatedAddress field if non-nil, zero value otherwise.

### GetUseDedicatedAddressOk

`func (o *OrderLinkBusinessInfo) GetUseDedicatedAddressOk() (*bool, bool)`

GetUseDedicatedAddressOk returns a tuple with the UseDedicatedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDedicatedAddress

`func (o *OrderLinkBusinessInfo) SetUseDedicatedAddress(v bool)`

SetUseDedicatedAddress sets UseDedicatedAddress field to given value.

### HasUseDedicatedAddress

`func (o *OrderLinkBusinessInfo) HasUseDedicatedAddress() bool`

HasUseDedicatedAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


