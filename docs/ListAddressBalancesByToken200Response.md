# ListAddressBalancesByToken200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AddressBalance**](AddressBalance.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListAddressBalancesByToken200Response

`func NewListAddressBalancesByToken200Response() *ListAddressBalancesByToken200Response`

NewListAddressBalancesByToken200Response instantiates a new ListAddressBalancesByToken200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAddressBalancesByToken200ResponseWithDefaults

`func NewListAddressBalancesByToken200ResponseWithDefaults() *ListAddressBalancesByToken200Response`

NewListAddressBalancesByToken200ResponseWithDefaults instantiates a new ListAddressBalancesByToken200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAddressBalancesByToken200Response) GetData() []AddressBalance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAddressBalancesByToken200Response) GetDataOk() (*[]AddressBalance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAddressBalancesByToken200Response) SetData(v []AddressBalance)`

SetData sets Data field to given value.

### HasData

`func (o *ListAddressBalancesByToken200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListAddressBalancesByToken200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListAddressBalancesByToken200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListAddressBalancesByToken200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListAddressBalancesByToken200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


