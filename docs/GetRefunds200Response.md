# GetRefunds200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Refund**](Refund.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGetRefunds200Response

`func NewGetRefunds200Response() *GetRefunds200Response`

NewGetRefunds200Response instantiates a new GetRefunds200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRefunds200ResponseWithDefaults

`func NewGetRefunds200ResponseWithDefaults() *GetRefunds200Response`

NewGetRefunds200ResponseWithDefaults instantiates a new GetRefunds200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRefunds200Response) GetData() []Refund`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRefunds200Response) GetDataOk() (*[]Refund, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRefunds200Response) SetData(v []Refund)`

SetData sets Data field to given value.

### HasData

`func (o *GetRefunds200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *GetRefunds200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetRefunds200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetRefunds200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetRefunds200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


