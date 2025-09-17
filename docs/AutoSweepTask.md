# AutoSweepTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Auto-sweep task ID. | 
**WalletId** | **string** | Wallet ID. | 
**TokenId** | **string** | Token ID of the swept token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
**Status** | [**AutoSweepTaskStatus**](AutoSweepTaskStatus.md) |  | 
**TransactionIds** | Pointer to **[]string** | IDs of the transactions triggered by the task. | [optional] 
**FailedReasons** | Pointer to **[]string** | Reasons why the task creation failed. | [optional] 
**CreatedTimestamp** | **int64** | The time when the task was created, in Unix timestamp format, measured in milliseconds. | 
**UpdatedTimestamp** | Pointer to **int64** | The time when the task was updated, in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewAutoSweepTask

`func NewAutoSweepTask(taskId string, walletId string, tokenId string, status AutoSweepTaskStatus, createdTimestamp int64, ) *AutoSweepTask`

NewAutoSweepTask instantiates a new AutoSweepTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoSweepTaskWithDefaults

`func NewAutoSweepTaskWithDefaults() *AutoSweepTask`

NewAutoSweepTaskWithDefaults instantiates a new AutoSweepTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *AutoSweepTask) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *AutoSweepTask) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *AutoSweepTask) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetWalletId

`func (o *AutoSweepTask) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *AutoSweepTask) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *AutoSweepTask) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetTokenId

`func (o *AutoSweepTask) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *AutoSweepTask) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *AutoSweepTask) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetStatus

`func (o *AutoSweepTask) GetStatus() AutoSweepTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutoSweepTask) GetStatusOk() (*AutoSweepTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutoSweepTask) SetStatus(v AutoSweepTaskStatus)`

SetStatus sets Status field to given value.


### GetTransactionIds

`func (o *AutoSweepTask) GetTransactionIds() []string`

GetTransactionIds returns the TransactionIds field if non-nil, zero value otherwise.

### GetTransactionIdsOk

`func (o *AutoSweepTask) GetTransactionIdsOk() (*[]string, bool)`

GetTransactionIdsOk returns a tuple with the TransactionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIds

`func (o *AutoSweepTask) SetTransactionIds(v []string)`

SetTransactionIds sets TransactionIds field to given value.

### HasTransactionIds

`func (o *AutoSweepTask) HasTransactionIds() bool`

HasTransactionIds returns a boolean if a field has been set.

### GetFailedReasons

`func (o *AutoSweepTask) GetFailedReasons() []string`

GetFailedReasons returns the FailedReasons field if non-nil, zero value otherwise.

### GetFailedReasonsOk

`func (o *AutoSweepTask) GetFailedReasonsOk() (*[]string, bool)`

GetFailedReasonsOk returns a tuple with the FailedReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReasons

`func (o *AutoSweepTask) SetFailedReasons(v []string)`

SetFailedReasons sets FailedReasons field to given value.

### HasFailedReasons

`func (o *AutoSweepTask) HasFailedReasons() bool`

HasFailedReasons returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *AutoSweepTask) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *AutoSweepTask) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *AutoSweepTask) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *AutoSweepTask) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *AutoSweepTask) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *AutoSweepTask) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *AutoSweepTask) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


