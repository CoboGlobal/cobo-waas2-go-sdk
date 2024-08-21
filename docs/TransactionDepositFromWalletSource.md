# TransactionDepositFromWalletSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**TransactionSourceType**](TransactionSourceType.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**WalletSubtype** | [**WalletSubtype**](WalletSubtype.md) |  | 
**TradingAccountType** | Pointer to **string** | The exchange trading account or a sub-wallet ID. | [optional] 

## Methods

### NewTransactionDepositFromWalletSource

`func NewTransactionDepositFromWalletSource(sourceType TransactionSourceType, walletId string, walletType WalletType, walletSubtype WalletSubtype, ) *TransactionDepositFromWalletSource`

NewTransactionDepositFromWalletSource instantiates a new TransactionDepositFromWalletSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDepositFromWalletSourceWithDefaults

`func NewTransactionDepositFromWalletSourceWithDefaults() *TransactionDepositFromWalletSource`

NewTransactionDepositFromWalletSourceWithDefaults instantiates a new TransactionDepositFromWalletSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TransactionDepositFromWalletSource) GetSourceType() TransactionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransactionDepositFromWalletSource) GetSourceTypeOk() (*TransactionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransactionDepositFromWalletSource) SetSourceType(v TransactionSourceType)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *TransactionDepositFromWalletSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionDepositFromWalletSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionDepositFromWalletSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetWalletType

`func (o *TransactionDepositFromWalletSource) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *TransactionDepositFromWalletSource) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *TransactionDepositFromWalletSource) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetWalletSubtype

`func (o *TransactionDepositFromWalletSource) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *TransactionDepositFromWalletSource) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *TransactionDepositFromWalletSource) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.


### GetTradingAccountType

`func (o *TransactionDepositFromWalletSource) GetTradingAccountType() string`

GetTradingAccountType returns the TradingAccountType field if non-nil, zero value otherwise.

### GetTradingAccountTypeOk

`func (o *TransactionDepositFromWalletSource) GetTradingAccountTypeOk() (*string, bool)`

GetTradingAccountTypeOk returns a tuple with the TradingAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountType

`func (o *TransactionDepositFromWalletSource) SetTradingAccountType(v string)`

SetTradingAccountType sets TradingAccountType field to given value.

### HasTradingAccountType

`func (o *TransactionDepositFromWalletSource) HasTradingAccountType() bool`

HasTradingAccountType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


