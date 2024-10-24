# TransactionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**TransactionSourceType**](TransactionSourceType.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**Address** | **string** | The wallet address. | 
**IncludedUtxos** | Pointer to [**[]TransactionUtxo**](TransactionUtxo.md) |  | [optional] 
**ExcludedUtxos** | Pointer to [**[]TransactionUtxo**](TransactionUtxo.md) |  | [optional] 
**SignerKeyShareHolderGroupId** | Pointer to **string** | The ID of the key share holder group that is selected to sign the transaction. | [optional] 
**Delegate** | Pointer to [**CoboSafeDelegate**](CoboSafeDelegate.md) |  | [optional] 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**TradingAccountType** | Pointer to **string** | The exchange trading account or a sub-wallet ID. | [optional] 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**Addresses** | **[]string** | A list of addresses. | 

## Methods

### NewTransactionSource

`func NewTransactionSource(sourceType TransactionSourceType, walletId string, address string, exchangeId ExchangeId, walletType WalletType, walletSubtype WalletSubtype, addresses []string, ) *TransactionSource`

NewTransactionSource instantiates a new TransactionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSourceWithDefaults

`func NewTransactionSourceWithDefaults() *TransactionSource`

NewTransactionSourceWithDefaults instantiates a new TransactionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TransactionSource) GetSourceType() TransactionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransactionSource) GetSourceTypeOk() (*TransactionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransactionSource) SetSourceType(v TransactionSourceType)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *TransactionSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *TransactionSource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionSource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionSource) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetIncludedUtxos

`func (o *TransactionSource) GetIncludedUtxos() []TransactionUtxo`

GetIncludedUtxos returns the IncludedUtxos field if non-nil, zero value otherwise.

### GetIncludedUtxosOk

`func (o *TransactionSource) GetIncludedUtxosOk() (*[]TransactionUtxo, bool)`

GetIncludedUtxosOk returns a tuple with the IncludedUtxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUtxos

`func (o *TransactionSource) SetIncludedUtxos(v []TransactionUtxo)`

SetIncludedUtxos sets IncludedUtxos field to given value.

### HasIncludedUtxos

`func (o *TransactionSource) HasIncludedUtxos() bool`

HasIncludedUtxos returns a boolean if a field has been set.

### GetExcludedUtxos

`func (o *TransactionSource) GetExcludedUtxos() []TransactionUtxo`

GetExcludedUtxos returns the ExcludedUtxos field if non-nil, zero value otherwise.

### GetExcludedUtxosOk

`func (o *TransactionSource) GetExcludedUtxosOk() (*[]TransactionUtxo, bool)`

GetExcludedUtxosOk returns a tuple with the ExcludedUtxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUtxos

`func (o *TransactionSource) SetExcludedUtxos(v []TransactionUtxo)`

SetExcludedUtxos sets ExcludedUtxos field to given value.

### HasExcludedUtxos

`func (o *TransactionSource) HasExcludedUtxos() bool`

HasExcludedUtxos returns a boolean if a field has been set.

### GetSignerKeyShareHolderGroupId

`func (o *TransactionSource) GetSignerKeyShareHolderGroupId() string`

GetSignerKeyShareHolderGroupId returns the SignerKeyShareHolderGroupId field if non-nil, zero value otherwise.

### GetSignerKeyShareHolderGroupIdOk

`func (o *TransactionSource) GetSignerKeyShareHolderGroupIdOk() (*string, bool)`

GetSignerKeyShareHolderGroupIdOk returns a tuple with the SignerKeyShareHolderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerKeyShareHolderGroupId

`func (o *TransactionSource) SetSignerKeyShareHolderGroupId(v string)`

SetSignerKeyShareHolderGroupId sets SignerKeyShareHolderGroupId field to given value.

### HasSignerKeyShareHolderGroupId

`func (o *TransactionSource) HasSignerKeyShareHolderGroupId() bool`

HasSignerKeyShareHolderGroupId returns a boolean if a field has been set.

### GetDelegate

`func (o *TransactionSource) GetDelegate() CoboSafeDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *TransactionSource) GetDelegateOk() (*CoboSafeDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *TransactionSource) SetDelegate(v CoboSafeDelegate)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *TransactionSource) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.

### GetExchangeId

`func (o *TransactionSource) GetExchangeId() ExchangeId`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *TransactionSource) GetExchangeIdOk() (*ExchangeId, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *TransactionSource) SetExchangeId(v ExchangeId)`

SetExchangeId sets ExchangeId field to given value.


### GetTradingAccountType

`func (o *TransactionSource) GetTradingAccountType() string`

GetTradingAccountType returns the TradingAccountType field if non-nil, zero value otherwise.

### GetTradingAccountTypeOk

`func (o *TransactionSource) GetTradingAccountTypeOk() (*string, bool)`

GetTradingAccountTypeOk returns a tuple with the TradingAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountType

`func (o *TransactionSource) SetTradingAccountType(v string)`

SetTradingAccountType sets TradingAccountType field to given value.

### HasTradingAccountType

`func (o *TransactionSource) HasTradingAccountType() bool`

HasTradingAccountType returns a boolean if a field has been set.

### GetWalletType

`func (o *TransactionSource) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *TransactionSource) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *TransactionSource) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *TransactionSource) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *TransactionSource) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *TransactionSource) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetAddresses

`func (o *TransactionSource) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *TransactionSource) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *TransactionSource) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


