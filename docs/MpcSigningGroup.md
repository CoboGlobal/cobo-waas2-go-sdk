# MpcSigningGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedKeyShareHolderGroupId** | **string** | The ID of the Signing Group. | 
**UsedTssNodeIds** | Pointer to **[]string** | The ID of the TSS Nodes that are required to participate in the signature. | [optional] 

## Methods

### NewMpcSigningGroup

`func NewMpcSigningGroup(usedKeyShareHolderGroupId string, ) *MpcSigningGroup`

NewMpcSigningGroup instantiates a new MpcSigningGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpcSigningGroupWithDefaults

`func NewMpcSigningGroupWithDefaults() *MpcSigningGroup`

NewMpcSigningGroupWithDefaults instantiates a new MpcSigningGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsedKeyShareHolderGroupId

`func (o *MpcSigningGroup) GetUsedKeyShareHolderGroupId() string`

GetUsedKeyShareHolderGroupId returns the UsedKeyShareHolderGroupId field if non-nil, zero value otherwise.

### GetUsedKeyShareHolderGroupIdOk

`func (o *MpcSigningGroup) GetUsedKeyShareHolderGroupIdOk() (*string, bool)`

GetUsedKeyShareHolderGroupIdOk returns a tuple with the UsedKeyShareHolderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedKeyShareHolderGroupId

`func (o *MpcSigningGroup) SetUsedKeyShareHolderGroupId(v string)`

SetUsedKeyShareHolderGroupId sets UsedKeyShareHolderGroupId field to given value.


### GetUsedTssNodeIds

`func (o *MpcSigningGroup) GetUsedTssNodeIds() []string`

GetUsedTssNodeIds returns the UsedTssNodeIds field if non-nil, zero value otherwise.

### GetUsedTssNodeIdsOk

`func (o *MpcSigningGroup) GetUsedTssNodeIdsOk() (*[]string, bool)`

GetUsedTssNodeIdsOk returns a tuple with the UsedTssNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTssNodeIds

`func (o *MpcSigningGroup) SetUsedTssNodeIds(v []string)`

SetUsedTssNodeIds sets UsedTssNodeIds field to given value.

### HasUsedTssNodeIds

`func (o *MpcSigningGroup) HasUsedTssNodeIds() bool`

HasUsedTssNodeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


