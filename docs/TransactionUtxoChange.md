# TransactionUtxoChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The receiving address of the UTXO change output. | [optional] 
**Value** | Pointer to **string** | The amount of the UTXO change output. | [optional] 

## Methods

### NewTransactionUtxoChange

`func NewTransactionUtxoChange() *TransactionUtxoChange`

NewTransactionUtxoChange instantiates a new TransactionUtxoChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionUtxoChangeWithDefaults

`func NewTransactionUtxoChangeWithDefaults() *TransactionUtxoChange`

NewTransactionUtxoChangeWithDefaults instantiates a new TransactionUtxoChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TransactionUtxoChange) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionUtxoChange) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionUtxoChange) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransactionUtxoChange) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetValue

`func (o *TransactionUtxoChange) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TransactionUtxoChange) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TransactionUtxoChange) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TransactionUtxoChange) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


