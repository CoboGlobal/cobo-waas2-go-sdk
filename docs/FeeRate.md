# FeeRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 
**FeeAmount** | Pointer to **string** | The transaction fee that you need to pay for the transaction. | [optional] 
**Slow** | Pointer to [**UtxoFeeBasePrice**](UtxoFeeBasePrice.md) |  | [optional] 
**Recommended** | [**UtxoFeeBasePrice**](UtxoFeeBasePrice.md) |  | 
**Fast** | Pointer to [**UtxoFeeBasePrice**](UtxoFeeBasePrice.md) |  | [optional] 

## Methods

### NewFeeRate

`func NewFeeRate(feeType FeeType, tokenId string, recommended UtxoFeeBasePrice, ) *FeeRate`

NewFeeRate instantiates a new FeeRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeRateWithDefaults

`func NewFeeRateWithDefaults() *FeeRate`

NewFeeRateWithDefaults instantiates a new FeeRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *FeeRate) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *FeeRate) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *FeeRate) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *FeeRate) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *FeeRate) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *FeeRate) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetFeeAmount

`func (o *FeeRate) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *FeeRate) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *FeeRate) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *FeeRate) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetSlow

`func (o *FeeRate) GetSlow() UtxoFeeBasePrice`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *FeeRate) GetSlowOk() (*UtxoFeeBasePrice, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *FeeRate) SetSlow(v UtxoFeeBasePrice)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *FeeRate) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *FeeRate) GetRecommended() UtxoFeeBasePrice`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *FeeRate) GetRecommendedOk() (*UtxoFeeBasePrice, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *FeeRate) SetRecommended(v UtxoFeeBasePrice)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *FeeRate) GetFast() UtxoFeeBasePrice`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *FeeRate) GetFastOk() (*UtxoFeeBasePrice, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *FeeRate) SetFast(v UtxoFeeBasePrice)`

SetFast sets Fast field to given value.

### HasFast

`func (o *FeeRate) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


