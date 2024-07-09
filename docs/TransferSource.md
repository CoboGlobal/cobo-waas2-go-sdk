# TransferSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**WalletSubtype**](WalletSubtype.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**AccountInput** | Pointer to [**TransactionMPCWalletSourceAccountInput**](TransactionMPCWalletSourceAccountInput.md) |  | [optional] 
**UtxoInputs** | Pointer to [**MpcTransferSourceAllOfUtxoInputs**](MpcTransferSourceAllOfUtxoInputs.md) |  | [optional] 
**MpcUsedKeyGroup** | [**MpcSigningGroup**](MpcSigningGroup.md) |  | 
**Address** | **string** | The wallet address. | 
**Delegate** | [**SafeTransferSourceAllOfDelegate**](SafeTransferSourceAllOfDelegate.md) |  | 
**SubWalletId** | **string** | The exchange trading account or the sub-wallet ID. | 

## Methods

### NewTransferSource

`func NewTransferSource(sourceType WalletSubtype, walletId string, mpcUsedKeyGroup MpcSigningGroup, address string, delegate SafeTransferSourceAllOfDelegate, subWalletId string, ) *TransferSource`

NewTransferSource instantiates a new TransferSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferSourceWithDefaults

`func NewTransferSourceWithDefaults() *TransferSource`

NewTransferSourceWithDefaults instantiates a new TransferSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *TransferSource) GetSourceType() WalletSubtype`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransferSource) GetSourceTypeOk() (*WalletSubtype, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransferSource) SetSourceType(v WalletSubtype)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *TransferSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransferSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransferSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAccountInput

`func (o *TransferSource) GetAccountInput() TransactionMPCWalletSourceAccountInput`

GetAccountInput returns the AccountInput field if non-nil, zero value otherwise.

### GetAccountInputOk

`func (o *TransferSource) GetAccountInputOk() (*TransactionMPCWalletSourceAccountInput, bool)`

GetAccountInputOk returns a tuple with the AccountInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInput

`func (o *TransferSource) SetAccountInput(v TransactionMPCWalletSourceAccountInput)`

SetAccountInput sets AccountInput field to given value.

### HasAccountInput

`func (o *TransferSource) HasAccountInput() bool`

HasAccountInput returns a boolean if a field has been set.

### GetUtxoInputs

`func (o *TransferSource) GetUtxoInputs() MpcTransferSourceAllOfUtxoInputs`

GetUtxoInputs returns the UtxoInputs field if non-nil, zero value otherwise.

### GetUtxoInputsOk

`func (o *TransferSource) GetUtxoInputsOk() (*MpcTransferSourceAllOfUtxoInputs, bool)`

GetUtxoInputsOk returns a tuple with the UtxoInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoInputs

`func (o *TransferSource) SetUtxoInputs(v MpcTransferSourceAllOfUtxoInputs)`

SetUtxoInputs sets UtxoInputs field to given value.

### HasUtxoInputs

`func (o *TransferSource) HasUtxoInputs() bool`

HasUtxoInputs returns a boolean if a field has been set.

### GetMpcUsedKeyGroup

`func (o *TransferSource) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *TransferSource) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *TransferSource) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.


### GetAddress

`func (o *TransferSource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransferSource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransferSource) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetDelegate

`func (o *TransferSource) GetDelegate() SafeTransferSourceAllOfDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *TransferSource) GetDelegateOk() (*SafeTransferSourceAllOfDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *TransferSource) SetDelegate(v SafeTransferSourceAllOfDelegate)`

SetDelegate sets Delegate field to given value.


### GetSubWalletId

`func (o *TransferSource) GetSubWalletId() string`

GetSubWalletId returns the SubWalletId field if non-nil, zero value otherwise.

### GetSubWalletIdOk

`func (o *TransferSource) GetSubWalletIdOk() (*string, bool)`

GetSubWalletIdOk returns a tuple with the SubWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWalletId

`func (o *TransferSource) SetSubWalletId(v string)`

SetSubWalletId sets SubWalletId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


