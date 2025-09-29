# TransactionSolContractDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**Instructions** | Pointer to [**[]TransactionSolContractInstruction**](TransactionSolContractInstruction.md) |  | [optional] 
**AddressLookupTableAccounts** | Pointer to [**[]TransactionSolContractAddressLookupTableAccount**](TransactionSolContractAddressLookupTableAccount.md) |  | [optional] 

## Methods

### NewTransactionSolContractDestination

`func NewTransactionSolContractDestination(destinationType TransactionDestinationType, ) *TransactionSolContractDestination`

NewTransactionSolContractDestination instantiates a new TransactionSolContractDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSolContractDestinationWithDefaults

`func NewTransactionSolContractDestinationWithDefaults() *TransactionSolContractDestination`

NewTransactionSolContractDestinationWithDefaults instantiates a new TransactionSolContractDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionSolContractDestination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionSolContractDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionSolContractDestination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetInstructions

`func (o *TransactionSolContractDestination) GetInstructions() []TransactionSolContractInstruction`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *TransactionSolContractDestination) GetInstructionsOk() (*[]TransactionSolContractInstruction, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *TransactionSolContractDestination) SetInstructions(v []TransactionSolContractInstruction)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *TransactionSolContractDestination) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetAddressLookupTableAccounts

`func (o *TransactionSolContractDestination) GetAddressLookupTableAccounts() []TransactionSolContractAddressLookupTableAccount`

GetAddressLookupTableAccounts returns the AddressLookupTableAccounts field if non-nil, zero value otherwise.

### GetAddressLookupTableAccountsOk

`func (o *TransactionSolContractDestination) GetAddressLookupTableAccountsOk() (*[]TransactionSolContractAddressLookupTableAccount, bool)`

GetAddressLookupTableAccountsOk returns a tuple with the AddressLookupTableAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLookupTableAccounts

`func (o *TransactionSolContractDestination) SetAddressLookupTableAccounts(v []TransactionSolContractAddressLookupTableAccount)`

SetAddressLookupTableAccounts sets AddressLookupTableAccounts field to given value.

### HasAddressLookupTableAccounts

`func (o *TransactionSolContractDestination) HasAddressLookupTableAccounts() bool`

HasAddressLookupTableAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


