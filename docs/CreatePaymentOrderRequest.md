# CreatePaymentOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | The merchant ID. | 
**MerchantOrderCode** | Pointer to **string** | A unique reference code assigned by the merchant to identify this order in their system. | [optional] 
**PspOrderCode** | **string** | A unique reference code assigned by the developer to identify this order in their system. | 
**PricingCurrency** | Pointer to **string** | The pricing currency that denominates &#x60;pricing_amount&#x60; and &#x60;fee_amount&#x60;. If left empty, both values will be denominated in &#x60;payable_currency&#x60;.  Currently, For a complete list of supported currencies, see [Supported chains and tokens](https://www.cobo.com//payments/en/guides/supported-chains-and-tokens#pricing-currency).  | [optional] 
**PricingAmount** | Pointer to **string** | The base amount of the order, excluding the developer fee (specified in &#x60;fee_amount&#x60;). Values must be greater than &#x60;0&#x60; and contain two decimal places. | [optional] 
**FeeAmount** | **string** | The developer fee for the order. It is added to the base amount (&#x60;pricing_amount&#x60;) to determine the final charge. For example, if &#x60;pricing_amount&#x60; is \&quot;100.00\&quot; and &#x60;fee_amount&#x60; is \&quot;2.00\&quot;, the payer will be charged \&quot;102.00\&quot; in total, with \&quot;100.00\&quot; being settled to the merchant account and \&quot;2.00\&quot; settled to the developer account. Values must be greater than 0 and contain two decimal places. | 
**PayableCurrency** | **string** | The ID of the cryptocurrency used for payment. Supported values:   - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDC&#x60;, &#x60;SOL_USDC&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC&#x60;, &#x60;BSC_USDC&#x60;   - USDT: &#x60;TRON_USDT&#x60;, &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;SOL_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**PayableAmount** | Pointer to **string** | The total amount the payer needs to pay, denominated in the specified &#x60;payable_currency&#x60;. If this field is left blank, the system will automatically calculate the amount at order creation using the following formula: (&#x60;pricing_amount&#x60; + &#x60;fee_amount&#x60;) / current exchange rate.  Values must be greater than 0 and contain two decimal places.  | [optional] 
**ExpiredIn** | Pointer to **int32** | The number of seconds until the pay-in order expires, counted from when the request is sent. For example, if set to &#x60;1800&#x60;, the order will expire in 30 minutes. Must be greater than zero and cannot exceed 3 hours (10800 seconds). After expiration:  - The order status becomes final and cannot be changed - The &#x60;received_token_amount&#x60; field will no longer be updated - Funds received after expiration will be categorized as late payments and can only be settled from the developer balance. - A late payment will trigger a &#x60;transactionLate&#x60; webhook event.  | [optional] [default to 1800]
**AmountTolerance** | Pointer to **string** | The allowed amount deviation, with precision up to 1 decimal place.  For example, if &#x60;payable_amount&#x60; is &#x60;100.00&#x60; and &#x60;amount_tolerance&#x60; is &#x60;0.50&#x60;: - Payer pays 99.55 → Success (difference of 0.45 ≤ 0.5) - Payer pays 99.40 → Underpaid (difference of 0.60 &gt; 0.5)  | [optional] 
**Currency** | Pointer to **string** | This field has been deprecated. Please use &#x60;pricing_currency&#x60; instead. | [optional] [default to ""]
**OrderAmount** | Pointer to **string** | This field has been deprecated. Please use &#x60;pricing_amount&#x60; instead. | [optional] 
**TokenId** | Pointer to **string** | This field has been deprecated. Please use &#x60;payable_currency&#x60; instead. | [optional] 
**UseDedicatedAddress** | Pointer to **bool** | This field has been deprecated. | [optional] 
**CustomExchangeRate** | Pointer to **string** | This field has been deprecated. | [optional] 

## Methods

### NewCreatePaymentOrderRequest

`func NewCreatePaymentOrderRequest(merchantId string, pspOrderCode string, feeAmount string, payableCurrency string, ) *CreatePaymentOrderRequest`

NewCreatePaymentOrderRequest instantiates a new CreatePaymentOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentOrderRequestWithDefaults

`func NewCreatePaymentOrderRequestWithDefaults() *CreatePaymentOrderRequest`

NewCreatePaymentOrderRequestWithDefaults instantiates a new CreatePaymentOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *CreatePaymentOrderRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CreatePaymentOrderRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CreatePaymentOrderRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetMerchantOrderCode

`func (o *CreatePaymentOrderRequest) GetMerchantOrderCode() string`

GetMerchantOrderCode returns the MerchantOrderCode field if non-nil, zero value otherwise.

### GetMerchantOrderCodeOk

`func (o *CreatePaymentOrderRequest) GetMerchantOrderCodeOk() (*string, bool)`

GetMerchantOrderCodeOk returns a tuple with the MerchantOrderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderCode

`func (o *CreatePaymentOrderRequest) SetMerchantOrderCode(v string)`

SetMerchantOrderCode sets MerchantOrderCode field to given value.

### HasMerchantOrderCode

`func (o *CreatePaymentOrderRequest) HasMerchantOrderCode() bool`

HasMerchantOrderCode returns a boolean if a field has been set.

### GetPspOrderCode

`func (o *CreatePaymentOrderRequest) GetPspOrderCode() string`

GetPspOrderCode returns the PspOrderCode field if non-nil, zero value otherwise.

### GetPspOrderCodeOk

`func (o *CreatePaymentOrderRequest) GetPspOrderCodeOk() (*string, bool)`

GetPspOrderCodeOk returns a tuple with the PspOrderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspOrderCode

`func (o *CreatePaymentOrderRequest) SetPspOrderCode(v string)`

SetPspOrderCode sets PspOrderCode field to given value.


### GetPricingCurrency

`func (o *CreatePaymentOrderRequest) GetPricingCurrency() string`

GetPricingCurrency returns the PricingCurrency field if non-nil, zero value otherwise.

### GetPricingCurrencyOk

`func (o *CreatePaymentOrderRequest) GetPricingCurrencyOk() (*string, bool)`

GetPricingCurrencyOk returns a tuple with the PricingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingCurrency

`func (o *CreatePaymentOrderRequest) SetPricingCurrency(v string)`

SetPricingCurrency sets PricingCurrency field to given value.

### HasPricingCurrency

`func (o *CreatePaymentOrderRequest) HasPricingCurrency() bool`

HasPricingCurrency returns a boolean if a field has been set.

### GetPricingAmount

`func (o *CreatePaymentOrderRequest) GetPricingAmount() string`

GetPricingAmount returns the PricingAmount field if non-nil, zero value otherwise.

### GetPricingAmountOk

`func (o *CreatePaymentOrderRequest) GetPricingAmountOk() (*string, bool)`

GetPricingAmountOk returns a tuple with the PricingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingAmount

`func (o *CreatePaymentOrderRequest) SetPricingAmount(v string)`

SetPricingAmount sets PricingAmount field to given value.

### HasPricingAmount

`func (o *CreatePaymentOrderRequest) HasPricingAmount() bool`

HasPricingAmount returns a boolean if a field has been set.

### GetFeeAmount

`func (o *CreatePaymentOrderRequest) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *CreatePaymentOrderRequest) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *CreatePaymentOrderRequest) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetPayableCurrency

`func (o *CreatePaymentOrderRequest) GetPayableCurrency() string`

GetPayableCurrency returns the PayableCurrency field if non-nil, zero value otherwise.

### GetPayableCurrencyOk

`func (o *CreatePaymentOrderRequest) GetPayableCurrencyOk() (*string, bool)`

GetPayableCurrencyOk returns a tuple with the PayableCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableCurrency

`func (o *CreatePaymentOrderRequest) SetPayableCurrency(v string)`

SetPayableCurrency sets PayableCurrency field to given value.


### GetPayableAmount

`func (o *CreatePaymentOrderRequest) GetPayableAmount() string`

GetPayableAmount returns the PayableAmount field if non-nil, zero value otherwise.

### GetPayableAmountOk

`func (o *CreatePaymentOrderRequest) GetPayableAmountOk() (*string, bool)`

GetPayableAmountOk returns a tuple with the PayableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableAmount

`func (o *CreatePaymentOrderRequest) SetPayableAmount(v string)`

SetPayableAmount sets PayableAmount field to given value.

### HasPayableAmount

`func (o *CreatePaymentOrderRequest) HasPayableAmount() bool`

HasPayableAmount returns a boolean if a field has been set.

### GetExpiredIn

`func (o *CreatePaymentOrderRequest) GetExpiredIn() int32`

GetExpiredIn returns the ExpiredIn field if non-nil, zero value otherwise.

### GetExpiredInOk

`func (o *CreatePaymentOrderRequest) GetExpiredInOk() (*int32, bool)`

GetExpiredInOk returns a tuple with the ExpiredIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredIn

`func (o *CreatePaymentOrderRequest) SetExpiredIn(v int32)`

SetExpiredIn sets ExpiredIn field to given value.

### HasExpiredIn

`func (o *CreatePaymentOrderRequest) HasExpiredIn() bool`

HasExpiredIn returns a boolean if a field has been set.

### GetAmountTolerance

`func (o *CreatePaymentOrderRequest) GetAmountTolerance() string`

GetAmountTolerance returns the AmountTolerance field if non-nil, zero value otherwise.

### GetAmountToleranceOk

`func (o *CreatePaymentOrderRequest) GetAmountToleranceOk() (*string, bool)`

GetAmountToleranceOk returns a tuple with the AmountTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountTolerance

`func (o *CreatePaymentOrderRequest) SetAmountTolerance(v string)`

SetAmountTolerance sets AmountTolerance field to given value.

### HasAmountTolerance

`func (o *CreatePaymentOrderRequest) HasAmountTolerance() bool`

HasAmountTolerance returns a boolean if a field has been set.

### GetCurrency

`func (o *CreatePaymentOrderRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreatePaymentOrderRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreatePaymentOrderRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreatePaymentOrderRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetOrderAmount

`func (o *CreatePaymentOrderRequest) GetOrderAmount() string`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *CreatePaymentOrderRequest) GetOrderAmountOk() (*string, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *CreatePaymentOrderRequest) SetOrderAmount(v string)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *CreatePaymentOrderRequest) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetTokenId

`func (o *CreatePaymentOrderRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CreatePaymentOrderRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CreatePaymentOrderRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *CreatePaymentOrderRequest) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetUseDedicatedAddress

`func (o *CreatePaymentOrderRequest) GetUseDedicatedAddress() bool`

GetUseDedicatedAddress returns the UseDedicatedAddress field if non-nil, zero value otherwise.

### GetUseDedicatedAddressOk

`func (o *CreatePaymentOrderRequest) GetUseDedicatedAddressOk() (*bool, bool)`

GetUseDedicatedAddressOk returns a tuple with the UseDedicatedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDedicatedAddress

`func (o *CreatePaymentOrderRequest) SetUseDedicatedAddress(v bool)`

SetUseDedicatedAddress sets UseDedicatedAddress field to given value.

### HasUseDedicatedAddress

`func (o *CreatePaymentOrderRequest) HasUseDedicatedAddress() bool`

HasUseDedicatedAddress returns a boolean if a field has been set.

### GetCustomExchangeRate

`func (o *CreatePaymentOrderRequest) GetCustomExchangeRate() string`

GetCustomExchangeRate returns the CustomExchangeRate field if non-nil, zero value otherwise.

### GetCustomExchangeRateOk

`func (o *CreatePaymentOrderRequest) GetCustomExchangeRateOk() (*string, bool)`

GetCustomExchangeRateOk returns a tuple with the CustomExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomExchangeRate

`func (o *CreatePaymentOrderRequest) SetCustomExchangeRate(v string)`

SetCustomExchangeRate sets CustomExchangeRate field to given value.

### HasCustomExchangeRate

`func (o *CreatePaymentOrderRequest) HasCustomExchangeRate() bool`

HasCustomExchangeRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


