# CreateSwapActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The unique identifier of the wallet. | 
**QuoteId** | **string** | The unique identifier of the quote. | 

## Methods

### NewCreateSwapActivityRequest

`func NewCreateSwapActivityRequest(walletId string, quoteId string, ) *CreateSwapActivityRequest`

NewCreateSwapActivityRequest instantiates a new CreateSwapActivityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSwapActivityRequestWithDefaults

`func NewCreateSwapActivityRequestWithDefaults() *CreateSwapActivityRequest`

NewCreateSwapActivityRequestWithDefaults instantiates a new CreateSwapActivityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *CreateSwapActivityRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CreateSwapActivityRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CreateSwapActivityRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetQuoteId

`func (o *CreateSwapActivityRequest) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *CreateSwapActivityRequest) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *CreateSwapActivityRequest) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


