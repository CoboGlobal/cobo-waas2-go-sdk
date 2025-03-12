# SolContractCallInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]SolContractCallAccount**](SolContractCallAccount.md) |  | 
**Data** | **string** | The Base64-encoded instruction data used for interacting with a Solana program.  | 
**ProgramId** | **string** | The address of the Solana program (smart contract).   | 

## Methods

### NewSolContractCallInstruction

`func NewSolContractCallInstruction(accounts []SolContractCallAccount, data string, programId string, ) *SolContractCallInstruction`

NewSolContractCallInstruction instantiates a new SolContractCallInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolContractCallInstructionWithDefaults

`func NewSolContractCallInstructionWithDefaults() *SolContractCallInstruction`

NewSolContractCallInstructionWithDefaults instantiates a new SolContractCallInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *SolContractCallInstruction) GetAccounts() []SolContractCallAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *SolContractCallInstruction) GetAccountsOk() (*[]SolContractCallAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *SolContractCallInstruction) SetAccounts(v []SolContractCallAccount)`

SetAccounts sets Accounts field to given value.


### GetData

`func (o *SolContractCallInstruction) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SolContractCallInstruction) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SolContractCallInstruction) SetData(v string)`

SetData sets Data field to given value.


### GetProgramId

`func (o *SolContractCallInstruction) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *SolContractCallInstruction) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *SolContractCallInstruction) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


