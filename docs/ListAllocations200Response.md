# ListAllocations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AllocationRecord**](AllocationRecord.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListAllocations200Response

`func NewListAllocations200Response() *ListAllocations200Response`

NewListAllocations200Response instantiates a new ListAllocations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllocations200ResponseWithDefaults

`func NewListAllocations200ResponseWithDefaults() *ListAllocations200Response`

NewListAllocations200ResponseWithDefaults instantiates a new ListAllocations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAllocations200Response) GetData() []AllocationRecord`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAllocations200Response) GetDataOk() (*[]AllocationRecord, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAllocations200Response) SetData(v []AllocationRecord)`

SetData sets Data field to given value.

### HasData

`func (o *ListAllocations200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListAllocations200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListAllocations200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListAllocations200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListAllocations200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


