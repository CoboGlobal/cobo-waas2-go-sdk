# CreateSwapQuote201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayAmount** | **string** | The amount of tokens to pay. | 
**ReceiveAmount** | **string** | The amount of tokens to receive. | 
**FeeAmount** | **string** | The amount of tokens to pay for fee. | 
**MinPayAmount** | Pointer to **string** | The minimum amount of tokens to pay. | [optional] 
**MaxPayAmount** | Pointer to **string** | The maximum amount of tokens to pay. | [optional] 
**MinReceiveAmount** | Pointer to **string** | The minimum amount of tokens to receive. | [optional] 
**MaxReceiveAmount** | Pointer to **string** | The maximum amount of tokens to receive. | [optional] 
**QuoteExpiredTimestamp** | **int32** | The time when the quote will expire, in Unix timestamp format, measured in milliseconds. | 
**QuoteId** | **string** | The unique identifier of this quote. | 

## Methods

### NewCreateSwapQuote201Response

`func NewCreateSwapQuote201Response(payAmount string, receiveAmount string, feeAmount string, quoteExpiredTimestamp int32, quoteId string, ) *CreateSwapQuote201Response`

NewCreateSwapQuote201Response instantiates a new CreateSwapQuote201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSwapQuote201ResponseWithDefaults

`func NewCreateSwapQuote201ResponseWithDefaults() *CreateSwapQuote201Response`

NewCreateSwapQuote201ResponseWithDefaults instantiates a new CreateSwapQuote201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayAmount

`func (o *CreateSwapQuote201Response) GetPayAmount() string`

GetPayAmount returns the PayAmount field if non-nil, zero value otherwise.

### GetPayAmountOk

`func (o *CreateSwapQuote201Response) GetPayAmountOk() (*string, bool)`

GetPayAmountOk returns a tuple with the PayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayAmount

`func (o *CreateSwapQuote201Response) SetPayAmount(v string)`

SetPayAmount sets PayAmount field to given value.


### GetReceiveAmount

`func (o *CreateSwapQuote201Response) GetReceiveAmount() string`

GetReceiveAmount returns the ReceiveAmount field if non-nil, zero value otherwise.

### GetReceiveAmountOk

`func (o *CreateSwapQuote201Response) GetReceiveAmountOk() (*string, bool)`

GetReceiveAmountOk returns a tuple with the ReceiveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveAmount

`func (o *CreateSwapQuote201Response) SetReceiveAmount(v string)`

SetReceiveAmount sets ReceiveAmount field to given value.


### GetFeeAmount

`func (o *CreateSwapQuote201Response) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *CreateSwapQuote201Response) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *CreateSwapQuote201Response) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetMinPayAmount

`func (o *CreateSwapQuote201Response) GetMinPayAmount() string`

GetMinPayAmount returns the MinPayAmount field if non-nil, zero value otherwise.

### GetMinPayAmountOk

`func (o *CreateSwapQuote201Response) GetMinPayAmountOk() (*string, bool)`

GetMinPayAmountOk returns a tuple with the MinPayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPayAmount

`func (o *CreateSwapQuote201Response) SetMinPayAmount(v string)`

SetMinPayAmount sets MinPayAmount field to given value.

### HasMinPayAmount

`func (o *CreateSwapQuote201Response) HasMinPayAmount() bool`

HasMinPayAmount returns a boolean if a field has been set.

### GetMaxPayAmount

`func (o *CreateSwapQuote201Response) GetMaxPayAmount() string`

GetMaxPayAmount returns the MaxPayAmount field if non-nil, zero value otherwise.

### GetMaxPayAmountOk

`func (o *CreateSwapQuote201Response) GetMaxPayAmountOk() (*string, bool)`

GetMaxPayAmountOk returns a tuple with the MaxPayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayAmount

`func (o *CreateSwapQuote201Response) SetMaxPayAmount(v string)`

SetMaxPayAmount sets MaxPayAmount field to given value.

### HasMaxPayAmount

`func (o *CreateSwapQuote201Response) HasMaxPayAmount() bool`

HasMaxPayAmount returns a boolean if a field has been set.

### GetMinReceiveAmount

`func (o *CreateSwapQuote201Response) GetMinReceiveAmount() string`

GetMinReceiveAmount returns the MinReceiveAmount field if non-nil, zero value otherwise.

### GetMinReceiveAmountOk

`func (o *CreateSwapQuote201Response) GetMinReceiveAmountOk() (*string, bool)`

GetMinReceiveAmountOk returns a tuple with the MinReceiveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReceiveAmount

`func (o *CreateSwapQuote201Response) SetMinReceiveAmount(v string)`

SetMinReceiveAmount sets MinReceiveAmount field to given value.

### HasMinReceiveAmount

`func (o *CreateSwapQuote201Response) HasMinReceiveAmount() bool`

HasMinReceiveAmount returns a boolean if a field has been set.

### GetMaxReceiveAmount

`func (o *CreateSwapQuote201Response) GetMaxReceiveAmount() string`

GetMaxReceiveAmount returns the MaxReceiveAmount field if non-nil, zero value otherwise.

### GetMaxReceiveAmountOk

`func (o *CreateSwapQuote201Response) GetMaxReceiveAmountOk() (*string, bool)`

GetMaxReceiveAmountOk returns a tuple with the MaxReceiveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReceiveAmount

`func (o *CreateSwapQuote201Response) SetMaxReceiveAmount(v string)`

SetMaxReceiveAmount sets MaxReceiveAmount field to given value.

### HasMaxReceiveAmount

`func (o *CreateSwapQuote201Response) HasMaxReceiveAmount() bool`

HasMaxReceiveAmount returns a boolean if a field has been set.

### GetQuoteExpiredTimestamp

`func (o *CreateSwapQuote201Response) GetQuoteExpiredTimestamp() int32`

GetQuoteExpiredTimestamp returns the QuoteExpiredTimestamp field if non-nil, zero value otherwise.

### GetQuoteExpiredTimestampOk

`func (o *CreateSwapQuote201Response) GetQuoteExpiredTimestampOk() (*int32, bool)`

GetQuoteExpiredTimestampOk returns a tuple with the QuoteExpiredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteExpiredTimestamp

`func (o *CreateSwapQuote201Response) SetQuoteExpiredTimestamp(v int32)`

SetQuoteExpiredTimestamp sets QuoteExpiredTimestamp field to given value.


### GetQuoteId

`func (o *CreateSwapQuote201Response) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *CreateSwapQuote201Response) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *CreateSwapQuote201Response) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


