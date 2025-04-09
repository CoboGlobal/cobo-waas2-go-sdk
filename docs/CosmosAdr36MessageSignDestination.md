# CosmosAdr36MessageSignDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**MessageSignDestinationType**](MessageSignDestinationType.md) |  | 
**MessageCosmosAdr36** | **string** | Message to be signed, in hexadecimal format. | 

## Methods

### NewCosmosAdr36MessageSignDestination

`func NewCosmosAdr36MessageSignDestination(destinationType MessageSignDestinationType, messageCosmosAdr36 string, ) *CosmosAdr36MessageSignDestination`

NewCosmosAdr36MessageSignDestination instantiates a new CosmosAdr36MessageSignDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCosmosAdr36MessageSignDestinationWithDefaults

`func NewCosmosAdr36MessageSignDestinationWithDefaults() *CosmosAdr36MessageSignDestination`

NewCosmosAdr36MessageSignDestinationWithDefaults instantiates a new CosmosAdr36MessageSignDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *CosmosAdr36MessageSignDestination) GetDestinationType() MessageSignDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *CosmosAdr36MessageSignDestination) GetDestinationTypeOk() (*MessageSignDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *CosmosAdr36MessageSignDestination) SetDestinationType(v MessageSignDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetMessageCosmosAdr36

`func (o *CosmosAdr36MessageSignDestination) GetMessageCosmosAdr36() string`

GetMessageCosmosAdr36 returns the MessageCosmosAdr36 field if non-nil, zero value otherwise.

### GetMessageCosmosAdr36Ok

`func (o *CosmosAdr36MessageSignDestination) GetMessageCosmosAdr36Ok() (*string, bool)`

GetMessageCosmosAdr36Ok returns a tuple with the MessageCosmosAdr36 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCosmosAdr36

`func (o *CosmosAdr36MessageSignDestination) SetMessageCosmosAdr36(v string)`

SetMessageCosmosAdr36 sets MessageCosmosAdr36 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


