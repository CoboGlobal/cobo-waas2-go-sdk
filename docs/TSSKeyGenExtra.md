# TSSKeyGenExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to [**OrgInfo**](OrgInfo.md) |  | [optional] 
**Project** | Pointer to [**MPCProject**](MPCProject.md) |  | [optional] 
**Vault** | Pointer to [**MPCVault**](MPCVault.md) |  | [optional] 
**TargetKeyShareHolderGroup** | Pointer to [**KeyShareHolderGroup**](KeyShareHolderGroup.md) |  | [optional] 
**TssRequest** | Pointer to [**TSSRequest**](TSSRequest.md) |  | [optional] 

## Methods

### NewTSSKeyGenExtra

`func NewTSSKeyGenExtra() *TSSKeyGenExtra`

NewTSSKeyGenExtra instantiates a new TSSKeyGenExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSKeyGenExtraWithDefaults

`func NewTSSKeyGenExtraWithDefaults() *TSSKeyGenExtra`

NewTSSKeyGenExtraWithDefaults instantiates a new TSSKeyGenExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *TSSKeyGenExtra) GetOrg() OrgInfo`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *TSSKeyGenExtra) GetOrgOk() (*OrgInfo, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *TSSKeyGenExtra) SetOrg(v OrgInfo)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *TSSKeyGenExtra) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *TSSKeyGenExtra) GetProject() MPCProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TSSKeyGenExtra) GetProjectOk() (*MPCProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TSSKeyGenExtra) SetProject(v MPCProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *TSSKeyGenExtra) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetVault

`func (o *TSSKeyGenExtra) GetVault() MPCVault`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *TSSKeyGenExtra) GetVaultOk() (*MPCVault, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *TSSKeyGenExtra) SetVault(v MPCVault)`

SetVault sets Vault field to given value.

### HasVault

`func (o *TSSKeyGenExtra) HasVault() bool`

HasVault returns a boolean if a field has been set.

### GetTargetKeyShareHolderGroup

`func (o *TSSKeyGenExtra) GetTargetKeyShareHolderGroup() KeyShareHolderGroup`

GetTargetKeyShareHolderGroup returns the TargetKeyShareHolderGroup field if non-nil, zero value otherwise.

### GetTargetKeyShareHolderGroupOk

`func (o *TSSKeyGenExtra) GetTargetKeyShareHolderGroupOk() (*KeyShareHolderGroup, bool)`

GetTargetKeyShareHolderGroupOk returns a tuple with the TargetKeyShareHolderGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKeyShareHolderGroup

`func (o *TSSKeyGenExtra) SetTargetKeyShareHolderGroup(v KeyShareHolderGroup)`

SetTargetKeyShareHolderGroup sets TargetKeyShareHolderGroup field to given value.

### HasTargetKeyShareHolderGroup

`func (o *TSSKeyGenExtra) HasTargetKeyShareHolderGroup() bool`

HasTargetKeyShareHolderGroup returns a boolean if a field has been set.

### GetTssRequest

`func (o *TSSKeyGenExtra) GetTssRequest() TSSRequest`

GetTssRequest returns the TssRequest field if non-nil, zero value otherwise.

### GetTssRequestOk

`func (o *TSSKeyGenExtra) GetTssRequestOk() (*TSSRequest, bool)`

GetTssRequestOk returns a tuple with the TssRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTssRequest

`func (o *TSSKeyGenExtra) SetTssRequest(v TSSRequest)`

SetTssRequest sets TssRequest field to given value.

### HasTssRequest

`func (o *TSSKeyGenExtra) HasTssRequest() bool`

HasTssRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


