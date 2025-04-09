# BTCBIP137MessageSignDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**MessageSignDestinationType**](MessageSignDestinationType.md) |  | 
**MessageBip137** | **string** | Message to be signed, in hexadecimal format. | 

## Methods

### NewBTCBIP137MessageSignDestination

`func NewBTCBIP137MessageSignDestination(destinationType MessageSignDestinationType, messageBip137 string, ) *BTCBIP137MessageSignDestination`

NewBTCBIP137MessageSignDestination instantiates a new BTCBIP137MessageSignDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTCBIP137MessageSignDestinationWithDefaults

`func NewBTCBIP137MessageSignDestinationWithDefaults() *BTCBIP137MessageSignDestination`

NewBTCBIP137MessageSignDestinationWithDefaults instantiates a new BTCBIP137MessageSignDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *BTCBIP137MessageSignDestination) GetDestinationType() MessageSignDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *BTCBIP137MessageSignDestination) GetDestinationTypeOk() (*MessageSignDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *BTCBIP137MessageSignDestination) SetDestinationType(v MessageSignDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetMessageBip137

`func (o *BTCBIP137MessageSignDestination) GetMessageBip137() string`

GetMessageBip137 returns the MessageBip137 field if non-nil, zero value otherwise.

### GetMessageBip137Ok

`func (o *BTCBIP137MessageSignDestination) GetMessageBip137Ok() (*string, bool)`

GetMessageBip137Ok returns a tuple with the MessageBip137 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBip137

`func (o *BTCBIP137MessageSignDestination) SetMessageBip137(v string)`

SetMessageBip137 sets MessageBip137 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


