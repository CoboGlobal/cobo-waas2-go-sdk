# TransactionReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | 
**TransactionHash** | **string** | The transaction hash, returned with the &#x60;0x&#x60; prefix. | 
**Status** | **int32** | The execution result of the transaction. - &#x60;1&#x60;: The transaction was executed successfully. - &#x60;0&#x60;: The transaction failed, for example, because it was reverted by the contract or ran out of gas.  | 
**BlockNumber** | **int64** | The number of the block that contains the transaction. | 
**BlockHash** | **string** | The hash of the block that contains the transaction. | 
**TransactionIndex** | **int64** | The index position of the transaction within the block. | 
**FromAddress** | **string** | The address that sent the transaction. | 
**ToAddress** | Pointer to **NullableString** | The address that received the transaction. The value is &#x60;null&#x60; if the transaction created a contract. | [optional] 
**ContractAddress** | Pointer to **NullableString** | The address of the contract created by the transaction. The value is &#x60;null&#x60; if the transaction did not create a contract. | [optional] 
**GasUsed** | **string** | The number of gas units consumed by the transaction. | 
**CumulativeGasUsed** | Pointer to **string** | The total number of gas units consumed by all the transactions up to and including this transaction in the block. | [optional] 
**EffectiveGasPrice** | Pointer to **string** | The gas price actually paid for each gas unit consumed by the transaction, in the smallest unit of the chain&#39;s native token. For example, the value is in wei for Ethereum. | [optional] 
**EvmTransactionType** | Pointer to **int32** | The transaction envelope type defined by the chain. This property describes the on-chain transaction format, not the Cobo transaction type returned by [Get transaction information](https://www.cobo.com/developers/v2/api-references/transactions/get-transaction-information). - &#x60;0&#x60;: A legacy transaction. - &#x60;1&#x60;: An access list transaction, as defined in EIP-2930. - &#x60;2&#x60;: A dynamic fee transaction, as defined in EIP-1559.  | [optional] 
**LogsBloom** | Pointer to **string** | The bloom filter of the event logs emitted during the execution of the transaction, which can be used to quickly check whether the transaction emitted a specific log. The value is a 256-byte hexadecimal string. | [optional] 
**Logs** | [**[]TransactionReceiptLog**](TransactionReceiptLog.md) | The event logs emitted during the execution of the transaction, in the order in which they were emitted. The array is empty if the transaction emitted no logs. A maximum of 1,000 logs are returned. If the transaction emitted more logs, only the first 1,000 are returned and &#x60;logs_truncated&#x60; is &#x60;true&#x60;. | 
**LogsTruncated** | **bool** | Whether the event logs returned in &#x60;logs&#x60; were truncated. - &#x60;true&#x60;: The transaction emitted more than 1,000 logs and only the first 1,000 are returned. - &#x60;false&#x60;: All the logs emitted by the transaction are returned.  | 

## Methods

### NewTransactionReceipt

`func NewTransactionReceipt(chainId string, transactionHash string, status int32, blockNumber int64, blockHash string, transactionIndex int64, fromAddress string, gasUsed string, logs []TransactionReceiptLog, logsTruncated bool, ) *TransactionReceipt`

NewTransactionReceipt instantiates a new TransactionReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionReceiptWithDefaults

`func NewTransactionReceiptWithDefaults() *TransactionReceipt`

NewTransactionReceiptWithDefaults instantiates a new TransactionReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *TransactionReceipt) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TransactionReceipt) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TransactionReceipt) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetTransactionHash

`func (o *TransactionReceipt) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *TransactionReceipt) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *TransactionReceipt) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetStatus

`func (o *TransactionReceipt) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionReceipt) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionReceipt) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetBlockNumber

`func (o *TransactionReceipt) GetBlockNumber() int64`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *TransactionReceipt) GetBlockNumberOk() (*int64, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *TransactionReceipt) SetBlockNumber(v int64)`

SetBlockNumber sets BlockNumber field to given value.


### GetBlockHash

`func (o *TransactionReceipt) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TransactionReceipt) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TransactionReceipt) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetTransactionIndex

`func (o *TransactionReceipt) GetTransactionIndex() int64`

GetTransactionIndex returns the TransactionIndex field if non-nil, zero value otherwise.

### GetTransactionIndexOk

`func (o *TransactionReceipt) GetTransactionIndexOk() (*int64, bool)`

GetTransactionIndexOk returns a tuple with the TransactionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIndex

`func (o *TransactionReceipt) SetTransactionIndex(v int64)`

SetTransactionIndex sets TransactionIndex field to given value.


### GetFromAddress

`func (o *TransactionReceipt) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *TransactionReceipt) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *TransactionReceipt) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetToAddress

`func (o *TransactionReceipt) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *TransactionReceipt) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *TransactionReceipt) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *TransactionReceipt) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### SetToAddressNil

`func (o *TransactionReceipt) SetToAddressNil(b bool)`

 SetToAddressNil sets the value for ToAddress to be an explicit nil

### UnsetToAddress
`func (o *TransactionReceipt) UnsetToAddress()`

UnsetToAddress ensures that no value is present for ToAddress, not even an explicit nil
### GetContractAddress

`func (o *TransactionReceipt) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *TransactionReceipt) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *TransactionReceipt) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *TransactionReceipt) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### SetContractAddressNil

`func (o *TransactionReceipt) SetContractAddressNil(b bool)`

 SetContractAddressNil sets the value for ContractAddress to be an explicit nil

### UnsetContractAddress
`func (o *TransactionReceipt) UnsetContractAddress()`

UnsetContractAddress ensures that no value is present for ContractAddress, not even an explicit nil
### GetGasUsed

`func (o *TransactionReceipt) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *TransactionReceipt) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *TransactionReceipt) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetCumulativeGasUsed

`func (o *TransactionReceipt) GetCumulativeGasUsed() string`

GetCumulativeGasUsed returns the CumulativeGasUsed field if non-nil, zero value otherwise.

### GetCumulativeGasUsedOk

`func (o *TransactionReceipt) GetCumulativeGasUsedOk() (*string, bool)`

GetCumulativeGasUsedOk returns a tuple with the CumulativeGasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeGasUsed

`func (o *TransactionReceipt) SetCumulativeGasUsed(v string)`

SetCumulativeGasUsed sets CumulativeGasUsed field to given value.

### HasCumulativeGasUsed

`func (o *TransactionReceipt) HasCumulativeGasUsed() bool`

HasCumulativeGasUsed returns a boolean if a field has been set.

### GetEffectiveGasPrice

`func (o *TransactionReceipt) GetEffectiveGasPrice() string`

GetEffectiveGasPrice returns the EffectiveGasPrice field if non-nil, zero value otherwise.

### GetEffectiveGasPriceOk

`func (o *TransactionReceipt) GetEffectiveGasPriceOk() (*string, bool)`

GetEffectiveGasPriceOk returns a tuple with the EffectiveGasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveGasPrice

`func (o *TransactionReceipt) SetEffectiveGasPrice(v string)`

SetEffectiveGasPrice sets EffectiveGasPrice field to given value.

### HasEffectiveGasPrice

`func (o *TransactionReceipt) HasEffectiveGasPrice() bool`

HasEffectiveGasPrice returns a boolean if a field has been set.

### GetEvmTransactionType

`func (o *TransactionReceipt) GetEvmTransactionType() int32`

GetEvmTransactionType returns the EvmTransactionType field if non-nil, zero value otherwise.

### GetEvmTransactionTypeOk

`func (o *TransactionReceipt) GetEvmTransactionTypeOk() (*int32, bool)`

GetEvmTransactionTypeOk returns a tuple with the EvmTransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvmTransactionType

`func (o *TransactionReceipt) SetEvmTransactionType(v int32)`

SetEvmTransactionType sets EvmTransactionType field to given value.

### HasEvmTransactionType

`func (o *TransactionReceipt) HasEvmTransactionType() bool`

HasEvmTransactionType returns a boolean if a field has been set.

### GetLogsBloom

`func (o *TransactionReceipt) GetLogsBloom() string`

GetLogsBloom returns the LogsBloom field if non-nil, zero value otherwise.

### GetLogsBloomOk

`func (o *TransactionReceipt) GetLogsBloomOk() (*string, bool)`

GetLogsBloomOk returns a tuple with the LogsBloom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsBloom

`func (o *TransactionReceipt) SetLogsBloom(v string)`

SetLogsBloom sets LogsBloom field to given value.

### HasLogsBloom

`func (o *TransactionReceipt) HasLogsBloom() bool`

HasLogsBloom returns a boolean if a field has been set.

### GetLogs

`func (o *TransactionReceipt) GetLogs() []TransactionReceiptLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *TransactionReceipt) GetLogsOk() (*[]TransactionReceiptLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *TransactionReceipt) SetLogs(v []TransactionReceiptLog)`

SetLogs sets Logs field to given value.


### GetLogsTruncated

`func (o *TransactionReceipt) GetLogsTruncated() bool`

GetLogsTruncated returns the LogsTruncated field if non-nil, zero value otherwise.

### GetLogsTruncatedOk

`func (o *TransactionReceipt) GetLogsTruncatedOk() (*bool, bool)`

GetLogsTruncatedOk returns a tuple with the LogsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsTruncated

`func (o *TransactionReceipt) SetLogsTruncated(v bool)`

SetLogsTruncated sets LogsTruncated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


