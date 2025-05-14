# TSSKeyShareSignDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | The TSS key share group ID. | [optional] 
**Message** | Pointer to **string** | The message to sign by key share. | [optional] 

## Methods

### NewTSSKeyShareSignDetail

`func NewTSSKeyShareSignDetail() *TSSKeyShareSignDetail`

NewTSSKeyShareSignDetail instantiates a new TSSKeyShareSignDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSKeyShareSignDetailWithDefaults

`func NewTSSKeyShareSignDetailWithDefaults() *TSSKeyShareSignDetail`

NewTSSKeyShareSignDetailWithDefaults instantiates a new TSSKeyShareSignDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *TSSKeyShareSignDetail) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *TSSKeyShareSignDetail) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *TSSKeyShareSignDetail) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *TSSKeyShareSignDetail) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetMessage

`func (o *TSSKeyShareSignDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TSSKeyShareSignDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TSSKeyShareSignDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TSSKeyShareSignDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


