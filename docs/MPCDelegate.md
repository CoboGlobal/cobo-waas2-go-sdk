# MPCDelegate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegateType** | [**CoboSafeDelegateType**](CoboSafeDelegateType.md) |  | [default to COBOSAFEDELEGATETYPE_ORG_CONTROLLED]
**WalletId** | **string** | The wallet ID of the Delegate. This is required when initiating a transfer or contract call from Smart Contract Wallets (Safe{Wallet}). | 
**Address** | **string** | The wallet address of the Delegate. This is required when initiating a transfer or contract call from Smart Contract Wallets (Safe{Wallet}). | 

## Methods

### NewMPCDelegate

`func NewMPCDelegate(delegateType CoboSafeDelegateType, walletId string, address string, ) *MPCDelegate`

NewMPCDelegate instantiates a new MPCDelegate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMPCDelegateWithDefaults

`func NewMPCDelegateWithDefaults() *MPCDelegate`

NewMPCDelegateWithDefaults instantiates a new MPCDelegate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegateType

`func (o *MPCDelegate) GetDelegateType() CoboSafeDelegateType`

GetDelegateType returns the DelegateType field if non-nil, zero value otherwise.

### GetDelegateTypeOk

`func (o *MPCDelegate) GetDelegateTypeOk() (*CoboSafeDelegateType, bool)`

GetDelegateTypeOk returns a tuple with the DelegateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateType

`func (o *MPCDelegate) SetDelegateType(v CoboSafeDelegateType)`

SetDelegateType sets DelegateType field to given value.


### GetWalletId

`func (o *MPCDelegate) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *MPCDelegate) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *MPCDelegate) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *MPCDelegate) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MPCDelegate) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MPCDelegate) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


