# TransactionMPCWalletSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**TransactionSourceType**](TransactionSourceType.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**MpcUsedKeyGroup** | [**MpcSigningGroup**](MpcSigningGroup.md) |  | 
**AccountInput** | Pointer to [**TransactionMPCWalletSourceAccountInput**](TransactionMPCWalletSourceAccountInput.md) |  | [optional] 
**UtxoInputs** | Pointer to [**[]TransactionMPCWalletSourceUtxoInputsInner**](TransactionMPCWalletSourceUtxoInputsInner.md) |  | [optional] 

## Methods

### NewTransactionMPCWalletSource

`func NewTransactionMPCWalletSource(sourceType TransactionSourceType, walletId string, mpcUsedKeyGroup MpcSigningGroup, ) *TransactionMPCWalletSource`

NewTransactionMPCWalletSource instantiates a new TransactionMPCWalletSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionMPCWalletSourceWithDefaults

`func NewTransactionMPCWalletSourceWithDefaults() *TransactionMPCWalletSource`

NewTransactionMPCWalletSourceWithDefaults instantiates a new TransactionMPCWalletSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TransactionMPCWalletSource) GetSourceType() TransactionSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransactionMPCWalletSource) GetSourceTypeOk() (*TransactionSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransactionMPCWalletSource) SetSourceType(v TransactionSourceType)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *TransactionMPCWalletSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionMPCWalletSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionMPCWalletSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetMpcUsedKeyGroup

`func (o *TransactionMPCWalletSource) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *TransactionMPCWalletSource) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *TransactionMPCWalletSource) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.


### GetAccountInput

`func (o *TransactionMPCWalletSource) GetAccountInput() TransactionMPCWalletSourceAccountInput`

GetAccountInput returns the AccountInput field if non-nil, zero value otherwise.

### GetAccountInputOk

`func (o *TransactionMPCWalletSource) GetAccountInputOk() (*TransactionMPCWalletSourceAccountInput, bool)`

GetAccountInputOk returns a tuple with the AccountInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInput

`func (o *TransactionMPCWalletSource) SetAccountInput(v TransactionMPCWalletSourceAccountInput)`

SetAccountInput sets AccountInput field to given value.

### HasAccountInput

`func (o *TransactionMPCWalletSource) HasAccountInput() bool`

HasAccountInput returns a boolean if a field has been set.

### GetUtxoInputs

`func (o *TransactionMPCWalletSource) GetUtxoInputs() []TransactionMPCWalletSourceUtxoInputsInner`

GetUtxoInputs returns the UtxoInputs field if non-nil, zero value otherwise.

### GetUtxoInputsOk

`func (o *TransactionMPCWalletSource) GetUtxoInputsOk() (*[]TransactionMPCWalletSourceUtxoInputsInner, bool)`

GetUtxoInputsOk returns a tuple with the UtxoInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoInputs

`func (o *TransactionMPCWalletSource) SetUtxoInputs(v []TransactionMPCWalletSourceUtxoInputsInner)`

SetUtxoInputs sets UtxoInputs field to given value.

### HasUtxoInputs

`func (o *TransactionMPCWalletSource) HasUtxoInputs() bool`

HasUtxoInputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


