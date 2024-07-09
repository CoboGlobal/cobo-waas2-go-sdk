# EvmEip1559Fee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | Pointer to **string** | The token ID of the transaction fee. | [optional] 
**Slow** | Pointer to [**EvmEip1559FeeSlow**](EvmEip1559FeeSlow.md) |  | [optional] 
**Recommended** | [**EvmEip1559FeeSlow**](EvmEip1559FeeSlow.md) |  | 
**Fast** | Pointer to [**EvmEip1559FeeSlow**](EvmEip1559FeeSlow.md) |  | [optional] 

## Methods

### NewEvmEip1559Fee

`func NewEvmEip1559Fee(feeType FeeType, recommended EvmEip1559FeeSlow, ) *EvmEip1559Fee`

NewEvmEip1559Fee instantiates a new EvmEip1559Fee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmEip1559FeeWithDefaults

`func NewEvmEip1559FeeWithDefaults() *EvmEip1559Fee`

NewEvmEip1559FeeWithDefaults instantiates a new EvmEip1559Fee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *EvmEip1559Fee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EvmEip1559Fee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EvmEip1559Fee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *EvmEip1559Fee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EvmEip1559Fee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EvmEip1559Fee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *EvmEip1559Fee) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetSlow

`func (o *EvmEip1559Fee) GetSlow() EvmEip1559FeeSlow`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EvmEip1559Fee) GetSlowOk() (*EvmEip1559FeeSlow, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EvmEip1559Fee) SetSlow(v EvmEip1559FeeSlow)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EvmEip1559Fee) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *EvmEip1559Fee) GetRecommended() EvmEip1559FeeSlow`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *EvmEip1559Fee) GetRecommendedOk() (*EvmEip1559FeeSlow, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *EvmEip1559Fee) SetRecommended(v EvmEip1559FeeSlow)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *EvmEip1559Fee) GetFast() EvmEip1559FeeSlow`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EvmEip1559Fee) GetFastOk() (*EvmEip1559FeeSlow, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EvmEip1559Fee) SetFast(v EvmEip1559FeeSlow)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EvmEip1559Fee) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


