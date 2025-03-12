# SolContractCallDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**ContractCallDestinationType**](ContractCallDestinationType.md) |  | 
**Instructions** | [**[]SolContractCallInstruction**](SolContractCallInstruction.md) |  | 

## Methods

### NewSolContractCallDestination

`func NewSolContractCallDestination(destinationType ContractCallDestinationType, instructions []SolContractCallInstruction, ) *SolContractCallDestination`

NewSolContractCallDestination instantiates a new SolContractCallDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolContractCallDestinationWithDefaults

`func NewSolContractCallDestinationWithDefaults() *SolContractCallDestination`

NewSolContractCallDestinationWithDefaults instantiates a new SolContractCallDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *SolContractCallDestination) GetDestinationType() ContractCallDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *SolContractCallDestination) GetDestinationTypeOk() (*ContractCallDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *SolContractCallDestination) SetDestinationType(v ContractCallDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetInstructions

`func (o *SolContractCallDestination) GetInstructions() []SolContractCallInstruction`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *SolContractCallDestination) GetInstructionsOk() (*[]SolContractCallInstruction, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *SolContractCallDestination) SetInstructions(v []SolContractCallInstruction)`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


