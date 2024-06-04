# ChainFeePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**FeeTokenId** | Pointer to **string** | The token ID of the transaction fee. | [optional] 
**MaxPriorityFee** | **string** | The max priority fee, in gwei. The max priority fee represents the highest amount of miner tips you are willing to pay for your transaction. | 
**BaseFee** | **string** | The base fee of chain. | 
**GasPrice** | **string** | The gas price, in gwei. The gas price represents the amount of ETH that must be paid to validators for processing transactions. | 
**FeeRate** | **string** | The fee rate, in sats/vByte. The fee rate represents the satoshis you are willing to pay for each byte of data that your transaction will consume on the blockchain. | 
**MaxFeeAmount** | Pointer to **string** | The maximum fee amount in fee_coin. | [optional] 

## Methods

### NewChainFeePrice

`func NewChainFeePrice(feeType FeeType, maxPriorityFee string, baseFee string, gasPrice string, feeRate string, ) *ChainFeePrice`

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

### GetMaxPriorityFee

`func (o *ChainFeePrice) GetMaxPriorityFee() string`

GetMaxPriorityFee returns the MaxPriorityFee field if non-nil, zero value otherwise.

### GetMaxPriorityFeeOk

`func (o *ChainFeePrice) GetMaxPriorityFeeOk() (*string, bool)`

GetMaxPriorityFeeOk returns a tuple with the MaxPriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFee

`func (o *ChainFeePrice) SetMaxPriorityFee(v string)`

SetMaxPriorityFee sets MaxPriorityFee field to given value.


### GetBaseFee

`func (o *ChainFeePrice) GetBaseFee() string`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *ChainFeePrice) GetBaseFeeOk() (*string, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *ChainFeePrice) SetBaseFee(v string)`

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


### GetMaxFeeAmount

`func (o *ChainFeePrice) GetMaxFeeAmount() string`

GetMaxFeeAmount returns the MaxFeeAmount field if non-nil, zero value otherwise.

### GetMaxFeeAmountOk

`func (o *ChainFeePrice) GetMaxFeeAmountOk() (*string, bool)`

GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeAmount

`func (o *ChainFeePrice) SetMaxFeeAmount(v string)`

SetMaxFeeAmount sets MaxFeeAmount field to given value.

### HasMaxFeeAmount

`func (o *ChainFeePrice) HasMaxFeeAmount() bool`

HasMaxFeeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


