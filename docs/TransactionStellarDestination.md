# TransactionStellarDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**ContractParam** | [**TransactionStellarContractParam**](TransactionStellarContractParam.md) |  | 

## Methods

### NewTransactionStellarDestination

`func NewTransactionStellarDestination(destinationType TransactionDestinationType, contractParam TransactionStellarContractParam, ) *TransactionStellarDestination`

NewTransactionStellarDestination instantiates a new TransactionStellarDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionStellarDestinationWithDefaults

`func NewTransactionStellarDestinationWithDefaults() *TransactionStellarDestination`

NewTransactionStellarDestinationWithDefaults instantiates a new TransactionStellarDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionStellarDestination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionStellarDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionStellarDestination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetContractParam

`func (o *TransactionStellarDestination) GetContractParam() TransactionStellarContractParam`

GetContractParam returns the ContractParam field if non-nil, zero value otherwise.

### GetContractParamOk

`func (o *TransactionStellarDestination) GetContractParamOk() (*TransactionStellarContractParam, bool)`

GetContractParamOk returns a tuple with the ContractParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractParam

`func (o *TransactionStellarDestination) SetContractParam(v TransactionStellarContractParam)`

SetContractParam sets ContractParam field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


