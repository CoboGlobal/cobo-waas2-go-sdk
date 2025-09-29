# EstimatedEvmEip1559Fee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 
**Slow** | Pointer to [**EstimatedEvmEip1559FeeSlow**](EstimatedEvmEip1559FeeSlow.md) |  | [optional] 
**Recommended** | [**EstimatedEvmEip1559FeeSlow**](EstimatedEvmEip1559FeeSlow.md) |  | 
**Fast** | Pointer to [**EstimatedEvmEip1559FeeSlow**](EstimatedEvmEip1559FeeSlow.md) |  | [optional] 

## Methods

### NewEstimatedEvmEip1559Fee

`func NewEstimatedEvmEip1559Fee(feeType FeeType, tokenId string, recommended EstimatedEvmEip1559FeeSlow, ) *EstimatedEvmEip1559Fee`

NewEstimatedEvmEip1559Fee instantiates a new EstimatedEvmEip1559Fee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedEvmEip1559FeeWithDefaults

`func NewEstimatedEvmEip1559FeeWithDefaults() *EstimatedEvmEip1559Fee`

NewEstimatedEvmEip1559FeeWithDefaults instantiates a new EstimatedEvmEip1559Fee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *EstimatedEvmEip1559Fee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EstimatedEvmEip1559Fee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EstimatedEvmEip1559Fee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *EstimatedEvmEip1559Fee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EstimatedEvmEip1559Fee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EstimatedEvmEip1559Fee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetSlow

`func (o *EstimatedEvmEip1559Fee) GetSlow() EstimatedEvmEip1559FeeSlow`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EstimatedEvmEip1559Fee) GetSlowOk() (*EstimatedEvmEip1559FeeSlow, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EstimatedEvmEip1559Fee) SetSlow(v EstimatedEvmEip1559FeeSlow)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EstimatedEvmEip1559Fee) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *EstimatedEvmEip1559Fee) GetRecommended() EstimatedEvmEip1559FeeSlow`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *EstimatedEvmEip1559Fee) GetRecommendedOk() (*EstimatedEvmEip1559FeeSlow, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *EstimatedEvmEip1559Fee) SetRecommended(v EstimatedEvmEip1559FeeSlow)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *EstimatedEvmEip1559Fee) GetFast() EstimatedEvmEip1559FeeSlow`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EstimatedEvmEip1559Fee) GetFastOk() (*EstimatedEvmEip1559FeeSlow, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EstimatedEvmEip1559Fee) SetFast(v EstimatedEvmEip1559FeeSlow)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EstimatedEvmEip1559Fee) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


