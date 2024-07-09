# CreateMpcVaultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** | The project ID.  **Notes:** 1. If &#x60;vault_type&#x60; is &#x60;OrgControlled&#x60;, the value of &#x60;project_id&#x60; will be ignored. 2. If &#x60;vault_type&#x60; is &#x60;UserControlled&#x60;, then &#x60;project_id&#x60; is required.  | [optional] 
**Name** | **string** | The name of the new vault. | 
**VaultType** | [**MPCVaultType**](MPCVaultType.md) |  | 

## Methods

### NewCreateMpcVaultRequest

`func NewCreateMpcVaultRequest(name string, vaultType MPCVaultType, ) *CreateMpcVaultRequest`

NewCreateMpcVaultRequest instantiates a new CreateMpcVaultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMpcVaultRequestWithDefaults

`func NewCreateMpcVaultRequestWithDefaults() *CreateMpcVaultRequest`

NewCreateMpcVaultRequestWithDefaults instantiates a new CreateMpcVaultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *CreateMpcVaultRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateMpcVaultRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateMpcVaultRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateMpcVaultRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetName

`func (o *CreateMpcVaultRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMpcVaultRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMpcVaultRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVaultType

`func (o *CreateMpcVaultRequest) GetVaultType() MPCVaultType`

GetVaultType returns the VaultType field if non-nil, zero value otherwise.

### GetVaultTypeOk

`func (o *CreateMpcVaultRequest) GetVaultTypeOk() (*MPCVaultType, bool)`

GetVaultTypeOk returns a tuple with the VaultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultType

`func (o *CreateMpcVaultRequest) SetVaultType(v MPCVaultType)`

SetVaultType sets VaultType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


