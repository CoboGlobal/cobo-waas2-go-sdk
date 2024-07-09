# TransactionEvmLegacyFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasPrice** | **string** | The gas price, in gwei. The gas price represents the amount of ETH that must be paid to validators for processing transactions. | 
**GasLimit** | **string** | The gas limit. It represents the maximum number of gas units that you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. The gas unit cost of each operation varies. | [default to "21000"]
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**FeeUsed** | Pointer to **string** | The transaction fee. | [optional] 
**GasUsed** | Pointer to **string** | The gas units used in the transaction. | [optional] 

## Methods

### NewTransactionEvmLegacyFee

`func NewTransactionEvmLegacyFee(gasPrice string, gasLimit string, feeType FeeType, ) *TransactionEvmLegacyFee`

NewTransactionEvmLegacyFee instantiates a new TransactionEvmLegacyFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEvmLegacyFeeWithDefaults

`func NewTransactionEvmLegacyFeeWithDefaults() *TransactionEvmLegacyFee`

NewTransactionEvmLegacyFeeWithDefaults instantiates a new TransactionEvmLegacyFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasPrice

`func (o *TransactionEvmLegacyFee) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *TransactionEvmLegacyFee) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *TransactionEvmLegacyFee) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetGasLimit

`func (o *TransactionEvmLegacyFee) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *TransactionEvmLegacyFee) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *TransactionEvmLegacyFee) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetFeeType

`func (o *TransactionEvmLegacyFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *TransactionEvmLegacyFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *TransactionEvmLegacyFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetFeeUsed

`func (o *TransactionEvmLegacyFee) GetFeeUsed() string`

GetFeeUsed returns the FeeUsed field if non-nil, zero value otherwise.

### GetFeeUsedOk

`func (o *TransactionEvmLegacyFee) GetFeeUsedOk() (*string, bool)`

GetFeeUsedOk returns a tuple with the FeeUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeUsed

`func (o *TransactionEvmLegacyFee) SetFeeUsed(v string)`

SetFeeUsed sets FeeUsed field to given value.

### HasFeeUsed

`func (o *TransactionEvmLegacyFee) HasFeeUsed() bool`

HasFeeUsed returns a boolean if a field has been set.

### GetGasUsed

`func (o *TransactionEvmLegacyFee) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *TransactionEvmLegacyFee) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *TransactionEvmLegacyFee) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.

### HasGasUsed

`func (o *TransactionEvmLegacyFee) HasGasUsed() bool`

HasGasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

