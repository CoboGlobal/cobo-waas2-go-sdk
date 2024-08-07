# KeyShareHolderGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyShareHolderGroupId** | Pointer to **string** | The key share holder group ID. | [optional] 
**Type** | Pointer to [**KeyShareHolderGroupType**](KeyShareHolderGroupType.md) |  | [optional] 
**TssKeyShareGroups** | Pointer to [**[]TSSGroups**](TSSGroups.md) |  | [optional] 
**KeyShareHolders** | Pointer to [**[]KeyShareHolder**](KeyShareHolder.md) |  | [optional] 
**Participants** | Pointer to **int32** | The number of key share holders in this key share holder group. | [optional] 
**Threshold** | Pointer to **int32** | The number of key share holders required to approve each operation in this key share holder group. | [optional] 
**Status** | Pointer to [**KeyShareHolderGroupStatus**](KeyShareHolderGroupStatus.md) |  | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The key share holder group&#39;s creation time in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewKeyShareHolderGroup

`func NewKeyShareHolderGroup() *KeyShareHolderGroup`

NewKeyShareHolderGroup instantiates a new KeyShareHolderGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyShareHolderGroupWithDefaults

`func NewKeyShareHolderGroupWithDefaults() *KeyShareHolderGroup`

NewKeyShareHolderGroupWithDefaults instantiates a new KeyShareHolderGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyShareHolderGroupId

`func (o *KeyShareHolderGroup) GetKeyShareHolderGroupId() string`

GetKeyShareHolderGroupId returns the KeyShareHolderGroupId field if non-nil, zero value otherwise.

### GetKeyShareHolderGroupIdOk

`func (o *KeyShareHolderGroup) GetKeyShareHolderGroupIdOk() (*string, bool)`

GetKeyShareHolderGroupIdOk returns a tuple with the KeyShareHolderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyShareHolderGroupId

`func (o *KeyShareHolderGroup) SetKeyShareHolderGroupId(v string)`

SetKeyShareHolderGroupId sets KeyShareHolderGroupId field to given value.

### HasKeyShareHolderGroupId

`func (o *KeyShareHolderGroup) HasKeyShareHolderGroupId() bool`

HasKeyShareHolderGroupId returns a boolean if a field has been set.

### GetType

`func (o *KeyShareHolderGroup) GetType() KeyShareHolderGroupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyShareHolderGroup) GetTypeOk() (*KeyShareHolderGroupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyShareHolderGroup) SetType(v KeyShareHolderGroupType)`

SetType sets Type field to given value.

### HasType

`func (o *KeyShareHolderGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTssKeyShareGroups

`func (o *KeyShareHolderGroup) GetTssKeyShareGroups() []TSSGroups`

GetTssKeyShareGroups returns the TssKeyShareGroups field if non-nil, zero value otherwise.

### GetTssKeyShareGroupsOk

`func (o *KeyShareHolderGroup) GetTssKeyShareGroupsOk() (*[]TSSGroups, bool)`

GetTssKeyShareGroupsOk returns a tuple with the TssKeyShareGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTssKeyShareGroups

`func (o *KeyShareHolderGroup) SetTssKeyShareGroups(v []TSSGroups)`

SetTssKeyShareGroups sets TssKeyShareGroups field to given value.

### HasTssKeyShareGroups

`func (o *KeyShareHolderGroup) HasTssKeyShareGroups() bool`

HasTssKeyShareGroups returns a boolean if a field has been set.

### GetKeyShareHolders

`func (o *KeyShareHolderGroup) GetKeyShareHolders() []KeyShareHolder`

GetKeyShareHolders returns the KeyShareHolders field if non-nil, zero value otherwise.

### GetKeyShareHoldersOk

`func (o *KeyShareHolderGroup) GetKeyShareHoldersOk() (*[]KeyShareHolder, bool)`

GetKeyShareHoldersOk returns a tuple with the KeyShareHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyShareHolders

`func (o *KeyShareHolderGroup) SetKeyShareHolders(v []KeyShareHolder)`

SetKeyShareHolders sets KeyShareHolders field to given value.

### HasKeyShareHolders

`func (o *KeyShareHolderGroup) HasKeyShareHolders() bool`

HasKeyShareHolders returns a boolean if a field has been set.

### GetParticipants

`func (o *KeyShareHolderGroup) GetParticipants() int32`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *KeyShareHolderGroup) GetParticipantsOk() (*int32, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *KeyShareHolderGroup) SetParticipants(v int32)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *KeyShareHolderGroup) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetThreshold

`func (o *KeyShareHolderGroup) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *KeyShareHolderGroup) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *KeyShareHolderGroup) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *KeyShareHolderGroup) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetStatus

`func (o *KeyShareHolderGroup) GetStatus() KeyShareHolderGroupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyShareHolderGroup) GetStatusOk() (*KeyShareHolderGroupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyShareHolderGroup) SetStatus(v KeyShareHolderGroupStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeyShareHolderGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *KeyShareHolderGroup) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *KeyShareHolderGroup) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *KeyShareHolderGroup) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *KeyShareHolderGroup) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


