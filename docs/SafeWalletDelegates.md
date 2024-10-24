# SafeWalletDelegates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestType** | [**EstimateFeeRequestType**](EstimateFeeRequestType.md) |  | 
**Address** | Pointer to **string** | The address of the recipient. | [optional] 
**Value** | Pointer to **string** | The transfer amount. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | [optional] 
**Calldata** | Pointer to **string** | The data that is used to invoke a specific function or method within the specified contract at the destination address.  | [optional] 
**TokenId** | **string** | The token ID. | 
**Amount** | Pointer to **string** | The transfer amount. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;. | [optional] 

## Methods

### NewSafeWalletDelegates

`func NewSafeWalletDelegates(requestType EstimateFeeRequestType, tokenId string, ) *SafeWalletDelegates`

NewSafeWalletDelegates instantiates a new SafeWalletDelegates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeWalletDelegatesWithDefaults

`func NewSafeWalletDelegatesWithDefaults() *SafeWalletDelegates`

NewSafeWalletDelegatesWithDefaults instantiates a new SafeWalletDelegates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestType

`func (o *SafeWalletDelegates) GetRequestType() EstimateFeeRequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *SafeWalletDelegates) GetRequestTypeOk() (*EstimateFeeRequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *SafeWalletDelegates) SetRequestType(v EstimateFeeRequestType)`

SetRequestType sets RequestType field to given value.


### GetAddress

`func (o *SafeWalletDelegates) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SafeWalletDelegates) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SafeWalletDelegates) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SafeWalletDelegates) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetValue

`func (o *SafeWalletDelegates) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SafeWalletDelegates) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SafeWalletDelegates) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SafeWalletDelegates) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCalldata

`func (o *SafeWalletDelegates) GetCalldata() string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *SafeWalletDelegates) GetCalldataOk() (*string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *SafeWalletDelegates) SetCalldata(v string)`

SetCalldata sets Calldata field to given value.

### HasCalldata

`func (o *SafeWalletDelegates) HasCalldata() bool`

HasCalldata returns a boolean if a field has been set.

### GetTokenId

`func (o *SafeWalletDelegates) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *SafeWalletDelegates) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *SafeWalletDelegates) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *SafeWalletDelegates) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SafeWalletDelegates) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SafeWalletDelegates) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SafeWalletDelegates) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


