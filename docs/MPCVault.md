# MPCVault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The vault ID. | [optional] 
**Name** | Pointer to **string** | The vault name. | [optional] 
**Type** | Pointer to [**MPCVaultType**](MPCVaultType.md) |  | [optional] 
**RootPubkeys** | Pointer to [**[]RootPubkey**](RootPubkey.md) |  | [optional] 
**CreateTimestamp** | Pointer to **float32** | The vault&#39;s creation time in Unix timestamp format, measured in milliseconds. | [optional] 

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

### GetId

`func (o *MPCVault) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MPCVault) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MPCVault) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MPCVault) HasId() bool`

HasId returns a boolean if a field has been set.

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

### GetCreateTimestamp

`func (o *MPCVault) GetCreateTimestamp() float32`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *MPCVault) GetCreateTimestampOk() (*float32, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *MPCVault) SetCreateTimestamp(v float32)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *MPCVault) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


