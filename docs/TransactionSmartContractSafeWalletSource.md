# TransactionSmartContractSafeWalletSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**TransactionSourceType**](TransactionSourceType.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**Address** | **string** | The wallet address. | 
**Delegate** | [**CoboSafeDelegate**](CoboSafeDelegate.md) |  | 

## Methods

### NewTransactionSmartContractSafeWalletSource

`func NewTransactionSmartContractSafeWalletSource(sourceType TransactionSourceType, walletId string, address string, delegate CoboSafeDelegate, ) *TransactionSmartContractSafeWalletSource`

NewTransactionSmartContractSafeWalletSource instantiates a new TransactionSmartContractSafeWalletSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSmartContractSafeWalletSourceWithDefaults

`func NewTransactionSmartContractSafeWalletSourceWithDefaults() *TransactionSmartContractSafeWalletSource`

NewTransactionSmartContractSafeWalletSourceWithDefaults instantiates a new TransactionSmartContractSafeWalletSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TransactionSmartContractSafeWalletSource) GetSourceType() TransactionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransactionSmartContractSafeWalletSource) GetSourceTypeOk() (*TransactionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransactionSmartContractSafeWalletSource) SetSourceType(v TransactionSourceType)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *TransactionSmartContractSafeWalletSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionSmartContractSafeWalletSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionSmartContractSafeWalletSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *TransactionSmartContractSafeWalletSource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionSmartContractSafeWalletSource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionSmartContractSafeWalletSource) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetDelegate

`func (o *TransactionSmartContractSafeWalletSource) GetDelegate() CoboSafeDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *TransactionSmartContractSafeWalletSource) GetDelegateOk() (*CoboSafeDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *TransactionSmartContractSafeWalletSource) SetDelegate(v CoboSafeDelegate)`

SetDelegate sets Delegate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


