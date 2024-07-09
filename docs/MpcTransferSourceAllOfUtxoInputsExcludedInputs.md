# MpcTransferSourceAllOfUtxoInputsExcludedInputs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxHash** | Pointer to **string** | The transaction hash of UTXO not to be consumed in the transaction. | [optional] 
**VoutN** | Pointer to **int32** | The output index of the UTXO. | [optional] 

## Methods

### NewMpcTransferSourceAllOfUtxoInputsExcludedInputs

`func NewMpcTransferSourceAllOfUtxoInputsExcludedInputs() *MpcTransferSourceAllOfUtxoInputsExcludedInputs`

NewMpcTransferSourceAllOfUtxoInputsExcludedInputs instantiates a new MpcTransferSourceAllOfUtxoInputsExcludedInputs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpcTransferSourceAllOfUtxoInputsExcludedInputsWithDefaults

`func NewMpcTransferSourceAllOfUtxoInputsExcludedInputsWithDefaults() *MpcTransferSourceAllOfUtxoInputsExcludedInputs`

NewMpcTransferSourceAllOfUtxoInputsExcludedInputsWithDefaults instantiates a new MpcTransferSourceAllOfUtxoInputsExcludedInputs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxHash

`func (o *MpcTransferSourceAllOfUtxoInputsExcludedInputs) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *MpcTransferSourceAllOfUtxoInputsExcludedInputs) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *MpcTransferSourceAllOfUtxoInputsExcludedInputs) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *MpcTransferSourceAllOfUtxoInputsExcludedInputs) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetVoutN

`func (o *MpcTransferSourceAllOfUtxoInputsExcludedInputs) GetVoutN() int32`

GetVoutN returns the VoutN field if non-nil, zero value otherwise.

### GetVoutNOk

`func (o *MpcTransferSourceAllOfUtxoInputsExcludedInputs) GetVoutNOk() (*int32, bool)`

GetVoutNOk returns a tuple with the VoutN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoutN

`func (o *MpcTransferSourceAllOfUtxoInputsExcludedInputs) SetVoutN(v int32)`

SetVoutN sets VoutN field to given value.

### HasVoutN

`func (o *MpcTransferSourceAllOfUtxoInputsExcludedInputs) HasVoutN() bool`

HasVoutN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


