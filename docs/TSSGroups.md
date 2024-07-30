# TSSGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TssKeyShareGroupId** | Pointer to **string** | The TSS key share group ID. | [optional] 
**Curve** | Pointer to [**CurveType**](CurveType.md) |  | [optional] 

## Methods

### NewTSSGroups

`func NewTSSGroups() *TSSGroups`

NewTSSGroups instantiates a new TSSGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSGroupsWithDefaults

`func NewTSSGroupsWithDefaults() *TSSGroups`

NewTSSGroupsWithDefaults instantiates a new TSSGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTssKeyShareGroupId

`func (o *TSSGroups) GetTssKeyShareGroupId() string`

GetTssKeyShareGroupId returns the TssKeyShareGroupId field if non-nil, zero value otherwise.

### GetTssKeyShareGroupIdOk

`func (o *TSSGroups) GetTssKeyShareGroupIdOk() (*string, bool)`

GetTssKeyShareGroupIdOk returns a tuple with the TssKeyShareGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTssKeyShareGroupId

`func (o *TSSGroups) SetTssKeyShareGroupId(v string)`

SetTssKeyShareGroupId sets TssKeyShareGroupId field to given value.

### HasTssKeyShareGroupId

`func (o *TSSGroups) HasTssKeyShareGroupId() bool`

HasTssKeyShareGroupId returns a boolean if a field has been set.

### GetCurve

`func (o *TSSGroups) GetCurve() CurveType`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *TSSGroups) GetCurveOk() (*CurveType, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *TSSGroups) SetCurve(v CurveType)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *TSSGroups) HasCurve() bool`

HasCurve returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


