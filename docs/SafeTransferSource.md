# SafeTransferSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**WalletSubtype**](WalletSubtype.md) |  | 
**WalletId** | **string** | Unique id of the wallet to transfer from. | 
**AddressStr** | **string** | From address | 
**Delegate** | [**SafeTransferSourceAllOfDelegate**](SafeTransferSourceAllOfDelegate.md) |  | 

## Methods

### NewSafeTransferSource

`func NewSafeTransferSource(sourceType WalletSubtype, walletId string, addressStr string, delegate SafeTransferSourceAllOfDelegate, ) *SafeTransferSource`

NewSafeTransferSource instantiates a new SafeTransferSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeTransferSourceWithDefaults

`func NewSafeTransferSourceWithDefaults() *SafeTransferSource`

NewSafeTransferSourceWithDefaults instantiates a new SafeTransferSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *SafeTransferSource) GetSourceType() WalletSubtype`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *SafeTransferSource) GetSourceTypeOk() (*WalletSubtype, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *SafeTransferSource) SetSourceType(v WalletSubtype)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *SafeTransferSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *SafeTransferSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *SafeTransferSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddressStr

`func (o *SafeTransferSource) GetAddressStr() string`

GetAddressStr returns the AddressStr field if non-nil, zero value otherwise.

### GetAddressStrOk

`func (o *SafeTransferSource) GetAddressStrOk() (*string, bool)`

GetAddressStrOk returns a tuple with the AddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStr

`func (o *SafeTransferSource) SetAddressStr(v string)`

SetAddressStr sets AddressStr field to given value.


### GetDelegate

`func (o *SafeTransferSource) GetDelegate() SafeTransferSourceAllOfDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *SafeTransferSource) GetDelegateOk() (*SafeTransferSourceAllOfDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *SafeTransferSource) SetDelegate(v SafeTransferSourceAllOfDelegate)`

SetDelegate sets Delegate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


