# AddressTransferDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransferDestinationType**](TransferDestinationType.md) |  | 
**AccountOutput** | Pointer to [**AddressTransferDestinationAccountOutput**](AddressTransferDestinationAccountOutput.md) |  | [optional] 
**UtxoOutputs** | Pointer to [**AddressTransferDestinationUtxoOutputs**](AddressTransferDestinationUtxoOutputs.md) |  | [optional] 
**ForceInternal** | Pointer to **bool** | Whether the transfer request must be executed as an off-chain transfer.    - &#x60;true&#x60;: The transfer request must be executed an off-chain transfer.    - &#x60;false&#x60;: The transfer may not be executed as an off-chain transfer.  | [optional] 
**ForceExternal** | Pointer to **bool** | Whether the transfer must be executed as an on-chain transfer. - &#x60;true&#x60;: The transfer must be executed as an on-chain transfer. - &#x60;false&#x60;: The transfer may not be executed as an on-chain transfer.  | [optional] 

## Methods

### NewAddressTransferDestination

`func NewAddressTransferDestination(destinationType TransferDestinationType, ) *AddressTransferDestination`

NewAddressTransferDestination instantiates a new AddressTransferDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransferDestinationWithDefaults

`func NewAddressTransferDestinationWithDefaults() *AddressTransferDestination`

NewAddressTransferDestinationWithDefaults instantiates a new AddressTransferDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *AddressTransferDestination) GetDestinationType() TransferDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *AddressTransferDestination) GetDestinationTypeOk() (*TransferDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *AddressTransferDestination) SetDestinationType(v TransferDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetAccountOutput

`func (o *AddressTransferDestination) GetAccountOutput() AddressTransferDestinationAccountOutput`

GetAccountOutput returns the AccountOutput field if non-nil, zero value otherwise.

### GetAccountOutputOk

`func (o *AddressTransferDestination) GetAccountOutputOk() (*AddressTransferDestinationAccountOutput, bool)`

GetAccountOutputOk returns a tuple with the AccountOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOutput

`func (o *AddressTransferDestination) SetAccountOutput(v AddressTransferDestinationAccountOutput)`

SetAccountOutput sets AccountOutput field to given value.

### HasAccountOutput

`func (o *AddressTransferDestination) HasAccountOutput() bool`

HasAccountOutput returns a boolean if a field has been set.

### GetUtxoOutputs

`func (o *AddressTransferDestination) GetUtxoOutputs() AddressTransferDestinationUtxoOutputs`

GetUtxoOutputs returns the UtxoOutputs field if non-nil, zero value otherwise.

### GetUtxoOutputsOk

`func (o *AddressTransferDestination) GetUtxoOutputsOk() (*AddressTransferDestinationUtxoOutputs, bool)`

GetUtxoOutputsOk returns a tuple with the UtxoOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoOutputs

`func (o *AddressTransferDestination) SetUtxoOutputs(v AddressTransferDestinationUtxoOutputs)`

SetUtxoOutputs sets UtxoOutputs field to given value.

### HasUtxoOutputs

`func (o *AddressTransferDestination) HasUtxoOutputs() bool`

HasUtxoOutputs returns a boolean if a field has been set.

### GetForceInternal

`func (o *AddressTransferDestination) GetForceInternal() bool`

GetForceInternal returns the ForceInternal field if non-nil, zero value otherwise.

### GetForceInternalOk

`func (o *AddressTransferDestination) GetForceInternalOk() (*bool, bool)`

GetForceInternalOk returns a tuple with the ForceInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceInternal

`func (o *AddressTransferDestination) SetForceInternal(v bool)`

SetForceInternal sets ForceInternal field to given value.

### HasForceInternal

`func (o *AddressTransferDestination) HasForceInternal() bool`

HasForceInternal returns a boolean if a field has been set.

### GetForceExternal

`func (o *AddressTransferDestination) GetForceExternal() bool`

GetForceExternal returns the ForceExternal field if non-nil, zero value otherwise.

### GetForceExternalOk

`func (o *AddressTransferDestination) GetForceExternalOk() (*bool, bool)`

GetForceExternalOk returns a tuple with the ForceExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExternal

`func (o *AddressTransferDestination) SetForceExternal(v bool)`

SetForceExternal sets ForceExternal field to given value.

### HasForceExternal

`func (o *AddressTransferDestination) HasForceExternal() bool`

HasForceExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


