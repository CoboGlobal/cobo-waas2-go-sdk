# EvmLegacyTransactionFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeTokenId** | Pointer to **string** | The token ID of the transaction fee. | [optional] 
**GasPrice** | **string** | The gas price, in gwei. The gas price represents the amount of ETH that must be paid to validators for processing transactions. | 
**GasLimit** | Pointer to **string** | The gas limit, which represents the max number of gas units you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. Different operations require varying quantities of gas units. | [optional] [default to "21000"]
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]

## Methods

### NewEvmLegacyTransactionFee

`func NewEvmLegacyTransactionFee(gasPrice string, feeType FeeType, ) *EvmLegacyTransactionFee`

NewEvmLegacyTransactionFee instantiates a new EvmLegacyTransactionFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmLegacyTransactionFeeWithDefaults

`func NewEvmLegacyTransactionFeeWithDefaults() *EvmLegacyTransactionFee`

NewEvmLegacyTransactionFeeWithDefaults instantiates a new EvmLegacyTransactionFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeTokenId

`func (o *EvmLegacyTransactionFee) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *EvmLegacyTransactionFee) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *EvmLegacyTransactionFee) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *EvmLegacyTransactionFee) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetGasPrice

`func (o *EvmLegacyTransactionFee) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *EvmLegacyTransactionFee) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *EvmLegacyTransactionFee) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetGasLimit

`func (o *EvmLegacyTransactionFee) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *EvmLegacyTransactionFee) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *EvmLegacyTransactionFee) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *EvmLegacyTransactionFee) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetFeeType

`func (o *EvmLegacyTransactionFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EvmLegacyTransactionFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EvmLegacyTransactionFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


