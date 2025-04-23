# ListSettlementRequests200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Settlement**](Settlement.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListSettlementRequests200Response

`func NewListSettlementRequests200Response() *ListSettlementRequests200Response`

NewListSettlementRequests200Response instantiates a new ListSettlementRequests200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSettlementRequests200ResponseWithDefaults

`func NewListSettlementRequests200ResponseWithDefaults() *ListSettlementRequests200Response`

NewListSettlementRequests200ResponseWithDefaults instantiates a new ListSettlementRequests200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListSettlementRequests200Response) GetData() []Settlement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListSettlementRequests200Response) GetDataOk() (*[]Settlement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListSettlementRequests200Response) SetData(v []Settlement)`

SetData sets Data field to given value.

### HasData

`func (o *ListSettlementRequests200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListSettlementRequests200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSettlementRequests200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSettlementRequests200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListSettlementRequests200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


