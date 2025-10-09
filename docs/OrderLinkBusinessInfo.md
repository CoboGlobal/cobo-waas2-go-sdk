# OrderLinkBusinessInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenIds** | **[]string** | List of supported cryptocurrency token IDs for this payment. Each token ID must be from the supported values.  | 
**CustomExchangeRates** | Pointer to [**[]OrderLinkBusinessInfoCustomExchangeRatesInner**](OrderLinkBusinessInfoCustomExchangeRatesInner.md) | Optional list of final exchange rates for different tokens. If provided, these rates will be used instead of real-time market rates.  | [optional] 
**Currency** | Pointer to **string** | The fiat currency for the base order amount and the developer fee. Currently, only &#x60;USD&#x60; is supported.  If left empty, both &#x60;order_amount&#x60; and &#x60;fee_amount&#x60; will be denominated in the cryptocurrency specified by &#x60;token_id&#x60;  | [optional] [default to ""]
**FeeAmount** | **string** | The developer fee for the order, in the currency specified by &#x60;currency&#x60;. If &#x60;currency&#x60; is not specified, the fee is in the cryptocurrency specified by &#x60;token_id&#x60;.  If you are a merchant directly serving payers, set this field to &#x60;0&#x60;. Developer fees are only relevant for platforms like payment service providers (PSPs) that charge fees to their downstream merchants.  The developer fee is added to the base amount (&#x60;order_amount&#x60;) to determine the final charge. For example: - Base amount (&#x60;order_amount&#x60;): \&quot;100.00\&quot; - Developer fee (&#x60;fee_amount&#x60;): \&quot;2.00\&quot;  - Total charged to customer: \&quot;102.00\&quot;  Values can contain up to two decimal places.  | 
**MerchantId** | **string** | The merchant ID. | 
**OrderAmount** | **string** | The base amount of the order, excluding the developer fee (specified in &#x60;fee_amount&#x60;), in the currency specified by &#x60;currency&#x60;. If &#x60;currency&#x60; is not specified, the amount is in the cryptocurrency specified by &#x60;token_id&#x60;.   Values must be greater than &#x60;0&#x60; and contain two decimal places.   | 
**MerchantOrderCode** | Pointer to **string** | A unique reference code assigned by the merchant to identify this order in their system. The code should have a maximum length of 128 characters. | [optional] 
**PspOrderCode** | **string** | A unique reference code assigned by you as a developer to identify this order in your system. This code must be unique across all orders in your system. The code should have a maximum length of 128 characters.  | 
**ExpiredIn** | Pointer to **int32** | The number of seconds until the pay-in order expires, counted from when the request is sent. For example, if set to &#x60;1800&#x60;, the order will expire in 30 minutes. Must be greater than zero and cannot exceed 3 hours (10800 seconds). After expiration:  - The order status becomes final and cannot be changed - The &#x60;received_token_amount&#x60; field will no longer be updated - Funds received after expiration will be categorized as late payments and can only be settled from the developer balance. - A late payment will trigger a &#x60;transactionLate&#x60; webhook event.  | [optional] [default to 1800]
**UseDedicatedAddress** | Pointer to **bool** | Whether to allocate a dedicated address for this order.  - &#x60;true&#x60;: A dedicated address will be allocated for this order. - &#x60;false&#x60;: A shared address from the address pool will be used.  | [optional] 

## Methods

### NewOrderLinkBusinessInfo

`func NewOrderLinkBusinessInfo(tokenIds []string, feeAmount string, merchantId string, orderAmount string, pspOrderCode string, ) *OrderLinkBusinessInfo`

NewOrderLinkBusinessInfo instantiates a new OrderLinkBusinessInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderLinkBusinessInfoWithDefaults

`func NewOrderLinkBusinessInfoWithDefaults() *OrderLinkBusinessInfo`

NewOrderLinkBusinessInfoWithDefaults instantiates a new OrderLinkBusinessInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


