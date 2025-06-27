# CreateSwapActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The unique identifier of the wallet to pay. | 
**Address** | Pointer to **string** | The wallet address. | [optional] 
**QuoteId** | **string** | The unique identifier of the quote. | 
**AppInitiator** | Pointer to **string** | The initiator of the app activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator. | [optional] 
**RequestId** | Pointer to **string** | The request id of the swap activity. | [optional] 
**Destination** | Pointer to [**AddressTransferDestination**](AddressTransferDestination.md) |  | [optional] 
**Fee** | Pointer to [**EstimatedFee**](EstimatedFee.md) |  | [optional] 

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


### GetAddress

`func (o *CreateSwapActivityRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateSwapActivityRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateSwapActivityRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CreateSwapActivityRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

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

### GetRequestId

`func (o *CreateSwapActivityRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateSwapActivityRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateSwapActivityRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateSwapActivityRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetDestination

`func (o *CreateSwapActivityRequest) GetDestination() AddressTransferDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CreateSwapActivityRequest) GetDestinationOk() (*AddressTransferDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CreateSwapActivityRequest) SetDestination(v AddressTransferDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *CreateSwapActivityRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetFee

`func (o *CreateSwapActivityRequest) GetFee() EstimatedFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateSwapActivityRequest) GetFeeOk() (*EstimatedFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateSwapActivityRequest) SetFee(v EstimatedFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *CreateSwapActivityRequest) HasFee() bool`

HasFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


