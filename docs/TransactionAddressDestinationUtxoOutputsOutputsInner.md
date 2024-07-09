# TransactionAddressDestinationUtxoOutputsOutputsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The destination address. | [optional] 
**Amount** | Pointer to **string** | The transaction amount. For example, if you trade 1.5 ETH, then the amount is &#x60;1.5&#x60;.  | [optional] 
**Script** | Pointer to **string** | The script of the output. It is a programmable code fragment that defines the conditions under which the UTXO can be spent. | [optional] 

## Methods

### NewTransactionAddressDestinationUtxoOutputsOutputsInner

`func NewTransactionAddressDestinationUtxoOutputsOutputsInner() *TransactionAddressDestinationUtxoOutputsOutputsInner`

NewTransactionAddressDestinationUtxoOutputsOutputsInner instantiates a new TransactionAddressDestinationUtxoOutputsOutputsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionAddressDestinationUtxoOutputsOutputsInnerWithDefaults

`func NewTransactionAddressDestinationUtxoOutputsOutputsInnerWithDefaults() *TransactionAddressDestinationUtxoOutputsOutputsInner`

NewTransactionAddressDestinationUtxoOutputsOutputsInnerWithDefaults instantiates a new TransactionAddressDestinationUtxoOutputsOutputsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetScript

`func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) HasScript() bool`

HasScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


