# PayoutEstimatedFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 
**IsLoop** | Pointer to **bool** | Whether the transaction was executed as a [Cobo Loop](https://manuals.cobo.com/en/portal/custodial-wallets/cobo-loop) transfer. - &#x60;true&#x60;: The transaction was executed as a Cobo Loop transfer. - &#x60;false&#x60;: The transaction was not executed as a Cobo Loop transfer.  | [optional] 
**FeeAmount** | **string** | The transaction fee that you need to pay for the transaction. | 
**Slow** | Pointer to [**EstimatedUtxoFeeSlow**](EstimatedUtxoFeeSlow.md) |  | [optional] 
**Recommended** | [**EstimatedUtxoFeeSlow**](EstimatedUtxoFeeSlow.md) |  | 
**Fast** | Pointer to [**EstimatedUtxoFeeSlow**](EstimatedUtxoFeeSlow.md) |  | [optional] 

## Methods

### NewPayoutEstimatedFee

`func NewPayoutEstimatedFee(feeType FeeType, tokenId string, feeAmount string, recommended EstimatedUtxoFeeSlow, ) *PayoutEstimatedFee`

NewPayoutEstimatedFee instantiates a new PayoutEstimatedFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutEstimatedFeeWithDefaults

`func NewPayoutEstimatedFeeWithDefaults() *PayoutEstimatedFee`

NewPayoutEstimatedFeeWithDefaults instantiates a new PayoutEstimatedFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *PayoutEstimatedFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *PayoutEstimatedFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *PayoutEstimatedFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *PayoutEstimatedFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PayoutEstimatedFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PayoutEstimatedFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetIsLoop

`func (o *PayoutEstimatedFee) GetIsLoop() bool`

GetIsLoop returns the IsLoop field if non-nil, zero value otherwise.

### GetIsLoopOk

`func (o *PayoutEstimatedFee) GetIsLoopOk() (*bool, bool)`

GetIsLoopOk returns a tuple with the IsLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoop

`func (o *PayoutEstimatedFee) SetIsLoop(v bool)`

SetIsLoop sets IsLoop field to given value.

### HasIsLoop

`func (o *PayoutEstimatedFee) HasIsLoop() bool`

HasIsLoop returns a boolean if a field has been set.

### GetFeeAmount

`func (o *PayoutEstimatedFee) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *PayoutEstimatedFee) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *PayoutEstimatedFee) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetSlow

`func (o *PayoutEstimatedFee) GetSlow() EstimatedUtxoFeeSlow`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *PayoutEstimatedFee) GetSlowOk() (*EstimatedUtxoFeeSlow, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *PayoutEstimatedFee) SetSlow(v EstimatedUtxoFeeSlow)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *PayoutEstimatedFee) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *PayoutEstimatedFee) GetRecommended() EstimatedUtxoFeeSlow`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *PayoutEstimatedFee) GetRecommendedOk() (*EstimatedUtxoFeeSlow, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *PayoutEstimatedFee) SetRecommended(v EstimatedUtxoFeeSlow)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *PayoutEstimatedFee) GetFast() EstimatedUtxoFeeSlow`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *PayoutEstimatedFee) GetFastOk() (*EstimatedUtxoFeeSlow, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *PayoutEstimatedFee) SetFast(v EstimatedUtxoFeeSlow)`

SetFast sets Fast field to given value.

### HasFast

`func (o *PayoutEstimatedFee) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


