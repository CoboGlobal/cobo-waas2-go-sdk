# AddressTransferDestinationUtxoOutputsOutputsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The external address. | [optional] 
**Amount** | Pointer to **string** | The quantity of the token in the transaction. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | [optional] 
**Script** | Pointer to **string** | The script of the output. It is a programmable code fragment that defines the conditions under which the UTXO can be spent. | [optional] 

## Methods

### NewAddressTransferDestinationUtxoOutputsOutputsInner

`func NewAddressTransferDestinationUtxoOutputsOutputsInner() *AddressTransferDestinationUtxoOutputsOutputsInner`

NewAddressTransferDestinationUtxoOutputsOutputsInner instantiates a new AddressTransferDestinationUtxoOutputsOutputsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransferDestinationUtxoOutputsOutputsInnerWithDefaults

`func NewAddressTransferDestinationUtxoOutputsOutputsInnerWithDefaults() *AddressTransferDestinationUtxoOutputsOutputsInner`

NewAddressTransferDestinationUtxoOutputsOutputsInnerWithDefaults instantiates a new AddressTransferDestinationUtxoOutputsOutputsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddressTransferDestinationUtxoOutputsOutputsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressTransferDestinationUtxoOutputsOutputsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressTransferDestinationUtxoOutputsOutputsInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddressTransferDestinationUtxoOutputsOutputsInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAmount

`func (o *AddressTransferDestinationUtxoOutputsOutputsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressTransferDestinationUtxoOutputsOutputsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressTransferDestinationUtxoOutputsOutputsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AddressTransferDestinationUtxoOutputsOutputsInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetScript

`func (o *AddressTransferDestinationUtxoOutputsOutputsInner) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *AddressTransferDestinationUtxoOutputsOutputsInner) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *AddressTransferDestinationUtxoOutputsOutputsInner) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *AddressTransferDestinationUtxoOutputsOutputsInner) HasScript() bool`

HasScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


