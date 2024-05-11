# CreateKeyGroupRequestKeyHoldersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**KeyHolderType**](KeyHolderType.md) |  | [optional] 
**TssNodeId** | Pointer to **string** | The ID of the tss node. | [optional] 

## Methods

### NewCreateKeyGroupRequestKeyHoldersInner

`func NewCreateKeyGroupRequestKeyHoldersInner() *CreateKeyGroupRequestKeyHoldersInner`

NewCreateKeyGroupRequestKeyHoldersInner instantiates a new CreateKeyGroupRequestKeyHoldersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeyGroupRequestKeyHoldersInnerWithDefaults

`func NewCreateKeyGroupRequestKeyHoldersInnerWithDefaults() *CreateKeyGroupRequestKeyHoldersInner`

NewCreateKeyGroupRequestKeyHoldersInnerWithDefaults instantiates a new CreateKeyGroupRequestKeyHoldersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateKeyGroupRequestKeyHoldersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateKeyGroupRequestKeyHoldersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateKeyGroupRequestKeyHoldersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateKeyGroupRequestKeyHoldersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CreateKeyGroupRequestKeyHoldersInner) GetType() KeyHolderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateKeyGroupRequestKeyHoldersInner) GetTypeOk() (*KeyHolderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateKeyGroupRequestKeyHoldersInner) SetType(v KeyHolderType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateKeyGroupRequestKeyHoldersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTssNodeId

`func (o *CreateKeyGroupRequestKeyHoldersInner) GetTssNodeId() string`

GetTssNodeId returns the TssNodeId field if non-nil, zero value otherwise.

### GetTssNodeIdOk

`func (o *CreateKeyGroupRequestKeyHoldersInner) GetTssNodeIdOk() (*string, bool)`

GetTssNodeIdOk returns a tuple with the TssNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTssNodeId

`func (o *CreateKeyGroupRequestKeyHoldersInner) SetTssNodeId(v string)`

SetTssNodeId sets TssNodeId field to given value.

### HasTssNodeId

`func (o *CreateKeyGroupRequestKeyHoldersInner) HasTssNodeId() bool`

HasTssNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


