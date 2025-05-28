# UTXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxHash** | Pointer to **string** | The transaction hash of the UTXO. | [optional] 
**VoutN** | Pointer to **int32** | The output index of the UTXO. | [optional] 
**Address** | Pointer to **string** | The address of the UTXO. | [optional] 
**TokenId** | Pointer to **string** | The token ID, which is the unique identifier of a token. | [optional] 
**Value** | Pointer to **string** | The value of the UTXO. | [optional] 
**IsCoinbase** | Pointer to **bool** | Whether the UTXO comes from a coinbase transaction. | [optional] 
**IsLocked** | Pointer to **bool** | Whether the UTXO is locked. | [optional] 
**ConfirmedNumber** | Pointer to **int32** | The number of confirmations for the UTXO. | [optional] 
**IsFrozen** | Pointer to **bool** | Whether the UTXO is frozen. | [optional] 

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

### GetTokenId

`func (o *UTXO) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *UTXO) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *UTXO) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *UTXO) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetValue

`func (o *UTXO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UTXO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UTXO) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UTXO) HasValue() bool`

HasValue returns a boolean if a field has been set.

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

### GetIsLocked

`func (o *UTXO) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *UTXO) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *UTXO) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *UTXO) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

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

### GetIsFrozen

`func (o *UTXO) GetIsFrozen() bool`

GetIsFrozen returns the IsFrozen field if non-nil, zero value otherwise.

### GetIsFrozenOk

`func (o *UTXO) GetIsFrozenOk() (*bool, bool)`

GetIsFrozenOk returns a tuple with the IsFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFrozen

`func (o *UTXO) SetIsFrozen(v bool)`

SetIsFrozen sets IsFrozen field to given value.

### HasIsFrozen

`func (o *UTXO) HasIsFrozen() bool`

HasIsFrozen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


