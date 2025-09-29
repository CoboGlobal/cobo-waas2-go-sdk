# TransactionRequestSOLFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeUnitPrice** | **string** | The cost per compute unit. Transactions consume computational resources measured in compute units, and this price helps determine the cost of executing transactions, especially complex ones involving smart contracts. | 
**ComputeUnitLimit** | **string** | The maximum number of compute units allowed for a transaction. This limits the resources any single transaction can consume, preventing excessive resource usage that could impact network performance negatively. | 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**TokenId** | **string** | The token ID of the transaction fee. | 

## Methods

### NewTransactionRequestSOLFee

`func NewTransactionRequestSOLFee(computeUnitPrice string, computeUnitLimit string, feeType FeeType, tokenId string, ) *TransactionRequestSOLFee`

NewTransactionRequestSOLFee instantiates a new TransactionRequestSOLFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRequestSOLFeeWithDefaults

`func NewTransactionRequestSOLFeeWithDefaults() *TransactionRequestSOLFee`

NewTransactionRequestSOLFeeWithDefaults instantiates a new TransactionRequestSOLFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeUnitPrice

`func (o *TransactionRequestSOLFee) GetComputeUnitPrice() string`

GetComputeUnitPrice returns the ComputeUnitPrice field if non-nil, zero value otherwise.

### GetComputeUnitPriceOk

`func (o *TransactionRequestSOLFee) GetComputeUnitPriceOk() (*string, bool)`

GetComputeUnitPriceOk returns a tuple with the ComputeUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeUnitPrice

`func (o *TransactionRequestSOLFee) SetComputeUnitPrice(v string)`

SetComputeUnitPrice sets ComputeUnitPrice field to given value.


### GetComputeUnitLimit

`func (o *TransactionRequestSOLFee) GetComputeUnitLimit() string`

GetComputeUnitLimit returns the ComputeUnitLimit field if non-nil, zero value otherwise.

### GetComputeUnitLimitOk

`func (o *TransactionRequestSOLFee) GetComputeUnitLimitOk() (*string, bool)`

GetComputeUnitLimitOk returns a tuple with the ComputeUnitLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeUnitLimit

`func (o *TransactionRequestSOLFee) SetComputeUnitLimit(v string)`

SetComputeUnitLimit sets ComputeUnitLimit field to given value.


### GetFeeType

`func (o *TransactionRequestSOLFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *TransactionRequestSOLFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *TransactionRequestSOLFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetTokenId

`func (o *TransactionRequestSOLFee) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionRequestSOLFee) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionRequestSOLFee) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


