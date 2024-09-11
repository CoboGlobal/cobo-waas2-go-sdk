# WalletBalanceSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | **int32** | The snapshot ID. | 
**SnapshotName** | Pointer to **string** | The snapshot name. | [optional] 

## Methods

### NewWalletBalanceSnapshot

`func NewWalletBalanceSnapshot(snapshotId int32, ) *WalletBalanceSnapshot`

NewWalletBalanceSnapshot instantiates a new WalletBalanceSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletBalanceSnapshotWithDefaults

`func NewWalletBalanceSnapshotWithDefaults() *WalletBalanceSnapshot`

NewWalletBalanceSnapshotWithDefaults instantiates a new WalletBalanceSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *WalletBalanceSnapshot) GetSnapshotId() int32`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *WalletBalanceSnapshot) GetSnapshotIdOk() (*int32, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *WalletBalanceSnapshot) SetSnapshotId(v int32)`

SetSnapshotId sets SnapshotId field to given value.


### GetSnapshotName

`func (o *WalletBalanceSnapshot) GetSnapshotName() string`

GetSnapshotName returns the SnapshotName field if non-nil, zero value otherwise.

### GetSnapshotNameOk

`func (o *WalletBalanceSnapshot) GetSnapshotNameOk() (*string, bool)`

GetSnapshotNameOk returns a tuple with the SnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotName

`func (o *WalletBalanceSnapshot) SetSnapshotName(v string)`

SetSnapshotName sets SnapshotName field to given value.

### HasSnapshotName

`func (o *WalletBalanceSnapshot) HasSnapshotName() bool`

HasSnapshotName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


