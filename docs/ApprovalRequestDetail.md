# ApprovalRequestDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalId** | Pointer to **string** |  | [optional] 
**Initiator** | [**ApprovalUser**](ApprovalUser.md) |  | 
**ApprovedList** | Pointer to [**[]ApprovalEntry**](ApprovalEntry.md) |  | [optional] 
**Status** | [**ApprovalStatus**](ApprovalStatus.md) |  | 
**ModifiedTimestamp** | Pointer to **int64** | The time when the approval was modified, in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewApprovalRequestDetail

`func NewApprovalRequestDetail(initiator ApprovalUser, status ApprovalStatus, ) *ApprovalRequestDetail`

NewApprovalRequestDetail instantiates a new ApprovalRequestDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestDetailWithDefaults

`func NewApprovalRequestDetailWithDefaults() *ApprovalRequestDetail`

NewApprovalRequestDetailWithDefaults instantiates a new ApprovalRequestDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalId

`func (o *ApprovalRequestDetail) GetApprovalId() string`

GetApprovalId returns the ApprovalId field if non-nil, zero value otherwise.

### GetApprovalIdOk

`func (o *ApprovalRequestDetail) GetApprovalIdOk() (*string, bool)`

GetApprovalIdOk returns a tuple with the ApprovalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalId

`func (o *ApprovalRequestDetail) SetApprovalId(v string)`

SetApprovalId sets ApprovalId field to given value.

### HasApprovalId

`func (o *ApprovalRequestDetail) HasApprovalId() bool`

HasApprovalId returns a boolean if a field has been set.

### GetInitiator

`func (o *ApprovalRequestDetail) GetInitiator() ApprovalUser`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *ApprovalRequestDetail) GetInitiatorOk() (*ApprovalUser, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *ApprovalRequestDetail) SetInitiator(v ApprovalUser)`

SetInitiator sets Initiator field to given value.


### GetApprovedList

`func (o *ApprovalRequestDetail) GetApprovedList() []ApprovalEntry`

GetApprovedList returns the ApprovedList field if non-nil, zero value otherwise.

### GetApprovedListOk

`func (o *ApprovalRequestDetail) GetApprovedListOk() (*[]ApprovalEntry, bool)`

GetApprovedListOk returns a tuple with the ApprovedList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedList

`func (o *ApprovalRequestDetail) SetApprovedList(v []ApprovalEntry)`

SetApprovedList sets ApprovedList field to given value.

### HasApprovedList

`func (o *ApprovalRequestDetail) HasApprovedList() bool`

HasApprovedList returns a boolean if a field has been set.

### GetStatus

`func (o *ApprovalRequestDetail) GetStatus() ApprovalStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalRequestDetail) GetStatusOk() (*ApprovalStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalRequestDetail) SetStatus(v ApprovalStatus)`

SetStatus sets Status field to given value.


### GetModifiedTimestamp

`func (o *ApprovalRequestDetail) GetModifiedTimestamp() int64`

GetModifiedTimestamp returns the ModifiedTimestamp field if non-nil, zero value otherwise.

### GetModifiedTimestampOk

`func (o *ApprovalRequestDetail) GetModifiedTimestampOk() (*int64, bool)`

GetModifiedTimestampOk returns a tuple with the ModifiedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTimestamp

`func (o *ApprovalRequestDetail) SetModifiedTimestamp(v int64)`

SetModifiedTimestamp sets ModifiedTimestamp field to given value.

### HasModifiedTimestamp

`func (o *ApprovalRequestDetail) HasModifiedTimestamp() bool`

HasModifiedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


