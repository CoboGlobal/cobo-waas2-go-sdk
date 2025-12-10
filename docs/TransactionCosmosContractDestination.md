# TransactionCosmosContractDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**CosmosMessages** | [**[]TransactionCosmosMessage**](TransactionCosmosMessage.md) |  | 

## Methods

### NewTransactionCosmosContractDestination

`func NewTransactionCosmosContractDestination(destinationType TransactionDestinationType, cosmosMessages []TransactionCosmosMessage, ) *TransactionCosmosContractDestination`

NewTransactionCosmosContractDestination instantiates a new TransactionCosmosContractDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCosmosContractDestinationWithDefaults

`func NewTransactionCosmosContractDestinationWithDefaults() *TransactionCosmosContractDestination`

NewTransactionCosmosContractDestinationWithDefaults instantiates a new TransactionCosmosContractDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionCosmosContractDestination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionCosmosContractDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionCosmosContractDestination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetCosmosMessages

`func (o *TransactionCosmosContractDestination) GetCosmosMessages() []TransactionCosmosMessage`

GetCosmosMessages returns the CosmosMessages field if non-nil, zero value otherwise.

### GetCosmosMessagesOk

`func (o *TransactionCosmosContractDestination) GetCosmosMessagesOk() (*[]TransactionCosmosMessage, bool)`

GetCosmosMessagesOk returns a tuple with the CosmosMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosmosMessages

`func (o *TransactionCosmosContractDestination) SetCosmosMessages(v []TransactionCosmosMessage)`

SetCosmosMessages sets CosmosMessages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


