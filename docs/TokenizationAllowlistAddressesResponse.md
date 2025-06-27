# TokenizationAllowlistAddressesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TokenizationAllowlistAddressNote**](TokenizationAllowlistAddressNote.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewTokenizationAllowlistAddressesResponse

`func NewTokenizationAllowlistAddressesResponse(data []TokenizationAllowlistAddressNote, pagination Pagination, ) *TokenizationAllowlistAddressesResponse`

NewTokenizationAllowlistAddressesResponse instantiates a new TokenizationAllowlistAddressesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationAllowlistAddressesResponseWithDefaults

`func NewTokenizationAllowlistAddressesResponseWithDefaults() *TokenizationAllowlistAddressesResponse`

NewTokenizationAllowlistAddressesResponseWithDefaults instantiates a new TokenizationAllowlistAddressesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TokenizationAllowlistAddressesResponse) GetData() []TokenizationAllowlistAddressNote`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenizationAllowlistAddressesResponse) GetDataOk() (*[]TokenizationAllowlistAddressNote, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenizationAllowlistAddressesResponse) SetData(v []TokenizationAllowlistAddressNote)`

SetData sets Data field to given value.


### GetPagination

`func (o *TokenizationAllowlistAddressesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TokenizationAllowlistAddressesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TokenizationAllowlistAddressesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


