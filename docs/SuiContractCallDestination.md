# SuiContractCallDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**ContractCallDestinationType**](ContractCallDestinationType.md) |  | 
**ContractParam** | [**SuiContractCallContractParam**](SuiContractCallContractParam.md) |  | 

## Methods

### NewSuiContractCallDestination

`func NewSuiContractCallDestination(destinationType ContractCallDestinationType, contractParam SuiContractCallContractParam, ) *SuiContractCallDestination`

NewSuiContractCallDestination instantiates a new SuiContractCallDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuiContractCallDestinationWithDefaults

`func NewSuiContractCallDestinationWithDefaults() *SuiContractCallDestination`

NewSuiContractCallDestinationWithDefaults instantiates a new SuiContractCallDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *SuiContractCallDestination) GetDestinationType() ContractCallDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *SuiContractCallDestination) GetDestinationTypeOk() (*ContractCallDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *SuiContractCallDestination) SetDestinationType(v ContractCallDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetContractParam

`func (o *SuiContractCallDestination) GetContractParam() SuiContractCallContractParam`

GetContractParam returns the ContractParam field if non-nil, zero value otherwise.

### GetContractParamOk

`func (o *SuiContractCallDestination) GetContractParamOk() (*SuiContractCallContractParam, bool)`

GetContractParamOk returns a tuple with the ContractParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractParam

`func (o *SuiContractCallDestination) SetContractParam(v SuiContractCallContractParam)`

SetContractParam sets ContractParam field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


