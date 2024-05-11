# CreateKeyGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupType** | [**KeyGroupType**](KeyGroupType.md) |  | 
**NodeCount** | **int32** | The count of tss node of the key group | 
**Threshold** | **int32** | The threshold number of tss node required for signature | 
**KeyHolders** | [**[]CreateKeyGroupRequestKeyHoldersInner**](CreateKeyGroupRequestKeyHoldersInner.md) |  | 

## Methods

### NewCreateKeyGroupRequest

`func NewCreateKeyGroupRequest(groupType KeyGroupType, nodeCount int32, threshold int32, keyHolders []CreateKeyGroupRequestKeyHoldersInner, ) *CreateKeyGroupRequest`

NewCreateKeyGroupRequest instantiates a new CreateKeyGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeyGroupRequestWithDefaults

`func NewCreateKeyGroupRequestWithDefaults() *CreateKeyGroupRequest`

NewCreateKeyGroupRequestWithDefaults instantiates a new CreateKeyGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupType

`func (o *CreateKeyGroupRequest) GetGroupType() KeyGroupType`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *CreateKeyGroupRequest) GetGroupTypeOk() (*KeyGroupType, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *CreateKeyGroupRequest) SetGroupType(v KeyGroupType)`

SetGroupType sets GroupType field to given value.


### GetNodeCount

`func (o *CreateKeyGroupRequest) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *CreateKeyGroupRequest) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *CreateKeyGroupRequest) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetThreshold

`func (o *CreateKeyGroupRequest) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *CreateKeyGroupRequest) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *CreateKeyGroupRequest) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetKeyHolders

`func (o *CreateKeyGroupRequest) GetKeyHolders() []CreateKeyGroupRequestKeyHoldersInner`

GetKeyHolders returns the KeyHolders field if non-nil, zero value otherwise.

### GetKeyHoldersOk

`func (o *CreateKeyGroupRequest) GetKeyHoldersOk() (*[]CreateKeyGroupRequestKeyHoldersInner, bool)`

GetKeyHoldersOk returns a tuple with the KeyHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyHolders

`func (o *CreateKeyGroupRequest) SetKeyHolders(v []CreateKeyGroupRequestKeyHoldersInner)`

SetKeyHolders sets KeyHolders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


