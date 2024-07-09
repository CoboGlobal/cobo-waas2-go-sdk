# EstimationFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 
**Slow** | Pointer to [**UtxoFeeSlow**](UtxoFeeSlow.md) |  | [optional] 
**Recommended** | [**UtxoFeeSlow**](UtxoFeeSlow.md) |  | 
**Fast** | Pointer to [**UtxoFeeSlow**](UtxoFeeSlow.md) |  | [optional] 
**MaxFeeAmount** | **string** | The maximum fee that you are willing to pay for the transaction. The transaction will fail if the transaction fee exceeds the maximum fee. | 

## Methods

### NewEstimationFee

`func NewEstimationFee(feeType FeeType, tokenId string, recommended UtxoFeeSlow, maxFeeAmount string, ) *EstimationFee`

NewEstimationFee instantiates a new EstimationFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimationFeeWithDefaults

`func NewEstimationFeeWithDefaults() *EstimationFee`

NewEstimationFeeWithDefaults instantiates a new EstimationFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *EstimationFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EstimationFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EstimationFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *EstimationFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EstimationFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EstimationFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetSlow

`func (o *EstimationFee) GetSlow() UtxoFeeSlow`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EstimationFee) GetSlowOk() (*UtxoFeeSlow, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EstimationFee) SetSlow(v UtxoFeeSlow)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EstimationFee) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *EstimationFee) GetRecommended() UtxoFeeSlow`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *EstimationFee) GetRecommendedOk() (*UtxoFeeSlow, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *EstimationFee) SetRecommended(v UtxoFeeSlow)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *EstimationFee) GetFast() UtxoFeeSlow`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EstimationFee) GetFastOk() (*UtxoFeeSlow, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EstimationFee) SetFast(v UtxoFeeSlow)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EstimationFee) HasFast() bool`

HasFast returns a boolean if a field has been set.

### GetMaxFeeAmount

`func (o *EstimationFee) GetMaxFeeAmount() string`

GetMaxFeeAmount returns the MaxFeeAmount field if non-nil, zero value otherwise.

### GetMaxFeeAmountOk

`func (o *EstimationFee) GetMaxFeeAmountOk() (*string, bool)`

GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeAmount

`func (o *EstimationFee) SetMaxFeeAmount(v string)`

SetMaxFeeAmount sets MaxFeeAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


