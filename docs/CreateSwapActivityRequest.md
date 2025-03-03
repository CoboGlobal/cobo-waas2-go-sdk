# CreateSwapActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The unique identifier of the wallet. | 
**QuoteId** | **string** | The unique identifier of the quote. | 
**SlippageTolerance** | **string** | The slippage tolerance for the swap. | 
**AppInitiator** | Pointer to **string** | The initiator of the app activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator. | [optional] 

## Methods

### NewCreateSwapActivityRequest

`func NewCreateSwapActivityRequest(walletId string, quoteId string, slippageTolerance string, ) *CreateSwapActivityRequest`

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


### GetSlippageTolerance

`func (o *CreateSwapActivityRequest) GetSlippageTolerance() string`

GetSlippageTolerance returns the SlippageTolerance field if non-nil, zero value otherwise.

### GetSlippageToleranceOk

`func (o *CreateSwapActivityRequest) GetSlippageToleranceOk() (*string, bool)`

GetSlippageToleranceOk returns a tuple with the SlippageTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippageTolerance

`func (o *CreateSwapActivityRequest) SetSlippageTolerance(v string)`

SetSlippageTolerance sets SlippageTolerance field to given value.


### GetAppInitiator

`func (o *CreateSwapActivityRequest) GetAppInitiator() string`

GetAppInitiator returns the AppInitiator field if non-nil, zero value otherwise.

### GetAppInitiatorOk

`func (o *CreateSwapActivityRequest) GetAppInitiatorOk() (*string, bool)`

GetAppInitiatorOk returns a tuple with the AppInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInitiator

`func (o *CreateSwapActivityRequest) SetAppInitiator(v string)`

SetAppInitiator sets AppInitiator field to given value.

### HasAppInitiator

`func (o *CreateSwapActivityRequest) HasAppInitiator() bool`

HasAppInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


