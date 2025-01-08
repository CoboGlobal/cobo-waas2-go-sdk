# SwapActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | Pointer to **string** | The unique identifier of the swap activity. | [optional] 
**Status** | Pointer to **string** | The status of the swap activity. | [optional] 
**WalletId** | Pointer to **string** | The unique identifier of the wallet. | [optional] 
**PayTokenId** | Pointer to **string** | The token symbol to swap from. | [optional] 
**ReceiveTokenId** | Pointer to **string** | The token symbol to swap to. | [optional] 
**PayAmount** | Pointer to **string** | The amount of tokens to bridge. | [optional] 
**ReceiveAmount** | Pointer to **string** | The amount of tokens to receive. | [optional] 
**FeeAmount** | Pointer to **string** | The amount of fee. | [optional] 
**Initiator** | Pointer to **NullableString** | The initiator of the swap activity. | [optional] 
**InitiatorType** | Pointer to [**TransactionInitiatorType**](TransactionInitiatorType.md) |  | [optional] 
**CreatedTimestamp** | Pointer to **int32** | The time when the swap activity was created, in Unix timestamp format, measured in milliseconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The time when the swap activity was last updated, in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewSwapActivity

`func NewSwapActivity() *SwapActivity`

NewSwapActivity instantiates a new SwapActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapActivityWithDefaults

`func NewSwapActivityWithDefaults() *SwapActivity`

NewSwapActivityWithDefaults instantiates a new SwapActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityId

`func (o *SwapActivity) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *SwapActivity) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *SwapActivity) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *SwapActivity) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetStatus

`func (o *SwapActivity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwapActivity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwapActivity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SwapActivity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWalletId

`func (o *SwapActivity) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *SwapActivity) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *SwapActivity) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *SwapActivity) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetPayTokenId

`func (o *SwapActivity) GetPayTokenId() string`

GetPayTokenId returns the PayTokenId field if non-nil, zero value otherwise.

### GetPayTokenIdOk

`func (o *SwapActivity) GetPayTokenIdOk() (*string, bool)`

GetPayTokenIdOk returns a tuple with the PayTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayTokenId

`func (o *SwapActivity) SetPayTokenId(v string)`

SetPayTokenId sets PayTokenId field to given value.

### HasPayTokenId

`func (o *SwapActivity) HasPayTokenId() bool`

HasPayTokenId returns a boolean if a field has been set.

### GetReceiveTokenId

`func (o *SwapActivity) GetReceiveTokenId() string`

GetReceiveTokenId returns the ReceiveTokenId field if non-nil, zero value otherwise.

### GetReceiveTokenIdOk

`func (o *SwapActivity) GetReceiveTokenIdOk() (*string, bool)`

GetReceiveTokenIdOk returns a tuple with the ReceiveTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveTokenId

`func (o *SwapActivity) SetReceiveTokenId(v string)`

SetReceiveTokenId sets ReceiveTokenId field to given value.

### HasReceiveTokenId

`func (o *SwapActivity) HasReceiveTokenId() bool`

HasReceiveTokenId returns a boolean if a field has been set.

### GetPayAmount

`func (o *SwapActivity) GetPayAmount() string`

GetPayAmount returns the PayAmount field if non-nil, zero value otherwise.

### GetPayAmountOk

`func (o *SwapActivity) GetPayAmountOk() (*string, bool)`

GetPayAmountOk returns a tuple with the PayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayAmount

`func (o *SwapActivity) SetPayAmount(v string)`

SetPayAmount sets PayAmount field to given value.

### HasPayAmount

`func (o *SwapActivity) HasPayAmount() bool`

HasPayAmount returns a boolean if a field has been set.

### GetReceiveAmount

`func (o *SwapActivity) GetReceiveAmount() string`

GetReceiveAmount returns the ReceiveAmount field if non-nil, zero value otherwise.

### GetReceiveAmountOk

`func (o *SwapActivity) GetReceiveAmountOk() (*string, bool)`

GetReceiveAmountOk returns a tuple with the ReceiveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveAmount

`func (o *SwapActivity) SetReceiveAmount(v string)`

SetReceiveAmount sets ReceiveAmount field to given value.

### HasReceiveAmount

`func (o *SwapActivity) HasReceiveAmount() bool`

HasReceiveAmount returns a boolean if a field has been set.

### GetFeeAmount

`func (o *SwapActivity) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *SwapActivity) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *SwapActivity) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *SwapActivity) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetInitiator

`func (o *SwapActivity) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *SwapActivity) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *SwapActivity) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *SwapActivity) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### SetInitiatorNil

`func (o *SwapActivity) SetInitiatorNil(b bool)`

 SetInitiatorNil sets the value for Initiator to be an explicit nil

### UnsetInitiator
`func (o *SwapActivity) UnsetInitiator()`

UnsetInitiator ensures that no value is present for Initiator, not even an explicit nil
### GetInitiatorType

`func (o *SwapActivity) GetInitiatorType() TransactionInitiatorType`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *SwapActivity) GetInitiatorTypeOk() (*TransactionInitiatorType, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *SwapActivity) SetInitiatorType(v TransactionInitiatorType)`

SetInitiatorType sets InitiatorType field to given value.

### HasInitiatorType

`func (o *SwapActivity) HasInitiatorType() bool`

HasInitiatorType returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *SwapActivity) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *SwapActivity) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *SwapActivity) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *SwapActivity) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *SwapActivity) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *SwapActivity) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *SwapActivity) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *SwapActivity) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


