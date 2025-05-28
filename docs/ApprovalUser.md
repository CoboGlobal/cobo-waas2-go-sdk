# ApprovalUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The approval user email. | 
**Name** | Pointer to **string** | The approval user name. | [optional] 
**Status** | [**ApprovalStatus**](ApprovalStatus.md) |  | 
**CreatedTimestamp** | **int64** | The time when the approval was created, in Unix timestamp format, measured in milliseconds. | 

## Methods

### NewApprovalUser

`func NewApprovalUser(email string, status ApprovalStatus, createdTimestamp int64, ) *ApprovalUser`

NewApprovalUser instantiates a new ApprovalUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalUserWithDefaults

`func NewApprovalUserWithDefaults() *ApprovalUser`

NewApprovalUserWithDefaults instantiates a new ApprovalUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ApprovalUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApprovalUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApprovalUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *ApprovalUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApprovalUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApprovalUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApprovalUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ApprovalUser) GetStatus() ApprovalStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalUser) GetStatusOk() (*ApprovalStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalUser) SetStatus(v ApprovalStatus)`

SetStatus sets Status field to given value.


### GetCreatedTimestamp

`func (o *ApprovalUser) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ApprovalUser) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ApprovalUser) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


