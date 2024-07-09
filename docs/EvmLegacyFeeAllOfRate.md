# EvmLegacyFeeAllOfRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slow** | Pointer to [**EvmLegacyFeeBasePrice**](EvmLegacyFeeBasePrice.md) |  | [optional] 
**Recommended** | [**EvmLegacyFeeBasePrice**](EvmLegacyFeeBasePrice.md) |  | 
**Fast** | Pointer to [**EvmLegacyFeeBasePrice**](EvmLegacyFeeBasePrice.md) |  | [optional] 
**Customized** | Pointer to [**EvmLegacyFeeBasePrice**](EvmLegacyFeeBasePrice.md) |  | [optional] 

## Methods

### NewEvmLegacyFeeAllOfRate

`func NewEvmLegacyFeeAllOfRate(recommended EvmLegacyFeeBasePrice, ) *EvmLegacyFeeAllOfRate`

NewEvmLegacyFeeAllOfRate instantiates a new EvmLegacyFeeAllOfRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmLegacyFeeAllOfRateWithDefaults

`func NewEvmLegacyFeeAllOfRateWithDefaults() *EvmLegacyFeeAllOfRate`

NewEvmLegacyFeeAllOfRateWithDefaults instantiates a new EvmLegacyFeeAllOfRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlow

`func (o *EvmLegacyFeeAllOfRate) GetSlow() EvmLegacyFeeBasePrice`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EvmLegacyFeeAllOfRate) GetSlowOk() (*EvmLegacyFeeBasePrice, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EvmLegacyFeeAllOfRate) SetSlow(v EvmLegacyFeeBasePrice)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EvmLegacyFeeAllOfRate) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *EvmLegacyFeeAllOfRate) GetRecommended() EvmLegacyFeeBasePrice`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *EvmLegacyFeeAllOfRate) GetRecommendedOk() (*EvmLegacyFeeBasePrice, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *EvmLegacyFeeAllOfRate) SetRecommended(v EvmLegacyFeeBasePrice)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *EvmLegacyFeeAllOfRate) GetFast() EvmLegacyFeeBasePrice`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EvmLegacyFeeAllOfRate) GetFastOk() (*EvmLegacyFeeBasePrice, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EvmLegacyFeeAllOfRate) SetFast(v EvmLegacyFeeBasePrice)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EvmLegacyFeeAllOfRate) HasFast() bool`

HasFast returns a boolean if a field has been set.

### GetCustomized

`func (o *EvmLegacyFeeAllOfRate) GetCustomized() EvmLegacyFeeBasePrice`

GetCustomized returns the Customized field if non-nil, zero value otherwise.

### GetCustomizedOk

`func (o *EvmLegacyFeeAllOfRate) GetCustomizedOk() (*EvmLegacyFeeBasePrice, bool)`

GetCustomizedOk returns a tuple with the Customized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomized

`func (o *EvmLegacyFeeAllOfRate) SetCustomized(v EvmLegacyFeeBasePrice)`

SetCustomized sets Customized field to given value.

### HasCustomized

`func (o *EvmLegacyFeeAllOfRate) HasCustomized() bool`

HasCustomized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


