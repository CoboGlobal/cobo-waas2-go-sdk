# TransactionEvmEip1559Fee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPriorityFee** | **string** | The maximum priority fee, in gwei. The maximum priority fee represents the highest amount of miner tips that you are willing to pay for your transaction. | 
**BaseFee** | **string** | The base fee price of the chain, in gwei. | 
**GasLimit** | **string** | The gas limit. It represents the maximum number of gas units that you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. The gas unit cost of each operation varies. | [default to "21000"]
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**FeeUsed** | Pointer to **string** | The transaction fee. | [optional] 
**GasUsed** | Pointer to **string** | The number of gas units used in the transaction. | [optional] 

## Methods

### NewTransactionEvmEip1559Fee

`func NewTransactionEvmEip1559Fee(maxPriorityFee string, baseFee string, gasLimit string, feeType FeeType, ) *TransactionEvmEip1559Fee`

NewTransactionEvmEip1559Fee instantiates a new TransactionEvmEip1559Fee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEvmEip1559FeeWithDefaults

`func NewTransactionEvmEip1559FeeWithDefaults() *TransactionEvmEip1559Fee`

NewTransactionEvmEip1559FeeWithDefaults instantiates a new TransactionEvmEip1559Fee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPriorityFee

`func (o *TransactionEvmEip1559Fee) GetMaxPriorityFee() string`

GetMaxPriorityFee returns the MaxPriorityFee field if non-nil, zero value otherwise.

### GetMaxPriorityFeeOk

`func (o *TransactionEvmEip1559Fee) GetMaxPriorityFeeOk() (*string, bool)`

GetMaxPriorityFeeOk returns a tuple with the MaxPriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFee

`func (o *TransactionEvmEip1559Fee) SetMaxPriorityFee(v string)`

SetMaxPriorityFee sets MaxPriorityFee field to given value.


### GetBaseFee

`func (o *TransactionEvmEip1559Fee) GetBaseFee() string`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *TransactionEvmEip1559Fee) GetBaseFeeOk() (*string, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *TransactionEvmEip1559Fee) SetBaseFee(v string)`

SetBaseFee sets BaseFee field to given value.


### GetGasLimit

`func (o *TransactionEvmEip1559Fee) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *TransactionEvmEip1559Fee) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *TransactionEvmEip1559Fee) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetFeeType

`func (o *TransactionEvmEip1559Fee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *TransactionEvmEip1559Fee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *TransactionEvmEip1559Fee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetFeeUsed

`func (o *TransactionEvmEip1559Fee) GetFeeUsed() string`

GetFeeUsed returns the FeeUsed field if non-nil, zero value otherwise.

### GetFeeUsedOk

`func (o *TransactionEvmEip1559Fee) GetFeeUsedOk() (*string, bool)`

GetFeeUsedOk returns a tuple with the FeeUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeUsed

`func (o *TransactionEvmEip1559Fee) SetFeeUsed(v string)`

SetFeeUsed sets FeeUsed field to given value.

### HasFeeUsed

`func (o *TransactionEvmEip1559Fee) HasFeeUsed() bool`

HasFeeUsed returns a boolean if a field has been set.

### GetGasUsed

`func (o *TransactionEvmEip1559Fee) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *TransactionEvmEip1559Fee) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *TransactionEvmEip1559Fee) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.

### HasGasUsed

`func (o *TransactionEvmEip1559Fee) HasGasUsed() bool`

HasGasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


