# LockUtxosRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Utxos** | [**[]LockUtxosRequestUtxosInner**](LockUtxosRequestUtxosInner.md) |  | 

## Methods

### NewLockUtxosRequest

`func NewLockUtxosRequest(utxos []LockUtxosRequestUtxosInner, ) *LockUtxosRequest`

NewLockUtxosRequest instantiates a new LockUtxosRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockUtxosRequestWithDefaults

`func NewLockUtxosRequestWithDefaults() *LockUtxosRequest`

NewLockUtxosRequestWithDefaults instantiates a new LockUtxosRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUtxos

`func (o *LockUtxosRequest) GetUtxos() []LockUtxosRequestUtxosInner`

GetUtxos returns the Utxos field if non-nil, zero value otherwise.

### GetUtxosOk

`func (o *LockUtxosRequest) GetUtxosOk() (*[]LockUtxosRequestUtxosInner, bool)`

GetUtxosOk returns a tuple with the Utxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxos

`func (o *LockUtxosRequest) SetUtxos(v []LockUtxosRequestUtxosInner)`

SetUtxos sets Utxos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


