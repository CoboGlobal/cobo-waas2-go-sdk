# EstimatedUtxoFeeSlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeRate** | **string** | The fee rate in sat/vByte. The fee rate represents the satoshis you are willing to pay for each byte of data that your transaction will consume on the blockchain. | 
**FeeAmount** | **string** | The fee that you need to pay for the transaction. | 

## Methods

### NewEstimatedUtxoFeeSlow

`func NewEstimatedUtxoFeeSlow(feeRate string, feeAmount string, ) *EstimatedUtxoFeeSlow`

NewEstimatedUtxoFeeSlow instantiates a new EstimatedUtxoFeeSlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedUtxoFeeSlowWithDefaults

`func NewEstimatedUtxoFeeSlowWithDefaults() *EstimatedUtxoFeeSlow`

NewEstimatedUtxoFeeSlowWithDefaults instantiates a new EstimatedUtxoFeeSlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeRate

`func (o *EstimatedUtxoFeeSlow) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *EstimatedUtxoFeeSlow) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *EstimatedUtxoFeeSlow) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetFeeAmount

`func (o *EstimatedUtxoFeeSlow) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *EstimatedUtxoFeeSlow) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *EstimatedUtxoFeeSlow) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


