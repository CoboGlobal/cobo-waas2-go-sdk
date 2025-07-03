# EstimatedFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 
**IsLoop** | Pointer to **bool** | Whether the transaction was executed as a [Cobo Loop](https://manuals.cobo.com/en/portal/custodial-wallets/cobo-loop) transfer. - &#x60;true&#x60;: The transaction was executed as a Cobo Loop transfer. - &#x60;false&#x60;: The transaction was not executed as a Cobo Loop transfer.  | [optional] 
**FeeAmount** | **string** | The transaction fee that you need to pay for the transaction. | 
**Slow** | Pointer to [**EstimatedFILFeeSlow**](EstimatedFILFeeSlow.md) |  | [optional] 
**Recommended** | [**EstimatedFILFeeSlow**](EstimatedFILFeeSlow.md) |  | 
**Fast** | Pointer to [**EstimatedFILFeeSlow**](EstimatedFILFeeSlow.md) |  | [optional] 

## Methods

### NewEstimatedFee

`func NewEstimatedFee(feeType FeeType, tokenId string, feeAmount string, recommended EstimatedFILFeeSlow, ) *EstimatedFee`

NewEstimatedFee instantiates a new EstimatedFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedFeeWithDefaults

`func NewEstimatedFeeWithDefaults() *EstimatedFee`

NewEstimatedFeeWithDefaults instantiates a new EstimatedFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *EstimatedFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EstimatedFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EstimatedFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *EstimatedFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EstimatedFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EstimatedFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetIsLoop

`func (o *EstimatedFee) GetIsLoop() bool`

GetIsLoop returns the IsLoop field if non-nil, zero value otherwise.

### GetIsLoopOk

`func (o *EstimatedFee) GetIsLoopOk() (*bool, bool)`

GetIsLoopOk returns a tuple with the IsLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoop

`func (o *EstimatedFee) SetIsLoop(v bool)`

SetIsLoop sets IsLoop field to given value.

### HasIsLoop

`func (o *EstimatedFee) HasIsLoop() bool`

HasIsLoop returns a boolean if a field has been set.

### GetFeeAmount

`func (o *EstimatedFee) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *EstimatedFee) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *EstimatedFee) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetSlow

`func (o *EstimatedFee) GetSlow() EstimatedFILFeeSlow`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *EstimatedFee) GetSlowOk() (*EstimatedFILFeeSlow, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *EstimatedFee) SetSlow(v EstimatedFILFeeSlow)`

SetSlow sets Slow field to given value.

### HasSlow

`func (o *EstimatedFee) HasSlow() bool`

HasSlow returns a boolean if a field has been set.

### GetRecommended

`func (o *EstimatedFee) GetRecommended() EstimatedFILFeeSlow`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *EstimatedFee) GetRecommendedOk() (*EstimatedFILFeeSlow, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *EstimatedFee) SetRecommended(v EstimatedFILFeeSlow)`

SetRecommended sets Recommended field to given value.


### GetFast

`func (o *EstimatedFee) GetFast() EstimatedFILFeeSlow`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *EstimatedFee) GetFastOk() (*EstimatedFILFeeSlow, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *EstimatedFee) SetFast(v EstimatedFILFeeSlow)`

SetFast sets Fast field to given value.

### HasFast

`func (o *EstimatedFee) HasFast() bool`

HasFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


