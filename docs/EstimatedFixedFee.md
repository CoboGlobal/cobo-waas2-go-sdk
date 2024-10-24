# EstimatedFixedFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeAmount** | **string** | The fee that you need to pay for the transaction. | 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 
**IsLoop** | Pointer to **bool** | Whether the transaction was executed as a [Cobo Loop](https://manuals.cobo.com/en/portal/custodial-wallets/cobo-loop) transfer. - &#x60;true&#x60;: The transaction was executed as a Cobo Loop transfer. - &#x60;false&#x60;: The transaction was not executed as a Cobo Loop transfer.  | [optional] 

## Methods

### NewEstimatedFixedFee

`func NewEstimatedFixedFee(feeAmount string, feeType FeeType, tokenId string, ) *EstimatedFixedFee`

NewEstimatedFixedFee instantiates a new EstimatedFixedFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedFixedFeeWithDefaults

`func NewEstimatedFixedFeeWithDefaults() *EstimatedFixedFee`

NewEstimatedFixedFeeWithDefaults instantiates a new EstimatedFixedFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeAmount

`func (o *EstimatedFixedFee) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *EstimatedFixedFee) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *EstimatedFixedFee) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetFeeType

`func (o *EstimatedFixedFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *EstimatedFixedFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *EstimatedFixedFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *EstimatedFixedFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EstimatedFixedFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EstimatedFixedFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetIsLoop

`func (o *EstimatedFixedFee) GetIsLoop() bool`

GetIsLoop returns the IsLoop field if non-nil, zero value otherwise.

### GetIsLoopOk

`func (o *EstimatedFixedFee) GetIsLoopOk() (*bool, bool)`

GetIsLoopOk returns a tuple with the IsLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoop

`func (o *EstimatedFixedFee) SetIsLoop(v bool)`

SetIsLoop sets IsLoop field to given value.

### HasIsLoop

`func (o *EstimatedFixedFee) HasIsLoop() bool`

HasIsLoop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


