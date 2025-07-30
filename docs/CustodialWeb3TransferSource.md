# CustodialWeb3TransferSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**WalletSubtype**](WalletSubtype.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**Address** | Pointer to **string** | Indicates the wallet address to be used as the source of funds. - For UTXO-based chains: both &#x60;address&#x60; and &#x60;included_utxos&#x60; are optional. If both &#x60;address&#x60; and &#x60;included_utxos&#x60; are provided, the UTXOs must belong to the specified address. If neither &#x60;address&#x60; nor &#x60;included_utxos&#x60; is provided, the system will select UTXOs from the wallet associated with &#x60;wallet_id&#x60;. - For account-based chains: You need to provide &#x60;address&#x60; otherwise the token transfer will fail. However, when estimating fees for a transfer, &#x60;address&#x60; is not required.  For detailed rules on &#x60;address&#x60; and &#x60;included_utxos&#x60; in both regular and RBF transactions, see [Address and included_utxos usage](https://www.cobo.com/developers/v2/guides/transactions/sources-and-destinations#address-and-included-utxos-usage).  | [optional] 
**IncludedUtxos** | Pointer to [**[]TransactionUtxo**](TransactionUtxo.md) |  | [optional] 
**ExcludedUtxos** | Pointer to [**[]TransactionUtxo**](TransactionUtxo.md) |  | [optional] 

## Methods

### NewCustodialWeb3TransferSource

`func NewCustodialWeb3TransferSource(sourceType WalletSubtype, walletId string, ) *CustodialWeb3TransferSource`

NewCustodialWeb3TransferSource instantiates a new CustodialWeb3TransferSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustodialWeb3TransferSourceWithDefaults

`func NewCustodialWeb3TransferSourceWithDefaults() *CustodialWeb3TransferSource`

NewCustodialWeb3TransferSourceWithDefaults instantiates a new CustodialWeb3TransferSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *CustodialWeb3TransferSource) GetSourceType() WalletSubtype`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *CustodialWeb3TransferSource) GetSourceTypeOk() (*WalletSubtype, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *CustodialWeb3TransferSource) SetSourceType(v WalletSubtype)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *CustodialWeb3TransferSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CustodialWeb3TransferSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CustodialWeb3TransferSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *CustodialWeb3TransferSource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustodialWeb3TransferSource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustodialWeb3TransferSource) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CustodialWeb3TransferSource) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIncludedUtxos

`func (o *CustodialWeb3TransferSource) GetIncludedUtxos() []TransactionUtxo`

GetIncludedUtxos returns the IncludedUtxos field if non-nil, zero value otherwise.

### GetIncludedUtxosOk

`func (o *CustodialWeb3TransferSource) GetIncludedUtxosOk() (*[]TransactionUtxo, bool)`

GetIncludedUtxosOk returns a tuple with the IncludedUtxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUtxos

`func (o *CustodialWeb3TransferSource) SetIncludedUtxos(v []TransactionUtxo)`

SetIncludedUtxos sets IncludedUtxos field to given value.

### HasIncludedUtxos

`func (o *CustodialWeb3TransferSource) HasIncludedUtxos() bool`

HasIncludedUtxos returns a boolean if a field has been set.

### GetExcludedUtxos

`func (o *CustodialWeb3TransferSource) GetExcludedUtxos() []TransactionUtxo`

GetExcludedUtxos returns the ExcludedUtxos field if non-nil, zero value otherwise.

### GetExcludedUtxosOk

`func (o *CustodialWeb3TransferSource) GetExcludedUtxosOk() (*[]TransactionUtxo, bool)`

GetExcludedUtxosOk returns a tuple with the ExcludedUtxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUtxos

`func (o *CustodialWeb3TransferSource) SetExcludedUtxos(v []TransactionUtxo)`

SetExcludedUtxos sets ExcludedUtxos field to given value.

### HasExcludedUtxos

`func (o *CustodialWeb3TransferSource) HasExcludedUtxos() bool`

HasExcludedUtxos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


