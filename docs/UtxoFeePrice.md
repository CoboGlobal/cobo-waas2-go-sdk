# UtxoFeePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | Pointer to **string** | The token ID of the transaction fee. | [optional] 
**Slow** | Pointer to [**UtxoFeeBasePrice**](UtxoFeeBasePrice.md) |  | [optional] 
**Recommended** | [**UtxoFeeBasePrice**](UtxoFeeBasePrice.md) |  | 
**Fast** | Pointer to [**UtxoFeeBasePrice**](UtxoFeeBasePrice.md) |  | [optional] 

## Methods

### NewUtxoFeePrice

`func NewUtxoFeePrice(feeType FeeType, recommended UtxoFeeBasePrice, ) *UtxoFeePrice`

NewUtxoFeePrice instantiates a new UtxoFeePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtxoFeePriceWithDefaults

`func NewUtxoFeePriceWithDefaults() *UtxoFeePrice`

NewUtxoFeePriceWithDefaults instantiates a new UtxoFeePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *UtxoFeePrice) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *UtxoFeePrice) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *UtxoFeePrice) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *UtxoFeePrice) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *UtxoFeePrice) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *UtxoFeePrice) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *UtxoFeePrice) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetSlow

`func (o *UtxoFeePrice) GetSlow() UtxoFeeBasePrice`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *UtxoFeePrice) GetSlowOk() (*UtxoFeeBasePrice, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *UtxoFeePrice) SetSlow(v UtxoFeeBasePrice)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *UtxoFeePrice) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *UtxoFeePrice) GetRecommended() UtxoFeeBasePrice`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *UtxoFeePrice) GetRecommendedOk() (*UtxoFeeBasePrice, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *UtxoFeePrice) SetRecommended(v UtxoFeeBasePrice)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *UtxoFeePrice) GetFast() UtxoFeeBasePrice`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *UtxoFeePrice) GetFastOk() (*UtxoFeeBasePrice, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *UtxoFeePrice) SetFast(v UtxoFeeBasePrice)`

SetFast sets Fast field to given value.

### HasFast

`func (o *UtxoFeePrice) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


