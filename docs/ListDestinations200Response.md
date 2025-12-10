# ListDestinations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Destination**](Destination.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListDestinations200Response

`func NewListDestinations200Response() *ListDestinations200Response`

NewListDestinations200Response instantiates a new ListDestinations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDestinations200ResponseWithDefaults

`func NewListDestinations200ResponseWithDefaults() *ListDestinations200Response`

NewListDestinations200ResponseWithDefaults instantiates a new ListDestinations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListDestinations200Response) GetData() []Destination`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDestinations200Response) GetDataOk() (*[]Destination, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDestinations200Response) SetData(v []Destination)`

SetData sets Data field to given value.

### HasData

`func (o *ListDestinations200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListDestinations200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListDestinations200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListDestinations200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListDestinations200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


