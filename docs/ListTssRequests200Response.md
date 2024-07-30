# ListTssRequests200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TSSRequest**](TSSRequest.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListTssRequests200Response

`func NewListTssRequests200Response() *ListTssRequests200Response`

NewListTssRequests200Response instantiates a new ListTssRequests200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTssRequests200ResponseWithDefaults

`func NewListTssRequests200ResponseWithDefaults() *ListTssRequests200Response`

NewListTssRequests200ResponseWithDefaults instantiates a new ListTssRequests200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListTssRequests200Response) GetData() []TSSRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListTssRequests200Response) GetDataOk() (*[]TSSRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListTssRequests200Response) SetData(v []TSSRequest)`

SetData sets Data field to given value.

### HasData

`func (o *ListTssRequests200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListTssRequests200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTssRequests200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTssRequests200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListTssRequests200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


