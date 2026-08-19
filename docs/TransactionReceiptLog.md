# TransactionReceiptLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogIndex** | **int64** | The index position of the log within the block. | 
**Address** | **string** | The address of the contract that emitted the log. | 
**Topics** | **[]string** | The indexed log arguments. The first topic is the hash of the event signature, and the remaining topics are the indexed event parameters, with a maximum of three. | 
**Data** | **string** | The non-indexed log arguments, encoded as a hexadecimal string. | 
**BlockNumber** | Pointer to **int64** | The number of the block that contains the log. | [optional] 
**BlockHash** | Pointer to **string** | The hash of the block that contains the log. | [optional] 
**TransactionHash** | Pointer to **string** | The hash of the transaction that emitted the log. | [optional] 
**TransactionIndex** | Pointer to **int64** | The index position within the block of the transaction that emitted the log. | [optional] 
**Removed** | Pointer to **bool** | Whether the log was removed due to a chain reorganization. - &#x60;true&#x60;: The log was removed because the block that contains it was reorganized out of the canonical chain. - &#x60;false&#x60;: The log is still valid.  | [optional] 

## Methods

### NewTransactionReceiptLog

`func NewTransactionReceiptLog(logIndex int64, address string, topics []string, data string, ) *TransactionReceiptLog`

NewTransactionReceiptLog instantiates a new TransactionReceiptLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionReceiptLogWithDefaults

`func NewTransactionReceiptLogWithDefaults() *TransactionReceiptLog`

NewTransactionReceiptLogWithDefaults instantiates a new TransactionReceiptLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogIndex

`func (o *TransactionReceiptLog) GetLogIndex() int64`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *TransactionReceiptLog) GetLogIndexOk() (*int64, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *TransactionReceiptLog) SetLogIndex(v int64)`

SetLogIndex sets LogIndex field to given value.


### GetAddress

`func (o *TransactionReceiptLog) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransactionReceiptLog) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransactionReceiptLog) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTopics

`func (o *TransactionReceiptLog) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *TransactionReceiptLog) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *TransactionReceiptLog) SetTopics(v []string)`

SetTopics sets Topics field to given value.


### GetData

`func (o *TransactionReceiptLog) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionReceiptLog) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionReceiptLog) SetData(v string)`

SetData sets Data field to given value.


### GetBlockNumber

`func (o *TransactionReceiptLog) GetBlockNumber() int64`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *TransactionReceiptLog) GetBlockNumberOk() (*int64, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *TransactionReceiptLog) SetBlockNumber(v int64)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *TransactionReceiptLog) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetBlockHash

`func (o *TransactionReceiptLog) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TransactionReceiptLog) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TransactionReceiptLog) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *TransactionReceiptLog) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetTransactionHash

`func (o *TransactionReceiptLog) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *TransactionReceiptLog) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *TransactionReceiptLog) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *TransactionReceiptLog) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetTransactionIndex

`func (o *TransactionReceiptLog) GetTransactionIndex() int64`

GetTransactionIndex returns the TransactionIndex field if non-nil, zero value otherwise.

### GetTransactionIndexOk

`func (o *TransactionReceiptLog) GetTransactionIndexOk() (*int64, bool)`

GetTransactionIndexOk returns a tuple with the TransactionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIndex

`func (o *TransactionReceiptLog) SetTransactionIndex(v int64)`

SetTransactionIndex sets TransactionIndex field to given value.

### HasTransactionIndex

`func (o *TransactionReceiptLog) HasTransactionIndex() bool`

HasTransactionIndex returns a boolean if a field has been set.

### GetRemoved

`func (o *TransactionReceiptLog) GetRemoved() bool`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *TransactionReceiptLog) GetRemovedOk() (*bool, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *TransactionReceiptLog) SetRemoved(v bool)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *TransactionReceiptLog) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


