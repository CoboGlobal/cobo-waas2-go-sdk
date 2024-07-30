# TransactionRawTxInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedNonce** | Pointer to **int32** | The transaction nonce. | [optional] 
**SelectedUtxos** | Pointer to [**[]TransactionUtxo**](TransactionUtxo.md) | The selected UTXOs to be consumed in the transaction. | [optional] 
**RawTx** | Pointer to **string** | The raw transaction data. | [optional] 

## Methods

### NewTransactionRawTxInfo

`func NewTransactionRawTxInfo() *TransactionRawTxInfo`

NewTransactionRawTxInfo instantiates a new TransactionRawTxInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRawTxInfoWithDefaults

`func NewTransactionRawTxInfoWithDefaults() *TransactionRawTxInfo`

NewTransactionRawTxInfoWithDefaults instantiates a new TransactionRawTxInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsedNonce

`func (o *TransactionRawTxInfo) GetUsedNonce() int32`

GetUsedNonce returns the UsedNonce field if non-nil, zero value otherwise.

### GetUsedNonceOk

`func (o *TransactionRawTxInfo) GetUsedNonceOk() (*int32, bool)`

GetUsedNonceOk returns a tuple with the UsedNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedNonce

`func (o *TransactionRawTxInfo) SetUsedNonce(v int32)`

SetUsedNonce sets UsedNonce field to given value.

### HasUsedNonce

`func (o *TransactionRawTxInfo) HasUsedNonce() bool`

HasUsedNonce returns a boolean if a field has been set.

### GetSelectedUtxos

`func (o *TransactionRawTxInfo) GetSelectedUtxos() []TransactionUtxo`

GetSelectedUtxos returns the SelectedUtxos field if non-nil, zero value otherwise.

### GetSelectedUtxosOk

`func (o *TransactionRawTxInfo) GetSelectedUtxosOk() (*[]TransactionUtxo, bool)`

GetSelectedUtxosOk returns a tuple with the SelectedUtxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedUtxos

`func (o *TransactionRawTxInfo) SetSelectedUtxos(v []TransactionUtxo)`

SetSelectedUtxos sets SelectedUtxos field to given value.

### HasSelectedUtxos

`func (o *TransactionRawTxInfo) HasSelectedUtxos() bool`

HasSelectedUtxos returns a boolean if a field has been set.

### GetRawTx

`func (o *TransactionRawTxInfo) GetRawTx() string`

GetRawTx returns the RawTx field if non-nil, zero value otherwise.

### GetRawTxOk

`func (o *TransactionRawTxInfo) GetRawTxOk() (*string, bool)`

GetRawTxOk returns a tuple with the RawTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTx

`func (o *TransactionRawTxInfo) SetRawTx(v string)`

SetRawTx sets RawTx field to given value.

### HasRawTx

`func (o *TransactionRawTxInfo) HasRawTx() bool`

HasRawTx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


