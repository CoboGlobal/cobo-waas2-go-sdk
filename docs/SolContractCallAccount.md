# SolContractCallAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | **string** | The public key of the account. If the account is a signer of this transaction, this property must be the same as the value of the &#x60;source.address&#x60; property.  | 
**IsSigner** | **bool** | Whether the account is the signer of this transaction: - &#x60;true&#x60;: The account is a signer. - &#x60;false&#x60;: The account is not a signer.  | 
**IsWritable** | **bool** | Whether the account can be modified by the instruction: - &#x60;true&#x60;: The account can be modified by the instruction. - &#x60;false&#x60;: The account cannot be modified by the instruction.  | 

## Methods

### NewSolContractCallAccount

`func NewSolContractCallAccount(pubkey string, isSigner bool, isWritable bool, ) *SolContractCallAccount`

NewSolContractCallAccount instantiates a new SolContractCallAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolContractCallAccountWithDefaults

`func NewSolContractCallAccountWithDefaults() *SolContractCallAccount`

NewSolContractCallAccountWithDefaults instantiates a new SolContractCallAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *SolContractCallAccount) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *SolContractCallAccount) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *SolContractCallAccount) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.


### GetIsSigner

`func (o *SolContractCallAccount) GetIsSigner() bool`

GetIsSigner returns the IsSigner field if non-nil, zero value otherwise.

### GetIsSignerOk

`func (o *SolContractCallAccount) GetIsSignerOk() (*bool, bool)`

GetIsSignerOk returns a tuple with the IsSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSigner

`func (o *SolContractCallAccount) SetIsSigner(v bool)`

SetIsSigner sets IsSigner field to given value.


### GetIsWritable

`func (o *SolContractCallAccount) GetIsWritable() bool`

GetIsWritable returns the IsWritable field if non-nil, zero value otherwise.

### GetIsWritableOk

`func (o *SolContractCallAccount) GetIsWritableOk() (*bool, bool)`

GetIsWritableOk returns a tuple with the IsWritable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWritable

`func (o *SolContractCallAccount) SetIsWritable(v bool)`

SetIsWritable sets IsWritable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


