# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | The order ID. | 
**MerchantId** | Pointer to **string** | The merchant ID. | [optional] 
**MerchantOrderCode** | Pointer to **string** | A unique reference code assigned by the merchant to identify this order in their system. | [optional] 
**PspOrderCode** | **string** | A unique reference code assigned by the developer to identify this order in their system. | 
**PricingCurrency** | Pointer to **string** | The pricing currency of the order. | [optional] 
**PricingAmount** | Pointer to **string** | The base amount of the order, excluding the developer fee (specified in &#x60;fee_amount&#x60;). | [optional] 
**FeeAmount** | **string** | The developer fee for the order. It is added to the base amount to determine the final charge. | 
**PayableCurrency** | Pointer to **string** | The ID of the cryptocurrency used for payment. | [optional] 
**ChainId** | **string** | The ID of the blockchain network where the payment transaction should be made. | 
**PayableAmount** | **string** | The cryptocurrency amount to be paid for this order. | 
**ExchangeRate** | **string** | The exchange rate between &#x60;payable_currency&#x60; and &#x60;pricing_currency&#x60;, calculated as (&#x60;pricing_amount&#x60; + &#x60;fee_amount&#x60;) / &#x60;payable_amount&#x60;.    &lt;Note&gt;This field is only returned when &#x60;payable_amount&#x60; was not provided in the order creation request. &lt;/Note&gt;  | 
**AmountTolerance** | Pointer to **string** | The allowed amount deviation, with precision up to 1 decimal place.  For example, if &#x60;payable_amount&#x60; is &#x60;100.00&#x60; and &#x60;amount_tolerance&#x60; is &#x60;0.50&#x60;: - Payer pays 99.55 → Success (difference of 0.45 ≤ 0.5) - Payer pays 99.40 → Underpaid (difference of 0.60 &gt; 0.5)  | [optional] 
**ReceiveAddress** | **string** | The recipient wallet address to be used for the payment transaction. | 
**Status** | [**OrderStatus**](OrderStatus.md) |  | 
**ReceivedTokenAmount** | **string** | The total cryptocurrency amount received for this order. Updates until the expiration time. Precision matches the token standard (e.g., 6 decimals for USDT). | 
**ExpiredAt** | Pointer to **int32** | The expiration time of the pay-in order, represented as a UNIX timestamp in seconds. | [optional] 
**CreatedTimestamp** | Pointer to **int32** | The created time of the order, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The updated time of the order, represented as a UNIX timestamp in seconds. | [optional] 
**Transactions** | Pointer to [**[]PaymentTransaction**](PaymentTransaction.md) | An array of transactions associated with this pay-in order. Each transaction represents a separate blockchain operation related to the settlement process. | [optional] 
**Currency** | Pointer to **string** | This field has been deprecated. Please use &#x60;pricing_currency&#x60; instead. | [optional] 
**OrderAmount** | Pointer to **string** | This field has been deprecated. Please use &#x60;pricing_amount&#x60; instead. | [optional] 
**TokenId** | Pointer to **string** | This field has been deprecated. Please use &#x60;payable_currency&#x60; instead. | [optional] 
**SettlementStatus** | Pointer to [**SettleStatus**](SettleStatus.md) |  | [optional] 

## Methods

### NewOrder

`func NewOrder(orderId string, pspOrderCode string, feeAmount string, chainId string, payableAmount string, exchangeRate string, receiveAddress string, status OrderStatus, receivedTokenAmount string, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *Order) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Order) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Order) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetMerchantId

`func (o *Order) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *Order) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *Order) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *Order) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMerchantOrderCode

`func (o *Order) GetMerchantOrderCode() string`

GetMerchantOrderCode returns the MerchantOrderCode field if non-nil, zero value otherwise.

### GetMerchantOrderCodeOk

`func (o *Order) GetMerchantOrderCodeOk() (*string, bool)`

GetMerchantOrderCodeOk returns a tuple with the MerchantOrderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderCode

`func (o *Order) SetMerchantOrderCode(v string)`

SetMerchantOrderCode sets MerchantOrderCode field to given value.

### HasMerchantOrderCode

`func (o *Order) HasMerchantOrderCode() bool`

HasMerchantOrderCode returns a boolean if a field has been set.

### GetPspOrderCode

`func (o *Order) GetPspOrderCode() string`

GetPspOrderCode returns the PspOrderCode field if non-nil, zero value otherwise.

### GetPspOrderCodeOk

`func (o *Order) GetPspOrderCodeOk() (*string, bool)`

GetPspOrderCodeOk returns a tuple with the PspOrderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspOrderCode

`func (o *Order) SetPspOrderCode(v string)`

SetPspOrderCode sets PspOrderCode field to given value.


### GetPricingCurrency

`func (o *Order) GetPricingCurrency() string`

GetPricingCurrency returns the PricingCurrency field if non-nil, zero value otherwise.

### GetPricingCurrencyOk

`func (o *Order) GetPricingCurrencyOk() (*string, bool)`

GetPricingCurrencyOk returns a tuple with the PricingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingCurrency

`func (o *Order) SetPricingCurrency(v string)`

SetPricingCurrency sets PricingCurrency field to given value.

### HasPricingCurrency

`func (o *Order) HasPricingCurrency() bool`

HasPricingCurrency returns a boolean if a field has been set.

### GetPricingAmount

`func (o *Order) GetPricingAmount() string`

GetPricingAmount returns the PricingAmount field if non-nil, zero value otherwise.

### GetPricingAmountOk

`func (o *Order) GetPricingAmountOk() (*string, bool)`

GetPricingAmountOk returns a tuple with the PricingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingAmount

`func (o *Order) SetPricingAmount(v string)`

SetPricingAmount sets PricingAmount field to given value.

### HasPricingAmount

`func (o *Order) HasPricingAmount() bool`

HasPricingAmount returns a boolean if a field has been set.

### GetFeeAmount

`func (o *Order) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *Order) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *Order) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetPayableCurrency

`func (o *Order) GetPayableCurrency() string`

GetPayableCurrency returns the PayableCurrency field if non-nil, zero value otherwise.

### GetPayableCurrencyOk

`func (o *Order) GetPayableCurrencyOk() (*string, bool)`

GetPayableCurrencyOk returns a tuple with the PayableCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableCurrency

`func (o *Order) SetPayableCurrency(v string)`

SetPayableCurrency sets PayableCurrency field to given value.

### HasPayableCurrency

`func (o *Order) HasPayableCurrency() bool`

HasPayableCurrency returns a boolean if a field has been set.

### GetChainId

`func (o *Order) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Order) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Order) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetPayableAmount

`func (o *Order) GetPayableAmount() string`

GetPayableAmount returns the PayableAmount field if non-nil, zero value otherwise.

### GetPayableAmountOk

`func (o *Order) GetPayableAmountOk() (*string, bool)`

GetPayableAmountOk returns a tuple with the PayableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableAmount

`func (o *Order) SetPayableAmount(v string)`

SetPayableAmount sets PayableAmount field to given value.


### GetExchangeRate

`func (o *Order) GetExchangeRate() string`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *Order) GetExchangeRateOk() (*string, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *Order) SetExchangeRate(v string)`

SetExchangeRate sets ExchangeRate field to given value.


### GetAmountTolerance

`func (o *Order) GetAmountTolerance() string`

GetAmountTolerance returns the AmountTolerance field if non-nil, zero value otherwise.

### GetAmountToleranceOk

`func (o *Order) GetAmountToleranceOk() (*string, bool)`

GetAmountToleranceOk returns a tuple with the AmountTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountTolerance

`func (o *Order) SetAmountTolerance(v string)`

SetAmountTolerance sets AmountTolerance field to given value.

### HasAmountTolerance

`func (o *Order) HasAmountTolerance() bool`

HasAmountTolerance returns a boolean if a field has been set.

### GetReceiveAddress

`func (o *Order) GetReceiveAddress() string`

GetReceiveAddress returns the ReceiveAddress field if non-nil, zero value otherwise.

### GetReceiveAddressOk

`func (o *Order) GetReceiveAddressOk() (*string, bool)`

GetReceiveAddressOk returns a tuple with the ReceiveAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveAddress

`func (o *Order) SetReceiveAddress(v string)`

SetReceiveAddress sets ReceiveAddress field to given value.


### GetStatus

`func (o *Order) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.


### GetReceivedTokenAmount

`func (o *Order) GetReceivedTokenAmount() string`

GetReceivedTokenAmount returns the ReceivedTokenAmount field if non-nil, zero value otherwise.

### GetReceivedTokenAmountOk

`func (o *Order) GetReceivedTokenAmountOk() (*string, bool)`

GetReceivedTokenAmountOk returns a tuple with the ReceivedTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTokenAmount

`func (o *Order) SetReceivedTokenAmount(v string)`

SetReceivedTokenAmount sets ReceivedTokenAmount field to given value.


### GetExpiredAt

`func (o *Order) GetExpiredAt() int32`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *Order) GetExpiredAtOk() (*int32, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *Order) SetExpiredAt(v int32)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *Order) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *Order) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Order) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Order) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *Order) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *Order) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Order) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Order) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *Order) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetTransactions

`func (o *Order) GetTransactions() []PaymentTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *Order) GetTransactionsOk() (*[]PaymentTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *Order) SetTransactions(v []PaymentTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *Order) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetCurrency

`func (o *Order) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Order) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Order) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Order) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetOrderAmount

`func (o *Order) GetOrderAmount() string`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *Order) GetOrderAmountOk() (*string, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *Order) SetOrderAmount(v string)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *Order) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetTokenId

`func (o *Order) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Order) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Order) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *Order) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetSettlementStatus

`func (o *Order) GetSettlementStatus() SettleStatus`

GetSettlementStatus returns the SettlementStatus field if non-nil, zero value otherwise.

### GetSettlementStatusOk

`func (o *Order) GetSettlementStatusOk() (*SettleStatus, bool)`

GetSettlementStatusOk returns a tuple with the SettlementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementStatus

`func (o *Order) SetSettlementStatus(v SettleStatus)`

SetSettlementStatus sets SettlementStatus field to given value.

### HasSettlementStatus

`func (o *Order) HasSettlementStatus() bool`

HasSettlementStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


