# EstimatedFILFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token used to pay the transaction fee. | 
**Slow** | Pointer to [**EstimatedFILFeeSlow**](EstimatedFILFeeSlow.md) |  | [optional] 
**Recommended** | [**EstimatedFILFeeSlow**](EstimatedFILFeeSlow.md) |  | 
**Fast** | Pointer to [**EstimatedFILFeeSlow**](EstimatedFILFeeSlow.md) |  | [optional] 

## Methods

### NewEstimatedFILFee

`func NewEstimatedFILFee(feeType FeeType, tokenId string, recommended EstimatedFILFeeSlow, ) *EstimatedFILFee`

NewEstimatedFILFee instantiates a new EstimatedFILFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedFILFeeWithDefaults

`func NewEstimatedFILFeeWithDefaults() *EstimatedFILFee`

NewEstimatedFILFeeWithDefaults instantiates a new EstimatedFILFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *EstimatedFILFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EstimatedFILFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EstimatedFILFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *EstimatedFILFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EstimatedFILFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EstimatedFILFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetSlow

`func (o *EstimatedFILFee) GetSlow() EstimatedFILFeeSlow`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EstimatedFILFee) GetSlowOk() (*EstimatedFILFeeSlow, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EstimatedFILFee) SetSlow(v EstimatedFILFeeSlow)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EstimatedFILFee) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *EstimatedFILFee) GetRecommended() EstimatedFILFeeSlow`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *EstimatedFILFee) GetRecommendedOk() (*EstimatedFILFeeSlow, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *EstimatedFILFee) SetRecommended(v EstimatedFILFeeSlow)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *EstimatedFILFee) GetFast() EstimatedFILFeeSlow`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EstimatedFILFee) GetFastOk() (*EstimatedFILFeeSlow, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EstimatedFILFee) SetFast(v EstimatedFILFeeSlow)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EstimatedFILFee) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


