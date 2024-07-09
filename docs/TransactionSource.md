# TransactionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**TransactionSourceType**](TransactionSourceType.md) |  | 
**AccountInput** | Pointer to [**TransactionMPCWalletSourceAccountInput**](TransactionMPCWalletSourceAccountInput.md) |  | [optional] 
**UtxoInputs** | Pointer to [**[]TransactionMPCWalletSourceUtxoInputsInner**](TransactionMPCWalletSourceUtxoInputsInner.md) |  | [optional] 
**WalletId** | **string** | The Wallet ID. | 
**MpcUsedKeyGroup** | [**MpcSigningGroup**](MpcSigningGroup.md) |  | 
**Delegate** | [**TransactionSafeWalletSourceDelegate**](TransactionSafeWalletSourceDelegate.md) |  | 
**ExchangeId** | [**ExchangeId**](ExchangeId.md) |  | 
**SubWalletId** | Pointer to **string** | The exchange trading account or a sub-wallet ID. | [optional] 

## Methods

### NewTransactionSource

`func NewTransactionSource(sourceType TransactionSourceType, walletId string, mpcUsedKeyGroup MpcSigningGroup, delegate TransactionSafeWalletSourceDelegate, exchangeId ExchangeId, ) *TransactionSource`

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


### GetAccountInput

`func (o *TransactionSource) GetAccountInput() TransactionMPCWalletSourceAccountInput`

GetAccountInput returns the AccountInput field if non-nil, zero value otherwise.

### GetAccountInputOk

`func (o *TransactionSource) GetAccountInputOk() (*TransactionMPCWalletSourceAccountInput, bool)`

GetAccountInputOk returns a tuple with the AccountInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInput

`func (o *TransactionSource) SetAccountInput(v TransactionMPCWalletSourceAccountInput)`

SetAccountInput sets AccountInput field to given value.

### HasAccountInput

`func (o *TransactionSource) HasAccountInput() bool`

HasAccountInput returns a boolean if a field has been set.

### GetUtxoInputs

`func (o *TransactionSource) GetUtxoInputs() []TransactionMPCWalletSourceUtxoInputsInner`

GetUtxoInputs returns the UtxoInputs field if non-nil, zero value otherwise.

### GetUtxoInputsOk

`func (o *TransactionSource) GetUtxoInputsOk() (*[]TransactionMPCWalletSourceUtxoInputsInner, bool)`

GetUtxoInputsOk returns a tuple with the UtxoInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoInputs

`func (o *TransactionSource) SetUtxoInputs(v []TransactionMPCWalletSourceUtxoInputsInner)`

SetUtxoInputs sets UtxoInputs field to given value.

### HasUtxoInputs

`func (o *TransactionSource) HasUtxoInputs() bool`

HasUtxoInputs returns a boolean if a field has been set.

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


### GetMpcUsedKeyGroup

`func (o *TransactionSource) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *TransactionSource) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *TransactionSource) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.


### GetDelegate

`func (o *TransactionSource) GetDelegate() TransactionSafeWalletSourceDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *TransactionSource) GetDelegateOk() (*TransactionSafeWalletSourceDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *TransactionSource) SetDelegate(v TransactionSafeWalletSourceDelegate)`

SetDelegate sets Delegate field to given value.


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


### GetSubWalletId

`func (o *TransactionSource) GetSubWalletId() string`

GetSubWalletId returns the SubWalletId field if non-nil, zero value otherwise.

### GetSubWalletIdOk

`func (o *TransactionSource) GetSubWalletIdOk() (*string, bool)`

GetSubWalletIdOk returns a tuple with the SubWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWalletId

`func (o *TransactionSource) SetSubWalletId(v string)`

SetSubWalletId sets SubWalletId field to given value.

### HasSubWalletId

`func (o *TransactionSource) HasSubWalletId() bool`

HasSubWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


