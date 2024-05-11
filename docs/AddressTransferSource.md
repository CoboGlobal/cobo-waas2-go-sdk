# AddressTransferSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**WalletSubtype**](WalletSubtype.md) |  | 
**WalletId** | **string** | Unique id of the wallet to transfer from. | 
**AddressStr** | **string** | From address | 

## Methods

### NewAddressTransferSource

`func NewAddressTransferSource(sourceType WalletSubtype, walletId string, addressStr string, ) *AddressTransferSource`

NewAddressTransferSource instantiates a new AddressTransferSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransferSourceWithDefaults

`func NewAddressTransferSourceWithDefaults() *AddressTransferSource`

NewAddressTransferSourceWithDefaults instantiates a new AddressTransferSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *AddressTransferSource) GetSourceType() WalletSubtype`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *AddressTransferSource) GetSourceTypeOk() (*WalletSubtype, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *AddressTransferSource) SetSourceType(v WalletSubtype)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *AddressTransferSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *AddressTransferSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *AddressTransferSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddressStr

`func (o *AddressTransferSource) GetAddressStr() string`

GetAddressStr returns the AddressStr field if non-nil, zero value otherwise.

### GetAddressStrOk

`func (o *AddressTransferSource) GetAddressStrOk() (*string, bool)`

GetAddressStrOk returns a tuple with the AddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStr

`func (o *AddressTransferSource) SetAddressStr(v string)`

SetAddressStr sets AddressStr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


