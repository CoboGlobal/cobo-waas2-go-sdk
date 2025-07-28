# TokenizationSolContractCallParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TokenizationContractCallType**](TokenizationContractCallType.md) |  | [optional] 
**Instructions** | [**[]SolContractCallInstruction**](SolContractCallInstruction.md) |  | 

## Methods

### NewTokenizationSolContractCallParams

`func NewTokenizationSolContractCallParams(instructions []SolContractCallInstruction, ) *TokenizationSolContractCallParams`

NewTokenizationSolContractCallParams instantiates a new TokenizationSolContractCallParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationSolContractCallParamsWithDefaults

`func NewTokenizationSolContractCallParamsWithDefaults() *TokenizationSolContractCallParams`

NewTokenizationSolContractCallParamsWithDefaults instantiates a new TokenizationSolContractCallParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TokenizationSolContractCallParams) GetType() TokenizationContractCallType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenizationSolContractCallParams) GetTypeOk() (*TokenizationContractCallType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenizationSolContractCallParams) SetType(v TokenizationContractCallType)`

SetType sets Type field to given value.

### HasType

`func (o *TokenizationSolContractCallParams) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInstructions

`func (o *TokenizationSolContractCallParams) GetInstructions() []SolContractCallInstruction`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *TokenizationSolContractCallParams) GetInstructionsOk() (*[]SolContractCallInstruction, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *TokenizationSolContractCallParams) SetInstructions(v []SolContractCallInstruction)`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


