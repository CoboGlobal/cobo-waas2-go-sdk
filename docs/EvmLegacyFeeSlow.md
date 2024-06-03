# EvmLegacyFeeSlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeTokenId** | Pointer to **string** | The token ID of the transaction fee. Unique in all chains scope. | [optional] 
**GasPrice** | **string** | The gas price, in gwei. The gas price represents the amount of ETH that must be paid to validators for processing transactions. | 
**GasLimit** | Pointer to **string** | The gas limit, which represents the max number of gas units you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. Different operations require varying quantities of gas units. | [optional] [default to "21000"]

## Methods

### NewEvmLegacyFeeSlow

`func NewEvmLegacyFeeSlow(gasPrice string, ) *EvmLegacyFeeSlow`

NewEvmLegacyFeeSlow instantiates a new EvmLegacyFeeSlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmLegacyFeeSlowWithDefaults

`func NewEvmLegacyFeeSlowWithDefaults() *EvmLegacyFeeSlow`

NewEvmLegacyFeeSlowWithDefaults instantiates a new EvmLegacyFeeSlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeTokenId

`func (o *EvmLegacyFeeSlow) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *EvmLegacyFeeSlow) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *EvmLegacyFeeSlow) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *EvmLegacyFeeSlow) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetGasPrice

`func (o *EvmLegacyFeeSlow) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *EvmLegacyFeeSlow) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *EvmLegacyFeeSlow) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetGasLimit

`func (o *EvmLegacyFeeSlow) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *EvmLegacyFeeSlow) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *EvmLegacyFeeSlow) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *EvmLegacyFeeSlow) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


