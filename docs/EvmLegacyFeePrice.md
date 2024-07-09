# EvmLegacyFeePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | Pointer to **string** | The token ID of the transaction fee. | [optional] 
**Slow** | Pointer to [**EvmLegacyFeeBasePrice**](EvmLegacyFeeBasePrice.md) |  | [optional] 
**Recommended** | [**EvmLegacyFeeBasePrice**](EvmLegacyFeeBasePrice.md) |  | 
**Fast** | Pointer to [**EvmLegacyFeeBasePrice**](EvmLegacyFeeBasePrice.md) |  | [optional] 

## Methods

### NewEvmLegacyFeePrice

`func NewEvmLegacyFeePrice(feeType FeeType, recommended EvmLegacyFeeBasePrice, ) *EvmLegacyFeePrice`

NewEvmLegacyFeePrice instantiates a new EvmLegacyFeePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmLegacyFeePriceWithDefaults

`func NewEvmLegacyFeePriceWithDefaults() *EvmLegacyFeePrice`

NewEvmLegacyFeePriceWithDefaults instantiates a new EvmLegacyFeePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *EvmLegacyFeePrice) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EvmLegacyFeePrice) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EvmLegacyFeePrice) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *EvmLegacyFeePrice) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EvmLegacyFeePrice) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EvmLegacyFeePrice) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *EvmLegacyFeePrice) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetSlow

`func (o *EvmLegacyFeePrice) GetSlow() EvmLegacyFeeBasePrice`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EvmLegacyFeePrice) GetSlowOk() (*EvmLegacyFeeBasePrice, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EvmLegacyFeePrice) SetSlow(v EvmLegacyFeeBasePrice)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EvmLegacyFeePrice) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *EvmLegacyFeePrice) GetRecommended() EvmLegacyFeeBasePrice`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *EvmLegacyFeePrice) GetRecommendedOk() (*EvmLegacyFeeBasePrice, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *EvmLegacyFeePrice) SetRecommended(v EvmLegacyFeeBasePrice)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *EvmLegacyFeePrice) GetFast() EvmLegacyFeeBasePrice`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EvmLegacyFeePrice) GetFastOk() (*EvmLegacyFeeBasePrice, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EvmLegacyFeePrice) SetFast(v EvmLegacyFeeBasePrice)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EvmLegacyFeePrice) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


