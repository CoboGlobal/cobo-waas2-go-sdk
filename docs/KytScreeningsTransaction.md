# KytScreeningsTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The UUID of the transaction that was screened. | 
**TransactionType** | [**KytScreeningsTransactionType**](KytScreeningsTransactionType.md) |  | 
**ReviewStatus** | [**ReviewStatusType**](ReviewStatusType.md) |  | 
**FundsStatus** | [**FundsStatusType**](FundsStatusType.md) |  | 

## Methods

### NewKytScreeningsTransaction

`func NewKytScreeningsTransaction(transactionId string, transactionType KytScreeningsTransactionType, reviewStatus ReviewStatusType, fundsStatus FundsStatusType, ) *KytScreeningsTransaction`

NewKytScreeningsTransaction instantiates a new KytScreeningsTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKytScreeningsTransactionWithDefaults

`func NewKytScreeningsTransactionWithDefaults() *KytScreeningsTransaction`

NewKytScreeningsTransactionWithDefaults instantiates a new KytScreeningsTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *KytScreeningsTransaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *KytScreeningsTransaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *KytScreeningsTransaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTransactionType

`func (o *KytScreeningsTransaction) GetTransactionType() KytScreeningsTransactionType`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *KytScreeningsTransaction) GetTransactionTypeOk() (*KytScreeningsTransactionType, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *KytScreeningsTransaction) SetTransactionType(v KytScreeningsTransactionType)`

SetTransactionType sets TransactionType field to given value.


### GetReviewStatus

`func (o *KytScreeningsTransaction) GetReviewStatus() ReviewStatusType`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *KytScreeningsTransaction) GetReviewStatusOk() (*ReviewStatusType, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *KytScreeningsTransaction) SetReviewStatus(v ReviewStatusType)`

SetReviewStatus sets ReviewStatus field to given value.


### GetFundsStatus

`func (o *KytScreeningsTransaction) GetFundsStatus() FundsStatusType`

GetFundsStatus returns the FundsStatus field if non-nil, zero value otherwise.

### GetFundsStatusOk

`func (o *KytScreeningsTransaction) GetFundsStatusOk() (*FundsStatusType, bool)`

GetFundsStatusOk returns a tuple with the FundsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsStatus

`func (o *KytScreeningsTransaction) SetFundsStatus(v FundsStatusType)`

SetFundsStatus sets FundsStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


