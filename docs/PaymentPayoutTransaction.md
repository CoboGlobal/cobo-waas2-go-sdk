# PaymentPayoutTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The WaaS transaction ID. | 
**Amount** | [**PaymentAssetAmount**](PaymentAssetAmount.md) | The transaction amount. | 
**TxHash** | **NullableString** | The blockchain transaction hash. | 
**Result** | [**NullablePaymentPayoutTransactionResult**](PaymentPayoutTransactionResult.md) |  | 
**FailedReason** | [**NullablePaymentFailedReason**](PaymentFailedReason.md) |  | 

## Methods

### NewPaymentPayoutTransaction

`func NewPaymentPayoutTransaction(transactionId string, amount PaymentAssetAmount, txHash NullableString, result NullablePaymentPayoutTransactionResult, failedReason NullablePaymentFailedReason, ) *PaymentPayoutTransaction`

NewPaymentPayoutTransaction instantiates a new PaymentPayoutTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPayoutTransactionWithDefaults

`func NewPaymentPayoutTransactionWithDefaults() *PaymentPayoutTransaction`

NewPaymentPayoutTransactionWithDefaults instantiates a new PaymentPayoutTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *PaymentPayoutTransaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *PaymentPayoutTransaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *PaymentPayoutTransaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetAmount

`func (o *PaymentPayoutTransaction) GetAmount() PaymentAssetAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentPayoutTransaction) GetAmountOk() (*PaymentAssetAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentPayoutTransaction) SetAmount(v PaymentAssetAmount)`

SetAmount sets Amount field to given value.


### GetTxHash

`func (o *PaymentPayoutTransaction) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *PaymentPayoutTransaction) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *PaymentPayoutTransaction) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### SetTxHashNil

`func (o *PaymentPayoutTransaction) SetTxHashNil(b bool)`

 SetTxHashNil sets the value for TxHash to be an explicit nil

### UnsetTxHash
`func (o *PaymentPayoutTransaction) UnsetTxHash()`

UnsetTxHash ensures that no value is present for TxHash, not even an explicit nil
### GetResult

`func (o *PaymentPayoutTransaction) GetResult() PaymentPayoutTransactionResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PaymentPayoutTransaction) GetResultOk() (*PaymentPayoutTransactionResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PaymentPayoutTransaction) SetResult(v PaymentPayoutTransactionResult)`

SetResult sets Result field to given value.


### SetResultNil

`func (o *PaymentPayoutTransaction) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *PaymentPayoutTransaction) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetFailedReason

`func (o *PaymentPayoutTransaction) GetFailedReason() PaymentFailedReason`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *PaymentPayoutTransaction) GetFailedReasonOk() (*PaymentFailedReason, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *PaymentPayoutTransaction) SetFailedReason(v PaymentFailedReason)`

SetFailedReason sets FailedReason field to given value.


### SetFailedReasonNil

`func (o *PaymentPayoutTransaction) SetFailedReasonNil(b bool)`

 SetFailedReasonNil sets the value for FailedReason to be an explicit nil

### UnsetFailedReason
`func (o *PaymentPayoutTransaction) UnsetFailedReason()`

UnsetFailedReason ensures that no value is present for FailedReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


