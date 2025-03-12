# TransactionSolContractInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]TransactionSolContractAccount**](TransactionSolContractAccount.md) |  | [optional] 
**Data** | Pointer to **string** | data used for calling Solana contract..  | [optional] 
**ProgramId** | Pointer to **string** | contract address. when calling a Solana contract, the to_address parameter needs to match the program_id parameter. If multiple contracts are being called, then the to_address parameter should match the program_id parameter of the first instruction.  | [optional] 

## Methods

### NewTransactionSolContractInstruction

`func NewTransactionSolContractInstruction() *TransactionSolContractInstruction`

NewTransactionSolContractInstruction instantiates a new TransactionSolContractInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSolContractInstructionWithDefaults

`func NewTransactionSolContractInstructionWithDefaults() *TransactionSolContractInstruction`

NewTransactionSolContractInstructionWithDefaults instantiates a new TransactionSolContractInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *TransactionSolContractInstruction) GetAccounts() []TransactionSolContractAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *TransactionSolContractInstruction) GetAccountsOk() (*[]TransactionSolContractAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *TransactionSolContractInstruction) SetAccounts(v []TransactionSolContractAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *TransactionSolContractInstruction) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetData

`func (o *TransactionSolContractInstruction) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionSolContractInstruction) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionSolContractInstruction) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *TransactionSolContractInstruction) HasData() bool`

HasData returns a boolean if a field has been set.

### GetProgramId

`func (o *TransactionSolContractInstruction) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *TransactionSolContractInstruction) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *TransactionSolContractInstruction) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *TransactionSolContractInstruction) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


