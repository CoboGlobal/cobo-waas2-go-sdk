# SafeWalletDelegatesContractCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestType** | [**EstimateFeeRequestType**](EstimateFeeRequestType.md) |  | 
**Address** | Pointer to **string** | The destination address. | [optional] 
**Value** | Pointer to **string** | The transfer amount. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | [optional] 
**Calldata** | Pointer to **string** | The data used to invoke a specific function or method within the specified contract at the destination address, with a maximum length of 65,000 characters.  | [optional] 

## Methods

### NewSafeWalletDelegatesContractCall

`func NewSafeWalletDelegatesContractCall(requestType EstimateFeeRequestType, ) *SafeWalletDelegatesContractCall`

NewSafeWalletDelegatesContractCall instantiates a new SafeWalletDelegatesContractCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeWalletDelegatesContractCallWithDefaults

`func NewSafeWalletDelegatesContractCallWithDefaults() *SafeWalletDelegatesContractCall`

NewSafeWalletDelegatesContractCallWithDefaults instantiates a new SafeWalletDelegatesContractCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestType

`func (o *SafeWalletDelegatesContractCall) GetRequestType() EstimateFeeRequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *SafeWalletDelegatesContractCall) GetRequestTypeOk() (*EstimateFeeRequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *SafeWalletDelegatesContractCall) SetRequestType(v EstimateFeeRequestType)`

SetRequestType sets RequestType field to given value.


### GetAddress

`func (o *SafeWalletDelegatesContractCall) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SafeWalletDelegatesContractCall) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SafeWalletDelegatesContractCall) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SafeWalletDelegatesContractCall) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetValue

`func (o *SafeWalletDelegatesContractCall) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SafeWalletDelegatesContractCall) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SafeWalletDelegatesContractCall) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SafeWalletDelegatesContractCall) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCalldata

`func (o *SafeWalletDelegatesContractCall) GetCalldata() string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *SafeWalletDelegatesContractCall) GetCalldataOk() (*string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *SafeWalletDelegatesContractCall) SetCalldata(v string)`

SetCalldata sets Calldata field to given value.

### HasCalldata

`func (o *SafeWalletDelegatesContractCall) HasCalldata() bool`

HasCalldata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


