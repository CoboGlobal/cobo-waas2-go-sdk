# MpcMessageSignSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**MessageSignSourceType**](MessageSignSourceType.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**Address** | **string** | The wallet address. | 
**MpcUsedKeyShareHolderGroup** | Pointer to [**MpcSigningGroup**](MpcSigningGroup.md) |  | [optional] 

## Methods

### NewMpcMessageSignSource

`func NewMpcMessageSignSource(sourceType MessageSignSourceType, walletId string, address string, ) *MpcMessageSignSource`

NewMpcMessageSignSource instantiates a new MpcMessageSignSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpcMessageSignSourceWithDefaults

`func NewMpcMessageSignSourceWithDefaults() *MpcMessageSignSource`

NewMpcMessageSignSourceWithDefaults instantiates a new MpcMessageSignSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *MpcMessageSignSource) GetSourceType() MessageSignSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *MpcMessageSignSource) GetSourceTypeOk() (*MessageSignSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *MpcMessageSignSource) SetSourceType(v MessageSignSourceType)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *MpcMessageSignSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *MpcMessageSignSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *MpcMessageSignSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *MpcMessageSignSource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MpcMessageSignSource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MpcMessageSignSource) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMpcUsedKeyShareHolderGroup

`func (o *MpcMessageSignSource) GetMpcUsedKeyShareHolderGroup() MpcSigningGroup`

GetMpcUsedKeyShareHolderGroup returns the MpcUsedKeyShareHolderGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyShareHolderGroupOk

`func (o *MpcMessageSignSource) GetMpcUsedKeyShareHolderGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyShareHolderGroupOk returns a tuple with the MpcUsedKeyShareHolderGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyShareHolderGroup

`func (o *MpcMessageSignSource) SetMpcUsedKeyShareHolderGroup(v MpcSigningGroup)`

SetMpcUsedKeyShareHolderGroup sets MpcUsedKeyShareHolderGroup field to given value.

### HasMpcUsedKeyShareHolderGroup

`func (o *MpcMessageSignSource) HasMpcUsedKeyShareHolderGroup() bool`

HasMpcUsedKeyShareHolderGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


