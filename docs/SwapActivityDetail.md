# SwapActivityDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | Pointer to **string** | The unique identifier of the swap activity. | [optional] 
**SwapType** | Pointer to [**SwapType**](SwapType.md) |  | [optional] 
**Status** | Pointer to [**SwapActivityStatus**](SwapActivityStatus.md) |  | [optional] 
**RequestId** | Pointer to **string** | The request id of the swap activity. | [optional] 
**WalletId** | Pointer to **string** | The unique identifier of the wallet. | [optional] 
**PayTokenId** | Pointer to **string** | The token ID to pay. | [optional] 
**ReceiveTokenId** | Pointer to **string** | The token ID to receive. | [optional] 
**PayAmount** | Pointer to **string** | The amount of tokens to bridge. | [optional] 
**ReceiveAmount** | Pointer to **string** | The amount of tokens to receive. | [optional] 
**FeeTokenId** | Pointer to **string** | The fee token ID. | [optional] 
**FeeAmount** | Pointer to **string** | The amount of fee. | [optional] 
**Initiator** | Pointer to **NullableString** | The initiator of the swap activity. | [optional] 
**InitiatorType** | Pointer to [**TransactionInitiatorType**](TransactionInitiatorType.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the swap activity. | [optional] 
**CreatedTimestamp** | Pointer to **int32** | The time when the swap activity was created, in Unix timestamp format, measured in milliseconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The time when the swap activity was last updated, in Unix timestamp format, measured in milliseconds. | [optional] 
**Timeline** | Pointer to [**[]SwapActivityTimeline**](SwapActivityTimeline.md) |  | [optional] 

## Methods

### NewSwapActivityDetail

`func NewSwapActivityDetail() *SwapActivityDetail`

NewSwapActivityDetail instantiates a new SwapActivityDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapActivityDetailWithDefaults

`func NewSwapActivityDetailWithDefaults() *SwapActivityDetail`

NewSwapActivityDetailWithDefaults instantiates a new SwapActivityDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityId

`func (o *SwapActivityDetail) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *SwapActivityDetail) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *SwapActivityDetail) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *SwapActivityDetail) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetSwapType

`func (o *SwapActivityDetail) GetSwapType() SwapType`

GetSwapType returns the SwapType field if non-nil, zero value otherwise.

### GetSwapTypeOk

`func (o *SwapActivityDetail) GetSwapTypeOk() (*SwapType, bool)`

GetSwapTypeOk returns a tuple with the SwapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapType

`func (o *SwapActivityDetail) SetSwapType(v SwapType)`

SetSwapType sets SwapType field to given value.

### HasSwapType

`func (o *SwapActivityDetail) HasSwapType() bool`

HasSwapType returns a boolean if a field has been set.

### GetStatus

`func (o *SwapActivityDetail) GetStatus() SwapActivityStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwapActivityDetail) GetStatusOk() (*SwapActivityStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwapActivityDetail) SetStatus(v SwapActivityStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SwapActivityDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequestId

`func (o *SwapActivityDetail) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SwapActivityDetail) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SwapActivityDetail) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SwapActivityDetail) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetWalletId

`func (o *SwapActivityDetail) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *SwapActivityDetail) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *SwapActivityDetail) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *SwapActivityDetail) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetPayTokenId

`func (o *SwapActivityDetail) GetPayTokenId() string`

GetPayTokenId returns the PayTokenId field if non-nil, zero value otherwise.

### GetPayTokenIdOk

`func (o *SwapActivityDetail) GetPayTokenIdOk() (*string, bool)`

GetPayTokenIdOk returns a tuple with the PayTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayTokenId

`func (o *SwapActivityDetail) SetPayTokenId(v string)`

SetPayTokenId sets PayTokenId field to given value.

### HasPayTokenId

`func (o *SwapActivityDetail) HasPayTokenId() bool`

HasPayTokenId returns a boolean if a field has been set.

### GetReceiveTokenId

`func (o *SwapActivityDetail) GetReceiveTokenId() string`

GetReceiveTokenId returns the ReceiveTokenId field if non-nil, zero value otherwise.

### GetReceiveTokenIdOk

`func (o *SwapActivityDetail) GetReceiveTokenIdOk() (*string, bool)`

GetReceiveTokenIdOk returns a tuple with the ReceiveTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveTokenId

`func (o *SwapActivityDetail) SetReceiveTokenId(v string)`

SetReceiveTokenId sets ReceiveTokenId field to given value.

### HasReceiveTokenId

`func (o *SwapActivityDetail) HasReceiveTokenId() bool`

HasReceiveTokenId returns a boolean if a field has been set.

### GetPayAmount

`func (o *SwapActivityDetail) GetPayAmount() string`

GetPayAmount returns the PayAmount field if non-nil, zero value otherwise.

### GetPayAmountOk

`func (o *SwapActivityDetail) GetPayAmountOk() (*string, bool)`

GetPayAmountOk returns a tuple with the PayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayAmount

`func (o *SwapActivityDetail) SetPayAmount(v string)`

SetPayAmount sets PayAmount field to given value.

### HasPayAmount

`func (o *SwapActivityDetail) HasPayAmount() bool`

HasPayAmount returns a boolean if a field has been set.

### GetReceiveAmount

`func (o *SwapActivityDetail) GetReceiveAmount() string`

GetReceiveAmount returns the ReceiveAmount field if non-nil, zero value otherwise.

### GetReceiveAmountOk

`func (o *SwapActivityDetail) GetReceiveAmountOk() (*string, bool)`

GetReceiveAmountOk returns a tuple with the ReceiveAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveAmount

`func (o *SwapActivityDetail) SetReceiveAmount(v string)`

SetReceiveAmount sets ReceiveAmount field to given value.

### HasReceiveAmount

`func (o *SwapActivityDetail) HasReceiveAmount() bool`

HasReceiveAmount returns a boolean if a field has been set.

### GetFeeTokenId

`func (o *SwapActivityDetail) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *SwapActivityDetail) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *SwapActivityDetail) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *SwapActivityDetail) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetFeeAmount

`func (o *SwapActivityDetail) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *SwapActivityDetail) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *SwapActivityDetail) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *SwapActivityDetail) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetInitiator

`func (o *SwapActivityDetail) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *SwapActivityDetail) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *SwapActivityDetail) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *SwapActivityDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### SetInitiatorNil

`func (o *SwapActivityDetail) SetInitiatorNil(b bool)`

 SetInitiatorNil sets the value for Initiator to be an explicit nil

### UnsetInitiator
`func (o *SwapActivityDetail) UnsetInitiator()`

UnsetInitiator ensures that no value is present for Initiator, not even an explicit nil
### GetInitiatorType

`func (o *SwapActivityDetail) GetInitiatorType() TransactionInitiatorType`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *SwapActivityDetail) GetInitiatorTypeOk() (*TransactionInitiatorType, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *SwapActivityDetail) SetInitiatorType(v TransactionInitiatorType)`

SetInitiatorType sets InitiatorType field to given value.

### HasInitiatorType

`func (o *SwapActivityDetail) HasInitiatorType() bool`

HasInitiatorType returns a boolean if a field has been set.

### GetDescription

`func (o *SwapActivityDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SwapActivityDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SwapActivityDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SwapActivityDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *SwapActivityDetail) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *SwapActivityDetail) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *SwapActivityDetail) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *SwapActivityDetail) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *SwapActivityDetail) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *SwapActivityDetail) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *SwapActivityDetail) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *SwapActivityDetail) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetTimeline

`func (o *SwapActivityDetail) GetTimeline() []SwapActivityTimeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *SwapActivityDetail) GetTimelineOk() (*[]SwapActivityTimeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *SwapActivityDetail) SetTimeline(v []SwapActivityTimeline)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *SwapActivityDetail) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


