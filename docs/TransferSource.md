# TransferSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**WalletSubtype**](WalletSubtype.md) |  | 
**WalletId** | **string** | Unique id of the wallet to transfer from. | 
**MpcUsedKeyGroup** | Pointer to [**MpcSigningGroup**](MpcSigningGroup.md) |  | [optional] 
**AddressStr** | **string** | From address | 
**Delegate** | [**SafeTransferSourceAllOfDelegate**](SafeTransferSourceAllOfDelegate.md) |  | 
**SubWalletId** | **string** | Exchange trading account or any sub wallet info for transfer. | 

## Methods

### NewTransferSource

`func NewTransferSource(sourceType WalletSubtype, walletId string, addressStr string, delegate SafeTransferSourceAllOfDelegate, subWalletId string, ) *TransferSource`

NewTransferSource instantiates a new TransferSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferSourceWithDefaults

`func NewTransferSourceWithDefaults() *TransferSource`

NewTransferSourceWithDefaults instantiates a new TransferSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TransferSource) GetSourceType() WalletSubtype`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransferSource) GetSourceTypeOk() (*WalletSubtype, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransferSource) SetSourceType(v WalletSubtype)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *TransferSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransferSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransferSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetMpcUsedKeyGroup

`func (o *TransferSource) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *TransferSource) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *TransferSource) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.

### HasMpcUsedKeyGroup

`func (o *TransferSource) HasMpcUsedKeyGroup() bool`

HasMpcUsedKeyGroup returns a boolean if a field has been set.

### GetAddressStr

`func (o *TransferSource) GetAddressStr() string`

GetAddressStr returns the AddressStr field if non-nil, zero value otherwise.

### GetAddressStrOk

`func (o *TransferSource) GetAddressStrOk() (*string, bool)`

GetAddressStrOk returns a tuple with the AddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStr

`func (o *TransferSource) SetAddressStr(v string)`

SetAddressStr sets AddressStr field to given value.


### GetDelegate

`func (o *TransferSource) GetDelegate() SafeTransferSourceAllOfDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *TransferSource) GetDelegateOk() (*SafeTransferSourceAllOfDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *TransferSource) SetDelegate(v SafeTransferSourceAllOfDelegate)`

SetDelegate sets Delegate field to given value.


### GetSubWalletId

`func (o *TransferSource) GetSubWalletId() string`

GetSubWalletId returns the SubWalletId field if non-nil, zero value otherwise.

### GetSubWalletIdOk

`func (o *TransferSource) GetSubWalletIdOk() (*string, bool)`

GetSubWalletIdOk returns a tuple with the SubWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWalletId

`func (o *TransferSource) SetSubWalletId(v string)`

SetSubWalletId sets SubWalletId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


