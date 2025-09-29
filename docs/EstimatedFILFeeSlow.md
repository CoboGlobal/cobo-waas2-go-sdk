# EstimatedFILFeeSlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasPremium** | **string** | An optional additional fee that users can include to prioritize their transactions over others. It acts like a tip to incentivize miners to select and include your transaction over transactions with only the base fee. | 
**GasFeeCap** | **string** | The gas_fee_cap is a user-defined limit on how much they are willing to pay per unit of gas. | 
**GasLimit** | **string** | This defines the maximum amount of computational effort that a transaction is allowed to consume. It&#39;s a way to cap the resources that a transaction can use, ensuring it doesn&#39;t consume excessive network resources. | 
**GasBase** | **string** | This is the minimum fee required to include a transaction in a block. It is determined by the network&#39;s congestion level, which adjusts to maintain a target block utilization rate. The base fee is burned, reducing the total supply of Filecoin over time. | 

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


