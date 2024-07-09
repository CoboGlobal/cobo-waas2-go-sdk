# TransactionAddressDestinationUtxoOutputs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outputs** | Pointer to [**[]TransactionAddressDestinationUtxoOutputsOutputsInner**](TransactionAddressDestinationUtxoOutputsOutputsInner.md) |  | [optional] 
**ChangeAddress** | Pointer to **string** | The address used to receive the remaining funds or change from the transaction. | [optional] 

## Methods

### NewTransactionAddressDestinationUtxoOutputs

`func NewTransactionAddressDestinationUtxoOutputs() *TransactionAddressDestinationUtxoOutputs`

NewTransactionAddressDestinationUtxoOutputs instantiates a new TransactionAddressDestinationUtxoOutputs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionAddressDestinationUtxoOutputsWithDefaults

`func NewTransactionAddressDestinationUtxoOutputsWithDefaults() *TransactionAddressDestinationUtxoOutputs`

NewTransactionAddressDestinationUtxoOutputsWithDefaults instantiates a new TransactionAddressDestinationUtxoOutputs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutputs

`func (o *TransactionAddressDestinationUtxoOutputs) GetOutputs() []TransactionAddressDestinationUtxoOutputsOutputsInner`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *TransactionAddressDestinationUtxoOutputs) GetOutputsOk() (*[]TransactionAddressDestinationUtxoOutputsOutputsInner, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *TransactionAddressDestinationUtxoOutputs) SetOutputs(v []TransactionAddressDestinationUtxoOutputsOutputsInner)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *TransactionAddressDestinationUtxoOutputs) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetChangeAddress

`func (o *TransactionAddressDestinationUtxoOutputs) GetChangeAddress() string`

GetChangeAddress returns the ChangeAddress field if non-nil, zero value otherwise.

### GetChangeAddressOk

`func (o *TransactionAddressDestinationUtxoOutputs) GetChangeAddressOk() (*string, bool)`

GetChangeAddressOk returns a tuple with the ChangeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAddress

`func (o *TransactionAddressDestinationUtxoOutputs) SetChangeAddress(v string)`

SetChangeAddress sets ChangeAddress field to given value.

### HasChangeAddress

`func (o *TransactionAddressDestinationUtxoOutputs) HasChangeAddress() bool`

HasChangeAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


