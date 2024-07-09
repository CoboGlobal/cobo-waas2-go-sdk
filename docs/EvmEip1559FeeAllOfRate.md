# EvmEip1559FeeAllOfRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slow** | Pointer to [**EvmEip1559FeeBasePrice**](EvmEip1559FeeBasePrice.md) |  | [optional] 
**Recommended** | [**EvmEip1559FeeBasePrice**](EvmEip1559FeeBasePrice.md) |  | 
**Fast** | Pointer to [**EvmEip1559FeeBasePrice**](EvmEip1559FeeBasePrice.md) |  | [optional] 
**Customized** | Pointer to [**EvmEip1559FeeBasePrice**](EvmEip1559FeeBasePrice.md) |  | [optional] 

## Methods

### NewEvmEip1559FeeAllOfRate

`func NewEvmEip1559FeeAllOfRate(recommended EvmEip1559FeeBasePrice, ) *EvmEip1559FeeAllOfRate`

NewEvmEip1559FeeAllOfRate instantiates a new EvmEip1559FeeAllOfRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmEip1559FeeAllOfRateWithDefaults

`func NewEvmEip1559FeeAllOfRateWithDefaults() *EvmEip1559FeeAllOfRate`

NewEvmEip1559FeeAllOfRateWithDefaults instantiates a new EvmEip1559FeeAllOfRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlow

`func (o *EvmEip1559FeeAllOfRate) GetSlow() EvmEip1559FeeBasePrice`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EvmEip1559FeeAllOfRate) GetSlowOk() (*EvmEip1559FeeBasePrice, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EvmEip1559FeeAllOfRate) SetSlow(v EvmEip1559FeeBasePrice)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EvmEip1559FeeAllOfRate) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *EvmEip1559FeeAllOfRate) GetRecommended() EvmEip1559FeeBasePrice`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *EvmEip1559FeeAllOfRate) GetRecommendedOk() (*EvmEip1559FeeBasePrice, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *EvmEip1559FeeAllOfRate) SetRecommended(v EvmEip1559FeeBasePrice)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *EvmEip1559FeeAllOfRate) GetFast() EvmEip1559FeeBasePrice`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EvmEip1559FeeAllOfRate) GetFastOk() (*EvmEip1559FeeBasePrice, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EvmEip1559FeeAllOfRate) SetFast(v EvmEip1559FeeBasePrice)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EvmEip1559FeeAllOfRate) HasFast() bool`

HasFast returns a boolean if a field has been set.

### GetCustomized

`func (o *EvmEip1559FeeAllOfRate) GetCustomized() EvmEip1559FeeBasePrice`

GetCustomized returns the Customized field if non-nil, zero value otherwise.

### GetCustomizedOk

`func (o *EvmEip1559FeeAllOfRate) GetCustomizedOk() (*EvmEip1559FeeBasePrice, bool)`

GetCustomizedOk returns a tuple with the Customized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomized

`func (o *EvmEip1559FeeAllOfRate) SetCustomized(v EvmEip1559FeeBasePrice)`

SetCustomized sets Customized field to given value.

### HasCustomized

`func (o *EvmEip1559FeeAllOfRate) HasCustomized() bool`

HasCustomized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


