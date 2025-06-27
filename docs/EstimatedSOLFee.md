# EstimatedSOLFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 
**Slow** | Pointer to [**EstimatedSOLFeeSlow**](EstimatedSOLFeeSlow.md) |  | [optional] 
**Recommended** | [**EstimatedSOLFeeSlow**](EstimatedSOLFeeSlow.md) |  | 
**Fast** | Pointer to [**EstimatedSOLFeeSlow**](EstimatedSOLFeeSlow.md) |  | [optional] 

## Methods

### NewEstimatedSOLFee

`func NewEstimatedSOLFee(feeType FeeType, tokenId string, recommended EstimatedSOLFeeSlow, ) *EstimatedSOLFee`

NewEstimatedSOLFee instantiates a new EstimatedSOLFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedSOLFeeWithDefaults

`func NewEstimatedSOLFeeWithDefaults() *EstimatedSOLFee`

NewEstimatedSOLFeeWithDefaults instantiates a new EstimatedSOLFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *EstimatedSOLFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EstimatedSOLFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EstimatedSOLFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *EstimatedSOLFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EstimatedSOLFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EstimatedSOLFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetSlow

`func (o *EstimatedSOLFee) GetSlow() EstimatedSOLFeeSlow`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EstimatedSOLFee) GetSlowOk() (*EstimatedSOLFeeSlow, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EstimatedSOLFee) SetSlow(v EstimatedSOLFeeSlow)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EstimatedSOLFee) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *EstimatedSOLFee) GetRecommended() EstimatedSOLFeeSlow`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *EstimatedSOLFee) GetRecommendedOk() (*EstimatedSOLFeeSlow, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *EstimatedSOLFee) SetRecommended(v EstimatedSOLFeeSlow)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *EstimatedSOLFee) GetFast() EstimatedSOLFeeSlow`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EstimatedSOLFee) GetFastOk() (*EstimatedSOLFeeSlow, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EstimatedSOLFee) SetFast(v EstimatedSOLFeeSlow)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EstimatedSOLFee) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


