# PayoutFeeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 
**MaxFeeAmount** | Pointer to **string** | The maximum fee that you are willing to pay for the transaction. Provide the value without applying precision. The transaction will fail if the transaction fee exceeds the maximum fee. | [optional] 
**GasLimit** | Pointer to **string** | The gas limit. It represents the maximum number of gas units that you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. The gas unit cost of each operation varies. | [optional] 
**MaxFeePerGas** | **string** | The maximum gas fee per gas unit used on the chain, in wei. | 
**MaxPriorityFeePerGas** | **string** | The maximum priority fee per gas unit used, in wei. The maximum priority fee represents the highest amount of miner tips that you are willing to pay for your transaction. | 
**GasPrice** | **string** | The gas price, in wei. The gas price represents the amount of ETH that must be paid to validators for processing transactions per gas unit used. | 
**FeeRate** | Pointer to **string** | The fee rate in sat/vByte. The fee rate represents the satoshis you are willing to pay for each byte of data that your transaction will consume on the blockchain. | [optional] 
**Fallback** | Pointer to **bool** | Indicates whether the estimated fee is generated from Cobo’s fallback mechanism. When the estimated transaction belongs to a UTXO-based chain and the specified address does not have sufficient balance to cover the on-chain fee, this field will be set to &#x60;true&#x60;. In this case, the returned fee value is estimated by Cobo’s internal fallback strategy, which is typically higher than the actual on-chain fee. When &#x60;fallback&#x60; is &#x60;true&#x60;, please use the estimated fee value with caution. | [optional] 

## Methods

### NewPayoutFeeData

`func NewPayoutFeeData(feeType FeeType, tokenId string, maxFeePerGas string, maxPriorityFeePerGas string, gasPrice string, ) *PayoutFeeData`

NewPayoutFeeData instantiates a new PayoutFeeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutFeeDataWithDefaults

`func NewPayoutFeeDataWithDefaults() *PayoutFeeData`

NewPayoutFeeDataWithDefaults instantiates a new PayoutFeeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *PayoutFeeData) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *PayoutFeeData) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *PayoutFeeData) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *PayoutFeeData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PayoutFeeData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PayoutFeeData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetMaxFeeAmount

`func (o *PayoutFeeData) GetMaxFeeAmount() string`

GetMaxFeeAmount returns the MaxFeeAmount field if non-nil, zero value otherwise.

### GetMaxFeeAmountOk

`func (o *PayoutFeeData) GetMaxFeeAmountOk() (*string, bool)`

GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeAmount

`func (o *PayoutFeeData) SetMaxFeeAmount(v string)`

SetMaxFeeAmount sets MaxFeeAmount field to given value.

### HasMaxFeeAmount

`func (o *PayoutFeeData) HasMaxFeeAmount() bool`

HasMaxFeeAmount returns a boolean if a field has been set.

### GetGasLimit

`func (o *PayoutFeeData) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *PayoutFeeData) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *PayoutFeeData) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *PayoutFeeData) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *PayoutFeeData) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *PayoutFeeData) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *PayoutFeeData) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.


### GetMaxPriorityFeePerGas

`func (o *PayoutFeeData) GetMaxPriorityFeePerGas() string`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *PayoutFeeData) GetMaxPriorityFeePerGasOk() (*string, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *PayoutFeeData) SetMaxPriorityFeePerGas(v string)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.


### GetGasPrice

`func (o *PayoutFeeData) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *PayoutFeeData) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *PayoutFeeData) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetFeeRate

`func (o *PayoutFeeData) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *PayoutFeeData) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *PayoutFeeData) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *PayoutFeeData) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetFallback

`func (o *PayoutFeeData) GetFallback() bool`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *PayoutFeeData) GetFallbackOk() (*bool, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *PayoutFeeData) SetFallback(v bool)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *PayoutFeeData) HasFallback() bool`

HasFallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


