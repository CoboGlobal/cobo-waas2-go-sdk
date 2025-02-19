# MessageSignDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**MessageSignDestinationType**](MessageSignDestinationType.md) |  | 
**Message** | **string** | The raw data of the message to be signed, encoded in Base64 format. | 
**StructuredData** | **map[string]interface{}** | The structured data to be signed, formatted as a JSON object according to the EIP-712 standard. | 
**MsgHash** | **string** | Message hash to be signed, in hexadecimal format. | 

## Methods

### NewMessageSignDestination

`func NewMessageSignDestination(destinationType MessageSignDestinationType, message string, structuredData map[string]interface{}, msgHash string, ) *MessageSignDestination`

NewMessageSignDestination instantiates a new MessageSignDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageSignDestinationWithDefaults

`func NewMessageSignDestinationWithDefaults() *MessageSignDestination`

NewMessageSignDestinationWithDefaults instantiates a new MessageSignDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *MessageSignDestination) GetDestinationType() MessageSignDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *MessageSignDestination) GetDestinationTypeOk() (*MessageSignDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *MessageSignDestination) SetDestinationType(v MessageSignDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetMessage

`func (o *MessageSignDestination) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MessageSignDestination) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MessageSignDestination) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStructuredData

`func (o *MessageSignDestination) GetStructuredData() map[string]interface{}`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *MessageSignDestination) GetStructuredDataOk() (*map[string]interface{}, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *MessageSignDestination) SetStructuredData(v map[string]interface{})`

SetStructuredData sets StructuredData field to given value.


### GetMsgHash

`func (o *MessageSignDestination) GetMsgHash() string`

GetMsgHash returns the MsgHash field if non-nil, zero value otherwise.

### GetMsgHashOk

`func (o *MessageSignDestination) GetMsgHashOk() (*string, bool)`

GetMsgHashOk returns a tuple with the MsgHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgHash

`func (o *MessageSignDestination) SetMsgHash(v string)`

SetMsgHash sets MsgHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


