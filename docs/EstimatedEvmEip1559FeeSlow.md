# EstimatedEvmEip1559FeeSlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxFeePerGas** | **string** | The maximum gas fee per gas unit used on the chain, in wei. | 
**MaxPriorityFeePerGas** | **string** | The maximum priority fee per gas unit used, in wei. The maximum priority fee represents the highest amount of miner tips that you are willing to pay for your transaction. | 
**GasLimit** | **string** | The gas limit. It represents the maximum number of gas units that you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. The gas unit cost of each operation varies. | [default to "21000"]

## Methods

### NewEstimatedEvmEip1559FeeSlow

`func NewEstimatedEvmEip1559FeeSlow(maxFeePerGas string, maxPriorityFeePerGas string, gasLimit string, ) *EstimatedEvmEip1559FeeSlow`

NewEstimatedEvmEip1559FeeSlow instantiates a new EstimatedEvmEip1559FeeSlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedEvmEip1559FeeSlowWithDefaults

`func NewEstimatedEvmEip1559FeeSlowWithDefaults() *EstimatedEvmEip1559FeeSlow`

NewEstimatedEvmEip1559FeeSlowWithDefaults instantiates a new EstimatedEvmEip1559FeeSlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxFeePerGas

`func (o *EstimatedEvmEip1559FeeSlow) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *EstimatedEvmEip1559FeeSlow) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *EstimatedEvmEip1559FeeSlow) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.


### GetMaxPriorityFeePerGas

`func (o *EstimatedEvmEip1559FeeSlow) GetMaxPriorityFeePerGas() string`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *EstimatedEvmEip1559FeeSlow) GetMaxPriorityFeePerGasOk() (*string, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *EstimatedEvmEip1559FeeSlow) SetMaxPriorityFeePerGas(v string)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.


### GetGasLimit

`func (o *EstimatedEvmEip1559FeeSlow) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *EstimatedEvmEip1559FeeSlow) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *EstimatedEvmEip1559FeeSlow) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


