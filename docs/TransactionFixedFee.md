# TransactionFixedFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxFeeAmount** | Pointer to **string** | The maximum fee that you are willing to pay for the transaction. Provide the value without applying precision. The transaction will fail if the transaction fee exceeds the maximum fee. | [optional] 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | Pointer to **string** | The token used to pay the transaction fee. | [optional] 
**FeeUsed** | Pointer to **string** | The actually charged transaction fee. | [optional] 
**EstimatedFeeUsed** | Pointer to **string** | The estimated transaction fee. | [optional] 

## Methods

### NewTransactionFixedFee

`func NewTransactionFixedFee(feeType FeeType, ) *TransactionFixedFee`

NewTransactionFixedFee instantiates a new TransactionFixedFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFixedFeeWithDefaults

`func NewTransactionFixedFeeWithDefaults() *TransactionFixedFee`

NewTransactionFixedFeeWithDefaults instantiates a new TransactionFixedFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxFeeAmount

`func (o *TransactionFixedFee) GetMaxFeeAmount() string`

GetMaxFeeAmount returns the MaxFeeAmount field if non-nil, zero value otherwise.

### GetMaxFeeAmountOk

`func (o *TransactionFixedFee) GetMaxFeeAmountOk() (*string, bool)`

GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeAmount

`func (o *TransactionFixedFee) SetMaxFeeAmount(v string)`

SetMaxFeeAmount sets MaxFeeAmount field to given value.

### HasMaxFeeAmount

`func (o *TransactionFixedFee) HasMaxFeeAmount() bool`

HasMaxFeeAmount returns a boolean if a field has been set.

### GetFeeType

`func (o *TransactionFixedFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *TransactionFixedFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *TransactionFixedFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *TransactionFixedFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionFixedFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionFixedFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TransactionFixedFee) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetFeeUsed

`func (o *TransactionFixedFee) GetFeeUsed() string`

GetFeeUsed returns the FeeUsed field if non-nil, zero value otherwise.

### GetFeeUsedOk

`func (o *TransactionFixedFee) GetFeeUsedOk() (*string, bool)`

GetFeeUsedOk returns a tuple with the FeeUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeUsed

`func (o *TransactionFixedFee) SetFeeUsed(v string)`

SetFeeUsed sets FeeUsed field to given value.

### HasFeeUsed

`func (o *TransactionFixedFee) HasFeeUsed() bool`

HasFeeUsed returns a boolean if a field has been set.

### GetEstimatedFeeUsed

`func (o *TransactionFixedFee) GetEstimatedFeeUsed() string`

GetEstimatedFeeUsed returns the EstimatedFeeUsed field if non-nil, zero value otherwise.

### GetEstimatedFeeUsedOk

`func (o *TransactionFixedFee) GetEstimatedFeeUsedOk() (*string, bool)`

GetEstimatedFeeUsedOk returns a tuple with the EstimatedFeeUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedFeeUsed

`func (o *TransactionFixedFee) SetEstimatedFeeUsed(v string)`

SetEstimatedFeeUsed sets EstimatedFeeUsed field to given value.

### HasEstimatedFeeUsed

`func (o *TransactionFixedFee) HasEstimatedFeeUsed() bool`

HasEstimatedFeeUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


