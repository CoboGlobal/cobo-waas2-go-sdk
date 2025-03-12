# ContractCallDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**ContractCallDestinationType**](ContractCallDestinationType.md) |  | 
**Address** | **string** | The destination address. | 
**Value** | Pointer to **string** | The transfer amount. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | [optional] 
**Calldata** | **string** | The data that is used to invoke a specific function or method within the specified contract at the destination address.  | 
**Instructions** | [**[]SolContractCallInstruction**](SolContractCallInstruction.md) |  | 

## Methods

### NewContractCallDestination

`func NewContractCallDestination(destinationType ContractCallDestinationType, address string, calldata string, instructions []SolContractCallInstruction, ) *ContractCallDestination`

NewContractCallDestination instantiates a new ContractCallDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractCallDestinationWithDefaults

`func NewContractCallDestinationWithDefaults() *ContractCallDestination`

NewContractCallDestinationWithDefaults instantiates a new ContractCallDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *ContractCallDestination) GetDestinationType() ContractCallDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *ContractCallDestination) GetDestinationTypeOk() (*ContractCallDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *ContractCallDestination) SetDestinationType(v ContractCallDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetAddress

`func (o *ContractCallDestination) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractCallDestination) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractCallDestination) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetValue

`func (o *ContractCallDestination) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContractCallDestination) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContractCallDestination) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ContractCallDestination) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCalldata

`func (o *ContractCallDestination) GetCalldata() string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *ContractCallDestination) GetCalldataOk() (*string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *ContractCallDestination) SetCalldata(v string)`

SetCalldata sets Calldata field to given value.


### GetInstructions

`func (o *ContractCallDestination) GetInstructions() []SolContractCallInstruction`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ContractCallDestination) GetInstructionsOk() (*[]SolContractCallInstruction, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ContractCallDestination) SetInstructions(v []SolContractCallInstruction)`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


