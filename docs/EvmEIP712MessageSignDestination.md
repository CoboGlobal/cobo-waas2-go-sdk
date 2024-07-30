# EvmEIP712MessageSignDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**MessageSignDestinationType**](MessageSignDestinationType.md) |  | 
**StructuredData** | **map[string]interface{}** | The structured data to be signed, formatted as a JSON object according to the EIP-712 standard. | 

## Methods

### NewEvmEIP712MessageSignDestination

`func NewEvmEIP712MessageSignDestination(destinationType MessageSignDestinationType, structuredData map[string]interface{}, ) *EvmEIP712MessageSignDestination`

NewEvmEIP712MessageSignDestination instantiates a new EvmEIP712MessageSignDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmEIP712MessageSignDestinationWithDefaults

`func NewEvmEIP712MessageSignDestinationWithDefaults() *EvmEIP712MessageSignDestination`

NewEvmEIP712MessageSignDestinationWithDefaults instantiates a new EvmEIP712MessageSignDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *EvmEIP712MessageSignDestination) GetDestinationType() MessageSignDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *EvmEIP712MessageSignDestination) GetDestinationTypeOk() (*MessageSignDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *EvmEIP712MessageSignDestination) SetDestinationType(v MessageSignDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetStructuredData

`func (o *EvmEIP712MessageSignDestination) GetStructuredData() map[string]interface{}`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *EvmEIP712MessageSignDestination) GetStructuredDataOk() (*map[string]interface{}, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *EvmEIP712MessageSignDestination) SetStructuredData(v map[string]interface{})`

SetStructuredData sets StructuredData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


