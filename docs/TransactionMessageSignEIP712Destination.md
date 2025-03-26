# TransactionMessageSignEIP712Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**RawStructuredData** | Pointer to **string** | The raw structured data to be signed, formatted as a JSON string. | [optional] 
**StructuredData** | **map[string]interface{}** | The structured data to be signed, formatted as a JSON object according to the EIP-712 standard. | 
**SafeTxExtraData** | Pointer to [**SafeTxExtraData**](SafeTxExtraData.md) |  | [optional] 

## Methods

### NewTransactionMessageSignEIP712Destination

`func NewTransactionMessageSignEIP712Destination(destinationType TransactionDestinationType, structuredData map[string]interface{}, ) *TransactionMessageSignEIP712Destination`

NewTransactionMessageSignEIP712Destination instantiates a new TransactionMessageSignEIP712Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionMessageSignEIP712DestinationWithDefaults

`func NewTransactionMessageSignEIP712DestinationWithDefaults() *TransactionMessageSignEIP712Destination`

NewTransactionMessageSignEIP712DestinationWithDefaults instantiates a new TransactionMessageSignEIP712Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionMessageSignEIP712Destination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionMessageSignEIP712Destination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionMessageSignEIP712Destination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetRawStructuredData

`func (o *TransactionMessageSignEIP712Destination) GetRawStructuredData() string`

GetRawStructuredData returns the RawStructuredData field if non-nil, zero value otherwise.

### GetRawStructuredDataOk

`func (o *TransactionMessageSignEIP712Destination) GetRawStructuredDataOk() (*string, bool)`

GetRawStructuredDataOk returns a tuple with the RawStructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawStructuredData

`func (o *TransactionMessageSignEIP712Destination) SetRawStructuredData(v string)`

SetRawStructuredData sets RawStructuredData field to given value.

### HasRawStructuredData

`func (o *TransactionMessageSignEIP712Destination) HasRawStructuredData() bool`

HasRawStructuredData returns a boolean if a field has been set.

### GetStructuredData

`func (o *TransactionMessageSignEIP712Destination) GetStructuredData() map[string]interface{}`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *TransactionMessageSignEIP712Destination) GetStructuredDataOk() (*map[string]interface{}, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *TransactionMessageSignEIP712Destination) SetStructuredData(v map[string]interface{})`

SetStructuredData sets StructuredData field to given value.


### GetSafeTxExtraData

`func (o *TransactionMessageSignEIP712Destination) GetSafeTxExtraData() SafeTxExtraData`

GetSafeTxExtraData returns the SafeTxExtraData field if non-nil, zero value otherwise.

### GetSafeTxExtraDataOk

`func (o *TransactionMessageSignEIP712Destination) GetSafeTxExtraDataOk() (*SafeTxExtraData, bool)`

GetSafeTxExtraDataOk returns a tuple with the SafeTxExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeTxExtraData

`func (o *TransactionMessageSignEIP712Destination) SetSafeTxExtraData(v SafeTxExtraData)`

SetSafeTxExtraData sets SafeTxExtraData field to given value.

### HasSafeTxExtraData

`func (o *TransactionMessageSignEIP712Destination) HasSafeTxExtraData() bool`

HasSafeTxExtraData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


