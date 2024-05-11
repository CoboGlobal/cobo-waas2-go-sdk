# KeyHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**KeyHolderType**](KeyHolderType.md) |  | [optional] 
**TssNodeId** | Pointer to **string** | The IDs of the tss node. | [optional] 
**Online** | Pointer to **bool** | Indicates if the tss node online | [optional] 
**Status** | Pointer to [**KeyHolderStatus**](KeyHolderStatus.md) |  | [optional] 

## Methods

### NewKeyHolder

`func NewKeyHolder() *KeyHolder`

NewKeyHolder instantiates a new KeyHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyHolderWithDefaults

`func NewKeyHolderWithDefaults() *KeyHolder`

NewKeyHolderWithDefaults instantiates a new KeyHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyHolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyHolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyHolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyHolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *KeyHolder) GetType() KeyHolderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyHolder) GetTypeOk() (*KeyHolderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyHolder) SetType(v KeyHolderType)`

SetType sets Type field to given value.

### HasType

`func (o *KeyHolder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTssNodeId

`func (o *KeyHolder) GetTssNodeId() string`

GetTssNodeId returns the TssNodeId field if non-nil, zero value otherwise.

### GetTssNodeIdOk

`func (o *KeyHolder) GetTssNodeIdOk() (*string, bool)`

GetTssNodeIdOk returns a tuple with the TssNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTssNodeId

`func (o *KeyHolder) SetTssNodeId(v string)`

SetTssNodeId sets TssNodeId field to given value.

### HasTssNodeId

`func (o *KeyHolder) HasTssNodeId() bool`

HasTssNodeId returns a boolean if a field has been set.

### GetOnline

`func (o *KeyHolder) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *KeyHolder) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *KeyHolder) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *KeyHolder) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetStatus

`func (o *KeyHolder) GetStatus() KeyHolderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyHolder) GetStatusOk() (*KeyHolderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyHolder) SetStatus(v KeyHolderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeyHolder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


