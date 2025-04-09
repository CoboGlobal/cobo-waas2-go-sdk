# TransactionSelectedUtxo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxHash** | Pointer to **string** | The transaction hash of the UTXO. | [optional] 
**VoutN** | Pointer to **int32** | The output index of the UTXO. | [optional] 
**Address** | Pointer to **string** | The address of the UTXO. | [optional] 
**Value** | Pointer to **string** | The value of the UTXO. | [optional] 
**RedeemScript** | Pointer to **string** | The redeem script used in P2SH and P2WSH transactions. | [optional] 
**RevealedScript** | Pointer to **string** | The revealed script used for Taproot script-path spend transaction. | [optional] 

## Methods

### NewTransactionSelectedUtxo

`func NewTransactionSelectedUtxo() *TransactionSelectedUtxo`

NewTransactionSelectedUtxo instantiates a new TransactionSelectedUtxo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSelectedUtxoWithDefaults

`func NewTransactionSelectedUtxoWithDefaults() *TransactionSelectedUtxo`

NewTransactionSelectedUtxoWithDefaults instantiates a new TransactionSelectedUtxo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxHash

`func (o *TransactionSelectedUtxo) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *TransactionSelectedUtxo) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *TransactionSelectedUtxo) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *TransactionSelectedUtxo) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetVoutN

`func (o *TransactionSelectedUtxo) GetVoutN() int32`

GetVoutN returns the VoutN field if non-nil, zero value otherwise.

### GetVoutNOk

`func (o *TransactionSelectedUtxo) GetVoutNOk() (*int32, bool)`

GetVoutNOk returns a tuple with the VoutN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoutN

`func (o *TransactionSelectedUtxo) SetVoutN(v int32)`

SetVoutN sets VoutN field to given value.

### HasVoutN

`func (o *TransactionSelectedUtxo) HasVoutN() bool`

HasVoutN returns a boolean if a field has been set.

### GetAddress

`func (o *TransactionSelectedUtxo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionSelectedUtxo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionSelectedUtxo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransactionSelectedUtxo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetValue

`func (o *TransactionSelectedUtxo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TransactionSelectedUtxo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TransactionSelectedUtxo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TransactionSelectedUtxo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRedeemScript

`func (o *TransactionSelectedUtxo) GetRedeemScript() string`

GetRedeemScript returns the RedeemScript field if non-nil, zero value otherwise.

### GetRedeemScriptOk

`func (o *TransactionSelectedUtxo) GetRedeemScriptOk() (*string, bool)`

GetRedeemScriptOk returns a tuple with the RedeemScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemScript

`func (o *TransactionSelectedUtxo) SetRedeemScript(v string)`

SetRedeemScript sets RedeemScript field to given value.

### HasRedeemScript

`func (o *TransactionSelectedUtxo) HasRedeemScript() bool`

HasRedeemScript returns a boolean if a field has been set.

### GetRevealedScript

`func (o *TransactionSelectedUtxo) GetRevealedScript() string`

GetRevealedScript returns the RevealedScript field if non-nil, zero value otherwise.

### GetRevealedScriptOk

`func (o *TransactionSelectedUtxo) GetRevealedScriptOk() (*string, bool)`

GetRevealedScriptOk returns a tuple with the RevealedScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevealedScript

`func (o *TransactionSelectedUtxo) SetRevealedScript(v string)`

SetRevealedScript sets RevealedScript field to given value.

### HasRevealedScript

`func (o *TransactionSelectedUtxo) HasRevealedScript() bool`

HasRevealedScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


