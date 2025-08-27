# ListAutoSweepTask200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AutoSweepTask**](AutoSweepTask.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListAutoSweepTask200Response

`func NewListAutoSweepTask200Response() *ListAutoSweepTask200Response`

NewListAutoSweepTask200Response instantiates a new ListAutoSweepTask200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAutoSweepTask200ResponseWithDefaults

`func NewListAutoSweepTask200ResponseWithDefaults() *ListAutoSweepTask200Response`

NewListAutoSweepTask200ResponseWithDefaults instantiates a new ListAutoSweepTask200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAutoSweepTask200Response) GetData() []AutoSweepTask`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAutoSweepTask200Response) GetDataOk() (*[]AutoSweepTask, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAutoSweepTask200Response) SetData(v []AutoSweepTask)`

SetData sets Data field to given value.

### HasData

`func (o *ListAutoSweepTask200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListAutoSweepTask200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListAutoSweepTask200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListAutoSweepTask200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListAutoSweepTask200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


