# TSSRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TssRequestId** | Pointer to **string** | The TSS request ID. | [optional] 
**SourceKeyShareHolderGroup** | Pointer to [**SourceGroup**](SourceGroup.md) |  | [optional] 
**TargetKeyShareHolderGroupId** | Pointer to **string** | The target key share holder group ID. | [optional] 
**Type** | Pointer to [**TSSRequestType**](TSSRequestType.md) |  | [optional] 
**Status** | Pointer to [**TSSRequestStatus**](TSSRequestStatus.md) |  | [optional] 

## Methods

### NewTSSRequest

`func NewTSSRequest() *TSSRequest`

NewTSSRequest instantiates a new TSSRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSRequestWithDefaults

`func NewTSSRequestWithDefaults() *TSSRequest`

NewTSSRequestWithDefaults instantiates a new TSSRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTssRequestId

`func (o *TSSRequest) GetTssRequestId() string`

GetTssRequestId returns the TssRequestId field if non-nil, zero value otherwise.

### GetTssRequestIdOk

`func (o *TSSRequest) GetTssRequestIdOk() (*string, bool)`

GetTssRequestIdOk returns a tuple with the TssRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTssRequestId

`func (o *TSSRequest) SetTssRequestId(v string)`

SetTssRequestId sets TssRequestId field to given value.

### HasTssRequestId

`func (o *TSSRequest) HasTssRequestId() bool`

HasTssRequestId returns a boolean if a field has been set.

### GetSourceKeyShareHolderGroup

`func (o *TSSRequest) GetSourceKeyShareHolderGroup() SourceGroup`

GetSourceKeyShareHolderGroup returns the SourceKeyShareHolderGroup field if non-nil, zero value otherwise.

### GetSourceKeyShareHolderGroupOk

`func (o *TSSRequest) GetSourceKeyShareHolderGroupOk() (*SourceGroup, bool)`

GetSourceKeyShareHolderGroupOk returns a tuple with the SourceKeyShareHolderGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKeyShareHolderGroup

`func (o *TSSRequest) SetSourceKeyShareHolderGroup(v SourceGroup)`

SetSourceKeyShareHolderGroup sets SourceKeyShareHolderGroup field to given value.

### HasSourceKeyShareHolderGroup

`func (o *TSSRequest) HasSourceKeyShareHolderGroup() bool`

HasSourceKeyShareHolderGroup returns a boolean if a field has been set.

### GetTargetKeyShareHolderGroupId

`func (o *TSSRequest) GetTargetKeyShareHolderGroupId() string`

GetTargetKeyShareHolderGroupId returns the TargetKeyShareHolderGroupId field if non-nil, zero value otherwise.

### GetTargetKeyShareHolderGroupIdOk

`func (o *TSSRequest) GetTargetKeyShareHolderGroupIdOk() (*string, bool)`

GetTargetKeyShareHolderGroupIdOk returns a tuple with the TargetKeyShareHolderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKeyShareHolderGroupId

`func (o *TSSRequest) SetTargetKeyShareHolderGroupId(v string)`

SetTargetKeyShareHolderGroupId sets TargetKeyShareHolderGroupId field to given value.

### HasTargetKeyShareHolderGroupId

`func (o *TSSRequest) HasTargetKeyShareHolderGroupId() bool`

HasTargetKeyShareHolderGroupId returns a boolean if a field has been set.

### GetType

`func (o *TSSRequest) GetType() TSSRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TSSRequest) GetTypeOk() (*TSSRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TSSRequest) SetType(v TSSRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *TSSRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *TSSRequest) GetStatus() TSSRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TSSRequest) GetStatusOk() (*TSSRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TSSRequest) SetStatus(v TSSRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TSSRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


