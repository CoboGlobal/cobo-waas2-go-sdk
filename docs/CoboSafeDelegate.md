# CoboSafeDelegate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegateType** | [**CoboSafeDelegateType**](CoboSafeDelegateType.md) |  | [default to COBOSAFEDELEGATETYPE_ORG_CONTROLLED]
**WalletId** | **string** | The wallet ID of the Delegate. This is required when initiating a transfer from Smart Contract Wallets (Safe{Wallet}). | 
**Address** | **string** | The wallet address of the Delegate. This is required when initiating a transfer from Smart Contract Wallets (Safe{Wallet}). | 

## Methods

### NewCoboSafeDelegate

`func NewCoboSafeDelegate(delegateType CoboSafeDelegateType, walletId string, address string, ) *CoboSafeDelegate`

NewCoboSafeDelegate instantiates a new CoboSafeDelegate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoboSafeDelegateWithDefaults

`func NewCoboSafeDelegateWithDefaults() *CoboSafeDelegate`

NewCoboSafeDelegateWithDefaults instantiates a new CoboSafeDelegate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegateType

`func (o *CoboSafeDelegate) GetDelegateType() CoboSafeDelegateType`

GetDelegateType returns the DelegateType field if non-nil, zero value otherwise.

### GetDelegateTypeOk

`func (o *CoboSafeDelegate) GetDelegateTypeOk() (*CoboSafeDelegateType, bool)`

GetDelegateTypeOk returns a tuple with the DelegateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateType

`func (o *CoboSafeDelegate) SetDelegateType(v CoboSafeDelegateType)`

SetDelegateType sets DelegateType field to given value.


### GetWalletId

`func (o *CoboSafeDelegate) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CoboSafeDelegate) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CoboSafeDelegate) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *CoboSafeDelegate) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CoboSafeDelegate) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CoboSafeDelegate) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


