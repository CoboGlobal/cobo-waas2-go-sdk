# TransactionBIP322Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**MessageBip322** | **string** | Message to be signed, in hexadecimal format. | 

## Methods

### NewTransactionBIP322Destination

`func NewTransactionBIP322Destination(destinationType TransactionDestinationType, messageBip322 string, ) *TransactionBIP322Destination`

NewTransactionBIP322Destination instantiates a new TransactionBIP322Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionBIP322DestinationWithDefaults

`func NewTransactionBIP322DestinationWithDefaults() *TransactionBIP322Destination`

NewTransactionBIP322DestinationWithDefaults instantiates a new TransactionBIP322Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionBIP322Destination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionBIP322Destination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionBIP322Destination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetMessageBip322

`func (o *TransactionBIP322Destination) GetMessageBip322() string`

GetMessageBip322 returns the MessageBip322 field if non-nil, zero value otherwise.

### GetMessageBip322Ok

`func (o *TransactionBIP322Destination) GetMessageBip322Ok() (*string, bool)`

GetMessageBip322Ok returns a tuple with the MessageBip322 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBip322

`func (o *TransactionBIP322Destination) SetMessageBip322(v string)`

SetMessageBip322 sets MessageBip322 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


