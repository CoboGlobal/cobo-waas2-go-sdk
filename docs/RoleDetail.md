# RoleDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ApprovalTransactionResult**](ApprovalTransactionResult.md) |  | [optional] 
**Threshold** | Pointer to **int32** | The threshold for the transaction approval. | [optional] 
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

### GetThreshold

`func (o *RoleDetail) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *RoleDetail) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *RoleDetail) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *RoleDetail) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

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


