# TransactionFILFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasBase** | Pointer to **string** | The minimum fee required for a transaction to be included in a block. The base fee is dynamically adjusted based on network congestion to maintain target block utilization. It is burned rather than paid to miners, reducing the total Filecoin supply over time. | [optional] 
**GasPremium** | Pointer to **string** | An optional tip you can include to prioritize your transaction. The gas premium incentivizes miners to include your transaction sooner than those offering only the base fee. | [optional] 
**GasFeeCap** | Pointer to **string** | The maximum gas price you are willing to pay per unit of gas. | [optional] 
**GasLimit** | Pointer to **string** | The maximum amount of gas your transaction is allowed to consume. | [optional] 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | Pointer to **string** | The token used to pay the transaction fee. | [optional] 
**FeeUsed** | Pointer to **string** | The actually charged transaction fee. | [optional] 
**EstimatedFeeUsed** | Pointer to **string** | The estimated transaction fee. | [optional] 

## Methods

### NewTransactionFILFee

`func NewTransactionFILFee(feeType FeeType, ) *TransactionFILFee`

NewTransactionFILFee instantiates a new TransactionFILFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFILFeeWithDefaults

`func NewTransactionFILFeeWithDefaults() *TransactionFILFee`

NewTransactionFILFeeWithDefaults instantiates a new TransactionFILFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasBase

`func (o *TransactionFILFee) GetGasBase() string`

GetGasBase returns the GasBase field if non-nil, zero value otherwise.

### GetGasBaseOk

`func (o *TransactionFILFee) GetGasBaseOk() (*string, bool)`

GetGasBaseOk returns a tuple with the GasBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasBase

`func (o *TransactionFILFee) SetGasBase(v string)`

SetGasBase sets GasBase field to given value.

### HasGasBase

`func (o *TransactionFILFee) HasGasBase() bool`

HasGasBase returns a boolean if a field has been set.

### GetGasPremium

`func (o *TransactionFILFee) GetGasPremium() string`

GetGasPremium returns the GasPremium field if non-nil, zero value otherwise.

### GetGasPremiumOk

`func (o *TransactionFILFee) GetGasPremiumOk() (*string, bool)`

GetGasPremiumOk returns a tuple with the GasPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPremium

`func (o *TransactionFILFee) SetGasPremium(v string)`

SetGasPremium sets GasPremium field to given value.

### HasGasPremium

`func (o *TransactionFILFee) HasGasPremium() bool`

HasGasPremium returns a boolean if a field has been set.

### GetGasFeeCap

`func (o *TransactionFILFee) GetGasFeeCap() string`

GetGasFeeCap returns the GasFeeCap field if non-nil, zero value otherwise.

### GetGasFeeCapOk

`func (o *TransactionFILFee) GetGasFeeCapOk() (*string, bool)`

GetGasFeeCapOk returns a tuple with the GasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeeCap

`func (o *TransactionFILFee) SetGasFeeCap(v string)`

SetGasFeeCap sets GasFeeCap field to given value.

### HasGasFeeCap

`func (o *TransactionFILFee) HasGasFeeCap() bool`

HasGasFeeCap returns a boolean if a field has been set.

### GetGasLimit

`func (o *TransactionFILFee) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *TransactionFILFee) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *TransactionFILFee) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *TransactionFILFee) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetFeeType

`func (o *TransactionFILFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *TransactionFILFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *TransactionFILFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *TransactionFILFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionFILFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionFILFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TransactionFILFee) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetFeeUsed

`func (o *TransactionFILFee) GetFeeUsed() string`

GetFeeUsed returns the FeeUsed field if non-nil, zero value otherwise.

### GetFeeUsedOk

`func (o *TransactionFILFee) GetFeeUsedOk() (*string, bool)`

GetFeeUsedOk returns a tuple with the FeeUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeUsed

`func (o *TransactionFILFee) SetFeeUsed(v string)`

SetFeeUsed sets FeeUsed field to given value.

### HasFeeUsed

`func (o *TransactionFILFee) HasFeeUsed() bool`

HasFeeUsed returns a boolean if a field has been set.

### GetEstimatedFeeUsed

`func (o *TransactionFILFee) GetEstimatedFeeUsed() string`

GetEstimatedFeeUsed returns the EstimatedFeeUsed field if non-nil, zero value otherwise.

### GetEstimatedFeeUsedOk

`func (o *TransactionFILFee) GetEstimatedFeeUsedOk() (*string, bool)`

GetEstimatedFeeUsedOk returns a tuple with the EstimatedFeeUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedFeeUsed

`func (o *TransactionFILFee) SetEstimatedFeeUsed(v string)`

SetEstimatedFeeUsed sets EstimatedFeeUsed field to given value.

### HasEstimatedFeeUsed

`func (o *TransactionFILFee) HasEstimatedFeeUsed() bool`

HasEstimatedFeeUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


