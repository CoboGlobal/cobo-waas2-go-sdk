# MPCProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** | The project ID. | [optional] 
**OrgId** | Pointer to **string** | The [organization](https://manuals.cobo.com/en/portal/organization/introduction) ID. | [optional] 
**Name** | Pointer to **string** | The project name. | [optional] 
**Participants** | Pointer to **int32** | The number of key share holders in the project. | [optional] 
**Threshold** | Pointer to **int32** | The number of key share holders required to sign an operation in the project. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The project&#39;s creation time in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewMPCProject

`func NewMPCProject() *MPCProject`

NewMPCProject instantiates a new MPCProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMPCProjectWithDefaults

`func NewMPCProjectWithDefaults() *MPCProject`

NewMPCProjectWithDefaults instantiates a new MPCProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *MPCProject) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MPCProject) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MPCProject) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *MPCProject) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetOrgId

`func (o *MPCProject) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *MPCProject) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *MPCProject) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *MPCProject) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetName

`func (o *MPCProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MPCProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MPCProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MPCProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParticipants

`func (o *MPCProject) GetParticipants() int32`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *MPCProject) GetParticipantsOk() (*int32, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *MPCProject) SetParticipants(v int32)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *MPCProject) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetThreshold

`func (o *MPCProject) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MPCProject) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MPCProject) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *MPCProject) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *MPCProject) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *MPCProject) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *MPCProject) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *MPCProject) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


