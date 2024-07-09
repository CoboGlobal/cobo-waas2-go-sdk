# TransactionSafeWalletSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**TransactionSourceType**](TransactionSourceType.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**Delegate** | [**TransactionSafeWalletSourceDelegate**](TransactionSafeWalletSourceDelegate.md) |  | 

## Methods

### NewTransactionSafeWalletSource

`func NewTransactionSafeWalletSource(sourceType TransactionSourceType, walletId string, delegate TransactionSafeWalletSourceDelegate, ) *TransactionSafeWalletSource`

NewTransactionSafeWalletSource instantiates a new TransactionSafeWalletSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSafeWalletSourceWithDefaults

`func NewTransactionSafeWalletSourceWithDefaults() *TransactionSafeWalletSource`

NewTransactionSafeWalletSourceWithDefaults instantiates a new TransactionSafeWalletSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TransactionSafeWalletSource) GetSourceType() TransactionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransactionSafeWalletSource) GetSourceTypeOk() (*TransactionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransactionSafeWalletSource) SetSourceType(v TransactionSourceType)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *TransactionSafeWalletSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionSafeWalletSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionSafeWalletSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetDelegate

`func (o *TransactionSafeWalletSource) GetDelegate() TransactionSafeWalletSourceDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *TransactionSafeWalletSource) GetDelegateOk() (*TransactionSafeWalletSourceDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *TransactionSafeWalletSource) SetDelegate(v TransactionSafeWalletSourceDelegate)`

SetDelegate sets Delegate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


