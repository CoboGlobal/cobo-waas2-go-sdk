# CreateTssRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TSSRequestType**](TSSRequestType.md) |  | 
**TargetKeyGroupId** | **string** | The target key share group ID. | 
**SourceKeyGroupId** | Pointer to **string** | The used key share group ID.  **Note:** &#x60;used_key_group_id&#x60; is used only when the action &#x60;type&#x60; is either &#x60;KeyGenfromKeyGroup&#x60; or &#x60;Recovery&#x60;. This is to specify the key share group to be used as the source group to create a new &#x60;target_key_group&#x60;.  | [optional] 
**DetailParams** | Pointer to [**CreateTssRequestRequestDetailParams**](CreateTssRequestRequestDetailParams.md) |  | [optional] 

## Methods

### NewCreateTssRequestRequest

`func NewCreateTssRequestRequest(type_ TSSRequestType, targetKeyGroupId string, ) *CreateTssRequestRequest`

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


### GetTargetKeyGroupId

`func (o *CreateTssRequestRequest) GetTargetKeyGroupId() string`

GetTargetKeyGroupId returns the TargetKeyGroupId field if non-nil, zero value otherwise.

### GetTargetKeyGroupIdOk

`func (o *CreateTssRequestRequest) GetTargetKeyGroupIdOk() (*string, bool)`

GetTargetKeyGroupIdOk returns a tuple with the TargetKeyGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKeyGroupId

`func (o *CreateTssRequestRequest) SetTargetKeyGroupId(v string)`

SetTargetKeyGroupId sets TargetKeyGroupId field to given value.


### GetSourceKeyGroupId

`func (o *CreateTssRequestRequest) GetSourceKeyGroupId() string`

GetSourceKeyGroupId returns the SourceKeyGroupId field if non-nil, zero value otherwise.

### GetSourceKeyGroupIdOk

`func (o *CreateTssRequestRequest) GetSourceKeyGroupIdOk() (*string, bool)`

GetSourceKeyGroupIdOk returns a tuple with the SourceKeyGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKeyGroupId

`func (o *CreateTssRequestRequest) SetSourceKeyGroupId(v string)`

SetSourceKeyGroupId sets SourceKeyGroupId field to given value.

### HasSourceKeyGroupId

`func (o *CreateTssRequestRequest) HasSourceKeyGroupId() bool`

HasSourceKeyGroupId returns a boolean if a field has been set.

### GetDetailParams

`func (o *CreateTssRequestRequest) GetDetailParams() CreateTssRequestRequestDetailParams`

GetDetailParams returns the DetailParams field if non-nil, zero value otherwise.

### GetDetailParamsOk

`func (o *CreateTssRequestRequest) GetDetailParamsOk() (*CreateTssRequestRequestDetailParams, bool)`

GetDetailParamsOk returns a tuple with the DetailParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailParams

`func (o *CreateTssRequestRequest) SetDetailParams(v CreateTssRequestRequestDetailParams)`

SetDetailParams sets DetailParams field to given value.

### HasDetailParams

`func (o *CreateTssRequestRequest) HasDetailParams() bool`

HasDetailParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


