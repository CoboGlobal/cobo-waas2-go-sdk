# PaymentBankWithdrawalTimelineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**PaymentBankWithdrawalStatus**](PaymentBankWithdrawalStatus.md) |  | 
**Timestamp** | **int32** | The time when the bank withdrawal entered this status, represented as a UNIX timestamp in seconds. | 

## Methods

### NewPaymentBankWithdrawalTimelineItem

`func NewPaymentBankWithdrawalTimelineItem(status PaymentBankWithdrawalStatus, timestamp int32, ) *PaymentBankWithdrawalTimelineItem`

NewPaymentBankWithdrawalTimelineItem instantiates a new PaymentBankWithdrawalTimelineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentBankWithdrawalTimelineItemWithDefaults

`func NewPaymentBankWithdrawalTimelineItemWithDefaults() *PaymentBankWithdrawalTimelineItem`

NewPaymentBankWithdrawalTimelineItemWithDefaults instantiates a new PaymentBankWithdrawalTimelineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PaymentBankWithdrawalTimelineItem) GetStatus() PaymentBankWithdrawalStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentBankWithdrawalTimelineItem) GetStatusOk() (*PaymentBankWithdrawalStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentBankWithdrawalTimelineItem) SetStatus(v PaymentBankWithdrawalStatus)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *PaymentBankWithdrawalTimelineItem) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PaymentBankWithdrawalTimelineItem) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PaymentBankWithdrawalTimelineItem) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


