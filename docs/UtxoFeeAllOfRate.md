# UtxoFeeAllOfRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slow** | Pointer to [**UtxoFeeBasePrice**](UtxoFeeBasePrice.md) |  | [optional] 
**Recommended** | [**UtxoFeeBasePrice**](UtxoFeeBasePrice.md) |  | 
**Fast** | Pointer to [**UtxoFeeBasePrice**](UtxoFeeBasePrice.md) |  | [optional] 
**Customized** | Pointer to [**UtxoFeeBasePrice**](UtxoFeeBasePrice.md) |  | [optional] 

## Methods

### NewUtxoFeeAllOfRate

`func NewUtxoFeeAllOfRate(recommended UtxoFeeBasePrice, ) *UtxoFeeAllOfRate`

NewUtxoFeeAllOfRate instantiates a new UtxoFeeAllOfRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtxoFeeAllOfRateWithDefaults

`func NewUtxoFeeAllOfRateWithDefaults() *UtxoFeeAllOfRate`

NewUtxoFeeAllOfRateWithDefaults instantiates a new UtxoFeeAllOfRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlow

`func (o *UtxoFeeAllOfRate) GetSlow() UtxoFeeBasePrice`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *UtxoFeeAllOfRate) GetSlowOk() (*UtxoFeeBasePrice, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *UtxoFeeAllOfRate) SetSlow(v UtxoFeeBasePrice)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *UtxoFeeAllOfRate) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *UtxoFeeAllOfRate) GetRecommended() UtxoFeeBasePrice`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *UtxoFeeAllOfRate) GetRecommendedOk() (*UtxoFeeBasePrice, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *UtxoFeeAllOfRate) SetRecommended(v UtxoFeeBasePrice)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *UtxoFeeAllOfRate) GetFast() UtxoFeeBasePrice`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *UtxoFeeAllOfRate) GetFastOk() (*UtxoFeeBasePrice, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *UtxoFeeAllOfRate) SetFast(v UtxoFeeBasePrice)`

SetFast sets Fast field to given value.

### HasFast

`func (o *UtxoFeeAllOfRate) HasFast() bool`

HasFast returns a boolean if a field has been set.

### GetCustomized

`func (o *UtxoFeeAllOfRate) GetCustomized() UtxoFeeBasePrice`

GetCustomized returns the Customized field if non-nil, zero value otherwise.

### GetCustomizedOk

`func (o *UtxoFeeAllOfRate) GetCustomizedOk() (*UtxoFeeBasePrice, bool)`

GetCustomizedOk returns a tuple with the Customized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomized

`func (o *UtxoFeeAllOfRate) SetCustomized(v UtxoFeeBasePrice)`

SetCustomized sets Customized field to given value.

### HasCustomized

`func (o *UtxoFeeAllOfRate) HasCustomized() bool`

HasCustomized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


