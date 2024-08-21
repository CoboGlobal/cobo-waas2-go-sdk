# AddressTransferDestinationUtxoOutputsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The destination address. | 
**Amount** | Pointer to **string** | The transfer amount. For example, if you trade 1.5 ETH, then the value is &#x60;1.5&#x60;.  | [optional] 
**Script** | Pointer to **string** | The script of the output. It is a programmable code fragment that defines the conditions under which the UTXO can be spent. | [optional] 

## Methods

### NewAddressTransferDestinationUtxoOutputsInner

`func NewAddressTransferDestinationUtxoOutputsInner(address string, ) *AddressTransferDestinationUtxoOutputsInner`

NewAddressTransferDestinationUtxoOutputsInner instantiates a new AddressTransferDestinationUtxoOutputsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransferDestinationUtxoOutputsInnerWithDefaults

`func NewAddressTransferDestinationUtxoOutputsInnerWithDefaults() *AddressTransferDestinationUtxoOutputsInner`

NewAddressTransferDestinationUtxoOutputsInnerWithDefaults instantiates a new AddressTransferDestinationUtxoOutputsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddressTransferDestinationUtxoOutputsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressTransferDestinationUtxoOutputsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressTransferDestinationUtxoOutputsInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *AddressTransferDestinationUtxoOutputsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressTransferDestinationUtxoOutputsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressTransferDestinationUtxoOutputsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AddressTransferDestinationUtxoOutputsInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetScript

`func (o *AddressTransferDestinationUtxoOutputsInner) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *AddressTransferDestinationUtxoOutputsInner) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *AddressTransferDestinationUtxoOutputsInner) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *AddressTransferDestinationUtxoOutputsInner) HasScript() bool`

HasScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


