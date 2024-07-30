# TransactionDepositFromAddressSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**TransactionSourceType**](TransactionSourceType.md) |  | 
**WalletId** | Pointer to **string** | The wallet ID. | [optional] 
**WalletType** | Pointer to [**WalletType**](WalletType.md) |  | [optional] 
**WalletSubtype** | Pointer to [**WalletSubtype**](WalletSubtype.md) |  | [optional] 
**Addresses** | **[]string** | A list of addresses. | 

## Methods

### NewTransactionDepositFromAddressSource

`func NewTransactionDepositFromAddressSource(sourceType TransactionSourceType, addresses []string, ) *TransactionDepositFromAddressSource`

NewTransactionDepositFromAddressSource instantiates a new TransactionDepositFromAddressSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDepositFromAddressSourceWithDefaults

`func NewTransactionDepositFromAddressSourceWithDefaults() *TransactionDepositFromAddressSource`

NewTransactionDepositFromAddressSourceWithDefaults instantiates a new TransactionDepositFromAddressSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TransactionDepositFromAddressSource) GetSourceType() TransactionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransactionDepositFromAddressSource) GetSourceTypeOk() (*TransactionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransactionDepositFromAddressSource) SetSourceType(v TransactionSourceType)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *TransactionDepositFromAddressSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionDepositFromAddressSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionDepositFromAddressSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *TransactionDepositFromAddressSource) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetWalletType

`func (o *TransactionDepositFromAddressSource) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *TransactionDepositFromAddressSource) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *TransactionDepositFromAddressSource) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *TransactionDepositFromAddressSource) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.

### GetWalletSubtype

`func (o *TransactionDepositFromAddressSource) GetWalletSubtype() WalletSubtype`

GetWalletSubtype returns the WalletSubtype field if non-nil, zero value otherwise.

### GetWalletSubtypeOk

`func (o *TransactionDepositFromAddressSource) GetWalletSubtypeOk() (*WalletSubtype, bool)`

GetWalletSubtypeOk returns a tuple with the WalletSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSubtype

`func (o *TransactionDepositFromAddressSource) SetWalletSubtype(v WalletSubtype)`

SetWalletSubtype sets WalletSubtype field to given value.

### HasWalletSubtype

`func (o *TransactionDepositFromAddressSource) HasWalletSubtype() bool`

HasWalletSubtype returns a boolean if a field has been set.

### GetAddresses

`func (o *TransactionDepositFromAddressSource) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *TransactionDepositFromAddressSource) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *TransactionDepositFromAddressSource) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


