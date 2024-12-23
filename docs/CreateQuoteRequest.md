# CreateQuoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayTokenId** | **string** | Unique id of the token to pay. | 
**ReceiveTokenId** | **string** | Unique id of the token to receive. | 
**PayAmount** | Pointer to **string** | The amount of token to swap. | [optional] 
**ReceiveAmount** | Pointer to **string** | The amount of token to receive. | [optional] 

## Methods

### NewCreateQuoteRequest

`func NewCreateQuoteRequest(payTokenId string, receiveTokenId string, ) *CreateQuoteRequest`

NewCreateQuoteRequest instantiates a new CreateQuoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuoteRequestWithDefaults

`func NewCreateQuoteRequestWithDefaults() *CreateQuoteRequest`

NewCreateQuoteRequestWithDefaults instantiates a new CreateQuoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayTokenId

`func (o *CreateQuoteRequest) GetPayTokenId() string`

GetPayTokenId returns the PayTokenId field if non-nil, zero value otherwise.

### GetPayTokenIdOk

`func (o *CreateQuoteRequest) GetPayTokenIdOk() (*string, bool)`

GetPayTokenIdOk returns a tuple with the PayTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayTokenId

`func (o *CreateQuoteRequest) SetPayTokenId(v string)`

SetPayTokenId sets PayTokenId field to given value.


### GetReceiveTokenId

`func (o *CreateQuoteRequest) GetReceiveTokenId() string`

GetReceiveTokenId returns the ReceiveTokenId field if non-nil, zero value otherwise.

### GetReceiveTokenIdOk

`func (o *CreateQuoteRequest) GetReceiveTokenIdOk() (*string, bool)`

GetReceiveTokenIdOk returns a tuple with the ReceiveTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveTokenId

`func (o *CreateQuoteRequest) SetReceiveTokenId(v string)`

SetReceiveTokenId sets ReceiveTokenId field to given value.


### GetPayAmount

`func (o *CreateQuoteRequest) GetPayAmount() string`

GetPayAmount returns the PayAmount field if non-nil, zero value otherwise.

### GetPayAmountOk

`func (o *CreateQuoteRequest) GetPayAmountOk() (*string, bool)`

GetPayAmountOk returns a tuple with the PayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayAmount

`func (o *CreateQuoteRequest) SetPayAmount(v string)`

SetPayAmount sets PayAmount field to given value.

### HasPayAmount

`func (o *CreateQuoteRequest) HasPayAmount() bool`

HasPayAmount returns a boolean if a field has been set.

### GetReceiveAmount

`func (o *CreateQuoteRequest) GetReceiveAmount() string`

GetReceiveAmount returns the ReceiveAmount field if non-nil, zero value otherwise.

### GetReceiveAmountOk

`func (o *CreateQuoteRequest) GetReceiveAmountOk() (*string, bool)`

GetReceiveAmountOk returns a tuple with the ReceiveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveAmount

`func (o *CreateQuoteRequest) SetReceiveAmount(v string)`

SetReceiveAmount sets ReceiveAmount field to given value.

### HasReceiveAmount

`func (o *CreateQuoteRequest) HasReceiveAmount() bool`

HasReceiveAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


