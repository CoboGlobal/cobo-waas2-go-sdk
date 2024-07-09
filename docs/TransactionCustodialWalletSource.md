# TransactionCustodialWalletSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**TransactionSourceType**](TransactionSourceType.md) |  | 
**WalletId** | **string** | The wallet ID. | 

## Methods

### NewTransactionCustodialWalletSource

`func NewTransactionCustodialWalletSource(sourceType TransactionSourceType, walletId string, ) *TransactionCustodialWalletSource`

NewTransactionCustodialWalletSource instantiates a new TransactionCustodialWalletSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCustodialWalletSourceWithDefaults

`func NewTransactionCustodialWalletSourceWithDefaults() *TransactionCustodialWalletSource`

NewTransactionCustodialWalletSourceWithDefaults instantiates a new TransactionCustodialWalletSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TransactionCustodialWalletSource) GetSourceType() TransactionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransactionCustodialWalletSource) GetSourceTypeOk() (*TransactionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransactionCustodialWalletSource) SetSourceType(v TransactionSourceType)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *TransactionCustodialWalletSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionCustodialWalletSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionCustodialWalletSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


