# OrgInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | The organization ID. | 
**BizOrgId** | Pointer to **int32** | An internal business ID assigned by Cobo. Used mainly by Cobo&#39;s customer support to locate the organization. | [optional] 
**Name** | Pointer to **string** | The organization name. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The organization&#39;s creation time in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewOrgInfo

`func NewOrgInfo(orgId string, ) *OrgInfo`

NewOrgInfo instantiates a new OrgInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgInfoWithDefaults

`func NewOrgInfoWithDefaults() *OrgInfo`

NewOrgInfoWithDefaults instantiates a new OrgInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *OrgInfo) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgInfo) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgInfo) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetBizOrgId

`func (o *OrgInfo) GetBizOrgId() int32`

GetBizOrgId returns the BizOrgId field if non-nil, zero value otherwise.

### GetBizOrgIdOk

`func (o *OrgInfo) GetBizOrgIdOk() (*int32, bool)`

GetBizOrgIdOk returns a tuple with the BizOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizOrgId

`func (o *OrgInfo) SetBizOrgId(v int32)`

SetBizOrgId sets BizOrgId field to given value.

### HasBizOrgId

`func (o *OrgInfo) HasBizOrgId() bool`

HasBizOrgId returns a boolean if a field has been set.

### GetName

`func (o *OrgInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrgInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *OrgInfo) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *OrgInfo) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *OrgInfo) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *OrgInfo) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


