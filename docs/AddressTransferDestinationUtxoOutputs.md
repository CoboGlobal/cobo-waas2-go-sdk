# AddressTransferDestinationUtxoOutputs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outputs** | Pointer to [**[]AddressTransferDestinationUtxoOutputsOutputsInner**](AddressTransferDestinationUtxoOutputsOutputsInner.md) |  | [optional] 
**ChangeAddress** | Pointer to **string** | Change address | [optional] 

## Methods

### NewAddressTransferDestinationUtxoOutputs

`func NewAddressTransferDestinationUtxoOutputs() *AddressTransferDestinationUtxoOutputs`

NewAddressTransferDestinationUtxoOutputs instantiates a new AddressTransferDestinationUtxoOutputs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransferDestinationUtxoOutputsWithDefaults

`func NewAddressTransferDestinationUtxoOutputsWithDefaults() *AddressTransferDestinationUtxoOutputs`

NewAddressTransferDestinationUtxoOutputsWithDefaults instantiates a new AddressTransferDestinationUtxoOutputs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutputs

`func (o *AddressTransferDestinationUtxoOutputs) GetOutputs() []AddressTransferDestinationUtxoOutputsOutputsInner`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *AddressTransferDestinationUtxoOutputs) GetOutputsOk() (*[]AddressTransferDestinationUtxoOutputsOutputsInner, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *AddressTransferDestinationUtxoOutputs) SetOutputs(v []AddressTransferDestinationUtxoOutputsOutputsInner)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *AddressTransferDestinationUtxoOutputs) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetChangeAddress

`func (o *AddressTransferDestinationUtxoOutputs) GetChangeAddress() string`

GetChangeAddress returns the ChangeAddress field if non-nil, zero value otherwise.

### GetChangeAddressOk

`func (o *AddressTransferDestinationUtxoOutputs) GetChangeAddressOk() (*string, bool)`

GetChangeAddressOk returns a tuple with the ChangeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAddress

`func (o *AddressTransferDestinationUtxoOutputs) SetChangeAddress(v string)`

SetChangeAddress sets ChangeAddress field to given value.

### HasChangeAddress

`func (o *AddressTransferDestinationUtxoOutputs) HasChangeAddress() bool`

HasChangeAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


