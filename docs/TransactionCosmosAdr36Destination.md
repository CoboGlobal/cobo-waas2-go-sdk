# TransactionCosmosAdr36Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**MessageCosmosAdr36** | **string** | Message to be signed, in hexadecimal format. | 

## Methods

### NewTransactionCosmosAdr36Destination

`func NewTransactionCosmosAdr36Destination(destinationType TransactionDestinationType, messageCosmosAdr36 string, ) *TransactionCosmosAdr36Destination`

NewTransactionCosmosAdr36Destination instantiates a new TransactionCosmosAdr36Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCosmosAdr36DestinationWithDefaults

`func NewTransactionCosmosAdr36DestinationWithDefaults() *TransactionCosmosAdr36Destination`

NewTransactionCosmosAdr36DestinationWithDefaults instantiates a new TransactionCosmosAdr36Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionCosmosAdr36Destination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionCosmosAdr36Destination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionCosmosAdr36Destination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetMessageCosmosAdr36

`func (o *TransactionCosmosAdr36Destination) GetMessageCosmosAdr36() string`

GetMessageCosmosAdr36 returns the MessageCosmosAdr36 field if non-nil, zero value otherwise.

### GetMessageCosmosAdr36Ok

`func (o *TransactionCosmosAdr36Destination) GetMessageCosmosAdr36Ok() (*string, bool)`

GetMessageCosmosAdr36Ok returns a tuple with the MessageCosmosAdr36 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCosmosAdr36

`func (o *TransactionCosmosAdr36Destination) SetMessageCosmosAdr36(v string)`

SetMessageCosmosAdr36 sets MessageCosmosAdr36 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


