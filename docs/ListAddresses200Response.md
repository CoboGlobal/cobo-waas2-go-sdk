# ListAddresses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AddressInfo**](AddressInfo.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListAddresses200Response

`func NewListAddresses200Response() *ListAddresses200Response`

NewListAddresses200Response instantiates a new ListAddresses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAddresses200ResponseWithDefaults

`func NewListAddresses200ResponseWithDefaults() *ListAddresses200Response`

NewListAddresses200ResponseWithDefaults instantiates a new ListAddresses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAddresses200Response) GetData() []AddressInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAddresses200Response) GetDataOk() (*[]AddressInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAddresses200Response) SetData(v []AddressInfo)`

SetData sets Data field to given value.

### HasData

`func (o *ListAddresses200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListAddresses200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListAddresses200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListAddresses200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListAddresses200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


