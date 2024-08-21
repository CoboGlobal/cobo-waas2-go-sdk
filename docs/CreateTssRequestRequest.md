# CreateTssRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TSSRequestType**](TSSRequestType.md) |  | 
**TargetKeyShareHolderGroupId** | **string** | The target key share holder group ID. | 
**SourceKeyShareHolderGroup** | Pointer to [**SourceGroup**](SourceGroup.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the TSS request. | [optional] 

## Methods

### NewCreateTssRequestRequest

`func NewCreateTssRequestRequest(type_ TSSRequestType, targetKeyShareHolderGroupId string, ) *CreateTssRequestRequest`

NewCreateTssRequestRequest instantiates a new CreateTssRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTssRequestRequestWithDefaults

`func NewCreateTssRequestRequestWithDefaults() *CreateTssRequestRequest`

NewCreateTssRequestRequestWithDefaults instantiates a new CreateTssRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateTssRequestRequest) GetType() TSSRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTssRequestRequest) GetTypeOk() (*TSSRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTssRequestRequest) SetType(v TSSRequestType)`

SetType sets Type field to given value.


### GetTargetKeyShareHolderGroupId

`func (o *CreateTssRequestRequest) GetTargetKeyShareHolderGroupId() string`

GetTargetKeyShareHolderGroupId returns the TargetKeyShareHolderGroupId field if non-nil, zero value otherwise.

### GetTargetKeyShareHolderGroupIdOk

`func (o *CreateTssRequestRequest) GetTargetKeyShareHolderGroupIdOk() (*string, bool)`

GetTargetKeyShareHolderGroupIdOk returns a tuple with the TargetKeyShareHolderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKeyShareHolderGroupId

`func (o *CreateTssRequestRequest) SetTargetKeyShareHolderGroupId(v string)`

SetTargetKeyShareHolderGroupId sets TargetKeyShareHolderGroupId field to given value.


### GetSourceKeyShareHolderGroup

`func (o *CreateTssRequestRequest) GetSourceKeyShareHolderGroup() SourceGroup`

GetSourceKeyShareHolderGroup returns the SourceKeyShareHolderGroup field if non-nil, zero value otherwise.

### GetSourceKeyShareHolderGroupOk

`func (o *CreateTssRequestRequest) GetSourceKeyShareHolderGroupOk() (*SourceGroup, bool)`

GetSourceKeyShareHolderGroupOk returns a tuple with the SourceKeyShareHolderGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKeyShareHolderGroup

`func (o *CreateTssRequestRequest) SetSourceKeyShareHolderGroup(v SourceGroup)`

SetSourceKeyShareHolderGroup sets SourceKeyShareHolderGroup field to given value.

### HasSourceKeyShareHolderGroup

`func (o *CreateTssRequestRequest) HasSourceKeyShareHolderGroup() bool`

HasSourceKeyShareHolderGroup returns a boolean if a field has been set.

### GetDescription

`func (o *CreateTssRequestRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTssRequestRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTssRequestRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTssRequestRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


