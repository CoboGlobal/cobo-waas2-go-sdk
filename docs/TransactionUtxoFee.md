# TransactionUtxoFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeRate** | **string** | The fee rate in sat/vByte. The fee rate represents the satoshis you are willing to pay for each byte of data that your transaction will consume on the blockchain. | 
**MaxFeeAmount** | **string** | The maximum fee that you are willing to pay for the transaction. The transaction will fail if the transaction fee exceeds the maximum fee. | 
**FeeType** | [**FeeType**](FeeType.md) |  | [default to FEETYPE_EVM_EIP_1559]
**FeeUsed** | Pointer to **string** | The transaction fee. | [optional] 

## Methods

### NewTransactionUtxoFee

`func NewTransactionUtxoFee(feeRate string, maxFeeAmount string, feeType FeeType, ) *TransactionUtxoFee`

NewTransactionUtxoFee instantiates a new TransactionUtxoFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionUtxoFeeWithDefaults

`func NewTransactionUtxoFeeWithDefaults() *TransactionUtxoFee`

NewTransactionUtxoFeeWithDefaults instantiates a new TransactionUtxoFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeRate

`func (o *TransactionUtxoFee) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *TransactionUtxoFee) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *TransactionUtxoFee) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetMaxFeeAmount

`func (o *TransactionUtxoFee) GetMaxFeeAmount() string`

GetMaxFeeAmount returns the MaxFeeAmount field if non-nil, zero value otherwise.

### GetMaxFeeAmountOk

`func (o *TransactionUtxoFee) GetMaxFeeAmountOk() (*string, bool)`

GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeAmount

`func (o *TransactionUtxoFee) SetMaxFeeAmount(v string)`

SetMaxFeeAmount sets MaxFeeAmount field to given value.


### GetFeeType

`func (o *TransactionUtxoFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *TransactionUtxoFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *TransactionUtxoFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.


### GetFeeUsed

`func (o *TransactionUtxoFee) GetFeeUsed() string`

GetFeeUsed returns the FeeUsed field if non-nil, zero value otherwise.

### GetFeeUsedOk

`func (o *TransactionUtxoFee) GetFeeUsedOk() (*string, bool)`

GetFeeUsedOk returns a tuple with the FeeUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeUsed

`func (o *TransactionUtxoFee) SetFeeUsed(v string)`

SetFeeUsed sets FeeUsed field to given value.

### HasFeeUsed

`func (o *TransactionUtxoFee) HasFeeUsed() bool`

HasFeeUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


