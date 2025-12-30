# ListKyaScreenings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]KyaScreeningResult**](KyaScreeningResult.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewListKyaScreenings200Response

`func NewListKyaScreenings200Response(data []KyaScreeningResult, pagination Pagination, ) *ListKyaScreenings200Response`

NewListKyaScreenings200Response instantiates a new ListKyaScreenings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKyaScreenings200ResponseWithDefaults

`func NewListKyaScreenings200ResponseWithDefaults() *ListKyaScreenings200Response`

NewListKyaScreenings200ResponseWithDefaults instantiates a new ListKyaScreenings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListKyaScreenings200Response) GetData() []KyaScreeningResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListKyaScreenings200Response) GetDataOk() (*[]KyaScreeningResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListKyaScreenings200Response) SetData(v []KyaScreeningResult)`

SetData sets Data field to given value.


### GetPagination

`func (o *ListKyaScreenings200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListKyaScreenings200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListKyaScreenings200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


