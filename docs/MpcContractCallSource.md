# MpcContractCallSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** |  | 
**WalletId** | **string** | Unique id of the wallet to initiate contract call from. | 
**AddressStr** | **string** | From address | 
**MpcUsedKeyGroup** | [**MpcSigningGroup**](MpcSigningGroup.md) |  | 

## Methods

### NewMpcContractCallSource

`func NewMpcContractCallSource(sourceType string, walletId string, addressStr string, mpcUsedKeyGroup MpcSigningGroup, ) *MpcContractCallSource`

NewMpcContractCallSource instantiates a new MpcContractCallSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpcContractCallSourceWithDefaults

`func NewMpcContractCallSourceWithDefaults() *MpcContractCallSource`

NewMpcContractCallSourceWithDefaults instantiates a new MpcContractCallSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *MpcContractCallSource) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *MpcContractCallSource) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *MpcContractCallSource) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *MpcContractCallSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *MpcContractCallSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *MpcContractCallSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddressStr

`func (o *MpcContractCallSource) GetAddressStr() string`

GetAddressStr returns the AddressStr field if non-nil, zero value otherwise.

### GetAddressStrOk

`func (o *MpcContractCallSource) GetAddressStrOk() (*string, bool)`

GetAddressStrOk returns a tuple with the AddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStr

`func (o *MpcContractCallSource) SetAddressStr(v string)`

SetAddressStr sets AddressStr field to given value.


### GetMpcUsedKeyGroup

`func (o *MpcContractCallSource) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *MpcContractCallSource) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *MpcContractCallSource) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


