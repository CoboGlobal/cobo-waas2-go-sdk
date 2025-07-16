# TransactionRequestFILFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasPremium** | **string** | An optional tip you can include to prioritize your transaction. The gas premium incentivizes miners to include your transaction sooner than those offering only the base fee. | 
**GasFeeCap** | **string** | The maximum gas price you are willing to pay per unit of gas. | 
**GasLimit** | Pointer to **string** | The maximum amount of gas your transaction is allowed to consume. | [optional] 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token used to pay the transaction fee. | 

## Methods

### NewTransactionRequestFILFee

`func NewTransactionRequestFILFee(gasPremium string, gasFeeCap string, feeType FeeType, tokenId string, ) *TransactionRequestFILFee`

NewTransactionRequestFILFee instantiates a new TransactionRequestFILFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRequestFILFeeWithDefaults

`func NewTransactionRequestFILFeeWithDefaults() *TransactionRequestFILFee`

NewTransactionRequestFILFeeWithDefaults instantiates a new TransactionRequestFILFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasPremium

`func (o *TransactionRequestFILFee) GetGasPremium() string`

GetGasPremium returns the GasPremium field if non-nil, zero value otherwise.

### GetGasPremiumOk

`func (o *TransactionRequestFILFee) GetGasPremiumOk() (*string, bool)`

GetGasPremiumOk returns a tuple with the GasPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPremium

`func (o *TransactionRequestFILFee) SetGasPremium(v string)`

SetGasPremium sets GasPremium field to given value.


### GetGasFeeCap

`func (o *TransactionRequestFILFee) GetGasFeeCap() string`

GetGasFeeCap returns the GasFeeCap field if non-nil, zero value otherwise.

### GetGasFeeCapOk

`func (o *TransactionRequestFILFee) GetGasFeeCapOk() (*string, bool)`

GetGasFeeCapOk returns a tuple with the GasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeeCap

`func (o *TransactionRequestFILFee) SetGasFeeCap(v string)`

SetGasFeeCap sets GasFeeCap field to given value.


### GetGasLimit

`func (o *TransactionRequestFILFee) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *TransactionRequestFILFee) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *TransactionRequestFILFee) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *TransactionRequestFILFee) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetFeeType

`func (o *TransactionRequestFILFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *TransactionRequestFILFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *TransactionRequestFILFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *TransactionRequestFILFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionRequestFILFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionRequestFILFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


