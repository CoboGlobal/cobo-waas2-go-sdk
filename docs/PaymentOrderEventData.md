# PaymentOrderEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  The data type of the event. - &#x60;Transaction&#x60;: The transaction event data. - &#x60;TSSRequest&#x60;: The TSS request event data. - &#x60;Addresses&#x60;: The addresses event data. - &#x60;WalletInfo&#x60;: The wallet information event data. - &#x60;MPCVault&#x60;: The MPC vault event data. - &#x60;Chains&#x60;: The enabled chain event data. - &#x60;Tokens&#x60;: The enabled token event data. - &#x60;TokenListing&#x60;: The token listing event data.        - &#x60;PaymentOrder&#x60;: The payment order event data. - &#x60;PaymentRefund&#x60;: The payment refund event data. - &#x60;PaymentSettlement&#x60;: The payment settlement event data. | 
**OrderId** | **string** | The order ID. | 
**MerchantId** | Pointer to **string** | The merchant ID. | [optional] 
**TokenId** | **string** | The ID of the cryptocurrency used for payment. | 
**ChainId** | **string** | The ID of the blockchain network where the payment transaction should be made. | 
**PayableAmount** | **string** | The cryptocurrency amount to be paid for this order. | 
**ReceiveAddress** | **string** | The recipient wallet address to be used for the payment transaction. | 
**Currency** | **string** | The fiat currency of the order. | 
**OrderAmount** | **string** | The base amount of the order in fiat currency, excluding the developer fee (specified in &#x60;fee_amount&#x60;). | 
**FeeAmount** | **string** | The developer fee for the order in fiat currency. It is added to the base amount (&#x60;order_amount&#x60;) to determine the final charge. | 
**ExchangeRate** | **string** | The exchange rate between a currency pair. Expressed as the amount of fiat currency per one unit of cryptocurrency. For example, if the cryptocurrency is USDT and the fiat currency is USD, a rate of \&quot;0.99\&quot; means 1 USDT &#x3D; 0.99 USD. | 
**ExpiredAt** | Pointer to **int32** | The expiration time of the pay-in order, represented as a UNIX timestamp in seconds. | [optional] 
**MerchantOrderCode** | Pointer to **string** | A unique reference code assigned by the merchant to identify this order in their system. | [optional] 
**PspOrderCode** | **string** | A unique reference code assigned by the developer to identify this order in their system. | 
**Status** | [**OrderStatus**](OrderStatus.md) |  | 
**ReceivedTokenAmount** | **string** | The total cryptocurrency amount received for this order. Updates until the expiration time. Precision matches the token standard (e.g., 6 decimals for USDT). | 
**CreatedTimestamp** | Pointer to **int32** | The created time of the order, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The updated time of the order, represented as a UNIX timestamp in seconds. | [optional] 
**Transactions** | Pointer to [**[]PaymentTransaction**](PaymentTransaction.md) | An array of transactions associated with this pay-in order. Each transaction represents a separate blockchain operation related to the settlement process. | [optional] 

## Methods

### NewPaymentOrderEventData

`func NewPaymentOrderEventData(dataType string, orderId string, tokenId string, chainId string, payableAmount string, receiveAddress string, currency string, orderAmount string, feeAmount string, exchangeRate string, pspOrderCode string, status OrderStatus, receivedTokenAmount string, ) *PaymentOrderEventData`

NewPaymentOrderEventData instantiates a new PaymentOrderEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOrderEventDataWithDefaults

`func NewPaymentOrderEventDataWithDefaults() *PaymentOrderEventData`

NewPaymentOrderEventDataWithDefaults instantiates a new PaymentOrderEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *PaymentOrderEventData) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PaymentOrderEventData) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PaymentOrderEventData) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetOrderId

`func (o *PaymentOrderEventData) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PaymentOrderEventData) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PaymentOrderEventData) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetMerchantId

`func (o *PaymentOrderEventData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *PaymentOrderEventData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *PaymentOrderEventData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *PaymentOrderEventData) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetTokenId

`func (o *PaymentOrderEventData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentOrderEventData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentOrderEventData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *PaymentOrderEventData) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *PaymentOrderEventData) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *PaymentOrderEventData) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetPayableAmount

`func (o *PaymentOrderEventData) GetPayableAmount() string`

GetPayableAmount returns the PayableAmount field if non-nil, zero value otherwise.

### GetPayableAmountOk

`func (o *PaymentOrderEventData) GetPayableAmountOk() (*string, bool)`

GetPayableAmountOk returns a tuple with the PayableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableAmount

`func (o *PaymentOrderEventData) SetPayableAmount(v string)`

SetPayableAmount sets PayableAmount field to given value.


### GetReceiveAddress

`func (o *PaymentOrderEventData) GetReceiveAddress() string`

GetReceiveAddress returns the ReceiveAddress field if non-nil, zero value otherwise.

### GetReceiveAddressOk

`func (o *PaymentOrderEventData) GetReceiveAddressOk() (*string, bool)`

GetReceiveAddressOk returns a tuple with the ReceiveAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveAddress

`func (o *PaymentOrderEventData) SetReceiveAddress(v string)`

SetReceiveAddress sets ReceiveAddress field to given value.


### GetCurrency

`func (o *PaymentOrderEventData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentOrderEventData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentOrderEventData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetOrderAmount

`func (o *PaymentOrderEventData) GetOrderAmount() string`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *PaymentOrderEventData) GetOrderAmountOk() (*string, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *PaymentOrderEventData) SetOrderAmount(v string)`

SetOrderAmount sets OrderAmount field to given value.


### GetFeeAmount

`func (o *PaymentOrderEventData) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *PaymentOrderEventData) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *PaymentOrderEventData) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetExchangeRate

`func (o *PaymentOrderEventData) GetExchangeRate() string`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *PaymentOrderEventData) GetExchangeRateOk() (*string, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *PaymentOrderEventData) SetExchangeRate(v string)`

SetExchangeRate sets ExchangeRate field to given value.


### GetExpiredAt

`func (o *PaymentOrderEventData) GetExpiredAt() int32`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *PaymentOrderEventData) GetExpiredAtOk() (*int32, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *PaymentOrderEventData) SetExpiredAt(v int32)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *PaymentOrderEventData) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetMerchantOrderCode

`func (o *PaymentOrderEventData) GetMerchantOrderCode() string`

GetMerchantOrderCode returns the MerchantOrderCode field if non-nil, zero value otherwise.

### GetMerchantOrderCodeOk

`func (o *PaymentOrderEventData) GetMerchantOrderCodeOk() (*string, bool)`

GetMerchantOrderCodeOk returns a tuple with the MerchantOrderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderCode

`func (o *PaymentOrderEventData) SetMerchantOrderCode(v string)`

SetMerchantOrderCode sets MerchantOrderCode field to given value.

### HasMerchantOrderCode

`func (o *PaymentOrderEventData) HasMerchantOrderCode() bool`

HasMerchantOrderCode returns a boolean if a field has been set.

### GetPspOrderCode

`func (o *PaymentOrderEventData) GetPspOrderCode() string`

GetPspOrderCode returns the PspOrderCode field if non-nil, zero value otherwise.

### GetPspOrderCodeOk

`func (o *PaymentOrderEventData) GetPspOrderCodeOk() (*string, bool)`

GetPspOrderCodeOk returns a tuple with the PspOrderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspOrderCode

`func (o *PaymentOrderEventData) SetPspOrderCode(v string)`

SetPspOrderCode sets PspOrderCode field to given value.


### GetStatus

`func (o *PaymentOrderEventData) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentOrderEventData) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentOrderEventData) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.


### GetReceivedTokenAmount

`func (o *PaymentOrderEventData) GetReceivedTokenAmount() string`

GetReceivedTokenAmount returns the ReceivedTokenAmount field if non-nil, zero value otherwise.

### GetReceivedTokenAmountOk

`func (o *PaymentOrderEventData) GetReceivedTokenAmountOk() (*string, bool)`

GetReceivedTokenAmountOk returns a tuple with the ReceivedTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTokenAmount

`func (o *PaymentOrderEventData) SetReceivedTokenAmount(v string)`

SetReceivedTokenAmount sets ReceivedTokenAmount field to given value.


### GetCreatedTimestamp

`func (o *PaymentOrderEventData) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentOrderEventData) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentOrderEventData) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *PaymentOrderEventData) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *PaymentOrderEventData) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PaymentOrderEventData) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PaymentOrderEventData) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *PaymentOrderEventData) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetTransactions

`func (o *PaymentOrderEventData) GetTransactions() []PaymentTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *PaymentOrderEventData) GetTransactionsOk() (*[]PaymentTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *PaymentOrderEventData) SetTransactions(v []PaymentTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *PaymentOrderEventData) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


