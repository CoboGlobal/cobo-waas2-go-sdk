# UTXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxHash** | Pointer to **string** | Transaction hash of the UTXO. | [optional] 
**VoutN** | Pointer to **int32** | Output index of the UTXO. | [optional] 
**Address** | Pointer to **string** | Address of the UTXO. | [optional] 
**Amount** | Pointer to **string** | UTXO amount in decimal places (e.g. one bitcoin is divisible to eight decimal places, and 100000000 represents 1 BTC). | [optional] 
**IsCoinbase** | Pointer to **bool** | Whether the UTXO is a coinbase transaction. | [optional] 
**ConfirmedNumber** | Pointer to **int32** | Number of confirmations for the UTXO. | [optional] 

## Methods

### NewUTXO

`func NewUTXO() *UTXO`

NewUTXO instantiates a new UTXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUTXOWithDefaults

`func NewUTXOWithDefaults() *UTXO`

NewUTXOWithDefaults instantiates a new UTXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxHash

`func (o *UTXO) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *UTXO) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *UTXO) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *UTXO) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetVoutN

`func (o *UTXO) GetVoutN() int32`

GetVoutN returns the VoutN field if non-nil, zero value otherwise.

### GetVoutNOk

`func (o *UTXO) GetVoutNOk() (*int32, bool)`

GetVoutNOk returns a tuple with the VoutN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoutN

`func (o *UTXO) SetVoutN(v int32)`

SetVoutN sets VoutN field to given value.

### HasVoutN

`func (o *UTXO) HasVoutN() bool`

HasVoutN returns a boolean if a field has been set.

### GetAddress

`func (o *UTXO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UTXO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UTXO) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UTXO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAmount

`func (o *UTXO) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UTXO) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UTXO) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UTXO) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetIsCoinbase

`func (o *UTXO) GetIsCoinbase() bool`

GetIsCoinbase returns the IsCoinbase field if non-nil, zero value otherwise.

### GetIsCoinbaseOk

`func (o *UTXO) GetIsCoinbaseOk() (*bool, bool)`

GetIsCoinbaseOk returns a tuple with the IsCoinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCoinbase

`func (o *UTXO) SetIsCoinbase(v bool)`

SetIsCoinbase sets IsCoinbase field to given value.

### HasIsCoinbase

`func (o *UTXO) HasIsCoinbase() bool`

HasIsCoinbase returns a boolean if a field has been set.

### GetConfirmedNumber

`func (o *UTXO) GetConfirmedNumber() int32`

GetConfirmedNumber returns the ConfirmedNumber field if non-nil, zero value otherwise.

### GetConfirmedNumberOk

`func (o *UTXO) GetConfirmedNumberOk() (*int32, bool)`

GetConfirmedNumberOk returns a tuple with the ConfirmedNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedNumber

`func (o *UTXO) SetConfirmedNumber(v int32)`

SetConfirmedNumber sets ConfirmedNumber field to given value.

### HasConfirmedNumber

`func (o *UTXO) HasConfirmedNumber() bool`

HasConfirmedNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


