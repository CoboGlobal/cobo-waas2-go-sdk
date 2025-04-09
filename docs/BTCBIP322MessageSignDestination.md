# BTCBIP322MessageSignDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**MessageSignDestinationType**](MessageSignDestinationType.md) |  | 
**MessageBip322** | **string** | Message to be signed, in hexadecimal format. | 

## Methods

### NewBTCBIP322MessageSignDestination

`func NewBTCBIP322MessageSignDestination(destinationType MessageSignDestinationType, messageBip322 string, ) *BTCBIP322MessageSignDestination`

NewBTCBIP322MessageSignDestination instantiates a new BTCBIP322MessageSignDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTCBIP322MessageSignDestinationWithDefaults

`func NewBTCBIP322MessageSignDestinationWithDefaults() *BTCBIP322MessageSignDestination`

NewBTCBIP322MessageSignDestinationWithDefaults instantiates a new BTCBIP322MessageSignDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *BTCBIP322MessageSignDestination) GetDestinationType() MessageSignDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *BTCBIP322MessageSignDestination) GetDestinationTypeOk() (*MessageSignDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *BTCBIP322MessageSignDestination) SetDestinationType(v MessageSignDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetMessageBip322

`func (o *BTCBIP322MessageSignDestination) GetMessageBip322() string`

GetMessageBip322 returns the MessageBip322 field if non-nil, zero value otherwise.

### GetMessageBip322Ok

`func (o *BTCBIP322MessageSignDestination) GetMessageBip322Ok() (*string, bool)`

GetMessageBip322Ok returns a tuple with the MessageBip322 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBip322

`func (o *BTCBIP322MessageSignDestination) SetMessageBip322(v string)`

SetMessageBip322 sets MessageBip322 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


