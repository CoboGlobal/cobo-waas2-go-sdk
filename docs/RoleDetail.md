# RoleDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ApprovalTransactionResult**](ApprovalTransactionResult.md) |  | [optional] 
**ReviewThreshold** | Pointer to **int32** | The threshold for the transaction approval. | [optional] 
**Initiator** | Pointer to **string** | The initiator of the transaction. | [optional] 
**IsUpgraded** | Pointer to **bool** | Indicates whether the transaction approval has been upgraded. | [optional] 
**CompleteTime** | Pointer to **string** | Time to complete the review. | [optional] 
**UserDetails** | Pointer to [**[]ApprovalUserDetail**](ApprovalUserDetail.md) |  | [optional] 

## Methods

### NewRoleDetail

`func NewRoleDetail() *RoleDetail`

NewRoleDetail instantiates a new RoleDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleDetailWithDefaults

`func NewRoleDetailWithDefaults() *RoleDetail`

NewRoleDetailWithDefaults instantiates a new RoleDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *RoleDetail) GetResult() ApprovalTransactionResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RoleDetail) GetResultOk() (*ApprovalTransactionResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RoleDetail) SetResult(v ApprovalTransactionResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *RoleDetail) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetReviewThreshold

`func (o *RoleDetail) GetReviewThreshold() int32`

GetReviewThreshold returns the ReviewThreshold field if non-nil, zero value otherwise.

### GetReviewThresholdOk

`func (o *RoleDetail) GetReviewThresholdOk() (*int32, bool)`

GetReviewThresholdOk returns a tuple with the ReviewThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewThreshold

`func (o *RoleDetail) SetReviewThreshold(v int32)`

SetReviewThreshold sets ReviewThreshold field to given value.

### HasReviewThreshold

`func (o *RoleDetail) HasReviewThreshold() bool`

HasReviewThreshold returns a boolean if a field has been set.

### GetInitiator

`func (o *RoleDetail) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *RoleDetail) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *RoleDetail) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *RoleDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetIsUpgraded

`func (o *RoleDetail) GetIsUpgraded() bool`

GetIsUpgraded returns the IsUpgraded field if non-nil, zero value otherwise.

### GetIsUpgradedOk

`func (o *RoleDetail) GetIsUpgradedOk() (*bool, bool)`

GetIsUpgradedOk returns a tuple with the IsUpgraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgraded

`func (o *RoleDetail) SetIsUpgraded(v bool)`

SetIsUpgraded sets IsUpgraded field to given value.

### HasIsUpgraded

`func (o *RoleDetail) HasIsUpgraded() bool`

HasIsUpgraded returns a boolean if a field has been set.

### GetCompleteTime

`func (o *RoleDetail) GetCompleteTime() string`

GetCompleteTime returns the CompleteTime field if non-nil, zero value otherwise.

### GetCompleteTimeOk

`func (o *RoleDetail) GetCompleteTimeOk() (*string, bool)`

GetCompleteTimeOk returns a tuple with the CompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteTime

`func (o *RoleDetail) SetCompleteTime(v string)`

SetCompleteTime sets CompleteTime field to given value.

### HasCompleteTime

`func (o *RoleDetail) HasCompleteTime() bool`

HasCompleteTime returns a boolean if a field has been set.

### GetUserDetails

`func (o *RoleDetail) GetUserDetails() []ApprovalUserDetail`

GetUserDetails returns the UserDetails field if non-nil, zero value otherwise.

### GetUserDetailsOk

`func (o *RoleDetail) GetUserDetailsOk() (*[]ApprovalUserDetail, bool)`

GetUserDetailsOk returns a tuple with the UserDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDetails

`func (o *RoleDetail) SetUserDetails(v []ApprovalUserDetail)`

SetUserDetails sets UserDetails field to given value.

### HasUserDetails

`func (o *RoleDetail) HasUserDetails() bool`

HasUserDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


