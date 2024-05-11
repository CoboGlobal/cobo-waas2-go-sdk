# EvmEip1559FeeSlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeTokenId** | Pointer to **string** | ID of the fee token. Unique in all chains scope. | [optional] 
**MaxFee** | **string** | The highest Gas price paid for the transfer, unit GWei. | 
**MaxPriorityFee** | **int32** | The maximum Gas price paid to miners, the higher it is, the faster it is likely to be packaged into the block, unit GWei. | 
**BaseFee** | **int32** | The Base Fee of chain. | 
**GasLimit** | Pointer to **int32** | The Limit of gas. | [optional] [default to 21000]
**FeeAmount** | Pointer to **string** | The estimated fee amount in fee_coin. | [optional] 

## Methods

### NewEvmEip1559FeeSlow

`func NewEvmEip1559FeeSlow(maxFee string, maxPriorityFee int32, baseFee int32, ) *EvmEip1559FeeSlow`

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

### GetMaxFee

`func (o *EvmEip1559FeeSlow) GetMaxFee() string`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *EvmEip1559FeeSlow) GetMaxFeeOk() (*string, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *EvmEip1559FeeSlow) SetMaxFee(v string)`

SetMaxFee sets MaxFee field to given value.


### GetMaxPriorityFee

`func (o *EvmEip1559FeeSlow) GetMaxPriorityFee() int32`

GetMaxPriorityFee returns the MaxPriorityFee field if non-nil, zero value otherwise.

### GetMaxPriorityFeeOk

`func (o *EvmEip1559FeeSlow) GetMaxPriorityFeeOk() (*int32, bool)`

GetMaxPriorityFeeOk returns a tuple with the MaxPriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFee

`func (o *EvmEip1559FeeSlow) SetMaxPriorityFee(v int32)`

SetMaxPriorityFee sets MaxPriorityFee field to given value.


### GetBaseFee

`func (o *EvmEip1559FeeSlow) GetBaseFee() int32`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *EvmEip1559FeeSlow) GetBaseFeeOk() (*int32, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *EvmEip1559FeeSlow) SetBaseFee(v int32)`

SetBaseFee sets BaseFee field to given value.


### GetGasLimit

`func (o *EvmEip1559FeeSlow) GetGasLimit() int32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *EvmEip1559FeeSlow) GetGasLimitOk() (*int32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *EvmEip1559FeeSlow) SetGasLimit(v int32)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *EvmEip1559FeeSlow) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetFeeAmount

`func (o *EvmEip1559FeeSlow) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *EvmEip1559FeeSlow) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *EvmEip1559FeeSlow) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *EvmEip1559FeeSlow) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


