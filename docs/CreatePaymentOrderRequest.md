# CreatePaymentOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | The merchant ID. | 
**TokenId** | **string** | The ID of the cryptocurrency used for payment. Supported values:    - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDC&#x60;, &#x60;SOL_USDC&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC&#x60;, &#x60;BSC_USDC&#x60;   - USDT: &#x60;TRON_USDT&#x60;, &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;SOL_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**Currency** | Pointer to **string** | The fiat currency of the order. | [optional] [default to "USD"]
**OrderAmount** | **string** | The base amount of the order in fiat currency, excluding the developer fee (specified in &#x60;fee_amount&#x60;). Values must be greater than &#x60;0&#x60; and contain two decimal places. | 
**FeeAmount** | **string** | The developer fee for the order in fiat currency. It is added to the base amount (&#x60;order_amount&#x60;) to determine the final charge. For example, if order_amount is \&quot;100.00\&quot; and fee_amount is \&quot;2.00\&quot;, the customer will be charged \&quot;102.00\&quot; in total, with \&quot;100.00\&quot; being settled to the merchant and \&quot;2.00\&quot; settled to the developer. Values must be greater than 0 and contain two decimal places. | 
**MerchantOrderCode** | Pointer to **string** | A unique reference code assigned by the merchant to identify this order in their system. | [optional] 
**PspOrderCode** | **string** | A unique reference code assigned by the developer to identify this order in their system. | 
**ExpiredIn** | Pointer to **int32** | The number of seconds after which the pay-in order will expire. After expiration: - The order status becomes final and cannot be changed - The &#x60;received_token_amount&#x60; field will no longer be updated - Funds received after expiration will be categorized as late payments and can only be settled from the developer balance. - A late payment will trigger a &#x60;transactionLate&#x60; webhook event.  | [optional] 
**UseDedicatedAddress** | Pointer to **bool** | Whether to allocate a dedicated address for this order.  - &#x60;true&#x60;: A dedicated address will be allocated for this order. - &#x60;false&#x60;: A shared address from the address pool will be used.  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


