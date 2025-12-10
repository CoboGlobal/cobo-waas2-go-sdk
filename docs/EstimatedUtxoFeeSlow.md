# EstimatedUtxoFeeSlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeRate** | **string** | The fee rate in sat/vByte. The fee rate represents the satoshis you are willing to pay for each byte of data that your transaction will consume on the blockchain. | 
**Fallback** | Pointer to **bool** | Indicates whether the estimated fee is generated from Cobo’s fallback mechanism. When the estimated transaction belongs to a UTXO-based chain and the specified address does not have sufficient balance to cover the on-chain fee, this field will be set to &#x60;true&#x60;. In this case, the returned fee value is estimated by Cobo’s internal fallback strategy, which is typically higher than the actual on-chain fee. When &#x60;fallback&#x60; is &#x60;true&#x60;, please use the estimated fee value with caution. | [optional] 
**FeeAmount** | **string** | The transaction fee that you need to pay for the transaction. | 

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


### GetFallback

`func (o *EstimatedUtxoFeeSlow) GetFallback() bool`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *EstimatedUtxoFeeSlow) GetFallbackOk() (*bool, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *EstimatedUtxoFeeSlow) SetFallback(v bool)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *EstimatedUtxoFeeSlow) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

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


