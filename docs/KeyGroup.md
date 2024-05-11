# KeyGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique id of the key group | [optional] 
**Type** | Pointer to [**KeyGroupType**](KeyGroupType.md) |  | [optional] 
**TssGroupIds** | Pointer to [**[]TSSGroupId**](TSSGroupId.md) |  | [optional] 
**KeyHolders** | Pointer to [**[]KeyHolder**](KeyHolder.md) |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**Threshold** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**KeyGroupStatus**](KeyGroupStatus.md) |  | [optional] 

## Methods

### NewKeyGroup

`func NewKeyGroup() *KeyGroup`

NewKeyGroup instantiates a new KeyGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyGroupWithDefaults

`func NewKeyGroupWithDefaults() *KeyGroup`

NewKeyGroupWithDefaults instantiates a new KeyGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *KeyGroup) GetType() KeyGroupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyGroup) GetTypeOk() (*KeyGroupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyGroup) SetType(v KeyGroupType)`

SetType sets Type field to given value.

### HasType

`func (o *KeyGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTssGroupIds

`func (o *KeyGroup) GetTssGroupIds() []TSSGroupId`

GetTssGroupIds returns the TssGroupIds field if non-nil, zero value otherwise.

### GetTssGroupIdsOk

`func (o *KeyGroup) GetTssGroupIdsOk() (*[]TSSGroupId, bool)`

GetTssGroupIdsOk returns a tuple with the TssGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTssGroupIds

`func (o *KeyGroup) SetTssGroupIds(v []TSSGroupId)`

SetTssGroupIds sets TssGroupIds field to given value.

### HasTssGroupIds

`func (o *KeyGroup) HasTssGroupIds() bool`

HasTssGroupIds returns a boolean if a field has been set.

### GetKeyHolders

`func (o *KeyGroup) GetKeyHolders() []KeyHolder`

GetKeyHolders returns the KeyHolders field if non-nil, zero value otherwise.

### GetKeyHoldersOk

`func (o *KeyGroup) GetKeyHoldersOk() (*[]KeyHolder, bool)`

GetKeyHoldersOk returns a tuple with the KeyHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyHolders

`func (o *KeyGroup) SetKeyHolders(v []KeyHolder)`

SetKeyHolders sets KeyHolders field to given value.

### HasKeyHolders

`func (o *KeyGroup) HasKeyHolders() bool`

HasKeyHolders returns a boolean if a field has been set.

### GetNodeCount

`func (o *KeyGroup) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *KeyGroup) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *KeyGroup) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *KeyGroup) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetThreshold

`func (o *KeyGroup) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *KeyGroup) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *KeyGroup) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *KeyGroup) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetStatus

`func (o *KeyGroup) GetStatus() KeyGroupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyGroup) GetStatusOk() (*KeyGroupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyGroup) SetStatus(v KeyGroupStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeyGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


