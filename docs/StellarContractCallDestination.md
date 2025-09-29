# StellarContractCallDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**ContractCallDestinationType**](ContractCallDestinationType.md) |  | 
**ContractParam** | [**StellarContractCallContractParam**](StellarContractCallContractParam.md) |  | 

## Methods

### NewStellarContractCallDestination

`func NewStellarContractCallDestination(destinationType ContractCallDestinationType, contractParam StellarContractCallContractParam, ) *StellarContractCallDestination`

NewStellarContractCallDestination instantiates a new StellarContractCallDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStellarContractCallDestinationWithDefaults

`func NewStellarContractCallDestinationWithDefaults() *StellarContractCallDestination`

NewStellarContractCallDestinationWithDefaults instantiates a new StellarContractCallDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *StellarContractCallDestination) GetDestinationType() ContractCallDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *StellarContractCallDestination) GetDestinationTypeOk() (*ContractCallDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *StellarContractCallDestination) SetDestinationType(v ContractCallDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetContractParam

`func (o *StellarContractCallDestination) GetContractParam() StellarContractCallContractParam`

GetContractParam returns the ContractParam field if non-nil, zero value otherwise.

### GetContractParamOk

`func (o *StellarContractCallDestination) GetContractParamOk() (*StellarContractCallContractParam, bool)`

GetContractParamOk returns a tuple with the ContractParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractParam

`func (o *StellarContractCallDestination) SetContractParam(v StellarContractCallContractParam)`

SetContractParam sets ContractParam field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


