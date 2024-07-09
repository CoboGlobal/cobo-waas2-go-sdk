# TransactionSafeWalletSourceDelegate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | Pointer to **string** | The wallet ID of the Delegate. This is required for initiating a transfer from Smart Contract Wallets (Safe{Wallet}).  | [optional] 
**WalletType** | Pointer to **string** | The wallet type of the Delegate. This is required for initiating a transfer from Smart Contract Wallets (Safe{Wallet}). | [optional] 
**WalletAddress** | Pointer to **string** | The wallet address of the Delegate. This is required for initiating a transfer from Smart Contract Wallets (Safe{Wallet}). | [optional] 
**MpcUsedKeyGroup** | Pointer to [**MpcSigningGroup**](MpcSigningGroup.md) |  | [optional] 

## Methods

### NewTransactionSafeWalletSourceDelegate

`func NewTransactionSafeWalletSourceDelegate() *TransactionSafeWalletSourceDelegate`

NewTransactionSafeWalletSourceDelegate instantiates a new TransactionSafeWalletSourceDelegate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSafeWalletSourceDelegateWithDefaults

`func NewTransactionSafeWalletSourceDelegateWithDefaults() *TransactionSafeWalletSourceDelegate`

NewTransactionSafeWalletSourceDelegateWithDefaults instantiates a new TransactionSafeWalletSourceDelegate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *TransactionSafeWalletSourceDelegate) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransactionSafeWalletSourceDelegate) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransactionSafeWalletSourceDelegate) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *TransactionSafeWalletSourceDelegate) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetWalletType

`func (o *TransactionSafeWalletSourceDelegate) GetWalletType() string`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *TransactionSafeWalletSourceDelegate) GetWalletTypeOk() (*string, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *TransactionSafeWalletSourceDelegate) SetWalletType(v string)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *TransactionSafeWalletSourceDelegate) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.

### GetWalletAddress

`func (o *TransactionSafeWalletSourceDelegate) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *TransactionSafeWalletSourceDelegate) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *TransactionSafeWalletSourceDelegate) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.

### HasWalletAddress

`func (o *TransactionSafeWalletSourceDelegate) HasWalletAddress() bool`

HasWalletAddress returns a boolean if a field has been set.

### GetMpcUsedKeyGroup

`func (o *TransactionSafeWalletSourceDelegate) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *TransactionSafeWalletSourceDelegate) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *TransactionSafeWalletSourceDelegate) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.

### HasMpcUsedKeyGroup

`func (o *TransactionSafeWalletSourceDelegate) HasMpcUsedKeyGroup() bool`

HasMpcUsedKeyGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


