# ListSettlementDetails200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SettlementDetail**](SettlementDetail.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListSettlementDetails200Response

`func NewListSettlementDetails200Response() *ListSettlementDetails200Response`

NewListSettlementDetails200Response instantiates a new ListSettlementDetails200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSettlementDetails200ResponseWithDefaults

`func NewListSettlementDetails200ResponseWithDefaults() *ListSettlementDetails200Response`

NewListSettlementDetails200ResponseWithDefaults instantiates a new ListSettlementDetails200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListSettlementDetails200Response) GetData() []SettlementDetail`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListSettlementDetails200Response) GetDataOk() (*[]SettlementDetail, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListSettlementDetails200Response) SetData(v []SettlementDetail)`

SetData sets Data field to given value.

### HasData

`func (o *ListSettlementDetails200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListSettlementDetails200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSettlementDetails200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSettlementDetails200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListSettlementDetails200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


