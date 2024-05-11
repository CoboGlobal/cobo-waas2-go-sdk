# TSSGroupId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | the group id of the tss group. | [optional] 
**Curve** | Pointer to [**CurveType**](CurveType.md) |  | [optional] 

## Methods

### NewTSSGroupId

`func NewTSSGroupId() *TSSGroupId`

NewTSSGroupId instantiates a new TSSGroupId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSGroupIdWithDefaults

`func NewTSSGroupIdWithDefaults() *TSSGroupId`

NewTSSGroupIdWithDefaults instantiates a new TSSGroupId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *TSSGroupId) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *TSSGroupId) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *TSSGroupId) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *TSSGroupId) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetCurve

`func (o *TSSGroupId) GetCurve() CurveType`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *TSSGroupId) GetCurveOk() (*CurveType, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *TSSGroupId) SetCurve(v CurveType)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *TSSGroupId) HasCurve() bool`

HasCurve returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


