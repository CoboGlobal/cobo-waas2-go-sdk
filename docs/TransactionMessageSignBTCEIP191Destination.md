# TransactionMessageSignBTCEIP191Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**Message** | **string** | The raw data of the message to be signed, encoded in Base64 format. | 

## Methods

### NewTransactionMessageSignBTCEIP191Destination

`func NewTransactionMessageSignBTCEIP191Destination(destinationType TransactionDestinationType, message string, ) *TransactionMessageSignBTCEIP191Destination`

NewTransactionMessageSignBTCEIP191Destination instantiates a new TransactionMessageSignBTCEIP191Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionMessageSignBTCEIP191DestinationWithDefaults

`func NewTransactionMessageSignBTCEIP191DestinationWithDefaults() *TransactionMessageSignBTCEIP191Destination`

NewTransactionMessageSignBTCEIP191DestinationWithDefaults instantiates a new TransactionMessageSignBTCEIP191Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionMessageSignBTCEIP191Destination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionMessageSignBTCEIP191Destination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionMessageSignBTCEIP191Destination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetMessage

`func (o *TransactionMessageSignBTCEIP191Destination) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransactionMessageSignBTCEIP191Destination) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransactionMessageSignBTCEIP191Destination) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


