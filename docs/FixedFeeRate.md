# FixedFeeRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeAmount** | Pointer to **string** | The transaction fee that you need to pay for the transaction. | [optional] 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 

## Methods

### NewFixedFeeRate

`func NewFixedFeeRate(feeType FeeType, tokenId string, ) *FixedFeeRate`

NewFixedFeeRate instantiates a new FixedFeeRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedFeeRateWithDefaults

`func NewFixedFeeRateWithDefaults() *FixedFeeRate`

NewFixedFeeRateWithDefaults instantiates a new FixedFeeRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeAmount

`func (o *FixedFeeRate) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *FixedFeeRate) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *FixedFeeRate) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *FixedFeeRate) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetFeeType

`func (o *FixedFeeRate) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *FixedFeeRate) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *FixedFeeRate) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *FixedFeeRate) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *FixedFeeRate) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *FixedFeeRate) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


