# TransactionTransferToAddressDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationType** | [**TransactionDestinationType**](TransactionDestinationType.md) |  | 
**AccountOutput** | Pointer to [**TransactionTransferToAddressDestinationAccountOutput**](TransactionTransferToAddressDestinationAccountOutput.md) |  | [optional] 
**UtxoOutputs** | Pointer to [**[]TransactionTransferToAddressDestinationUtxoOutputsInner**](TransactionTransferToAddressDestinationUtxoOutputsInner.md) |  | [optional] 
**ChangeAddress** | Pointer to **string** | The address used to receive the remaining funds or change from the transaction. | [optional] 
**ForceInternal** | Pointer to **bool** | Whether the transaction request must be executed as a [Cobo Loop](https://manuals.cobo.com/en/portal/custodial-wallets/cobo-loop) transfer.   - &#x60;true&#x60;: The transaction request must be executed as a Cobo Loop transfer.   - &#x60;false&#x60;: The transaction request may not be executed as a Cobo Loop transfer.  | [optional] 
**ForceExternal** | Pointer to **bool** | Whether the transaction request must not be executed as a [Cobo Loop](https://manuals.cobo.com/en/portal/custodial-wallets/cobo-loop) transfer.   - &#x60;true&#x60;: The transaction request must not be executed as a Cobo Loop transfer.   - &#x60;false&#x60;: The transaction request can be executed as a Cobo Loop transfer.  | [optional] 

## Methods

### NewTransactionTransferToAddressDestination

`func NewTransactionTransferToAddressDestination(destinationType TransactionDestinationType, ) *TransactionTransferToAddressDestination`

NewTransactionTransferToAddressDestination instantiates a new TransactionTransferToAddressDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTransferToAddressDestinationWithDefaults

`func NewTransactionTransferToAddressDestinationWithDefaults() *TransactionTransferToAddressDestination`

NewTransactionTransferToAddressDestinationWithDefaults instantiates a new TransactionTransferToAddressDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationType

`func (o *TransactionTransferToAddressDestination) GetDestinationType() TransactionDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionTransferToAddressDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionTransferToAddressDestination) SetDestinationType(v TransactionDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetAccountOutput

`func (o *TransactionTransferToAddressDestination) GetAccountOutput() TransactionTransferToAddressDestinationAccountOutput`

GetAccountOutput returns the AccountOutput field if non-nil, zero value otherwise.

### GetAccountOutputOk

`func (o *TransactionTransferToAddressDestination) GetAccountOutputOk() (*TransactionTransferToAddressDestinationAccountOutput, bool)`

GetAccountOutputOk returns a tuple with the AccountOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOutput

`func (o *TransactionTransferToAddressDestination) SetAccountOutput(v TransactionTransferToAddressDestinationAccountOutput)`

SetAccountOutput sets AccountOutput field to given value.

### HasAccountOutput

`func (o *TransactionTransferToAddressDestination) HasAccountOutput() bool`

HasAccountOutput returns a boolean if a field has been set.

### GetUtxoOutputs

`func (o *TransactionTransferToAddressDestination) GetUtxoOutputs() []TransactionTransferToAddressDestinationUtxoOutputsInner`

GetUtxoOutputs returns the UtxoOutputs field if non-nil, zero value otherwise.

### GetUtxoOutputsOk

`func (o *TransactionTransferToAddressDestination) GetUtxoOutputsOk() (*[]TransactionTransferToAddressDestinationUtxoOutputsInner, bool)`

GetUtxoOutputsOk returns a tuple with the UtxoOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoOutputs

`func (o *TransactionTransferToAddressDestination) SetUtxoOutputs(v []TransactionTransferToAddressDestinationUtxoOutputsInner)`

SetUtxoOutputs sets UtxoOutputs field to given value.

### HasUtxoOutputs

`func (o *TransactionTransferToAddressDestination) HasUtxoOutputs() bool`

HasUtxoOutputs returns a boolean if a field has been set.

### GetChangeAddress

`func (o *TransactionTransferToAddressDestination) GetChangeAddress() string`

GetChangeAddress returns the ChangeAddress field if non-nil, zero value otherwise.

### GetChangeAddressOk

`func (o *TransactionTransferToAddressDestination) GetChangeAddressOk() (*string, bool)`

GetChangeAddressOk returns a tuple with the ChangeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeAddress

`func (o *TransactionTransferToAddressDestination) SetChangeAddress(v string)`

SetChangeAddress sets ChangeAddress field to given value.

### HasChangeAddress

`func (o *TransactionTransferToAddressDestination) HasChangeAddress() bool`

HasChangeAddress returns a boolean if a field has been set.

### GetForceInternal

`func (o *TransactionTransferToAddressDestination) GetForceInternal() bool`

GetForceInternal returns the ForceInternal field if non-nil, zero value otherwise.

### GetForceInternalOk

`func (o *TransactionTransferToAddressDestination) GetForceInternalOk() (*bool, bool)`

GetForceInternalOk returns a tuple with the ForceInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceInternal

`func (o *TransactionTransferToAddressDestination) SetForceInternal(v bool)`

SetForceInternal sets ForceInternal field to given value.

### HasForceInternal

`func (o *TransactionTransferToAddressDestination) HasForceInternal() bool`

HasForceInternal returns a boolean if a field has been set.

### GetForceExternal

`func (o *TransactionTransferToAddressDestination) GetForceExternal() bool`

GetForceExternal returns the ForceExternal field if non-nil, zero value otherwise.

### GetForceExternalOk

`func (o *TransactionTransferToAddressDestination) GetForceExternalOk() (*bool, bool)`

GetForceExternalOk returns a tuple with the ForceExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceExternal

`func (o *TransactionTransferToAddressDestination) SetForceExternal(v bool)`

SetForceExternal sets ForceExternal field to given value.

### HasForceExternal

`func (o *TransactionTransferToAddressDestination) HasForceExternal() bool`

HasForceExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


