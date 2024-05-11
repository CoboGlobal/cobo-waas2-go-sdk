# MpcTransferSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**WalletSubtype**](WalletSubtype.md) |  | 
**WalletId** | **string** | Unique id of the wallet to transfer from. | 
**AddressStr** | **string** | From address | 
**MpcUsedKeyGroup** | Pointer to [**MpcSigningGroup**](MpcSigningGroup.md) |  | [optional] 

## Methods

### NewMpcTransferSource

`func NewMpcTransferSource(sourceType WalletSubtype, walletId string, addressStr string, ) *MpcTransferSource`

NewMpcTransferSource instantiates a new MpcTransferSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpcTransferSourceWithDefaults

`func NewMpcTransferSourceWithDefaults() *MpcTransferSource`

NewMpcTransferSourceWithDefaults instantiates a new MpcTransferSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *MpcTransferSource) GetSourceType() WalletSubtype`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *MpcTransferSource) GetSourceTypeOk() (*WalletSubtype, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *MpcTransferSource) SetSourceType(v WalletSubtype)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *MpcTransferSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *MpcTransferSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *MpcTransferSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddressStr

`func (o *MpcTransferSource) GetAddressStr() string`

GetAddressStr returns the AddressStr field if non-nil, zero value otherwise.

### GetAddressStrOk

`func (o *MpcTransferSource) GetAddressStrOk() (*string, bool)`

GetAddressStrOk returns a tuple with the AddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStr

`func (o *MpcTransferSource) SetAddressStr(v string)`

SetAddressStr sets AddressStr field to given value.


### GetMpcUsedKeyGroup

`func (o *MpcTransferSource) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *MpcTransferSource) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *MpcTransferSource) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.

### HasMpcUsedKeyGroup

`func (o *MpcTransferSource) HasMpcUsedKeyGroup() bool`

HasMpcUsedKeyGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


