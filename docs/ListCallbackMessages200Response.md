# ListCallbackMessages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CallbackMessage**](CallbackMessage.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewListCallbackMessages200Response

`func NewListCallbackMessages200Response(data []CallbackMessage, pagination Pagination, ) *ListCallbackMessages200Response`

NewListCallbackMessages200Response instantiates a new ListCallbackMessages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCallbackMessages200ResponseWithDefaults

`func NewListCallbackMessages200ResponseWithDefaults() *ListCallbackMessages200Response`

NewListCallbackMessages200ResponseWithDefaults instantiates a new ListCallbackMessages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCallbackMessages200Response) GetData() []CallbackMessage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCallbackMessages200Response) GetDataOk() (*[]CallbackMessage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCallbackMessages200Response) SetData(v []CallbackMessage)`

SetData sets Data field to given value.


### GetPagination

`func (o *ListCallbackMessages200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListCallbackMessages200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListCallbackMessages200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


