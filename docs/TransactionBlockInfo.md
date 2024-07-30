# TransactionBlockInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockNumber** | Pointer to **int64** | The block number. | [optional] 
**BlockTime** | Pointer to **int64** | The time when the block was created, in Unix timestamp format, measured in milliseconds. | [optional] 
**BlockHash** | Pointer to **string** | The block hash. | [optional] 

## Methods

### NewTransactionBlockInfo

`func NewTransactionBlockInfo() *TransactionBlockInfo`

NewTransactionBlockInfo instantiates a new TransactionBlockInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionBlockInfoWithDefaults

`func NewTransactionBlockInfoWithDefaults() *TransactionBlockInfo`

NewTransactionBlockInfoWithDefaults instantiates a new TransactionBlockInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockNumber

`func (o *TransactionBlockInfo) GetBlockNumber() int64`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *TransactionBlockInfo) GetBlockNumberOk() (*int64, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *TransactionBlockInfo) SetBlockNumber(v int64)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *TransactionBlockInfo) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetBlockTime

`func (o *TransactionBlockInfo) GetBlockTime() int64`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *TransactionBlockInfo) GetBlockTimeOk() (*int64, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *TransactionBlockInfo) SetBlockTime(v int64)`

SetBlockTime sets BlockTime field to given value.

### HasBlockTime

`func (o *TransactionBlockInfo) HasBlockTime() bool`

HasBlockTime returns a boolean if a field has been set.

### GetBlockHash

`func (o *TransactionBlockInfo) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TransactionBlockInfo) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TransactionBlockInfo) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *TransactionBlockInfo) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


