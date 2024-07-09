# EvmEip1559FeePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | Pointer to **string** | The token ID of the transaction fee. | [optional] 
**Slow** | Pointer to [**EvmEip1559FeeBasePrice**](EvmEip1559FeeBasePrice.md) |  | [optional] 
**Recommended** | [**EvmEip1559FeeBasePrice**](EvmEip1559FeeBasePrice.md) |  | 
**Fast** | Pointer to [**EvmEip1559FeeBasePrice**](EvmEip1559FeeBasePrice.md) |  | [optional] 

## Methods

### NewEvmEip1559FeePrice

`func NewEvmEip1559FeePrice(feeType FeeType, recommended EvmEip1559FeeBasePrice, ) *EvmEip1559FeePrice`

NewEvmEip1559FeePrice instantiates a new EvmEip1559FeePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvmEip1559FeePriceWithDefaults

`func NewEvmEip1559FeePriceWithDefaults() *EvmEip1559FeePrice`

NewEvmEip1559FeePriceWithDefaults instantiates a new EvmEip1559FeePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *EvmEip1559FeePrice) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EvmEip1559FeePrice) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EvmEip1559FeePrice) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *EvmEip1559FeePrice) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EvmEip1559FeePrice) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EvmEip1559FeePrice) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *EvmEip1559FeePrice) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetSlow

`func (o *EvmEip1559FeePrice) GetSlow() EvmEip1559FeeBasePrice`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EvmEip1559FeePrice) GetSlowOk() (*EvmEip1559FeeBasePrice, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EvmEip1559FeePrice) SetSlow(v EvmEip1559FeeBasePrice)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EvmEip1559FeePrice) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *EvmEip1559FeePrice) GetRecommended() EvmEip1559FeeBasePrice`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *EvmEip1559FeePrice) GetRecommendedOk() (*EvmEip1559FeeBasePrice, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *EvmEip1559FeePrice) SetRecommended(v EvmEip1559FeeBasePrice)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *EvmEip1559FeePrice) GetFast() EvmEip1559FeeBasePrice`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EvmEip1559FeePrice) GetFastOk() (*EvmEip1559FeeBasePrice, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EvmEip1559FeePrice) SetFast(v EvmEip1559FeeBasePrice)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EvmEip1559FeePrice) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


