# TokenizationContractCallParamsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TokenizationContractCallType**](TokenizationContractCallType.md) |  | [optional] 
**Calldata** | **string** | The data that is used to invoke a specific function or method within the specified contract at the destination address.  | 
**Value** | Pointer to **string** | The transfer amount. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | [optional] 
**Instructions** | [**[]SolContractCallInstruction**](SolContractCallInstruction.md) |  | 

## Methods

### NewTokenizationContractCallParamsData

`func NewTokenizationContractCallParamsData(calldata string, instructions []SolContractCallInstruction, ) *TokenizationContractCallParamsData`

NewTokenizationContractCallParamsData instantiates a new TokenizationContractCallParamsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationContractCallParamsDataWithDefaults

`func NewTokenizationContractCallParamsDataWithDefaults() *TokenizationContractCallParamsData`

NewTokenizationContractCallParamsDataWithDefaults instantiates a new TokenizationContractCallParamsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TokenizationContractCallParamsData) GetType() TokenizationContractCallType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenizationContractCallParamsData) GetTypeOk() (*TokenizationContractCallType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenizationContractCallParamsData) SetType(v TokenizationContractCallType)`

SetType sets Type field to given value.

### HasType

`func (o *TokenizationContractCallParamsData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCalldata

`func (o *TokenizationContractCallParamsData) GetCalldata() string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *TokenizationContractCallParamsData) GetCalldataOk() (*string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *TokenizationContractCallParamsData) SetCalldata(v string)`

SetCalldata sets Calldata field to given value.


### GetValue

`func (o *TokenizationContractCallParamsData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TokenizationContractCallParamsData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TokenizationContractCallParamsData) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TokenizationContractCallParamsData) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetInstructions

`func (o *TokenizationContractCallParamsData) GetInstructions() []SolContractCallInstruction`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *TokenizationContractCallParamsData) GetInstructionsOk() (*[]SolContractCallInstruction, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *TokenizationContractCallParamsData) SetInstructions(v []SolContractCallInstruction)`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


