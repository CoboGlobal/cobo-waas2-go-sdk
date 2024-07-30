# EstimatedEvmLegacyFeeSlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasPrice** | **string** | The gas price, in wei. The gas price represents the amount of ETH that must be paid to validators for processing transactions per gas unit used. | 
**GasLimit** | **string** | The gas limit. It represents the maximum number of gas units that you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. The gas unit cost of each operation varies. | [default to "21000"]

## Methods

### NewEstimatedEvmLegacyFeeSlow

`func NewEstimatedEvmLegacyFeeSlow(gasPrice string, gasLimit string, ) *EstimatedEvmLegacyFeeSlow`

NewEstimatedEvmLegacyFeeSlow instantiates a new EstimatedEvmLegacyFeeSlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedEvmLegacyFeeSlowWithDefaults

`func NewEstimatedEvmLegacyFeeSlowWithDefaults() *EstimatedEvmLegacyFeeSlow`

NewEstimatedEvmLegacyFeeSlowWithDefaults instantiates a new EstimatedEvmLegacyFeeSlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasPrice

`func (o *EstimatedEvmLegacyFeeSlow) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *EstimatedEvmLegacyFeeSlow) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *EstimatedEvmLegacyFeeSlow) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetGasLimit

`func (o *EstimatedEvmLegacyFeeSlow) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *EstimatedEvmLegacyFeeSlow) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *EstimatedEvmLegacyFeeSlow) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


