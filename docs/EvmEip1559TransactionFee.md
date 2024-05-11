# EvmEip1559TransactionFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeTokenId** | Pointer to **string** | ID of the fee token. Unique in all chains scope. | [optional] 
**MaxFee** | **string** | The highest Gas price paid for the transfer, unit GWei. | 
**MaxPriorityFee** | **int32** | The maximum Gas price paid to miners, the higher it is, the faster it is likely to be packaged into the block, unit GWei. | 
**BaseFee** | **int32** | The Base Fee of chain. | 
**GasLimit** | Pointer to **int32** | The Limit of gas. | [optional] [default to 21000]
**FeeAmount** | Pointer to **string** | The estimated fee amount in fee_coin. | [optional] 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]

## Methods

### NewEvmEip1559TransactionFee

`func NewEvmEip1559TransactionFee(maxFee string, maxPriorityFee int32, baseFee int32, feeType FeeType, ) *EvmEip1559TransactionFee`

NewEvmEip1559TransactionFee instantiates a new EvmEip1559TransactionFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmEip1559TransactionFeeWithDefaults

`func NewEvmEip1559TransactionFeeWithDefaults() *EvmEip1559TransactionFee`

NewEvmEip1559TransactionFeeWithDefaults instantiates a new EvmEip1559TransactionFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeTokenId

`func (o *EvmEip1559TransactionFee) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *EvmEip1559TransactionFee) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *EvmEip1559TransactionFee) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *EvmEip1559TransactionFee) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetMaxFee

`func (o *EvmEip1559TransactionFee) GetMaxFee() string`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *EvmEip1559TransactionFee) GetMaxFeeOk() (*string, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *EvmEip1559TransactionFee) SetMaxFee(v string)`

SetMaxFee sets MaxFee field to given value.


### GetMaxPriorityFee

`func (o *EvmEip1559TransactionFee) GetMaxPriorityFee() int32`

GetMaxPriorityFee returns the MaxPriorityFee field if non-nil, zero value otherwise.

### GetMaxPriorityFeeOk

`func (o *EvmEip1559TransactionFee) GetMaxPriorityFeeOk() (*int32, bool)`

GetMaxPriorityFeeOk returns a tuple with the MaxPriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFee

`func (o *EvmEip1559TransactionFee) SetMaxPriorityFee(v int32)`

SetMaxPriorityFee sets MaxPriorityFee field to given value.


### GetBaseFee

`func (o *EvmEip1559TransactionFee) GetBaseFee() int32`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *EvmEip1559TransactionFee) GetBaseFeeOk() (*int32, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *EvmEip1559TransactionFee) SetBaseFee(v int32)`

SetBaseFee sets BaseFee field to given value.


### GetGasLimit

`func (o *EvmEip1559TransactionFee) GetGasLimit() int32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *EvmEip1559TransactionFee) GetGasLimitOk() (*int32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *EvmEip1559TransactionFee) SetGasLimit(v int32)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *EvmEip1559TransactionFee) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetFeeAmount

`func (o *EvmEip1559TransactionFee) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *EvmEip1559TransactionFee) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *EvmEip1559TransactionFee) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *EvmEip1559TransactionFee) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetFeeType

`func (o *EvmEip1559TransactionFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EvmEip1559TransactionFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EvmEip1559TransactionFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


