# CreateKeyShareHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Key share holder&#39;s name. | [optional] 
**Type** | Pointer to [**KeyShareHolderType**](KeyShareHolderType.md) |  | [optional] 
**TssNodeId** | Pointer to **string** | The TSS Node ID. | [optional] 

## Methods

### NewCreateKeyShareHolder

`func NewCreateKeyShareHolder() *CreateKeyShareHolder`

NewCreateKeyShareHolder instantiates a new CreateKeyShareHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeyShareHolderWithDefaults

`func NewCreateKeyShareHolderWithDefaults() *CreateKeyShareHolder`

NewCreateKeyShareHolderWithDefaults instantiates a new CreateKeyShareHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateKeyShareHolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateKeyShareHolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateKeyShareHolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateKeyShareHolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CreateKeyShareHolder) GetType() KeyShareHolderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateKeyShareHolder) GetTypeOk() (*KeyShareHolderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateKeyShareHolder) SetType(v KeyShareHolderType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateKeyShareHolder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTssNodeId

`func (o *CreateKeyShareHolder) GetTssNodeId() string`

GetTssNodeId returns the TssNodeId field if non-nil, zero value otherwise.

### GetTssNodeIdOk

`func (o *CreateKeyShareHolder) GetTssNodeIdOk() (*string, bool)`

GetTssNodeIdOk returns a tuple with the TssNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTssNodeId

`func (o *CreateKeyShareHolder) SetTssNodeId(v string)`

SetTssNodeId sets TssNodeId field to given value.

### HasTssNodeId

`func (o *CreateKeyShareHolder) HasTssNodeId() bool`

HasTssNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


