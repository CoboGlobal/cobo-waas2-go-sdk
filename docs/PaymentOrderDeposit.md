# PaymentOrderDeposit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The incoming Payment transaction ID. | 
**TxHash** | **NullableString** | The blockchain transaction hash. | 
**Amount** | [**PaymentAssetAmount**](PaymentAssetAmount.md) |  | 
**SourceAddress** | **NullableString** | The source address of the incoming transaction. | 
**Status** | [**PaymentPayInStatus**](PaymentPayInStatus.md) |  | 
**ScreeningMode** | [**NullablePaymentScreeningMode**](PaymentScreeningMode.md) |  | 
**IsLate** | **bool** | Whether the incoming transaction completed after the order reached a terminal status. | 
**Message** | **NullableString** | The action prompt for an &#x60;ActionRequired&#x60; status. | 
**FailedReason** | [**NullablePaymentFailedReason**](PaymentFailedReason.md) |  | 

## Methods

### NewPaymentOrderDeposit

`func NewPaymentOrderDeposit(transactionId string, txHash NullableString, amount PaymentAssetAmount, sourceAddress NullableString, status PaymentPayInStatus, screeningMode NullablePaymentScreeningMode, isLate bool, message NullableString, failedReason NullablePaymentFailedReason, ) *PaymentOrderDeposit`

NewPaymentOrderDeposit instantiates a new PaymentOrderDeposit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOrderDepositWithDefaults

`func NewPaymentOrderDepositWithDefaults() *PaymentOrderDeposit`

NewPaymentOrderDepositWithDefaults instantiates a new PaymentOrderDeposit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *PaymentOrderDeposit) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *PaymentOrderDeposit) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *PaymentOrderDeposit) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTxHash

`func (o *PaymentOrderDeposit) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *PaymentOrderDeposit) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *PaymentOrderDeposit) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### SetTxHashNil

`func (o *PaymentOrderDeposit) SetTxHashNil(b bool)`

 SetTxHashNil sets the value for TxHash to be an explicit nil

### UnsetTxHash
`func (o *PaymentOrderDeposit) UnsetTxHash()`

UnsetTxHash ensures that no value is present for TxHash, not even an explicit nil
### GetAmount

`func (o *PaymentOrderDeposit) GetAmount() PaymentAssetAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentOrderDeposit) GetAmountOk() (*PaymentAssetAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentOrderDeposit) SetAmount(v PaymentAssetAmount)`

SetAmount sets Amount field to given value.


### GetSourceAddress

`func (o *PaymentOrderDeposit) GetSourceAddress() string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *PaymentOrderDeposit) GetSourceAddressOk() (*string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *PaymentOrderDeposit) SetSourceAddress(v string)`

SetSourceAddress sets SourceAddress field to given value.


### SetSourceAddressNil

`func (o *PaymentOrderDeposit) SetSourceAddressNil(b bool)`

 SetSourceAddressNil sets the value for SourceAddress to be an explicit nil

### UnsetSourceAddress
`func (o *PaymentOrderDeposit) UnsetSourceAddress()`

UnsetSourceAddress ensures that no value is present for SourceAddress, not even an explicit nil
### GetStatus

`func (o *PaymentOrderDeposit) GetStatus() PaymentPayInStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentOrderDeposit) GetStatusOk() (*PaymentPayInStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentOrderDeposit) SetStatus(v PaymentPayInStatus)`

SetStatus sets Status field to given value.


### GetScreeningMode

`func (o *PaymentOrderDeposit) GetScreeningMode() PaymentScreeningMode`

GetScreeningMode returns the ScreeningMode field if non-nil, zero value otherwise.

### GetScreeningModeOk

`func (o *PaymentOrderDeposit) GetScreeningModeOk() (*PaymentScreeningMode, bool)`

GetScreeningModeOk returns a tuple with the ScreeningMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreeningMode

`func (o *PaymentOrderDeposit) SetScreeningMode(v PaymentScreeningMode)`

SetScreeningMode sets ScreeningMode field to given value.


### SetScreeningModeNil

`func (o *PaymentOrderDeposit) SetScreeningModeNil(b bool)`

 SetScreeningModeNil sets the value for ScreeningMode to be an explicit nil

### UnsetScreeningMode
`func (o *PaymentOrderDeposit) UnsetScreeningMode()`

UnsetScreeningMode ensures that no value is present for ScreeningMode, not even an explicit nil
### GetIsLate

`func (o *PaymentOrderDeposit) GetIsLate() bool`

GetIsLate returns the IsLate field if non-nil, zero value otherwise.

### GetIsLateOk

`func (o *PaymentOrderDeposit) GetIsLateOk() (*bool, bool)`

GetIsLateOk returns a tuple with the IsLate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLate

`func (o *PaymentOrderDeposit) SetIsLate(v bool)`

SetIsLate sets IsLate field to given value.


### GetMessage

`func (o *PaymentOrderDeposit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentOrderDeposit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentOrderDeposit) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *PaymentOrderDeposit) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *PaymentOrderDeposit) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetFailedReason

`func (o *PaymentOrderDeposit) GetFailedReason() PaymentFailedReason`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *PaymentOrderDeposit) GetFailedReasonOk() (*PaymentFailedReason, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *PaymentOrderDeposit) SetFailedReason(v PaymentFailedReason)`

SetFailedReason sets FailedReason field to given value.


### SetFailedReasonNil

`func (o *PaymentOrderDeposit) SetFailedReasonNil(b bool)`

 SetFailedReasonNil sets the value for FailedReason to be an explicit nil

### UnsetFailedReason
`func (o *PaymentOrderDeposit) UnsetFailedReason()`

UnsetFailedReason ensures that no value is present for FailedReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


