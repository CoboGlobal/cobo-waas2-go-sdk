# MPCProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The project ID. | [optional] 
**OrgId** | Pointer to **string** | The organization ID. | [optional] 
**Name** | Pointer to **string** | The project name. | [optional] 
**NodeCount** | Pointer to **int32** | The number of key share holders in the project. | [optional] 
**Threshold** | Pointer to **int32** | The number of key share holders required to sign an operation in the project. | [optional] 
**CreateTimestamp** | Pointer to **int64** | The project&#39;s creation time in Unix timestamp format, measured in milliseconds. | [optional] 

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

### GetId

`func (o *MPCProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MPCProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MPCProject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MPCProject) HasId() bool`

HasId returns a boolean if a field has been set.

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

### GetNodeCount

`func (o *MPCProject) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *MPCProject) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *MPCProject) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *MPCProject) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

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

### GetCreateTimestamp

`func (o *MPCProject) GetCreateTimestamp() int64`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *MPCProject) GetCreateTimestampOk() (*int64, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *MPCProject) SetCreateTimestamp(v int64)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *MPCProject) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


