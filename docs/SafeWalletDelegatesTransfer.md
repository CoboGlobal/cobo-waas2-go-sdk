# SafeWalletDelegatesTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestType** | [**EstimateFeeRequestType**](EstimateFeeRequestType.md) |  | 
**TokenId** | **string** | The token ID. | 
**Amount** | Pointer to **string** | The transfer amount. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;. | [optional] 
**Address** | Pointer to **string** | The address of the recipient. | [optional] 

## Methods

### NewSafeWalletDelegatesTransfer

`func NewSafeWalletDelegatesTransfer(requestType EstimateFeeRequestType, tokenId string, ) *SafeWalletDelegatesTransfer`

NewSafeWalletDelegatesTransfer instantiates a new SafeWalletDelegatesTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeWalletDelegatesTransferWithDefaults

`func NewSafeWalletDelegatesTransferWithDefaults() *SafeWalletDelegatesTransfer`

NewSafeWalletDelegatesTransferWithDefaults instantiates a new SafeWalletDelegatesTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestType

`func (o *SafeWalletDelegatesTransfer) GetRequestType() EstimateFeeRequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *SafeWalletDelegatesTransfer) GetRequestTypeOk() (*EstimateFeeRequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *SafeWalletDelegatesTransfer) SetRequestType(v EstimateFeeRequestType)`

SetRequestType sets RequestType field to given value.


### GetTokenId

`func (o *SafeWalletDelegatesTransfer) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *SafeWalletDelegatesTransfer) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *SafeWalletDelegatesTransfer) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *SafeWalletDelegatesTransfer) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SafeWalletDelegatesTransfer) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SafeWalletDelegatesTransfer) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SafeWalletDelegatesTransfer) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAddress

`func (o *SafeWalletDelegatesTransfer) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SafeWalletDelegatesTransfer) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SafeWalletDelegatesTransfer) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SafeWalletDelegatesTransfer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


