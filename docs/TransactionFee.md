# TransactionFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**FeeTokenId** | Pointer to **string** | ID of the fee token. Unique in all chains scope. | [optional] 
**MaxFee** | **string** | The highest Gas price paid for the transfer, unit GWei. | 
**MaxPriorityFee** | **int32** | The maximum Gas price paid to miners, the higher it is, the faster it is likely to be packaged into the block, unit GWei. | 
**BaseFee** | **int32** | The Base Fee of chain. | 
**GasLimit** | Pointer to **int32** | The Limit of gas. | [optional] [default to 21000]
**FeeAmount** | Pointer to **string** | The estimated fee amount in fee_coin. | [optional] 
**GasPrice** | **string** | The Price of Gas, unit GWei. | 
**FeeRate** | **string** | The fee rate, unit sat/vB. | 

## Methods

### NewTransactionFee

`func NewTransactionFee(feeType FeeType, maxFee string, maxPriorityFee int32, baseFee int32, gasPrice string, feeRate string, ) *TransactionFee`

NewTransactionFee instantiates a new TransactionFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFeeWithDefaults

`func NewTransactionFeeWithDefaults() *TransactionFee`

NewTransactionFeeWithDefaults instantiates a new TransactionFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *TransactionFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *TransactionFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *TransactionFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetFeeTokenId

`func (o *TransactionFee) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *TransactionFee) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *TransactionFee) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *TransactionFee) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetMaxFee

`func (o *TransactionFee) GetMaxFee() string`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *TransactionFee) GetMaxFeeOk() (*string, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *TransactionFee) SetMaxFee(v string)`

SetMaxFee sets MaxFee field to given value.


### GetMaxPriorityFee

`func (o *TransactionFee) GetMaxPriorityFee() int32`

GetMaxPriorityFee returns the MaxPriorityFee field if non-nil, zero value otherwise.

### GetMaxPriorityFeeOk

`func (o *TransactionFee) GetMaxPriorityFeeOk() (*int32, bool)`

GetMaxPriorityFeeOk returns a tuple with the MaxPriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFee

`func (o *TransactionFee) SetMaxPriorityFee(v int32)`

SetMaxPriorityFee sets MaxPriorityFee field to given value.


### GetBaseFee

`func (o *TransactionFee) GetBaseFee() int32`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *TransactionFee) GetBaseFeeOk() (*int32, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *TransactionFee) SetBaseFee(v int32)`

SetBaseFee sets BaseFee field to given value.


### GetGasLimit

`func (o *TransactionFee) GetGasLimit() int32`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *TransactionFee) GetGasLimitOk() (*int32, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *TransactionFee) SetGasLimit(v int32)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *TransactionFee) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetFeeAmount

`func (o *TransactionFee) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *TransactionFee) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *TransactionFee) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *TransactionFee) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetGasPrice

`func (o *TransactionFee) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *TransactionFee) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *TransactionFee) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetFeeRate

`func (o *TransactionFee) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *TransactionFee) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *TransactionFee) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


