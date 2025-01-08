# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Before** | **string** | An object ID used to retrieve records before the specified object, indicating earlier or smaller records relative to the current dataset. You can use it to paginate backwards.   If empty, it means you have reached the start of the data.    Most API endpoints sort by object ID, but some use other fields depending on the endpoint.  | 
**After** | **string** | An object ID used to retrieve records after the specified object, indicating newer or larger records relative to the current dataset. You can use it to paginate forwards.    If empty, it means you have reached the end of the data.    Most API endpoints sort by object ID, but some use other fields depending on the endpoint.  | 
**TotalCount** | **int32** | The total number of records that match the query criteria, unaffected by the pagination parameters (&#x60;before&#x60; , &#x60;after&#x60;, and &#x60;limit&#x60;). | 

## Methods

### NewPagination

`func NewPagination(before string, after string, totalCount int32, ) *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBefore

`func (o *Pagination) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *Pagination) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *Pagination) SetBefore(v string)`

SetBefore sets Before field to given value.


### GetAfter

`func (o *Pagination) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *Pagination) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *Pagination) SetAfter(v string)`

SetAfter sets After field to given value.


### GetTotalCount

`func (o *Pagination) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Pagination) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Pagination) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


