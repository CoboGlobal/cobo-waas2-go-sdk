# ListTopUpPayers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ListTopUpPayers200ResponseDataInner**](ListTopUpPayers200ResponseDataInner.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListTopUpPayers200Response

`func NewListTopUpPayers200Response() *ListTopUpPayers200Response`

NewListTopUpPayers200Response instantiates a new ListTopUpPayers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTopUpPayers200ResponseWithDefaults

`func NewListTopUpPayers200ResponseWithDefaults() *ListTopUpPayers200Response`

NewListTopUpPayers200ResponseWithDefaults instantiates a new ListTopUpPayers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListTopUpPayers200Response) GetData() []ListTopUpPayers200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListTopUpPayers200Response) GetDataOk() (*[]ListTopUpPayers200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListTopUpPayers200Response) SetData(v []ListTopUpPayers200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ListTopUpPayers200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListTopUpPayers200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTopUpPayers200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTopUpPayers200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListTopUpPayers200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


