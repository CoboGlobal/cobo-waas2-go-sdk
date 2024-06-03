# SignMessageDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Raw data to be signed, Base 64 encoded | [optional] 
**StructuredData** | Pointer to **string** | Structured data to be signed, JSON encoded | [optional] 

## Methods

### NewSignMessageDestination

`func NewSignMessageDestination() *SignMessageDestination`

NewSignMessageDestination instantiates a new SignMessageDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignMessageDestinationWithDefaults

`func NewSignMessageDestinationWithDefaults() *SignMessageDestination`

NewSignMessageDestinationWithDefaults instantiates a new SignMessageDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SignMessageDestination) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SignMessageDestination) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SignMessageDestination) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SignMessageDestination) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStructuredData

`func (o *SignMessageDestination) GetStructuredData() string`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *SignMessageDestination) GetStructuredDataOk() (*string, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *SignMessageDestination) SetStructuredData(v string)`

SetStructuredData sets StructuredData field to given value.

### HasStructuredData

`func (o *SignMessageDestination) HasStructuredData() bool`

HasStructuredData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


