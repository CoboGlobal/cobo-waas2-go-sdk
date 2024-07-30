# ListMpcProjects200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MPCProject**](MPCProject.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListMpcProjects200Response

`func NewListMpcProjects200Response() *ListMpcProjects200Response`

NewListMpcProjects200Response instantiates a new ListMpcProjects200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMpcProjects200ResponseWithDefaults

`func NewListMpcProjects200ResponseWithDefaults() *ListMpcProjects200Response`

NewListMpcProjects200ResponseWithDefaults instantiates a new ListMpcProjects200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListMpcProjects200Response) GetData() []MPCProject`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListMpcProjects200Response) GetDataOk() (*[]MPCProject, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListMpcProjects200Response) SetData(v []MPCProject)`

SetData sets Data field to given value.

### HasData

`func (o *ListMpcProjects200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListMpcProjects200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListMpcProjects200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListMpcProjects200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListMpcProjects200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


