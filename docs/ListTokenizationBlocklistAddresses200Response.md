# ListTokenizationBlocklistAddresses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TokenizationBlocklistAddressNote**](TokenizationBlocklistAddressNote.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewListTokenizationBlocklistAddresses200Response

`func NewListTokenizationBlocklistAddresses200Response(data []TokenizationBlocklistAddressNote, pagination Pagination, ) *ListTokenizationBlocklistAddresses200Response`

NewListTokenizationBlocklistAddresses200Response instantiates a new ListTokenizationBlocklistAddresses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTokenizationBlocklistAddresses200ResponseWithDefaults

`func NewListTokenizationBlocklistAddresses200ResponseWithDefaults() *ListTokenizationBlocklistAddresses200Response`

NewListTokenizationBlocklistAddresses200ResponseWithDefaults instantiates a new ListTokenizationBlocklistAddresses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListTokenizationBlocklistAddresses200Response) GetData() []TokenizationBlocklistAddressNote`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListTokenizationBlocklistAddresses200Response) GetDataOk() (*[]TokenizationBlocklistAddressNote, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListTokenizationBlocklistAddresses200Response) SetData(v []TokenizationBlocklistAddressNote)`

SetData sets Data field to given value.


### GetPagination

`func (o *ListTokenizationBlocklistAddresses200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTokenizationBlocklistAddresses200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTokenizationBlocklistAddresses200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


