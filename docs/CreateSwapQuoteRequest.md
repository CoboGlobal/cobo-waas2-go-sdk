# CreateSwapQuoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The unique identifier of the wallet. | 
**PayTokenId** | **string** | Unique id of the token to pay. | 
**ReceiveTokenId** | **string** | Unique id of the token to receive. | 
**PayAmount** | Pointer to **string** | Amount of tokens to pay. For example \&quot;0.5 BTC\&quot;. Note: Either pay_amount or receive_amount must be provided, but not both.  | [optional] 
**ReceiveAmount** | Pointer to **string** | Amount of tokens to receive. For example \&quot;0.5 ETH_WBTC\&quot;. Note: Either pay_amount or receive_amount must be provided, but not both.  | [optional] 

## Methods

### NewCreateSwapQuoteRequest

`func NewCreateSwapQuoteRequest(walletId string, payTokenId string, receiveTokenId string, ) *CreateSwapQuoteRequest`

NewCreateSwapQuoteRequest instantiates a new CreateSwapQuoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSwapQuoteRequestWithDefaults

`func NewCreateSwapQuoteRequestWithDefaults() *CreateSwapQuoteRequest`

NewCreateSwapQuoteRequestWithDefaults instantiates a new CreateSwapQuoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *CreateSwapQuoteRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CreateSwapQuoteRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CreateSwapQuoteRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetPayTokenId

`func (o *CreateSwapQuoteRequest) GetPayTokenId() string`

GetPayTokenId returns the PayTokenId field if non-nil, zero value otherwise.

### GetPayTokenIdOk

`func (o *CreateSwapQuoteRequest) GetPayTokenIdOk() (*string, bool)`

GetPayTokenIdOk returns a tuple with the PayTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayTokenId

`func (o *CreateSwapQuoteRequest) SetPayTokenId(v string)`

SetPayTokenId sets PayTokenId field to given value.


### GetReceiveTokenId

`func (o *CreateSwapQuoteRequest) GetReceiveTokenId() string`

GetReceiveTokenId returns the ReceiveTokenId field if non-nil, zero value otherwise.

### GetReceiveTokenIdOk

`func (o *CreateSwapQuoteRequest) GetReceiveTokenIdOk() (*string, bool)`

GetReceiveTokenIdOk returns a tuple with the ReceiveTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveTokenId

`func (o *CreateSwapQuoteRequest) SetReceiveTokenId(v string)`

SetReceiveTokenId sets ReceiveTokenId field to given value.


### GetPayAmount

`func (o *CreateSwapQuoteRequest) GetPayAmount() string`

GetPayAmount returns the PayAmount field if non-nil, zero value otherwise.

### GetPayAmountOk

`func (o *CreateSwapQuoteRequest) GetPayAmountOk() (*string, bool)`

GetPayAmountOk returns a tuple with the PayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayAmount

`func (o *CreateSwapQuoteRequest) SetPayAmount(v string)`

SetPayAmount sets PayAmount field to given value.

### HasPayAmount

`func (o *CreateSwapQuoteRequest) HasPayAmount() bool`

HasPayAmount returns a boolean if a field has been set.

### GetReceiveAmount

`func (o *CreateSwapQuoteRequest) GetReceiveAmount() string`

GetReceiveAmount returns the ReceiveAmount field if non-nil, zero value otherwise.

### GetReceiveAmountOk

`func (o *CreateSwapQuoteRequest) GetReceiveAmountOk() (*string, bool)`

GetReceiveAmountOk returns a tuple with the ReceiveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveAmount

`func (o *CreateSwapQuoteRequest) SetReceiveAmount(v string)`

SetReceiveAmount sets ReceiveAmount field to given value.

### HasReceiveAmount

`func (o *CreateSwapQuoteRequest) HasReceiveAmount() bool`

HasReceiveAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


