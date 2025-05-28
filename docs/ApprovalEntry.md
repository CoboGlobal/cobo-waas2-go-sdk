# ApprovalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApprovalStatus**](ApprovalStatus.md) |  | 
**CreatedTimestamp** | **int64** | The time when the approval was created, in Unix timestamp format, measured in milliseconds. | 
**ApprovalUsers** | Pointer to [**[]ApprovalUser**](ApprovalUser.md) |  | [optional] 

## Methods

### NewApprovalEntry

`func NewApprovalEntry(status ApprovalStatus, createdTimestamp int64, ) *ApprovalEntry`

NewApprovalEntry instantiates a new ApprovalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalEntryWithDefaults

`func NewApprovalEntryWithDefaults() *ApprovalEntry`

NewApprovalEntryWithDefaults instantiates a new ApprovalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApprovalEntry) GetStatus() ApprovalStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalEntry) GetStatusOk() (*ApprovalStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalEntry) SetStatus(v ApprovalStatus)`

SetStatus sets Status field to given value.


### GetCreatedTimestamp

`func (o *ApprovalEntry) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ApprovalEntry) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ApprovalEntry) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetApprovalUsers

`func (o *ApprovalEntry) GetApprovalUsers() []ApprovalUser`

GetApprovalUsers returns the ApprovalUsers field if non-nil, zero value otherwise.

### GetApprovalUsersOk

`func (o *ApprovalEntry) GetApprovalUsersOk() (*[]ApprovalUser, bool)`

GetApprovalUsersOk returns a tuple with the ApprovalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUsers

`func (o *ApprovalEntry) SetApprovalUsers(v []ApprovalUser)`

SetApprovalUsers sets ApprovalUsers field to given value.

### HasApprovalUsers

`func (o *ApprovalEntry) HasApprovalUsers() bool`

HasApprovalUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


