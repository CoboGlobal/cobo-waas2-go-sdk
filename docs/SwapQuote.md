# SwapQuote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** | The unique id of quote. | 
**PayTokenId** | **string** | The token ID to pay. | 
**PayAmount** | **string** | The amount of tokens to pay. | 
**ReceiveTokenId** | **string** | The token ID to receive. | 
**ReceiveAmount** | **string** | The amount of tokens to receive. | 
**FeeTokenId** | **string** | The fee token ID. | 
**FeeAmount** | **string** | The amount of tokens to pay for fee. | 
**MinReceiveAmount** | Pointer to **string** | The minimum amount of tokens to receive if the pay amount is specified. | [optional] 
**MaxPayAmount** | Pointer to **string** | The maximum amount of tokens to pay if the receive amount is specified. | [optional] 
**QuoteExpiredTimestamp** | **int32** | The time when the quote will expire, in Unix timestamp format, measured in milliseconds. | 

## Methods

### NewSwapQuote

`func NewSwapQuote(quoteId string, payTokenId string, payAmount string, receiveTokenId string, receiveAmount string, feeTokenId string, feeAmount string, quoteExpiredTimestamp int32, ) *SwapQuote`

NewSwapQuote instantiates a new SwapQuote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapQuoteWithDefaults

`func NewSwapQuoteWithDefaults() *SwapQuote`

NewSwapQuoteWithDefaults instantiates a new SwapQuote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *SwapQuote) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *SwapQuote) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *SwapQuote) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetPayTokenId

`func (o *SwapQuote) GetPayTokenId() string`

GetPayTokenId returns the PayTokenId field if non-nil, zero value otherwise.

### GetPayTokenIdOk

`func (o *SwapQuote) GetPayTokenIdOk() (*string, bool)`

GetPayTokenIdOk returns a tuple with the PayTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayTokenId

`func (o *SwapQuote) SetPayTokenId(v string)`

SetPayTokenId sets PayTokenId field to given value.


### GetPayAmount

`func (o *SwapQuote) GetPayAmount() string`

GetPayAmount returns the PayAmount field if non-nil, zero value otherwise.

### GetPayAmountOk

`func (o *SwapQuote) GetPayAmountOk() (*string, bool)`

GetPayAmountOk returns a tuple with the PayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayAmount

`func (o *SwapQuote) SetPayAmount(v string)`

SetPayAmount sets PayAmount field to given value.


### GetReceiveTokenId

`func (o *SwapQuote) GetReceiveTokenId() string`

GetReceiveTokenId returns the ReceiveTokenId field if non-nil, zero value otherwise.

### GetReceiveTokenIdOk

`func (o *SwapQuote) GetReceiveTokenIdOk() (*string, bool)`

GetReceiveTokenIdOk returns a tuple with the ReceiveTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveTokenId

`func (o *SwapQuote) SetReceiveTokenId(v string)`

SetReceiveTokenId sets ReceiveTokenId field to given value.


### GetReceiveAmount

`func (o *SwapQuote) GetReceiveAmount() string`

GetReceiveAmount returns the ReceiveAmount field if non-nil, zero value otherwise.

### GetReceiveAmountOk

`func (o *SwapQuote) GetReceiveAmountOk() (*string, bool)`

GetReceiveAmountOk returns a tuple with the ReceiveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveAmount

`func (o *SwapQuote) SetReceiveAmount(v string)`

SetReceiveAmount sets ReceiveAmount field to given value.


### GetFeeTokenId

`func (o *SwapQuote) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *SwapQuote) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *SwapQuote) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.


### GetFeeAmount

`func (o *SwapQuote) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *SwapQuote) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *SwapQuote) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetMinReceiveAmount

`func (o *SwapQuote) GetMinReceiveAmount() string`

GetMinReceiveAmount returns the MinReceiveAmount field if non-nil, zero value otherwise.

### GetMinReceiveAmountOk

`func (o *SwapQuote) GetMinReceiveAmountOk() (*string, bool)`

GetMinReceiveAmountOk returns a tuple with the MinReceiveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReceiveAmount

`func (o *SwapQuote) SetMinReceiveAmount(v string)`

SetMinReceiveAmount sets MinReceiveAmount field to given value.

### HasMinReceiveAmount

`func (o *SwapQuote) HasMinReceiveAmount() bool`

HasMinReceiveAmount returns a boolean if a field has been set.

### GetMaxPayAmount

`func (o *SwapQuote) GetMaxPayAmount() string`

GetMaxPayAmount returns the MaxPayAmount field if non-nil, zero value otherwise.

### GetMaxPayAmountOk

`func (o *SwapQuote) GetMaxPayAmountOk() (*string, bool)`

GetMaxPayAmountOk returns a tuple with the MaxPayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayAmount

`func (o *SwapQuote) SetMaxPayAmount(v string)`

SetMaxPayAmount sets MaxPayAmount field to given value.

### HasMaxPayAmount

`func (o *SwapQuote) HasMaxPayAmount() bool`

HasMaxPayAmount returns a boolean if a field has been set.

### GetQuoteExpiredTimestamp

`func (o *SwapQuote) GetQuoteExpiredTimestamp() int32`

GetQuoteExpiredTimestamp returns the QuoteExpiredTimestamp field if non-nil, zero value otherwise.

### GetQuoteExpiredTimestampOk

`func (o *SwapQuote) GetQuoteExpiredTimestampOk() (*int32, bool)`

GetQuoteExpiredTimestampOk returns a tuple with the QuoteExpiredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteExpiredTimestamp

`func (o *SwapQuote) SetQuoteExpiredTimestamp(v int32)`

SetQuoteExpiredTimestamp sets QuoteExpiredTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


