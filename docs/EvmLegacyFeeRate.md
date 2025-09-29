# EvmLegacyFeeRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 
**Slow** | Pointer to [**EvmLegacyFeeBasePrice**](EvmLegacyFeeBasePrice.md) |  | [optional] 
**Recommended** | [**EvmLegacyFeeBasePrice**](EvmLegacyFeeBasePrice.md) |  | 
**Fast** | Pointer to [**EvmLegacyFeeBasePrice**](EvmLegacyFeeBasePrice.md) |  | [optional] 

## Methods

### NewEvmLegacyFeeRate

`func NewEvmLegacyFeeRate(feeType FeeType, tokenId string, recommended EvmLegacyFeeBasePrice, ) *EvmLegacyFeeRate`

NewEvmLegacyFeeRate instantiates a new EvmLegacyFeeRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmLegacyFeeRateWithDefaults

`func NewEvmLegacyFeeRateWithDefaults() *EvmLegacyFeeRate`

NewEvmLegacyFeeRateWithDefaults instantiates a new EvmLegacyFeeRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *EvmLegacyFeeRate) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EvmLegacyFeeRate) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EvmLegacyFeeRate) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *EvmLegacyFeeRate) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EvmLegacyFeeRate) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EvmLegacyFeeRate) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetSlow

`func (o *EvmLegacyFeeRate) GetSlow() EvmLegacyFeeBasePrice`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EvmLegacyFeeRate) GetSlowOk() (*EvmLegacyFeeBasePrice, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EvmLegacyFeeRate) SetSlow(v EvmLegacyFeeBasePrice)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EvmLegacyFeeRate) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *EvmLegacyFeeRate) GetRecommended() EvmLegacyFeeBasePrice`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *EvmLegacyFeeRate) GetRecommendedOk() (*EvmLegacyFeeBasePrice, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *EvmLegacyFeeRate) SetRecommended(v EvmLegacyFeeBasePrice)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *EvmLegacyFeeRate) GetFast() EvmLegacyFeeBasePrice`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EvmLegacyFeeRate) GetFastOk() (*EvmLegacyFeeBasePrice, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EvmLegacyFeeRate) SetFast(v EvmLegacyFeeBasePrice)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EvmLegacyFeeRate) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


