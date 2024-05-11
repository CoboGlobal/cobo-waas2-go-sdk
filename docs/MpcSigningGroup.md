# MpcSigningGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedKeyGroupId** | Pointer to **string** | Unique id of the using key group. | [optional] 
**UsedNodeIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMpcSigningGroup

`func NewMpcSigningGroup() *MpcSigningGroup`

NewMpcSigningGroup instantiates a new MpcSigningGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpcSigningGroupWithDefaults

`func NewMpcSigningGroupWithDefaults() *MpcSigningGroup`

NewMpcSigningGroupWithDefaults instantiates a new MpcSigningGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsedKeyGroupId

`func (o *MpcSigningGroup) GetUsedKeyGroupId() string`

GetUsedKeyGroupId returns the UsedKeyGroupId field if non-nil, zero value otherwise.

### GetUsedKeyGroupIdOk

`func (o *MpcSigningGroup) GetUsedKeyGroupIdOk() (*string, bool)`

GetUsedKeyGroupIdOk returns a tuple with the UsedKeyGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKeyGroupId

`func (o *MpcSigningGroup) SetUsedKeyGroupId(v string)`

SetUsedKeyGroupId sets UsedKeyGroupId field to given value.

### HasUsedKeyGroupId

`func (o *MpcSigningGroup) HasUsedKeyGroupId() bool`

HasUsedKeyGroupId returns a boolean if a field has been set.

### GetUsedNodeIds

`func (o *MpcSigningGroup) GetUsedNodeIds() []string`

GetUsedNodeIds returns the UsedNodeIds field if non-nil, zero value otherwise.

### GetUsedNodeIdsOk

`func (o *MpcSigningGroup) GetUsedNodeIdsOk() (*[]string, bool)`

GetUsedNodeIdsOk returns a tuple with the UsedNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedNodeIds

`func (o *MpcSigningGroup) SetUsedNodeIds(v []string)`

SetUsedNodeIds sets UsedNodeIds field to given value.

### HasUsedNodeIds

`func (o *MpcSigningGroup) HasUsedNodeIds() bool`

HasUsedNodeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


