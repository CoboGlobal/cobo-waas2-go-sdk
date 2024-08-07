# MPCVault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaultId** | Pointer to **string** | The vault ID. | [optional] 
**ProjectId** | Pointer to **string** | The project ID. | [optional] 
**Name** | Pointer to **string** | The vault name. | [optional] 
**Type** | Pointer to [**MPCVaultType**](MPCVaultType.md) |  | [optional] 
**RootPubkeys** | Pointer to [**[]RootPubkey**](RootPubkey.md) |  | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The vault&#39;s creation time in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewMPCVault

`func NewMPCVault() *MPCVault`

NewMPCVault instantiates a new MPCVault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMPCVaultWithDefaults

`func NewMPCVaultWithDefaults() *MPCVault`

NewMPCVaultWithDefaults instantiates a new MPCVault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVaultId

`func (o *MPCVault) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *MPCVault) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *MPCVault) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *MPCVault) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### GetProjectId

`func (o *MPCVault) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MPCVault) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MPCVault) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *MPCVault) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetName

`func (o *MPCVault) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MPCVault) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MPCVault) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MPCVault) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *MPCVault) GetType() MPCVaultType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MPCVault) GetTypeOk() (*MPCVaultType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MPCVault) SetType(v MPCVaultType)`

SetType sets Type field to given value.

### HasType

`func (o *MPCVault) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRootPubkeys

`func (o *MPCVault) GetRootPubkeys() []RootPubkey`

GetRootPubkeys returns the RootPubkeys field if non-nil, zero value otherwise.

### GetRootPubkeysOk

`func (o *MPCVault) GetRootPubkeysOk() (*[]RootPubkey, bool)`

GetRootPubkeysOk returns a tuple with the RootPubkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPubkeys

`func (o *MPCVault) SetRootPubkeys(v []RootPubkey)`

SetRootPubkeys sets RootPubkeys field to given value.

### HasRootPubkeys

`func (o *MPCVault) HasRootPubkeys() bool`

HasRootPubkeys returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *MPCVault) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *MPCVault) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *MPCVault) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *MPCVault) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


