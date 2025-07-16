# TransactionRequestEvmLegacyFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasPrice** | **string** | The gas price, in wei. The gas price represents the amount of ETH that must be paid to validators for processing transactions per gas unit used. | 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token used to pay the transaction fee. | 
**GasLimit** | Pointer to **string** | The gas limit. It represents the maximum number of gas units that you are willing to pay for the execution of a transaction or Ethereum Virtual Machine (EVM) operation. The gas unit cost of each operation varies. | [optional] 

## Methods

### NewTransactionRequestEvmLegacyFee

`func NewTransactionRequestEvmLegacyFee(gasPrice string, feeType FeeType, tokenId string, ) *TransactionRequestEvmLegacyFee`

NewTransactionRequestEvmLegacyFee instantiates a new TransactionRequestEvmLegacyFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRequestEvmLegacyFeeWithDefaults

`func NewTransactionRequestEvmLegacyFeeWithDefaults() *TransactionRequestEvmLegacyFee`

NewTransactionRequestEvmLegacyFeeWithDefaults instantiates a new TransactionRequestEvmLegacyFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasPrice

`func (o *TransactionRequestEvmLegacyFee) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *TransactionRequestEvmLegacyFee) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *TransactionRequestEvmLegacyFee) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetFeeType

`func (o *TransactionRequestEvmLegacyFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *TransactionRequestEvmLegacyFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *TransactionRequestEvmLegacyFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *TransactionRequestEvmLegacyFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionRequestEvmLegacyFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionRequestEvmLegacyFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetGasLimit

`func (o *TransactionRequestEvmLegacyFee) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *TransactionRequestEvmLegacyFee) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *TransactionRequestEvmLegacyFee) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *TransactionRequestEvmLegacyFee) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


