# TransactionMessageSignDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**Message** | Pointer to **string** | The raw data to be signed that are encoded in Base64. | [optional] 
**StructuredData** | Pointer to **string** | The structured data to be signed that are encoded in JSON. | [optional] 

## Methods

### NewTransactionMessageSignDestination

`func NewTransactionMessageSignDestination(destinationType TransactionDestinationType, ) *TransactionMessageSignDestination`

NewTransactionMessageSignDestination instantiates a new TransactionMessageSignDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionMessageSignDestinationWithDefaults

`func NewTransactionMessageSignDestinationWithDefaults() *TransactionMessageSignDestination`

NewTransactionMessageSignDestinationWithDefaults instantiates a new TransactionMessageSignDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionMessageSignDestination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionMessageSignDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionMessageSignDestination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetMessage

`func (o *TransactionMessageSignDestination) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransactionMessageSignDestination) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransactionMessageSignDestination) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TransactionMessageSignDestination) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStructuredData

`func (o *TransactionMessageSignDestination) GetStructuredData() string`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *TransactionMessageSignDestination) GetStructuredDataOk() (*string, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *TransactionMessageSignDestination) SetStructuredData(v string)`

SetStructuredData sets StructuredData field to given value.

### HasStructuredData

`func (o *TransactionMessageSignDestination) HasStructuredData() bool`

HasStructuredData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


