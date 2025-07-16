# EstimatedFILFeeSlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasPremium** | **string** | An optional tip you can include to prioritize your transaction. The gas premium incentivizes miners to include your transaction sooner than those offering only the base fee. | 
**GasFeeCap** | **string** | The maximum gas price you are willing to pay per unit of gas. | 
**GasLimit** | **string** | The maximum amount of gas your transaction is allowed to consume. | 
**GasBase** | **string** | The minimum fee required for a transaction to be included in a block. The base fee is dynamically adjusted based on network congestion to maintain target block utilization. It is burned rather than paid to miners, reducing the total Filecoin supply over time. | 

## Methods

### NewEstimatedFILFeeSlow

`func NewEstimatedFILFeeSlow(gasPremium string, gasFeeCap string, gasLimit string, gasBase string, ) *EstimatedFILFeeSlow`

NewEstimatedFILFeeSlow instantiates a new EstimatedFILFeeSlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedFILFeeSlowWithDefaults

`func NewEstimatedFILFeeSlowWithDefaults() *EstimatedFILFeeSlow`

NewEstimatedFILFeeSlowWithDefaults instantiates a new EstimatedFILFeeSlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasPremium

`func (o *EstimatedFILFeeSlow) GetGasPremium() string`

GetGasPremium returns the GasPremium field if non-nil, zero value otherwise.

### GetGasPremiumOk

`func (o *EstimatedFILFeeSlow) GetGasPremiumOk() (*string, bool)`

GetGasPremiumOk returns a tuple with the GasPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPremium

`func (o *EstimatedFILFeeSlow) SetGasPremium(v string)`

SetGasPremium sets GasPremium field to given value.


### GetGasFeeCap

`func (o *EstimatedFILFeeSlow) GetGasFeeCap() string`

GetGasFeeCap returns the GasFeeCap field if non-nil, zero value otherwise.

### GetGasFeeCapOk

`func (o *EstimatedFILFeeSlow) GetGasFeeCapOk() (*string, bool)`

GetGasFeeCapOk returns a tuple with the GasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeeCap

`func (o *EstimatedFILFeeSlow) SetGasFeeCap(v string)`

SetGasFeeCap sets GasFeeCap field to given value.


### GetGasLimit

`func (o *EstimatedFILFeeSlow) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *EstimatedFILFeeSlow) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *EstimatedFILFeeSlow) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasBase

`func (o *EstimatedFILFeeSlow) GetGasBase() string`

GetGasBase returns the GasBase field if non-nil, zero value otherwise.

### GetGasBaseOk

`func (o *EstimatedFILFeeSlow) GetGasBaseOk() (*string, bool)`

GetGasBaseOk returns a tuple with the GasBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasBase

`func (o *EstimatedFILFeeSlow) SetGasBase(v string)`

SetGasBase sets GasBase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


