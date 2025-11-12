# CreatePaymentOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | The merchant ID. This ID is assigned by Cobo when you create the merchant. | 
**TokenId** | **string** | The ID of the cryptocurrency used for payment. Supported values:    - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDCOIN&#x60;, &#x60;SOL_USDC&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC2&#x60;, &#x60;BSC_USDC&#x60;   - USDT: &#x60;TRON_USDT&#x60;, &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;SOL_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**Currency** | Pointer to **string** | The fiat currency for the base order amount and the developer fee. If left empty, both &#x60;order_amount&#x60; and &#x60;fee_amount&#x60; will be denominated in the cryptocurrency specified by &#x60;token_id&#x60;  Currently, only &#x60;USD&#x60; is supported.  | [optional] [default to ""]
**OrderAmount** | **string** |  The base amount of the order, excluding the developer fee (specified in &#x60;fee_amount&#x60;). The denomination of the amount depends on if &#x60;currency&#x60; is specified: - If &#x60;currency&#x60; is specified, the amount is in the currency specified by &#x60;currency&#x60;, e.g. \&quot;USD\&quot;. - If &#x60;currency&#x60; is not specified, the amount is in the cryptocurrency specified by &#x60;token_id&#x60;, e.g. \&quot;ETH_USDT\&quot;.   Values must be greater than &#x60;0&#x60; and contain two decimal places.   | 
**FeeAmount** | **string** |  The developer fee for the order.  - **When to set:**     If you are a merchant serving payers directly, set this field to &#x60;0&#x60;.     Developer fees are only relevant for platforms like payment service providers (PSPs) that charge fees to their downstream merchants.   For details, see [Funds allocation and balances](https://www.cobo.com/developers/v2/payments/amounts-and-balances).  - **Denomination:**     The denomination of &#x60;fee_amount&#x60; depends on the presence of &#x60;currency&#x60;:       - If &#x60;currency&#x60; is set, the amount is denominated in that currency (e.g., \&quot;USD\&quot;).       - If &#x60;currency&#x60; is not set, the amount is denominated in the cryptocurrency specified by &#x60;token_id&#x60; (e.g., \&quot;ETH_USDT\&quot;).  - **Calculation:**     The developer fee is added to the base order amount (&#x60;order_amount&#x60;) to determine the final amount charged to the customer.     For example:       - Base amount (&#x60;order_amount&#x60;): \&quot;100.00\&quot;       - Developer fee (&#x60;fee_amount&#x60;): \&quot;2.00\&quot;       - **Total charged:** \&quot;102.00\&quot;  - **Formatting:**     The value can contain up to two decimal places.  | 
**MerchantOrderCode** | Pointer to **string** | A unique reference code assigned by the merchant to identify this order in their system. The code should have a maximum length of 128 characters. | [optional] 
**PspOrderCode** | **string** | A unique reference code assigned by you as a developer to identify this order in your system. This code must be unique across all orders in your system. The code should have a maximum length of 128 characters.  | 
**ExpiredIn** | Pointer to **int32** | The number of seconds until the pay-in order expires, counted from when the request is sent. For example, if set to &#x60;1800&#x60;, the order will expire in 30 minutes. Must be greater than zero and cannot exceed 3 hours (10800 seconds). After expiration:  - The order status becomes final and cannot be changed - The &#x60;received_token_amount&#x60; field will no longer be updated - Funds received after expiration will be categorized as late payments and can only be settled from the developer balance. - A late payment will trigger a &#x60;transactionLate&#x60; webhook event.  | [optional] [default to 1800]
**UseDedicatedAddress** | Pointer to **bool** | This field has been deprecated.  | [optional] 
**CustomExchangeRate** | Pointer to **string** |  A custom exchange rate that defines how much fiat currency equals 1 unit of cryptocurrency. If not provided, the system&#39;s default exchange rate will be used.  For example, if the fiat currency is USD and the cryptocurrency is USDT, setting &#x60;custom_exchange_rate&#x60; to &#x60;\&quot;0.99\&quot;&#x60; means that 1 USDT will be valued at 0.99 USD.  | [optional] 
**AmountTolerance** | Pointer to **string** | The maximum allowed deviation from the payable amount in the case of underpayment, specified as a positive value with up to one decimal place. If you provide more than one decimal place, an error will occur.  When the actual received amount is within this deviation (inclusive) of the payable amount, the order status will be set to &#x60;Completed&#x60; rather than &#x60;Underpaid&#x60;.  | [optional] 

## Methods

### NewCreatePaymentOrderRequest

`func NewCreatePaymentOrderRequest(merchantId string, tokenId string, orderAmount string, feeAmount string, pspOrderCode string, ) *CreatePaymentOrderRequest`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


