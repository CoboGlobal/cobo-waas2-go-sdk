# TransactionRawMessageSignDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**MsgHash** | Pointer to **string** | Message hash to be signed, in hexadecimal format. | [optional] 

## Methods

### NewTransactionRawMessageSignDestination

`func NewTransactionRawMessageSignDestination(destinationType TransactionDestinationType, ) *TransactionRawMessageSignDestination`

NewTransactionRawMessageSignDestination instantiates a new TransactionRawMessageSignDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRawMessageSignDestinationWithDefaults

`func NewTransactionRawMessageSignDestinationWithDefaults() *TransactionRawMessageSignDestination`

NewTransactionRawMessageSignDestinationWithDefaults instantiates a new TransactionRawMessageSignDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionRawMessageSignDestination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionRawMessageSignDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionRawMessageSignDestination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetMsgHash

`func (o *TransactionRawMessageSignDestination) GetMsgHash() string`

GetMsgHash returns the MsgHash field if non-nil, zero value otherwise.

### GetMsgHashOk

`func (o *TransactionRawMessageSignDestination) GetMsgHashOk() (*string, bool)`

GetMsgHashOk returns a tuple with the MsgHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgHash

`func (o *TransactionRawMessageSignDestination) SetMsgHash(v string)`

SetMsgHash sets MsgHash field to given value.

### HasMsgHash

`func (o *TransactionRawMessageSignDestination) HasMsgHash() bool`

HasMsgHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


