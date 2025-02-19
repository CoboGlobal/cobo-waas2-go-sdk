# BTCEIP191MessageSignDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**MessageSignDestinationType**](MessageSignDestinationType.md) |  | 
**Message** | **string** | The raw data of the message to be signed, encoded in Base64 format. | 

## Methods

### NewBTCEIP191MessageSignDestination

`func NewBTCEIP191MessageSignDestination(destinationType MessageSignDestinationType, message string, ) *BTCEIP191MessageSignDestination`

NewBTCEIP191MessageSignDestination instantiates a new BTCEIP191MessageSignDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTCEIP191MessageSignDestinationWithDefaults

`func NewBTCEIP191MessageSignDestinationWithDefaults() *BTCEIP191MessageSignDestination`

NewBTCEIP191MessageSignDestinationWithDefaults instantiates a new BTCEIP191MessageSignDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *BTCEIP191MessageSignDestination) GetDestinationType() MessageSignDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *BTCEIP191MessageSignDestination) GetDestinationTypeOk() (*MessageSignDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *BTCEIP191MessageSignDestination) SetDestinationType(v MessageSignDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetMessage

`func (o *BTCEIP191MessageSignDestination) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BTCEIP191MessageSignDestination) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BTCEIP191MessageSignDestination) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


