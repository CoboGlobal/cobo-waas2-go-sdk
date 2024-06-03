# LockSpendableListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxHashes** | Pointer to **[]string** |  | [optional] 
**IsLocked** | Pointer to **bool** | True if to lock the UTXOs, False to unlock. | [optional] 

## Methods

### NewLockSpendableListRequest

`func NewLockSpendableListRequest() *LockSpendableListRequest`

NewLockSpendableListRequest instantiates a new LockSpendableListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockSpendableListRequestWithDefaults

`func NewLockSpendableListRequestWithDefaults() *LockSpendableListRequest`

NewLockSpendableListRequestWithDefaults instantiates a new LockSpendableListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxHashes

`func (o *LockSpendableListRequest) GetTxHashes() []string`

GetTxHashes returns the TxHashes field if non-nil, zero value otherwise.

### GetTxHashesOk

`func (o *LockSpendableListRequest) GetTxHashesOk() (*[]string, bool)`

GetTxHashesOk returns a tuple with the TxHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHashes

`func (o *LockSpendableListRequest) SetTxHashes(v []string)`

SetTxHashes sets TxHashes field to given value.

### HasTxHashes

`func (o *LockSpendableListRequest) HasTxHashes() bool`

HasTxHashes returns a boolean if a field has been set.

### GetIsLocked

`func (o *LockSpendableListRequest) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *LockSpendableListRequest) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *LockSpendableListRequest) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *LockSpendableListRequest) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


