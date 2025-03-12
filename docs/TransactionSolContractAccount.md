# TransactionSolContractAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | Pointer to **string** | account address. If the account is signer, pubkey needs to match the from_address parameter.  | [optional] 
**IsSigner** | Pointer to **bool** | boolean value indicating whether the account can sign transactions.  | [optional] 
**IsWritable** | Pointer to **bool** | boolean value indicating whether the account can be modified.  | [optional] 

## Methods

### NewTransactionSolContractAccount

`func NewTransactionSolContractAccount() *TransactionSolContractAccount`

NewTransactionSolContractAccount instantiates a new TransactionSolContractAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSolContractAccountWithDefaults

`func NewTransactionSolContractAccountWithDefaults() *TransactionSolContractAccount`

NewTransactionSolContractAccountWithDefaults instantiates a new TransactionSolContractAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *TransactionSolContractAccount) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *TransactionSolContractAccount) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *TransactionSolContractAccount) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *TransactionSolContractAccount) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetIsSigner

`func (o *TransactionSolContractAccount) GetIsSigner() bool`

GetIsSigner returns the IsSigner field if non-nil, zero value otherwise.

### GetIsSignerOk

`func (o *TransactionSolContractAccount) GetIsSignerOk() (*bool, bool)`

GetIsSignerOk returns a tuple with the IsSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSigner

`func (o *TransactionSolContractAccount) SetIsSigner(v bool)`

SetIsSigner sets IsSigner field to given value.

### HasIsSigner

`func (o *TransactionSolContractAccount) HasIsSigner() bool`

HasIsSigner returns a boolean if a field has been set.

### GetIsWritable

`func (o *TransactionSolContractAccount) GetIsWritable() bool`

GetIsWritable returns the IsWritable field if non-nil, zero value otherwise.

### GetIsWritableOk

`func (o *TransactionSolContractAccount) GetIsWritableOk() (*bool, bool)`

GetIsWritableOk returns a tuple with the IsWritable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWritable

`func (o *TransactionSolContractAccount) SetIsWritable(v bool)`

SetIsWritable sets IsWritable field to given value.

### HasIsWritable

`func (o *TransactionSolContractAccount) HasIsWritable() bool`

HasIsWritable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


