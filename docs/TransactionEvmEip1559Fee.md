# TransactionEvmEip1559Fee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxFeePerGas** | Pointer to **string** | The maximum gas fee per gas unit used on the chain, in wei. | [optional] 
**MaxPriorityFeePerGas** | Pointer to **string** | The maximum priority fee per gas unit used, in wei. The maximum priority fee represents the highest amount of miner tips that you are willing to pay for your transaction. | [optional] 
**GasLimit** | Pointer to **string** | The gas limit. It represents the maximum number of gas units that you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. The gas unit cost of each operation varies. | [optional] [default to "21000"]
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | Pointer to **string** | The token ID of the transaction fee. | [optional] 
**EffectiveGasPrice** | Pointer to **string** | The gas price (gas fee per gas unit) on the chain, in wei. The gas price represents the amount of ETH that must be paid to validators for processing transactions. | [optional] 
**FeeUsed** | Pointer to **string** | The transaction fee. | [optional] 
**GasUsed** | Pointer to **string** | The number of gas units used in the transaction. | [optional] 

## Methods

### NewTransactionEvmEip1559Fee

`func NewTransactionEvmEip1559Fee(feeType FeeType, ) *TransactionEvmEip1559Fee`

NewTransactionEvmEip1559Fee instantiates a new TransactionEvmEip1559Fee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEvmEip1559FeeWithDefaults

`func NewTransactionEvmEip1559FeeWithDefaults() *TransactionEvmEip1559Fee`

NewTransactionEvmEip1559FeeWithDefaults instantiates a new TransactionEvmEip1559Fee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxFeePerGas

`func (o *TransactionEvmEip1559Fee) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *TransactionEvmEip1559Fee) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *TransactionEvmEip1559Fee) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *TransactionEvmEip1559Fee) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### GetMaxPriorityFeePerGas

`func (o *TransactionEvmEip1559Fee) GetMaxPriorityFeePerGas() string`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *TransactionEvmEip1559Fee) GetMaxPriorityFeePerGasOk() (*string, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *TransactionEvmEip1559Fee) SetMaxPriorityFeePerGas(v string)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.

### HasMaxPriorityFeePerGas

`func (o *TransactionEvmEip1559Fee) HasMaxPriorityFeePerGas() bool`

HasMaxPriorityFeePerGas returns a boolean if a field has been set.

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

### HasGasLimit

`func (o *TransactionEvmEip1559Fee) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

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


### GetTokenId

`func (o *TransactionEvmEip1559Fee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionEvmEip1559Fee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionEvmEip1559Fee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TransactionEvmEip1559Fee) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetEffectiveGasPrice

`func (o *TransactionEvmEip1559Fee) GetEffectiveGasPrice() string`

GetEffectiveGasPrice returns the EffectiveGasPrice field if non-nil, zero value otherwise.

### GetEffectiveGasPriceOk

`func (o *TransactionEvmEip1559Fee) GetEffectiveGasPriceOk() (*string, bool)`

GetEffectiveGasPriceOk returns a tuple with the EffectiveGasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveGasPrice

`func (o *TransactionEvmEip1559Fee) SetEffectiveGasPrice(v string)`

SetEffectiveGasPrice sets EffectiveGasPrice field to given value.

### HasEffectiveGasPrice

`func (o *TransactionEvmEip1559Fee) HasEffectiveGasPrice() bool`

HasEffectiveGasPrice returns a boolean if a field has been set.

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


