# ListApprovalRequests200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ApprovalRequest**](ApprovalRequest.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListApprovalRequests200Response

`func NewListApprovalRequests200Response() *ListApprovalRequests200Response`

NewListApprovalRequests200Response instantiates a new ListApprovalRequests200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApprovalRequests200ResponseWithDefaults

`func NewListApprovalRequests200ResponseWithDefaults() *ListApprovalRequests200Response`

NewListApprovalRequests200ResponseWithDefaults instantiates a new ListApprovalRequests200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListApprovalRequests200Response) GetData() []ApprovalRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListApprovalRequests200Response) GetDataOk() (*[]ApprovalRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListApprovalRequests200Response) SetData(v []ApprovalRequest)`

SetData sets Data field to given value.

### HasData

`func (o *ListApprovalRequests200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListApprovalRequests200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListApprovalRequests200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListApprovalRequests200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListApprovalRequests200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


