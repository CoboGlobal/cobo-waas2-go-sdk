# MpcTransferSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**WalletSubtype**](WalletSubtype.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**Address** | Pointer to **string** | The wallet address. If you want to specify the UTXOs to be used, please provide the &#x60;included_utxos&#x60; property. If you specify both the &#x60;address&#x60; and &#x60;included_utxos&#x60; properties, the specified included UTXOs must belong to the address. It is recommended to specify no more than 100 included UTXOs to ensure optimal transaction processing.  You need to provide either the &#x60;address&#x60; or &#x60;included_utxos&#x60; property. If neither property is provided, the transfer will fail.  | [optional] 
**IncludedUtxos** | Pointer to [**[]TransactionUtxo**](TransactionUtxo.md) |  | [optional] 
**ExcludedUtxos** | Pointer to [**[]TransactionUtxo**](TransactionUtxo.md) |  | [optional] 
**MpcUsedKeyShareHolderGroup** | Pointer to [**MpcSigningGroup**](MpcSigningGroup.md) |  | [optional] 

## Methods

### NewMpcTransferSource

`func NewMpcTransferSource(sourceType WalletSubtype, walletId string, ) *MpcTransferSource`

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


### GetAddress

`func (o *MpcTransferSource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MpcTransferSource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MpcTransferSource) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MpcTransferSource) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIncludedUtxos

`func (o *MpcTransferSource) GetIncludedUtxos() []TransactionUtxo`

GetIncludedUtxos returns the IncludedUtxos field if non-nil, zero value otherwise.

### GetIncludedUtxosOk

`func (o *MpcTransferSource) GetIncludedUtxosOk() (*[]TransactionUtxo, bool)`

GetIncludedUtxosOk returns a tuple with the IncludedUtxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUtxos

`func (o *MpcTransferSource) SetIncludedUtxos(v []TransactionUtxo)`

SetIncludedUtxos sets IncludedUtxos field to given value.

### HasIncludedUtxos

`func (o *MpcTransferSource) HasIncludedUtxos() bool`

HasIncludedUtxos returns a boolean if a field has been set.

### GetExcludedUtxos

`func (o *MpcTransferSource) GetExcludedUtxos() []TransactionUtxo`

GetExcludedUtxos returns the ExcludedUtxos field if non-nil, zero value otherwise.

### GetExcludedUtxosOk

`func (o *MpcTransferSource) GetExcludedUtxosOk() (*[]TransactionUtxo, bool)`

GetExcludedUtxosOk returns a tuple with the ExcludedUtxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUtxos

`func (o *MpcTransferSource) SetExcludedUtxos(v []TransactionUtxo)`

SetExcludedUtxos sets ExcludedUtxos field to given value.

### HasExcludedUtxos

`func (o *MpcTransferSource) HasExcludedUtxos() bool`

HasExcludedUtxos returns a boolean if a field has been set.

### GetMpcUsedKeyShareHolderGroup

`func (o *MpcTransferSource) GetMpcUsedKeyShareHolderGroup() MpcSigningGroup`

GetMpcUsedKeyShareHolderGroup returns the MpcUsedKeyShareHolderGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyShareHolderGroupOk

`func (o *MpcTransferSource) GetMpcUsedKeyShareHolderGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyShareHolderGroupOk returns a tuple with the MpcUsedKeyShareHolderGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyShareHolderGroup

`func (o *MpcTransferSource) SetMpcUsedKeyShareHolderGroup(v MpcSigningGroup)`

SetMpcUsedKeyShareHolderGroup sets MpcUsedKeyShareHolderGroup field to given value.

### HasMpcUsedKeyShareHolderGroup

`func (o *MpcTransferSource) HasMpcUsedKeyShareHolderGroup() bool`

HasMpcUsedKeyShareHolderGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


