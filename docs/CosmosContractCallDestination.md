# CosmosContractCallDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**ContractCallDestinationType**](ContractCallDestinationType.md) |  | 
**CosmosMessages** | [**[]CosmosContractCallMessage**](CosmosContractCallMessage.md) |  | 

## Methods

### NewCosmosContractCallDestination

`func NewCosmosContractCallDestination(destinationType ContractCallDestinationType, cosmosMessages []CosmosContractCallMessage, ) *CosmosContractCallDestination`

NewCosmosContractCallDestination instantiates a new CosmosContractCallDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCosmosContractCallDestinationWithDefaults

`func NewCosmosContractCallDestinationWithDefaults() *CosmosContractCallDestination`

NewCosmosContractCallDestinationWithDefaults instantiates a new CosmosContractCallDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *CosmosContractCallDestination) GetDestinationType() ContractCallDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *CosmosContractCallDestination) GetDestinationTypeOk() (*ContractCallDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *CosmosContractCallDestination) SetDestinationType(v ContractCallDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetCosmosMessages

`func (o *CosmosContractCallDestination) GetCosmosMessages() []CosmosContractCallMessage`

GetCosmosMessages returns the CosmosMessages field if non-nil, zero value otherwise.

### GetCosmosMessagesOk

`func (o *CosmosContractCallDestination) GetCosmosMessagesOk() (*[]CosmosContractCallMessage, bool)`

GetCosmosMessagesOk returns a tuple with the CosmosMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosmosMessages

`func (o *CosmosContractCallDestination) SetCosmosMessages(v []CosmosContractCallMessage)`

SetCosmosMessages sets CosmosMessages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


