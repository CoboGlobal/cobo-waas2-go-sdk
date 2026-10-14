# TransactionSuiDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**ContractParam** | [**SuiContractCallContractParam**](SuiContractCallContractParam.md) |  | 

## Methods

### NewTransactionSuiDestination

`func NewTransactionSuiDestination(destinationType TransactionDestinationType, contractParam SuiContractCallContractParam, ) *TransactionSuiDestination`

NewTransactionSuiDestination instantiates a new TransactionSuiDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSuiDestinationWithDefaults

`func NewTransactionSuiDestinationWithDefaults() *TransactionSuiDestination`

NewTransactionSuiDestinationWithDefaults instantiates a new TransactionSuiDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionSuiDestination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionSuiDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionSuiDestination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetContractParam

`func (o *TransactionSuiDestination) GetContractParam() SuiContractCallContractParam`

GetContractParam returns the ContractParam field if non-nil, zero value otherwise.

### GetContractParamOk

`func (o *TransactionSuiDestination) GetContractParamOk() (*SuiContractCallContractParam, bool)`

GetContractParamOk returns a tuple with the ContractParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractParam

`func (o *TransactionSuiDestination) SetContractParam(v SuiContractCallContractParam)`

SetContractParam sets ContractParam field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


