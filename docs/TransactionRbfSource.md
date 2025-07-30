# TransactionRbfSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**WalletSubtype**](WalletSubtype.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**Address** | Pointer to **string** | Indicates the wallet address to be used as the source of funds. - For UTXO-based chains: both &#x60;address&#x60; and &#x60;included_utxos&#x60; are optional. If both &#x60;address&#x60; and &#x60;included_utxos&#x60; are provided, the UTXOs must belong to the specified address. If neither &#x60;address&#x60; nor &#x60;included_utxos&#x60; is provided, the system will select UTXOs from the wallet associated with &#x60;wallet_id&#x60;. - For account-based chains: You need to provide &#x60;address&#x60; otherwise the token transfer will fail. However, when estimating fees for a transfer, &#x60;address&#x60; is not required.  For detailed rules on &#x60;address&#x60; and &#x60;included_utxos&#x60; in both regular and RBF transactions, see [Address and included_utxos usage](https://www.cobo.com/developers/v2/guides/transactions/sources-and-destinations#address-and-included-utxos-usage).  | [optional] 
**IncludedUtxos** | Pointer to [**[]TransactionUtxo**](TransactionUtxo.md) |  | [optional] 
**ExcludedUtxos** | Pointer to [**[]TransactionUtxo**](TransactionUtxo.md) |  | [optional] 
**MpcUsedKeyShareHolderGroup** | Pointer to [**MpcSigningGroup**](MpcSigningGroup.md) |  | [optional] 

## Methods

### NewTransactionRbfSource

`func NewTransactionRbfSource(sourceType WalletSubtype, walletId string, ) *TransactionRbfSource`

NewTransactionRbfSource instantiates a new TransactionRbfSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRbfSourceWithDefaults

`func NewTransactionRbfSourceWithDefaults() *TransactionRbfSource`

NewTransactionRbfSourceWithDefaults instantiates a new TransactionRbfSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TransactionRbfSource) GetSourceType() WalletSubtype`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransactionRbfSource) GetSourceTypeOk() (*WalletSubtype, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransactionRbfSource) SetSourceType(v WalletSubtype)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *TransactionRbfSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionRbfSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionRbfSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *TransactionRbfSource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionRbfSource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionRbfSource) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransactionRbfSource) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIncludedUtxos

`func (o *TransactionRbfSource) GetIncludedUtxos() []TransactionUtxo`

GetIncludedUtxos returns the IncludedUtxos field if non-nil, zero value otherwise.

### GetIncludedUtxosOk

`func (o *TransactionRbfSource) GetIncludedUtxosOk() (*[]TransactionUtxo, bool)`

GetIncludedUtxosOk returns a tuple with the IncludedUtxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUtxos

`func (o *TransactionRbfSource) SetIncludedUtxos(v []TransactionUtxo)`

SetIncludedUtxos sets IncludedUtxos field to given value.

### HasIncludedUtxos

`func (o *TransactionRbfSource) HasIncludedUtxos() bool`

HasIncludedUtxos returns a boolean if a field has been set.

### GetExcludedUtxos

`func (o *TransactionRbfSource) GetExcludedUtxos() []TransactionUtxo`

GetExcludedUtxos returns the ExcludedUtxos field if non-nil, zero value otherwise.

### GetExcludedUtxosOk

`func (o *TransactionRbfSource) GetExcludedUtxosOk() (*[]TransactionUtxo, bool)`

GetExcludedUtxosOk returns a tuple with the ExcludedUtxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUtxos

`func (o *TransactionRbfSource) SetExcludedUtxos(v []TransactionUtxo)`

SetExcludedUtxos sets ExcludedUtxos field to given value.

### HasExcludedUtxos

`func (o *TransactionRbfSource) HasExcludedUtxos() bool`

HasExcludedUtxos returns a boolean if a field has been set.

### GetMpcUsedKeyShareHolderGroup

`func (o *TransactionRbfSource) GetMpcUsedKeyShareHolderGroup() MpcSigningGroup`

GetMpcUsedKeyShareHolderGroup returns the MpcUsedKeyShareHolderGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyShareHolderGroupOk

`func (o *TransactionRbfSource) GetMpcUsedKeyShareHolderGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyShareHolderGroupOk returns a tuple with the MpcUsedKeyShareHolderGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyShareHolderGroup

`func (o *TransactionRbfSource) SetMpcUsedKeyShareHolderGroup(v MpcSigningGroup)`

SetMpcUsedKeyShareHolderGroup sets MpcUsedKeyShareHolderGroup field to given value.

### HasMpcUsedKeyShareHolderGroup

`func (o *TransactionRbfSource) HasMpcUsedKeyShareHolderGroup() bool`

HasMpcUsedKeyShareHolderGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


