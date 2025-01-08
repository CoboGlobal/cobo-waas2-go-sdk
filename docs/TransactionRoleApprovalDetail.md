# TransactionRoleApprovalDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**TransactionApprovalResult**](TransactionApprovalResult.md) |  | [optional] 
**ReviewThreshold** | Pointer to **int32** | The threshold for passing the review of this role. | [optional] 
**Initiator** | Pointer to **string** | The initiator of the transaction. | [optional] 
**CompleteTime** | Pointer to **string** | Time to complete the review. | [optional] 
**UserDetails** | Pointer to [**[]TransactionUserApprovalDetail**](TransactionUserApprovalDetail.md) |  | [optional] 

## Methods

### NewTransactionRoleApprovalDetail

`func NewTransactionRoleApprovalDetail() *TransactionRoleApprovalDetail`

NewTransactionRoleApprovalDetail instantiates a new TransactionRoleApprovalDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRoleApprovalDetailWithDefaults

`func NewTransactionRoleApprovalDetailWithDefaults() *TransactionRoleApprovalDetail`

NewTransactionRoleApprovalDetailWithDefaults instantiates a new TransactionRoleApprovalDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *TransactionRoleApprovalDetail) GetResult() TransactionApprovalResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TransactionRoleApprovalDetail) GetResultOk() (*TransactionApprovalResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TransactionRoleApprovalDetail) SetResult(v TransactionApprovalResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *TransactionRoleApprovalDetail) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetReviewThreshold

`func (o *TransactionRoleApprovalDetail) GetReviewThreshold() int32`

GetReviewThreshold returns the ReviewThreshold field if non-nil, zero value otherwise.

### GetReviewThresholdOk

`func (o *TransactionRoleApprovalDetail) GetReviewThresholdOk() (*int32, bool)`

GetReviewThresholdOk returns a tuple with the ReviewThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewThreshold

`func (o *TransactionRoleApprovalDetail) SetReviewThreshold(v int32)`

SetReviewThreshold sets ReviewThreshold field to given value.

### HasReviewThreshold

`func (o *TransactionRoleApprovalDetail) HasReviewThreshold() bool`

HasReviewThreshold returns a boolean if a field has been set.

### GetInitiator

`func (o *TransactionRoleApprovalDetail) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *TransactionRoleApprovalDetail) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *TransactionRoleApprovalDetail) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *TransactionRoleApprovalDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetCompleteTime

`func (o *TransactionRoleApprovalDetail) GetCompleteTime() string`

GetCompleteTime returns the CompleteTime field if non-nil, zero value otherwise.

### GetCompleteTimeOk

`func (o *TransactionRoleApprovalDetail) GetCompleteTimeOk() (*string, bool)`

GetCompleteTimeOk returns a tuple with the CompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteTime

`func (o *TransactionRoleApprovalDetail) SetCompleteTime(v string)`

SetCompleteTime sets CompleteTime field to given value.

### HasCompleteTime

`func (o *TransactionRoleApprovalDetail) HasCompleteTime() bool`

HasCompleteTime returns a boolean if a field has been set.

### GetUserDetails

`func (o *TransactionRoleApprovalDetail) GetUserDetails() []TransactionUserApprovalDetail`

GetUserDetails returns the UserDetails field if non-nil, zero value otherwise.

### GetUserDetailsOk

`func (o *TransactionRoleApprovalDetail) GetUserDetailsOk() (*[]TransactionUserApprovalDetail, bool)`

GetUserDetailsOk returns a tuple with the UserDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDetails

`func (o *TransactionRoleApprovalDetail) SetUserDetails(v []TransactionUserApprovalDetail)`

SetUserDetails sets UserDetails field to given value.

### HasUserDetails

`func (o *TransactionRoleApprovalDetail) HasUserDetails() bool`

HasUserDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


