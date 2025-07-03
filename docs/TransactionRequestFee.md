# TransactionRequestFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 
**MaxFeeAmount** | Pointer to **string** | The maximum fee that you are willing to pay for the transaction. Provide the value without applying precision. The transaction will fail if the transaction fee exceeds the maximum fee. | [optional] 
**GasLimit** | Pointer to **string** | This defines the maximum amount of computational effort that a transaction is allowed to consume. It&#39;s a way to cap the resources that a transaction can use, ensuring it doesn&#39;t consume excessive network resources. | [optional] 
**MaxFeePerGas** | **string** | The maximum gas fee per gas unit used on the chain, in wei. | 
**MaxPriorityFeePerGas** | **string** | The maximum priority fee per gas unit used, in wei. The maximum priority fee represents the highest amount of miner tips that you are willing to pay for your transaction. | 
**GasPrice** | **string** | The gas price, in wei. The gas price represents the amount of ETH that must be paid to validators for processing transactions per gas unit used. | 
**FeeRate** | Pointer to **string** | The fee rate in sat/vByte. The fee rate represents the satoshis you are willing to pay for each byte of data that your transaction will consume on the blockchain. | [optional] 
**ComputeUnitPrice** | **string** | The cost per compute unit. Transactions consume computational resources measured in compute units, and this price helps determine the cost of executing transactions, especially complex ones involving smart contracts. | 
**ComputeUnitLimit** | **string** | The maximum number of compute units allowed for a transaction. This limits the resources any single transaction can consume, preventing excessive resource usage that could impact network performance negatively. | 
**GasPremium** | **string** | An optional additional fee that users can include to prioritize their transactions over others. It acts like a tip to incentivize miners to select and include your transaction over transactions with only the base fee. | 
**GasFeeCap** | **string** | The gas_fee_cap is a user-defined limit on how much they are willing to pay per unit of gas. | 

## Methods

### NewTransactionRequestFee

`func NewTransactionRequestFee(feeType FeeType, tokenId string, maxFeePerGas string, maxPriorityFeePerGas string, gasPrice string, computeUnitPrice string, computeUnitLimit string, gasPremium string, gasFeeCap string, ) *TransactionRequestFee`

NewTransactionRequestFee instantiates a new TransactionRequestFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRequestFeeWithDefaults

`func NewTransactionRequestFeeWithDefaults() *TransactionRequestFee`

NewTransactionRequestFeeWithDefaults instantiates a new TransactionRequestFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *TransactionRequestFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *TransactionRequestFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *TransactionRequestFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *TransactionRequestFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionRequestFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionRequestFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetMaxFeeAmount

`func (o *TransactionRequestFee) GetMaxFeeAmount() string`

GetMaxFeeAmount returns the MaxFeeAmount field if non-nil, zero value otherwise.

### GetMaxFeeAmountOk

`func (o *TransactionRequestFee) GetMaxFeeAmountOk() (*string, bool)`

GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeAmount

`func (o *TransactionRequestFee) SetMaxFeeAmount(v string)`

SetMaxFeeAmount sets MaxFeeAmount field to given value.

### HasMaxFeeAmount

`func (o *TransactionRequestFee) HasMaxFeeAmount() bool`

HasMaxFeeAmount returns a boolean if a field has been set.

### GetGasLimit

`func (o *TransactionRequestFee) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *TransactionRequestFee) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *TransactionRequestFee) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *TransactionRequestFee) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *TransactionRequestFee) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *TransactionRequestFee) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *TransactionRequestFee) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.


### GetMaxPriorityFeePerGas

`func (o *TransactionRequestFee) GetMaxPriorityFeePerGas() string`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *TransactionRequestFee) GetMaxPriorityFeePerGasOk() (*string, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *TransactionRequestFee) SetMaxPriorityFeePerGas(v string)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.


### GetGasPrice

`func (o *TransactionRequestFee) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *TransactionRequestFee) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *TransactionRequestFee) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetFeeRate

`func (o *TransactionRequestFee) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *TransactionRequestFee) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *TransactionRequestFee) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *TransactionRequestFee) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetComputeUnitPrice

`func (o *TransactionRequestFee) GetComputeUnitPrice() string`

GetComputeUnitPrice returns the ComputeUnitPrice field if non-nil, zero value otherwise.

### GetComputeUnitPriceOk

`func (o *TransactionRequestFee) GetComputeUnitPriceOk() (*string, bool)`

GetComputeUnitPriceOk returns a tuple with the ComputeUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeUnitPrice

`func (o *TransactionRequestFee) SetComputeUnitPrice(v string)`

SetComputeUnitPrice sets ComputeUnitPrice field to given value.


### GetComputeUnitLimit

`func (o *TransactionRequestFee) GetComputeUnitLimit() string`

GetComputeUnitLimit returns the ComputeUnitLimit field if non-nil, zero value otherwise.

### GetComputeUnitLimitOk

`func (o *TransactionRequestFee) GetComputeUnitLimitOk() (*string, bool)`

GetComputeUnitLimitOk returns a tuple with the ComputeUnitLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeUnitLimit

`func (o *TransactionRequestFee) SetComputeUnitLimit(v string)`

SetComputeUnitLimit sets ComputeUnitLimit field to given value.


### GetGasPremium

`func (o *TransactionRequestFee) GetGasPremium() string`

GetGasPremium returns the GasPremium field if non-nil, zero value otherwise.

### GetGasPremiumOk

`func (o *TransactionRequestFee) GetGasPremiumOk() (*string, bool)`

GetGasPremiumOk returns a tuple with the GasPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPremium

`func (o *TransactionRequestFee) SetGasPremium(v string)`

SetGasPremium sets GasPremium field to given value.


### GetGasFeeCap

`func (o *TransactionRequestFee) GetGasFeeCap() string`

GetGasFeeCap returns the GasFeeCap field if non-nil, zero value otherwise.

### GetGasFeeCapOk

`func (o *TransactionRequestFee) GetGasFeeCapOk() (*string, bool)`

GetGasFeeCapOk returns a tuple with the GasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeeCap

`func (o *TransactionRequestFee) SetGasFeeCap(v string)`

SetGasFeeCap sets GasFeeCap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


