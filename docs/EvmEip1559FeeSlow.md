# EvmEip1559FeeSlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeTokenId** | Pointer to **string** | The token ID of the transaction fee. Unique in all chains scope. | [optional] 
**MaxPriorityFee** | **string** | The max priority fee, in gwei. The max priority fee represents the highest amount of miner tips you are willing to pay for your transaction. | 
**BaseFee** | **string** | The base fee of chain. | 
**GasLimit** | Pointer to **string** | The gas limit, which represents the max number of gas units you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. Different operations require varying quantities of gas units. | [optional] [default to "21000"]

## Methods

### NewEvmEip1559FeeSlow

`func NewEvmEip1559FeeSlow(maxPriorityFee string, baseFee string, ) *EvmEip1559FeeSlow`

NewEvmEip1559FeeSlow instantiates a new EvmEip1559FeeSlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmEip1559FeeSlowWithDefaults

`func NewEvmEip1559FeeSlowWithDefaults() *EvmEip1559FeeSlow`

NewEvmEip1559FeeSlowWithDefaults instantiates a new EvmEip1559FeeSlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeTokenId

`func (o *EvmEip1559FeeSlow) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *EvmEip1559FeeSlow) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *EvmEip1559FeeSlow) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *EvmEip1559FeeSlow) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetMaxPriorityFee

`func (o *EvmEip1559FeeSlow) GetMaxPriorityFee() string`

GetMaxPriorityFee returns the MaxPriorityFee field if non-nil, zero value otherwise.

### GetMaxPriorityFeeOk

`func (o *EvmEip1559FeeSlow) GetMaxPriorityFeeOk() (*string, bool)`

GetMaxPriorityFeeOk returns a tuple with the MaxPriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFee

`func (o *EvmEip1559FeeSlow) SetMaxPriorityFee(v string)`

SetMaxPriorityFee sets MaxPriorityFee field to given value.


### GetBaseFee

`func (o *EvmEip1559FeeSlow) GetBaseFee() string`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *EvmEip1559FeeSlow) GetBaseFeeOk() (*string, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *EvmEip1559FeeSlow) SetBaseFee(v string)`

SetBaseFee sets BaseFee field to given value.


### GetGasLimit

`func (o *EvmEip1559FeeSlow) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *EvmEip1559FeeSlow) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *EvmEip1559FeeSlow) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *EvmEip1559FeeSlow) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


