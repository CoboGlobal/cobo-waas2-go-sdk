# ForcedSweepRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request id of the force sweep. | 
**WalletId** | **string** | The wallet ID to force sweep, which is the unique identifier of a wallet. | 
**TokenId** | **string** | The token ID to force sweep, which is the unique identifier of a token. | 
**Amount** | **string** | The amount of needing force sweep. | 

## Methods

### NewForcedSweepRequest

`func NewForcedSweepRequest(requestId string, walletId string, tokenId string, amount string, ) *ForcedSweepRequest`

NewForcedSweepRequest instantiates a new ForcedSweepRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForcedSweepRequestWithDefaults

`func NewForcedSweepRequestWithDefaults() *ForcedSweepRequest`

NewForcedSweepRequestWithDefaults instantiates a new ForcedSweepRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ForcedSweepRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ForcedSweepRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ForcedSweepRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetWalletId

`func (o *ForcedSweepRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ForcedSweepRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ForcedSweepRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetTokenId

`func (o *ForcedSweepRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ForcedSweepRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ForcedSweepRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *ForcedSweepRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ForcedSweepRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ForcedSweepRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


