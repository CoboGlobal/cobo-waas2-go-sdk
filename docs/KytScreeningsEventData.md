# KytScreeningsEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The transaction ID. | 
**TransactionType** | [**KytScreeningsTransactionType**](KytScreeningsTransactionType.md) |  | 
**ReviewStatus** | [**ReviewStatusType**](ReviewStatusType.md) |  | 
**FundsStatus** | [**FundsStatusType**](FundsStatusType.md) |  | 
**UpdatedTimestamp** | **int64** | The time when the kyt screening was updated, in Unix timestamp format, measured in milliseconds. | 

## Methods

### NewKytScreeningsEventData

`func NewKytScreeningsEventData(transactionId string, transactionType KytScreeningsTransactionType, reviewStatus ReviewStatusType, fundsStatus FundsStatusType, updatedTimestamp int64, ) *KytScreeningsEventData`

NewKytScreeningsEventData instantiates a new KytScreeningsEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKytScreeningsEventDataWithDefaults

`func NewKytScreeningsEventDataWithDefaults() *KytScreeningsEventData`

NewKytScreeningsEventDataWithDefaults instantiates a new KytScreeningsEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *KytScreeningsEventData) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *KytScreeningsEventData) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *KytScreeningsEventData) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTransactionType

`func (o *KytScreeningsEventData) GetTransactionType() KytScreeningsTransactionType`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *KytScreeningsEventData) GetTransactionTypeOk() (*KytScreeningsTransactionType, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *KytScreeningsEventData) SetTransactionType(v KytScreeningsTransactionType)`

SetTransactionType sets TransactionType field to given value.


### GetReviewStatus

`func (o *KytScreeningsEventData) GetReviewStatus() ReviewStatusType`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *KytScreeningsEventData) GetReviewStatusOk() (*ReviewStatusType, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *KytScreeningsEventData) SetReviewStatus(v ReviewStatusType)`

SetReviewStatus sets ReviewStatus field to given value.


### GetFundsStatus

`func (o *KytScreeningsEventData) GetFundsStatus() FundsStatusType`

GetFundsStatus returns the FundsStatus field if non-nil, zero value otherwise.

### GetFundsStatusOk

`func (o *KytScreeningsEventData) GetFundsStatusOk() (*FundsStatusType, bool)`

GetFundsStatusOk returns a tuple with the FundsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsStatus

`func (o *KytScreeningsEventData) SetFundsStatus(v FundsStatusType)`

SetFundsStatus sets FundsStatus field to given value.


### GetUpdatedTimestamp

`func (o *KytScreeningsEventData) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *KytScreeningsEventData) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *KytScreeningsEventData) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


