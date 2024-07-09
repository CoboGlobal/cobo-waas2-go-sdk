# GetExchangeSupportedAssets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AssetInfo**](AssetInfo.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGetExchangeSupportedAssets200Response

`func NewGetExchangeSupportedAssets200Response() *GetExchangeSupportedAssets200Response`

NewGetExchangeSupportedAssets200Response instantiates a new GetExchangeSupportedAssets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExchangeSupportedAssets200ResponseWithDefaults

`func NewGetExchangeSupportedAssets200ResponseWithDefaults() *GetExchangeSupportedAssets200Response`

NewGetExchangeSupportedAssets200ResponseWithDefaults instantiates a new GetExchangeSupportedAssets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetExchangeSupportedAssets200Response) GetData() []AssetInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetExchangeSupportedAssets200Response) GetDataOk() (*[]AssetInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetExchangeSupportedAssets200Response) SetData(v []AssetInfo)`

SetData sets Data field to given value.

### HasData

`func (o *GetExchangeSupportedAssets200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *GetExchangeSupportedAssets200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetExchangeSupportedAssets200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetExchangeSupportedAssets200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetExchangeSupportedAssets200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


