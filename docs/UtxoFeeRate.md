# UtxoFeeRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 
**Slow** | Pointer to [**UtxoFeeBasePrice**](UtxoFeeBasePrice.md) |  | [optional] 
**Recommended** | [**UtxoFeeBasePrice**](UtxoFeeBasePrice.md) |  | 
**Fast** | Pointer to [**UtxoFeeBasePrice**](UtxoFeeBasePrice.md) |  | [optional] 

## Methods

### NewUtxoFeeRate

`func NewUtxoFeeRate(feeType FeeType, tokenId string, recommended UtxoFeeBasePrice, ) *UtxoFeeRate`

NewUtxoFeeRate instantiates a new UtxoFeeRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtxoFeeRateWithDefaults

`func NewUtxoFeeRateWithDefaults() *UtxoFeeRate`

NewUtxoFeeRateWithDefaults instantiates a new UtxoFeeRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *UtxoFeeRate) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *UtxoFeeRate) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *UtxoFeeRate) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *UtxoFeeRate) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *UtxoFeeRate) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *UtxoFeeRate) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetSlow

`func (o *UtxoFeeRate) GetSlow() UtxoFeeBasePrice`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *UtxoFeeRate) GetSlowOk() (*UtxoFeeBasePrice, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *UtxoFeeRate) SetSlow(v UtxoFeeBasePrice)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *UtxoFeeRate) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *UtxoFeeRate) GetRecommended() UtxoFeeBasePrice`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *UtxoFeeRate) GetRecommendedOk() (*UtxoFeeBasePrice, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *UtxoFeeRate) SetRecommended(v UtxoFeeBasePrice)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *UtxoFeeRate) GetFast() UtxoFeeBasePrice`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *UtxoFeeRate) GetFastOk() (*UtxoFeeBasePrice, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *UtxoFeeRate) SetFast(v UtxoFeeBasePrice)`

SetFast sets Fast field to given value.

### HasFast

`func (o *UtxoFeeRate) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


