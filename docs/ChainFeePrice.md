# ChainFeePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**FeeTokenId** | Pointer to **string** | ID of the fee token. Unique in all chains scope. | [optional] 
**MaxFee** | **string** | The highest Gas price paid for the transfer, unit GWei. | 
**MaxPriorityFee** | **int32** | The maximum Gas price paid to miners, the higher it is, the faster it is likely to be packaged into the block, unit GWei. | 
**BaseFee** | **int32** | The Base Fee of chain. | 
**GasPrice** | **string** | The Price of Gas, unit GWei. | 
**FeeRate** | **string** | The fee rate, unit sat/vB. | 
**FeeAmount** | Pointer to **string** | The estimated fee amount in fee_coin. | [optional] 

## Methods

### NewChainFeePrice

`func NewChainFeePrice(feeType FeeType, maxFee string, maxPriorityFee int32, baseFee int32, gasPrice string, feeRate string, ) *ChainFeePrice`

NewChainFeePrice instantiates a new ChainFeePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainFeePriceWithDefaults

`func NewChainFeePriceWithDefaults() *ChainFeePrice`

NewChainFeePriceWithDefaults instantiates a new ChainFeePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *ChainFeePrice) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *ChainFeePrice) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *ChainFeePrice) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetFeeTokenId

`func (o *ChainFeePrice) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *ChainFeePrice) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *ChainFeePrice) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *ChainFeePrice) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetMaxFee

`func (o *ChainFeePrice) GetMaxFee() string`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *ChainFeePrice) GetMaxFeeOk() (*string, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *ChainFeePrice) SetMaxFee(v string)`

SetMaxFee sets MaxFee field to given value.


### GetMaxPriorityFee

`func (o *ChainFeePrice) GetMaxPriorityFee() int32`

GetMaxPriorityFee returns the MaxPriorityFee field if non-nil, zero value otherwise.

### GetMaxPriorityFeeOk

`func (o *ChainFeePrice) GetMaxPriorityFeeOk() (*int32, bool)`

GetMaxPriorityFeeOk returns a tuple with the MaxPriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFee

`func (o *ChainFeePrice) SetMaxPriorityFee(v int32)`

SetMaxPriorityFee sets MaxPriorityFee field to given value.


### GetBaseFee

`func (o *ChainFeePrice) GetBaseFee() int32`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *ChainFeePrice) GetBaseFeeOk() (*int32, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *ChainFeePrice) SetBaseFee(v int32)`

SetBaseFee sets BaseFee field to given value.


### GetGasPrice

`func (o *ChainFeePrice) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ChainFeePrice) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ChainFeePrice) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetFeeRate

`func (o *ChainFeePrice) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *ChainFeePrice) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *ChainFeePrice) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetFeeAmount

`func (o *ChainFeePrice) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *ChainFeePrice) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *ChainFeePrice) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *ChainFeePrice) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


