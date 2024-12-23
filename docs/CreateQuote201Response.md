# CreateQuote201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** | The unique identifier of this quote. | 
**PayAmount** | **string** | The amount of tokens to pay. | 
**PayUsdValue** | **string** | The USD value of pay_amount. | 
**ReceiveAmount** | **string** | The amount of tokens to receive. | 
**ReceiveUsdValue** | **string** | The USD value of receive_amount. | 
**FeeTokenId** | Pointer to **string** | The token to pay for fee. | [optional] 
**FeeAmount** | Pointer to **string** | The amount of tokens to pay for fee. | [optional] 
**FeeUsdValue** | Pointer to **string** | The USD value of fee_amount. | [optional] 
**MinPayAmount** | Pointer to **string** | The minimum amount of tokens to pay. | [optional] 
**MaxPayAmount** | Pointer to **string** | The maximum amount of tokens to pay. | [optional] 
**MinReceiveAmount** | Pointer to **string** | The minimum amount of tokens to receive. | [optional] 
**MaxReceiveAmount** | Pointer to **string** | The maximum amount of tokens to receive. | [optional] 
**QuoteCreatedTime** | **int32** | The time when the quote was created, in Unix timestamp format, measured in milliseconds. | 
**QuoteExpiredTime** | **int32** | The time when the quote will expire, in Unix timestamp format, measured in milliseconds. | 

## Methods

### NewCreateQuote201Response

`func NewCreateQuote201Response(quoteId string, payAmount string, payUsdValue string, receiveAmount string, receiveUsdValue string, quoteCreatedTime int32, quoteExpiredTime int32, ) *CreateQuote201Response`

NewCreateQuote201Response instantiates a new CreateQuote201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuote201ResponseWithDefaults

`func NewCreateQuote201ResponseWithDefaults() *CreateQuote201Response`

NewCreateQuote201ResponseWithDefaults instantiates a new CreateQuote201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *CreateQuote201Response) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *CreateQuote201Response) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *CreateQuote201Response) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetPayAmount

`func (o *CreateQuote201Response) GetPayAmount() string`

GetPayAmount returns the PayAmount field if non-nil, zero value otherwise.

### GetPayAmountOk

`func (o *CreateQuote201Response) GetPayAmountOk() (*string, bool)`

GetPayAmountOk returns a tuple with the PayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayAmount

`func (o *CreateQuote201Response) SetPayAmount(v string)`

SetPayAmount sets PayAmount field to given value.


### GetPayUsdValue

`func (o *CreateQuote201Response) GetPayUsdValue() string`

GetPayUsdValue returns the PayUsdValue field if non-nil, zero value otherwise.

### GetPayUsdValueOk

`func (o *CreateQuote201Response) GetPayUsdValueOk() (*string, bool)`

GetPayUsdValueOk returns a tuple with the PayUsdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayUsdValue

`func (o *CreateQuote201Response) SetPayUsdValue(v string)`

SetPayUsdValue sets PayUsdValue field to given value.


### GetReceiveAmount

`func (o *CreateQuote201Response) GetReceiveAmount() string`

GetReceiveAmount returns the ReceiveAmount field if non-nil, zero value otherwise.

### GetReceiveAmountOk

`func (o *CreateQuote201Response) GetReceiveAmountOk() (*string, bool)`

GetReceiveAmountOk returns a tuple with the ReceiveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveAmount

`func (o *CreateQuote201Response) SetReceiveAmount(v string)`

SetReceiveAmount sets ReceiveAmount field to given value.


### GetReceiveUsdValue

`func (o *CreateQuote201Response) GetReceiveUsdValue() string`

GetReceiveUsdValue returns the ReceiveUsdValue field if non-nil, zero value otherwise.

### GetReceiveUsdValueOk

`func (o *CreateQuote201Response) GetReceiveUsdValueOk() (*string, bool)`

GetReceiveUsdValueOk returns a tuple with the ReceiveUsdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveUsdValue

`func (o *CreateQuote201Response) SetReceiveUsdValue(v string)`

SetReceiveUsdValue sets ReceiveUsdValue field to given value.


### GetFeeTokenId

`func (o *CreateQuote201Response) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *CreateQuote201Response) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *CreateQuote201Response) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *CreateQuote201Response) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetFeeAmount

`func (o *CreateQuote201Response) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *CreateQuote201Response) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *CreateQuote201Response) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *CreateQuote201Response) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetFeeUsdValue

`func (o *CreateQuote201Response) GetFeeUsdValue() string`

GetFeeUsdValue returns the FeeUsdValue field if non-nil, zero value otherwise.

### GetFeeUsdValueOk

`func (o *CreateQuote201Response) GetFeeUsdValueOk() (*string, bool)`

GetFeeUsdValueOk returns a tuple with the FeeUsdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeUsdValue

`func (o *CreateQuote201Response) SetFeeUsdValue(v string)`

SetFeeUsdValue sets FeeUsdValue field to given value.

### HasFeeUsdValue

`func (o *CreateQuote201Response) HasFeeUsdValue() bool`

HasFeeUsdValue returns a boolean if a field has been set.

### GetMinPayAmount

`func (o *CreateQuote201Response) GetMinPayAmount() string`

GetMinPayAmount returns the MinPayAmount field if non-nil, zero value otherwise.

### GetMinPayAmountOk

`func (o *CreateQuote201Response) GetMinPayAmountOk() (*string, bool)`

GetMinPayAmountOk returns a tuple with the MinPayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPayAmount

`func (o *CreateQuote201Response) SetMinPayAmount(v string)`

SetMinPayAmount sets MinPayAmount field to given value.

### HasMinPayAmount

`func (o *CreateQuote201Response) HasMinPayAmount() bool`

HasMinPayAmount returns a boolean if a field has been set.

### GetMaxPayAmount

`func (o *CreateQuote201Response) GetMaxPayAmount() string`

GetMaxPayAmount returns the MaxPayAmount field if non-nil, zero value otherwise.

### GetMaxPayAmountOk

`func (o *CreateQuote201Response) GetMaxPayAmountOk() (*string, bool)`

GetMaxPayAmountOk returns a tuple with the MaxPayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayAmount

`func (o *CreateQuote201Response) SetMaxPayAmount(v string)`

SetMaxPayAmount sets MaxPayAmount field to given value.

### HasMaxPayAmount

`func (o *CreateQuote201Response) HasMaxPayAmount() bool`

HasMaxPayAmount returns a boolean if a field has been set.

### GetMinReceiveAmount

`func (o *CreateQuote201Response) GetMinReceiveAmount() string`

GetMinReceiveAmount returns the MinReceiveAmount field if non-nil, zero value otherwise.

### GetMinReceiveAmountOk

`func (o *CreateQuote201Response) GetMinReceiveAmountOk() (*string, bool)`

GetMinReceiveAmountOk returns a tuple with the MinReceiveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReceiveAmount

`func (o *CreateQuote201Response) SetMinReceiveAmount(v string)`

SetMinReceiveAmount sets MinReceiveAmount field to given value.

### HasMinReceiveAmount

`func (o *CreateQuote201Response) HasMinReceiveAmount() bool`

HasMinReceiveAmount returns a boolean if a field has been set.

### GetMaxReceiveAmount

`func (o *CreateQuote201Response) GetMaxReceiveAmount() string`

GetMaxReceiveAmount returns the MaxReceiveAmount field if non-nil, zero value otherwise.

### GetMaxReceiveAmountOk

`func (o *CreateQuote201Response) GetMaxReceiveAmountOk() (*string, bool)`

GetMaxReceiveAmountOk returns a tuple with the MaxReceiveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReceiveAmount

`func (o *CreateQuote201Response) SetMaxReceiveAmount(v string)`

SetMaxReceiveAmount sets MaxReceiveAmount field to given value.

### HasMaxReceiveAmount

`func (o *CreateQuote201Response) HasMaxReceiveAmount() bool`

HasMaxReceiveAmount returns a boolean if a field has been set.

### GetQuoteCreatedTime

`func (o *CreateQuote201Response) GetQuoteCreatedTime() int32`

GetQuoteCreatedTime returns the QuoteCreatedTime field if non-nil, zero value otherwise.

### GetQuoteCreatedTimeOk

`func (o *CreateQuote201Response) GetQuoteCreatedTimeOk() (*int32, bool)`

GetQuoteCreatedTimeOk returns a tuple with the QuoteCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCreatedTime

`func (o *CreateQuote201Response) SetQuoteCreatedTime(v int32)`

SetQuoteCreatedTime sets QuoteCreatedTime field to given value.


### GetQuoteExpiredTime

`func (o *CreateQuote201Response) GetQuoteExpiredTime() int32`

GetQuoteExpiredTime returns the QuoteExpiredTime field if non-nil, zero value otherwise.

### GetQuoteExpiredTimeOk

`func (o *CreateQuote201Response) GetQuoteExpiredTimeOk() (*int32, bool)`

GetQuoteExpiredTimeOk returns a tuple with the QuoteExpiredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteExpiredTime

`func (o *CreateQuote201Response) SetQuoteExpiredTime(v int32)`

SetQuoteExpiredTime sets QuoteExpiredTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


