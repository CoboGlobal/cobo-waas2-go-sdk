# RawMessageSignDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**MessageSignDestinationType**](MessageSignDestinationType.md) |  | 
**MsgHash** | **string** | Message hash to be signed, in hexadecimal format. | 

## Methods

### NewRawMessageSignDestination

`func NewRawMessageSignDestination(destinationType MessageSignDestinationType, msgHash string, ) *RawMessageSignDestination`

NewRawMessageSignDestination instantiates a new RawMessageSignDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawMessageSignDestinationWithDefaults

`func NewRawMessageSignDestinationWithDefaults() *RawMessageSignDestination`

NewRawMessageSignDestinationWithDefaults instantiates a new RawMessageSignDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *RawMessageSignDestination) GetDestinationType() MessageSignDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *RawMessageSignDestination) GetDestinationTypeOk() (*MessageSignDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *RawMessageSignDestination) SetDestinationType(v MessageSignDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetMsgHash

`func (o *RawMessageSignDestination) GetMsgHash() string`

GetMsgHash returns the MsgHash field if non-nil, zero value otherwise.

### GetMsgHashOk

`func (o *RawMessageSignDestination) GetMsgHashOk() (*string, bool)`

GetMsgHashOk returns a tuple with the MsgHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgHash

`func (o *RawMessageSignDestination) SetMsgHash(v string)`

SetMsgHash sets MsgHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


