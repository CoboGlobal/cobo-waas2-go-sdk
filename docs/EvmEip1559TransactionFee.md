# EvmEip1559TransactionFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPriorityFee** | **string** | The maximum priority fee, in gwei. The maximum priority fee represents the highest amount of miner tips that you are willing to pay for your transaction. | 
**BaseFee** | **string** | The base fee price of the chain, in gwei. | 
**GasLimit** | **string** | The gas limit. It represents the maximum number of gas units that you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. The gas unit cost of each operation varies. | [default to "21000"]
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 

## Methods

### NewEvmEip1559TransactionFee

`func NewEvmEip1559TransactionFee(maxPriorityFee string, baseFee string, gasLimit string, feeType FeeType, tokenId string, ) *EvmEip1559TransactionFee`

NewEvmEip1559TransactionFee instantiates a new EvmEip1559TransactionFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmEip1559TransactionFeeWithDefaults

`func NewEvmEip1559TransactionFeeWithDefaults() *EvmEip1559TransactionFee`

NewEvmEip1559TransactionFeeWithDefaults instantiates a new EvmEip1559TransactionFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPriorityFee

`func (o *EvmEip1559TransactionFee) GetMaxPriorityFee() string`

GetMaxPriorityFee returns the MaxPriorityFee field if non-nil, zero value otherwise.

### GetMaxPriorityFeeOk

`func (o *EvmEip1559TransactionFee) GetMaxPriorityFeeOk() (*string, bool)`

GetMaxPriorityFeeOk returns a tuple with the MaxPriorityFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFee

`func (o *EvmEip1559TransactionFee) SetMaxPriorityFee(v string)`

SetMaxPriorityFee sets MaxPriorityFee field to given value.


### GetBaseFee

`func (o *EvmEip1559TransactionFee) GetBaseFee() string`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *EvmEip1559TransactionFee) GetBaseFeeOk() (*string, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *EvmEip1559TransactionFee) SetBaseFee(v string)`

SetBaseFee sets BaseFee field to given value.


### GetGasLimit

`func (o *EvmEip1559TransactionFee) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *EvmEip1559TransactionFee) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *EvmEip1559TransactionFee) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


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


### GetTokenId

`func (o *EvmEip1559TransactionFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EvmEip1559TransactionFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EvmEip1559TransactionFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


