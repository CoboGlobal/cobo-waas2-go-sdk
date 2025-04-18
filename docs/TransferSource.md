# TransferSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**WalletSubtype**](WalletSubtype.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**Address** | **string** | The wallet address. | 
**IncludedUtxos** | Pointer to [**[]TransactionUtxo**](TransactionUtxo.md) |  | [optional] 
**ExcludedUtxos** | Pointer to [**[]TransactionUtxo**](TransactionUtxo.md) |  | [optional] 
**MpcUsedKeyShareHolderGroup** | Pointer to [**MpcSigningGroup**](MpcSigningGroup.md) |  | [optional] 
**Delegate** | [**CoboSafeDelegate**](CoboSafeDelegate.md) |  | 
**TradingAccountType** | **string** | The trading account type. | 

## Methods

### NewTransferSource

`func NewTransferSource(sourceType WalletSubtype, walletId string, address string, delegate CoboSafeDelegate, tradingAccountType string, ) *TransferSource`

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


### GetAddress

`func (o *TransferSource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransferSource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransferSource) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetIncludedUtxos

`func (o *TransferSource) GetIncludedUtxos() []TransactionUtxo`

GetIncludedUtxos returns the IncludedUtxos field if non-nil, zero value otherwise.

### GetIncludedUtxosOk

`func (o *TransferSource) GetIncludedUtxosOk() (*[]TransactionUtxo, bool)`

GetIncludedUtxosOk returns a tuple with the IncludedUtxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUtxos

`func (o *TransferSource) SetIncludedUtxos(v []TransactionUtxo)`

SetIncludedUtxos sets IncludedUtxos field to given value.

### HasIncludedUtxos

`func (o *TransferSource) HasIncludedUtxos() bool`

HasIncludedUtxos returns a boolean if a field has been set.

### GetExcludedUtxos

`func (o *TransferSource) GetExcludedUtxos() []TransactionUtxo`

GetExcludedUtxos returns the ExcludedUtxos field if non-nil, zero value otherwise.

### GetExcludedUtxosOk

`func (o *TransferSource) GetExcludedUtxosOk() (*[]TransactionUtxo, bool)`

GetExcludedUtxosOk returns a tuple with the ExcludedUtxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUtxos

`func (o *TransferSource) SetExcludedUtxos(v []TransactionUtxo)`

SetExcludedUtxos sets ExcludedUtxos field to given value.

### HasExcludedUtxos

`func (o *TransferSource) HasExcludedUtxos() bool`

HasExcludedUtxos returns a boolean if a field has been set.

### GetMpcUsedKeyShareHolderGroup

`func (o *TransferSource) GetMpcUsedKeyShareHolderGroup() MpcSigningGroup`

GetMpcUsedKeyShareHolderGroup returns the MpcUsedKeyShareHolderGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyShareHolderGroupOk

`func (o *TransferSource) GetMpcUsedKeyShareHolderGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyShareHolderGroupOk returns a tuple with the MpcUsedKeyShareHolderGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyShareHolderGroup

`func (o *TransferSource) SetMpcUsedKeyShareHolderGroup(v MpcSigningGroup)`

SetMpcUsedKeyShareHolderGroup sets MpcUsedKeyShareHolderGroup field to given value.

### HasMpcUsedKeyShareHolderGroup

`func (o *TransferSource) HasMpcUsedKeyShareHolderGroup() bool`

HasMpcUsedKeyShareHolderGroup returns a boolean if a field has been set.

### GetDelegate

`func (o *TransferSource) GetDelegate() CoboSafeDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *TransferSource) GetDelegateOk() (*CoboSafeDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *TransferSource) SetDelegate(v CoboSafeDelegate)`

SetDelegate sets Delegate field to given value.


### GetTradingAccountType

`func (o *TransferSource) GetTradingAccountType() string`

GetTradingAccountType returns the TradingAccountType field if non-nil, zero value otherwise.

### GetTradingAccountTypeOk

`func (o *TransferSource) GetTradingAccountTypeOk() (*string, bool)`

GetTradingAccountTypeOk returns a tuple with the TradingAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountType

`func (o *TransferSource) SetTradingAccountType(v string)`

SetTradingAccountType sets TradingAccountType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


