# EstimatedUtxoFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token used to pay the transaction fee. | 
**Slow** | Pointer to [**EstimatedUtxoFeeSlow**](EstimatedUtxoFeeSlow.md) |  | [optional] 
**Recommended** | [**EstimatedUtxoFeeSlow**](EstimatedUtxoFeeSlow.md) |  | 
**Fast** | Pointer to [**EstimatedUtxoFeeSlow**](EstimatedUtxoFeeSlow.md) |  | [optional] 

## Methods

### NewEstimatedUtxoFee

`func NewEstimatedUtxoFee(feeType FeeType, tokenId string, recommended EstimatedUtxoFeeSlow, ) *EstimatedUtxoFee`

NewEstimatedUtxoFee instantiates a new EstimatedUtxoFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedUtxoFeeWithDefaults

`func NewEstimatedUtxoFeeWithDefaults() *EstimatedUtxoFee`

NewEstimatedUtxoFeeWithDefaults instantiates a new EstimatedUtxoFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *EstimatedUtxoFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EstimatedUtxoFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EstimatedUtxoFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *EstimatedUtxoFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EstimatedUtxoFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EstimatedUtxoFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetSlow

`func (o *EstimatedUtxoFee) GetSlow() EstimatedUtxoFeeSlow`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EstimatedUtxoFee) GetSlowOk() (*EstimatedUtxoFeeSlow, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EstimatedUtxoFee) SetSlow(v EstimatedUtxoFeeSlow)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EstimatedUtxoFee) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *EstimatedUtxoFee) GetRecommended() EstimatedUtxoFeeSlow`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *EstimatedUtxoFee) GetRecommendedOk() (*EstimatedUtxoFeeSlow, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *EstimatedUtxoFee) SetRecommended(v EstimatedUtxoFeeSlow)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *EstimatedUtxoFee) GetFast() EstimatedUtxoFeeSlow`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EstimatedUtxoFee) GetFastOk() (*EstimatedUtxoFeeSlow, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EstimatedUtxoFee) SetFast(v EstimatedUtxoFeeSlow)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EstimatedUtxoFee) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


