# EstimatedEvmLegacyFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 
**Slow** | Pointer to [**EstimatedEvmLegacyFeeSlow**](EstimatedEvmLegacyFeeSlow.md) |  | [optional] 
**Recommended** | [**EstimatedEvmLegacyFeeSlow**](EstimatedEvmLegacyFeeSlow.md) |  | 
**Fast** | Pointer to [**EstimatedEvmLegacyFeeSlow**](EstimatedEvmLegacyFeeSlow.md) |  | [optional] 

## Methods

### NewEstimatedEvmLegacyFee

`func NewEstimatedEvmLegacyFee(feeType FeeType, tokenId string, recommended EstimatedEvmLegacyFeeSlow, ) *EstimatedEvmLegacyFee`

NewEstimatedEvmLegacyFee instantiates a new EstimatedEvmLegacyFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedEvmLegacyFeeWithDefaults

`func NewEstimatedEvmLegacyFeeWithDefaults() *EstimatedEvmLegacyFee`

NewEstimatedEvmLegacyFeeWithDefaults instantiates a new EstimatedEvmLegacyFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *EstimatedEvmLegacyFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EstimatedEvmLegacyFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EstimatedEvmLegacyFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *EstimatedEvmLegacyFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EstimatedEvmLegacyFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EstimatedEvmLegacyFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetSlow

`func (o *EstimatedEvmLegacyFee) GetSlow() EstimatedEvmLegacyFeeSlow`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EstimatedEvmLegacyFee) GetSlowOk() (*EstimatedEvmLegacyFeeSlow, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EstimatedEvmLegacyFee) SetSlow(v EstimatedEvmLegacyFeeSlow)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EstimatedEvmLegacyFee) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *EstimatedEvmLegacyFee) GetRecommended() EstimatedEvmLegacyFeeSlow`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *EstimatedEvmLegacyFee) GetRecommendedOk() (*EstimatedEvmLegacyFeeSlow, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *EstimatedEvmLegacyFee) SetRecommended(v EstimatedEvmLegacyFeeSlow)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *EstimatedEvmLegacyFee) GetFast() EstimatedEvmLegacyFeeSlow`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EstimatedEvmLegacyFee) GetFastOk() (*EstimatedEvmLegacyFeeSlow, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EstimatedEvmLegacyFee) SetFast(v EstimatedEvmLegacyFeeSlow)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EstimatedEvmLegacyFee) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


